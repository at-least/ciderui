//! The kit's stylesheets, embedded (the `css` feature): the root ones
//! (`index.css` and its parts, `fonts.css`) and every
//! `components/<name>/<name>.css`, by their paths in the kit. A site writes
//! them out before its Tailwind build, so the crate's commit is the one pin
//! for its markup and its CSS:
//!
//! ```no_run
//! ciderui::css::write("target/ciderui").unwrap();
//! ```
//!
//! and its entry imports from there (`@import "../target/ciderui/index.css";`,
//! `@import "../target/ciderui/components/button/button.css";`). The Inter
//! font files are not embedded: `fonts.css` refers to them as
//! `url("fonts/…")`, relative to the compiled stylesheet, so a site copies
//! the kit's `fonts/` directory beside its built CSS once.
use std::fs;
use std::io;
use std::path::Path;

include!(concat!(env!("OUT_DIR"), "/css_files.rs"));

/// Writes every stylesheet of [`FILES`] under `dir`, at its path in the
/// kit, creating the directories.
pub fn write(dir: impl AsRef<Path>) -> io::Result<()> {
    let dir = dir.as_ref();
    for (rel, text) in FILES {
        let path = dir.join(rel);
        if let Some(parent) = path.parent() {
            fs::create_dir_all(parent)?;
        }
        fs::write(path, text)?;
    }
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn the_kit_is_embedded() {
        let paths: Vec<&str> = FILES.iter().map(|(p, _)| *p).collect();
        for want in [
            "index.css",
            "tokens.css",
            "type.css",
            "fonts.css",
            "components/button/button.css",
        ] {
            assert!(paths.contains(&want), "{want} is not embedded");
        }
        let (_, tokens) = FILES.iter().find(|(p, _)| *p == "tokens.css").unwrap();
        assert!(tokens.contains("--cui-c-accent:"));
    }

    #[test]
    fn writes_the_kits_layout() {
        let dir = std::env::temp_dir().join(format!("ciderui-css-{}", std::process::id()));
        write(&dir).unwrap();
        let button = fs::read_to_string(dir.join("components/button/button.css")).unwrap();
        assert!(button.contains(".cui-btn"));
        fs::remove_dir_all(&dir).unwrap();
    }
}
