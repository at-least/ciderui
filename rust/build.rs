//! With the `css` feature, the kit's stylesheets as `include_str!` items:
//! every root stylesheet and every component's, listed by their paths in the
//! kit, so `ciderui::css::write` can lay them out for a site's Tailwind
//! build. The paths are absolute, from the crate's directory: a git checkout
//! of the repository has the kit beside `rust/`.
use std::env;
use std::fs;
use std::path::{Path, PathBuf};

fn css_files(dir: &Path, root: &Path, out: &mut Vec<(String, PathBuf)>) {
    let mut entries: Vec<_> = fs::read_dir(dir)
        .unwrap_or_else(|e| panic!("{}: {e}", dir.display()))
        .map(|e| e.unwrap().path())
        .collect();
    entries.sort();
    for path in entries {
        if path.is_dir() {
            css_files(&path, root, out);
        } else if path.extension().is_some_and(|x| x == "css") {
            let rel = path
                .strip_prefix(root)
                .unwrap()
                .to_string_lossy()
                .replace('\\', "/");
            out.push((rel, path));
        }
    }
}

fn main() {
    let manifest = PathBuf::from(env::var("CARGO_MANIFEST_DIR").unwrap());
    let root = manifest.parent().unwrap();
    println!("cargo:rerun-if-changed=build.rs");
    let mut files = Vec::new();
    for name in [
        "index.css",
        "theme.css",
        "tokens.css",
        "base.css",
        "type.css",
        "fonts.css",
        "site.css",
        "console.css",
        "docs.css",
    ] {
        files.push((name.to_string(), root.join(name)));
    }
    css_files(&root.join("components"), root, &mut files);
    println!(
        "cargo:rerun-if-changed={}",
        root.join("components").display()
    );
    let out_dir = PathBuf::from(env::var("OUT_DIR").unwrap());
    let texts: Vec<String> = files
        .iter()
        .map(|(_, path)| {
            fs::read_to_string(path).unwrap_or_else(|e| panic!("{}: {e}", path.display()))
        })
        .collect();
    fs::write(out_dir.join("words.rs"), words(&texts)).unwrap();
    for (_, path) in &files {
        println!("cargo:rerun-if-changed={}", path.display());
    }
    if env::var_os("CARGO_FEATURE_CSS").is_none() {
        return;
    }
    let mut out = String::from(
        "/// Every stylesheet of the kit: its path in the kit, and its text.\npub const FILES: &[(&str, &str)] = &[\n",
    );
    for (rel, path) in &files {
        out.push_str(&format!(
            "    ({rel:?}, include_str!({:?})),\n",
            path.display()
        ));
    }
    out.push_str("];\n");
    fs::write(out_dir.join("css_files.rs"), out).unwrap();
}

/// The words of ciderui's markup the kit's stylesheets know, for
/// `check_markup`: `CLASSES`, every `cui-*` name they write anywhere (a
/// rule's, a utility's, a header's, as a part with no rule of its own,
/// `cui-post-content`, is named in its header); `VALUES`, each
/// `data-variant`, `data-size` and `data-color` value a rule selects, with
/// the class it is written beside (the last `.cui-*` before it in the
/// selector, nesting resolved).
fn words(texts: &[String]) -> String {
    let mut classes = Vec::new();
    let mut values = Vec::new();
    for text in texts {
        let mut at = 0;
        while let Some(i) = text[at..].find("cui-").map(|i| at + i) {
            at = i + 4;
            if text[..i].ends_with(|c: char| c.is_ascii_alphanumeric() || c == '-') {
                continue;
            }
            classes.push(name(&text[i..]).to_string());
        }
        for selector in selectors(text) {
            let mut at = 0;
            while let Some(i) = selector[at..].find("[data-").map(|i| at + i) {
                at = i + 1;
                let rest = &selector[i + 1..];
                let Some((attribute, value)) = rest
                    .split_once(']')
                    .map(|(a, _)| a)
                    .and_then(|a| a.split_once('='))
                else {
                    continue;
                };
                if !["data-variant", "data-size", "data-color"].contains(&attribute) {
                    continue;
                }
                let value = value.trim_matches('"');
                let Some(class) = selector[..i]
                    .rfind(".cui-")
                    .map(|c| name(&selector[c + 1..]))
                else {
                    panic!("{selector}: {attribute}={value} with no cui-* class before it");
                };
                values.push((class.to_string(), attribute.to_string(), value.to_string()));
            }
        }
    }
    classes.sort();
    classes.dedup();
    values.sort();
    values.dedup();
    let mut out = String::from("pub(crate) const CLASSES: &[&str] = &[\n");
    for class in &classes {
        out.push_str(&format!("    {class:?},\n"));
    }
    out.push_str("];\npub(crate) const VALUES: &[(&str, &str, &str)] = &[\n");
    for (class, attribute, value) in &values {
        out.push_str(&format!("    ({class:?}, {attribute:?}, {value:?}),\n"));
    }
    out.push_str("];\n");
    out
}

/// The class name `text` begins with.
fn name(text: &str) -> &str {
    let end = text
        .find(|c: char| !(c.is_ascii_alphanumeric() || c == '-'))
        .unwrap_or(text.len());
    &text[..end]
}

/// Every style rule's selector in `css`, nesting resolved (`&` the parent,
/// or the parent and a space before it); at-rules see-through.
fn selectors(css: &str) -> Vec<String> {
    let mut text = String::new();
    let mut rest = css;
    while let Some(i) = rest.find("/*") {
        text.push_str(&rest[..i]);
        rest = &rest[i + rest[i..].find("*/").expect("an unclosed comment") + 2..];
    }
    text.push_str(rest);
    let mut out = Vec::new();
    let mut stack: Vec<Vec<String>> = vec![Vec::new()];
    let mut head = String::new();
    for c in text.chars() {
        match c {
            '{' => {
                let parents = stack.last().unwrap().clone();
                let prelude = head.trim().to_string();
                head.clear();
                if prelude.starts_with('@') {
                    stack.push(parents);
                    continue;
                }
                let mut full = Vec::new();
                for part in split(&prelude) {
                    if parents.is_empty() {
                        full.push(part.to_string());
                    } else {
                        for parent in &parents {
                            full.push(if part.contains('&') {
                                part.replace('&', parent)
                            } else {
                                format!("{parent} {part}")
                            });
                        }
                    }
                }
                out.extend(full.iter().cloned());
                stack.push(full);
            }
            '}' => {
                stack.pop();
                head.clear();
            }
            ';' => head.clear(),
            _ => head.push(c),
        }
    }
    out
}

/// `selector`'s list split at its top-level commas.
fn split(selector: &str) -> Vec<&str> {
    let mut parts = Vec::new();
    let mut depth = 0;
    let mut start = 0;
    for (i, c) in selector.char_indices() {
        match c {
            '(' | '[' => depth += 1,
            ')' | ']' => depth -= 1,
            ',' if depth == 0 => {
                parts.push(selector[start..i].trim());
                start = i + 1;
            }
            _ => {}
        }
    }
    parts.push(selector[start..].trim());
    parts
}
