use std::fmt::Write;

use hypertext::{Buffer, Raw, prelude::*};

use crate::Class;
use crate::console::{ConsoleColor, ConsoleStatus};

/// One series of a [`ConsoleChart`]'s legend.
pub struct ConsoleChartSeries<'a> {
    /// Its name.
    pub name: &'a str,
    /// Its marks' fill, a CSS color.
    pub color: &'a str,
}

/// `cui-console-chart` (`components/console/console-chart`): App Store
/// Connect's analytics chart, what is around the page's plot (`children`:
/// its `<svg class="cui-console-chart-svg">`, gridlines and marks, which
/// the page draws from its data): the scale over it, the empty note, the
/// axis's first and last labels under it, and the legend, a console status
/// for each series.
///
/// ```
/// use hypertext::{Raw, prelude::*};
/// use ciderui::console::ConsoleChart;
///
/// let html = rsx! {
///     <ConsoleChart peak="" half="" empty="No data in this window" axis_start="Oct 22" axis_end="Oct 28" legend=(&[]) class="">
///         (Raw::<_>::dangerously_create(r#"<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" role="img" aria-label="Events: 0"></svg>"#))
///     </ConsoleChart>
/// }
/// .render();
/// assert!(html.as_inner().starts_with(r#"<figure class="cui-console-chart"><div class="cui-console-chart-plot"><p class="cui-console-chart-empty">No data in this window</p><svg"#));
/// ```
pub struct ConsoleChart<'a, R: Renderable> {
    /// The top gridline's value, at the plot's top end; `""` for no scale.
    pub peak: &'a str,
    /// The middle gridline's value; `""` where `peak` is.
    pub half: &'a str,
    /// What the plot has none of, centered on it; `""` for a plot with
    /// data.
    pub empty: &'a str,
    /// The first bucket's label under the plot; `""` with `axis_end` for
    /// no axis.
    pub axis_start: &'a str,
    /// The last bucket's label.
    pub axis_end: &'a str,
    /// The series, for a chart of more than one; `&[]` for no legend.
    pub legend: &'a [ConsoleChartSeries<'a>],
    /// The page's utilities on the `<figure>`; `""` for none.
    pub class: &'a str,
    /// The plot.
    pub children: R,
}

impl<R: Renderable> Renderable for ConsoleChart<'_, R> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            self.peak.is_empty() == self.half.is_empty(),
            "ciderui::console::ConsoleChart's scale is a peak and a half, both or neither"
        );
        assert!(
            self.axis_start.is_empty() == self.axis_end.is_empty(),
            "ciderui::console::ConsoleChart's axis is a start and an end, both or neither"
        );
        rsx! {
            <figure class=(Class("cui-console-chart", self.class))>
                <div class="cui-console-chart-plot">
                    @if !self.peak.is_empty() {
                        <span class="cui-console-chart-scale" aria-hidden="true">(self.peak)</span>
                        <span class="cui-console-chart-scale" data-at="half" aria-hidden="true">(self.half)</span>
                    }
                    @if !self.empty.is_empty() {
                        <p class="cui-console-chart-empty">(self.empty)</p>
                    }
                    (self.children)
                </div>
                @if !self.axis_start.is_empty() {
                    <div class="cui-console-chart-axis" aria-hidden="true"><span>(self.axis_start)</span><span>(self.axis_end)</span></div>
                }
                @if !self.legend.is_empty() {
                    <div class="cui-console-chart-legend">
                        @for series in self.legend {
                            <ConsoleStatus label=(series.name) color=(ConsoleColor::Neutral) series=(series.color) class="" />
                        }
                    </div>
                }
            </figure>
        }
        .render_to(buffer);
    }
}

/// `cui-console-chart-sparkline` (`components/console/console-chart`): a
/// small line chart, its line in the brand's blue over its area in the
/// blue's tint, down to its foot.
///
/// ```
/// use hypertext::prelude::*;
/// use ciderui::console::ConsoleSparkline;
///
/// let html = rsx! { <ConsoleSparkline label="3 events over 2 days" width=40 height=28 points=(&[(0.0, 20.0), (40.0, 8.0)]) class="" /> }.render();
/// assert_eq!(
///     html.as_inner(),
///     r#"<svg class="cui-console-chart-sparkline" viewBox="0 0 40 28" width="40" height="28" role="img" aria-label="3 events over 2 days"><polygon class="cui-console-chart-sparkline-area" points="0,28 0,20 40,8 40,28"/><polyline class="cui-console-chart-sparkline-line" points="0,20 40,8"/></svg>"#
/// );
/// ```
pub struct ConsoleSparkline<'a> {
    /// What it shows, in words: the svg's label. Required: an empty one
    /// panics.
    pub label: &'a str,
    /// Its width, in its own units and in CSS pixels.
    pub width: u32,
    /// Its height, the same.
    pub height: u32,
    /// The line's points, `(x, y)` in its units, from its start (`x` 0) to
    /// its end (`x` `width`); y down.
    pub points: &'a [(f64, f64)],
    /// The page's utilities on the svg; `""` for none.
    pub class: &'a str,
}

impl Renderable for ConsoleSparkline<'_> {
    fn render_to(&self, buffer: &mut Buffer) {
        assert!(
            !self.label.is_empty(),
            "ciderui::console::ConsoleSparkline needs a label"
        );
        let mut line = String::new();
        for (i, (x, y)) in self.points.iter().enumerate() {
            if i > 0 {
                line.push(' ');
            }
            write!(line, "{x},{y}").expect("a String takes any write");
        }
        let (w, h) = (self.width, self.height);
        // a string, not rsx!: hypertext's rsx! knows no SVG element
        let mut svg = String::from(r#"<svg class="cui-console-chart-sparkline"#);
        if !self.class.is_empty() {
            svg.push(' ');
            escape_attribute(self.class, &mut svg);
        }
        write!(
            svg,
            r#"" viewBox="0 0 {w} {h}" width="{w}" height="{h}" role="img" aria-label=""#
        )
        .expect("a String takes any write");
        escape_attribute(self.label, &mut svg);
        write!(
            svg,
            r#""><polygon class="cui-console-chart-sparkline-area" points="0,{h} {line} {w},{h}"/><polyline class="cui-console-chart-sparkline-line" points="{line}"/></svg>"#
        )
        .expect("a String takes any write");
        Raw::<_>::dangerously_create(svg).render_to(buffer);
    }
}

/// `value` into `out` as a double-quoted attribute's value.
fn escape_attribute(value: &str, out: &mut String) {
    for ch in value.chars() {
        match ch {
            '&' => out.push_str("&amp;"),
            '"' => out.push_str("&quot;"),
            '<' => out.push_str("&lt;"),
            _ => out.push(ch),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::example::example;

    const GRID: &str = r#"<line class="cui-console-chart-grid" x1="0" x2="600" y1="105" y2="105"/><line class="cui-console-chart-grid" x1="0" x2="600" y1="70" y2="70"/><line class="cui-console-chart-grid" x1="0" x2="600" y1="35" y2="35"/>"#;

    #[test]
    fn renders_the_examples() {
        let bars = concat!(
            r#"<g class="cui-console-chart-bar"><title>Oct 22: 8</title><rect class="cui-console-chart-mark" x="3" y="94" width="80" height="22" fill="var(--cui-c-console-brand-accent-primary)"/><rect class="cui-console-chart-mark" x="3" y="116" width="80" height="24" fill="var(--cui-c-console-neutral-label-secondary)"/></g>"#,
            r#"<g class="cui-console-chart-bar"><title>Oct 24: 14</title><rect class="cui-console-chart-mark" x="89" y="60" width="80" height="40" fill="var(--cui-c-console-brand-accent-primary)"/><rect class="cui-console-chart-mark" x="89" y="100" width="80" height="40" fill="var(--cui-c-console-neutral-label-secondary)"/></g>"#,
            r#"<g class="cui-console-chart-bar"><title>Oct 27: 24</title><rect class="cui-console-chart-mark" x="175" y="4" width="80" height="96" fill="var(--cui-c-console-brand-accent-primary)"/><rect class="cui-console-chart-mark" x="175" y="100" width="80" height="40" fill="var(--cui-c-console-neutral-label-secondary)"/></g>"#,
        );
        let legend = [
            ConsoleChartSeries {
                name: "2.4.1",
                color: "var(--cui-c-console-brand-accent-primary)",
            },
            ConsoleChartSeries {
                name: "2.4.0",
                color: "var(--cui-c-console-neutral-label-secondary)",
            },
        ];
        let chart = rsx! {
            <div class="cui-console-page max-w-xl">
                <ConsoleChart peak="peak 24 · Oct 27" half="12" empty="" axis_start="Oct 22" axis_end="Oct 28" legend=(&legend) class="">
                    (Raw::<_>::dangerously_create(format!(r#"<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none" role="img" aria-label="Events: 96, peak 24 in one day, by release: 2.4.1 ×60, 2.4.0 ×36">{GRID}{bars}</svg>"#)))
                </ConsoleChart>
            </div>
        }
        .render();
        assert_eq!(
            *chart.as_inner(),
            example(include_str!(
                "../../../components/console/console-chart/examples/chart.html"
            ))
        );
        let empty = rsx! {
            <div class="cui-console-page max-w-xl">
                <ConsoleChart peak="" half="" empty="No unhandled errors in this window" axis_start="Oct 22" axis_end="Oct 28" legend=(&[]) class="">
                    (Raw::<_>::dangerously_create(format!(r#"<svg class="cui-console-chart-svg" viewBox="0 0 600 140" height="140" preserveAspectRatio="none" role="img" aria-label="Events: 0">{GRID}</svg>"#)))
                </ConsoleChart>
            </div>
        }
        .render();
        assert_eq!(
            *empty.as_inner(),
            example(include_str!(
                "../../../components/console/console-chart/examples/empty.html"
            ))
        );
        let points = [
            (0.0, 20.0),
            (20.0, 16.0),
            (40.0, 22.0),
            (60.0, 8.0),
            (80.0, 12.0),
            (100.0, 4.0),
            (120.0, 10.0),
        ];
        let sparkline = rsx! {
            <div class="cui-console-page">
                <ConsoleSparkline label="42 events over 7 days" width=120 height=28 points=(&points) class="" />
            </div>
        }
        .render();
        assert_eq!(
            *sparkline.as_inner(),
            example(include_str!(
                "../../../components/console/console-chart/examples/sparkline.html"
            ))
        );
    }

    #[test]
    #[should_panic(expected = "a peak and a half, both or neither")]
    fn a_peak_without_its_half_panics() {
        let _ = rsx! { <ConsoleChart peak="peak 3" half="" empty="" axis_start="" axis_end="" legend=(&[]) class="">""</ConsoleChart> }.render();
    }

    #[test]
    #[should_panic(expected = "a start and an end, both or neither")]
    fn an_axis_without_its_end_panics() {
        let _ = rsx! { <ConsoleChart peak="" half="" empty="" axis_start="Oct 22" axis_end="" legend=(&[]) class="">""</ConsoleChart> }.render();
    }

    #[test]
    #[should_panic(expected = "needs a label")]
    fn a_sparkline_without_a_label_panics() {
        let _ =
            rsx! { <ConsoleSparkline label="" width=1 height=1 points=(&[]) class="" /> }.render();
    }
}
