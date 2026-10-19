//! `ciderui-css <dir>`: writes the kit's stylesheets under `dir`, in the
//! kit's layout, for a Tailwind build to import. Needs the `css` feature: in
//! this repository `cargo run --features css --bin ciderui-css --
//! target/ciderui`; in a site whose dependency turns it on, `cargo build -p
//! ciderui --bin ciderui-css`, then the binary.
fn main() {
    let dir = std::env::args().nth(1).expect("usage: ciderui-css <dir>");
    ciderui::css::write(&dir).unwrap_or_else(|e| panic!("writing {dir}: {e}"));
    println!("ciderui's stylesheets written under {dir}");
}
