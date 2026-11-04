+++
title = "Console Play Button"
description = "A round play and pause button for one audio clip."

[extra]
category = "Console"
image = "pictures/console-play.png"
+++

A round button that plays one audio clip and pauses it. It holds both glyphs, Lucide’s play and pause filled; the page’s script sets `data-state="playing"` while the clip plays and changes the `aria-label` to say what pressing it does now. App Store Connect’s toolkit has no play button: this is its gray icon button at its capsule shape (`cui-console-btn` with `data-shape="capsule"`).

```css
@import "ciderui/components/console/console-play/console-play.css";
```

{{ <cui_example file="console-play/play" title="Paused and playing" pad={true} /> }}

| class or attribute | what |
| --- | --- |
| `cui-console-play` | on a `cui-console-btn cui-console-btn-icon` with `data-shape="capsule"` |
| `cui-console-play-paused` | the play glyph, shown while paused |
| `cui-console-play-playing` | the pause glyph, shown while playing |
| `data-state="playing"` | the clip is playing; the page’s script sets it |
| `aria-label` | what pressing it does now: “Play”, or “Pause” while playing |

## Sizes

The console button’s sizes: `data-size="sm"`, `"lg"` and `"xl"`, 22px, 40px and 50px across.

{{ <cui_example file="console-play/sizes" title="Console play button sizes" pad={true} /> }}

The Go module writes it with `console.ConsolePlay`, its state a `console.ConsolePlayState` and its size a `console.ConsoleBtnSize`.
