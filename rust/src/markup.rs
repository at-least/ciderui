//! [`check_markup`]: the words of ciderui's that a page writes, held to the
//! kit's stylesheets.

include!(concat!(env!("OUT_DIR"), "/words.rs"));

/// Each word of ciderui's in `html` that the kit's stylesheets do not know,
/// as a line saying what and where; `Ok` when there is none. Checked: every
/// `cui-*` class (a misspelt one styles nothing), and on an element with a
/// `cui-*` class, every `data-variant`, `data-size` and `data-color` value
/// (a value no rule selects renders the default, silently). For a site's
/// tests, on each page it renders:
///
/// ```
/// ciderui::check_markup(r#"<a class="cui-console-btn" data-variant="plain" href="/">Open</a>"#).unwrap();
/// let err = ciderui::check_markup(r#"<a class="cui-console-btn" data-variant="plian" href="/">Open</a>"#).unwrap_err();
/// assert_eq!(err, ["data-variant=\"plian\" on <a class=\"cui-console-btn\">: no rule of ciderui's selects it (cui-console-btn's: filled, plain)"]);
/// ```
pub fn check_markup(html: &str) -> Result<(), Vec<String>> {
    let mut problems = Vec::new();
    for (tag, attributes) in start_tags(html) {
        let class = attributes
            .iter()
            .find(|(n, _)| *n == "class")
            .map_or("", |(_, v)| *v);
        let ours: Vec<&str> = class
            .split_whitespace()
            .filter(|c| c.starts_with("cui-"))
            .collect();
        for c in &ours {
            if CLASSES.binary_search(c).is_err() {
                problems.push(format!(
                    "class {c:?} on <{tag} class={class:?}>: ciderui has no such class"
                ));
            }
        }
        if ours.is_empty() {
            continue;
        }
        for (attribute, value) in &attributes {
            if !["data-variant", "data-size", "data-color"].contains(attribute) {
                continue;
            }
            let selected = VALUES
                .iter()
                .any(|(c, a, v)| a == attribute && v == value && ours.contains(c));
            if !selected {
                let known: Vec<String> = ours
                    .iter()
                    .map(|c| {
                        let values: Vec<&str> = VALUES
                            .iter()
                            .filter(|(vc, a, _)| vc == c && a == attribute)
                            .map(|(_, _, v)| *v)
                            .collect();
                        format!(
                            "{c}'s: {}",
                            if values.is_empty() {
                                "none".to_string()
                            } else {
                                values.join(", ")
                            }
                        )
                    })
                    .collect();
                problems.push(format!(
                    "{attribute}={value:?} on <{tag} class={class:?}>: no rule of ciderui's selects it ({})",
                    known.join("; ")
                ));
            }
        }
    }
    if problems.is_empty() {
        Ok(())
    } else {
        Err(problems)
    }
}

/// Each start tag in `html`: its name and its attributes, names lowercase
/// as written, values unquoted; comments, and what a `<script>` or `<style>`
/// holds, skipped.
fn start_tags(html: &str) -> Vec<(&str, Vec<(&str, &str)>)> {
    let mut tags = Vec::new();
    let mut rest = html;
    while let Some(i) = rest.find('<') {
        rest = &rest[i + 1..];
        if let Some(after) = rest.strip_prefix("!--") {
            rest = &after[after.find("-->").map_or(after.len(), |e| e + 3)..];
            continue;
        }
        if !rest.starts_with(|c: char| c.is_ascii_alphabetic()) {
            continue;
        }
        let end = rest
            .find(|c: char| c.is_ascii_whitespace() || c == '>' || c == '/')
            .unwrap_or(rest.len());
        let tag = &rest[..end];
        rest = &rest[end..];
        let mut attributes = Vec::new();
        loop {
            rest = rest.trim_start_matches(|c: char| c.is_ascii_whitespace() || c == '/');
            if rest.is_empty() {
                break;
            }
            if let Some(after) = rest.strip_prefix('>') {
                rest = after;
                break;
            }
            let end = rest
                .find(|c: char| c.is_ascii_whitespace() || c == '=' || c == '>' || c == '/')
                .unwrap_or(rest.len());
            let name = &rest[..end];
            rest = &rest[end..];
            let mut value = "";
            if let Some(after) = rest.strip_prefix('=') {
                if let Some(quote) = after.chars().next().filter(|q| *q == '"' || *q == '\'') {
                    let close = after[1..].find(quote).map_or(after.len(), |e| e + 1);
                    value = &after[1..close];
                    rest = after.get(close + 1..).unwrap_or("");
                } else {
                    let end = after
                        .find(|c: char| c.is_ascii_whitespace() || c == '>')
                        .unwrap_or(after.len());
                    value = &after[..end];
                    rest = &after[end..];
                }
            }
            attributes.push((name, value));
        }
        if tag.eq_ignore_ascii_case("script") || tag.eq_ignore_ascii_case("style") {
            let close = format!("</{tag}");
            rest = &rest[rest.find(&close).unwrap_or(rest.len())..];
        }
        tags.push((tag, attributes));
    }
    tags
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs;
    use std::path::Path;

    // every component's examples write only words the stylesheets know
    #[test]
    fn every_example_passes() {
        let components = Path::new(env!("CARGO_MANIFEST_DIR")).join("../components");
        let mut read = 0;
        for set in fs::read_dir(&components).unwrap() {
            for component in fs::read_dir(set.unwrap().path()).unwrap() {
                let Ok(examples) = fs::read_dir(component.unwrap().path().join("examples")) else {
                    continue;
                };
                for example in examples {
                    let path = example.unwrap().path();
                    let html = fs::read_to_string(&path).unwrap();
                    assert_eq!(check_markup(&html), Ok(()), "{}", path.display());
                    read += 1;
                }
            }
        }
        assert!(read > 100, "{read} examples read");
    }

    #[test]
    fn a_misspelt_class_fails() {
        assert_eq!(
            check_markup(r#"<p>a</p><span class="cui-console-badg mt-1">b</span>"#),
            Err(vec![
                r#"class "cui-console-badg" on <span class="cui-console-badg mt-1">: ciderui has no such class"#.to_string()
            ])
        );
    }

    #[test]
    fn a_value_no_rule_selects_fails() {
        assert_eq!(
            check_markup(r#"<span class="cui-console-badge" data-color="purple">b</span>"#),
            Err(vec![
                r#"data-color="purple" on <span class="cui-console-badge">: no rule of ciderui's selects it (cui-console-badge's: brand, error, info, success, warning)"#
                    .to_string()
            ])
        );
    }

    // a value is held to every cui-* class on its element: cui-console-btn-icon's
    // sizes are its own
    #[test]
    fn a_value_of_another_class_on_the_element_passes() {
        assert_eq!(
            check_markup(
                r#"<button class="cui-console-btn cui-console-btn-icon" data-size="sm">x</button>"#
            ),
            Ok(())
        );
    }

    // a page's own data-* and classes, and what a script holds, are not ciderui's
    #[test]
    fn what_is_not_ciderui_s_passes() {
        assert_eq!(
            check_markup(
                r#"<div class="mt-2" data-variant="whatever"><!-- <b class="cui-nope"> --><script>let s = '<a class="cui-nope">';</script></div>"#
            ),
            Ok(())
        );
    }
}
