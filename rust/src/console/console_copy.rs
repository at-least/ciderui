use hypertext::{Buffer, prelude::*};

use crate::lucide;

/// `cui-console-copy` (`components/console/console-copy`): the two icons of
/// a copy button, Lucide's copy and, while `data-state="copied"`, its check,
/// which go in the page's own
/// `<button type="button" class="cui-console-copy-button">` (or
/// `cui-console-copy-inline`), so the button keeps whatever attributes the
/// page's script reads.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleCopyIcons;
///
/// let html = rsx! {
///     <button type="button" class="cui-console-copy-inline" aria-label="Copy issue ID"><ConsoleCopyIcons /></button>
/// }
/// .render();
/// assert!(html.as_inner().contains(r#"aria-label="Copy issue ID"><svg class="cui-console-copy-idle" aria-hidden="true""#));
/// ```
pub struct ConsoleCopyIcons;

impl Renderable for ConsoleCopyIcons {
    fn render_to(&self, buffer: &mut Buffer) {
        rsx! {
            (lucide::COPY.class("cui-console-copy-idle"))(lucide::CHECK.class("cui-console-copy-done"))
        }
        .render_to(buffer);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    #[test]
    fn renders_the_examples() {
        let copy = rsx! {
            <div class="cui-console-page max-w-xl">
                <div class="cui-console-card grid grid-cols-1 gap-3">
                    <div class="cui-console-copy">
                        <pre class="cui-console-listing"><code id="copy-cli">"sentry-cli debug-files upload --org acme --project shop ./build"</code></pre>
                        <button type="button" class="cui-console-copy-button" aria-label="Copy"><ConsoleCopyIcons /></button>
                    </div>
                    <div class="cui-console-copy">
                        <pre class="cui-console-listing"><code id="copy-dsn">"https://4f1c@crash.example.com/2"</code></pre>
                        <button type="button" class="cui-console-copy-button" aria-label="Copied" data-state="copied"><ConsoleCopyIcons /></button>
                    </div>
                </div>
            </div>
        }
        .render();
        assert_eq!(
            *copy.as_inner(),
            example(include_str!(
                "../../../components/console/console-copy/examples/card.html"
            ))
        );
        let inline = rsx! {
            <div class="cui-console-page grid gap-3">
                <p>"Issue ID "<span class="font-mono" id="issue-id">"60674cfe203893d8499985f493f86663"</span>" "<button type="button" class="cui-console-copy-inline" aria-label="Copy issue ID"><ConsoleCopyIcons /></button></p>
                <p>"Event ID "<span class="font-mono" id="event-id">"e1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6"</span>" "<button type="button" class="cui-console-copy-inline" aria-label="Copied" data-state="copied"><ConsoleCopyIcons /></button></p>
                <p>"Release "<span class="font-mono">"60674cfe203893d8499985f493f86663"</span></p>
            </div>
        }
        .render();
        assert_eq!(
            *inline.as_inner(),
            example(include_str!(
                "../../../components/console/console-copy/examples/inline.html"
            ))
        );
    }
}
