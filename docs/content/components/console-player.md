+++
title = "Console Player"
description = "tv.apple.com’s video player: its controls over the video, the centre play, the scrubber and its times."

[extra]
category = "Console"
image = "pictures/console-player.png"
+++

A video’s controls laid over it, as tv.apple.com’s player draws them: the volume at the top, play or pause in the middle, and at the foot a title, the buttons for subtitles, voices and full screen, and the scrubber with the time played and the time left.

```css
@import "ciderui/components/console/console-player/console-player.css";
```

{{ <cui_example file="console-player/paused" title="Console player, paused" /> }}

Playing, the centre button shows pause:

{{ <cui_example file="console-player/playing" title="Console player, playing" /> }}

Its controls hidden while it plays (`data-controls="hidden"`), the video alone:

{{ <cui_example file="console-player/hidden" title="Console player, controls hidden" /> }}

| class | what |
| --- | --- |
| `cui-console-player` | the player: black, its `<video>` and overlay in one place; its size is the page’s (`aspect-video` here) |
| `cui-console-player-overlay` | the controls, 30px from the sides and 20px from the top and foot, over a gradient |
| `cui-console-player-header` | the volume: a range and the mute button |
| `cui-console-player-play` | the centre button, 66px; it holds both glyphs |
| `cui-console-player-glyph-play`, `cui-console-player-glyph-pause` | the glyph shown while paused, and while playing |
| `cui-console-player-footer` | the title, the buttons and the scrubber |
| `cui-console-player-title` | the title, 26px bold |
| `cui-console-player-buttons` | the buttons’ row, 14px apart |
| `cui-console-player-btn` | a button of one glyph, 24px |
| `cui-console-player-progress` | the scrubber over the two times |
| `cui-console-player-range` | an `<input type="range">`: the scrubber, or the volume in the header |
| `cui-console-player-time` | the time played, then the time left |

## States

ciderui has no script: the page’s plays the video and keeps the markup in step with it, as tv.apple.com’s does.

| attribute | on | what |
| --- | --- | --- |
| `data-state="playing"` | `cui-console-player` | playing: the centre button shows pause. None, paused |
| `data-controls="hidden"` | `cui-console-player` | the controls faded out: tv.apple.com hides them while the video plays and the pointer is still, and keeps them while it is paused. They show again while the focus is inside the player |
| `--cui-console-player-value` | `cui-console-player-range` | the part played, or the volume, as a percentage (`style="--cui-console-player-value: 8%"`) |

The script also changes the centre button’s `aria-label` (Play, Pause) with the state.

## Menus

The subtitles and voices buttons open a popover by `popovertarget`: a `cui-console-popover` holding a `cui-console-menu`, the page’s. tv.apple.com’s own menu is drawn otherwise (a blurred light ground, 6px corners, 32px rows with a check at their start: `reference/apple.com/`’s `tv-player-menu`); the console keeps its one menu.

## Go

The Go module writes it with `console.ConsolePlayer`: the page’s `<video>`, the state, the labels, the volume, the position and duration in seconds (it writes the times), and each menu’s id (none, no button).
