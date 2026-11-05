+++
title = "Console Player"
description = "tv.apple.com’s video player: its controls over the video, the centre play, the scrubber and its times."

[extra]
category = "Console"
image = "pictures/console-player.png"
+++

A video’s controls laid over it, as tv.apple.com’s player draws them: the volume at the top, play or pause in the middle, and at the foot a title, the buttons for subtitles and full screen, and the scrubber with the time played and the time left.

```css
@import "ciderui/components/console/console-player/console-player.css";
```

{{ <cui_example file="console-player/paused" title="Console player, paused" /> }}

Playing, the centre button shows pause:

{{ <cui_example file="console-player/playing" title="Console player, playing" /> }}

Its controls hidden while it plays (`data-controls="hidden"`), the video alone:

{{ <cui_example file="console-player/hidden" title="Console player, controls hidden" /> }}

Its picture can be any element in the video’s place, its first child: here a `<div>` of stills the page’s script shows one at a time, with a subtitle over it:

{{ <cui_example file="console-player/stills" title="Console player, stills and a subtitle" /> }}

Filling the screen (`data-fullscreen`), its bar adds the page’s own controls before its buttons:

{{ <cui_example file="console-player/full-screen" title="Console player, full screen" /> }}

| class | what |
| --- | --- |
| `cui-console-player` | the player: black, its picture and overlay in one place; its size is the page’s (`aspect-video` here) |
| its first child | the picture, the player’s whole box: a `<video>` (`object-fit: contain`), or an element the page draws its stills in |
| `cui-console-player-overlay` | the controls, 30px from the sides and 20px from the top and foot, over a gradient |
| `cui-console-player-header` | the top row, the volume at its end |
| `cui-console-player-volume` | the volume: the mute button alone until the pointer is over it (or its range is held or focused from the keyboard), then its range opens, 100px, to the button’s left over 0.3s |
| `cui-console-player-glyph-volume`, `cui-console-player-glyph-muted` | the mute button’s glyph, the speaker, and struck through while muted |
| `cui-console-player-play` | the centre button, 66px; it holds both glyphs |
| `cui-console-player-glyph-play`, `cui-console-player-glyph-pause` | the glyph shown while paused, and while playing |
| `cui-console-player-subtitles` | the subtitles over the picture: above the footer while the controls show, 20px from the foot while they hide, moving in the footer’s 0.4s; it takes no click |
| `cui-console-player-cue` | the cue in it, a `<span>` in its `<p>`: 2.4% of the width (15px at least), so it grows with the player and in full screen, 1.6 lines, white under a soft shadow (`--cui-video-cue-shadow`), no box; empty, none |
| `cui-console-player-footer` | the title, the buttons and the scrubber |
| `cui-console-player-title` | the title, 26px bold |
| `cui-console-player-buttons` | the buttons’ row, 14px apart |
| `cui-console-player-full-screen-controls` | the page’s own controls, first in the buttons’ row, shown only while the player fills the screen |
| `cui-console-player-text-btn` | a control of words in the bar: a 24px capsule, white on a faint white fill |
| `cui-console-player-btn` | a button of one glyph, 24px |
| `cui-console-player-progress` | the scrubber over the two times |
| `cui-console-player-range` | an `<input type="range">`: the scrubber, or the volume in the header; 44px a pointer hits for an 11px row, its line 5px, 8px while held |
| `cui-console-player-preview` | over the scrubber, at the pointer: a picture (`cui-console-player-preview-picture`, 160px at 16:9) over a line (`cui-console-player-preview-label`), hidden until the page’s script shows it; it stays inside the bar’s ends and takes no pointer |
| `cui-console-player-time` | the time played, then the time left |

## States

ciderui has no script: the page’s plays the video and keeps the markup in step with it, as tv.apple.com’s does.

| attribute | on | what |
| --- | --- | --- |
| `data-state="playing"` | `cui-console-player` | playing: the centre button shows pause. None, paused |
| `data-fullscreen` | `cui-console-player` | it fills the screen: the bar shows the page’s full-screen controls |
| `data-muted` | `cui-console-player` | muted: the mute button shows the struck speaker |
| `data-controls="hidden"` | `cui-console-player` | the controls faded out: tv.apple.com hides them while the video plays and the pointer is still, and keeps them while it is paused. They show again while the focus is inside the player |
| `--cui-console-player-value` | `cui-console-player-range` | the part played, or the volume, as a percentage (`style="--cui-console-player-value: 8%"`) |
| `hidden` | `cui-console-player-preview` | written hidden; the script takes it away while the pointer is over the scrubber or its thumb is held |
| `--cui-console-player-preview-x` | `cui-console-player-preview` | the pointer’s x in the bar, a length (`300px`): the preview’s centre |

The script also changes the centre button’s `aria-label` (Play, Pause) with the state, and writes the cue’s text (and its `lang`) as the picture plays.

## Hooks

The script finds each control by its `data-control`, the same on every player:

| `data-control` | what |
| --- | --- |
| `play` | the centre button |
| `volume`, `mute` | the volume’s range, and its button |
| `subtitles` | the button that opens the menu |
| `full-screen` | the full-screen button |
| `scrubber` | the scrubber’s range |
| `elapsed`, `remaining` | the time played, and the time left |
| `preview` | the scrubber’s preview: the script sets its `<img>`’s `src` and its `<p>`’s text |

The cue is `cui-console-player-cue`.

## Menus

The subtitles button opens a popover by `popovertarget`: a menu of choices, a [Console Menu](@/components/console-menu.md#choices)’s, the voices and the subtitles in columns side by side (one button, one panel, as the Apple TV app’s Audio and Subtitles; tv.apple.com’s own audio is a platform `<select>`), written inside the player after its overlay, so the controls stay while the focus is in it (not inside the overlay, which places its own parts). tv.apple.com’s own menu is drawn otherwise (a blurred light ground, 6px corners, 32px rows: `reference/apple.com/`’s `tv-player-menu`); the console keeps its one menu, and takes tv.apple.com’s check at each choice’s start. A menu opens upward inside the player, as tv.apple.com’s: its foot and its right edge on its button’s centre, at most as tall as the player leaves above it (20px from its top), its choices scrolling inside it.

## Go

The Go module writes it with `console.ConsolePlayer`: the picture (`Video`), the state, the labels, the volume, whether it is muted or fills the screen, the position and duration in seconds (it writes the times), the cue and its language, the menu, a `console.ConsoleChoiceMenu` (none, no button), the page’s full-screen controls (`FullScreenControls`) and its own popovers (`Popovers`, written inside the player after its menus and placed as they are).
