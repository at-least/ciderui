+++
title = "Console Play Button"
description = "A button that plays one audio clip and stops it, a ring round its glyph filling as it plays."

[extra]
category = "Console"
image = "pictures/console-play.png"
+++

A button that plays one audio clip and stops it. It holds both glyphs, Lucide’s play and stop filled; the page’s script sets `data-state="playing"` while the clip plays, and `--cui-console-play-progress` as it goes: a ring round the glyph fills with how much has played, as in Podcasts’ and Voice Memos’ small buttons. A round one of the glyph alone changes its `aria-label` to say what pressing it does now. App Store Connect’s toolkit has no play button: this is its gray icon button at its capsule shape (`cui-console-btn` with `data-shape="capsule"`).

```css
@import "ciderui/components/console/console-play/console-play.css";
```

{{ <cui_example file="console-play/play" title="Paused and playing" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-play` | on a `cui-console-btn` with `data-shape="capsule"` (and `cui-console-btn-icon`, the glyph alone) |
| `cui-console-play-glyph` | a `<span>` holding both glyphs; playing, the ring round it |
| `cui-console-play-paused` | the play glyph, and the words, shown while paused |
| `cui-console-play-playing` | the stop glyph, and the words, shown while playing |
| `data-state="playing"` | the clip is playing; the page’s script sets it |
| `--cui-console-play-progress` | how much has played, 0% to 100%: the ring; the page’s script sets it |
| `aria-label` | the glyph alone: what pressing it does now, “Play”, or “Stop” while playing |

With words, the button says what a press does in them, its glyph before them, and needs no `aria-label`:

{{ <cui_example file="console-play/labelled" title="With words, paused and playing" pad={true} /> }}

## Sizes

The console button’s sizes: `data-size="sm"`, `"lg"` and `"xl"`, 22px, 40px and 50px across.

{{ <cui_example file="console-play/sizes" title="Console play button sizes" pad={true} /> }}

The Go module writes it with `console.ConsolePlay`: its state a `console.ConsolePlayState`, its `AriaLabel` (the glyph alone) or its `Label` and `PlayingLabel` (with words), and its size a `console.ConsoleBtnSize`.
