// Package lucide is every Lucide icon (https://lucide.dev/icons/; under its
// license, the repository's LICENSE-Lucide), lucide-static 1.61.0's, by its
// name in Go's capitals: shopping-bag is [ShoppingBag]. Each is an [icon.Icon]
// of Lucide's own elements in currentColor, 1em square, so it takes the text's
// color and size; a utility on it sizes it otherwise. In a line of text a page
// adds inline: Tailwind's preflight makes an <svg> a block. A part apple.com
// draws an icon in draws ciderui's own (package icon), not Lucide's.
//
// Written by scripts/build-lucide.mjs from lucide-static's icon-nodes.json and
// icons/, not by hand: npm test fails while this differs from what it writes.
package lucide

import "github.com/at-least/ciderui/go/icon"

const (
	viewBox = "0 0 24 24"
	paint   = `fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"`
)

// AArrowDown is [a-arrow-down](https://lucide.dev/icons/a-arrow-down).
var AArrowDown = icon.Lucide(viewBox, paint, `<path d="m14 12 4 4 4-4"/><path d="M18 16V7"/><path d="m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"/><path d="M3.304 13h6.392"/>`)

// AArrowUp is [a-arrow-up](https://lucide.dev/icons/a-arrow-up).
var AArrowUp = icon.Lucide(viewBox, paint, `<path d="m14 11 4-4 4 4"/><path d="M18 16V7"/><path d="m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"/><path d="M3.304 13h6.392"/>`)

// ALargeSmall is [a-large-small](https://lucide.dev/icons/a-large-small).
var ALargeSmall = icon.Lucide(viewBox, paint, `<path d="m15 16 2.536-7.328a1.02 1.02 1 0 1 1.928 0L22 16"/><path d="M15.697 14h5.606"/><path d="m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"/><path d="M3.304 13h6.392"/>`)

// Accessibility is [accessibility](https://lucide.dev/icons/accessibility).
var Accessibility = icon.Lucide(viewBox, paint, `<circle cx="16" cy="4" r="1"/><path d="m18 19 1-7-6 1"/><path d="m5 8 3-3 5.5 3-2.36 3.5"/><path d="M4.24 14.5a5 5 0 0 0 6.88 6"/><path d="M13.76 17.5a5 5 0 0 0-6.88-6"/>`)

// Activity is [activity](https://lucide.dev/icons/activity).
var Activity = icon.Lucide(viewBox, paint, `<path d="M22 12h-2.48a2 2 0 0 0-1.93 1.46l-2.35 8.36a.25.25 0 0 1-.48 0L9.24 2.18a.25.25 0 0 0-.48 0l-2.35 8.36A2 2 0 0 1 4.49 12H2"/>`)

// Ad is [ad](https://lucide.dev/icons/ad).
var Ad = icon.Lucide(viewBox, paint, `<path d="M10 13H6"/><path d="M10 15v-4a2 2 0 0 0-4 0v4"/><path d="M14 14.5a.5.5 0 0 0 .5.5h1a2.5 2.5 0 0 0 2.5-2.5v-1A2.5 2.5 0 0 0 15.5 9h-1a.5.5 0 0 0-.5.5z"/><rect x="2" y="5" width="20" height="14" rx="2"/>`)

// AirVent is [air-vent](https://lucide.dev/icons/air-vent).
var AirVent = icon.Lucide(viewBox, paint, `<path d="M18 17.5a2.5 2.5 0 1 1-4 2.03V12"/><path d="M6 12H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"/><path d="M6 8h12"/><path d="M6.6 15.572A2 2 0 1 0 10 17v-5"/>`)

// Airplay is [airplay](https://lucide.dev/icons/airplay).
var Airplay = icon.Lucide(viewBox, paint, `<path d="M5 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1"/><path d="m12 15 5 6H7Z"/>`)

// AlarmClockCheck is [alarm-clock-check](https://lucide.dev/icons/alarm-clock-check).
var AlarmClockCheck = icon.Lucide(viewBox, paint, `<circle cx="12" cy="13" r="8"/><path d="M5 3 2 6"/><path d="m22 6-3-3"/><path d="M6.38 18.7 4 21"/><path d="M17.64 18.67 20 21"/><path d="m9 13 2 2 4-4"/>`)

// AlarmClockMinus is [alarm-clock-minus](https://lucide.dev/icons/alarm-clock-minus).
var AlarmClockMinus = icon.Lucide(viewBox, paint, `<circle cx="12" cy="13" r="8"/><path d="M5 3 2 6"/><path d="m22 6-3-3"/><path d="M6.38 18.7 4 21"/><path d="M17.64 18.67 20 21"/><path d="M9 13h6"/>`)

// AlarmClockOff is [alarm-clock-off](https://lucide.dev/icons/alarm-clock-off).
var AlarmClockOff = icon.Lucide(viewBox, paint, `<path d="M6.87 6.87a8 8 0 1 0 11.26 11.26"/><path d="M19.9 14.25a8 8 0 0 0-9.15-9.15"/><path d="m22 6-3-3"/><path d="M6.26 18.67 4 21"/><path d="m2 2 20 20"/><path d="M4 4 2 6"/>`)

// AlarmClockPlus is [alarm-clock-plus](https://lucide.dev/icons/alarm-clock-plus).
var AlarmClockPlus = icon.Lucide(viewBox, paint, `<circle cx="12" cy="13" r="8"/><path d="M5 3 2 6"/><path d="m22 6-3-3"/><path d="M6.38 18.7 4 21"/><path d="M17.64 18.67 20 21"/><path d="M12 10v6"/><path d="M9 13h6"/>`)

// AlarmClock is [alarm-clock](https://lucide.dev/icons/alarm-clock).
var AlarmClock = icon.Lucide(viewBox, paint, `<circle cx="12" cy="13" r="8"/><path d="M12 9v4l2 2"/><path d="M5 3 2 6"/><path d="m22 6-3-3"/><path d="M6.38 18.7 4 21"/><path d="M17.64 18.67 20 21"/>`)

// AlarmSmoke is [alarm-smoke](https://lucide.dev/icons/alarm-smoke).
var AlarmSmoke = icon.Lucide(viewBox, paint, `<path d="M11 21c0-2.5 2-2.5 2-5"/><path d="M16 21c0-2.5 2-2.5 2-5"/><path d="m19 8-.8 3a1.25 1.25 0 0 1-1.2 1H7a1.25 1.25 0 0 1-1.2-1L5 8"/><path d="M21 3a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a1 1 0 0 1 1-1z"/><path d="M6 21c0-2.5 2-2.5 2-5"/>`)

// AlignCenterHorizontal is [align-center-horizontal](https://lucide.dev/icons/align-center-horizontal).
var AlignCenterHorizontal = icon.Lucide(viewBox, paint, `<path d="M2 12h20"/><path d="M10 16v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4"/><path d="M10 8V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4"/><path d="M20 16v1a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-1"/><path d="M14 8V7c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1"/>`)

// AlignCenterVertical is [align-center-vertical](https://lucide.dev/icons/align-center-vertical).
var AlignCenterVertical = icon.Lucide(viewBox, paint, `<path d="M12 2v20"/><path d="M8 10H4a2 2 0 0 1-2-2V6c0-1.1.9-2 2-2h4"/><path d="M16 10h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4"/><path d="M8 20H7a2 2 0 0 1-2-2v-2c0-1.1.9-2 2-2h1"/><path d="M16 14h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1"/>`)

// AlignEndHorizontal is [align-end-horizontal](https://lucide.dev/icons/align-end-horizontal).
var AlignEndHorizontal = icon.Lucide(viewBox, paint, `<rect width="6" height="16" x="4" y="2" rx="2"/><rect width="6" height="9" x="14" y="9" rx="2"/><path d="M22 22H2"/>`)

// AlignEndVertical is [align-end-vertical](https://lucide.dev/icons/align-end-vertical).
var AlignEndVertical = icon.Lucide(viewBox, paint, `<rect width="16" height="6" x="2" y="4" rx="2"/><rect width="9" height="6" x="9" y="14" rx="2"/><path d="M22 22V2"/>`)

// AlignHorizontalDistributeCenter is [align-horizontal-distribute-center](https://lucide.dev/icons/align-horizontal-distribute-center).
var AlignHorizontalDistributeCenter = icon.Lucide(viewBox, paint, `<rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M17 22v-5"/><path d="M17 7V2"/><path d="M7 22v-3"/><path d="M7 5V2"/>`)

// AlignHorizontalDistributeEnd is [align-horizontal-distribute-end](https://lucide.dev/icons/align-horizontal-distribute-end).
var AlignHorizontalDistributeEnd = icon.Lucide(viewBox, paint, `<rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M10 2v20"/><path d="M20 2v20"/>`)

// AlignHorizontalDistributeStart is [align-horizontal-distribute-start](https://lucide.dev/icons/align-horizontal-distribute-start).
var AlignHorizontalDistributeStart = icon.Lucide(viewBox, paint, `<rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M4 2v20"/><path d="M14 2v20"/>`)

// AlignHorizontalJustifyCenter is [align-horizontal-justify-center](https://lucide.dev/icons/align-horizontal-justify-center).
var AlignHorizontalJustifyCenter = icon.Lucide(viewBox, paint, `<rect width="6" height="14" x="2" y="5" rx="2"/><rect width="6" height="10" x="16" y="7" rx="2"/><path d="M12 2v20"/>`)

// AlignHorizontalJustifyEnd is [align-horizontal-justify-end](https://lucide.dev/icons/align-horizontal-justify-end).
var AlignHorizontalJustifyEnd = icon.Lucide(viewBox, paint, `<rect width="6" height="14" x="2" y="5" rx="2"/><rect width="6" height="10" x="12" y="7" rx="2"/><path d="M22 2v20"/>`)

// AlignHorizontalJustifyStart is [align-horizontal-justify-start](https://lucide.dev/icons/align-horizontal-justify-start).
var AlignHorizontalJustifyStart = icon.Lucide(viewBox, paint, `<rect width="6" height="14" x="6" y="5" rx="2"/><rect width="6" height="10" x="16" y="7" rx="2"/><path d="M2 2v20"/>`)

// AlignHorizontalSpaceAround is [align-horizontal-space-around](https://lucide.dev/icons/align-horizontal-space-around).
var AlignHorizontalSpaceAround = icon.Lucide(viewBox, paint, `<rect width="6" height="10" x="9" y="7" rx="2"/><path d="M4 22V2"/><path d="M20 22V2"/>`)

// AlignHorizontalSpaceBetween is [align-horizontal-space-between](https://lucide.dev/icons/align-horizontal-space-between).
var AlignHorizontalSpaceBetween = icon.Lucide(viewBox, paint, `<rect width="6" height="14" x="3" y="5" rx="2"/><rect width="6" height="10" x="15" y="7" rx="2"/><path d="M3 2v20"/><path d="M21 2v20"/>`)

// AlignStartHorizontal is [align-start-horizontal](https://lucide.dev/icons/align-start-horizontal).
var AlignStartHorizontal = icon.Lucide(viewBox, paint, `<rect width="6" height="16" x="4" y="6" rx="2"/><rect width="6" height="9" x="14" y="6" rx="2"/><path d="M22 2H2"/>`)

// AlignStartVertical is [align-start-vertical](https://lucide.dev/icons/align-start-vertical).
var AlignStartVertical = icon.Lucide(viewBox, paint, `<rect width="9" height="6" x="6" y="14" rx="2"/><rect width="16" height="6" x="6" y="4" rx="2"/><path d="M2 2v20"/>`)

// AlignVerticalDistributeCenter is [align-vertical-distribute-center](https://lucide.dev/icons/align-vertical-distribute-center).
var AlignVerticalDistributeCenter = icon.Lucide(viewBox, paint, `<path d="M22 17h-3"/><path d="M22 7h-5"/><path d="M5 17H2"/><path d="M7 7H2"/><rect x="5" y="14" width="14" height="6" rx="2"/><rect x="7" y="4" width="10" height="6" rx="2"/>`)

// AlignVerticalDistributeEnd is [align-vertical-distribute-end](https://lucide.dev/icons/align-vertical-distribute-end).
var AlignVerticalDistributeEnd = icon.Lucide(viewBox, paint, `<rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M2 20h20"/><path d="M2 10h20"/>`)

// AlignVerticalDistributeStart is [align-vertical-distribute-start](https://lucide.dev/icons/align-vertical-distribute-start).
var AlignVerticalDistributeStart = icon.Lucide(viewBox, paint, `<rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M2 14h20"/><path d="M2 4h20"/>`)

// AlignVerticalJustifyCenter is [align-vertical-justify-center](https://lucide.dev/icons/align-vertical-justify-center).
var AlignVerticalJustifyCenter = icon.Lucide(viewBox, paint, `<rect width="14" height="6" x="5" y="16" rx="2"/><rect width="10" height="6" x="7" y="2" rx="2"/><path d="M2 12h20"/>`)

// AlignVerticalJustifyEnd is [align-vertical-justify-end](https://lucide.dev/icons/align-vertical-justify-end).
var AlignVerticalJustifyEnd = icon.Lucide(viewBox, paint, `<rect width="14" height="6" x="5" y="12" rx="2"/><rect width="10" height="6" x="7" y="2" rx="2"/><path d="M2 22h20"/>`)

// AlignVerticalJustifyStart is [align-vertical-justify-start](https://lucide.dev/icons/align-vertical-justify-start).
var AlignVerticalJustifyStart = icon.Lucide(viewBox, paint, `<rect width="14" height="6" x="5" y="16" rx="2"/><rect width="10" height="6" x="7" y="6" rx="2"/><path d="M2 2h20"/>`)

// AlignVerticalSpaceAround is [align-vertical-space-around](https://lucide.dev/icons/align-vertical-space-around).
var AlignVerticalSpaceAround = icon.Lucide(viewBox, paint, `<rect width="10" height="6" x="7" y="9" rx="2"/><path d="M22 20H2"/><path d="M22 4H2"/>`)

// AlignVerticalSpaceBetween is [align-vertical-space-between](https://lucide.dev/icons/align-vertical-space-between).
var AlignVerticalSpaceBetween = icon.Lucide(viewBox, paint, `<rect width="14" height="6" x="5" y="15" rx="2"/><rect width="10" height="6" x="7" y="3" rx="2"/><path d="M2 21h20"/><path d="M2 3h20"/>`)

// Ambulance is [ambulance](https://lucide.dev/icons/ambulance).
var Ambulance = icon.Lucide(viewBox, paint, `<path d="M10 10H6"/><path d="M14 18V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v11a1 1 0 0 0 1 1h2"/><path d="M19 18h2a1 1 0 0 0 1-1v-3.28a1 1 0 0 0-.684-.948l-1.923-.641a1 1 0 0 1-.578-.502l-1.539-3.076A1 1 0 0 0 16.382 8H14"/><path d="M8 8v4"/><path d="M9 18h6"/><circle cx="17" cy="18" r="2"/><circle cx="7" cy="18" r="2"/>`)

// Ampersand is [ampersand](https://lucide.dev/icons/ampersand).
var Ampersand = icon.Lucide(viewBox, paint, `<path d="M16 12h3"/><path d="M17.5 12a8 8 0 0 1-8 8A4.5 4.5 0 0 1 5 15.5c0-6 8-4 8-8.5a3 3 0 1 0-6 0c0 3 2.5 8.5 12 13"/>`)

// Ampersands is [ampersands](https://lucide.dev/icons/ampersands).
var Ampersands = icon.Lucide(viewBox, paint, `<path d="M10 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6 0 1.7 1.3 3 3 3 2.8 0 5-2.2 5-5"/><path d="M22 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6 0 1.7 1.3 3 3 3 2.8 0 5-2.2 5-5"/>`)

// Amphora is [amphora](https://lucide.dev/icons/amphora).
var Amphora = icon.Lucide(viewBox, paint, `<path d="M10 2v5.632c0 .424-.272.795-.653.982A6 6 0 0 0 6 14c.006 4 3 7 5 8"/><path d="M10 5H8a2 2 0 0 0 0 4h.68"/><path d="M14 2v5.632c0 .424.272.795.652.982A6 6 0 0 1 18 14c0 4-3 7-5 8"/><path d="M14 5h2a2 2 0 0 1 0 4h-.68"/><path d="M18 22H6"/><path d="M9 2h6"/>`)

// Anchor is [anchor](https://lucide.dev/icons/anchor).
var Anchor = icon.Lucide(viewBox, paint, `<path d="M12 6v16"/><path d="m19 13 2-1a9 9 0 0 1-18 0l2 1"/><path d="M9 11h6"/><circle cx="12" cy="4" r="2"/>`)

// Angle is [angle](https://lucide.dev/icons/angle).
var Angle = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M3 11a10 10 0 0 1 10 10"/>`)

// Antenna is [antenna](https://lucide.dev/icons/antenna).
var Antenna = icon.Lucide(viewBox, paint, `<path d="M2 12 7 2"/><path d="m7 12 5-10"/><path d="m12 12 5-10"/><path d="m17 12 5-10"/><path d="M4.5 7h15"/><path d="M12 16v6"/>`)

// Anvil is [anvil](https://lucide.dev/icons/anvil).
var Anvil = icon.Lucide(viewBox, paint, `<path d="M7 10H6a4 4 0 0 1-4-4 1 1 0 0 1 1-1h4"/><path d="M7 5a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1 7 7 0 0 1-7 7H8a1 1 0 0 1-1-1z"/><path d="M9 12v5"/><path d="M15 12v5"/><path d="M5 20a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3 1 1 0 0 1-1 1H6a1 1 0 0 1-1-1"/>`)

// Aperture is [aperture](https://lucide.dev/icons/aperture).
var Aperture = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m14.31 8 5.74 9.94"/><path d="M9.69 8h11.48"/><path d="m7.38 12 5.74-9.94"/><path d="M9.69 16 3.95 6.06"/><path d="M14.31 16H2.83"/><path d="m16.62 12-5.74 9.94"/>`)

// AppWindowMac is [app-window-mac](https://lucide.dev/icons/app-window-mac).
var AppWindowMac = icon.Lucide(viewBox, paint, `<rect width="20" height="16" x="2" y="4" rx="2"/><path d="M6 8h.01"/><path d="M10 8h.01"/><path d="M14 8h.01"/>`)

// AppWindow is [app-window](https://lucide.dev/icons/app-window).
var AppWindow = icon.Lucide(viewBox, paint, `<rect x="2" y="4" width="20" height="16" rx="2"/><path d="M10 4v4"/><path d="M2 8h20"/><path d="M6 4v4"/>`)

// Apple is [apple](https://lucide.dev/icons/apple).
var Apple = icon.Lucide(viewBox, paint, `<path d="M12 6.528V3a1 1 0 0 1 1-1h0"/><path d="M18.237 21A15 15 0 0 0 22 11a6 6 0 0 0-10-4.472A6 6 0 0 0 2 11a15.1 15.1 0 0 0 3.763 10 3 3 0 0 0 3.648.648 5.5 5.5 0 0 1 5.178 0A3 3 0 0 0 18.237 21"/>`)

// ArchiveRestore is [archive-restore](https://lucide.dev/icons/archive-restore).
var ArchiveRestore = icon.Lucide(viewBox, paint, `<rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h2"/><path d="M20 8v11a2 2 0 0 1-2 2h-2"/><path d="m9 15 3-3 3 3"/><path d="M12 12v9"/>`)

// ArchiveX is [archive-x](https://lucide.dev/icons/archive-x).
var ArchiveX = icon.Lucide(viewBox, paint, `<rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8"/><path d="m9.5 17 5-5"/><path d="m9.5 12 5 5"/>`)

// Archive is [archive](https://lucide.dev/icons/archive).
var Archive = icon.Lucide(viewBox, paint, `<rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8"/><path d="M10 12h4"/>`)

// Armchair is [armchair](https://lucide.dev/icons/armchair).
var Armchair = icon.Lucide(viewBox, paint, `<path d="M19 9V6a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v3"/><path d="M3 16a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v1.5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5V11a2 2 0 0 0-4 0z"/><path d="M5 18v2"/><path d="M19 18v2"/>`)

// ArmenianDram is [armenian-dram](https://lucide.dev/icons/armenian-dram).
var ArmenianDram = icon.Lucide(viewBox, paint, `<path d="M11 10h8"/><path d="M11 14h8"/><path d="M17 20V10a6 6 0 0 0-12 0"/>`)

// ArrowBigDownDash is [arrow-big-down-dash](https://lucide.dev/icons/arrow-big-down-dash).
var ArrowBigDownDash = icon.Lucide(viewBox, paint, `<path d="M14 8a1 1 0 0 1 1 1v2a1 1 0 0 0 1 1h3.293a.707.707 0 0 1 .5 1.207l-6.939 6.939a1.207 1.207 0 0 1-1.708 0l-6.94-6.94a.707.707 0 0 1 .5-1.206H8a1 1 0 0 0 1-1V9a1 1 0 0 1 1-1z"/><path d="M9 4h6"/>`)

// ArrowBigDown is [arrow-big-down](https://lucide.dev/icons/arrow-big-down).
var ArrowBigDown = icon.Lucide(viewBox, paint, `<path d="M9 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v6a1 1 0 0 0 1 1h3.293a.707.707 0 0 1 .5 1.207l-7.086 7.086a1 1 0 0 1-1.414 0l-7.086-7.086a.707.707 0 0 1 .5-1.207H8a1 1 0 0 0 1-1z"/>`)

// ArrowBigLeftDash is [arrow-big-left-dash](https://lucide.dev/icons/arrow-big-left-dash).
var ArrowBigLeftDash = icon.Lucide(viewBox, paint, `<path d="M13 9a1 1 0 0 1-1-1V4.707a.707.707 0 0 0-1.207-.5l-6.94 6.94a1.207 1.207 0 0 0 0 1.707l6.94 6.94a.707.707 0 0 0 1.207-.5V16a1 1 0 0 1 1-1h2a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1z"/><path d="M20 9v6"/>`)

// ArrowBigLeft is [arrow-big-left](https://lucide.dev/icons/arrow-big-left).
var ArrowBigLeft = icon.Lucide(viewBox, paint, `<path d="M10.793 19.793a.707.707 0 0 0 1.207-.5V16a1 1 0 0 1 1-1h6a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-6a1 1 0 0 1-1-1V4.707a.707.707 0 0 0-1.207-.5l-6.94 6.94a1.207 1.207 0 0 0 0 1.707z"/>`)

// ArrowBigRightDash is [arrow-big-right-dash](https://lucide.dev/icons/arrow-big-right-dash).
var ArrowBigRightDash = icon.Lucide(viewBox, paint, `<path d="M11 9a1 1 0 0 0 1-1V4.707a.707.707 0 0 1 1.207-.5l6.94 6.94a1.207 1.207 0 0 1 0 1.707l-6.94 6.94a.707.707 0 0 1-1.207-.5V16a1 1 0 0 0-1-1H9a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"/><path d="M4 9v6"/>`)

// ArrowBigRight is [arrow-big-right](https://lucide.dev/icons/arrow-big-right).
var ArrowBigRight = icon.Lucide(viewBox, paint, `<path d="M13.207 19.793a.707.707 0 0 1-1.207-.5V16a1 1 0 0 0-1-1H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h6a1 1 0 0 0 1-1V4.707a.707.707 0 0 1 1.207-.5l6.94 6.94a1.207 1.207 0 0 1 0 1.707z"/>`)

// ArrowBigUpDash is [arrow-big-up-dash](https://lucide.dev/icons/arrow-big-up-dash).
var ArrowBigUpDash = icon.Lucide(viewBox, paint, `<path d="M14 16a1 1 0 0 0 1-1v-2a1 1 0 0 1 1-1h3.293a.707.707 0 0 0 .5-1.207l-6.939-6.939a1.207 1.207 0 0 0-1.708 0l-6.94 6.94a.707.707 0 0 0 .5 1.206H8a1 1 0 0 1 1 1v2a1 1 0 0 0 1 1z"/><path d="M9 20h6"/>`)

// ArrowBigUp is [arrow-big-up](https://lucide.dev/icons/arrow-big-up).
var ArrowBigUp = icon.Lucide(viewBox, paint, `<path d="M9 19a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6a1 1 0 0 1 1-1h3.293a.707.707 0 0 0 .5-1.207l-7.086-7.086a1 1 0 0 0-1.414 0l-7.086 7.086a.707.707 0 0 0 .5 1.207H8a1 1 0 0 1 1 1z"/>`)

// ArrowDown01 is [arrow-down-0-1](https://lucide.dev/icons/arrow-down-0-1).
var ArrowDown01 = icon.Lucide(viewBox, paint, `<path d="m3 16 4 4 4-4"/><path d="M7 20V4"/><rect x="15" y="4" width="4" height="6" ry="2"/><path d="M17 20v-6h-2"/><path d="M15 20h4"/>`)

// ArrowDown10 is [arrow-down-1-0](https://lucide.dev/icons/arrow-down-1-0).
var ArrowDown10 = icon.Lucide(viewBox, paint, `<path d="m3 16 4 4 4-4"/><path d="M7 20V4"/><path d="M17 10V4h-2"/><path d="M15 10h4"/><rect x="15" y="14" width="4" height="6" ry="2"/>`)

// ArrowDownAZ is [arrow-down-a-z](https://lucide.dev/icons/arrow-down-a-z).
var ArrowDownAZ = icon.Lucide(viewBox, paint, `<path d="m3 16 4 4 4-4"/><path d="M7 20V4"/><path d="M20 8h-5"/><path d="M15 10V6.5a2.5 2.5 0 0 1 5 0V10"/><path d="M15 14h5l-5 6h5"/>`)

// ArrowDownFromLine is [arrow-down-from-line](https://lucide.dev/icons/arrow-down-from-line).
var ArrowDownFromLine = icon.Lucide(viewBox, paint, `<path d="M19 3H5"/><path d="M12 21V7"/><path d="m6 15 6 6 6-6"/>`)

// ArrowDownLeft is [arrow-down-left](https://lucide.dev/icons/arrow-down-left).
var ArrowDownLeft = icon.Lucide(viewBox, paint, `<path d="M17 7 7 17"/><path d="M17 17H7V7"/>`)

// ArrowDownNarrowWide is [arrow-down-narrow-wide](https://lucide.dev/icons/arrow-down-narrow-wide).
var ArrowDownNarrowWide = icon.Lucide(viewBox, paint, `<path d="m3 16 4 4 4-4"/><path d="M7 20V4"/><path d="M11 4h4"/><path d="M11 8h7"/><path d="M11 12h10"/>`)

// ArrowDownRight is [arrow-down-right](https://lucide.dev/icons/arrow-down-right).
var ArrowDownRight = icon.Lucide(viewBox, paint, `<path d="m7 7 10 10"/><path d="M17 7v10H7"/>`)

// ArrowDownToDot is [arrow-down-to-dot](https://lucide.dev/icons/arrow-down-to-dot).
var ArrowDownToDot = icon.Lucide(viewBox, paint, `<path d="M12 2v14"/><path d="m19 9-7 7-7-7"/><circle cx="12" cy="21" r="1"/>`)

// ArrowDownToLine is [arrow-down-to-line](https://lucide.dev/icons/arrow-down-to-line).
var ArrowDownToLine = icon.Lucide(viewBox, paint, `<path d="M12 17V3"/><path d="m6 11 6 6 6-6"/><path d="M19 21H5"/>`)

// ArrowDownUp is [arrow-down-up](https://lucide.dev/icons/arrow-down-up).
var ArrowDownUp = icon.Lucide(viewBox, paint, `<path d="m3 16 4 4 4-4"/><path d="M7 20V4"/><path d="m21 8-4-4-4 4"/><path d="M17 4v16"/>`)

// ArrowDownWideNarrow is [arrow-down-wide-narrow](https://lucide.dev/icons/arrow-down-wide-narrow).
var ArrowDownWideNarrow = icon.Lucide(viewBox, paint, `<path d="m3 16 4 4 4-4"/><path d="M7 20V4"/><path d="M11 4h10"/><path d="M11 8h7"/><path d="M11 12h4"/>`)

// ArrowDownZA is [arrow-down-z-a](https://lucide.dev/icons/arrow-down-z-a).
var ArrowDownZA = icon.Lucide(viewBox, paint, `<path d="m3 16 4 4 4-4"/><path d="M7 4v16"/><path d="M15 4h5l-5 6h5"/><path d="M15 20v-3.5a2.5 2.5 0 0 1 5 0V20"/><path d="M20 18h-5"/>`)

// ArrowDown is [arrow-down](https://lucide.dev/icons/arrow-down).
var ArrowDown = icon.Lucide(viewBox, paint, `<path d="M12 5v14"/><path d="m19 12-7 7-7-7"/>`)

// ArrowLeftFromLine is [arrow-left-from-line](https://lucide.dev/icons/arrow-left-from-line).
var ArrowLeftFromLine = icon.Lucide(viewBox, paint, `<path d="m9 6-6 6 6 6"/><path d="M3 12h14"/><path d="M21 19V5"/>`)

// ArrowLeftRight is [arrow-left-right](https://lucide.dev/icons/arrow-left-right).
var ArrowLeftRight = icon.Lucide(viewBox, paint, `<path d="M8 3 4 7l4 4"/><path d="M4 7h16"/><path d="m16 21 4-4-4-4"/><path d="M20 17H4"/>`)

// ArrowLeftToLine is [arrow-left-to-line](https://lucide.dev/icons/arrow-left-to-line).
var ArrowLeftToLine = icon.Lucide(viewBox, paint, `<path d="M3 19V5"/><path d="m13 6-6 6 6 6"/><path d="M7 12h14"/>`)

// ArrowLeft is [arrow-left](https://lucide.dev/icons/arrow-left).
var ArrowLeft = icon.Lucide(viewBox, paint, `<path d="m12 19-7-7 7-7"/><path d="M19 12H5"/>`)

// ArrowRightFromLine is [arrow-right-from-line](https://lucide.dev/icons/arrow-right-from-line).
var ArrowRightFromLine = icon.Lucide(viewBox, paint, `<path d="M3 5v14"/><path d="M21 12H7"/><path d="m15 18 6-6-6-6"/>`)

// ArrowRightLeft is [arrow-right-left](https://lucide.dev/icons/arrow-right-left).
var ArrowRightLeft = icon.Lucide(viewBox, paint, `<path d="m16 3 4 4-4 4"/><path d="M20 7H4"/><path d="m8 21-4-4 4-4"/><path d="M4 17h16"/>`)

// ArrowRightToLine is [arrow-right-to-line](https://lucide.dev/icons/arrow-right-to-line).
var ArrowRightToLine = icon.Lucide(viewBox, paint, `<path d="M17 12H3"/><path d="m11 18 6-6-6-6"/><path d="M21 5v14"/>`)

// ArrowRight is [arrow-right](https://lucide.dev/icons/arrow-right).
var ArrowRight = icon.Lucide(viewBox, paint, `<path d="M5 12h14"/><path d="m12 5 7 7-7 7"/>`)

// ArrowUp01 is [arrow-up-0-1](https://lucide.dev/icons/arrow-up-0-1).
var ArrowUp01 = icon.Lucide(viewBox, paint, `<path d="m3 8 4-4 4 4"/><path d="M7 4v16"/><rect x="15" y="4" width="4" height="6" ry="2"/><path d="M17 20v-6h-2"/><path d="M15 20h4"/>`)

// ArrowUp10 is [arrow-up-1-0](https://lucide.dev/icons/arrow-up-1-0).
var ArrowUp10 = icon.Lucide(viewBox, paint, `<path d="m3 8 4-4 4 4"/><path d="M7 4v16"/><path d="M17 10V4h-2"/><path d="M15 10h4"/><rect x="15" y="14" width="4" height="6" ry="2"/>`)

// ArrowUpAZ is [arrow-up-a-z](https://lucide.dev/icons/arrow-up-a-z).
var ArrowUpAZ = icon.Lucide(viewBox, paint, `<path d="m3 8 4-4 4 4"/><path d="M7 4v16"/><path d="M20 8h-5"/><path d="M15 10V6.5a2.5 2.5 0 0 1 5 0V10"/><path d="M15 14h5l-5 6h5"/>`)

// ArrowUpDown is [arrow-up-down](https://lucide.dev/icons/arrow-up-down).
var ArrowUpDown = icon.Lucide(viewBox, paint, `<path d="m21 16-4 4-4-4"/><path d="M17 20V4"/><path d="m3 8 4-4 4 4"/><path d="M7 4v16"/>`)

// ArrowUpFromDot is [arrow-up-from-dot](https://lucide.dev/icons/arrow-up-from-dot).
var ArrowUpFromDot = icon.Lucide(viewBox, paint, `<path d="m5 9 7-7 7 7"/><path d="M12 16V2"/><circle cx="12" cy="21" r="1"/>`)

// ArrowUpFromLine is [arrow-up-from-line](https://lucide.dev/icons/arrow-up-from-line).
var ArrowUpFromLine = icon.Lucide(viewBox, paint, `<path d="m18 9-6-6-6 6"/><path d="M12 3v14"/><path d="M5 21h14"/>`)

// ArrowUpLeft is [arrow-up-left](https://lucide.dev/icons/arrow-up-left).
var ArrowUpLeft = icon.Lucide(viewBox, paint, `<path d="M7 17V7h10"/><path d="M17 17 7 7"/>`)

// ArrowUpNarrowWide is [arrow-up-narrow-wide](https://lucide.dev/icons/arrow-up-narrow-wide).
var ArrowUpNarrowWide = icon.Lucide(viewBox, paint, `<path d="m3 8 4-4 4 4"/><path d="M7 4v16"/><path d="M11 12h4"/><path d="M11 16h7"/><path d="M11 20h10"/>`)

// ArrowUpRight is [arrow-up-right](https://lucide.dev/icons/arrow-up-right).
var ArrowUpRight = icon.Lucide(viewBox, paint, `<path d="M7 7h10v10"/><path d="M7 17 17 7"/>`)

// ArrowUpToLine is [arrow-up-to-line](https://lucide.dev/icons/arrow-up-to-line).
var ArrowUpToLine = icon.Lucide(viewBox, paint, `<path d="M5 3h14"/><path d="m18 13-6-6-6 6"/><path d="M12 7v14"/>`)

// ArrowUpWideNarrow is [arrow-up-wide-narrow](https://lucide.dev/icons/arrow-up-wide-narrow).
var ArrowUpWideNarrow = icon.Lucide(viewBox, paint, `<path d="m3 8 4-4 4 4"/><path d="M7 4v16"/><path d="M11 12h10"/><path d="M11 16h7"/><path d="M11 20h4"/>`)

// ArrowUpZA is [arrow-up-z-a](https://lucide.dev/icons/arrow-up-z-a).
var ArrowUpZA = icon.Lucide(viewBox, paint, `<path d="m3 8 4-4 4 4"/><path d="M7 4v16"/><path d="M15 4h5l-5 6h5"/><path d="M15 20v-3.5a2.5 2.5 0 0 1 5 0V20"/><path d="M20 18h-5"/>`)

// ArrowUp is [arrow-up](https://lucide.dev/icons/arrow-up).
var ArrowUp = icon.Lucide(viewBox, paint, `<path d="m5 12 7-7 7 7"/><path d="M12 19V5"/>`)

// ArrowsUpFromLine is [arrows-up-from-line](https://lucide.dev/icons/arrows-up-from-line).
var ArrowsUpFromLine = icon.Lucide(viewBox, paint, `<path d="m4 6 3-3 3 3"/><path d="M7 17V3"/><path d="m14 6 3-3 3 3"/><path d="M17 17V3"/><path d="M4 21h16"/>`)

// Asterisk is [asterisk](https://lucide.dev/icons/asterisk).
var Asterisk = icon.Lucide(viewBox, paint, `<path d="M12 5v14"/><path d="m18.065 8.496-12.125 7"/><path d="m5.94 8.504 12.125 7"/>`)

// Astroid is [astroid](https://lucide.dev/icons/astroid).
var Astroid = icon.Lucide(viewBox, paint, `<path d="M12.983 21.186a1 1 0 0 1-1.966 0 10 10 0 0 0-8.203-8.203 1 1 0 0 1 0-1.966 10 10 0 0 0 8.203-8.203 1 1 0 0 1 1.966 0 10 10 0 0 0 8.203 8.203 1 1 0 0 1 0 1.966 10 10 0 0 0-8.203 8.203"/>`)

// AtSign is [at-sign](https://lucide.dev/icons/at-sign).
var AtSign = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="4"/><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-4 8"/>`)

// Atom is [atom](https://lucide.dev/icons/atom).
var Atom = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="1"/><path d="M20.2 20.2c2.04-2.03.02-7.36-4.5-11.9-4.54-4.52-9.87-6.54-11.9-4.5-2.04 2.03-.02 7.36 4.5 11.9 4.54 4.52 9.87 6.54 11.9 4.5Z"/><path d="M15.7 15.7c4.52-4.54 6.54-9.87 4.5-11.9-2.03-2.04-7.36-.02-11.9 4.5-4.52 4.54-6.54 9.87-4.5 11.9 2.03 2.04 7.36.02 11.9-4.5Z"/>`)

// AudioLinesOff is [audio-lines-off](https://lucide.dev/icons/audio-lines-off).
var AudioLinesOff = icon.Lucide(viewBox, paint, `<path d="M10 10v11"/><path d="M10 3v1.35"/><path d="M14 14v1"/><path d="M14 8v.35"/><path d="M18 5v7.35"/><path d="M2 10v3"/><path d="m2 2 20 20"/><path d="M22 10v3"/><path d="M6 6v11"/>`)

// AudioLinesX is [audio-lines-x](https://lucide.dev/icons/audio-lines-x).
var AudioLinesX = icon.Lucide(viewBox, paint, `<path d="M10 3v18"/><path d="M14 8v6.35"/><path d="m17 17 5 5"/><path d="M18 5v8.1"/><path d="M2 10v3"/><path d="M22 10v3"/><path d="m22 17-5 5"/><path d="M6 6v11"/>`)

// AudioLines is [audio-lines](https://lucide.dev/icons/audio-lines).
var AudioLines = icon.Lucide(viewBox, paint, `<path d="M2 10v3"/><path d="M6 6v11"/><path d="M10 3v18"/><path d="M14 8v7"/><path d="M18 5v13"/><path d="M22 10v3"/>`)

// AudioWaveform is [audio-waveform](https://lucide.dev/icons/audio-waveform).
var AudioWaveform = icon.Lucide(viewBox, paint, `<path d="M2 13a2 2 0 0 0 2-2V7a2 2 0 0 1 4 0v13a2 2 0 0 0 4 0V4a2 2 0 0 1 4 0v13a2 2 0 0 0 4 0v-4a2 2 0 0 1 2-2"/>`)

// Avocado is [avocado](https://lucide.dev/icons/avocado).
var Avocado = icon.Lucide(viewBox, paint, `<path d="M13.465 3.464a5 5 0 017.07 0 5 5 0 010 7.072c-1.204 1.203-2.38 2.89-2.614 4.577A8 8 0 118.886 6.08c1.687-.235 3.375-1.41 4.579-2.615"/><circle cx="10" cy="14" r="3"/>`)

// Award is [award](https://lucide.dev/icons/award).
var Award = icon.Lucide(viewBox, paint, `<path d="m15.477 12.89 1.515 8.526a.5.5 0 0 1-.81.47l-3.58-2.687a1 1 0 0 0-1.197 0l-3.586 2.686a.5.5 0 0 1-.81-.469l1.514-8.526"/><circle cx="12" cy="8" r="6"/>`)

// Axe is [axe](https://lucide.dev/icons/axe).
var Axe = icon.Lucide(viewBox, paint, `<path d="m14 12-8.381 8.38a1 1 0 0 1-3.001-3L11 9"/><path d="M15 15.5a.5.5 0 0 0 .5.5A6.5 6.5 0 0 0 22 9.5a.5.5 0 0 0-.5-.5h-1.672a2 2 0 0 1-1.414-.586l-5.062-5.062a1.205 1.205 0 0 0-1.704 0L9.352 5.648a1.205 1.205 0 0 0 0 1.704l5.062 5.062A2 2 0 0 1 15 13.828z"/>`)

// Axis3d is [axis-3d](https://lucide.dev/icons/axis-3d).
var Axis3d = icon.Lucide(viewBox, paint, `<path d="M13.5 10.5 15 9"/><path d="M4 4v15a1 1 0 0 0 1 1h15"/><path d="M4.293 19.707 6 18"/><path d="m9 15 1.5-1.5"/>`)

// Baby is [baby](https://lucide.dev/icons/baby).
var Baby = icon.Lucide(viewBox, paint, `<path d="M10 16c.5.3 1.2.5 2 .5s1.5-.2 2-.5"/><path d="M15 12h.01"/><path d="M19.38 6.813A9 9 0 0 1 20.8 10.2a2 2 0 0 1 0 3.6 9 9 0 0 1-17.6 0 2 2 0 0 1 0-3.6A9 9 0 0 1 12 3c2 0 3.5 1.1 3.5 2.5s-.9 2.5-2 2.5c-.8 0-1.5-.4-1.5-1"/><path d="M9 12h.01"/>`)

// Backpack is [backpack](https://lucide.dev/icons/backpack).
var Backpack = icon.Lucide(viewBox, paint, `<path d="M4 10a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M8 10h8"/><path d="M8 18h8"/><path d="M8 22v-6a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v6"/><path d="M9 6V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2"/>`)

// Bacon is [bacon](https://lucide.dev/icons/bacon).
var Bacon = icon.Lucide(viewBox, paint, `<path d="M2.795 16.804c-.789-.79-1.09-2.018-.442-2.928 1.933-2.713 4.414-1.05 6.145-5.378 1.868-4.67 4.608-2.366 6.592-6.107.218-.412.78-.53 1.11-.2l5.005 5.005c.789.79 1.09 2.018.442 2.928-1.933 2.713-4.414 1.05-6.145 5.378-1.868 4.67-4.608 2.366-6.591 6.107-.22.412-.78.53-1.11.2z"/><path d="M4.5 18.5c2-5 5-2 7-7s5-2 7-7"/>`)

// BadgeAlert is [badge-alert](https://lucide.dev/icons/badge-alert).
var BadgeAlert = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12.01" y1="16" y2="16"/>`)

// BadgeCent is [badge-cent](https://lucide.dev/icons/badge-cent).
var BadgeCent = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M12 7v10"/><path d="M15.4 10a4 4 0 1 0 0 4"/>`)

// BadgeCheck is [badge-check](https://lucide.dev/icons/badge-check).
var BadgeCheck = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="m16 9-5.5 5.5L8 12"/>`)

// BadgeDollarSign is [badge-dollar-sign](https://lucide.dev/icons/badge-dollar-sign).
var BadgeDollarSign = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8"/><path d="M12 18V6"/>`)

// BadgeEuro is [badge-euro](https://lucide.dev/icons/badge-euro).
var BadgeEuro = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M7 12h5"/><path d="M15 9.4a4 4 0 1 0 0 5.2"/>`)

// BadgeIndianRupee is [badge-indian-rupee](https://lucide.dev/icons/badge-indian-rupee).
var BadgeIndianRupee = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M8 8h8"/><path d="M8 12h8"/><path d="m13 17-5-1h1a4 4 0 0 0 0-8"/>`)

// BadgeInfo is [badge-info](https://lucide.dev/icons/badge-info).
var BadgeInfo = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><line x1="12" x2="12" y1="16" y2="12"/><line x1="12" x2="12.01" y1="8" y2="8"/>`)

// BadgeJapaneseYen is [badge-japanese-yen](https://lucide.dev/icons/badge-japanese-yen).
var BadgeJapaneseYen = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="m9 8 3 3v7"/><path d="m12 11 3-3"/><path d="M9 12h6"/><path d="M9 16h6"/>`)

// BadgeMinus is [badge-minus](https://lucide.dev/icons/badge-minus).
var BadgeMinus = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><line x1="8" x2="16" y1="12" y2="12"/>`)

// BadgePercent is [badge-percent](https://lucide.dev/icons/badge-percent).
var BadgePercent = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="m15 9-6 6"/><path d="M9 9h.01"/><path d="M15 15h.01"/>`)

// BadgePlus is [badge-plus](https://lucide.dev/icons/badge-plus).
var BadgePlus = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><line x1="12" x2="12" y1="8" y2="16"/><line x1="8" x2="16" y1="12" y2="12"/>`)

// BadgePoundSterling is [badge-pound-sterling](https://lucide.dev/icons/badge-pound-sterling).
var BadgePoundSterling = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M8 12h4"/><path d="M10 16V9.5a2.5 2.5 0 0 1 5 0"/><path d="M8 16h7"/>`)

// BadgeQuestionMark is [badge-question-mark](https://lucide.dev/icons/badge-question-mark).
var BadgeQuestionMark = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><line x1="12" x2="12.01" y1="17" y2="17"/>`)

// BadgeRussianRuble is [badge-russian-ruble](https://lucide.dev/icons/badge-russian-ruble).
var BadgeRussianRuble = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M9 16h5"/><path d="M9 12h5a2 2 0 1 0 0-4h-3v9"/>`)

// BadgeSwissFranc is [badge-swiss-franc](https://lucide.dev/icons/badge-swiss-franc).
var BadgeSwissFranc = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><path d="M11 17V8h4"/><path d="M11 12h3"/><path d="M9 16h4"/>`)

// BadgeTurkishLira is [badge-turkish-lira](https://lucide.dev/icons/badge-turkish-lira).
var BadgeTurkishLira = icon.Lucide(viewBox, paint, `<path d="M11 7v10a5 5 0 0 0 5-5"/><path d="m15 8-6 3"/><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76"/>`)

// BadgeX is [badge-x](https://lucide.dev/icons/badge-x).
var BadgeX = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/><line x1="15" x2="9" y1="9" y2="15"/><line x1="9" x2="15" y1="9" y2="15"/>`)

// Badge is [badge](https://lucide.dev/icons/badge).
var Badge = icon.Lucide(viewBox, paint, `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"/>`)

// BaggageClaim is [baggage-claim](https://lucide.dev/icons/baggage-claim).
var BaggageClaim = icon.Lucide(viewBox, paint, `<path d="M22 18H6a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2"/><path d="M17 14V4a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v10"/><rect width="13" height="8" x="8" y="6" rx="1"/><circle cx="18" cy="20" r="2"/><circle cx="9" cy="20" r="2"/>`)

// Balloon is [balloon](https://lucide.dev/icons/balloon).
var Balloon = icon.Lucide(viewBox, paint, `<path d="M12 16v1a2 2 0 0 0 2 2h1a2 2 0 0 1 2 2v1"/><path d="M12 6a2 2 0 0 1 2 2"/><path d="M18 8c0 4-3.5 8-6 8s-6-4-6-8a6 6 0 0 1 12 0"/>`)

// Ban is [ban](https://lucide.dev/icons/ban).
var Ban = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M4.929 4.929 19.07 19.071"/>`)

// Banana is [banana](https://lucide.dev/icons/banana).
var Banana = icon.Lucide(viewBox, paint, `<path d="M4 13c3.5-2 8-2 10 2a5.5 5.5 0 0 1 8 5"/><path d="M5.15 17.89c5.52-1.52 8.65-6.89 7-12C11.55 4 11.5 2 13 2c3.22 0 5 5.5 5 8 0 6.5-4.2 12-10.49 12C5.11 22 2 22 2 20c0-1.5 1.14-1.55 3.15-2.11Z"/>`)

// Bandage is [bandage](https://lucide.dev/icons/bandage).
var Bandage = icon.Lucide(viewBox, paint, `<path d="M10 10.01h.01"/><path d="M10 14.01h.01"/><path d="M14 10.01h.01"/><path d="M14 14.01h.01"/><path d="M18 6v12"/><path d="M6 6v12"/><rect x="2" y="6" width="20" height="12" rx="2"/>`)

// BangladeshiTaka is [bangladeshi-taka](https://lucide.dev/icons/bangladeshi-taka).
var BangladeshiTaka = icon.Lucide(viewBox, paint, `<path d="M6 5a2 2 0 0 1 4 0v12a4 4 0 0 0 8 0 2 2 0 0 0-4 0"/><path d="M6 9h12"/>`)

// BanknoteArrowDown is [banknote-arrow-down](https://lucide.dev/icons/banknote-arrow-down).
var BanknoteArrowDown = icon.Lucide(viewBox, paint, `<path d="M12 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5"/><path d="m16 19 3 3 3-3"/><path d="M18 12h.01"/><path d="M19 16v6"/><path d="M6 12h.01"/><circle cx="12" cy="12" r="2"/>`)

// BanknoteArrowUp is [banknote-arrow-up](https://lucide.dev/icons/banknote-arrow-up).
var BanknoteArrowUp = icon.Lucide(viewBox, paint, `<path d="M12 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5"/><path d="M18 12h.01"/><path d="M19 22v-6"/><path d="m22 19-3-3-3 3"/><path d="M6 12h.01"/><circle cx="12" cy="12" r="2"/>`)

// BanknoteCheck is [banknote-check](https://lucide.dev/icons/banknote-check).
var BanknoteCheck = icon.Lucide(viewBox, paint, `<path d="M11.748 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4.875"/><path d="m16 19 2 2 4-4"/><path d="M18 12h.01"/><path d="M6 12h.01"/><circle cx="12" cy="12" r="2"/>`)

// BanknoteX is [banknote-x](https://lucide.dev/icons/banknote-x).
var BanknoteX = icon.Lucide(viewBox, paint, `<path d="M13 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5"/><path d="m17 17 5 5"/><path d="M18 12h.01"/><path d="m22 17-5 5"/><path d="M6 12h.01"/><circle cx="12" cy="12" r="2"/>`)

// Banknote is [banknote](https://lucide.dev/icons/banknote).
var Banknote = icon.Lucide(viewBox, paint, `<rect width="20" height="12" x="2" y="6" rx="2"/><circle cx="12" cy="12" r="2"/><path d="M6 12h.01M18 12h.01"/>`)

// Barcode is [barcode](https://lucide.dev/icons/barcode).
var Barcode = icon.Lucide(viewBox, paint, `<path d="M3 5v14"/><path d="M8 5v14"/><path d="M12 5v14"/><path d="M17 5v14"/><path d="M21 5v14"/>`)

// Barrel is [barrel](https://lucide.dev/icons/barrel).
var Barrel = icon.Lucide(viewBox, paint, `<path d="M10.5 2a34.5 34.5 0 000 20"/><path d="M13.5 2a34.5 34.5 0 010 20"/><path d="M4.67 17h14.67"/><path d="M4.67 7h14.67"/><path d="M6.23 3.06a19 19 0 000 17.88A2 2 0 008 22h8a2 2 0 001.77-1.06 19 19 0 000-17.88A2 2 0 0016 2H8a2 2 0 00-1.77 1.06"/>`)

// Baseline is [baseline](https://lucide.dev/icons/baseline).
var Baseline = icon.Lucide(viewBox, paint, `<path d="M4 20h16"/><path d="m6 16 6-12 6 12"/><path d="M8 12h8"/>`)

// Bath is [bath](https://lucide.dev/icons/bath).
var Bath = icon.Lucide(viewBox, paint, `<path d="M10 4 8 6"/><path d="M17 19v2"/><path d="M2 12h20"/><path d="M7 19v2"/><path d="M9 5 7.621 3.621A2.121 2.121 0 0 0 4 5v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5"/>`)

// BatteryCharging is [battery-charging](https://lucide.dev/icons/battery-charging).
var BatteryCharging = icon.Lucide(viewBox, paint, `<path d="m11 7-3 5h4l-3 5"/><path d="M14.856 6H16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2.935"/><path d="M22 14v-4"/><path d="M5.14 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2.936"/>`)

// BatteryFull is [battery-full](https://lucide.dev/icons/battery-full).
var BatteryFull = icon.Lucide(viewBox, paint, `<path d="M10 10v4"/><path d="M14 10v4"/><path d="M22 14v-4"/><path d="M6 10v4"/><rect x="2" y="6" width="16" height="12" rx="2"/>`)

// BatteryLow is [battery-low](https://lucide.dev/icons/battery-low).
var BatteryLow = icon.Lucide(viewBox, paint, `<path d="M22 14v-4"/><path d="M6 14v-4"/><rect x="2" y="6" width="16" height="12" rx="2"/>`)

// BatteryMedium is [battery-medium](https://lucide.dev/icons/battery-medium).
var BatteryMedium = icon.Lucide(viewBox, paint, `<path d="M10 14v-4"/><path d="M22 14v-4"/><path d="M6 14v-4"/><rect x="2" y="6" width="16" height="12" rx="2"/>`)

// BatteryPlus is [battery-plus](https://lucide.dev/icons/battery-plus).
var BatteryPlus = icon.Lucide(viewBox, paint, `<path d="M10 9v6"/><path d="M12.543 6H16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-3.605"/><path d="M22 14v-4"/><path d="M7 12h6"/><path d="M7.606 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3.606"/>`)

// BatteryWarning is [battery-warning](https://lucide.dev/icons/battery-warning).
var BatteryWarning = icon.Lucide(viewBox, paint, `<path d="M10 17h.01"/><path d="M10 7v6"/><path d="M14 6h2a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2"/><path d="M22 14v-4"/><path d="M6 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2"/>`)

// Battery is [battery](https://lucide.dev/icons/battery).
var Battery = icon.Lucide(viewBox, paint, `<path d="M 22 14 L 22 10"/><rect x="2" y="6" width="16" height="12" rx="2"/>`)

// Beaker is [beaker](https://lucide.dev/icons/beaker).
var Beaker = icon.Lucide(viewBox, paint, `<path d="M4.5 3h15"/><path d="M6 3v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V3"/><path d="M6 14h12"/>`)

// BeanOff is [bean-off](https://lucide.dev/icons/bean-off).
var BeanOff = icon.Lucide(viewBox, paint, `<path d="M9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22a13.96 13.96 0 0 0 9.9-4.1"/><path d="M10.75 5.093A6 6 0 0 1 22 8c0 2.411-.61 4.68-1.683 6.66"/><path d="M5.341 10.62a4 4 0 0 0 6.487 1.208M10.62 5.341a4.015 4.015 0 0 1 2.039 2.04"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Bean is [bean](https://lucide.dev/icons/bean).
var Bean = icon.Lucide(viewBox, paint, `<path d="M10.165 6.598C9.954 7.478 9.64 8.36 9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22c7.732 0 14-6.268 14-14a6 6 0 0 0-11.835-1.402Z"/><path d="M5.341 10.62a4 4 0 1 0 5.279-5.28"/>`)

// BedDouble is [bed-double](https://lucide.dev/icons/bed-double).
var BedDouble = icon.Lucide(viewBox, paint, `<path d="M2 20v-8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"/><path d="M4 10V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4"/><path d="M12 4v6"/><path d="M2 18h20"/>`)

// BedSingle is [bed-single](https://lucide.dev/icons/bed-single).
var BedSingle = icon.Lucide(viewBox, paint, `<path d="M3 20v-8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8"/><path d="M5 10V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4"/><path d="M3 18h18"/>`)

// Bed is [bed](https://lucide.dev/icons/bed).
var Bed = icon.Lucide(viewBox, paint, `<path d="M2 4v16"/><path d="M2 8h18a2 2 0 0 1 2 2v10"/><path d="M2 17h20"/><path d="M6 8v9"/>`)

// BeefOff is [beef-off](https://lucide.dev/icons/beef-off).
var BeefOff = icon.Lucide(viewBox, paint, `<path d="M11.771 6.109a2.5 2.5 0 0 1 3.12 3.12"/><path d="M17.852 12.185a6.5 6.5 0 0 0-9.035-9.04"/><path d="M18.013 18.013C15.029 20.349 10.831 22 7 22a3 3 0 0 1-2.68-1.66L2.4 16.5"/><path d="m18.5 6 2.19 4.5a6.48 6.48 0 0 1-.139 4.393"/><path d="m2 2 20 20"/><path d="M6.355 6.37a7 7 0 0 0-.075.23c-1.1 3.13-.78 3.9-3.18 6.08A3 3 0 0 0 5 18c3.356 0 6.993-1.267 9.85-3.151"/>`)

// Beef is [beef](https://lucide.dev/icons/beef).
var Beef = icon.Lucide(viewBox, paint, `<path d="M16.4 13.7A6.5 6.5 0 1 0 6.28 6.6c-1.1 3.13-.78 3.9-3.18 6.08A3 3 0 0 0 5 18c4 0 8.4-1.8 11.4-4.3"/><path d="m18.5 6 1.754 3.5a6.48 6.48 0 0 1-1.854 8.2C15.4 20.2 11 22 7 22a3 3 0 0 1-2.68-1.66L2.4 16.5"/><circle cx="12.5" cy="8.5" r="2.5"/>`)

// BeerOff is [beer-off](https://lucide.dev/icons/beer-off).
var BeerOff = icon.Lucide(viewBox, paint, `<path d="M13 13v5"/><path d="M17 11.47V8"/><path d="M17 11h1a3 3 0 0 1 2.745 4.211"/><path d="m2 2 20 20"/><path d="M5 8v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-3"/><path d="M7.536 7.535C6.766 7.649 6.154 8 5.5 8a2.5 2.5 0 0 1-1.768-4.268"/><path d="M8.727 3.204C9.306 2.767 9.885 2 11 2c1.56 0 2 1.5 3 1.5s1.72-.5 2.5-.5a1 1 0 1 1 0 5c-.78 0-1.5-.5-2.5-.5a3.149 3.149 0 0 0-.842.12"/><path d="M9 14.6V18"/>`)

// Beer is [beer](https://lucide.dev/icons/beer).
var Beer = icon.Lucide(viewBox, paint, `<path d="M17 11h1a3 3 0 0 1 0 6h-1"/><path d="M9 12v6"/><path d="M13 12v6"/><path d="M14 7.5c-1 0-1.44.5-3 .5s-2-.5-3-.5-1.72.5-2.5.5a2.5 2.5 0 0 1 0-5c.78 0 1.57.5 2.5.5S9.44 2 11 2s2 1.5 3 1.5 1.72-.5 2.5-.5a2.5 2.5 0 0 1 0 5c-.78 0-1.5-.5-2.5-.5Z"/><path d="M5 8v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V8"/>`)

// BelarusianRuble is [belarusian-ruble](https://lucide.dev/icons/belarusian-ruble).
var BelarusianRuble = icon.Lucide(viewBox, paint, `<path d="M3 16 L13 16"/><path d="M6 12h9a4 4 0 0 1 0 8H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h7"/>`)

// BellCheck is [bell-check](https://lucide.dev/icons/bell-check).
var BellCheck = icon.Lucide(viewBox, paint, `<path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="m15 8 2 2 4-4"/><path d="M16.8607 4.4824A6 6 0 0 0 6 8C6 12.499 4.589 13.956 3.262 15.326"/><path d="M3.262 15.326A1 1 0 0 0 4 17H20A1 1 0 0 0 20.74 15.327C20.209 14.779 19.665 14.218 19.203 13.454"/>`)

// BellDot is [bell-dot](https://lucide.dev/icons/bell-dot).
var BellDot = icon.Lucide(viewBox, paint, `<path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M11.68 2.009A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673c-.824-.85-1.678-1.731-2.21-3.348"/><circle cx="18" cy="5" r="3"/>`)

// BellElectric is [bell-electric](https://lucide.dev/icons/bell-electric).
var BellElectric = icon.Lucide(viewBox, paint, `<path d="M18.518 17.347A7 7 0 0 1 14 19"/><path d="M18.8 4A11 11 0 0 1 20 9"/><path d="M9 9h.01"/><circle cx="20" cy="16" r="2"/><circle cx="9" cy="9" r="7"/><rect x="4" y="16" width="10" height="6" rx="2"/>`)

// BellMinus is [bell-minus](https://lucide.dev/icons/bell-minus).
var BellMinus = icon.Lucide(viewBox, paint, `<path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M15 8h6"/><path d="M16.243 3.757A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673A9.4 9.4 0 0 1 18.667 12"/>`)

// BellOff is [bell-off](https://lucide.dev/icons/bell-off).
var BellOff = icon.Lucide(viewBox, paint, `<path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M17 17H4a1 1 0 0 1-.74-1.673C4.59 13.956 6 12.499 6 8a6 6 0 0 1 .258-1.742"/><path d="m2 2 20 20"/><path d="M8.668 3.01A6 6 0 0 1 18 8c0 2.687.77 4.653 1.707 6.05"/>`)

// BellPlus is [bell-plus](https://lucide.dev/icons/bell-plus).
var BellPlus = icon.Lucide(viewBox, paint, `<path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M15 8h6"/><path d="M18 5v6"/><path d="M20.002 14.464a9 9 0 0 0 .738.863A1 1 0 0 1 20 17H4a1 1 0 0 1-.74-1.673C4.59 13.956 6 12.499 6 8a6 6 0 0 1 8.75-5.332"/>`)

// BellRing is [bell-ring](https://lucide.dev/icons/bell-ring).
var BellRing = icon.Lucide(viewBox, paint, `<path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M22 8c0-2.3-.8-4.3-2-6"/><path d="M3.262 15.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673C19.41 13.956 18 12.499 18 8A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326"/><path d="M4 2C2.8 3.7 2 5.7 2 8"/>`)

// Bell is [bell](https://lucide.dev/icons/bell).
var Bell = icon.Lucide(viewBox, paint, `<path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M3.262 15.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673C19.41 13.956 18 12.499 18 8A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326"/>`)

// BetweenHorizontalEnd is [between-horizontal-end](https://lucide.dev/icons/between-horizontal-end).
var BetweenHorizontalEnd = icon.Lucide(viewBox, paint, `<rect width="13" height="7" x="3" y="3" rx="1"/><path d="m22 15-3-3 3-3"/><rect width="13" height="7" x="3" y="14" rx="1"/>`)

// BetweenHorizontalStart is [between-horizontal-start](https://lucide.dev/icons/between-horizontal-start).
var BetweenHorizontalStart = icon.Lucide(viewBox, paint, `<rect width="13" height="7" x="8" y="3" rx="1"/><path d="m2 9 3 3-3 3"/><rect width="13" height="7" x="8" y="14" rx="1"/>`)

// BetweenVerticalEnd is [between-vertical-end](https://lucide.dev/icons/between-vertical-end).
var BetweenVerticalEnd = icon.Lucide(viewBox, paint, `<rect width="7" height="13" x="3" y="3" rx="1"/><path d="m9 22 3-3 3 3"/><rect width="7" height="13" x="14" y="3" rx="1"/>`)

// BetweenVerticalStart is [between-vertical-start](https://lucide.dev/icons/between-vertical-start).
var BetweenVerticalStart = icon.Lucide(viewBox, paint, `<rect width="7" height="13" x="3" y="8" rx="1"/><path d="m15 2-3 3-3-3"/><rect width="7" height="13" x="14" y="8" rx="1"/>`)

// BicepsFlexed is [biceps-flexed](https://lucide.dev/icons/biceps-flexed).
var BicepsFlexed = icon.Lucide(viewBox, paint, `<path d="M12.409 13.017A5 5 0 0 1 22 15c0 3.866-4 7-9 7-4.077 0-8.153-.82-10.371-2.462-.426-.316-.631-.832-.62-1.362C2.118 12.723 2.627 2 10 2a3 3 0 0 1 3 3 2 2 0 0 1-2 2c-1.105 0-1.64-.444-2-1"/><path d="M15 14a5 5 0 0 0-7.584 2"/><path d="M9.964 6.825C8.019 7.977 9.5 13 8 15"/>`)

// Bike is [bike](https://lucide.dev/icons/bike).
var Bike = icon.Lucide(viewBox, paint, `<circle cx="18.5" cy="17.5" r="3.5"/><circle cx="5.5" cy="17.5" r="3.5"/><circle cx="15" cy="5" r="1"/><path d="M12 17.5V14l-3-3 4-3 2 3h2"/>`)

// Binary is [binary](https://lucide.dev/icons/binary).
var Binary = icon.Lucide(viewBox, paint, `<rect x="14" y="14" width="4" height="6" rx="2"/><rect x="6" y="4" width="4" height="6" rx="2"/><path d="M6 20h4"/><path d="M14 10h4"/><path d="M6 14h2v6"/><path d="M14 4h2v6"/>`)

// Binoculars is [binoculars](https://lucide.dev/icons/binoculars).
var Binoculars = icon.Lucide(viewBox, paint, `<path d="M10 10h4"/><path d="M19 7V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3"/><path d="M20 21a2 2 0 0 0 2-2v-3.851c0-1.39-2-2.962-2-4.829V8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v11a2 2 0 0 0 2 2z"/><path d="M 22 16 L 2 16"/><path d="M4 21a2 2 0 0 1-2-2v-3.851c0-1.39 2-2.962 2-4.829V8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v11a2 2 0 0 1-2 2z"/><path d="M9 7V4a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3"/>`)

// Biohazard is [biohazard](https://lucide.dev/icons/biohazard).
var Biohazard = icon.Lucide(viewBox, paint, `<circle cx="12" cy="11.9" r="2"/><path d="M6.7 3.4c-.9 2.5 0 5.2 2.2 6.7C6.5 9 3.7 9.6 2 11.6"/><path d="m8.9 10.1 1.4.8"/><path d="M17.3 3.4c.9 2.5 0 5.2-2.2 6.7 2.4-1.2 5.2-.6 6.9 1.5"/><path d="m15.1 10.1-1.4.8"/><path d="M16.7 20.8c-2.6-.4-4.6-2.6-4.7-5.3-.2 2.6-2.1 4.8-4.7 5.2"/><path d="M12 13.9v1.6"/><path d="M13.5 5.4c-1-.2-2-.2-3 0"/><path d="M17 16.4c.7-.7 1.2-1.6 1.5-2.5"/><path d="M5.5 13.9c.3.9.8 1.8 1.5 2.5"/>`)

// Bird is [bird](https://lucide.dev/icons/bird).
var Bird = icon.Lucide(viewBox, paint, `<path d="M16 7h.01"/><path d="M3.4 18H12a8 8 0 0 0 8-8V7a4 4 0 0 0-7.28-2.3L2 20"/><path d="m20 7 2 .5-2 .5"/><path d="M10 18v3"/><path d="M14 17.75V21"/><path d="M7 18a6 6 0 0 0 3.84-10.61"/>`)

// Birdhouse is [birdhouse](https://lucide.dev/icons/birdhouse).
var Birdhouse = icon.Lucide(viewBox, paint, `<path d="M12 18v4"/><path d="m17 18 1.956-11.468"/><path d="m3 8 7.82-5.615a2 2 0 0 1 2.36 0L21 8"/><path d="M4 18h16"/><path d="M7 18 5.044 6.532"/><circle cx="12" cy="10" r="2"/>`)

// Bitcoin is [bitcoin](https://lucide.dev/icons/bitcoin).
var Bitcoin = icon.Lucide(viewBox, paint, `<path d="M11.767 19.089c4.924.868 6.14-6.025 1.216-6.894m-1.216 6.894L5.86 18.047m5.908 1.042-.347 1.97m1.563-8.864c4.924.869 6.14-6.025 1.215-6.893m-1.215 6.893-3.94-.694m5.155-6.2L8.29 4.26m5.908 1.042.348-1.97M7.48 20.364l3.126-17.727"/>`)

// Blend is [blend](https://lucide.dev/icons/blend).
var Blend = icon.Lucide(viewBox, paint, `<circle cx="15" cy="9" r="7"/><circle cx="9" cy="15" r="7"/>`)

// Blender is [blender](https://lucide.dev/icons/blender).
var Blender = icon.Lucide(viewBox, paint, `<path d="M8 14a2 2 0 0 0-1.963 1.615l-1.018 5.193A1 1 0 0 0 6 22h12a1 1 0 0 0 .981-1.192l-1.018-5.193A2 2 0 0 0 16 14z"/><path d="m17 2-1 12"/><path d="M8.006 14 7 2"/><path d="M7.565 8.787A5 5 0 0 0 12 8a5 5 0 0 1 4.56-.75"/><path d="M19 2H5a2 2 0 0 0-2 2v5a2 2 0 0 0 .688 1.5"/><path d="M12 18h.01"/>`)

// Blinds is [blinds](https://lucide.dev/icons/blinds).
var Blinds = icon.Lucide(viewBox, paint, `<path d="M3 3h18"/><path d="M20 7H8"/><path d="M20 11H8"/><path d="M10 19h10"/><path d="M8 15h12"/><path d="M4 3v14"/><circle cx="4" cy="19" r="2"/>`)

// Blocks is [blocks](https://lucide.dev/icons/blocks).
var Blocks = icon.Lucide(viewBox, paint, `<path d="M10 22V7a1 1 0 0 0-1-1H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5a1 1 0 0 0-1-1H2"/><rect x="14" y="2" width="8" height="8" rx="1"/>`)

// BluetoothConnected is [bluetooth-connected](https://lucide.dev/icons/bluetooth-connected).
var BluetoothConnected = icon.Lucide(viewBox, paint, `<path d="m7 7 10 10-5 5V2l5 5L7 17"/><line x1="18" x2="21" y1="12" y2="12"/><line x1="3" x2="6" y1="12" y2="12"/>`)

// BluetoothOff is [bluetooth-off](https://lucide.dev/icons/bluetooth-off).
var BluetoothOff = icon.Lucide(viewBox, paint, `<path d="m17 17-5 5V12l-5 5"/><path d="m2 2 20 20"/><path d="M14.5 9.5 17 7l-5-5v4.5"/>`)

// BluetoothSearching is [bluetooth-searching](https://lucide.dev/icons/bluetooth-searching).
var BluetoothSearching = icon.Lucide(viewBox, paint, `<path d="m7 7 10 10-5 5V2l5 5L7 17"/><path d="M20.83 14.83a4 4 0 0 0 0-5.66"/><path d="M18 12h.01"/>`)

// Bluetooth is [bluetooth](https://lucide.dev/icons/bluetooth).
var Bluetooth = icon.Lucide(viewBox, paint, `<path d="m7 7 10 10-5 5V2l5 5L7 17"/>`)

// Bold is [bold](https://lucide.dev/icons/bold).
var Bold = icon.Lucide(viewBox, paint, `<path d="M6 12h9a4 4 0 0 1 0 8H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h7a4 4 0 0 1 0 8"/>`)

// Bolt is [bolt](https://lucide.dev/icons/bolt).
var Bolt = icon.Lucide(viewBox, paint, `<path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><circle cx="12" cy="12" r="4"/>`)

// Bomb is [bomb](https://lucide.dev/icons/bomb).
var Bomb = icon.Lucide(viewBox, paint, `<circle cx="11" cy="13" r="9"/><path d="M14.35 4.65 16.3 2.7a2.41 2.41 0 0 1 3.4 0l1.6 1.6a2.4 2.4 0 0 1 0 3.4l-1.95 1.95"/><path d="m22 2-1.5 1.5"/>`)

// BoneFracture is [bone-fracture](https://lucide.dev/icons/bone-fracture).
var BoneFracture = icon.Lucide(viewBox, paint, `<path d="M14 4.5a1 1 0 0 1 5 0 .5.5 0 0 0 .5.5 1 1 0 0 1 0 5c-.81 0-1.8-.7-2.5 0l-1.958 1.957a.15.15 0 0 1-.252-.072l-.493-2.07a.15.15 0 0 0-.111-.112l-2.072-.494a.15.15 0 0 1-.072-.252L14 7c.7-.7 0-1.69 0-2.5"/><path d="m16 20-1-2"/><path d="m20 16-2-1"/><path d="m4 8 2 1"/><path d="m8 4 1 2"/><path d="M9.698 14.19a.15.15 0 0 0 .112.112l2.074.489a.15.15 0 0 1 .072.252L10 17c-.7.7 0 1.69 0 2.5a1 1 0 0 1-5 0 .495.495 0 0 0-.5-.5 1 1 0 0 1 0-5c.81 0 1.8.7 2.5 0l1.956-1.957a.15.15 0 0 1 .252.072z"/>`)

// Bone is [bone](https://lucide.dev/icons/bone).
var Bone = icon.Lucide(viewBox, paint, `<path d="M17 10c.7-.7 1.69 0 2.5 0a2.5 2.5 0 1 0 0-5 .5.5 0 0 1-.5-.5 2.5 2.5 0 1 0-5 0c0 .81.7 1.8 0 2.5l-7 7c-.7.7-1.69 0-2.5 0a2.5 2.5 0 0 0 0 5c.28 0 .5.22.5.5a2.5 2.5 0 1 0 5 0c0-.81-.7-1.8 0-2.5Z"/>`)

// BookA is [book-a](https://lucide.dev/icons/book-a).
var BookA = icon.Lucide(viewBox, paint, `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="m8 13 4-7 4 7"/><path d="M9.1 11h5.7"/>`)

// BookAlert is [book-alert](https://lucide.dev/icons/book-alert).
var BookAlert = icon.Lucide(viewBox, paint, `<path d="M12 13h.01"/><path d="M12 6v3"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/>`)

// BookAudio is [book-audio](https://lucide.dev/icons/book-audio).
var BookAudio = icon.Lucide(viewBox, paint, `<path d="M12 6v7"/><path d="M16 8v3"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M8 8v3"/>`)

// BookBookmark is [book-bookmark](https://lucide.dev/icons/book-bookmark).
var BookBookmark = icon.Lucide(viewBox, paint, `<path d="M10 2v7.751a.25.25 0 00.407.195l2.28-1.834a.5.5 0 01.627 0l2.28 1.834A.25.25 0 0016 9.751V2"/><path d="M4 19.5v-15A2.5 2.5 0 016.5 2H19a1 1 0 011 1v18a1 1 0 01-1 1H6.5a1 1 0 010-5H20"/>`)

// BookCheck is [book-check](https://lucide.dev/icons/book-check).
var BookCheck = icon.Lucide(viewBox, paint, `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="m9 9.5 2 2 4-4"/>`)

// BookCopy is [book-copy](https://lucide.dev/icons/book-copy).
var BookCopy = icon.Lucide(viewBox, paint, `<path d="M5 7a2 2 0 0 0-2 2v11"/><path d="M5.803 18H5a2 2 0 0 0 0 4h9.5a.5.5 0 0 0 .5-.5V21"/><path d="M9 15V4a2 2 0 0 1 2-2h9.5a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.5.5H11a2 2 0 0 1 0-4h10"/>`)

// BookDashed is [book-dashed](https://lucide.dev/icons/book-dashed).
var BookDashed = icon.Lucide(viewBox, paint, `<path d="M12 17h1.5"/><path d="M12 22h1.5"/><path d="M12 2h1.5"/><path d="M17.5 22H19a1 1 0 0 0 1-1"/><path d="M17.5 2H19a1 1 0 0 1 1 1v1.5"/><path d="M20 14v3h-2.5"/><path d="M20 8.5V10"/><path d="M4 10V8.5"/><path d="M4 19.5V14"/><path d="M4 4.5A2.5 2.5 0 0 1 6.5 2H8"/><path d="M8 22H6.5a1 1 0 0 1 0-5H8"/>`)

// BookDown is [book-down](https://lucide.dev/icons/book-down).
var BookDown = icon.Lucide(viewBox, paint, `<path d="M12 13V7"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="m9 10 3 3 3-3"/>`)

// BookHeadphones is [book-headphones](https://lucide.dev/icons/book-headphones).
var BookHeadphones = icon.Lucide(viewBox, paint, `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M8 12v-2a4 4 0 0 1 8 0v2"/><circle cx="15" cy="12" r="1"/><circle cx="9" cy="12" r="1"/>`)

// BookHeart is [book-heart](https://lucide.dev/icons/book-heart).
var BookHeart = icon.Lucide(viewBox, paint, `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M8.62 9.8A2.25 2.25 0 1 1 12 6.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"/>`)

// BookImage is [book-image](https://lucide.dev/icons/book-image).
var BookImage = icon.Lucide(viewBox, paint, `<path d="m20 13.7-2.1-2.1a2 2 0 0 0-2.8 0L9.7 17"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><circle cx="10" cy="8" r="2"/>`)

// BookKey is [book-key](https://lucide.dev/icons/book-key).
var BookKey = icon.Lucide(viewBox, paint, `<path d="M13 2H6.5A2.5 2.5 0 0 0 4 4.5v15"/><path d="M17 2v6"/><path d="M17 4h2"/><path d="M20 15.2V21a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><circle cx="17" cy="10" r="2"/>`)

// BookLock is [book-lock](https://lucide.dev/icons/book-lock).
var BookLock = icon.Lucide(viewBox, paint, `<path d="M18 6V4a2 2 0 1 0-4 0v2"/><path d="M20 15v6a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H10"/><rect x="12" y="6" width="8" height="5" rx="1"/>`)

// BookMinus is [book-minus](https://lucide.dev/icons/book-minus).
var BookMinus = icon.Lucide(viewBox, paint, `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M9 10h6"/>`)

// BookOpenCheck is [book-open-check](https://lucide.dev/icons/book-open-check).
var BookOpenCheck = icon.Lucide(viewBox, paint, `<path d="M12 5v16"/><path d="m16 12 2 2 4-4"/><path d="M22 6V5a2 2 0 00-1.999-2L16 3.002A5 5 0 0012 5a5 5 0 00-4-2H4a2 2 0 00-2 2v12a2 2 0 001.999 2H8a5 5 0 014 2 5 5 0 014-2h4.001A2 2 0 0022 17v-1.344"/>`)

// BookOpenText is [book-open-text](https://lucide.dev/icons/book-open-text).
var BookOpenText = icon.Lucide(viewBox, paint, `<path d="M12 5v16"/><path d="M16 13h2"/><path d="M16 9h2"/><path d="M20.001 19A2 2 0 0022 17V5a2 2 0 00-1.999-2L16 3.002A5 5 0 0012 5a5 5 0 00-4-2H4a2 2 0 00-2 2v12a2 2 0 001.999 2H8a5 5 0 014 2 5 5 0 014-2z"/><path d="M6 13h2"/><path d="M6 9h2"/>`)

// BookOpen is [book-open](https://lucide.dev/icons/book-open).
var BookOpen = icon.Lucide(viewBox, paint, `<path d="M12 5v16"/><path d="M20.001 19A2 2 0 0022 17V5a2 2 0 00-1.999-2L16 3.002A5 5 0 0012 5a5 5 0 00-4-2H4a2 2 0 00-2 2v12a2 2 0 001.999 2H8a5 5 0 014 2 5 5 0 014-2z"/>`)

// BookPlus is [book-plus](https://lucide.dev/icons/book-plus).
var BookPlus = icon.Lucide(viewBox, paint, `<path d="M12 7v6"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M9 10h6"/>`)

// BookSearch is [book-search](https://lucide.dev/icons/book-search).
var BookSearch = icon.Lucide(viewBox, paint, `<path d="M11 22H5.5a1 1 0 0 1 0-5h4.501"/><path d="m21 22-1.879-1.878"/><path d="M3 19.5v-15A2.5 2.5 0 0 1 5.5 2H18a1 1 0 0 1 1 1v8"/><circle cx="17" cy="18" r="3"/>`)

// BookText is [book-text](https://lucide.dev/icons/book-text).
var BookText = icon.Lucide(viewBox, paint, `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M8 11h8"/><path d="M8 7h6"/>`)

// BookType is [book-type](https://lucide.dev/icons/book-type).
var BookType = icon.Lucide(viewBox, paint, `<path d="M10 13h4"/><path d="M12 6v7"/><path d="M16 8V6H8v2"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/>`)

// BookUp2 is [book-up-2](https://lucide.dev/icons/book-up-2).
var BookUp2 = icon.Lucide(viewBox, paint, `<path d="M12 13V7"/><path d="M18 2h1a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2"/><path d="m9 10 3-3 3 3"/><path d="m9 5 3-3 3 3"/>`)

// BookUp is [book-up](https://lucide.dev/icons/book-up).
var BookUp = icon.Lucide(viewBox, paint, `<path d="M12 13V7"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="m9 10 3-3 3 3"/>`)

// BookUser is [book-user](https://lucide.dev/icons/book-user).
var BookUser = icon.Lucide(viewBox, paint, `<path d="M15 13a3 3 0 1 0-6 0"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><circle cx="12" cy="8" r="2"/>`)

// BookX is [book-x](https://lucide.dev/icons/book-x).
var BookX = icon.Lucide(viewBox, paint, `<path d="m14.5 7.5-5 5"/><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/><path d="m9.5 7.5 5 5"/>`)

// Book is [book](https://lucide.dev/icons/book).
var Book = icon.Lucide(viewBox, paint, `<path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"/>`)

// BookmarkCheck is [bookmark-check](https://lucide.dev/icons/bookmark-check).
var BookmarkCheck = icon.Lucide(viewBox, paint, `<path d="M17 3a2 2 0 0 1 2 2v15a1 1 0 0 1-1.496.868l-4.512-2.578a2 2 0 0 0-1.984 0l-4.512 2.578A1 1 0 0 1 5 20V5a2 2 0 0 1 2-2z"/><path d="m9 10 2 2 4-4"/>`)

// BookmarkMinus is [bookmark-minus](https://lucide.dev/icons/bookmark-minus).
var BookmarkMinus = icon.Lucide(viewBox, paint, `<path d="M15 10H9"/><path d="M17 3a2 2 0 0 1 2 2v15a1 1 0 0 1-1.496.868l-4.512-2.578a2 2 0 0 0-1.984 0l-4.512 2.578A1 1 0 0 1 5 20V5a2 2 0 0 1 2-2z"/>`)

// BookmarkOff is [bookmark-off](https://lucide.dev/icons/bookmark-off).
var BookmarkOff = icon.Lucide(viewBox, paint, `<path d="M19 19v1a1 1 0 0 1-1.496.868l-4.512-2.578a2 2 0 0 0-1.984 0l-4.512 2.578A1 1 0 0 1 5 20V5"/><path d="m2 2 20 20"/><path d="M8.656 3H17a2 2 0 0 1 2 2v8.344"/>`)

// BookmarkPlus is [bookmark-plus](https://lucide.dev/icons/bookmark-plus).
var BookmarkPlus = icon.Lucide(viewBox, paint, `<path d="M12 7v6"/><path d="M15 10H9"/><path d="M17 3a2 2 0 0 1 2 2v15a1 1 0 0 1-1.496.868l-4.512-2.578a2 2 0 0 0-1.984 0l-4.512 2.578A1 1 0 0 1 5 20V5a2 2 0 0 1 2-2z"/>`)

// BookmarkX is [bookmark-x](https://lucide.dev/icons/bookmark-x).
var BookmarkX = icon.Lucide(viewBox, paint, `<path d="m14.5 7.5-5 5"/><path d="M17 3a2 2 0 0 1 2 2v15a1 1 0 0 1-1.496.868l-4.512-2.578a2 2 0 0 0-1.984 0l-4.512 2.578A1 1 0 0 1 5 20V5a2 2 0 0 1 2-2z"/><path d="m9.5 7.5 5 5"/>`)

// Bookmark is [bookmark](https://lucide.dev/icons/bookmark).
var Bookmark = icon.Lucide(viewBox, paint, `<path d="M17 3a2 2 0 0 1 2 2v15a1 1 0 0 1-1.496.868l-4.512-2.578a2 2 0 0 0-1.984 0l-4.512 2.578A1 1 0 0 1 5 20V5a2 2 0 0 1 2-2z"/>`)

// BoomBox is [boom-box](https://lucide.dev/icons/boom-box).
var BoomBox = icon.Lucide(viewBox, paint, `<path d="M4 9V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4"/><path d="M8 8v1"/><path d="M12 8v1"/><path d="M16 8v1"/><rect width="20" height="12" x="2" y="9" rx="2"/><circle cx="8" cy="15" r="2"/><circle cx="16" cy="15" r="2"/>`)

// BotMessageSquare is [bot-message-square](https://lucide.dev/icons/bot-message-square).
var BotMessageSquare = icon.Lucide(viewBox, paint, `<path d="M12 6V2H8"/><path d="M15 11v2"/><path d="M2 12h2"/><path d="M20 12h2"/><path d="M20 16a2 2 0 0 1-2 2H8.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 4 20.286V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2z"/><path d="M9 11v2"/>`)

// BotOff is [bot-off](https://lucide.dev/icons/bot-off).
var BotOff = icon.Lucide(viewBox, paint, `<path d="M13.67 8H18a2 2 0 0 1 2 2v4.33"/><path d="M2 14h2"/><path d="M20 14h2"/><path d="M22 22 2 2"/><path d="M8 8H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 1.414-.586"/><path d="M9 13v2"/><path d="M9.67 4H12v2.33"/>`)

// Bot is [bot](https://lucide.dev/icons/bot).
var Bot = icon.Lucide(viewBox, paint, `<path d="M12 8V4H8"/><rect width="16" height="12" x="4" y="8" rx="2"/><path d="M2 14h2"/><path d="M20 14h2"/><path d="M15 13v2"/><path d="M9 13v2"/>`)

// BottleWine is [bottle-wine](https://lucide.dev/icons/bottle-wine).
var BottleWine = icon.Lucide(viewBox, paint, `<path d="M10 3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a6 6 0 0 0 1.2 3.6l.6.8A6 6 0 0 1 17 13v8a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-8a6 6 0 0 1 1.2-3.6l.6-.8A6 6 0 0 0 10 5z"/><path d="M17 13h-4a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h4"/>`)

// BowArrow is [bow-arrow](https://lucide.dev/icons/bow-arrow).
var BowArrow = icon.Lucide(viewBox, paint, `<path d="M17 3h4v4"/><path d="M18.575 11.082a13 13 0 0 1 1.048 9.027 1.17 1.17 0 0 1-1.914.597L14 17"/><path d="M7 10 3.29 6.29a1.17 1.17 0 0 1 .6-1.91 13 13 0 0 1 9.03 1.05"/><path d="M7 14a1.7 1.7 0 0 0-1.207.5l-2.646 2.646A.5.5 0 0 0 3.5 18H5a1 1 0 0 1 1 1v1.5a.5.5 0 0 0 .854.354L9.5 18.207A1.7 1.7 0 0 0 10 17v-2a1 1 0 0 0-1-1z"/><path d="M9.707 14.293 21 3"/>`)

// Box is [box](https://lucide.dev/icons/box).
var Box = icon.Lucide(viewBox, paint, `<path d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"/><path d="m3.3 7 8.7 5 8.7-5"/><path d="M12 22V12"/>`)

// Boxes is [boxes](https://lucide.dev/icons/boxes).
var Boxes = icon.Lucide(viewBox, paint, `<path d="M2.97 12.92A2 2 0 0 0 2 14.63v3.24a2 2 0 0 0 .97 1.71l3 1.8a2 2 0 0 0 2.06 0L12 19v-5.5l-5-3-4.03 2.42Z"/><path d="m7 16.5-4.74-2.85"/><path d="m7 16.5 5-3"/><path d="M7 16.5v5.17"/><path d="M12 13.5V19l3.97 2.38a2 2 0 0 0 2.06 0l3-1.8a2 2 0 0 0 .97-1.71v-3.24a2 2 0 0 0-.97-1.71L17 10.5l-5 3Z"/><path d="m17 16.5-5-3"/><path d="m17 16.5 4.74-2.85"/><path d="M17 16.5v5.17"/><path d="M7.97 4.42A2 2 0 0 0 7 6.13v4.37l5 3 5-3V6.13a2 2 0 0 0-.97-1.71l-3-1.8a2 2 0 0 0-2.06 0l-3 1.8Z"/><path d="M12 8 7.26 5.15"/><path d="m12 8 4.74-2.85"/><path d="M12 13.5V8"/>`)

// Braces is [braces](https://lucide.dev/icons/braces).
var Braces = icon.Lucide(viewBox, paint, `<path d="M8 3H7a2 2 0 0 0-2 2v5a2 2 0 0 1-2 2 2 2 0 0 1 2 2v5c0 1.1.9 2 2 2h1"/><path d="M16 21h1a2 2 0 0 0 2-2v-5c0-1.1.9-2 2-2a2 2 0 0 1-2-2V5a2 2 0 0 0-2-2h-1"/>`)

// Brackets is [brackets](https://lucide.dev/icons/brackets).
var Brackets = icon.Lucide(viewBox, paint, `<path d="M16 3h3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-3"/><path d="M8 21H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h3"/>`)

// BrainCircuit is [brain-circuit](https://lucide.dev/icons/brain-circuit).
var BrainCircuit = icon.Lucide(viewBox, paint, `<path d="M12 13h4"/><path d="M12 18h6a2 2 0 012 2v1"/><path d="M12 8h8"/><path d="M12 9a4.17 4.17 0 01-3 4"/><path d="M16 8V5a2 2 0 012-2"/><path d="M6 18a4 4 0 01-2-7.465"/><path d="M6.004 5.125a4 4 0 00-2.526 5.77"/><path d="M6.294 6.294A3 3 0 1112 5v13a4 4 0 11-7.967-.518"/><circle cx="16" cy="13" r="0.5" fill="currentColor"/><circle cx="18" cy="3" r="0.5" fill="currentColor"/><circle cx="20" cy="21" r="0.5" fill="currentColor"/><circle cx="20" cy="8" r="0.5" fill="currentColor"/>`)

// BrainCog is [brain-cog](https://lucide.dev/icons/brain-cog).
var BrainCog = icon.Lucide(viewBox, paint, `<path d="m10.852 14.772-.383.923"/><path d="m10.852 9.228-.383-.923"/><path d="m13.148 14.772.382.924"/><path d="m13.53 8.305-.382.923"/><path d="m14.772 10.852.923-.383"/><path d="m14.772 13.148.923.383"/><path d="M17.997 5.125a4 4 0 012.526 5.77"/><path d="M18 18a4 4 0 002-7.464"/><path d="M4.033 17.483A4 4 0 0011.46 20c.184-.31.896-.31 1.076 0a4 4 0 007.43-2.517"/><path d="M6 18a4 4 0 01-2-7.464"/><path d="M6.002 5.125a4 4 0 00-2.525 5.77"/><path d="M6.293 6.294a3 3 0 01-.287-1.098l-.004-.07a3 3 0 015.968-.543c.05.02.057.01.059 0a3 3 0 015.968.542l-.004.071a3 3 0 01-.287 1.098"/><path d="m9.228 10.852-.923-.383"/><path d="m9.228 13.148-.923.383"/><circle cx="12" cy="12" r="3"/>`)

// BrainOff is [brain-off](https://lucide.dev/icons/brain-off).
var BrainOff = icon.Lucide(viewBox, paint, `<path d="M12 18v-6"/><path d="M12 5v1.344"/><path d="M17.707 6.293A3 3 0 1012 5a3 3 0 00-4.124-2.781"/><path d="M17.997 5.125a4 4 0 012.526 5.77"/><path d="M19.644 19.644A4 4 0 0112 18a4 4 0 11-7.967-.517"/><path d="m2 2 20 20"/><path d="M20 10.536a4 4 0 011.533 5.34"/><path d="M3.477 10.895a4 4 0 011.876-5.542"/><path d="M6 18a4 4 0 01-2-7.464"/><path d="M9 13a4.17 4.17 0 002.166-1.5 4 4 0 00.138-.196"/>`)

// Brain is [brain](https://lucide.dev/icons/brain).
var Brain = icon.Lucide(viewBox, paint, `<path d="M12 18V5"/><path d="M15 13a4.17 4.17 0 01-3-4 4.17 4.17 0 01-3 4"/><path d="M17.707 6.294A3 3 0 1012 5a3 3 0 10-5.707 1.294"/><path d="M17.996 5.125a4 4 0 012.526 5.77"/><path d="M18 18a4 4 0 002-7.465"/><path d="M19.967 17.482a4 4 0 11-7.967.517 4 4 0 11-7.967-.517"/><path d="M6 18a4 4 0 01-2-7.465"/><path d="M6.004 5.125a4 4 0 00-2.526 5.77"/>`)

// BrickWallFire is [brick-wall-fire](https://lucide.dev/icons/brick-wall-fire).
var BrickWallFire = icon.Lucide(viewBox, paint, `<path d="M16 3v2.107"/><path d="M17 9c1 3 2.5 3.5 3.5 4.5A5 5 0 0 1 22 17a5 5 0 0 1-10 0c0-.3 0-.6.1-.9a2 2 0 1 0 3.3-2C13 11.5 16 9 17 9"/><path d="M21 8.274V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3.938"/><path d="M3 15h5.253"/><path d="M3 9h8.228"/><path d="M8 15v6"/><path d="M8 3v6"/>`)

// BrickWallShield is [brick-wall-shield](https://lucide.dev/icons/brick-wall-shield).
var BrickWallShield = icon.Lucide(viewBox, paint, `<path d="M12 9v1.258"/><path d="M16 3v5.46"/><path d="M21 9.118V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h5.75"/><path d="M22 17.5c0 2.499-1.75 3.749-3.83 4.474a.5.5 0 0 1-.335-.005c-2.085-.72-3.835-1.97-3.835-4.47V14a.5.5 0 0 1 .5-.499c1 0 2.25-.6 3.12-1.36a.6.6 0 0 1 .76-.001c.875.765 2.12 1.36 3.12 1.36a.5.5 0 0 1 .5.5z"/><path d="M3 15h7"/><path d="M3 9h12.142"/><path d="M8 15v6"/><path d="M8 3v6"/>`)

// BrickWall is [brick-wall](https://lucide.dev/icons/brick-wall).
var BrickWall = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 9v6"/><path d="M16 15v6"/><path d="M16 3v6"/><path d="M3 15h18"/><path d="M3 9h18"/><path d="M8 15v6"/><path d="M8 3v6"/>`)

// Bridge is [bridge](https://lucide.dev/icons/bridge).
var Bridge = icon.Lucide(viewBox, paint, `<path d="M10 9.728V16"/><path d="M14 9.728V16"/><path d="M18 20V4"/><path d="m22 11-4-4A7.5 7.5 0 0 1 6 7l-4 4"/><path d="M22 16H2"/><path d="M6 20V4"/>`)

// BriefcaseBusiness is [briefcase-business](https://lucide.dev/icons/briefcase-business).
var BriefcaseBusiness = icon.Lucide(viewBox, paint, `<path d="M12 12h.01"/><path d="M16 6V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2"/><path d="M22 13a18.15 18.15 0 0 1-20 0"/><rect width="20" height="14" x="2" y="6" rx="2"/>`)

// BriefcaseConveyorBelt is [briefcase-conveyor-belt](https://lucide.dev/icons/briefcase-conveyor-belt).
var BriefcaseConveyorBelt = icon.Lucide(viewBox, paint, `<path d="M10 20v2"/><path d="M14 20v2"/><path d="M18 20v2"/><path d="M21 20H3"/><path d="M6 20v2"/><path d="M8 16V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v12"/><rect x="4" y="6" width="16" height="10" rx="2"/>`)

// BriefcaseMedical is [briefcase-medical](https://lucide.dev/icons/briefcase-medical).
var BriefcaseMedical = icon.Lucide(viewBox, paint, `<path d="M12 11v4"/><path d="M14 13h-4"/><path d="M16 6V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2"/><path d="M18 6v14"/><path d="M6 6v14"/><rect width="20" height="14" x="2" y="6" rx="2"/>`)

// BriefcasePlus is [briefcase-plus](https://lucide.dev/icons/briefcase-plus).
var BriefcasePlus = icon.Lucide(viewBox, paint, `<path d="M13.354 20H4a2 2 0 01-2-2V8a2 2 0 012-2h16a2 2 0 012 2v3.354"/><path d="M16 11.354V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v16"/><path d="M16 17h6"/><path d="M19 14v6"/>`)

// Briefcase is [briefcase](https://lucide.dev/icons/briefcase).
var Briefcase = icon.Lucide(viewBox, paint, `<path d="M16 20V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"/><rect width="20" height="14" x="2" y="6" rx="2"/>`)

// BringToFront is [bring-to-front](https://lucide.dev/icons/bring-to-front).
var BringToFront = icon.Lucide(viewBox, paint, `<rect x="8" y="8" width="8" height="8" rx="2"/><path d="M4 10a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2"/><path d="M14 20a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2"/>`)

// Broccoli is [broccoli](https://lucide.dev/icons/broccoli).
var Broccoli = icon.Lucide(viewBox, paint, `<path d="M10 13a3 3 0 01-2.121-5.121"/><path d="M15.606 14.204c-3.5 1.5-5.899 4.503-8.899 7.503A1 1 0 016 22c-2 0-4-2-4-4a1 1 0 01.293-.707c1.911-1.911 3.823-3.578 5.347-5.441"/><path d="M16.573 14.737A4 4 0 0114 11"/><path d="M7.14 10.907a4 4 0 112.756-7.43A4 4 0 0116.7 4.48a2 2 0 012.82 2.82 4 4 0 011.002 6.805 4 4 0 11-7.51 1.59"/>`)

// BroomSparkles is [broom-sparkles](https://lucide.dev/icons/broom-sparkles).
var BroomSparkles = icon.Lucide(viewBox, paint, `<path d="M11 2v2"/><path d="M12 3h-2"/><path d="M13.5 10.5 22 2"/><path d="M14.734 13.841a2 2 0 00-.314-2.42L12.58 9.58a2 2 0 00-2.421-.314l-7.657 4.461A1 1 0 002.3 15.3l6.403 6.403a1 1 0 001.571-.204z"/><path d="M20 15v4"/><path d="M22 17h-4"/><path d="M4 4v4"/><path d="m5 18 2-2"/><path d="M6 6H2"/><path d="m7.699 10.7 5.602 5.601"/>`)

// Broom is [broom](https://lucide.dev/icons/broom).
var Broom = icon.Lucide(viewBox, paint, `<path d="M13.5 10.5 22 2"/><path d="M14.734 13.841a2 2 0 00-.314-2.42L12.58 9.58a2 2 0 00-2.421-.314l-7.657 4.461A1 1 0 002.3 15.3l6.403 6.403a1 1 0 001.571-.204z"/><path d="m5 18 2-2"/><path d="m7.699 10.7 5.602 5.601"/>`)

// BrushCleaning is [brush-cleaning](https://lucide.dev/icons/brush-cleaning).
var BrushCleaning = icon.Lucide(viewBox, paint, `<path d="m16 22-1-4"/><path d="M19 14a1 1 0 0 0 1-1v-1a2 2 0 0 0-2-2h-3a1 1 0 0 1-1-1V4a2 2 0 0 0-4 0v5a1 1 0 0 1-1 1H6a2 2 0 0 0-2 2v1a1 1 0 0 0 1 1"/><path d="M19 14H5l-1.973 6.767A1 1 0 0 0 4 22h16a1 1 0 0 0 .973-1.233z"/><path d="m8 22 1-4"/>`)

// Brush is [brush](https://lucide.dev/icons/brush).
var Brush = icon.Lucide(viewBox, paint, `<path d="m11 10 3 3"/><path d="M6.5 21A3.5 3.5 0 1 0 3 17.5a2.62 2.62 0 0 1-.708 1.792A1 1 0 0 0 3 21z"/><path d="M9.969 17.031 21.378 5.624a1 1 0 0 0-3.002-3.002L6.967 14.031"/>`)

// Bubbles is [bubbles](https://lucide.dev/icons/bubbles).
var Bubbles = icon.Lucide(viewBox, paint, `<path d="M7.001 15.085A1.5 1.5 0 0 1 9 16.5"/><circle cx="18.5" cy="8.5" r="3.5"/><circle cx="7.5" cy="16.5" r="5.5"/><circle cx="7.5" cy="4.5" r="2.5"/>`)

// BugOff is [bug-off](https://lucide.dev/icons/bug-off).
var BugOff = icon.Lucide(viewBox, paint, `<path d="M12 20v-8"/><path d="M12.656 7H14a4 4 0 0 1 4 4v1.344"/><path d="M14.12 3.88 16 2"/><path d="M17.123 17.123A6 6 0 0 1 6 14v-3a4 4 0 0 1 1.72-3.287"/><path d="m2 2 20 20"/><path d="M21 5a4 4 0 0 1-3.55 3.97"/><path d="M22 13h-3.344"/><path d="M3 21a4 4 0 0 1 3.81-4"/><path d="M3 5a4 4 0 0 0 3.55 3.97"/><path d="M6 13H2"/><path d="m8 2 1.88 1.88"/><path d="M9.712 4.06A3 3 0 0 1 15 6v1.13"/>`)

// BugPlay is [bug-play](https://lucide.dev/icons/bug-play).
var BugPlay = icon.Lucide(viewBox, paint, `<path d="M10 19.655A6 6 0 0 1 6 14v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 3.97"/><path d="M14 15.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z"/><path d="M14.12 3.88 16 2"/><path d="M21 5a4 4 0 0 1-3.55 3.97"/><path d="M3 21a4 4 0 0 1 3.81-4"/><path d="M3 5a4 4 0 0 0 3.55 3.97"/><path d="M6 13H2"/><path d="m8 2 1.88 1.88"/><path d="M9 7.13V6a3 3 0 1 1 6 0v1.13"/>`)

// Bug is [bug](https://lucide.dev/icons/bug).
var Bug = icon.Lucide(viewBox, paint, `<path d="M12 20v-9"/><path d="M14 7a4 4 0 0 1 4 4v3a6 6 0 0 1-12 0v-3a4 4 0 0 1 4-4z"/><path d="M14.12 3.88 16 2"/><path d="M21 21a4 4 0 0 0-3.81-4"/><path d="M21 5a4 4 0 0 1-3.55 3.97"/><path d="M22 13h-4"/><path d="M3 21a4 4 0 0 1 3.81-4"/><path d="M3 5a4 4 0 0 0 3.55 3.97"/><path d="M6 13H2"/><path d="m8 2 1.88 1.88"/><path d="M9 7.13V6a3 3 0 1 1 6 0v1.13"/>`)

// BuildingComplexPlus is [building-complex-plus](https://lucide.dev/icons/building-complex-plus).
var BuildingComplexPlus = icon.Lucide(viewBox, paint, `<path d="M10 12h4"/><path d="M10 21v-3a2 2 0 013.05-1.702"/><path d="M10 8h4"/><path d="M16 19h6"/><path d="M18 7h2a2 2 0 012 2v4.355"/><path d="M19 16v6"/><path d="M6 10H4a2 2 0 00-2 2v7a2 2 0 002 2h8.535"/><path d="M6 21V5a2 2 0 012-2h8a2 2 0 012 2v7.126"/>`)

// BuildingComplex is [building-complex](https://lucide.dev/icons/building-complex).
var BuildingComplex = icon.Lucide(viewBox, paint, `<path d="M10 12h4"/><path d="M10 8h4"/><path d="M14 21v-3a2 2 0 0 0-4 0v3"/><path d="M6 10H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-2"/><path d="M6 21V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v16"/>`)

// Building is [building](https://lucide.dev/icons/building).
var Building = icon.Lucide(viewBox, paint, `<path d="M12 10h.01"/><path d="M12 14h.01"/><path d="M12 6h.01"/><path d="M16 10h.01"/><path d="M16 14h.01"/><path d="M16 6h.01"/><path d="M8 10h.01"/><path d="M8 14h.01"/><path d="M8 6h.01"/><path d="M9 22v-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3"/><rect x="4" y="2" width="16" height="20" rx="2"/>`)

// BusFront is [bus-front](https://lucide.dev/icons/bus-front).
var BusFront = icon.Lucide(viewBox, paint, `<path d="M4 6 2 7"/><path d="M10 6h4"/><path d="m22 7-2-1"/><rect width="16" height="16" x="4" y="3" rx="2"/><path d="M4 11h16"/><path d="M8 15h.01"/><path d="M16 15h.01"/><path d="M6 19v2"/><path d="M18 21v-2"/>`)

// Bus is [bus](https://lucide.dev/icons/bus).
var Bus = icon.Lucide(viewBox, paint, `<path d="M8 6v6"/><path d="M15 6v6"/><path d="M2 12h19.6"/><path d="M18 18h3s.5-1.7.8-2.8c.1-.4.2-.8.2-1.2 0-.4-.1-.8-.2-1.2l-1.4-5C20.1 6.8 19.1 6 18 6H4a2 2 0 0 0-2 2v10h3"/><circle cx="7" cy="18" r="2"/><path d="M9 18h5"/><circle cx="16" cy="18" r="2"/>`)

// CableCar is [cable-car](https://lucide.dev/icons/cable-car).
var CableCar = icon.Lucide(viewBox, paint, `<path d="M10 3h.01"/><path d="M14 2h.01"/><path d="m2 9 20-5"/><path d="M12 12V6.5"/><rect width="16" height="10" x="4" y="12" rx="3"/><path d="M9 12v5"/><path d="M15 12v5"/><path d="M4 17h16"/>`)

// Cable is [cable](https://lucide.dev/icons/cable).
var Cable = icon.Lucide(viewBox, paint, `<path d="M17 19a1 1 0 0 1-1-1v-2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a1 1 0 0 1-1 1z"/><path d="M17 21v-2"/><path d="M19 14V6.5a1 1 0 0 0-7 0v11a1 1 0 0 1-7 0V10"/><path d="M21 21v-2"/><path d="M3 5V3"/><path d="M4 10a2 2 0 0 1-2-2V6a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2z"/><path d="M7 5V3"/>`)

// CakeSlice is [cake-slice](https://lucide.dev/icons/cake-slice).
var CakeSlice = icon.Lucide(viewBox, paint, `<path d="M16 13H3"/><path d="M16 17H3"/><path d="m7.2 7.9-3.388 2.5A2 2 0 0 0 3 12.01V20a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-8.654c0-2-2.44-6.026-6.44-8.026a1 1 0 0 0-1.082.057L10.4 5.6"/><circle cx="9" cy="7" r="2"/>`)

// Cake is [cake](https://lucide.dev/icons/cake).
var Cake = icon.Lucide(viewBox, paint, `<path d="M20 21v-8a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v8"/><path d="M4 16s.5-1 2-1 2.5 2 4 2 2.5-2 4-2 2.5 2 4 2 2-1 2-1"/><path d="M2 21h20"/><path d="M7 8v3"/><path d="M12 8v3"/><path d="M17 8v3"/><path d="M7 4h.01"/><path d="M12 4h.01"/><path d="M17 4h.01"/>`)

// Calculator is [calculator](https://lucide.dev/icons/calculator).
var Calculator = icon.Lucide(viewBox, paint, `<rect width="16" height="20" x="4" y="2" rx="2"/><line x1="8" x2="16" y1="6" y2="6"/><line x1="16" x2="16" y1="14" y2="18"/><path d="M16 10h.01"/><path d="M12 10h.01"/><path d="M8 10h.01"/><path d="M12 14h.01"/><path d="M8 14h.01"/><path d="M12 18h.01"/><path d="M8 18h.01"/>`)

// Calendar1 is [calendar-1](https://lucide.dev/icons/calendar-1).
var Calendar1 = icon.Lucide(viewBox, paint, `<path d="M11 13h1v4"/><path d="M16 2v3"/><path d="M3 9h18"/><path d="M8 2v3"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// CalendarArrowDown is [calendar-arrow-down](https://lucide.dev/icons/calendar-arrow-down).
var CalendarArrowDown = icon.Lucide(viewBox, paint, `<path d="m14 17 4 4 4-4"/><path d="M16 2v3"/><path d="M18 13v8"/><path d="M21 10.354V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h7.343"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarArrowUp is [calendar-arrow-up](https://lucide.dev/icons/calendar-arrow-up).
var CalendarArrowUp = icon.Lucide(viewBox, paint, `<path d="m14 17 4-4 4 4"/><path d="M16 2v3"/><path d="M18 21v-8"/><path d="M21 10.343V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h9"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarCheck2 is [calendar-check-2](https://lucide.dev/icons/calendar-check-2).
var CalendarCheck2 = icon.Lucide(viewBox, paint, `<path d="M 19 3 L 5 3"/><path d="M 21 13 L 21 5"/><path d="M 21 5 A2 2 0 0 0 19 3"/><path d="M 3 19 A2 2 0 0 0 5 21"/><path d="M 3 5 L 3 19"/><path d="M 5 3 A2 2 0 0 0 3 5"/><path d="m16 19 2 2 4-4"/><path d="M16 2v3"/><path d="M3 9h18"/><path d="M5 21 L12.5 21"/><path d="M8 2v3"/>`)

// CalendarCheck is [calendar-check](https://lucide.dev/icons/calendar-check).
var CalendarCheck = icon.Lucide(viewBox, paint, `<path d="M8 2v3"/><path d="M16 2v3"/><rect x="3" y="3" width="18" height="18" rx="2"/><path d="M3 9h18"/><path d="m9 15 2 2 4-4"/>`)

// CalendarChevronsLeft is [calendar-chevrons-left](https://lucide.dev/icons/calendar-chevrons-left).
var CalendarChevronsLeft = icon.Lucide(viewBox, paint, `<path d="m11 21-3-3 3-3"/><path d="M16 2v3"/><path d="M21 9H3"/><path d="M3 11.5V5a2 2 0 012-2h14a2 2 0 012 2v14a2 2 0 01-2 2h-4"/><path d="m5 21-3-3 3-3"/><path d="M8 2v3"/>`)

// CalendarChevronsRight is [calendar-chevrons-right](https://lucide.dev/icons/calendar-chevrons-right).
var CalendarChevronsRight = icon.Lucide(viewBox, paint, `<path d="m13 21 3-3-3-3"/><path d="M16 2v3"/><path d="m19 21 3-3-3-3"/><path d="M21 11.5V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h4"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarClock is [calendar-clock](https://lucide.dev/icons/calendar-clock).
var CalendarClock = icon.Lucide(viewBox, paint, `<path d="M16 14v2.2l1.6 1"/><path d="M16 2v3"/><path d="M21 7.338V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h2.338"/><path d="M3 9h5.859"/><path d="M8 2v3"/><circle cx="16" cy="16" r="6"/>`)

// CalendarCog is [calendar-cog](https://lucide.dev/icons/calendar-cog).
var CalendarCog = icon.Lucide(viewBox, paint, `<path d="m15.228 16.852-.923-.383"/><path d="m15.228 19.148-.923.383"/><path d="M16 2v3"/><path d="m16.47 14.305.382.923"/><path d="m16.852 20.772-.383.924"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.696-.382-.924"/><path d="m20.773 16.852.924-.383"/><path d="m20.773 19.148.924.383"/><path d="M21 10.5V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h5.5"/><path d="M3 9h18"/><path d="M8 2v3"/><circle cx="18" cy="18" r="3"/>`)

// CalendarDays is [calendar-days](https://lucide.dev/icons/calendar-days).
var CalendarDays = icon.Lucide(viewBox, paint, `<path d="M8 2v3"/><path d="M16 2v3"/><rect x="3" y="3" width="18" height="18" rx="2"/><path d="M3 9h18"/><path d="M8 13h.01"/><path d="M12 13h.01"/><path d="M16 13h.01"/><path d="M8 17h.01"/><path d="M12 17h.01"/><path d="M16 17h.01"/>`)

// CalendarFold is [calendar-fold](https://lucide.dev/icons/calendar-fold).
var CalendarFold = icon.Lucide(viewBox, paint, `<path d="M16 2v3"/><path d="M21 15V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h10v-5a1 1 0 011-1za2.4 2.4 0 01-.706 1.706l-3.588 3.588A2.4 2.4 0 0115 21"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarHeart is [calendar-heart](https://lucide.dev/icons/calendar-heart).
var CalendarHeart = icon.Lucide(viewBox, paint, `<path d="M12.127 21H5a2 2 0 01-2-2V5a2 2 0 012-2h14a2 2 0 012 2v5.125"/><path d="M14.62 17.8A2.25 2.25 0 1118 14.836a2.25 2.25 0 113.38 2.966l-2.626 2.856a.998.998 0 01-1.507 0z"/><path d="M16 2v3"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarMinus2 is [calendar-minus-2](https://lucide.dev/icons/calendar-minus-2).
var CalendarMinus2 = icon.Lucide(viewBox, paint, `<path d="M8 2v3"/><path d="M16 2v3"/><rect x="3" y="3" width="18" height="18" rx="2"/><path d="M3 9h18"/><path d="M10 15h4"/>`)

// CalendarMinus is [calendar-minus](https://lucide.dev/icons/calendar-minus).
var CalendarMinus = icon.Lucide(viewBox, paint, `<path d="M16 18h6"/><path d="M16 2v3"/><path d="M21 14V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h8.3"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarOff is [calendar-off](https://lucide.dev/icons/calendar-off).
var CalendarOff = icon.Lucide(viewBox, paint, `<path d="M16 2v3"/><path d="m2 2 20 20"/><path d="M21 9h-5.5"/><path d="M3 9h6"/><path d="M3.586 3.586A2 2 0 003 5v14a2 2 0 002 2h14a2 2 0 001.414-.586"/><path d="M8.656 3H19a2 2 0 012 2v10.344"/>`)

// CalendarPlus2 is [calendar-plus-2](https://lucide.dev/icons/calendar-plus-2).
var CalendarPlus2 = icon.Lucide(viewBox, paint, `<path d="M8 2v3"/><path d="M16 2v3"/><rect x="3" y="3" width="18" height="18" rx="2"/><path d="M3 9h18"/><path d="M10 15h4"/><path d="M12 13v4"/>`)

// CalendarPlus is [calendar-plus](https://lucide.dev/icons/calendar-plus).
var CalendarPlus = icon.Lucide(viewBox, paint, `<path d="M16 18h6"/><path d="M16 2v3"/><path d="M19 15v6"/><path d="M21 11.5V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h8.3"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarRange is [calendar-range](https://lucide.dev/icons/calendar-range).
var CalendarRange = icon.Lucide(viewBox, paint, `<rect x="3" y="3" width="18" height="18" rx="2"/><path d="M16 2v3"/><path d="M3 9h18"/><path d="M8 2v3"/><path d="M17 13h-6"/><path d="M13 17H7"/><path d="M7 13h.01"/><path d="M17 17h.01"/>`)

// CalendarSearch is [calendar-search](https://lucide.dev/icons/calendar-search).
var CalendarSearch = icon.Lucide(viewBox, paint, `<path d="M16 2v3"/><path d="M21 10.69V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h7.25"/><path d="m22 21-1.875-1.875"/><path d="M3 9h18"/><path d="M8 2v3"/><circle cx="18" cy="17" r="3"/>`)

// CalendarSync is [calendar-sync](https://lucide.dev/icons/calendar-sync).
var CalendarSync = icon.Lucide(viewBox, paint, `<path d="M11 10v4h4"/><path d="m11 14 1.535-1.605a5 5 0 018 1.5"/><path d="M16 2v3"/><path d="m21 18-1.535 1.605a5 5 0 01-8-1.5"/><path d="M21 22v-4h-4"/><path d="M21 8.517V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h3.517"/><path d="M3 9h4"/><path d="M8 2v3"/>`)

// CalendarX2 is [calendar-x-2](https://lucide.dev/icons/calendar-x-2).
var CalendarX2 = icon.Lucide(viewBox, paint, `<path d="M16 2v3"/><path d="m17 16 5 5"/><path d="m17 21 5-5"/><path d="M21 12V5a2 2 0 00-2-2H5a2 2 0 00-2 2v14a2 2 0 002 2h8"/><path d="M3 9h18"/><path d="M8 2v3"/>`)

// CalendarX is [calendar-x](https://lucide.dev/icons/calendar-x).
var CalendarX = icon.Lucide(viewBox, paint, `<path d="M8 2v3"/><path d="M16 2v3"/><rect x="3" y="3" width="18" height="18" rx="2"/><path d="M3 9h18"/><path d="m14 13-4 4"/><path d="m10 13 4 4"/>`)

// Calendar is [calendar](https://lucide.dev/icons/calendar).
var Calendar = icon.Lucide(viewBox, paint, `<path d="M8 2v3"/><path d="M16 2v3"/><rect x="3" y="3" width="18" height="18" rx="2"/><path d="M3 9h18"/>`)

// Calendars is [calendars](https://lucide.dev/icons/calendars).
var Calendars = icon.Lucide(viewBox, paint, `<path d="M12 2v2"/><path d="M15.726 21.01A2 2 0 0 1 14 22H4a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2"/><path d="M18 2v2"/><path d="M2 13h2"/><path d="M8 8h14"/><rect x="8" y="3" width="14" height="14" rx="2"/>`)

// CameraOff is [camera-off](https://lucide.dev/icons/camera-off).
var CameraOff = icon.Lucide(viewBox, paint, `<path d="M14.564 14.558a3 3 0 1 1-4.122-4.121"/><path d="m2 2 20 20"/><path d="M20 20H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1.997a2 2 0 0 0 .819-.175"/><path d="M9.695 4.024A2 2 0 0 1 10.004 4h3.993a2 2 0 0 1 1.76 1.05l.486.9A2 2 0 0 0 18.003 7H20a2 2 0 0 1 2 2v7.344"/>`)

// Camera is [camera](https://lucide.dev/icons/camera).
var Camera = icon.Lucide(viewBox, paint, `<path d="M13.997 4a2 2 0 0 1 1.76 1.05l.486.9A2 2 0 0 0 18.003 7H20a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1.997a2 2 0 0 0 1.759-1.048l.489-.904A2 2 0 0 1 10.004 4z"/><circle cx="12" cy="13" r="3"/>`)

// CanSoda is [can-soda](https://lucide.dev/icons/can-soda).
var CanSoda = icon.Lucide(viewBox, paint, `<path d="m17 22 1.664-2.496a2 2 0 00.336-1.11V5.606a2 2 0 00-.336-1.11L17 2"/><path d="M18 22H6"/><path d="M18 2H6"/><path d="M5 17h14"/><path d="M5 7h14"/><path d="m7 22-1.664-2.496A2 2 0 015 18.394V5.606a2 2 0 01.336-1.11L7 2"/>`)

// Can is [can](https://lucide.dev/icons/can).
var Can = icon.Lucide(viewBox, paint, `<path d="M21 10.5a9 2.5 0 01-18 0v8a9 2.5 0 0018 0z"/><path d="M21 10.5A9 2.5 25.32 004.59 3.47 9 2.5 25.32 0021 10.5"/><path d="M3 10.5a9 2.5 0 016.527-2.405"/><path d="M9 16.858a31 31 0 006 0"/>`)

// CandyCane is [candy-cane](https://lucide.dev/icons/candy-cane).
var CandyCane = icon.Lucide(viewBox, paint, `<path d="m10.8 5 2.111 4.223"/><path d="M17.75 7 15 2.1"/><path d="m4.874 14.647 2.12 4.24"/><path d="M5.7 21a2 2 0 0 1-3.5-2l8.6-14a6 6 0 0 1 10.4 6 2 2 0 1 1-3.464-2 2 2 0 1 0-3.464-2z"/><path d="m7.906 9.712 2.005 4.411"/>`)

// CandyOff is [candy-off](https://lucide.dev/icons/candy-off).
var CandyOff = icon.Lucide(viewBox, paint, `<path d="M10 10v7.9"/><path d="M11.802 6.145a5 5 0 0 1 6.053 6.053"/><path d="M14 6.1v2.243"/><path d="m15.5 15.571-.964.964a5 5 0 0 1-7.071 0 5 5 0 0 1 0-7.07l.964-.965"/><path d="M16 7V3a1 1 0 0 1 1.707-.707 2.5 2.5 0 0 0 2.152.717 1 1 0 0 1 1.131 1.131 2.5 2.5 0 0 0 .717 2.152A1 1 0 0 1 21 8h-4"/><path d="m2 2 20 20"/><path d="M8 17v4a1 1 0 0 1-1.707.707 2.5 2.5 0 0 0-2.152-.717 1 1 0 0 1-1.131-1.131 2.5 2.5 0 0 0-.717-2.152A1 1 0 0 1 3 16h4"/>`)

// Candy is [candy](https://lucide.dev/icons/candy).
var Candy = icon.Lucide(viewBox, paint, `<path d="M10 7v10.9"/><path d="M14 6.1V17"/><path d="M16 7V3a1 1 0 0 1 1.707-.707 2.5 2.5 0 0 0 2.152.717 1 1 0 0 1 1.131 1.131 2.5 2.5 0 0 0 .717 2.152A1 1 0 0 1 21 8h-4"/><path d="M16.536 7.465a5 5 0 0 0-7.072 0l-2 2a5 5 0 0 0 0 7.07 5 5 0 0 0 7.072 0l2-2a5 5 0 0 0 0-7.07"/><path d="M8 17v4a1 1 0 0 1-1.707.707 2.5 2.5 0 0 0-2.152-.717 1 1 0 0 1-1.131-1.131 2.5 2.5 0 0 0-.717-2.152A1 1 0 0 1 3 16h4"/>`)

// CannabisOff is [cannabis-off](https://lucide.dev/icons/cannabis-off).
var CannabisOff = icon.Lucide(viewBox, paint, `<path d="M12 22v-4c1.5 1.5 3.5 3 6 3 0-1.5-.5-3.5-2-5"/><path d="M13.988 8.327C13.902 6.054 13.365 3.82 12 2a9.3 9.3 0 0 0-1.445 2.9"/><path d="M17.375 11.725C18.882 10.53 21 7.841 21 6c-2.324 0-5.08 1.296-6.662 2.684"/><path d="m2 2 20 20"/><path d="M21.024 15.378A15 15 0 0 0 22 15c-.426-1.279-2.67-2.557-4.25-2.907"/><path d="M6.995 6.992C5.714 6.4 4.29 6 3 6c0 2 2.5 5 4 6-1.5 0-4.5 1.5-5 3 3.5 1.5 6 1 6 1-1.5 1.5-2 3.5-2 5 2.5 0 4.5-1.5 6-3"/>`)

// Cannabis is [cannabis](https://lucide.dev/icons/cannabis).
var Cannabis = icon.Lucide(viewBox, paint, `<path d="M12 22v-4"/><path d="M7 12c-1.5 0-4.5 1.5-5 3 3.5 1.5 6 1 6 1-1.5 1.5-2 3.5-2 5 2.5 0 4.5-1.5 6-3 1.5 1.5 3.5 3 6 3 0-1.5-.5-3.5-2-5 0 0 2.5.5 6-1-.5-1.5-3.5-3-5-3 1.5-1 4-4 4-6-2.5 0-5.5 1.5-7 3 0-2.5-.5-5-2-7-1.5 2-2 4.5-2 7-1.5-1.5-4.5-3-7-3 0 2 2.5 5 4 6"/>`)

// CaptionsOff is [captions-off](https://lucide.dev/icons/captions-off).
var CaptionsOff = icon.Lucide(viewBox, paint, `<path d="M10.5 5H19a2 2 0 0 1 2 2v8.5"/><path d="M17 11h-.5"/><path d="M19 19H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2"/><path d="m2 2 20 20"/><path d="M7 11h4"/><path d="M7 15h2.5"/>`)

// Captions is [captions](https://lucide.dev/icons/captions).
var Captions = icon.Lucide(viewBox, paint, `<rect width="18" height="14" x="3" y="5" rx="2" ry="2"/><path d="M7 15h4M15 15h2M7 11h2M13 11h4"/>`)

// CarBattery is [car-battery](https://lucide.dev/icons/car-battery).
var CarBattery = icon.Lucide(viewBox, paint, `<path d="M14 13h4"/><path d="M16 15v-4"/><path d="M18 5v2"/><path d="M6 13h4"/><path d="M6 5v2"/><rect x="2" y="7" width="20" height="12" rx="2"/>`)

// CarFront is [car-front](https://lucide.dev/icons/car-front).
var CarFront = icon.Lucide(viewBox, paint, `<path d="m21 8-2 2-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10 3 8"/><path d="M7 14h.01"/><path d="M17 14h.01"/><rect width="18" height="8" x="3" y="10" rx="2"/><path d="M5 18v2"/><path d="M19 18v2"/>`)

// CarTaxiFront is [car-taxi-front](https://lucide.dev/icons/car-taxi-front).
var CarTaxiFront = icon.Lucide(viewBox, paint, `<path d="M10 2h4"/><path d="m21 8-2 2-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10 3 8"/><path d="M7 14h.01"/><path d="M17 14h.01"/><rect width="18" height="8" x="3" y="10" rx="2"/><path d="M5 18v2"/><path d="M19 18v2"/>`)

// Car is [car](https://lucide.dev/icons/car).
var Car = icon.Lucide(viewBox, paint, `<path d="M19 17h2c.6 0 1-.4 1-1v-3c0-.9-.7-1.7-1.5-1.9C18.7 10.6 16 10 16 10s-1.3-1.4-2.2-2.3c-.5-.4-1.1-.7-1.8-.7H5c-.6 0-1.1.4-1.4.9l-1.4 2.9A3.7 3.7 0 0 0 2 12v4c0 .6.4 1 1 1h2"/><circle cx="7" cy="17" r="2"/><path d="M9 17h6"/><circle cx="17" cy="17" r="2"/>`)

// Caravan is [caravan](https://lucide.dev/icons/caravan).
var Caravan = icon.Lucide(viewBox, paint, `<path d="M18 19V9a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v8a2 2 0 0 0 2 2h2"/><path d="M2 9h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H2"/><path d="M22 17v1a1 1 0 0 1-1 1H10v-9a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v9"/><circle cx="8" cy="19" r="2"/>`)

// CardSim is [card-sim](https://lucide.dev/icons/card-sim).
var CardSim = icon.Lucide(viewBox, paint, `<path d="M12 14v4"/><path d="M14.172 2a2 2 0 0 1 1.414.586l3.828 3.828A2 2 0 0 1 20 7.828V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2z"/><path d="M8 14h8"/><rect x="8" y="10" width="8" height="8" rx="2"/>`)

// Carrot is [carrot](https://lucide.dev/icons/carrot).
var Carrot = icon.Lucide(viewBox, paint, `<path d="M15 16a1 1 0 0 0-7-7q-4 4-5.987 12.385a.5.5 0 0 0 .602.602Q11 20 15 16l-3-3"/><path d="M15 9q4 4 7 0-3-4-7 0 4-4 0-7-4 3 0 7"/><path d="m8 15-2.58-2.58"/>`)

// CartonOff is [carton-off](https://lucide.dev/icons/carton-off).
var CartonOff = icon.Lucide(viewBox, paint, `<path d="M10 10H5v10a2 2 0 002 2h10a2 2 0 002-2v-1"/><path d="M13 22v-9"/><path d="M13.902 8.245 16 6h-4.343"/><path d="M19 13.343V10a2 2 0 00-.539-1.367L16 6V3a1 1 0 00-1-1H9a1 1 0 00-.857.486"/><path d="m2 2 20 20"/><path d="M7.034 7.034 5.539 8.633A2 2 0 005 10"/>`)

// Carton is [carton](https://lucide.dev/icons/carton).
var Carton = icon.Lucide(viewBox, paint, `<path d="M13 22V10a2 2 0 01.539-1.367L16 6H8L5.539 8.633A2 2 0 005 10v10a2 2 0 002 2h10a2 2 0 002-2V10a2 2 0 00-.539-1.367L16 6V3a1 1 0 00-1-1H9a1 1 0 00-1 1v3"/><path d="M5 10h8"/>`)

// CaseLower is [case-lower](https://lucide.dev/icons/case-lower).
var CaseLower = icon.Lucide(viewBox, paint, `<path d="M10 9v7"/><path d="M14 6v10"/><circle cx="17.5" cy="12.5" r="3.5"/><circle cx="6.5" cy="12.5" r="3.5"/>`)

// CaseSensitive is [case-sensitive](https://lucide.dev/icons/case-sensitive).
var CaseSensitive = icon.Lucide(viewBox, paint, `<path d="m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"/><path d="M22 9v7"/><path d="M3.304 13h6.392"/><circle cx="18.5" cy="12.5" r="3.5"/>`)

// CaseUpper is [case-upper](https://lucide.dev/icons/case-upper).
var CaseUpper = icon.Lucide(viewBox, paint, `<path d="M15 11h4.5a1 1 0 0 1 0 5h-4a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h3a1 1 0 0 1 0 5"/><path d="m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16"/><path d="M3.304 13h6.392"/>`)

// CassetteTape is [cassette-tape](https://lucide.dev/icons/cassette-tape).
var CassetteTape = icon.Lucide(viewBox, paint, `<rect width="20" height="16" x="2" y="4" rx="2"/><circle cx="8" cy="10" r="2"/><path d="M8 12h8"/><circle cx="16" cy="10" r="2"/><path d="m6 20 .7-2.9A1.4 1.4 0 0 1 8.1 16h7.8a1.4 1.4 0 0 1 1.4 1l.7 3"/>`)

// Cast is [cast](https://lucide.dev/icons/cast).
var Cast = icon.Lucide(viewBox, paint, `<path d="M2 8V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6"/><path d="M2 12a9 9 0 0 1 8 8"/><path d="M2 16a5 5 0 0 1 4 4"/><line x1="2" x2="2.01" y1="20" y2="20"/>`)

// Castle is [castle](https://lucide.dev/icons/castle).
var Castle = icon.Lucide(viewBox, paint, `<path d="M10 5V3"/><path d="M14 5V3"/><path d="M15 21v-3a3 3 0 0 0-6 0v3"/><path d="M18 3v8"/><path d="M18 5H6"/><path d="M22 11H2"/><path d="M22 9v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9"/><path d="M6 3v8"/>`)

// Cat is [cat](https://lucide.dev/icons/cat).
var Cat = icon.Lucide(viewBox, paint, `<path d="M12 5c.67 0 1.35.09 2 .26 1.78-2 5.03-2.84 6.42-2.26 1.4.58-.42 7-.42 7 .57 1.07 1 2.24 1 3.44C21 17.9 16.97 21 12 21s-9-3-9-7.56c0-1.25.5-2.4 1-3.44 0 0-1.89-6.42-.5-7 1.39-.58 4.72.23 6.5 2.23A9.04 9.04 0 0 1 12 5Z"/><path d="M8 14v.5"/><path d="M16 14v.5"/><path d="M11.25 16.25h1.5L12 17l-.75-.75Z"/>`)

// CctvOff is [cctv-off](https://lucide.dev/icons/cctv-off).
var CctvOff = icon.Lucide(viewBox, paint, `<path d="m12.309 6.652 4.797 2.401a1 1 0 0 1 .447 1.341l-.501 1.001.605.605h2.725a1 1 0 0 1 .894 1.447l-.724 1.448"/><path d="m15.166 15.166-.719 1.439a1 1 0 0 1-1.342.447L3.61 12.3a2.92 2.92 0 0 1-1.3-3.91L3.69 5.6a2.9 2.9 0 0 1 .873-1.037"/><path d="M2 19h3.76a2 2 0 0 0 1.8-1.1l1.441-2.902"/><path d="m2 2 20 20"/><path d="M2 21v-4"/><path d="M7 9h.01"/>`)

// Cctv is [cctv](https://lucide.dev/icons/cctv).
var Cctv = icon.Lucide(viewBox, paint, `<path d="M16.75 12h3.632a1 1 0 0 1 .894 1.447l-2.034 4.069a1 1 0 0 1-1.708.134l-2.124-2.97"/><path d="M17.106 9.053a1 1 0 0 1 .447 1.341l-3.106 6.211a1 1 0 0 1-1.342.447L3.61 12.3a2.92 2.92 0 0 1-1.3-3.91L3.69 5.6a2.92 2.92 0 0 1 3.92-1.3z"/><path d="M2 19h3.76a2 2 0 0 0 1.8-1.1L9 15"/><path d="M2 21v-4"/><path d="M7 9h.01"/>`)

// Certificate is [certificate](https://lucide.dev/icons/certificate).
var Certificate = icon.Lucide(viewBox, paint, `<path d="M12 10H6"/><path d="m20.69 15.479 1.29 4.88a.5.5 0 01-.698.591l-1.843-.849a1 1 0 00-.879.001l-1.846.85a.5.5 0 01-.692-.593l1.29-4.88"/><path d="M22 6.675V6a2 2 0 00-2-2H4a2 2 0 00-2 2v12a2 2 0 002 2h8"/><path d="M6 14h4"/><circle cx="19" cy="13" r="3"/>`)

// ChartArea is [chart-area](https://lucide.dev/icons/chart-area).
var ChartArea = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M7 11.207a.5.5 0 0 1 .146-.353l2-2a.5.5 0 0 1 .708 0l3.292 3.292a.5.5 0 0 0 .708 0l4.292-4.292a.5.5 0 0 1 .854.353V16a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1z"/>`)

// ChartBarBig is [chart-bar-big](https://lucide.dev/icons/chart-bar-big).
var ChartBarBig = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><rect x="7" y="13" width="9" height="4" rx="1"/><rect x="7" y="5" width="12" height="4" rx="1"/>`)

// ChartBarDecreasing is [chart-bar-decreasing](https://lucide.dev/icons/chart-bar-decreasing).
var ChartBarDecreasing = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M7 11h8"/><path d="M7 16h3"/><path d="M7 6h12"/>`)

// ChartBarIncreasing is [chart-bar-increasing](https://lucide.dev/icons/chart-bar-increasing).
var ChartBarIncreasing = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M7 11h8"/><path d="M7 16h12"/><path d="M7 6h3"/>`)

// ChartBarStacked is [chart-bar-stacked](https://lucide.dev/icons/chart-bar-stacked).
var ChartBarStacked = icon.Lucide(viewBox, paint, `<path d="M11 13v4"/><path d="M15 5v4"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/><rect x="7" y="13" width="9" height="4" rx="1"/><rect x="7" y="5" width="12" height="4" rx="1"/>`)

// ChartBar is [chart-bar](https://lucide.dev/icons/chart-bar).
var ChartBar = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M7 16h8"/><path d="M7 11h12"/><path d="M7 6h3"/>`)

// ChartCandlestick is [chart-candlestick](https://lucide.dev/icons/chart-candlestick).
var ChartCandlestick = icon.Lucide(viewBox, paint, `<path d="M9 5v4"/><rect width="4" height="6" x="7" y="9" rx="1"/><path d="M9 15v2"/><path d="M17 3v2"/><rect width="4" height="8" x="15" y="5" rx="1"/><path d="M17 13v3"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/>`)

// ChartColumnBig is [chart-column-big](https://lucide.dev/icons/chart-column-big).
var ChartColumnBig = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><rect x="15" y="5" width="4" height="12" rx="1"/><rect x="7" y="8" width="4" height="9" rx="1"/>`)

// ChartColumnDecreasing is [chart-column-decreasing](https://lucide.dev/icons/chart-column-decreasing).
var ChartColumnDecreasing = icon.Lucide(viewBox, paint, `<path d="M13 17V9"/><path d="M18 17v-3"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M8 17V5"/>`)

// ChartColumnIncreasing is [chart-column-increasing](https://lucide.dev/icons/chart-column-increasing).
var ChartColumnIncreasing = icon.Lucide(viewBox, paint, `<path d="M13 17V9"/><path d="M18 17V5"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M8 17v-3"/>`)

// ChartColumnStacked is [chart-column-stacked](https://lucide.dev/icons/chart-column-stacked).
var ChartColumnStacked = icon.Lucide(viewBox, paint, `<path d="M11 13H7"/><path d="M19 9h-4"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/><rect x="15" y="5" width="4" height="12" rx="1"/><rect x="7" y="8" width="4" height="9" rx="1"/>`)

// ChartColumn is [chart-column](https://lucide.dev/icons/chart-column).
var ChartColumn = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M18 17V9"/><path d="M13 17V5"/><path d="M8 17v-3"/>`)

// ChartGantt is [chart-gantt](https://lucide.dev/icons/chart-gantt).
var ChartGantt = icon.Lucide(viewBox, paint, `<path d="M10 6h8"/><path d="M12 16h6"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M8 11h7"/>`)

// ChartLine is [chart-line](https://lucide.dev/icons/chart-line).
var ChartLine = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="m19 9-5 5-4-4-3 3"/>`)

// ChartNetwork is [chart-network](https://lucide.dev/icons/chart-network).
var ChartNetwork = icon.Lucide(viewBox, paint, `<path d="m13.11 7.664 1.78 2.672"/><path d="m14.162 12.788-3.324 1.424"/><path d="m20 4-6.06 1.515"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/><circle cx="12" cy="6" r="2"/><circle cx="16" cy="12" r="2"/><circle cx="9" cy="15" r="2"/>`)

// ChartNoAxesColumnDecreasing is [chart-no-axes-column-decreasing](https://lucide.dev/icons/chart-no-axes-column-decreasing).
var ChartNoAxesColumnDecreasing = icon.Lucide(viewBox, paint, `<path d="M5 21V3"/><path d="M12 21V9"/><path d="M19 21v-6"/>`)

// ChartNoAxesColumnIncreasing is [chart-no-axes-column-increasing](https://lucide.dev/icons/chart-no-axes-column-increasing).
var ChartNoAxesColumnIncreasing = icon.Lucide(viewBox, paint, `<path d="M5 21v-6"/><path d="M12 21V9"/><path d="M19 21V3"/>`)

// ChartNoAxesColumn is [chart-no-axes-column](https://lucide.dev/icons/chart-no-axes-column).
var ChartNoAxesColumn = icon.Lucide(viewBox, paint, `<path d="M5 21v-6"/><path d="M12 21V3"/><path d="M19 21V9"/>`)

// ChartNoAxesCombined is [chart-no-axes-combined](https://lucide.dev/icons/chart-no-axes-combined).
var ChartNoAxesCombined = icon.Lucide(viewBox, paint, `<path d="M12 16v5"/><path d="M16 14.639V21"/><path d="M20 10.656V21"/><path d="m22 3-8.646 8.646a.5.5 0 0 1-.708 0L9.354 8.354a.5.5 0 0 0-.707 0L2 15"/><path d="M4 18.463V21"/><path d="M8 14.656V21"/>`)

// ChartNoAxesGantt is [chart-no-axes-gantt](https://lucide.dev/icons/chart-no-axes-gantt).
var ChartNoAxesGantt = icon.Lucide(viewBox, paint, `<path d="M6 5h12"/><path d="M4 12h10"/><path d="M12 19h8"/>`)

// ChartPie is [chart-pie](https://lucide.dev/icons/chart-pie).
var ChartPie = icon.Lucide(viewBox, paint, `<path d="M21 12c.552 0 1.005-.449.95-.998a10 10 0 0 0-8.953-8.951c-.55-.055-.998.398-.998.95v8a1 1 0 0 0 1 1z"/><path d="M21.21 15.89A10 10 0 1 1 8 2.83"/>`)

// ChartScatter is [chart-scatter](https://lucide.dev/icons/chart-scatter).
var ChartScatter = icon.Lucide(viewBox, paint, `<circle cx="7.5" cy="7.5" r=".5" fill="currentColor"/><circle cx="18.5" cy="5.5" r=".5" fill="currentColor"/><circle cx="11.5" cy="11.5" r=".5" fill="currentColor"/><circle cx="7.5" cy="16.5" r=".5" fill="currentColor"/><circle cx="17.5" cy="14.5" r=".5" fill="currentColor"/><path d="M3 3v16a2 2 0 0 0 2 2h16"/>`)

// ChartSpline is [chart-spline](https://lucide.dev/icons/chart-spline).
var ChartSpline = icon.Lucide(viewBox, paint, `<path d="M3 3v16a2 2 0 0 0 2 2h16"/><path d="M7 16c.5-2 1.5-7 4-7 2 0 2 3 4 3 2.5 0 4.5-5 5-7"/>`)

// CheckCheck is [check-check](https://lucide.dev/icons/check-check).
var CheckCheck = icon.Lucide(viewBox, paint, `<path d="M18 6 7 17l-5-5"/><path d="m22 10-7.5 7.5L13 16"/>`)

// CheckLine is [check-line](https://lucide.dev/icons/check-line).
var CheckLine = icon.Lucide(viewBox, paint, `<path d="M20 4L9 15"/><path d="M21 19L3 19"/><path d="M9 15L4 10"/>`)

// Check is [check](https://lucide.dev/icons/check).
var Check = icon.Lucide(viewBox, paint, `<path d="M20 6 9 17l-5-5"/>`)

// ChefHat is [chef-hat](https://lucide.dev/icons/chef-hat).
var ChefHat = icon.Lucide(viewBox, paint, `<path d="M17 21a1 1 0 0 0 1-1v-5.35c0-.457.316-.844.727-1.041a4 4 0 0 0-2.134-7.589 5 5 0 0 0-9.186 0 4 4 0 0 0-2.134 7.588c.411.198.727.585.727 1.041V20a1 1 0 0 0 1 1Z"/><path d="M6 17h12"/>`)

// Cherry is [cherry](https://lucide.dev/icons/cherry).
var Cherry = icon.Lucide(viewBox, paint, `<path d="M2 17a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3-2.5-2-5 .24-5 3Z"/><path d="M12 17a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3-2.5-2-5 .24-5 3Z"/><path d="M7 14c3.22-2.91 4.29-8.75 5-12 1.66 2.38 4.94 9 5 12"/><path d="M22 9c-4.29 0-7.14-2.33-10-7 5.71 0 10 4.67 10 7Z"/>`)

// ChessBishop is [chess-bishop](https://lucide.dev/icons/chess-bishop).
var ChessBishop = icon.Lucide(viewBox, paint, `<path d="M5 20a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z"/><path d="M15 18c1.5-.615 3-2.461 3-4.923C18 8.769 14.5 4.462 12 2 9.5 4.462 6 8.77 6 13.077 6 15.539 7.5 17.385 9 18"/><path d="m16 7-2.5 2.5"/><path d="M9 2h6"/>`)

// ChessKing is [chess-king](https://lucide.dev/icons/chess-king).
var ChessKing = icon.Lucide(viewBox, paint, `<path d="M4 20a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z"/><path d="m6.7 18-1-1C4.35 15.682 3 14.09 3 12a5 5 0 0 1 4.95-5c1.584 0 2.7.455 4.05 1.818C13.35 7.455 14.466 7 16.05 7A5 5 0 0 1 21 12c0 2.082-1.359 3.673-2.7 5l-1 1"/><path d="M10 4h4"/><path d="M12 2v6.818"/>`)

// ChessKnight is [chess-knight](https://lucide.dev/icons/chess-knight).
var ChessKnight = icon.Lucide(viewBox, paint, `<path d="M5 20a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z"/><path d="M16.5 18c1-2 2.5-5 2.5-9a7 7 0 0 0-7-7H6.635a1 1 0 0 0-.768 1.64L7 5l-2.32 5.802a2 2 0 0 0 .95 2.526l2.87 1.456"/><path d="m15 5 1.425-1.425"/><path d="m17 8 1.53-1.53"/><path d="M9.713 12.185 7 18"/>`)

// ChessPawn is [chess-pawn](https://lucide.dev/icons/chess-pawn).
var ChessPawn = icon.Lucide(viewBox, paint, `<path d="M5 20a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z"/><path d="m14.5 10 1.5 8"/><path d="M7 10h10"/><path d="m8 18 1.5-8"/><circle cx="12" cy="6" r="4"/>`)

// ChessQueen is [chess-queen](https://lucide.dev/icons/chess-queen).
var ChessQueen = icon.Lucide(viewBox, paint, `<path d="m10.68 5.499-.996 4.982a.65.65 0 01-1.131.295L5.885 7.663"/><path d="m13.32 5.5.995 4.98a.65.65 0 001.131.296l2.668-3.113"/><path d="M20.428 8.952 17.5 18"/><path d="M4 20a2 2 0 012-2h12a2 2 0 012 2v1a1 1 0 01-1 1H5a1 1 0 01-1-1z"/><path d="M6.5 18 3.571 8.952"/><circle cx="12" cy="4" r="2"/><circle cx="20" cy="7" r="2"/><circle cx="4" cy="7" r="2"/>`)

// ChessRook is [chess-rook](https://lucide.dev/icons/chess-rook).
var ChessRook = icon.Lucide(viewBox, paint, `<path d="M5 20a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1z"/><path d="M10 2v2"/><path d="M14 2v2"/><path d="m17 18-1-9"/><path d="M6 2v5a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2"/><path d="M6 4h12"/><path d="m7 18 1-9"/>`)

// ChevronDown is [chevron-down](https://lucide.dev/icons/chevron-down).
var ChevronDown = icon.Lucide(viewBox, paint, `<path d="m6 9 6 6 6-6"/>`)

// ChevronFirst is [chevron-first](https://lucide.dev/icons/chevron-first).
var ChevronFirst = icon.Lucide(viewBox, paint, `<path d="m17 18-6-6 6-6"/><path d="M7 6v12"/>`)

// ChevronLast is [chevron-last](https://lucide.dev/icons/chevron-last).
var ChevronLast = icon.Lucide(viewBox, paint, `<path d="m7 18 6-6-6-6"/><path d="M17 6v12"/>`)

// ChevronLeft is [chevron-left](https://lucide.dev/icons/chevron-left).
var ChevronLeft = icon.Lucide(viewBox, paint, `<path d="m15 18-6-6 6-6"/>`)

// ChevronRight is [chevron-right](https://lucide.dev/icons/chevron-right).
var ChevronRight = icon.Lucide(viewBox, paint, `<path d="m9 18 6-6-6-6"/>`)

// ChevronUp is [chevron-up](https://lucide.dev/icons/chevron-up).
var ChevronUp = icon.Lucide(viewBox, paint, `<path d="m18 15-6-6-6 6"/>`)

// ChevronsDownUp is [chevrons-down-up](https://lucide.dev/icons/chevrons-down-up).
var ChevronsDownUp = icon.Lucide(viewBox, paint, `<path d="m7 20 5-5 5 5"/><path d="m7 4 5 5 5-5"/>`)

// ChevronsDown is [chevrons-down](https://lucide.dev/icons/chevrons-down).
var ChevronsDown = icon.Lucide(viewBox, paint, `<path d="m7 6 5 5 5-5"/><path d="m7 13 5 5 5-5"/>`)

// ChevronsLeftRightEllipsis is [chevrons-left-right-ellipsis](https://lucide.dev/icons/chevrons-left-right-ellipsis).
var ChevronsLeftRightEllipsis = icon.Lucide(viewBox, paint, `<path d="M12 12h.01"/><path d="M16 12h.01"/><path d="m17 7 5 5-5 5"/><path d="m7 7-5 5 5 5"/><path d="M8 12h.01"/>`)

// ChevronsLeftRight is [chevrons-left-right](https://lucide.dev/icons/chevrons-left-right).
var ChevronsLeftRight = icon.Lucide(viewBox, paint, `<path d="m9 7-5 5 5 5"/><path d="m15 7 5 5-5 5"/>`)

// ChevronsLeft is [chevrons-left](https://lucide.dev/icons/chevrons-left).
var ChevronsLeft = icon.Lucide(viewBox, paint, `<path d="m11 17-5-5 5-5"/><path d="m18 17-5-5 5-5"/>`)

// ChevronsRightLeft is [chevrons-right-left](https://lucide.dev/icons/chevrons-right-left).
var ChevronsRightLeft = icon.Lucide(viewBox, paint, `<path d="m20 17-5-5 5-5"/><path d="m4 17 5-5-5-5"/>`)

// ChevronsRight is [chevrons-right](https://lucide.dev/icons/chevrons-right).
var ChevronsRight = icon.Lucide(viewBox, paint, `<path d="m6 17 5-5-5-5"/><path d="m13 17 5-5-5-5"/>`)

// ChevronsUpDown is [chevrons-up-down](https://lucide.dev/icons/chevrons-up-down).
var ChevronsUpDown = icon.Lucide(viewBox, paint, `<path d="m7 15 5 5 5-5"/><path d="m7 9 5-5 5 5"/>`)

// ChevronsUp is [chevrons-up](https://lucide.dev/icons/chevrons-up).
var ChevronsUp = icon.Lucide(viewBox, paint, `<path d="m17 11-5-5-5 5"/><path d="m17 18-5-5-5 5"/>`)

// Church is [church](https://lucide.dev/icons/church).
var Church = icon.Lucide(viewBox, paint, `<path d="M10 9h4"/><path d="M12 7v5"/><path d="M14 21v-3a2 2 0 0 0-4 0v3"/><path d="m18 9 3.52 2.147a1 1 0 0 1 .48.854V19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-6.999a1 1 0 0 1 .48-.854L6 9"/><path d="M6 21V7a1 1 0 0 1 .376-.782l5-3.999a1 1 0 0 1 1.249.001l5 4A1 1 0 0 1 18 7v14"/>`)

// CigaretteOff is [cigarette-off](https://lucide.dev/icons/cigarette-off).
var CigaretteOff = icon.Lucide(viewBox, paint, `<path d="M12 12H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h13"/><path d="M18 8c0-2.5-2-2.5-2-5"/><path d="m2 2 20 20"/><path d="M21 12a1 1 0 0 1 1 1v2a1 1 0 0 1-.5.866"/><path d="M22 8c0-2.5-2-2.5-2-5"/><path d="M7 12v4"/>`)

// Cigarette is [cigarette](https://lucide.dev/icons/cigarette).
var Cigarette = icon.Lucide(viewBox, paint, `<path d="M17 12H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h14"/><path d="M18 8c0-2.5-2-2.5-2-5"/><path d="M21 16a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1"/><path d="M22 8c0-2.5-2-2.5-2-5"/><path d="M7 12v4"/>`)

// CircleAlert is [circle-alert](https://lucide.dev/icons/circle-alert).
var CircleAlert = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12.01" y1="16" y2="16"/>`)

// CircleArrowDown is [circle-arrow-down](https://lucide.dev/icons/circle-arrow-down).
var CircleArrowDown = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 8v8"/><path d="m8 12 4 4 4-4"/>`)

// CircleArrowLeft is [circle-arrow-left](https://lucide.dev/icons/circle-arrow-left).
var CircleArrowLeft = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m12 8-4 4 4 4"/><path d="M16 12H8"/>`)

// CircleArrowOutDownLeft is [circle-arrow-out-down-left](https://lucide.dev/icons/circle-arrow-out-down-left).
var CircleArrowOutDownLeft = icon.Lucide(viewBox, paint, `<path d="M2 12a10 10 0 1 1 10 10"/><path d="m2 22 10-10"/><path d="M8 22H2v-6"/>`)

// CircleArrowOutDownRight is [circle-arrow-out-down-right](https://lucide.dev/icons/circle-arrow-out-down-right).
var CircleArrowOutDownRight = icon.Lucide(viewBox, paint, `<path d="M12 22a10 10 0 1 1 10-10"/><path d="M22 22 12 12"/><path d="M22 16v6h-6"/>`)

// CircleArrowOutUpLeft is [circle-arrow-out-up-left](https://lucide.dev/icons/circle-arrow-out-up-left).
var CircleArrowOutUpLeft = icon.Lucide(viewBox, paint, `<path d="M2 8V2h6"/><path d="m2 2 10 10"/><path d="M12 2A10 10 0 1 1 2 12"/>`)

// CircleArrowOutUpRight is [circle-arrow-out-up-right](https://lucide.dev/icons/circle-arrow-out-up-right).
var CircleArrowOutUpRight = icon.Lucide(viewBox, paint, `<path d="M22 12A10 10 0 1 1 12 2"/><path d="M22 2 12 12"/><path d="M16 2h6v6"/>`)

// CircleArrowRight is [circle-arrow-right](https://lucide.dev/icons/circle-arrow-right).
var CircleArrowRight = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m12 16 4-4-4-4"/><path d="M8 12h8"/>`)

// CircleArrowUp is [circle-arrow-up](https://lucide.dev/icons/circle-arrow-up).
var CircleArrowUp = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m16 12-4-4-4 4"/><path d="M12 16V8"/>`)

// CircleCheckBig is [circle-check-big](https://lucide.dev/icons/circle-check-big).
var CircleCheckBig = icon.Lucide(viewBox, paint, `<path d="M21.801 10A10 10 0 1 1 17 3.335"/><path d="m9 11 3 3L22 4"/>`)

// CircleCheck is [circle-check](https://lucide.dev/icons/circle-check).
var CircleCheck = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m16 9-5.5 5.5L8 12"/>`)

// CircleChevronDown is [circle-chevron-down](https://lucide.dev/icons/circle-chevron-down).
var CircleChevronDown = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m16 10-4 4-4-4"/>`)

// CircleChevronLeft is [circle-chevron-left](https://lucide.dev/icons/circle-chevron-left).
var CircleChevronLeft = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m14 16-4-4 4-4"/>`)

// CircleChevronRight is [circle-chevron-right](https://lucide.dev/icons/circle-chevron-right).
var CircleChevronRight = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m10 8 4 4-4 4"/>`)

// CircleChevronUp is [circle-chevron-up](https://lucide.dev/icons/circle-chevron-up).
var CircleChevronUp = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m8 14 4-4 4 4"/>`)

// CircleChevronsDown is [circle-chevrons-down](https://lucide.dev/icons/circle-chevrons-down).
var CircleChevronsDown = icon.Lucide(viewBox, paint, `<path d="M9 14L12 17L15 14"/><path d="M9 8L12 11L15 8"/><circle cx="12" cy="12" r="10"/>`)

// CircleChevronsLeft is [circle-chevrons-left](https://lucide.dev/icons/circle-chevrons-left).
var CircleChevronsLeft = icon.Lucide(viewBox, paint, `<path d="M10 9L7 12L10 15"/><path d="M16 9L13 12L16 15"/><circle cx="12" cy="12" r="10"/>`)

// CircleChevronsRight is [circle-chevrons-right](https://lucide.dev/icons/circle-chevrons-right).
var CircleChevronsRight = icon.Lucide(viewBox, paint, `<path d="M14 15L17 12L14 9"/><path d="M8 15L11 12L8 9"/><circle cx="12" cy="12" r="10"/>`)

// CircleChevronsUp is [circle-chevrons-up](https://lucide.dev/icons/circle-chevrons-up).
var CircleChevronsUp = icon.Lucide(viewBox, paint, `<path d="M15 10L12 7L9 10"/><path d="M15 16L12 13L9 16"/><circle cx="12" cy="12" r="10"/>`)

// CircleDashedCheck is [circle-dashed-check](https://lucide.dev/icons/circle-dashed-check).
var CircleDashedCheck = icon.Lucide(viewBox, paint, `<path d="M10.1 2.182a10 10 0 013.8 0"/><path d="M13.9 21.818a10 10 0 01-3.8 0"/><path d="m16 9-5.5 5.5L8 12"/><path d="M17.609 3.721a10 10 0 012.69 2.7"/><path d="M2.182 13.9a10 10 0 010-3.8"/><path d="M20.279 17.609a10 10 0 01-2.7 2.69"/><path d="M21.818 10.1a10 10 0 010 3.8"/><path d="M3.721 6.391a10 10 0 012.7-2.69"/><path d="M6.391 20.279a10 10 0 01-2.69-2.7"/>`)

// CircleDashed is [circle-dashed](https://lucide.dev/icons/circle-dashed).
var CircleDashed = icon.Lucide(viewBox, paint, `<path d="M10.1 2.182a10 10 0 0 1 3.8 0"/><path d="M13.9 21.818a10 10 0 0 1-3.8 0"/><path d="M17.609 3.721a10 10 0 0 1 2.69 2.7"/><path d="M2.182 13.9a10 10 0 0 1 0-3.8"/><path d="M20.279 17.609a10 10 0 0 1-2.7 2.69"/><path d="M21.818 10.1a10 10 0 0 1 0 3.8"/><path d="M3.721 6.391a10 10 0 0 1 2.7-2.69"/><path d="M6.391 20.279a10 10 0 0 1-2.69-2.7"/>`)

// CircleDivide is [circle-divide](https://lucide.dev/icons/circle-divide).
var CircleDivide = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><line x1="8" x2="16" y1="12" y2="12"/><line x1="12" x2="12" y1="16" y2="16"/><line x1="12" x2="12" y1="8" y2="8"/>`)

// CircleDollarSign is [circle-dollar-sign](https://lucide.dev/icons/circle-dollar-sign).
var CircleDollarSign = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8"/><path d="M12 18V6"/>`)

// CircleDotDashed is [circle-dot-dashed](https://lucide.dev/icons/circle-dot-dashed).
var CircleDotDashed = icon.Lucide(viewBox, paint, `<path d="M10.1 2.18a9.93 9.93 0 0 1 3.8 0"/><path d="M17.6 3.71a9.95 9.95 0 0 1 2.69 2.7"/><path d="M21.82 10.1a9.93 9.93 0 0 1 0 3.8"/><path d="M20.29 17.6a9.95 9.95 0 0 1-2.7 2.69"/><path d="M13.9 21.82a9.94 9.94 0 0 1-3.8 0"/><path d="M6.4 20.29a9.95 9.95 0 0 1-2.69-2.7"/><path d="M2.18 13.9a9.93 9.93 0 0 1 0-3.8"/><path d="M3.71 6.4a9.95 9.95 0 0 1 2.7-2.69"/><circle cx="12" cy="12" r="1"/>`)

// CircleDot is [circle-dot](https://lucide.dev/icons/circle-dot).
var CircleDot = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="1"/><circle cx="12" cy="12" r="10"/>`)

// CircleEllipsis is [circle-ellipsis](https://lucide.dev/icons/circle-ellipsis).
var CircleEllipsis = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M17 12h.01"/><path d="M12 12h.01"/><path d="M7 12h.01"/>`)

// CircleEqual is [circle-equal](https://lucide.dev/icons/circle-equal).
var CircleEqual = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M7 10h10"/><path d="M7 14h10"/>`)

// CircleEuro is [circle-euro](https://lucide.dev/icons/circle-euro).
var CircleEuro = icon.Lucide(viewBox, paint, `<path d="M15 9.4a4 4 0 1 0 0 5.2"/><path d="M7 12h5"/><circle cx="12" cy="12" r="10"/>`)

// CircleFadingArrowUp is [circle-fading-arrow-up](https://lucide.dev/icons/circle-fading-arrow-up).
var CircleFadingArrowUp = icon.Lucide(viewBox, paint, `<path d="M12 2a10 10 0 0 1 7.38 16.75"/><path d="m16 12-4-4-4 4"/><path d="M12 16V8"/><path d="M2.5 8.875a10 10 0 0 0-.5 3"/><path d="M2.83 16a10 10 0 0 0 2.43 3.4"/><path d="M4.636 5.235a10 10 0 0 1 .891-.857"/><path d="M8.644 21.42a10 10 0 0 0 7.631-.38"/>`)

// CircleFadingPlay is [circle-fading-play](https://lucide.dev/icons/circle-fading-play).
var CircleFadingPlay = icon.Lucide(viewBox, paint, `<path d="M12 2a10 10 0 0 1 7.38 16.75"/><path d="M9 9.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997A1 1 0 0 1 9 14.996z"/><path d="M2.5 8.875a10 10 0 0 0-.5 3"/><path d="M2.83 16a10 10 0 0 0 2.43 3.4"/><path d="M4.636 5.235a10 10 0 0 1 .891-.857"/><path d="M8.644 21.42a10 10 0 0 0 7.631-.38"/>`)

// CircleFadingPlus is [circle-fading-plus](https://lucide.dev/icons/circle-fading-plus).
var CircleFadingPlus = icon.Lucide(viewBox, paint, `<path d="M12 2a10 10 0 0 1 7.38 16.75"/><path d="M12 8v8"/><path d="M16 12H8"/><path d="M2.5 8.875a10 10 0 0 0-.5 3"/><path d="M2.83 16a10 10 0 0 0 2.43 3.4"/><path d="M4.636 5.235a10 10 0 0 1 .891-.857"/><path d="M8.644 21.42a10 10 0 0 0 7.631-.38"/>`)

// CircleGauge is [circle-gauge](https://lucide.dev/icons/circle-gauge).
var CircleGauge = icon.Lucide(viewBox, paint, `<path d="M15.6 2.7a10 10 0 1 0 5.7 5.7"/><circle cx="12" cy="12" r="2"/><path d="M13.4 10.6 19 5"/>`)

// CircleMinus is [circle-minus](https://lucide.dev/icons/circle-minus).
var CircleMinus = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M8 12h8"/>`)

// CircleOff is [circle-off](https://lucide.dev/icons/circle-off).
var CircleOff = icon.Lucide(viewBox, paint, `<path d="m2 2 20 20"/><path d="M8.35 2.69A10 10 0 0 1 21.3 15.65"/><path d="M19.08 19.08A10 10 0 1 1 4.92 4.92"/>`)

// CircleParkingOff is [circle-parking-off](https://lucide.dev/icons/circle-parking-off).
var CircleParkingOff = icon.Lucide(viewBox, paint, `<path d="M12.656 7H13a3 3 0 0 1 2.984 3.307"/><path d="M13 13H9"/><path d="M19.071 19.071A1 1 0 0 1 4.93 4.93"/><path d="m2 2 20 20"/><path d="M8.357 2.687a10 10 0 0 1 12.956 12.956"/><path d="M9 17V9"/>`)

// CircleParking is [circle-parking](https://lucide.dev/icons/circle-parking).
var CircleParking = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M9 17V7h4a3 3 0 0 1 0 6H9"/>`)

// CirclePause is [circle-pause](https://lucide.dev/icons/circle-pause).
var CirclePause = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><line x1="10" x2="10" y1="15" y2="9"/><line x1="14" x2="14" y1="15" y2="9"/>`)

// CirclePercent is [circle-percent](https://lucide.dev/icons/circle-percent).
var CirclePercent = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m15 9-6 6"/><path d="M9 9h.01"/><path d="M15 15h.01"/>`)

// CirclePile is [circle-pile](https://lucide.dev/icons/circle-pile).
var CirclePile = icon.Lucide(viewBox, paint, `<circle cx="12" cy="19" r="2"/><circle cx="12" cy="5" r="2"/><circle cx="16" cy="12" r="2"/><circle cx="20" cy="19" r="2"/><circle cx="4" cy="19" r="2"/><circle cx="8" cy="12" r="2"/>`)

// CirclePlay is [circle-play](https://lucide.dev/icons/circle-play).
var CirclePlay = icon.Lucide(viewBox, paint, `<path d="M9 9.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997A1 1 0 0 1 9 14.996z"/><circle cx="12" cy="12" r="10"/>`)

// CirclePlus is [circle-plus](https://lucide.dev/icons/circle-plus).
var CirclePlus = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M8 12h8"/><path d="M12 8v8"/>`)

// CirclePoundSterling is [circle-pound-sterling](https://lucide.dev/icons/circle-pound-sterling).
var CirclePoundSterling = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M10 16V9.5a1 1 0 0 1 5 0"/><path d="M8 12h4"/><path d="M8 16h7"/>`)

// CirclePower is [circle-power](https://lucide.dev/icons/circle-power).
var CirclePower = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 7v4"/><path d="M7.998 9.003a5 5 0 1 0 8-.005"/>`)

// CircleQuestionMark is [circle-question-mark](https://lucide.dev/icons/circle-question-mark).
var CircleQuestionMark = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><path d="M12 17h.01"/>`)

// CircleSlash2 is [circle-slash-2](https://lucide.dev/icons/circle-slash-2).
var CircleSlash2 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M22 2 2 22"/>`)

// CircleSlash is [circle-slash](https://lucide.dev/icons/circle-slash).
var CircleSlash = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><line x1="9" x2="15" y1="15" y2="9"/>`)

// CircleSmall is [circle-small](https://lucide.dev/icons/circle-small).
var CircleSmall = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="6"/>`)

// CircleStar is [circle-star](https://lucide.dev/icons/circle-star).
var CircleStar = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M11.051 7.616a1 1 0 0 1 1.909.024l.737 1.452a1 1 0 0 0 .737.535l1.634.256a1 1 0 0 1 .588 1.806l-1.172 1.168a1 1 0 0 0-.282.866l.259 1.613a1 1 0 0 1-1.541 1.134l-1.465-.75a1 1 0 0 0-.912 0l-1.465.75a1 1 0 0 1-1.539-1.133l.258-1.613a1 1 0 0 0-.282-.867l-1.156-1.152a1 1 0 0 1 .572-1.822l1.633-.256a1 1 0 0 0 .737-.535z"/>`)

// CircleStop is [circle-stop](https://lucide.dev/icons/circle-stop).
var CircleStop = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><rect x="9" y="9" width="6" height="6" rx="1"/>`)

// CircleUserRound is [circle-user-round](https://lucide.dev/icons/circle-user-round).
var CircleUserRound = icon.Lucide(viewBox, paint, `<path d="M17.925 20.056a6 6 0 0 0-11.851.001"/><circle cx="12" cy="11" r="4"/><circle cx="12" cy="12" r="10"/>`)

// CircleUser is [circle-user](https://lucide.dev/icons/circle-user).
var CircleUser = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><circle cx="12" cy="10" r="3"/><path d="M7 20.662V19a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1.662"/>`)

// CircleX is [circle-x](https://lucide.dev/icons/circle-x).
var CircleX = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m15 9-6 6"/><path d="m9 9 6 6"/>`)

// Circle is [circle](https://lucide.dev/icons/circle).
var Circle = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/>`)

// CircuitBoard is [circuit-board](https://lucide.dev/icons/circuit-board).
var CircuitBoard = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M11 9h4a2 2 0 0 0 2-2V3"/><circle cx="9" cy="9" r="2"/><path d="M7 21v-4a2 2 0 0 1 2-2h4"/><circle cx="15" cy="15" r="2"/>`)

// Citrus is [citrus](https://lucide.dev/icons/citrus).
var Citrus = icon.Lucide(viewBox, paint, `<path d="M21.66 17.67a1.08 1.08 0 0 1-.04 1.6A12 12 0 0 1 4.73 2.38a1.1 1.1 0 0 1 1.61-.04z"/><path d="M19.65 15.66A8 8 0 0 1 8.35 4.34"/><path d="m14 10-5.5 5.5"/><path d="M14 17.85V10H6.15"/>`)

// Clapperboard is [clapperboard](https://lucide.dev/icons/clapperboard).
var Clapperboard = icon.Lucide(viewBox, paint, `<path d="m12.296 3.464 3.02 3.956"/><path d="M20.2 6 3 11l-.9-2.4c-.3-1.1.3-2.2 1.3-2.5l13.5-4c1.1-.3 2.2.3 2.5 1.3z"/><path d="M3 11h18v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="m6.18 5.276 3.1 3.899"/>`)

// ClefAlto is [clef-alto](https://lucide.dev/icons/clef-alto).
var ClefAlto = icon.Lucide(viewBox, paint, `<path d="M10 4v16"/><path d="M14 4.764a3 3 0 1 1-.152 4.327A4 4 0 0 1 10 12a4 4 0 0 1 3.848 2.909A3 3 0 1 1 14 19.236"/><path d="M6 4v16"/>`)

// ClefBass is [clef-bass](https://lucide.dev/icons/clef-bass).
var ClefBass = icon.Lucide(viewBox, paint, `<path d="M19 11h.01"/><path d="M19 6h.01"/><path d="M5 8c0-4 4-4 4-4 6 0 6 6 6 6 0 7-10 11-10 11"/><circle cx="7" cy="8" r="2"/>`)

// ClefTreble is [clef-treble](https://lucide.dev/icons/clef-treble).
var ClefTreble = icon.Lucide(viewBox, paint, `<path d="M10.586 21.414a2 2 0 0 0 3.378-1.791L11.036 4.377a2 2 0 1 1 3.378 1.037C12.414 7.414 7 8 7 13a5 5 0 0 0 5 5 5 4 0 0 0 5-4 3 3 0 0 0-3-3 3 2 0 0 0-3 2"/>`)

// ClipboardCheck is [clipboard-check](https://lucide.dev/icons/clipboard-check).
var ClipboardCheck = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="m9 14 2 2 4-4"/>`)

// ClipboardClock is [clipboard-clock](https://lucide.dev/icons/clipboard-clock).
var ClipboardClock = icon.Lucide(viewBox, paint, `<path d="M16 14v2.2l1.6 1"/><path d="M16 4h2a2 2 0 0 1 2 2v.832"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h2"/><circle cx="16" cy="16" r="6"/><rect x="8" y="2" width="8" height="4" rx="1"/>`)

// ClipboardCopy is [clipboard-copy](https://lucide.dev/icons/clipboard-copy).
var ClipboardCopy = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2"/><path d="M16 4h2a2 2 0 0 1 2 2v4"/><path d="M21 14H11"/><path d="m15 10-4 4 4 4"/>`)

// ClipboardList is [clipboard-list](https://lucide.dev/icons/clipboard-list).
var ClipboardList = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="M12 11h4"/><path d="M12 16h4"/><path d="M8 11h.01"/><path d="M8 16h.01"/>`)

// ClipboardMinus is [clipboard-minus](https://lucide.dev/icons/clipboard-minus).
var ClipboardMinus = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="M9 14h6"/>`)

// ClipboardPaste is [clipboard-paste](https://lucide.dev/icons/clipboard-paste).
var ClipboardPaste = icon.Lucide(viewBox, paint, `<path d="M11 14h10"/><path d="M16 4h2a2 2 0 0 1 2 2v1.344"/><path d="m17 18 4-4-4-4"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 1.793-1.113"/><rect x="8" y="2" width="8" height="4" rx="1"/>`)

// ClipboardPenLine is [clipboard-pen-line](https://lucide.dev/icons/clipboard-pen-line).
var ClipboardPenLine = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-.5"/><path d="M16 4h2a2 2 0 0 1 1.73 1"/><path d="M8 18h1"/><path d="M21.378 12.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/>`)

// ClipboardPen is [clipboard-pen](https://lucide.dev/icons/clipboard-pen).
var ClipboardPen = icon.Lucide(viewBox, paint, `<path d="M16 4h2a2 2 0 0 1 2 2v2"/><path d="M21.34 15.664a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/><path d="M8 22H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><rect x="8" y="2" width="8" height="4" rx="1"/>`)

// ClipboardPlus is [clipboard-plus](https://lucide.dev/icons/clipboard-plus).
var ClipboardPlus = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="M9 14h6"/><path d="M12 17v-6"/>`)

// ClipboardType is [clipboard-type](https://lucide.dev/icons/clipboard-type).
var ClipboardType = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="M9 12v-1h6v1"/><path d="M11 17h2"/><path d="M12 11v6"/>`)

// ClipboardX is [clipboard-x](https://lucide.dev/icons/clipboard-x).
var ClipboardX = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="m14.5 11.5-5 5"/><path d="m9.5 11.5 5 5"/>`)

// Clipboard is [clipboard](https://lucide.dev/icons/clipboard).
var Clipboard = icon.Lucide(viewBox, paint, `<rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/>`)

// Clock1 is [clock-1](https://lucide.dev/icons/clock-1).
var Clock1 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l2-4"/>`)

// Clock10 is [clock-10](https://lucide.dev/icons/clock-10).
var Clock10 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l-4-2"/>`)

// Clock11 is [clock-11](https://lucide.dev/icons/clock-11).
var Clock11 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l-2-4"/>`)

// Clock12 is [clock-12](https://lucide.dev/icons/clock-12).
var Clock12 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6"/>`)

// Clock2 is [clock-2](https://lucide.dev/icons/clock-2).
var Clock2 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l4-2"/>`)

// Clock3 is [clock-3](https://lucide.dev/icons/clock-3).
var Clock3 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6h4"/>`)

// Clock4 is [clock-4](https://lucide.dev/icons/clock-4).
var Clock4 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/>`)

// Clock5 is [clock-5](https://lucide.dev/icons/clock-5).
var Clock5 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l2 4"/>`)

// Clock6 is [clock-6](https://lucide.dev/icons/clock-6).
var Clock6 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v10"/>`)

// Clock7 is [clock-7](https://lucide.dev/icons/clock-7).
var Clock7 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l-2 4"/>`)

// Clock8 is [clock-8](https://lucide.dev/icons/clock-8).
var Clock8 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l-4 2"/>`)

// Clock9 is [clock-9](https://lucide.dev/icons/clock-9).
var Clock9 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6H8"/>`)

// ClockAlert is [clock-alert](https://lucide.dev/icons/clock-alert).
var ClockAlert = icon.Lucide(viewBox, paint, `<path d="M12 6v6l4 2"/><path d="M20 12v5"/><path d="M20 21h.01"/><path d="M21.25 8.2A10 10 0 1 0 16 21.16"/>`)

// ClockArrowDown is [clock-arrow-down](https://lucide.dev/icons/clock-arrow-down).
var ClockArrowDown = icon.Lucide(viewBox, paint, `<path d="M12 6v6l2 1"/><path d="M12.337 21.994a10 10 0 1 1 9.588-8.767"/><path d="m14 18 4 4 4-4"/><path d="M18 14v8"/>`)

// ClockArrowLeft is [clock-arrow-left](https://lucide.dev/icons/clock-arrow-left).
var ClockArrowLeft = icon.Lucide(viewBox, paint, `<path d="M12 6v6l1.5.8"/><path d="M12.338 21.994a10 10 0 1 1 9.587-8.767"/><path d="M14 18h8"/><path d="m18 22-4-4 4-4"/>`)

// ClockArrowRight is [clock-arrow-right](https://lucide.dev/icons/clock-arrow-right).
var ClockArrowRight = icon.Lucide(viewBox, paint, `<path d="M12 6v6l2 1"/><path d="M13.5 21.885A10 10 0 1 1 22 12"/><path d="M14 18h8"/><path d="m18 22 4-4-4-4"/>`)

// ClockArrowUp is [clock-arrow-up](https://lucide.dev/icons/clock-arrow-up).
var ClockArrowUp = icon.Lucide(viewBox, paint, `<path d="M12 6v6l1.56.78"/><path d="M13.227 21.925a10 10 0 1 1 8.767-9.588"/><path d="m14 18 4-4 4 4"/><path d="M18 22v-8"/>`)

// ClockCheck is [clock-check](https://lucide.dev/icons/clock-check).
var ClockCheck = icon.Lucide(viewBox, paint, `<path d="M21.95 13a10 10 0 1 0-8.685 8.92"/><path d="M12 6v6l4 2"/><path d="m16 19 2 2 4-4"/>`)

// ClockFading is [clock-fading](https://lucide.dev/icons/clock-fading).
var ClockFading = icon.Lucide(viewBox, paint, `<path d="M12 2a10 10 0 0 1 7.38 16.75"/><path d="M12 6v6l4 2"/><path d="M2.5 8.875a10 10 0 0 0-.5 3"/><path d="M2.83 16a10 10 0 0 0 2.43 3.4"/><path d="M4.636 5.235a10 10 0 0 1 .891-.857"/><path d="M8.644 21.42a10 10 0 0 0 7.631-.38"/>`)

// ClockPlus is [clock-plus](https://lucide.dev/icons/clock-plus).
var ClockPlus = icon.Lucide(viewBox, paint, `<path d="M12 6v6l3.644 1.822"/><path d="M16 19h6"/><path d="M19 16v6"/><path d="M21.92 13.267a10 10 0 1 0-8.653 8.653"/>`)

// Clock is [clock](https://lucide.dev/icons/clock).
var Clock = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/>`)

// ClosedCaption is [closed-caption](https://lucide.dev/icons/closed-caption).
var ClosedCaption = icon.Lucide(viewBox, paint, `<path d="M10 9.17a3 3 0 1 0 0 5.66"/><path d="M17 9.17a3 3 0 1 0 0 5.66"/><rect x="2" y="5" width="20" height="14" rx="2"/>`)

// CloudAlert is [cloud-alert](https://lucide.dev/icons/cloud-alert).
var CloudAlert = icon.Lucide(viewBox, paint, `<path d="M12 12v4"/><path d="M12 20h.01"/><path d="M8.128 16.949A7 7 0 1 1 15.71 8h1.79a1 1 0 0 1 0 9h-1.642"/>`)

// CloudBackup is [cloud-backup](https://lucide.dev/icons/cloud-backup).
var CloudBackup = icon.Lucide(viewBox, paint, `<path d="M21 15.251A4.5 4.5 0 0 0 17.5 8h-1.79A7 7 0 1 0 3 13.607"/><path d="M7 11v4h4"/><path d="M8 19a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5 4.82 4.82 0 0 0-3.41 1.41L7 15"/>`)

// CloudCheck is [cloud-check](https://lucide.dev/icons/cloud-check).
var CloudCheck = icon.Lucide(viewBox, paint, `<path d="m17 15-5.5 5.5L9 18"/><path d="M5.516 16.07A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 3.501 7.327"/>`)

// CloudCog is [cloud-cog](https://lucide.dev/icons/cloud-cog).
var CloudCog = icon.Lucide(viewBox, paint, `<path d="m10.852 19.772-.383.924"/><path d="m13.148 14.228.383-.923"/><path d="M13.148 19.772a3 3 0 1 0-2.296-5.544l-.383-.923"/><path d="m13.53 20.696-.382-.924a3 3 0 1 1-2.296-5.544"/><path d="m14.772 15.852.923-.383"/><path d="m14.772 18.148.923.383"/><path d="M4.2 15.1a7 7 0 1 1 9.93-9.858A7 7 0 0 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.2"/><path d="m9.228 15.852-.923-.383"/><path d="m9.228 18.148-.923.383"/>`)

// CloudDownload is [cloud-download](https://lucide.dev/icons/cloud-download).
var CloudDownload = icon.Lucide(viewBox, paint, `<path d="M12 13v8l-4-4"/><path d="m12 21 4-4"/><path d="M4.393 15.269A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.436 8.284"/>`)

// CloudDrizzle is [cloud-drizzle](https://lucide.dev/icons/cloud-drizzle).
var CloudDrizzle = icon.Lucide(viewBox, paint, `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="M8 19v1"/><path d="M8 14v1"/><path d="M16 19v1"/><path d="M16 14v1"/><path d="M12 21v1"/><path d="M12 16v1"/>`)

// CloudFog is [cloud-fog](https://lucide.dev/icons/cloud-fog).
var CloudFog = icon.Lucide(viewBox, paint, `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="M16 17H7"/><path d="M17 21H9"/>`)

// CloudHail is [cloud-hail](https://lucide.dev/icons/cloud-hail).
var CloudHail = icon.Lucide(viewBox, paint, `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="M16 14v2"/><path d="M8 14v2"/><path d="M16 20h.01"/><path d="M8 20h.01"/><path d="M12 16v2"/><path d="M12 22h.01"/>`)

// CloudLightning is [cloud-lightning](https://lucide.dev/icons/cloud-lightning).
var CloudLightning = icon.Lucide(viewBox, paint, `<path d="M6 16.326A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 .5 8.973"/><path d="m13 12-3 5h4l-3 5"/>`)

// CloudMoonRain is [cloud-moon-rain](https://lucide.dev/icons/cloud-moon-rain).
var CloudMoonRain = icon.Lucide(viewBox, paint, `<path d="M11 20v2"/><path d="M18.376 14.512a6 6 0 0 0 3.461-4.127c.148-.625-.659-.97-1.248-.714a4 4 0 0 1-5.259-5.26c.255-.589-.09-1.395-.716-1.248a6 6 0 0 0-4.594 5.36"/><path d="M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24"/><path d="M7 19v2"/>`)

// CloudMoon is [cloud-moon](https://lucide.dev/icons/cloud-moon).
var CloudMoon = icon.Lucide(viewBox, paint, `<path d="M13 16a3 3 0 0 1 0 6H7a5 5 0 1 1 4.9-6z"/><path d="M18.376 14.512a6 6 0 0 0 3.461-4.127c.148-.625-.659-.97-1.248-.714a4 4 0 0 1-5.259-5.26c.255-.589-.09-1.395-.716-1.248a6 6 0 0 0-4.594 5.36"/>`)

// CloudOff is [cloud-off](https://lucide.dev/icons/cloud-off).
var CloudOff = icon.Lucide(viewBox, paint, `<path d="M10.94 5.274A7 7 0 0 1 15.71 10h1.79a4.5 4.5 0 0 1 4.222 6.057"/><path d="M18.796 18.81A4.5 4.5 0 0 1 17.5 19H9A7 7 0 0 1 5.79 5.78"/><path d="m2 2 20 20"/>`)

// CloudRainWind is [cloud-rain-wind](https://lucide.dev/icons/cloud-rain-wind).
var CloudRainWind = icon.Lucide(viewBox, paint, `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="m9.2 22 3-7"/><path d="m9 13-3 7"/><path d="m17 13-3 7"/>`)

// CloudRain is [cloud-rain](https://lucide.dev/icons/cloud-rain).
var CloudRain = icon.Lucide(viewBox, paint, `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="M16 14v6"/><path d="M8 14v6"/><path d="M12 16v6"/>`)

// CloudSnow is [cloud-snow](https://lucide.dev/icons/cloud-snow).
var CloudSnow = icon.Lucide(viewBox, paint, `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="M8 15h.01"/><path d="M8 19h.01"/><path d="M12 17h.01"/><path d="M12 21h.01"/><path d="M16 15h.01"/><path d="M16 19h.01"/>`)

// CloudSunRain is [cloud-sun-rain](https://lucide.dev/icons/cloud-sun-rain).
var CloudSunRain = icon.Lucide(viewBox, paint, `<path d="M12 2v2"/><path d="m4.93 4.93 1.41 1.41"/><path d="M20 12h2"/><path d="m19.07 4.93-1.41 1.41"/><path d="M15.947 12.65a4 4 0 0 0-5.925-4.128"/><path d="M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24"/><path d="M11 20v2"/><path d="M7 19v2"/>`)

// CloudSun is [cloud-sun](https://lucide.dev/icons/cloud-sun).
var CloudSun = icon.Lucide(viewBox, paint, `<path d="M12 2v2"/><path d="m4.93 4.93 1.41 1.41"/><path d="M20 12h2"/><path d="m19.07 4.93-1.41 1.41"/><path d="M15.947 12.65a4 4 0 0 0-5.925-4.128"/><path d="M13 22H7a5 5 0 1 1 4.9-6H13a3 3 0 0 1 0 6Z"/>`)

// CloudSync is [cloud-sync](https://lucide.dev/icons/cloud-sync).
var CloudSync = icon.Lucide(viewBox, paint, `<path d="m17 18-1.535 1.605a5 5 0 0 1-8-1.5"/><path d="M17 22v-4h-4"/><path d="M20.996 15.251A4.5 4.5 0 0 0 17.495 8h-1.79a7 7 0 1 0-12.709 5.607"/><path d="M7 10v4h4"/><path d="m7 14 1.535-1.605a5 5 0 0 1 8 1.5"/>`)

// CloudUpload is [cloud-upload](https://lucide.dev/icons/cloud-upload).
var CloudUpload = icon.Lucide(viewBox, paint, `<path d="M12 13v8"/><path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"/><path d="m8 17 4-4 4 4"/>`)

// Cloud is [cloud](https://lucide.dev/icons/cloud).
var Cloud = icon.Lucide(viewBox, paint, `<path d="M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"/>`)

// Cloudy is [cloudy](https://lucide.dev/icons/cloudy).
var Cloudy = icon.Lucide(viewBox, paint, `<path d="M17.5 12a1 1 0 1 1 0 9H9.006a7 7 0 1 1 6.702-9z"/><path d="M21.832 9A3 3 0 0 0 19 7h-2.207a5.5 5.5 0 0 0-10.72.61"/>`)

// Clover is [clover](https://lucide.dev/icons/clover).
var Clover = icon.Lucide(viewBox, paint, `<path d="M16.17 7.83 2 22"/><path d="M4.02 12a2.827 2.827 0 1 1 3.81-4.17A2.827 2.827 0 1 1 12 4.02a2.827 2.827 0 1 1 4.17 3.81A2.827 2.827 0 1 1 19.98 12a2.827 2.827 0 1 1-3.81 4.17A2.827 2.827 0 1 1 12 19.98a2.827 2.827 0 1 1-4.17-3.81A1 1 0 1 1 4 12"/><path d="m7.83 7.83 8.34 8.34"/>`)

// Club is [club](https://lucide.dev/icons/club).
var Club = icon.Lucide(viewBox, paint, `<path d="M17.28 9.05a5.5 5.5 0 1 0-10.56 0A5.5 5.5 0 1 0 12 17.66a5.5 5.5 0 1 0 5.28-8.6Z"/><path d="M12 17.66L12 22"/>`)

// CodeXml is [code-xml](https://lucide.dev/icons/code-xml).
var CodeXml = icon.Lucide(viewBox, paint, `<path d="m18 16 4-4-4-4"/><path d="m6 8-4 4 4 4"/><path d="m14.5 4-5 16"/>`)

// Code is [code](https://lucide.dev/icons/code).
var Code = icon.Lucide(viewBox, paint, `<path d="m16 18 6-6-6-6"/><path d="m8 6-6 6 6 6"/>`)

// Coffee is [coffee](https://lucide.dev/icons/coffee).
var Coffee = icon.Lucide(viewBox, paint, `<path d="M10 2v2"/><path d="M14 2v2"/><path d="M16 8a1 1 0 0 1 1 1v8a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4V9a1 1 0 0 1 1-1h14a4 4 0 1 1 0 8h-1"/><path d="M6 2v2"/>`)

// Cog is [cog](https://lucide.dev/icons/cog).
var Cog = icon.Lucide(viewBox, paint, `<path d="M11 10.27 7 3.34"/><path d="m11 13.73-4 6.93"/><path d="M12 22v-2"/><path d="M12 2v2"/><path d="M14 12h8"/><path d="m17 20.66-1-1.73"/><path d="m17 3.34-1 1.73"/><path d="M2 12h2"/><path d="m20.66 17-1.73-1"/><path d="m20.66 7-1.73 1"/><path d="m3.34 17 1.73-1"/><path d="m3.34 7 1.73 1"/><circle cx="12" cy="12" r="2"/><circle cx="12" cy="12" r="8"/>`)

// Coins is [coins](https://lucide.dev/icons/coins).
var Coins = icon.Lucide(viewBox, paint, `<path d="M13.744 17.736a6 6 0 1 1-7.48-7.48"/><path d="M15 6h1v4"/><path d="m6.134 14.768.866-.5 2 3.464"/><circle cx="16" cy="8" r="6"/>`)

// Columns2 is [columns-2](https://lucide.dev/icons/columns-2).
var Columns2 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 3v18"/>`)

// Columns3Cog is [columns-3-cog](https://lucide.dev/icons/columns-3-cog).
var Columns3Cog = icon.Lucide(viewBox, paint, `<path d="M10.6 21H5a2 2 0 01-2-2V5a2 2 0 012-2h14a2 2 0 012 2v5.6"/><path d="m14.305 19.53.923-.382"/><path d="M15 3v7.6"/><path d="m15.229 16.852-.924-.383"/><path d="m16.852 15.228-.383-.923"/><path d="m16.852 20.772-.383.924"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.696-.382-.924"/><path d="m20.773 16.852.922-.383"/><path d="m20.773 19.148.922.383"/><path d="M9 3v18"/><circle cx="18" cy="18" r="3"/>`)

// Columns3 is [columns-3](https://lucide.dev/icons/columns-3).
var Columns3 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="M15 3v18"/>`)

// Columns4 is [columns-4](https://lucide.dev/icons/columns-4).
var Columns4 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7.5 3v18"/><path d="M12 3v18"/><path d="M16.5 3v18"/>`)

// Combine is [combine](https://lucide.dev/icons/combine).
var Combine = icon.Lucide(viewBox, paint, `<path d="M14 3a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1"/><path d="M19 3a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1"/><path d="m7 15 3 3"/><path d="m7 21 3-3H5a2 2 0 0 1-2-2v-2"/><rect x="14" y="14" width="7" height="7" rx="1"/><rect x="3" y="3" width="7" height="7" rx="1"/>`)

// Command is [command](https://lucide.dev/icons/command).
var Command = icon.Lucide(viewBox, paint, `<path d="M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"/>`)

// Compass is [compass](https://lucide.dev/icons/compass).
var Compass = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m16.24 7.76-1.804 5.411a2 2 0 0 1-1.265 1.265L7.76 16.24l1.804-5.411a2 2 0 0 1 1.265-1.265z"/>`)

// Component is [component](https://lucide.dev/icons/component).
var Component = icon.Lucide(viewBox, paint, `<path d="M15.536 11.293a1 1 0 0 0 0 1.414l2.376 2.377a1 1 0 0 0 1.414 0l2.377-2.377a1 1 0 0 0 0-1.414l-2.377-2.377a1 1 0 0 0-1.414 0z"/><path d="M2.297 11.293a1 1 0 0 0 0 1.414l2.377 2.377a1 1 0 0 0 1.414 0l2.377-2.377a1 1 0 0 0 0-1.414L6.088 8.916a1 1 0 0 0-1.414 0z"/><path d="M8.916 17.912a1 1 0 0 0 0 1.415l2.377 2.376a1 1 0 0 0 1.414 0l2.377-2.376a1 1 0 0 0 0-1.415l-2.377-2.376a1 1 0 0 0-1.414 0z"/><path d="M8.916 4.674a1 1 0 0 0 0 1.414l2.377 2.376a1 1 0 0 0 1.414 0l2.377-2.376a1 1 0 0 0 0-1.414l-2.377-2.377a1 1 0 0 0-1.414 0z"/>`)

// Computer is [computer](https://lucide.dev/icons/computer).
var Computer = icon.Lucide(viewBox, paint, `<path d="M12 18h6"/><path d="M6 18h.01"/><path d="M8 6h1"/><rect x="2" y="14" width="20" height="8" rx="2"/><rect x="4" y="2" width="16" height="12" rx="2"/>`)

// ConciergeBell is [concierge-bell](https://lucide.dev/icons/concierge-bell).
var ConciergeBell = icon.Lucide(viewBox, paint, `<path d="M3 20a1 1 0 0 1-1-1v-1a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1Z"/><path d="M20 16a8 8 0 1 0-16 0"/><path d="M12 4v4"/><path d="M10 4h4"/>`)

// Cone is [cone](https://lucide.dev/icons/cone).
var Cone = icon.Lucide(viewBox, paint, `<path d="m20.9 18.55-8-15.98a1 1 0 0 0-1.8 0l-8 15.98"/><ellipse cx="12" cy="19" rx="9" ry="3"/>`)

// Construction is [construction](https://lucide.dev/icons/construction).
var Construction = icon.Lucide(viewBox, paint, `<rect x="2" y="6" width="20" height="8" rx="1"/><path d="M17 14v7"/><path d="M7 14v7"/><path d="M17 3v3"/><path d="M7 3v3"/><path d="M10 14 2.3 6.3"/><path d="m14 6 7.7 7.7"/><path d="m8 6 8 8"/>`)

// ContactRound is [contact-round](https://lucide.dev/icons/contact-round).
var ContactRound = icon.Lucide(viewBox, paint, `<path d="M16 2v2"/><path d="M17.915 21a6 6 0 10-12 0"/><path d="M8 2v2"/><circle cx="12" cy="11" r="4"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// Contact is [contact](https://lucide.dev/icons/contact).
var Contact = icon.Lucide(viewBox, paint, `<path d="M16 2v2"/><path d="M7 21v-2a2 2 0 012-2h6a2 2 0 012 2v2"/><path d="M8 2v2"/><circle cx="12" cy="10" r="3"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// Container is [container](https://lucide.dev/icons/container).
var Container = icon.Lucide(viewBox, paint, `<path d="M22 7.7c0-.6-.4-1.2-.8-1.5l-6.3-3.9a1.72 1.72 0 0 0-1.7 0l-10.3 6c-.5.2-.9.8-.9 1.4v6.6c0 .5.4 1.2.8 1.5l6.3 3.9a1.72 1.72 0 0 0 1.7 0l10.3-6c.5-.3.9-1 .9-1.5Z"/><path d="M10 21.9V14L2.1 9.1"/><path d="m10 14 11.9-6.9"/><path d="M14 19.8v-8.1"/><path d="M18 17.5V9.4"/>`)

// Contrast is [contrast](https://lucide.dev/icons/contrast).
var Contrast = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 18a6 6 0 0 0 0-12v12z"/>`)

// Cookie is [cookie](https://lucide.dev/icons/cookie).
var Cookie = icon.Lucide(viewBox, paint, `<path d="M11 17h.01"/><path d="M11.496 2c.324-.016.558.292.529.615a4 4 0 004.235 4.368.713.713 0 01.758.757 4 4 0 004.366 4.237c.323-.03.63.204.614.527a10 10 0 01-2.915 6.566A1 1 0 114.93 4.918 10 10 0 0111.496 2"/><path d="M12 12h.01"/><path d="M16 16h.01"/><path d="M16 3h.01"/><path d="M21 4h.01"/><path d="M21 8h.01"/><path d="M7 14h.01"/><path d="M9 8h.01"/>`)

// CookingPot is [cooking-pot](https://lucide.dev/icons/cooking-pot).
var CookingPot = icon.Lucide(viewBox, paint, `<path d="M2 12h20"/><path d="M20 12v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/><path d="m4 8 16-4"/><path d="m8.86 6.78-.45-1.81a2 2 0 0 1 1.45-2.43l1.94-.48a2 2 0 0 1 2.43 1.46l.45 1.8"/>`)

// CopyCheck is [copy-check](https://lucide.dev/icons/copy-check).
var CopyCheck = icon.Lucide(viewBox, paint, `<path d="m12 15 2 2 4-4"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/>`)

// CopyMinus is [copy-minus](https://lucide.dev/icons/copy-minus).
var CopyMinus = icon.Lucide(viewBox, paint, `<line x1="12" x2="18" y1="15" y2="15"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/>`)

// CopyPlus is [copy-plus](https://lucide.dev/icons/copy-plus).
var CopyPlus = icon.Lucide(viewBox, paint, `<line x1="15" x2="15" y1="12" y2="18"/><line x1="12" x2="18" y1="15" y2="15"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/>`)

// CopySlash is [copy-slash](https://lucide.dev/icons/copy-slash).
var CopySlash = icon.Lucide(viewBox, paint, `<line x1="12" x2="18" y1="18" y2="12"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/>`)

// CopyX is [copy-x](https://lucide.dev/icons/copy-x).
var CopyX = icon.Lucide(viewBox, paint, `<path d="M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2"/><rect x="8" y="8" width="14" height="14" rx="2"/><path d="m12.5 12.5 5 5"/><path d="m12.5 17.5 5-5"/>`)

// Copy is [copy](https://lucide.dev/icons/copy).
var Copy = icon.Lucide(viewBox, paint, `<rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/>`)

// Copyleft is [copyleft](https://lucide.dev/icons/copyleft).
var Copyleft = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M9.17 14.83a4 4 0 1 0 0-5.66"/>`)

// Copyright is [copyright](https://lucide.dev/icons/copyright).
var Copyright = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M14.83 14.83a4 4 0 1 1 0-5.66"/>`)

// CornerDownLeft is [corner-down-left](https://lucide.dev/icons/corner-down-left).
var CornerDownLeft = icon.Lucide(viewBox, paint, `<path d="M20 4v7a4 4 0 0 1-4 4H4"/><path d="m9 10-5 5 5 5"/>`)

// CornerDownRight is [corner-down-right](https://lucide.dev/icons/corner-down-right).
var CornerDownRight = icon.Lucide(viewBox, paint, `<path d="m15 10 5 5-5 5"/><path d="M4 4v7a4 4 0 0 0 4 4h12"/>`)

// CornerLeftDown is [corner-left-down](https://lucide.dev/icons/corner-left-down).
var CornerLeftDown = icon.Lucide(viewBox, paint, `<path d="m14 15-5 5-5-5"/><path d="M20 4h-7a4 4 0 0 0-4 4v12"/>`)

// CornerLeftUp is [corner-left-up](https://lucide.dev/icons/corner-left-up).
var CornerLeftUp = icon.Lucide(viewBox, paint, `<path d="M14 9 9 4 4 9"/><path d="M20 20h-7a4 4 0 0 1-4-4V4"/>`)

// CornerRightDown is [corner-right-down](https://lucide.dev/icons/corner-right-down).
var CornerRightDown = icon.Lucide(viewBox, paint, `<path d="m10 15 5 5 5-5"/><path d="M4 4h7a4 4 0 0 1 4 4v12"/>`)

// CornerRightUp is [corner-right-up](https://lucide.dev/icons/corner-right-up).
var CornerRightUp = icon.Lucide(viewBox, paint, `<path d="m10 9 5-5 5 5"/><path d="M4 20h7a4 4 0 0 0 4-4V4"/>`)

// CornerUpLeft is [corner-up-left](https://lucide.dev/icons/corner-up-left).
var CornerUpLeft = icon.Lucide(viewBox, paint, `<path d="M20 20v-7a4 4 0 0 0-4-4H4"/><path d="M9 14 4 9l5-5"/>`)

// CornerUpRight is [corner-up-right](https://lucide.dev/icons/corner-up-right).
var CornerUpRight = icon.Lucide(viewBox, paint, `<path d="m15 14 5-5-5-5"/><path d="M4 20v-7a4 4 0 0 1 4-4h12"/>`)

// CpuAi is [cpu-ai](https://lucide.dev/icons/cpu-ai).
var CpuAi = icon.Lucide(viewBox, paint, `<path d="M12 20v2"/><path d="M12 2v2"/><path d="M16 9v6"/><path d="M17 20v2"/><path d="M17 2v2"/><path d="M2 12h2"/><path d="M2 17h2"/><path d="M2 7h2"/><path d="M20 12h2"/><path d="M20 17h2"/><path d="M20 7h2"/><path d="M7 20v2"/><path d="M7 2v2"/><path d="M8 13h4"/><path d="M8 15v-4a2 2 0 014 0v4"/><rect x="4" y="4" width="16" height="16" rx="2"/>`)

// Cpu is [cpu](https://lucide.dev/icons/cpu).
var Cpu = icon.Lucide(viewBox, paint, `<path d="M12 20v2"/><path d="M12 2v2"/><path d="M17 20v2"/><path d="M17 2v2"/><path d="M2 12h2"/><path d="M2 17h2"/><path d="M2 7h2"/><path d="M20 12h2"/><path d="M20 17h2"/><path d="M20 7h2"/><path d="M7 20v2"/><path d="M7 2v2"/><rect x="4" y="4" width="16" height="16" rx="2"/><rect x="8" y="8" width="8" height="8" rx="1"/>`)

// CreativeCommons is [creative-commons](https://lucide.dev/icons/creative-commons).
var CreativeCommons = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M10 9.3a2.8 2.8 0 0 0-3.5 1 3.1 3.1 0 0 0 0 3.4 2.7 2.7 0 0 0 3.5 1"/><path d="M17 9.3a2.8 2.8 0 0 0-3.5 1 3.1 3.1 0 0 0 0 3.4 2.7 2.7 0 0 0 3.5 1"/>`)

// CreditCardCheck is [credit-card-check](https://lucide.dev/icons/credit-card-check).
var CreditCardCheck = icon.Lucide(viewBox, paint, `<path d="M12.5 19H4a2 2 0 01-2-2V7a2 2 0 012-2h16a2 2 0 012 2v4"/><path d="M2 10h20"/><path d="M6 14h2"/><path d="m16 17 2 2 4-4"/>`)

// CreditCardMinus is [credit-card-minus](https://lucide.dev/icons/credit-card-minus).
var CreditCardMinus = icon.Lucide(viewBox, paint, `<path d="M22 13V7a2 2 0 00-2-2H4a2 2 0 00-2 2v10a2 2 0 002 2h8.536"/><path d="M22 10H2"/><path d="M6 14h2"/><path d="M16 17h6"/>`)

// CreditCardPlus is [credit-card-plus](https://lucide.dev/icons/credit-card-plus).
var CreditCardPlus = icon.Lucide(viewBox, paint, `<path d="M22 11.354V7a2 2 0 00-2-2H4a2 2 0 00-2 2v10a2 2 0 002 2h8.536"/><path d="M22 10H2"/><path d="M6 14h2"/><path d="M16 17h6"/><path d="M19 14v6"/>`)

// CreditCardReader is [credit-card-reader](https://lucide.dev/icons/credit-card-reader).
var CreditCardReader = icon.Lucide(viewBox, paint, `<path d="M15 16v1"/><path d="M16.963 7.734A1 1 0 0015.999 7H8.003a1 1 0 00-.964.734L4.073 18.467A2 2 0 006 21h12a2 2 0 001.927-2.532z"/><path d="M2.678 8.5A2 2 0 012 7V5a2 2 0 012-2h16a2 2 0 012 2v2a2 2 0 01-.676 1.499"/><path d="m9 21 2-14"/>`)

// CreditCardX is [credit-card-x](https://lucide.dev/icons/credit-card-x).
var CreditCardX = icon.Lucide(viewBox, paint, `<path d="M12.5 19H4a2 2 0 01-2-2V7a2 2 0 012-2h16a2 2 0 012 2v3.5"/><path d="M2 10h20"/><path d="M6 14h2"/><path d="m16.5 14.5 5 5"/><path d="m21.5 14.5-5 5"/>`)

// CreditCard is [credit-card](https://lucide.dev/icons/credit-card).
var CreditCard = icon.Lucide(viewBox, paint, `<rect width="20" height="14" x="2" y="5" rx="2"/><line x1="2" x2="22" y1="10" y2="10"/><path d="M6 14h2"/>`)

// Croissant is [croissant](https://lucide.dev/icons/croissant).
var Croissant = icon.Lucide(viewBox, paint, `<path d="M10.2 18H4.774a1.5 1.5 0 0 1-1.352-.97 11 11 0 0 1 .132-6.487"/><path d="M18 10.2V4.774a1.5 1.5 0 0 0-.97-1.352 11 11 0 0 0-6.486.132"/><path d="M18 5a4 3 0 0 1 4 3 2 2 0 0 1-2 2 10 10 0 0 0-5.139 1.42"/><path d="M5 18a3 4 0 0 0 3 4 2 2 0 0 0 2-2 10 10 0 0 1 1.42-5.14"/><path d="M8.709 2.554a10 10 0 0 0-6.155 6.155 1.5 1.5 0 0 0 .676 1.626l9.807 5.42a2 2 0 0 0 2.718-2.718l-5.42-9.807a1.5 1.5 0 0 0-1.626-.676"/>`)

// Crop is [crop](https://lucide.dev/icons/crop).
var Crop = icon.Lucide(viewBox, paint, `<path d="M6 2v14a2 2 0 0 0 2 2h14"/><path d="M18 22V8a2 2 0 0 0-2-2H2"/>`)

// Cross is [cross](https://lucide.dev/icons/cross).
var Cross = icon.Lucide(viewBox, paint, `<path d="M4 9a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h4a1 1 0 0 1 1 1v4a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-4a1 1 0 0 1 1-1h4a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-4a1 1 0 0 1-1-1V4a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a1 1 0 0 1-1 1z"/>`)

// Crosshair is [crosshair](https://lucide.dev/icons/crosshair).
var Crosshair = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><line x1="22" x2="18" y1="12" y2="12"/><line x1="6" x2="2" y1="12" y2="12"/><line x1="12" x2="12" y1="6" y2="2"/><line x1="12" x2="12" y1="22" y2="18"/>`)

// Crown is [crown](https://lucide.dev/icons/crown).
var Crown = icon.Lucide(viewBox, paint, `<path d="M11.562 3.266a.5.5 0 0 1 .876 0L15.39 8.87a1 1 0 0 0 1.516.294L21.183 5.5a.5.5 0 0 1 .798.519l-2.834 10.246a1 1 0 0 1-.956.734H5.81a1 1 0 0 1-.957-.734L2.02 6.02a.5.5 0 0 1 .798-.519l4.276 3.664a1 1 0 0 0 1.516-.294z"/><path d="M5 21h14"/>`)

// Cuboid is [cuboid](https://lucide.dev/icons/cuboid).
var Cuboid = icon.Lucide(viewBox, paint, `<path d="M10 22v-8"/><path d="M2.336 8.89 10 14l11.715-7.029"/><path d="M22 14a2 2 0 0 1-.971 1.715l-10 6a2 2 0 0 1-2.138-.05l-6-4A2 2 0 0 1 2 16v-6a2 2 0 0 1 .971-1.715l10-6a2 2 0 0 1 2.138.05l6 4A2 2 0 0 1 22 8z"/>`)

// CupSoda is [cup-soda](https://lucide.dev/icons/cup-soda).
var CupSoda = icon.Lucide(viewBox, paint, `<path d="m6 8 1.75 12.28a2 2 0 0 0 2 1.72h4.54a2 2 0 0 0 2-1.72L18 8"/><path d="M5 8h14"/><path d="M7 15a6.47 6.47 0 0 1 5 0 6.47 6.47 0 0 0 5 0"/><path d="m12 8 1-6h2"/>`)

// Cupcake is [cupcake](https://lucide.dev/icons/cupcake).
var Cupcake = icon.Lucide(viewBox, paint, `<path d="M12 22v-9"/><path d="M14 4h1a3 3 0 013 3l-.004.125A4 4 0 0121 11v2"/><path d="m15.5 22 1.5-9"/><path d="M21 13a1 1 0 01.919 1.394l-2.74 6.394A2 2 0 0117.34 22H6.659a2 2 0 01-1.838-1.212l-2.74-6.394A1 1 0 013 13z"/><path d="M3 13v-2a4 4 0 013.003-3.875L6 7a3 3 0 013-3h1"/><path d="M8.5 22 7 13"/><circle cx="12" cy="4" r="2"/>`)

// Currency is [currency](https://lucide.dev/icons/currency).
var Currency = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="8"/><line x1="3" x2="6" y1="3" y2="6"/><line x1="21" x2="18" y1="3" y2="6"/><line x1="3" x2="6" y1="21" y2="18"/><line x1="21" x2="18" y1="21" y2="18"/>`)

// Cylinder is [cylinder](https://lucide.dev/icons/cylinder).
var Cylinder = icon.Lucide(viewBox, paint, `<ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 18 0V5"/>`)

// Dam is [dam](https://lucide.dev/icons/dam).
var Dam = icon.Lucide(viewBox, paint, `<path d="M11 11.31c1.17.56 1.54 1.69 3.5 1.69 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"/><path d="M11.75 18c.35.5 1.45 1 2.75 1 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"/><path d="M2 10h4"/><path d="M2 14h4"/><path d="M2 18h4"/><path d="M2 6h4"/><path d="M7 3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1L10 4a1 1 0 0 0-1-1z"/>`)

// DatabaseArrowDown is [database-arrow-down](https://lucide.dev/icons/database-arrow-down).
var DatabaseArrowDown = icon.Lucide(viewBox, paint, `<path d="m16 19 3 3 3-3"/><path d="M19 16v6"/><path d="M21 12.536V5"/><path d="M3 12A9 3 0 0 0 15.182 14.806"/><path d="M3 5V19A9 3 0 0 0 13.318 21.968"/><ellipse cx="12" cy="5" rx="9" ry="3"/>`)

// DatabaseArrowUp is [database-arrow-up](https://lucide.dev/icons/database-arrow-up).
var DatabaseArrowUp = icon.Lucide(viewBox, paint, `<path d="M19 22v-6"/><path d="M21 12.536V5"/><path d="m22 19-3-3-3 3"/><path d="M3 12A9 3 0 0 0 14.457 14.886"/><path d="M3 5V19A9 3 0 0 0 13.318 21.968"/><ellipse cx="12" cy="5" rx="9" ry="3"/>`)

// DatabaseBackup is [database-backup](https://lucide.dev/icons/database-backup).
var DatabaseBackup = icon.Lucide(viewBox, paint, `<ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 12a9 3 0 0 0 5 2.69"/><path d="M21 9.3V5"/><path d="M3 5v14a9 3 0 0 0 6.47 2.88"/><path d="M12 12v4h4"/><path d="M13 20a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L12 16"/>`)

// DatabaseCheck is [database-check](https://lucide.dev/icons/database-check).
var DatabaseCheck = icon.Lucide(viewBox, paint, `<path d="m16 19 2 2 4-4"/><path d="M21 13.127V5"/><path d="M3 12A9 3 0 0 0 21 12"/><path d="M3 5V19A9 3 0 0 0 13.318 21.968"/><ellipse cx="12" cy="5" rx="9" ry="3"/>`)

// DatabaseMinus is [database-minus](https://lucide.dev/icons/database-minus).
var DatabaseMinus = icon.Lucide(viewBox, paint, `<path d="M21 15V5"/><path d="M22 19h-6"/><path d="M3 12A9 3 0 0 0 21 12"/><path d="M3 5V19A9 3 0 0 0 13.318 21.968"/><ellipse cx="12" cy="5" rx="9" ry="3"/>`)

// DatabasePlus is [database-plus](https://lucide.dev/icons/database-plus).
var DatabasePlus = icon.Lucide(viewBox, paint, `<path d="M19 16v6"/><path d="M21 12.536V5"/><path d="M22 19h-6"/><path d="M3 12A9 3 0 0 0 15.1824 14.8061"/><path d="M3 5V19A9 3 0 0 0 13.318 21.968"/><ellipse cx="12" cy="5" rx="9" ry="3"/>`)

// DatabaseSearch is [database-search](https://lucide.dev/icons/database-search).
var DatabaseSearch = icon.Lucide(viewBox, paint, `<path d="M21 11.693V5"/><path d="m22 22-1.875-1.875"/><path d="M3 12a9 3 0 0 0 8.697 2.998"/><path d="M3 5v14a9 3 0 0 0 9.28 2.999"/><circle cx="18" cy="18" r="3"/><ellipse cx="12" cy="5" rx="9" ry="3"/>`)

// DatabaseX is [database-x](https://lucide.dev/icons/database-x).
var DatabaseX = icon.Lucide(viewBox, paint, `<path d="m17 17 5 5"/><path d="M19.323 13.744A9 3 0 0 0 21 12"/><path d="M21 13.127V5"/><path d="m22 17-5 5"/><path d="M3 12A9 3 0 0 0 13.563 14.954"/><path d="M3 5V19A9 3 0 0 0 13 21.981"/><ellipse cx="12" cy="5" rx="9" ry="3"/>`)

// DatabaseZap is [database-zap](https://lucide.dev/icons/database-zap).
var DatabaseZap = icon.Lucide(viewBox, paint, `<ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5V19A9 3 0 0 0 15 21.84"/><path d="M21 5V8"/><path d="M21 12L18 17H22L19 22"/><path d="M3 12A9 3 0 0 0 14.59 14.87"/>`)

// Database is [database](https://lucide.dev/icons/database).
var Database = icon.Lucide(viewBox, paint, `<ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5V19A9 3 0 0 0 21 19V5"/><path d="M3 12A9 3 0 0 0 21 12"/>`)

// DecimalsArrowLeft is [decimals-arrow-left](https://lucide.dev/icons/decimals-arrow-left).
var DecimalsArrowLeft = icon.Lucide(viewBox, paint, `<path d="m13 21-3-3 3-3"/><path d="M20 18H10"/><path d="M3 11h.01"/><rect x="6" y="3" width="5" height="8" rx="2.5"/>`)

// DecimalsArrowRight is [decimals-arrow-right](https://lucide.dev/icons/decimals-arrow-right).
var DecimalsArrowRight = icon.Lucide(viewBox, paint, `<path d="M10 18h10"/><path d="m17 21 3-3-3-3"/><path d="M3 11h.01"/><rect x="15" y="3" width="5" height="8" rx="2.5"/><rect x="6" y="3" width="5" height="8" rx="2.5"/>`)

// Delete is [delete](https://lucide.dev/icons/delete).
var Delete = icon.Lucide(viewBox, paint, `<path d="M10 5a2 2 0 0 0-1.344.519l-6.328 5.74a1 1 0 0 0 0 1.481l6.328 5.741A2 2 0 0 0 10 19h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z"/><path d="m12 9 6 6"/><path d="m18 9-6 6"/>`)

// Dessert is [dessert](https://lucide.dev/icons/dessert).
var Dessert = icon.Lucide(viewBox, paint, `<path d="M10.162 3.167A10 10 0 0 0 2 13a2 2 0 0 0 4 0v-1a2 2 0 0 1 4 0v4a2 2 0 0 0 4 0v-4a2 2 0 0 1 4 0v1a2 2 0 0 0 4-.006 10 10 0 0 0-8.161-9.826"/><path d="M20.804 14.869a9 9 0 0 1-17.608 0"/><circle cx="12" cy="4" r="2"/>`)

// Diameter is [diameter](https://lucide.dev/icons/diameter).
var Diameter = icon.Lucide(viewBox, paint, `<circle cx="19" cy="19" r="2"/><circle cx="5" cy="5" r="2"/><path d="M6.48 3.66a10 10 0 0 1 13.86 13.86"/><path d="m6.41 6.41 11.18 11.18"/><path d="M3.66 6.48a10 10 0 0 0 13.86 13.86"/>`)

// DiamondMinus is [diamond-minus](https://lucide.dev/icons/diamond-minus).
var DiamondMinus = icon.Lucide(viewBox, paint, `<path d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0z"/><path d="M8 12h8"/>`)

// DiamondPercent is [diamond-percent](https://lucide.dev/icons/diamond-percent).
var DiamondPercent = icon.Lucide(viewBox, paint, `<path d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0Z"/><path d="M9.2 9.2h.01"/><path d="m14.5 9.5-5 5"/><path d="M14.7 14.8h.01"/>`)

// DiamondPlus is [diamond-plus](https://lucide.dev/icons/diamond-plus).
var DiamondPlus = icon.Lucide(viewBox, paint, `<path d="M12 8v8"/><path d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0z"/><path d="M8 12h8"/>`)

// Diamond is [diamond](https://lucide.dev/icons/diamond).
var Diamond = icon.Lucide(viewBox, paint, `<path d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41l-7.59-7.59a2.41 2.41 0 0 0-3.41 0Z"/>`)

// Dice1 is [dice-1](https://lucide.dev/icons/dice-1).
var Dice1 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M12 12h.01"/>`)

// Dice2 is [dice-2](https://lucide.dev/icons/dice-2).
var Dice2 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 9h.01"/><path d="M9 15h.01"/>`)

// Dice3 is [dice-3](https://lucide.dev/icons/dice-3).
var Dice3 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01"/><path d="M12 12h.01"/><path d="M8 16h.01"/>`)

// Dice4 is [dice-4](https://lucide.dev/icons/dice-4).
var Dice4 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01"/><path d="M8 8h.01"/><path d="M8 16h.01"/><path d="M16 16h.01"/>`)

// Dice5 is [dice-5](https://lucide.dev/icons/dice-5).
var Dice5 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01"/><path d="M8 8h.01"/><path d="M8 16h.01"/><path d="M16 16h.01"/><path d="M12 12h.01"/>`)

// Dice6 is [dice-6](https://lucide.dev/icons/dice-6).
var Dice6 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01"/><path d="M16 12h.01"/><path d="M16 16h.01"/><path d="M8 8h.01"/><path d="M8 12h.01"/><path d="M8 16h.01"/>`)

// Dices is [dices](https://lucide.dev/icons/dices).
var Dices = icon.Lucide(viewBox, paint, `<rect width="12" height="12" x="2" y="10" rx="2" ry="2"/><path d="m17.92 14 3.5-3.5a2.24 2.24 0 0 0 0-3l-5-4.92a2.24 2.24 0 0 0-3 0L10 6"/><path d="M6 18h.01"/><path d="M10 14h.01"/><path d="M15 6h.01"/><path d="M18 9h.01"/>`)

// Diff is [diff](https://lucide.dev/icons/diff).
var Diff = icon.Lucide(viewBox, paint, `<path d="M12 3v14"/><path d="M5 10h14"/><path d="M5 21h14"/>`)

// Disc2 is [disc-2](https://lucide.dev/icons/disc-2).
var Disc2 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="4"/><path d="M12 12h.01"/>`)

// Disc3 is [disc-3](https://lucide.dev/icons/disc-3).
var Disc3 = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M6 12c0-1.7.7-3.2 1.8-4.2"/><circle cx="12" cy="12" r="2"/><path d="M18 12c0 1.7-.7 3.2-1.8 4.2"/>`)

// DiscAlbum is [disc-album](https://lucide.dev/icons/disc-album).
var DiscAlbum = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="12" cy="12" r="5"/><path d="M12 12h.01"/>`)

// Disc is [disc](https://lucide.dev/icons/disc).
var Disc = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="2"/>`)

// Divide is [divide](https://lucide.dev/icons/divide).
var Divide = icon.Lucide(viewBox, paint, `<circle cx="12" cy="6" r="1"/><line x1="5" x2="19" y1="12" y2="12"/><circle cx="12" cy="18" r="1"/>`)

// DnaOff is [dna-off](https://lucide.dev/icons/dna-off).
var DnaOff = icon.Lucide(viewBox, paint, `<path d="M15 2c-1.35 1.5-2.092 3-2.5 4.5L14 8"/><path d="m17 6-2.891-2.891"/><path d="M2 15c3.333-3 6.667-3 10-3"/><path d="m2 2 20 20"/><path d="m20 9 .891.891"/><path d="M22 9c-1.5 1.35-3 2.092-4.5 2.5l-1-1"/><path d="M3.109 14.109 4 15"/><path d="m6.5 12.5 1 1"/><path d="m7 18 2.891 2.891"/><path d="M9 22c1.35-1.5 2.092-3 2.5-4.5L10 16"/>`)

// Dna is [dna](https://lucide.dev/icons/dna).
var Dna = icon.Lucide(viewBox, paint, `<path d="m10 16 1.5 1.5"/><path d="m14 8-1.5-1.5"/><path d="M15 2c-1.798 1.998-2.518 3.995-2.807 5.993"/><path d="m16.5 10.5 1 1"/><path d="m17 6-2.891-2.891"/><path d="M2 15c6.667-6 13.333 0 20-6"/><path d="m20 9 .891.891"/><path d="M3.109 14.109 4 15"/><path d="m6.5 12.5 1 1"/><path d="m7 18 2.891 2.891"/><path d="M9 22c1.798-1.998 2.518-3.995 2.807-5.993"/>`)

// Dock is [dock](https://lucide.dev/icons/dock).
var Dock = icon.Lucide(viewBox, paint, `<path d="M2 8h20"/><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M6 16h12"/>`)

// Dog is [dog](https://lucide.dev/icons/dog).
var Dog = icon.Lucide(viewBox, paint, `<path d="M11.25 16.25h1.5L12 17z"/><path d="M16 14v.5"/><path d="M4.42 11.247A13.152 13.152 0 0 0 4 14.556C4 18.728 7.582 21 12 21s8-2.272 8-6.444a11.702 11.702 0 0 0-.493-3.309"/><path d="M8 14v.5"/><path d="M8.5 8.5c-.384 1.05-1.083 2.028-2.344 2.5-1.931.722-3.576-.297-3.656-1-.113-.994 1.177-6.53 4-7 1.923-.321 3.651.845 3.651 2.235A7.497 7.497 0 0 1 14 5.277c0-1.39 1.844-2.598 3.767-2.277 2.823.47 4.113 6.006 4 7-.08.703-1.725 1.722-3.656 1-1.261-.472-1.855-1.45-2.239-2.5"/>`)

// DollarSign is [dollar-sign](https://lucide.dev/icons/dollar-sign).
var DollarSign = icon.Lucide(viewBox, paint, `<line x1="12" x2="12" y1="2" y2="22"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>`)

// Dome is [dome](https://lucide.dev/icons/dome).
var Dome = icon.Lucide(viewBox, paint, `<path d="M10 21v-3a2 2 0 014 0v3"/><path d="M12 2v2"/><path d="M18 12v9"/><path d="M22 19a2 2 0 01-2 2H4a2 2 0 01-2-2v-6a1 1 0 011-1h18a1 1 0 011 1z"/><path d="M4 12a8 8 0 0116 0"/><path d="M6 12v9"/>`)

// Donut is [donut](https://lucide.dev/icons/donut).
var Donut = icon.Lucide(viewBox, paint, `<path d="M20.5 10a2.5 2.5 0 0 1-2.4-3H18a2.95 2.95 0 0 1-2.6-4.4 10 10 0 1 0 6.3 7.1c-.3.2-.8.3-1.2.3"/><circle cx="12" cy="12" r="3"/>`)

// DoorClosedCog is [door-closed-cog](https://lucide.dev/icons/door-closed-cog).
var DoorClosedCog = icon.Lucide(viewBox, paint, `<path d="m14.305 19.53.923-.382"/><path d="m15.229 16.852-.924-.383"/><path d="m16.852 15.228-.383-.923"/><path d="m16.852 20.773-.383.924"/><path d="M19 10.35V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.697-.382-.924"/><path d="M2 21h8.58"/><path d="m20.773 16.852.922-.383"/><path d="m20.773 19.148.922.383"/><path d="M9 12h.01"/><circle cx="18" cy="18" r="3"/>`)

// DoorClosedLocked is [door-closed-locked](https://lucide.dev/icons/door-closed-locked).
var DoorClosedLocked = icon.Lucide(viewBox, paint, `<path d="M19 8V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16"/><path d="M2 21h8"/><path d="M20 16v-2a2 2 0 00-4 0v2"/><path d="M9 12h.01"/><rect x="14" y="16" width="8" height="5" rx="1"/>`)

// DoorClosedPackage is [door-closed-package](https://lucide.dev/icons/door-closed-package).
var DoorClosedPackage = icon.Lucide(viewBox, paint, `<path d="M18 13v3"/><path d="M19 9V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16"/><path d="M2 21h8"/><path d="M9 12h.01"/><rect x="14" y="13" width="8" height="8" rx="1"/>`)

// DoorClosed is [door-closed](https://lucide.dev/icons/door-closed).
var DoorClosed = icon.Lucide(viewBox, paint, `<path d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16"/><path d="M2 21h20"/><path d="M9 12h.01"/>`)

// DoorOpen is [door-open](https://lucide.dev/icons/door-open).
var DoorOpen = icon.Lucide(viewBox, paint, `<path d="M10 21H2"/><path d="M10 3H7a2 2 0 00-2 2v16"/><path d="M14 12h.01"/><path d="M19 21V5a2 2 0 00-1.675-1.974l-6.163-1.013A1 1 0 0010 3v18a1 1 0 001.124.992z"/><path d="M22 21h-3"/>`)

// DoorStairwell is [door-stairwell](https://lucide.dev/icons/door-stairwell).
var DoorStairwell = icon.Lucide(viewBox, paint, `<path d="M12 17v-3a1 1 0 011-1h6"/><path d="M19 17h-9a1 1 0 00-1 1v3"/><path d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16"/><path d="M19 9h-3a1 1 0 00-1 1v3"/><path d="M22 21H2"/>`)

// Dot is [dot](https://lucide.dev/icons/dot).
var Dot = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="1"/>`)

// Download is [download](https://lucide.dev/icons/download).
var Download = icon.Lucide(viewBox, paint, `<path d="M12 15V3"/><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><path d="m7 10 5 5 5-5"/>`)

// DraftingCompass is [drafting-compass](https://lucide.dev/icons/drafting-compass).
var DraftingCompass = icon.Lucide(viewBox, paint, `<path d="m12.99 6.74 1.93 3.44"/><path d="M19.136 12a10 10 0 0 1-14.271 0"/><path d="m21 21-2.16-3.84"/><path d="m3 21 8.02-14.26"/><circle cx="12" cy="5" r="2"/>`)

// Drama is [drama](https://lucide.dev/icons/drama).
var Drama = icon.Lucide(viewBox, paint, `<path d="M10 11h.01"/><path d="M14 6h.01"/><path d="M18 6h.01"/><path d="M6.5 13.1h.01"/><path d="M22 5c0 9-4 12-6 12s-6-3-6-12c0-2 2-3 6-3s6 1 6 3"/><path d="M17.4 9.9c-.8.8-2 .8-2.8 0"/><path d="M10.1 7.1C9 7.2 7.7 7.7 6 8.6c-3.5 2-4.7 3.9-3.7 5.6 4.5 7.8 9.5 8.4 11.2 7.4.9-.5 1.9-2.1 1.9-4.7"/><path d="M9.1 16.5c.3-1.1 1.4-1.7 2.4-1.4"/>`)

// Drill is [drill](https://lucide.dev/icons/drill).
var Drill = icon.Lucide(viewBox, paint, `<path d="M10 18a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a3 3 0 0 1-3-3 1 1 0 0 1 1-1z"/><path d="M13 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1l-.81 3.242a1 1 0 0 1-.97.758H8"/><path d="M14 4h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-3"/><path d="M18 6h4"/><path d="m5 10-2 8"/><path d="m7 18 2-8"/>`)

// Drone is [drone](https://lucide.dev/icons/drone).
var Drone = icon.Lucide(viewBox, paint, `<path d="M10 10 7 7"/><path d="m10 14-3 3"/><path d="m14 10 3-3"/><path d="m14 14 3 3"/><path d="M14.205 4.139a4 4 0 1 1 5.439 5.863"/><path d="M19.637 14a4 4 0 1 1-5.432 5.868"/><path d="M4.367 10a4 4 0 1 1 5.438-5.862"/><path d="M9.795 19.862a4 4 0 1 1-5.429-5.873"/><rect x="10" y="8" width="4" height="8" rx="1"/>`)

// DropletOff is [droplet-off](https://lucide.dev/icons/droplet-off).
var DropletOff = icon.Lucide(viewBox, paint, `<path d="M18.715 13.186C18.29 11.858 17.384 10.607 16 9.5c-2-1.6-3.5-4-4-6.5a10.7 10.7 0 0 1-.884 2.586"/><path d="m2 2 20 20"/><path d="M8.795 8.797A11 11 0 0 1 8 9.5C6 11.1 5 13 5 15a7 7 0 0 0 13.222 3.208"/>`)

// Droplet is [droplet](https://lucide.dev/icons/droplet).
var Droplet = icon.Lucide(viewBox, paint, `<path d="M12 22a7 7 0 0 0 7-7c0-2-1-3.9-3-5.5s-3.5-4-4-6.5c-.5 2.5-2 4.9-4 6.5C6 11.1 5 13 5 15a7 7 0 0 0 7 7z"/>`)

// Droplets is [droplets](https://lucide.dev/icons/droplets).
var Droplets = icon.Lucide(viewBox, paint, `<path d="M7 16.3c2.2 0 4-1.83 4-4.05 0-1.16-.57-2.26-1.71-3.19S7.29 6.75 7 5.3c-.29 1.45-1.14 2.84-2.29 3.76S3 11.1 3 12.25c0 2.22 1.8 4.05 4 4.05z"/><path d="M12.56 6.6A10.97 10.97 0 0 0 14 3.02c.5 2.5 2 4.9 4 6.5s3 3.5 3 5.5a6.98 6.98 0 0 1-11.91 4.97"/>`)

// Drum is [drum](https://lucide.dev/icons/drum).
var Drum = icon.Lucide(viewBox, paint, `<path d="m2 2 8 8"/><path d="m22 2-8 8"/><ellipse cx="12" cy="9" rx="10" ry="5"/><path d="M7 13.4v7.9"/><path d="M12 14v8"/><path d="M17 13.4v7.9"/><path d="M2 9v8a10 5 0 0 0 20 0V9"/>`)

// Drumstick is [drumstick](https://lucide.dev/icons/drumstick).
var Drumstick = icon.Lucide(viewBox, paint, `<path d="M15.4 15.63a7.875 6 135 1 1 6.23-6.23 4.5 3.43 135 0 0-6.23 6.23"/><path d="m8.29 12.71-2.6 2.6a2.5 2.5 0 1 0-1.65 4.65A2.5 2.5 0 1 0 8.7 18.3l2.59-2.59"/>`)

// Dumbbell is [dumbbell](https://lucide.dev/icons/dumbbell).
var Dumbbell = icon.Lucide(viewBox, paint, `<path d="M17.596 12.768a2 2 0 1 0 2.829-2.829l-1.768-1.767a2 2 0 0 0 2.828-2.829l-2.828-2.828a2 2 0 0 0-2.829 2.828l-1.767-1.768a2 2 0 1 0-2.829 2.829z"/><path d="m2.5 21.5 1.4-1.4"/><path d="m20.1 3.9 1.4-1.4"/><path d="M5.343 21.485a2 2 0 1 0 2.829-2.828l1.767 1.768a2 2 0 1 0 2.829-2.829l-6.364-6.364a2 2 0 1 0-2.829 2.829l1.768 1.767a2 2 0 0 0-2.828 2.829z"/><path d="m9.6 14.4 4.8-4.8"/>`)

// EarOff is [ear-off](https://lucide.dev/icons/ear-off).
var EarOff = icon.Lucide(viewBox, paint, `<path d="M6 18.5a3.5 3.5 0 1 0 7 0c0-1.57.92-2.52 2.04-3.46"/><path d="M6 8.5c0-.75.13-1.47.36-2.14"/><path d="M8.8 3.15A6.5 6.5 0 0 1 19 8.5c0 1.63-.44 2.81-1.09 3.76"/><path d="M12.5 6A2.5 2.5 0 0 1 15 8.5M10 13a2 2 0 0 0 1.82-1.18"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Ear is [ear](https://lucide.dev/icons/ear).
var Ear = icon.Lucide(viewBox, paint, `<path d="M6 8.5a6.5 6.5 0 1 1 13 0c0 6-6 6-6 10a3.5 3.5 0 1 1-7 0"/><path d="M15 8.5a2.5 2.5 0 0 0-5 0v1a2 2 0 1 1 0 4"/>`)

// EarthLock is [earth-lock](https://lucide.dev/icons/earth-lock).
var EarthLock = icon.Lucide(viewBox, paint, `<path d="M7 3.34V5a3 3 0 0 0 3 3"/><path d="M11 21.95V18a2 2 0 0 0-2-2 2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"/><path d="M21.54 15H17a2 2 0 0 0-2 2v4.54"/><path d="M12 2a10 10 0 1 0 9.54 13"/><path d="M20 6V4a2 2 0 1 0-4 0v2"/><rect width="8" height="5" x="14" y="6" rx="1"/>`)

// Earth is [earth](https://lucide.dev/icons/earth).
var Earth = icon.Lucide(viewBox, paint, `<path d="M21.54 15H17a2 2 0 0 0-2 2v4.54"/><path d="M7 3.34V5a3 3 0 0 0 3 3a2 2 0 0 1 2 2c0 1.1.9 2 2 2a2 2 0 0 0 2-2c0-1.1.9-2 2-2h3.17"/><path d="M11 21.95V18a2 2 0 0 0-2-2a2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"/><circle cx="12" cy="12" r="10"/>`)

// Eclipse is [eclipse](https://lucide.dev/icons/eclipse).
var Eclipse = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 2a7 7 0 1 0 10 10"/>`)

// EggFried is [egg-fried](https://lucide.dev/icons/egg-fried).
var EggFried = icon.Lucide(viewBox, paint, `<circle cx="11.5" cy="12.5" r="3.5"/><path d="M3 8c0-3.5 2.5-6 6.5-6 5 0 4.83 3 7.5 5s5 2 5 6c0 4.5-2.5 6.5-7 6.5-2.5 0-2.5 2.5-6 2.5s-7-2-7-5.5c0-3 1.5-3 1.5-5C3.5 10 3 9 3 8Z"/>`)

// EggOff is [egg-off](https://lucide.dev/icons/egg-off).
var EggOff = icon.Lucide(viewBox, paint, `<path d="m2 2 20 20"/><path d="M20 14.347V14c0-6-4-12-8-12-1.078 0-2.157.436-3.157 1.19"/><path d="M6.206 6.21C4.871 8.4 4 11.2 4 14a8 8 0 0 0 14.568 4.568"/>`)

// Egg is [egg](https://lucide.dev/icons/egg).
var Egg = icon.Lucide(viewBox, paint, `<path d="M12 2C8 2 4 8 4 14a8 8 0 0 0 16 0c0-6-4-12-8-12"/>`)

// EighthNote is [eighth-note](https://lucide.dev/icons/eighth-note).
var EighthNote = icon.Lucide(viewBox, paint, `<path d="m19 6-5.2-3.9a.5.5 0 0 0-.8.4V18"/><circle cx="9" cy="18" r="4"/>`)

// EighthNotes is [eighth-notes](https://lucide.dev/icons/eighth-notes).
var EighthNotes = icon.Lucide(viewBox, paint, `<path d="M21 16V3a1 1 0 0 0-1.28-.96l-10 2.917a1 1 0 0 0-.72.96V19"/><circle cx="18" cy="16" r="3"/><circle cx="6" cy="19" r="3"/>`)

// Eject is [eject](https://lucide.dev/icons/eject).
var Eject = icon.Lucide(viewBox, paint, `<path d="M4 13a1 1 0 0 1-.72-1.695l7.257-7.668a2 2 0 0 1 2.926 0l7.256 7.668A1 1 0 0 1 20 13z"/><rect x="3" y="17" width="18" height="4" rx="1"/>`)

// Ellipse is [ellipse](https://lucide.dev/icons/ellipse).
var Ellipse = icon.Lucide(viewBox, paint, `<ellipse cx="12" cy="12" rx="10" ry="6"/>`)

// EllipsisVertical is [ellipsis-vertical](https://lucide.dev/icons/ellipsis-vertical).
var EllipsisVertical = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="1"/><circle cx="12" cy="5" r="1"/><circle cx="12" cy="19" r="1"/>`)

// Ellipsis is [ellipsis](https://lucide.dev/icons/ellipsis).
var Ellipsis = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/>`)

// Engine is [engine](https://lucide.dev/icons/engine).
var Engine = icon.Lucide(viewBox, paint, `<path d="M10 3h6"/><path d="M13 3v4"/><path d="M2 10v6"/><path d="M2 13h4"/><path d="M6 16a2 2 0 002 2h1a2 2 0 011.6.8l.3.4a2 2 0 001.6.8h2.264a2 2 0 001.789-1.106l1.67-3.341a1 1 0 01.895-.553H21a1 1 0 001-1v-4a1 1 0 00-1-1h-3.5a1 1 0 01-.8-.4l-.9-1.2A1 1 0 0015 7h-4a1 1 0 00-.8.4l-.9 1.2a1 1 0 01-.8.4H7a1 1 0 00-1 1z"/>`)

// EqualApproximatelyNot is [equal-approximately-not](https://lucide.dev/icons/equal-approximately-not).
var EqualApproximatelyNot = icon.Lucide(viewBox, paint, `<path d="M5 15a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"/><path d="M5 9a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"/><line x1="19" x2="5" y1="5" y2="19"/>`)

// EqualApproximately is [equal-approximately](https://lucide.dev/icons/equal-approximately).
var EqualApproximately = icon.Lucide(viewBox, paint, `<path d="M5 15a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"/><path d="M5 9a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"/>`)

// EqualNot is [equal-not](https://lucide.dev/icons/equal-not).
var EqualNot = icon.Lucide(viewBox, paint, `<line x1="5" x2="19" y1="9" y2="9"/><line x1="5" x2="19" y1="15" y2="15"/><line x1="19" x2="5" y1="5" y2="19"/>`)

// Equal is [equal](https://lucide.dev/icons/equal).
var Equal = icon.Lucide(viewBox, paint, `<line x1="5" x2="19" y1="9" y2="9"/><line x1="5" x2="19" y1="15" y2="15"/>`)

// Eraser is [eraser](https://lucide.dev/icons/eraser).
var Eraser = icon.Lucide(viewBox, paint, `<path d="M21 21H8a2 2 0 0 1-1.42-.587l-3.994-3.999a2 2 0 0 1 0-2.828l10-10a2 2 0 0 1 2.829 0l5.999 6a2 2 0 0 1 0 2.828L12.834 21"/><path d="m5.082 11.09 8.828 8.828"/>`)

// EthernetPort is [ethernet-port](https://lucide.dev/icons/ethernet-port).
var EthernetPort = icon.Lucide(viewBox, paint, `<path d="M10 8v1"/><path d="M14 8v1"/><path d="M18 8v1"/><path d="M19 17a2 2 0 00-1.765 1.059l-.47.882A2 2 0 0115 20H9a2 2 0 01-1.765-1.059l-.47-.882A2 2 0 005 17H4a2 2 0 01-2-2V6a2 2 0 012-2h16a2 2 0 012 2v9a2 2 0 01-2 2z"/><path d="M6 8v1"/>`)

// Euro is [euro](https://lucide.dev/icons/euro).
var Euro = icon.Lucide(viewBox, paint, `<path d="M4 10h12"/><path d="M4 14h9"/><path d="M19 6a7.7 7.7 0 0 0-5.2-2A7.9 7.9 0 0 0 6 12c0 4.4 3.5 8 7.8 8 2 0 3.8-.8 5.2-2"/>`)

// EvCharger is [ev-charger](https://lucide.dev/icons/ev-charger).
var EvCharger = icon.Lucide(viewBox, paint, `<path d="M14 13h2a2 2 0 0 1 2 2v2a2 2 0 0 0 4 0v-6.998a2 2 0 0 0-.59-1.42L18 5"/><path d="M14 21V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16"/><path d="M2 21h13"/><path d="M3 7h11"/><path d="m9 11-2 3h3l-2 3"/>`)

// Expand is [expand](https://lucide.dev/icons/expand).
var Expand = icon.Lucide(viewBox, paint, `<path d="m15 15 6 6"/><path d="m15 9 6-6"/><path d="M21 16v5h-5"/><path d="M21 8V3h-5"/><path d="M3 16v5h5"/><path d="m3 21 6-6"/><path d="M3 8V3h5"/><path d="M9 9 3 3"/>`)

// ExternalLink is [external-link](https://lucide.dev/icons/external-link).
var ExternalLink = icon.Lucide(viewBox, paint, `<path d="M15 3h6v6"/><path d="M10 14 21 3"/><path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"/>`)

// EyeClosed is [eye-closed](https://lucide.dev/icons/eye-closed).
var EyeClosed = icon.Lucide(viewBox, paint, `<path d="m15 18-.722-3.25"/><path d="M2 8a10.645 10.645 0 0 0 20 0"/><path d="m20 15-1.726-2.05"/><path d="m4 15 1.726-2.05"/><path d="m9 18 .722-3.25"/>`)

// EyeDashed is [eye-dashed](https://lucide.dev/icons/eye-dashed).
var EyeDashed = icon.Lucide(viewBox, paint, `<path d="M13.054 18.946a11 11 0 0 1-2.11 0"/><path d="M13.054 5.054a11 11 0 0 0-2.11-.001"/><path d="M17.072 6.274a11 11 0 0 1 1.753 1.173"/><path d="M18.825 16.552a11 11 0 0 1-1.753 1.174"/><path d="M2.514 13.303a11 11 0 0 1-.452-.954 1 1 0 0 1 0-.697 11 11 0 0 1 .45-.955"/><path d="M21.485 10.697a11 11 0 0 1 .453.955 1 1 0 0 1 0 .697 11 11 0 0 1-.453.954"/><path d="M5.173 7.448a11 11 0 0 1 1.753-1.174"/><path d="M6.926 17.726a11 11 0 0 1-1.753-1.174"/><circle cx="12" cy="12" r="3"/>`)

// EyeOff is [eye-off](https://lucide.dev/icons/eye-off).
var EyeOff = icon.Lucide(viewBox, paint, `<path d="M10.733 5.076a10.744 10.744 0 0 1 11.205 6.575 1 1 0 0 1 0 .696 10.747 10.747 0 0 1-1.444 2.49"/><path d="M14.084 14.158a3 3 0 0 1-4.242-4.242"/><path d="M17.479 17.499a10.75 10.75 0 0 1-15.417-5.151 1 1 0 0 1 0-.696 10.75 10.75 0 0 1 4.446-5.143"/><path d="m2 2 20 20"/>`)

// Eye is [eye](https://lucide.dev/icons/eye).
var Eye = icon.Lucide(viewBox, paint, `<path d="M2.062 12.348a1 1 0 0 1 0-.696 10.75 10.75 0 0 1 19.876 0 1 1 0 0 1 0 .696 10.75 10.75 0 0 1-19.876 0"/><circle cx="12" cy="12" r="3"/>`)

// FaceAngry is [face-angry](https://lucide.dev/icons/face-angry).
var FaceAngry = icon.Lucide(viewBox, paint, `<path d="M15 12v-1.584"/><path d="M17 10a5 5 0 00-3 1"/><path d="M7 10a5 5 0 013 1"/><path d="M9 12v-1.584"/><path d="M9 17a5 5 0 016.001 0"/><circle cx="12" cy="12" r="10"/>`)

// FaceExpressionless is [face-expressionless](https://lucide.dev/icons/face-expressionless).
var FaceExpressionless = icon.Lucide(viewBox, paint, `<path d="M14 10h2"/><path d="M8 10h2"/><path d="M8 16h8"/><circle cx="12" cy="12" r="10"/>`)

// FaceGrinning is [face-grinning](https://lucide.dev/icons/face-grinning).
var FaceGrinning = icon.Lucide(viewBox, paint, `<path d="M15 10V9"/><path d="M7.084 14.302a5.12 5.12 0 009.833 0 .24.24 0 00-.235-.302H7.32a.24.24 0 00-.235.302"/><path d="M9 10V9"/><circle cx="12" cy="12" r="10"/>`)

// FaceNeutral is [face-neutral](https://lucide.dev/icons/face-neutral).
var FaceNeutral = icon.Lucide(viewBox, paint, `<path d="M15 10V9"/><path d="M8 16h8"/><path d="M9 10V9"/><circle cx="12" cy="12" r="10"/>`)

// FaceSlightlyFrowning is [face-slightly-frowning](https://lucide.dev/icons/face-slightly-frowning).
var FaceSlightlyFrowning = icon.Lucide(viewBox, paint, `<path d="M15 10V9"/><path d="M9 10V9"/><path d="M9 16a5 5 0 016 0"/><circle cx="12" cy="12" r="10"/>`)

// FaceSlightlySmilingPlus is [face-slightly-smiling-plus](https://lucide.dev/icons/face-slightly-smiling-plus).
var FaceSlightlySmilingPlus = icon.Lucide(viewBox, paint, `<path d="M13.267 2.08a10 10 0 108.653 8.653"/><path d="M15 10V9"/><path d="M16 5h6"/><path d="M16.472 15a6 6 0 01-8.943 0"/><path d="M19 2v6"/><path d="M9 10V9"/>`)

// FaceSlightlySmiling is [face-slightly-smiling](https://lucide.dev/icons/face-slightly-smiling).
var FaceSlightlySmiling = icon.Lucide(viewBox, paint, `<path d="M15 10V9"/><path d="M16.472 15a6 6 0 01-8.943 0"/><path d="M9 10V9"/><circle cx="12" cy="12" r="10"/>`)

// Factory is [factory](https://lucide.dev/icons/factory).
var Factory = icon.Lucide(viewBox, paint, `<path d="M12 16h.01"/><path d="M16 16h.01"/><path d="M3 19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8.5a.5.5 0 0 0-.769-.422l-4.462 2.844A.5.5 0 0 1 15 10.5v-2a.5.5 0 0 0-.769-.422L9.77 10.922A.5.5 0 0 1 9 10.5V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2z"/><path d="M8 16h.01"/>`)

// Fan is [fan](https://lucide.dev/icons/fan).
var Fan = icon.Lucide(viewBox, paint, `<path d="M10.827 16.379a6.082 6.082 0 0 1-8.618-7.002l5.412 1.45a6.082 6.082 0 0 1 7.002-8.618l-1.45 5.412a6.082 6.082 0 0 1 8.618 7.002l-5.412-1.45a6.082 6.082 0 0 1-7.002 8.618l1.45-5.412Z"/><path d="M12 12v.01"/>`)

// FastForward is [fast-forward](https://lucide.dev/icons/fast-forward).
var FastForward = icon.Lucide(viewBox, paint, `<path d="M12 6a2 2 0 0 1 3.414-1.414l6 6a2 2 0 0 1 0 2.828l-6 6A2 2 0 0 1 12 18z"/><path d="M2 6a2 2 0 0 1 3.414-1.414l6 6a2 2 0 0 1 0 2.828l-6 6A2 2 0 0 1 2 18z"/>`)

// Faucet is [faucet](https://lucide.dev/icons/faucet).
var Faucet = icon.Lucide(viewBox, paint, `<path d="M10.083 5.428 5.57 4.083a2 2 0 10.001 3.834l4.512-1.345"/><path d="M12 8v3"/><path d="m13.917 5.428 4.511-1.345a2 2 0 110 3.834l-4.51-1.345"/><path d="M18 17v-4.006"/><path d="M22 11v8"/><path d="M22 12h-3a1 1 0 00-1 .994h-2.539a4 4 0 00-6.915-.012L7 13a5 5 0 00-5 5v1a1 1 0 001 1h2a1 1 0 001-1v-1a1 1 0 01.995-1l1.552.018a4 4 0 006.907 0L18 17a1 1 0 001 1h3"/><circle cx="12" cy="6" r="2"/>`)

// Feather is [feather](https://lucide.dev/icons/feather).
var Feather = icon.Lucide(viewBox, paint, `<path d="M14.086 18.412A2 2 0 0112.67 19H5v-7.672a2 2 0 01.586-1.414L11.75 3.75a6 6 0 118.49 8.49z"/><path d="M16 8 2 22"/><path d="M17.488 15H9"/>`)

// Fence is [fence](https://lucide.dev/icons/fence).
var Fence = icon.Lucide(viewBox, paint, `<path d="M4 3 2 5v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"/><path d="M6 8h4"/><path d="M6 18h4"/><path d="m12 3-2 2v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"/><path d="M14 8h4"/><path d="M14 18h4"/><path d="m20 3-2 2v15c0 .6.4 1 1 1h2c.6 0 1-.4 1-1V5Z"/>`)

// FerrisWheel is [ferris-wheel](https://lucide.dev/icons/ferris-wheel).
var FerrisWheel = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="2"/><path d="M12 2v4"/><path d="m6.8 15-3.5 2"/><path d="m20.7 7-3.5 2"/><path d="M6.8 9 3.3 7"/><path d="m20.7 17-3.5-2"/><path d="m9 22 3-8 3 8"/><path d="M8 22h8"/><path d="M18 18.7a9 9 0 1 0-12 0"/>`)

// FileArchive is [file-archive](https://lucide.dev/icons/file-archive).
var FileArchive = icon.Lucide(viewBox, paint, `<path d="M13.659 22H18a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v11.5"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 12v-1"/><path d="M8 18v-2"/><path d="M8 7V6"/><circle cx="8" cy="20" r="2"/>`)

// FileAxis3d is [file-axis-3d](https://lucide.dev/icons/file-axis-3d).
var FileAxis3d = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m8 18 4-4"/><path d="M8 10v8h8"/>`)

// FileBadge is [file-badge](https://lucide.dev/icons/file-badge).
var FileBadge = icon.Lucide(viewBox, paint, `<path d="M13 22h5a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v3.3"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m7.69 16.479 1.29 4.88a.5.5 0 0 1-.698.591l-1.843-.849a1 1 0 0 0-.879.001l-1.846.85a.5.5 0 0 1-.692-.593l1.29-4.88"/><circle cx="6" cy="14" r="3"/>`)

// FileBox is [file-box](https://lucide.dev/icons/file-box).
var FileBox = icon.Lucide(viewBox, paint, `<path d="M14 2v5a1 1 0 001 1h5"/><path d="M14.692 22H18a2 2 0 002-2V8a2.4 2.4 0 00-.706-1.706l-3.588-3.588A2.4 2.4 0 0014 2H6a2 2 0 00-2 2v3.804"/><path d="M2.264 13.752 7 16.5l4.737-2.748"/><path d="M2.995 13.014A2 2 0 002 14.744v3.516a2 2 0 00.996 1.73l3 1.74a2 2 0 002.008 0l3-1.74A2 2 0 0012 18.26v-3.517a2 2 0 00-.995-1.73l-3-1.742a2 2 0 00-1.892-.064z"/><path d="M7 16.5V22"/>`)

// FileBracesCorner is [file-braces-corner](https://lucide.dev/icons/file-braces-corner).
var FileBracesCorner = icon.Lucide(viewBox, paint, `<path d="M14 22h4a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v6"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M5 14a1 1 0 0 0-1 1v2a1 1 0 0 1-1 1 1 1 0 0 1 1 1v2a1 1 0 0 0 1 1"/><path d="M9 22a1 1 0 0 0 1-1v-2a1 1 0 0 1 1-1 1 1 0 0 1-1-1v-2a1 1 0 0 0-1-1"/>`)

// FileBraces is [file-braces](https://lucide.dev/icons/file-braces).
var FileBraces = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M10 12a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1 1 1 0 0 1 1 1v1a1 1 0 0 0 1 1"/><path d="M14 18a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1 1 1 0 0 1-1-1v-1a1 1 0 0 0-1-1"/>`)

// FileChartColumnIncreasing is [file-chart-column-increasing](https://lucide.dev/icons/file-chart-column-increasing).
var FileChartColumnIncreasing = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 18v-2"/><path d="M12 18v-4"/><path d="M16 18v-6"/>`)

// FileChartColumn is [file-chart-column](https://lucide.dev/icons/file-chart-column).
var FileChartColumn = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 18v-1"/><path d="M12 18v-6"/><path d="M16 18v-3"/>`)

// FileChartLine is [file-chart-line](https://lucide.dev/icons/file-chart-line).
var FileChartLine = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m16 13-3.5 3.5-2-2L8 17"/>`)

// FileChartPie is [file-chart-pie](https://lucide.dev/icons/file-chart-pie).
var FileChartPie = icon.Lucide(viewBox, paint, `<path d="M15.941 22H18a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.704l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v3.512"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M4.017 11.512a6 6 0 1 0 8.466 8.475"/><path d="M9 16a1 1 0 0 1-1-1v-4c0-.552.45-1.008.995-.917a6 6 0 0 1 4.922 4.922c.091.544-.365.995-.917.995z"/>`)

// FileCheckCorner is [file-check-corner](https://lucide.dev/icons/file-check-corner).
var FileCheckCorner = icon.Lucide(viewBox, paint, `<path d="M10.5 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v6"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m14 20 2 2 4-4"/>`)

// FileCheck is [file-check](https://lucide.dev/icons/file-check).
var FileCheck = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m9 15 2 2 4-4"/>`)

// FileClock is [file-clock](https://lucide.dev/icons/file-clock).
var FileClock = icon.Lucide(viewBox, paint, `<path d="M16 22h2a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v2.85"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 14v2.2l1.6 1"/><circle cx="8" cy="16" r="6"/>`)

// FileCodeCorner is [file-code-corner](https://lucide.dev/icons/file-code-corner).
var FileCodeCorner = icon.Lucide(viewBox, paint, `<path d="M4 12.15V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2h-3.35"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m5 16-3 3 3 3"/><path d="m9 22 3-3-3-3"/>`)

// FileCode is [file-code](https://lucide.dev/icons/file-code).
var FileCode = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M10 12.5 8 15l2 2.5"/><path d="m14 12.5 2 2.5-2 2.5"/>`)

// FileCog is [file-cog](https://lucide.dev/icons/file-cog).
var FileCog = icon.Lucide(viewBox, paint, `<path d="M15 8a1 1 0 0 1-1-1V2a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8z"/><path d="M20 8v12a2 2 0 0 1-2 2h-4.182"/><path d="m3.305 19.53.923-.382"/><path d="M4 10.592V4a2 2 0 0 1 2-2h8"/><path d="m4.228 16.852-.924-.383"/><path d="m5.852 15.228-.383-.923"/><path d="m5.852 20.772-.383.924"/><path d="m8.148 15.228.383-.923"/><path d="m8.53 21.696-.382-.924"/><path d="m9.773 16.852.922-.383"/><path d="m9.773 19.148.922.383"/><circle cx="7" cy="18" r="3"/>`)

// FileDiff is [file-diff](https://lucide.dev/icons/file-diff).
var FileDiff = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M9 10h6"/><path d="M12 13V7"/><path d="M9 17h6"/>`)

// FileDigit is [file-digit](https://lucide.dev/icons/file-digit).
var FileDigit = icon.Lucide(viewBox, paint, `<path d="M4 12V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M10 16h2v6"/><path d="M10 22h4"/><rect x="2" y="16" width="4" height="6" rx="2"/>`)

// FileDown is [file-down](https://lucide.dev/icons/file-down).
var FileDown = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M12 18v-6"/><path d="m9 15 3 3 3-3"/>`)

// FileExclamationPoint is [file-exclamation-point](https://lucide.dev/icons/file-exclamation-point).
var FileExclamationPoint = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M12 9v4"/><path d="M12 17h.01"/>`)

// FileHeadphone is [file-headphone](https://lucide.dev/icons/file-headphone).
var FileHeadphone = icon.Lucide(viewBox, paint, `<path d="M4 6.835V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2h-.343"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M2 19a2 2 0 0 1 4 0v1a2 2 0 0 1-4 0v-4a6 6 0 0 1 12 0v4a2 2 0 0 1-4 0v-1a2 2 0 0 1 4 0"/>`)

// FileHeart is [file-heart](https://lucide.dev/icons/file-heart).
var FileHeart = icon.Lucide(viewBox, paint, `<path d="M13 22h5a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v7"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M3.62 18.8A2.25 2.25 0 1 1 7 15.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a1 1 0 0 1-1.507 0z"/>`)

// FileImage is [file-image](https://lucide.dev/icons/file-image).
var FileImage = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><circle cx="10" cy="12" r="2"/><path d="m20 17-1.296-1.296a2.41 2.41 0 0 0-3.408 0L9 22"/>`)

// FileInput is [file-input](https://lucide.dev/icons/file-input).
var FileInput = icon.Lucide(viewBox, paint, `<path d="M4 11V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-1"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M2 15h10"/><path d="m9 18 3-3-3-3"/>`)

// FileKey is [file-key](https://lucide.dev/icons/file-key).
var FileKey = icon.Lucide(viewBox, paint, `<path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M4 12v6"/><path d="M4 14h2"/><path d="M9.65 22H18a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v4"/><circle cx="4" cy="20" r="2"/>`)

// FileLock is [file-lock](https://lucide.dev/icons/file-lock).
var FileLock = icon.Lucide(viewBox, paint, `<path d="M4 9.8V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2h-3"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M9 17v-2a2 2 0 0 0-4 0v2"/><rect width="8" height="5" x="3" y="17" rx="1"/>`)

// FileMinusCorner is [file-minus-corner](https://lucide.dev/icons/file-minus-corner).
var FileMinusCorner = icon.Lucide(viewBox, paint, `<path d="M20 14V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M14 18h6"/>`)

// FileMinus is [file-minus](https://lucide.dev/icons/file-minus).
var FileMinus = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M9 15h6"/>`)

// FileMusic is [file-music](https://lucide.dev/icons/file-music).
var FileMusic = icon.Lucide(viewBox, paint, `<path d="M11.65 22H18a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v10.35"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 20v-7l3 1.474"/><circle cx="6" cy="20" r="2"/>`)

// FileOutput is [file-output](https://lucide.dev/icons/file-output).
var FileOutput = icon.Lucide(viewBox, paint, `<path d="M4.226 20.925A2 2 0 0 0 6 22h12a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v3.127"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m5 11-3 3"/><path d="m5 17-3-3h10"/>`)

// FilePenLine is [file-pen-line](https://lucide.dev/icons/file-pen-line).
var FilePenLine = icon.Lucide(viewBox, paint, `<path d="M14.364 13.634a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506l4.013-4.009a1 1 0 0 0-3.004-3.004z"/><path d="M14.487 7.858A1 1 0 0 1 14 7V2"/><path d="M20 19.645V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l2.516 2.516"/><path d="M8 18h1"/>`)

// FilePen is [file-pen](https://lucide.dev/icons/file-pen).
var FilePen = icon.Lucide(viewBox, paint, `<path d="M12.659 22H18a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v9.34"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M10.378 12.622a1 1 0 0 1 3 3.003L8.36 20.637a2 2 0 0 1-.854.506l-2.867.837a.5.5 0 0 1-.62-.62l.836-2.869a2 2 0 0 1 .506-.853z"/>`)

// FilePlay is [file-play](https://lucide.dev/icons/file-play).
var FilePlay = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M15.033 13.44a.647.647 0 0 1 0 1.12l-4.065 2.352a.645.645 0 0 1-.968-.56v-4.704a.645.645 0 0 1 .967-.56z"/>`)

// FilePlusCorner is [file-plus-corner](https://lucide.dev/icons/file-plus-corner).
var FilePlusCorner = icon.Lucide(viewBox, paint, `<path d="M11.35 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v5.35"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M14 19h6"/><path d="M17 16v6"/>`)

// FilePlus is [file-plus](https://lucide.dev/icons/file-plus).
var FilePlus = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M9 15h6"/><path d="M12 18v-6"/>`)

// FileQuestionMark is [file-question-mark](https://lucide.dev/icons/file-question-mark).
var FileQuestionMark = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M12 17h.01"/><path d="M9.1 9a3 3 0 0 1 5.82 1c0 2-3 3-3 3"/>`)

// FileScan is [file-scan](https://lucide.dev/icons/file-scan).
var FileScan = icon.Lucide(viewBox, paint, `<path d="M20 10V8a2.4 2.4 0 0 0-.706-1.704l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h4.35"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M16 14a2 2 0 0 0-2 2"/><path d="M16 22a2 2 0 0 1-2-2"/><path d="M20 14a2 2 0 0 1 2 2"/><path d="M20 22a2 2 0 0 0 2-2"/>`)

// FileSearchCorner is [file-search-corner](https://lucide.dev/icons/file-search-corner).
var FileSearchCorner = icon.Lucide(viewBox, paint, `<path d="M11.1 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.589 3.588A2.4 2.4 0 0 1 20 8v3.25"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m21 22-2.88-2.88"/><circle cx="16" cy="17" r="3"/>`)

// FileSearch is [file-search](https://lucide.dev/icons/file-search).
var FileSearch = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><circle cx="11.5" cy="14.5" r="2.5"/><path d="M13.3 16.3 15 18"/>`)

// FileSignal is [file-signal](https://lucide.dev/icons/file-signal).
var FileSignal = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 15h.01"/><path d="M11.5 13.5a2.5 2.5 0 0 1 0 3"/><path d="M15 12a5 5 0 0 1 0 6"/>`)

// FileSliders is [file-sliders](https://lucide.dev/icons/file-sliders).
var FileSliders = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 12h8"/><path d="M10 11v2"/><path d="M8 17h8"/><path d="M14 16v2"/>`)

// FileSpreadsheet is [file-spreadsheet](https://lucide.dev/icons/file-spreadsheet).
var FileSpreadsheet = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M8 13h2"/><path d="M14 13h2"/><path d="M8 17h2"/><path d="M14 17h2"/>`)

// FileStack is [file-stack](https://lucide.dev/icons/file-stack).
var FileStack = icon.Lucide(viewBox, paint, `<path d="M11 21a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1"/><path d="M16 16a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1"/><path d="M21 6a2 2 0 0 0-.586-1.414l-2-2A2 2 0 0 0 17 2h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1z"/>`)

// FileSymlink is [file-symlink](https://lucide.dev/icons/file-symlink).
var FileSymlink = icon.Lucide(viewBox, paint, `<path d="M4 11V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h7"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m10 18 3-3-3-3"/>`)

// FileTerminal is [file-terminal](https://lucide.dev/icons/file-terminal).
var FileTerminal = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m8 16 2-2-2-2"/><path d="M12 18h4"/>`)

// FileText is [file-text](https://lucide.dev/icons/file-text).
var FileText = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M10 9H8"/><path d="M16 13H8"/><path d="M16 17H8"/>`)

// FileTypeCorner is [file-type-corner](https://lucide.dev/icons/file-type-corner).
var FileTypeCorner = icon.Lucide(viewBox, paint, `<path d="M12 22h6a2 2 0 0 0 2-2V8a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 14 2H6a2 2 0 0 0-2 2v6"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M3 16v-1.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5V16"/><path d="M6 22h2"/><path d="M7 14v8"/>`)

// FileType is [file-type](https://lucide.dev/icons/file-type).
var FileType = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M11 18h2"/><path d="M12 12v6"/><path d="M9 13v-.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 .5.5v.5"/>`)

// FileUp is [file-up](https://lucide.dev/icons/file-up).
var FileUp = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M12 12v6"/><path d="m15 15-3-3-3 3"/>`)

// FileUser is [file-user](https://lucide.dev/icons/file-user).
var FileUser = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M16 22a4 4 0 0 0-8 0"/><circle cx="12" cy="15" r="3"/>`)

// FileVideoCamera is [file-video-camera](https://lucide.dev/icons/file-video-camera).
var FileVideoCamera = icon.Lucide(viewBox, paint, `<path d="M4 12V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m10 17.843 3.033-1.755a.64.64 0 0 1 .967.56v4.704a.65.65 0 0 1-.967.56L10 20.157"/><rect width="7" height="6" x="3" y="16" rx="1"/>`)

// FileVolume is [file-volume](https://lucide.dev/icons/file-volume).
var FileVolume = icon.Lucide(viewBox, paint, `<path d="M4 11.55V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2h-1.95"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M12 15a5 5 0 0 1 0 6"/><path d="M8 14.502a.5.5 0 0 0-.826-.381l-1.893 1.631a1 1 0 0 1-.651.243H3.5a.5.5 0 0 0-.5.501v3.006a.5.5 0 0 0 .5.501h1.129a1 1 0 0 1 .652.243l1.893 1.633a.5.5 0 0 0 .826-.38z"/>`)

// FileXCorner is [file-x-corner](https://lucide.dev/icons/file-x-corner).
var FileXCorner = icon.Lucide(viewBox, paint, `<path d="M11 22H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v5"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m15 17 5 5"/><path d="m20 17-5 5"/>`)

// FileX is [file-x](https://lucide.dev/icons/file-x).
var FileX = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="m14.5 12.5-5 5"/><path d="m9.5 12.5 5 5"/>`)

// File is [file](https://lucide.dev/icons/file).
var File = icon.Lucide(viewBox, paint, `<path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.704.706l3.588 3.588A2.4 2.4 0 0 1 20 8v12a2 2 0 0 1-2 2z"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/>`)

// Files is [files](https://lucide.dev/icons/files).
var Files = icon.Lucide(viewBox, paint, `<path d="M15 2h-4a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V8"/><path d="M16.706 2.706A2.4 2.4 0 0 0 15 2v5a1 1 0 0 0 1 1h5a2.4 2.4 0 0 0-.706-1.706z"/><path d="M5 7a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h8a2 2 0 0 0 1.732-1"/>`)

// Film is [film](https://lucide.dev/icons/film).
var Film = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 3v18"/><path d="M3 7.5h4"/><path d="M3 12h18"/><path d="M3 16.5h4"/><path d="M17 3v18"/><path d="M17 7.5h4"/><path d="M17 16.5h4"/>`)

// FingerprintPattern is [fingerprint-pattern](https://lucide.dev/icons/fingerprint-pattern).
var FingerprintPattern = icon.Lucide(viewBox, paint, `<path d="M12 10a2 2 0 0 0-2 2c0 1.02-.1 2.51-.26 4"/><path d="M14 13.12c0 2.38 0 6.38-1 8.88"/><path d="M17.29 21.02c.12-.6.43-2.3.5-3.02"/><path d="M2 12a10 10 0 0 1 18-6"/><path d="M2 16h.01"/><path d="M21.8 16c.2-2 .131-5.354 0-6"/><path d="M5 19.5C5.5 18 6 15 6 12a6 6 0 0 1 .34-2"/><path d="M8.65 22c.21-.66.45-1.32.57-2"/><path d="M9 6.8a6 6 0 0 1 9 5.2v2"/>`)

// FireExtinguisher is [fire-extinguisher](https://lucide.dev/icons/fire-extinguisher).
var FireExtinguisher = icon.Lucide(viewBox, paint, `<path d="M15 6.5V3a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3.5"/><path d="M9 18h8"/><path d="M18 3h-3"/><path d="M11 3a6 6 0 0 0-6 6v11"/><path d="M5 13h4"/><path d="M17 10a4 4 0 0 0-8 0v10a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2Z"/>`)

// FishOff is [fish-off](https://lucide.dev/icons/fish-off).
var FishOff = icon.Lucide(viewBox, paint, `<path d="M18 12.47v.03m0-.5v.47m-.475 5.056A6.744 6.744 0 0 1 15 18c-3.56 0-7.56-2.53-8.5-6 .348-1.28 1.114-2.433 2.121-3.38m3.444-2.088A8.802 8.802 0 0 1 15 6c3.56 0 6.06 2.54 7 6-.309 1.14-.786 2.177-1.413 3.058"/><path d="M7 10.67C7 8 5.58 5.97 2.73 5.5c-1 1.5-1 5 .23 6.5-1.24 1.5-1.24 5-.23 6.5C5.58 18.03 7 16 7 13.33m7.48-4.372A9.77 9.77 0 0 1 16 6.07m0 11.86a9.77 9.77 0 0 1-1.728-3.618"/><path d="m16.01 17.93-.23 1.4A2 2 0 0 1 13.8 21H9.5a5.96 5.96 0 0 0 1.49-3.98M8.53 3h5.27a2 2 0 0 1 1.98 1.67l.23 1.4M2 2l20 20"/>`)

// FishSymbol is [fish-symbol](https://lucide.dev/icons/fish-symbol).
var FishSymbol = icon.Lucide(viewBox, paint, `<path d="M2 16s9-15 20-4C11 23 2 8 2 8"/>`)

// Fish is [fish](https://lucide.dev/icons/fish).
var Fish = icon.Lucide(viewBox, paint, `<path d="M6.5 12c.94-3.46 4.94-6 8.5-6 3.56 0 6.06 2.54 7 6-.94 3.47-3.44 6-7 6s-7.56-2.53-8.5-6Z"/><path d="M18 12v.5"/><path d="M16 17.93a9.77 9.77 0 0 1 0-11.86"/><path d="M7 10.67C7 8 5.58 5.97 2.73 5.5c-1 1.5-1 5 .23 6.5-1.24 1.5-1.24 5-.23 6.5C5.58 18.03 7 16 7 13.33"/><path d="M10.46 7.26C10.2 5.88 9.17 4.24 8 3h5.8a2 2 0 0 1 1.98 1.67l.23 1.4"/><path d="m16.01 17.93-.23 1.4A2 2 0 0 1 13.8 21H9.5a5.96 5.96 0 0 0 1.49-3.98"/>`)

// FishingHook is [fishing-hook](https://lucide.dev/icons/fishing-hook).
var FishingHook = icon.Lucide(viewBox, paint, `<path d="m17.586 11.414-5.93 5.93a1 1 0 0 1-8-8l3.137-3.137a.707.707 0 0 1 1.207.5V10"/><path d="M20.414 8.586 22 7"/><circle cx="19" cy="10" r="2"/>`)

// FishingRod is [fishing-rod](https://lucide.dev/icons/fishing-rod).
var FishingRod = icon.Lucide(viewBox, paint, `<path d="M4 11h1"/><path d="M8 15a2 2 0 0 1-4 0V3a1 1 0 0 1 1-1h.5C14 2 20 9 20 18v4"/><circle cx="18" cy="18" r="2"/>`)

// FlagOff is [flag-off](https://lucide.dev/icons/flag-off).
var FlagOff = icon.Lucide(viewBox, paint, `<path d="M16 16c-3 0-5-2-8-2a6 6 0 0 0-4 1.528"/><path d="m2 2 20 20"/><path d="M4 22V4"/><path d="M7.656 2H8c3 0 5 2 7.333 2q2 0 3.067-.8A1 1 0 0 1 20 4v10.347"/>`)

// FlagTriangleLeft is [flag-triangle-left](https://lucide.dev/icons/flag-triangle-left).
var FlagTriangleLeft = icon.Lucide(viewBox, paint, `<path d="M18 22V2.8a.8.8 0 0 0-1.17-.71L5.45 7.78a.8.8 0 0 0 0 1.44L18 15.5"/>`)

// FlagTriangleRight is [flag-triangle-right](https://lucide.dev/icons/flag-triangle-right).
var FlagTriangleRight = icon.Lucide(viewBox, paint, `<path d="M6 22V2.8a.8.8 0 0 1 1.17-.71l11.38 5.69a.8.8 0 0 1 0 1.44L6 15.5"/>`)

// Flag is [flag](https://lucide.dev/icons/flag).
var Flag = icon.Lucide(viewBox, paint, `<path d="M4 22V4a1 1 0 0 1 .4-.8A6 6 0 0 1 8 2c3 0 5 2 7.333 2q2 0 3.067-.8A1 1 0 0 1 20 4v10a1 1 0 0 1-.4.8A6 6 0 0 1 16 16c-3 0-5-2-8-2a6 6 0 0 0-4 1.528"/>`)

// FlameKindling is [flame-kindling](https://lucide.dev/icons/flame-kindling).
var FlameKindling = icon.Lucide(viewBox, paint, `<path d="M12 2c1 3 2.5 3.5 3.5 4.5A5 5 0 0 1 17 10a5 5 0 1 1-10 0c0-.3 0-.6.1-.9a2 2 0 1 0 3.3-2C8 4.5 11 2 12 2Z"/><path d="m5 22 14-4"/><path d="m5 18 14 4"/>`)

// Flame is [flame](https://lucide.dev/icons/flame).
var Flame = icon.Lucide(viewBox, paint, `<path d="M12 3q1 4 4 6.5t3 5.5a1 1 0 0 1-14 0 5 5 0 0 1 1-3 1 1 0 0 0 5 0c0-2-1.5-3-1.5-5q0-2 2.5-4"/>`)

// FlashlightOff is [flashlight-off](https://lucide.dev/icons/flashlight-off).
var FlashlightOff = icon.Lucide(viewBox, paint, `<path d="M11.652 6H18"/><path d="M12 13v1"/><path d="M16 16v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-8a4 4 0 0 0-.8-2.4l-.6-.8A3 3 0 0 1 6 7V6"/><path d="m2 2 20 20"/><path d="M7.649 2H17a1 1 0 0 1 1 1v4a3 3 0 0 1-.6 1.8l-.6.8a4 4 0 0 0-.55 1.007"/>`)

// Flashlight is [flashlight](https://lucide.dev/icons/flashlight).
var Flashlight = icon.Lucide(viewBox, paint, `<path d="M12 13v1"/><path d="M17 2a1 1 0 0 1 1 1v4a3 3 0 0 1-.6 1.8l-.6.8A4 4 0 0 0 16 12v8a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2v-8a4 4 0 0 0-.8-2.4l-.6-.8A3 3 0 0 1 6 7V3a1 1 0 0 1 1-1z"/><path d="M6 6h12"/>`)

// FlaskConicalOff is [flask-conical-off](https://lucide.dev/icons/flask-conical-off).
var FlaskConicalOff = icon.Lucide(viewBox, paint, `<path d="M10 2v2.343"/><path d="M14 2v6.343"/><path d="m2 2 20 20"/><path d="M20 20a2 2 0 0 1-2 2H6a2 2 0 0 1-1.755-2.96l5.227-9.563"/><path d="M6.453 15H15"/><path d="M8.5 2h7"/>`)

// FlaskConical is [flask-conical](https://lucide.dev/icons/flask-conical).
var FlaskConical = icon.Lucide(viewBox, paint, `<path d="M14 2v6a2 2 0 0 0 .245.96l5.51 10.08A2 2 0 0 1 18 22H6a2 2 0 0 1-1.755-2.96l5.51-10.08A2 2 0 0 0 10 8V2"/><path d="M6.453 15h11.094"/><path d="M8.5 2h7"/>`)

// FlaskRound is [flask-round](https://lucide.dev/icons/flask-round).
var FlaskRound = icon.Lucide(viewBox, paint, `<path d="M10 2v6.292a7 7 0 1 0 4 0V2"/><path d="M5 15h14"/><path d="M8.5 2h7"/>`)

// Flower2 is [flower-2](https://lucide.dev/icons/flower-2).
var Flower2 = icon.Lucide(viewBox, paint, `<path d="M12 5a3 3 0 1 1 3 3m-3-3a3 3 0 1 0-3 3m3-3v1M9 8a3 3 0 1 0 3 3M9 8h1m5 0a3 3 0 1 1-3 3m3-3h-1m-2 3v-1"/><circle cx="12" cy="8" r="2"/><path d="M12 10v12"/><path d="M12 22c4.2 0 7-1.667 7-5-4.2 0-7 1.667-7 5Z"/><path d="M12 22c-4.2 0-7-1.667-7-5 4.2 0 7 1.667 7 5Z"/>`)

// Flower is [flower](https://lucide.dev/icons/flower).
var Flower = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="3"/><path d="M12 16.5A4.5 4.5 0 1 1 7.5 12 4.5 4.5 0 1 1 12 7.5a4.5 4.5 0 1 1 4.5 4.5 4.5 4.5 0 1 1-4.5 4.5"/><path d="M12 7.5V9"/><path d="M7.5 12H9"/><path d="M16.5 12H15"/><path d="M12 16.5V15"/><path d="m8 8 1.88 1.88"/><path d="M14.12 9.88 16 8"/><path d="m8 16 1.88-1.88"/><path d="M14.12 14.12 16 16"/>`)

// Focus is [focus](https://lucide.dev/icons/focus).
var Focus = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="3"/><path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/>`)

// FoldHorizontal is [fold-horizontal](https://lucide.dev/icons/fold-horizontal).
var FoldHorizontal = icon.Lucide(viewBox, paint, `<path d="M2 12h6"/><path d="M22 12h-6"/><path d="M12 2v2"/><path d="M12 8v2"/><path d="M12 14v2"/><path d="M12 20v2"/><path d="m19 9-3 3 3 3"/><path d="m5 15 3-3-3-3"/>`)

// FoldVertical is [fold-vertical](https://lucide.dev/icons/fold-vertical).
var FoldVertical = icon.Lucide(viewBox, paint, `<path d="M12 22v-6"/><path d="M12 8V2"/><path d="M4 12H2"/><path d="M10 12H8"/><path d="M16 12h-2"/><path d="M22 12h-2"/><path d="m15 19-3-3-3 3"/><path d="m15 5-3 3-3-3"/>`)

// FolderArchive is [folder-archive](https://lucide.dev/icons/folder-archive).
var FolderArchive = icon.Lucide(viewBox, paint, `<circle cx="15" cy="19" r="2"/><path d="M20.9 19.8A2 2 0 0 0 22 18V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h5.1"/><path d="M15 11v-1"/><path d="M15 17v-2"/>`)

// FolderBookmark is [folder-bookmark](https://lucide.dev/icons/folder-bookmark).
var FolderBookmark = icon.Lucide(viewBox, paint, `<path d="M12 6v7.751a.25.25 0 00.407.195l2.28-1.834a.5.5 0 01.627 0l2.28 1.834a.25.25 0 00.406-.195V6"/><path d="M20 20a2 2 0 002-2V8a2 2 0 00-2-2h-7.9a2 2 0 01-1.69-.9L9.6 3.9A2 2 0 007.93 3H4a2 2 0 00-2 2v13a2 2 0 002 2z"/>`)

// FolderCheck is [folder-check](https://lucide.dev/icons/folder-check).
var FolderCheck = icon.Lucide(viewBox, paint, `<path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="m9 13 2 2 4-4"/>`)

// FolderClock is [folder-clock](https://lucide.dev/icons/folder-clock).
var FolderClock = icon.Lucide(viewBox, paint, `<path d="M16 14v2.2l1.6 1"/><path d="M7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2"/><circle cx="16" cy="16" r="6"/>`)

// FolderClosed is [folder-closed](https://lucide.dev/icons/folder-closed).
var FolderClosed = icon.Lucide(viewBox, paint, `<path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="M2 10h20"/>`)

// FolderCode is [folder-code](https://lucide.dev/icons/folder-code).
var FolderCode = icon.Lucide(viewBox, paint, `<path d="M10 10.5 8 13l2 2.5"/><path d="m14 10.5 2 2.5-2 2.5"/><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2z"/>`)

// FolderCog is [folder-cog](https://lucide.dev/icons/folder-cog).
var FolderCog = icon.Lucide(viewBox, paint, `<path d="M10.3 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.98a2 2 0 0 1 1.69.9l.66 1.2A2 2 0 0 0 12 6h8a2 2 0 0 1 2 2v3.3"/><path d="m14.305 19.53.923-.382"/><path d="m15.228 16.852-.923-.383"/><path d="m16.852 15.228-.383-.923"/><path d="m16.852 20.772-.383.924"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.696-.382-.924"/><path d="m20.772 16.852.924-.383"/><path d="m20.772 19.148.924.383"/><circle cx="18" cy="18" r="3"/>`)

// FolderDot is [folder-dot](https://lucide.dev/icons/folder-dot).
var FolderDot = icon.Lucide(viewBox, paint, `<path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="1"/>`)

// FolderDown is [folder-down](https://lucide.dev/icons/folder-down).
var FolderDown = icon.Lucide(viewBox, paint, `<path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="M12 10v6"/><path d="m15 13-3 3-3-3"/>`)

// FolderGit2 is [folder-git-2](https://lucide.dev/icons/folder-git-2).
var FolderGit2 = icon.Lucide(viewBox, paint, `<path d="M18 19a5 5 0 0 1-5-5v8"/><path d="M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v5"/><circle cx="13" cy="12" r="2"/><circle cx="20" cy="19" r="2"/>`)

// FolderGit is [folder-git](https://lucide.dev/icons/folder-git).
var FolderGit = icon.Lucide(viewBox, paint, `<circle cx="12" cy="13" r="2"/><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="M14 13h3"/><path d="M7 13h3"/>`)

// FolderHeart is [folder-heart](https://lucide.dev/icons/folder-heart).
var FolderHeart = icon.Lucide(viewBox, paint, `<path d="M10.638 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v3.417"/><path d="M14.62 18.8A2.25 2.25 0 1 1 18 15.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"/>`)

// FolderInput is [folder-input](https://lucide.dev/icons/folder-input).
var FolderInput = icon.Lucide(viewBox, paint, `<path d="M2 9V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1"/><path d="M2 13h10"/><path d="m9 16 3-3-3-3"/>`)

// FolderKanban is [folder-kanban](https://lucide.dev/icons/folder-kanban).
var FolderKanban = icon.Lucide(viewBox, paint, `<path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><path d="M8 10v4"/><path d="M12 10v2"/><path d="M16 10v6"/>`)

// FolderKey is [folder-key](https://lucide.dev/icons/folder-key).
var FolderKey = icon.Lucide(viewBox, paint, `<path d="M13 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v1.36"/><path d="M19 12v6"/><path d="M19 14h2"/><circle cx="19" cy="20" r="2"/>`)

// FolderLock is [folder-lock](https://lucide.dev/icons/folder-lock).
var FolderLock = icon.Lucide(viewBox, paint, `<rect width="8" height="5" x="14" y="17" rx="1"/><path d="M10 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v2.5"/><path d="M20 17v-2a2 2 0 1 0-4 0v2"/>`)

// FolderMinus is [folder-minus](https://lucide.dev/icons/folder-minus).
var FolderMinus = icon.Lucide(viewBox, paint, `<path d="M9 13h6"/><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/>`)

// FolderOpenDot is [folder-open-dot](https://lucide.dev/icons/folder-open-dot).
var FolderOpenDot = icon.Lucide(viewBox, paint, `<path d="m6 14 1.45-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.55 6a2 2 0 0 1-1.94 1.5H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H18a2 2 0 0 1 2 2v2"/><circle cx="14" cy="15" r="1"/>`)

// FolderOpen is [folder-open](https://lucide.dev/icons/folder-open).
var FolderOpen = icon.Lucide(viewBox, paint, `<path d="m6 14 1.5-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.54 6a2 2 0 0 1-1.95 1.5H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H18a2 2 0 0 1 2 2v2"/>`)

// FolderOutput is [folder-output](https://lucide.dev/icons/folder-output).
var FolderOutput = icon.Lucide(viewBox, paint, `<path d="M2 7.5V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-1.5"/><path d="M2 13h10"/><path d="m5 10-3 3 3 3"/>`)

// FolderPen is [folder-pen](https://lucide.dev/icons/folder-pen).
var FolderPen = icon.Lucide(viewBox, paint, `<path d="M2 11.5V5a2 2 0 0 1 2-2h3.9c.7 0 1.3.3 1.7.9l.8 1.2c.4.6 1 .9 1.7.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-9.5"/><path d="M11.378 13.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/>`)

// FolderPlus is [folder-plus](https://lucide.dev/icons/folder-plus).
var FolderPlus = icon.Lucide(viewBox, paint, `<path d="M12 10v6"/><path d="M9 13h6"/><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/>`)

// FolderRoot is [folder-root](https://lucide.dev/icons/folder-root).
var FolderRoot = icon.Lucide(viewBox, paint, `<path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="2"/><path d="M12 15v5"/>`)

// FolderSearch2 is [folder-search-2](https://lucide.dev/icons/folder-search-2).
var FolderSearch2 = icon.Lucide(viewBox, paint, `<circle cx="11.5" cy="12.5" r="2.5"/><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="M13.3 14.3 15 16"/>`)

// FolderSearch is [folder-search](https://lucide.dev/icons/folder-search).
var FolderSearch = icon.Lucide(viewBox, paint, `<path d="M10.7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v4.1"/><path d="m21 21-1.9-1.9"/><circle cx="17" cy="17" r="3"/>`)

// FolderSymlink is [folder-symlink](https://lucide.dev/icons/folder-symlink).
var FolderSymlink = icon.Lucide(viewBox, paint, `<path d="M2 9.35V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h7"/><path d="m8 16 3-3-3-3"/>`)

// FolderSync is [folder-sync](https://lucide.dev/icons/folder-sync).
var FolderSync = icon.Lucide(viewBox, paint, `<path d="M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v.5"/><path d="M12 10v4h4"/><path d="m12 14 1.535-1.605a5 5 0 0 1 8 1.5"/><path d="M22 22v-4h-4"/><path d="m22 18-1.535 1.605a5 5 0 0 1-8-1.5"/>`)

// FolderTree is [folder-tree](https://lucide.dev/icons/folder-tree).
var FolderTree = icon.Lucide(viewBox, paint, `<path d="M20 10a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-2.5a1 1 0 0 1-.8-.4l-.9-1.2A1 1 0 0 0 15 3h-2a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1Z"/><path d="M20 21a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-2.9a1 1 0 0 1-.88-.55l-.42-.85a1 1 0 0 0-.92-.6H13a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1Z"/><path d="M3 5a2 2 0 0 0 2 2h3"/><path d="M3 3v13a2 2 0 0 0 2 2h3"/>`)

// FolderUp is [folder-up](https://lucide.dev/icons/folder-up).
var FolderUp = icon.Lucide(viewBox, paint, `<path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="M12 10v6"/><path d="m9 13 3-3 3 3"/>`)

// FolderUser is [folder-user](https://lucide.dev/icons/folder-user).
var FolderUser = icon.Lucide(viewBox, paint, `<path d="M10.254 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2V9.255"/><path d="M22 22a4 4 0 0 0-8 0"/><circle cx="18" cy="15" r="3"/>`)

// FolderUsers is [folder-users](https://lucide.dev/icons/folder-users).
var FolderUsers = icon.Lucide(viewBox, paint, `<path d="M5.254 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2V9.255"/><circle cx="13" cy="15" r="3"/><path d="M17 22a4 4 0 0 0-8 0"/><path d="M22 22A4 4 0 0 0 20.24 18.686 1 1 0 0 1 20.1 17.143 3 3 0 0 0 19.5 12.402"/>`)

// FolderX is [folder-x](https://lucide.dev/icons/folder-x).
var FolderX = icon.Lucide(viewBox, paint, `<path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="m9.5 10.5 5 5"/><path d="m14.5 10.5-5 5"/>`)

// Folder is [folder](https://lucide.dev/icons/folder).
var Folder = icon.Lucide(viewBox, paint, `<path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/>`)

// Folders is [folders](https://lucide.dev/icons/folders).
var Folders = icon.Lucide(viewBox, paint, `<path d="M20 5a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h2.5a1.5 1.5 0 0 1 1.2.6l.6.8a1.5 1.5 0 0 0 1.2.6z"/><path d="M3 8.268a2 2 0 0 0-1 1.738V19a2 2 0 0 0 2 2h11a2 2 0 0 0 1.732-1"/>`)

// Footprints is [footprints](https://lucide.dev/icons/footprints).
var Footprints = icon.Lucide(viewBox, paint, `<path d="M4 16v-2.38C4 11.5 2.97 10.5 3 8c.03-2.72 1.49-6 4.5-6C9.37 2 10 3.8 10 5.5c0 3.11-2 5.66-2 8.68V16a2 2 0 1 1-4 0Z"/><path d="M20 20v-2.38c0-2.12 1.03-3.12 1-5.62-.03-2.72-1.49-6-4.5-6C14.63 6 14 7.8 14 9.5c0 3.11 2 5.66 2 8.68V20a2 2 0 1 0 4 0Z"/><path d="M16 17h4"/><path d="M4 13h4"/>`)

// Forklift is [forklift](https://lucide.dev/icons/forklift).
var Forklift = icon.Lucide(viewBox, paint, `<path d="M12 12H5a2 2 0 0 0-2 2v5"/><path d="M15 19h7"/><path d="M16 19V2"/><path d="M6 12V7a2 2 0 0 1 2-2h2.172a2 2 0 0 1 1.414.586l3.828 3.828A2 2 0 0 1 16 10.828"/><path d="M7 19h4"/><circle cx="13" cy="19" r="2"/><circle cx="5" cy="19" r="2"/>`)

// Form is [form](https://lucide.dev/icons/form).
var Form = icon.Lucide(viewBox, paint, `<path d="M4 14h6"/><path d="M4 2h10"/><rect x="4" y="18" width="16" height="4" rx="1"/><rect x="4" y="6" width="16" height="4" rx="1"/>`)

// Forward is [forward](https://lucide.dev/icons/forward).
var Forward = icon.Lucide(viewBox, paint, `<path d="m15 17 5-5-5-5"/><path d="M4 18v-2a4 4 0 0 1 4-4h12"/>`)

// Frame is [frame](https://lucide.dev/icons/frame).
var Frame = icon.Lucide(viewBox, paint, `<line x1="22" x2="2" y1="6" y2="6"/><line x1="22" x2="2" y1="18" y2="18"/><line x1="6" x2="6" y1="2" y2="22"/><line x1="18" x2="18" y1="2" y2="22"/>`)

// Fuel is [fuel](https://lucide.dev/icons/fuel).
var Fuel = icon.Lucide(viewBox, paint, `<path d="M14 13h2a2 2 0 0 1 2 2v2a2 2 0 0 0 4 0v-6.998a2 2 0 0 0-.59-1.42L18 5"/><path d="M14 21V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16"/><path d="M2 21h13"/><path d="M3 9h11"/>`)

// Fullscreen is [fullscreen](https://lucide.dev/icons/fullscreen).
var Fullscreen = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><rect width="10" height="8" x="7" y="8" rx="1"/>`)

// FunnelPlus is [funnel-plus](https://lucide.dev/icons/funnel-plus).
var FunnelPlus = icon.Lucide(viewBox, paint, `<path d="M13.354 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14v6a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341l1.218-1.348"/><path d="M16 6h6"/><path d="M19 3v6"/>`)

// FunnelX is [funnel-x](https://lucide.dev/icons/funnel-x).
var FunnelX = icon.Lucide(viewBox, paint, `<path d="M12.531 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14v6a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341l.427-.473"/><path d="m16.5 3.5 5 5"/><path d="m21.5 3.5-5 5"/>`)

// Funnel is [funnel](https://lucide.dev/icons/funnel).
var Funnel = icon.Lucide(viewBox, paint, `<path d="M10 20a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341L21.74 4.67A1 1 0 0 0 21 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14z"/>`)

// Galaxy is [galaxy](https://lucide.dev/icons/galaxy).
var Galaxy = icon.Lucide(viewBox, paint, `<path d="M16.005 15.108a5.041 6.52 28.25 00-8.008-6.217 5.041 6.52 28.25 008.008 6.217A11.884 7.288-60.76 014.029 7.001"/><path d="M17 21h.01"/><path d="M7 3h.01"/><path d="M7.997 8.891a11.885 7.288-60.756 0111.977 8.107"/><circle cx="12" cy="12" r="1" fill="currentColor"/>`)

// GalleryHorizontalEnd is [gallery-horizontal-end](https://lucide.dev/icons/gallery-horizontal-end).
var GalleryHorizontalEnd = icon.Lucide(viewBox, paint, `<path d="M2 7v10"/><path d="M6 5v14"/><rect width="12" height="18" x="10" y="3" rx="2"/>`)

// GalleryHorizontal is [gallery-horizontal](https://lucide.dev/icons/gallery-horizontal).
var GalleryHorizontal = icon.Lucide(viewBox, paint, `<path d="M2 3v18"/><rect width="12" height="18" x="6" y="3" rx="2"/><path d="M22 3v18"/>`)

// GalleryThumbnails is [gallery-thumbnails](https://lucide.dev/icons/gallery-thumbnails).
var GalleryThumbnails = icon.Lucide(viewBox, paint, `<rect width="18" height="14" x="3" y="3" rx="2"/><path d="M4 21h1"/><path d="M9 21h1"/><path d="M14 21h1"/><path d="M19 21h1"/>`)

// GalleryVerticalEnd is [gallery-vertical-end](https://lucide.dev/icons/gallery-vertical-end).
var GalleryVerticalEnd = icon.Lucide(viewBox, paint, `<path d="M7 2h10"/><path d="M5 6h14"/><rect width="18" height="12" x="3" y="10" rx="2"/>`)

// GalleryVertical is [gallery-vertical](https://lucide.dev/icons/gallery-vertical).
var GalleryVertical = icon.Lucide(viewBox, paint, `<path d="M3 2h18"/><rect width="18" height="12" x="3" y="6" rx="2"/><path d="M3 22h18"/>`)

// Gamepad2 is [gamepad-2](https://lucide.dev/icons/gamepad-2).
var Gamepad2 = icon.Lucide(viewBox, paint, `<line x1="6" x2="10" y1="11" y2="11"/><line x1="8" x2="8" y1="9" y2="13"/><line x1="15" x2="15.01" y1="12" y2="12"/><line x1="18" x2="18.01" y1="10" y2="10"/><path d="M17.32 5H6.68a4 4 0 0 0-3.978 3.59c-.006.052-.01.101-.017.152C2.604 9.416 2 14.456 2 16a3 3 0 0 0 3 3c1 0 1.5-.5 2-1l1.414-1.414A2 2 0 0 1 9.828 16h4.344a2 2 0 0 1 1.414.586L17 18c.5.5 1 1 2 1a3 3 0 0 0 3-3c0-1.545-.604-6.584-.685-7.258-.007-.05-.011-.1-.017-.151A4 4 0 0 0 17.32 5z"/>`)

// GamepadDirectional is [gamepad-directional](https://lucide.dev/icons/gamepad-directional).
var GamepadDirectional = icon.Lucide(viewBox, paint, `<path d="M11.146 15.854a1.207 1.207 0 0 1 1.708 0l1.56 1.56A2 2 0 0 1 15 18.828V21a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2.172a2 2 0 0 1 .586-1.414z"/><path d="M18.828 15a2 2 0 0 1-1.414-.586l-1.56-1.56a1.207 1.207 0 0 1 0-1.708l1.56-1.56A2 2 0 0 1 18.828 9H21a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1z"/><path d="M6.586 14.414A2 2 0 0 1 5.172 15H3a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h2.172a2 2 0 0 1 1.414.586l1.56 1.56a1.207 1.207 0 0 1 0 1.708z"/><path d="M9 3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2.172a2 2 0 0 1-.586 1.414l-1.56 1.56a1.207 1.207 0 0 1-1.708 0l-1.56-1.56A2 2 0 0 1 9 5.172z"/>`)

// Gamepad is [gamepad](https://lucide.dev/icons/gamepad).
var Gamepad = icon.Lucide(viewBox, paint, `<line x1="6" x2="10" y1="12" y2="12"/><line x1="8" x2="8" y1="10" y2="14"/><line x1="15" x2="15.01" y1="13" y2="13"/><line x1="18" x2="18.01" y1="11" y2="11"/><rect width="20" height="12" x="2" y="6" rx="2"/>`)

// GapHorizontal is [gap-horizontal](https://lucide.dev/icons/gap-horizontal).
var GapHorizontal = icon.Lucide(viewBox, paint, `<path d="M12 2v2"/><path d="M12 8v2"/><path d="M12 14v2"/><path d="M12 20v2"/><path d="M21 3h-3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3"/><path d="M3 3h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H3"/>`)

// GapVertical is [gap-vertical](https://lucide.dev/icons/gap-vertical).
var GapVertical = icon.Lucide(viewBox, paint, `<path d="M2 12h2"/><path d="M8 12h2"/><path d="M14 12h2"/><path d="M20 12h2"/><path d="M3 21v-3a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3"/><path d="M3 3v3a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3"/>`)

// Gauge is [gauge](https://lucide.dev/icons/gauge).
var Gauge = icon.Lucide(viewBox, paint, `<path d="m12 14 4-4"/><path d="M3.34 19a10 10 0 1 1 17.32 0"/>`)

// Gavel is [gavel](https://lucide.dev/icons/gavel).
var Gavel = icon.Lucide(viewBox, paint, `<path d="m14 13-8.381 8.38a1 1 0 0 1-3.001-3l8.384-8.381"/><path d="m16 16 6-6"/><path d="m21.5 10.5-8-8"/><path d="m8 8 6-6"/><path d="m8.5 7.5 8 8"/>`)

// Gem is [gem](https://lucide.dev/icons/gem).
var Gem = icon.Lucide(viewBox, paint, `<path d="M10.5 3 8 9l4 13 4-13-2.5-6"/><path d="M17 3a2 2 0 0 1 1.6.8l3 4a2 2 0 0 1 .013 2.382l-7.99 10.986a2 2 0 0 1-3.247 0l-7.99-10.986A2 2 0 0 1 2.4 7.8l2.998-3.997A2 2 0 0 1 7 3z"/><path d="M2 9h20"/>`)

// GeorgianLari is [georgian-lari](https://lucide.dev/icons/georgian-lari).
var GeorgianLari = icon.Lucide(viewBox, paint, `<path d="M11.5 21a7.5 7.5 0 1 1 7.35-9"/><path d="M13 12V3"/><path d="M4 21h16"/><path d="M9 12V3"/>`)

// GermOff is [germ-off](https://lucide.dev/icons/germ-off).
var GermOff = icon.Lucide(viewBox, paint, `<path d="m11 2 .925 1.848"/><path d="M13 15h.01"/><path d="M13.424 7.768a2 2 0 112.808 2.808"/><path d="m16 21-1-2.472"/><path d="M16.988 16.988A12 12 0 019 20a5 5 0 01-2.759-9.171 8.8 8.8 0 002.307-2.28"/><path d="m19 2-1 1.804"/><path d="m2 19 2.746-1.373"/><path d="m2 2 20 20"/><path d="m22 16-2.474-2.13a12 12 0 001.376-3.786 6 6 0 00-10.313-5.151"/><path d="m22 5-1.804 1"/><path d="m3 10 2 2"/><path d="M9 16h.01"/><path d="M9 20v2"/>`)

// Germ is [germ](https://lucide.dev/icons/germ).
var Germ = icon.Lucide(viewBox, paint, `<path d="m11 2 .925 1.848"/><path d="M13 15h.01"/><path d="m16 21-1-2.472"/><path d="m19 2-1 1.804"/><path d="m2 19 2.746-1.373"/><path d="m22 16-2.474-2.13"/><path d="m22 5-1.804 1"/><path d="m3 10 2 2"/><path d="M9 16h.01"/><path d="M9 20v2"/><path d="M9.33 7.035c-.51 1.478-1.786 2.93-3.09 3.794A5 5 0 009 20a12.1 12.1 0 0011.902-9.916A6 6 0 009.33 7.035"/><circle cx="15" cy="9" r="2"/>`)

// Ghost is [ghost](https://lucide.dev/icons/ghost).
var Ghost = icon.Lucide(viewBox, paint, `<path d="M15 10v1"/><path d="M7.528 20.472a1.6 1.6 0 012.277 0l1.057 1.056a1.6 1.6 0 002.276 0l1.057-1.056a1.6 1.6 0 012.277 0l1.114 1.114a1.4 1.4 0 002.414-1V10a8 8 0 00-16 0v10.586a1.4 1.4 0 002.414 1z"/><path d="M9 10v1"/>`)

// Gift is [gift](https://lucide.dev/icons/gift).
var Gift = icon.Lucide(viewBox, paint, `<path d="M12 7v14"/><path d="M20 11v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"/><path d="M7.5 7a1 1 0 0 1 0-5A4.8 8 0 0 1 12 7a4.8 8 0 0 1 4.5-5 1 1 0 0 1 0 5"/><rect x="3" y="7" width="18" height="4" rx="1"/>`)

// GitBranchMinus is [git-branch-minus](https://lucide.dev/icons/git-branch-minus).
var GitBranchMinus = icon.Lucide(viewBox, paint, `<path d="M15 6a9 9 0 0 0-9 9V3"/><path d="M21 18h-6"/><circle cx="18" cy="6" r="3"/><circle cx="6" cy="18" r="3"/>`)

// GitBranchPlus is [git-branch-plus](https://lucide.dev/icons/git-branch-plus).
var GitBranchPlus = icon.Lucide(viewBox, paint, `<path d="M6 3v12"/><path d="M18 9a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/><path d="M6 21a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"/><path d="M15 6a9 9 0 0 0-9 9"/><path d="M18 15v6"/><path d="M21 18h-6"/>`)

// GitBranch is [git-branch](https://lucide.dev/icons/git-branch).
var GitBranch = icon.Lucide(viewBox, paint, `<path d="M15 6a9 9 0 0 0-9 9V3"/><circle cx="18" cy="6" r="3"/><circle cx="6" cy="18" r="3"/>`)

// GitCommitHorizontal is [git-commit-horizontal](https://lucide.dev/icons/git-commit-horizontal).
var GitCommitHorizontal = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="3"/><line x1="3" x2="9" y1="12" y2="12"/><line x1="15" x2="21" y1="12" y2="12"/>`)

// GitCommitVertical is [git-commit-vertical](https://lucide.dev/icons/git-commit-vertical).
var GitCommitVertical = icon.Lucide(viewBox, paint, `<path d="M12 3v6"/><circle cx="12" cy="12" r="3"/><path d="M12 15v6"/>`)

// GitCompareArrows is [git-compare-arrows](https://lucide.dev/icons/git-compare-arrows).
var GitCompareArrows = icon.Lucide(viewBox, paint, `<circle cx="5" cy="6" r="3"/><path d="M12 6h5a2 2 0 0 1 2 2v7"/><path d="m15 9-3-3 3-3"/><circle cx="19" cy="18" r="3"/><path d="M12 18H7a2 2 0 0 1-2-2V9"/><path d="m9 15 3 3-3 3"/>`)

// GitCompare is [git-compare](https://lucide.dev/icons/git-compare).
var GitCompare = icon.Lucide(viewBox, paint, `<circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M13 6h3a2 2 0 0 1 2 2v7"/><path d="M11 18H8a2 2 0 0 1-2-2V9"/>`)

// GitFork is [git-fork](https://lucide.dev/icons/git-fork).
var GitFork = icon.Lucide(viewBox, paint, `<circle cx="12" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><circle cx="18" cy="6" r="3"/><path d="M18 9v2c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V9"/><path d="M12 12v3"/>`)

// GitGraph is [git-graph](https://lucide.dev/icons/git-graph).
var GitGraph = icon.Lucide(viewBox, paint, `<circle cx="5" cy="6" r="3"/><path d="M5 9v6"/><circle cx="5" cy="18" r="3"/><path d="M12 3v18"/><circle cx="19" cy="6" r="3"/><path d="M16 15.7A9 9 0 0 0 19 9"/>`)

// GitMergeConflict is [git-merge-conflict](https://lucide.dev/icons/git-merge-conflict).
var GitMergeConflict = icon.Lucide(viewBox, paint, `<path d="M12 6h4a2 2 0 0 1 2 2v7"/><path d="M6 12v9"/><path d="m8.5 3.5-5 5"/><path d="m8.5 8.5-5-5"/><circle cx="18" cy="18" r="3"/>`)

// GitMerge is [git-merge](https://lucide.dev/icons/git-merge).
var GitMerge = icon.Lucide(viewBox, paint, `<circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M6 21V9a9 9 0 0 0 9 9"/>`)

// GitPullRequestArrow is [git-pull-request-arrow](https://lucide.dev/icons/git-pull-request-arrow).
var GitPullRequestArrow = icon.Lucide(viewBox, paint, `<circle cx="5" cy="6" r="3"/><path d="M5 9v12"/><circle cx="19" cy="18" r="3"/><path d="m15 9-3-3 3-3"/><path d="M12 6h5a2 2 0 0 1 2 2v7"/>`)

// GitPullRequestClosed is [git-pull-request-closed](https://lucide.dev/icons/git-pull-request-closed).
var GitPullRequestClosed = icon.Lucide(viewBox, paint, `<path d="m15.5 3.5 5 5"/><path d="m15.5 8.5 5-5"/><path d="M18 11.62V15"/><path d="M6 9v12"/><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/>`)

// GitPullRequestCreateArrow is [git-pull-request-create-arrow](https://lucide.dev/icons/git-pull-request-create-arrow).
var GitPullRequestCreateArrow = icon.Lucide(viewBox, paint, `<circle cx="5" cy="6" r="3"/><path d="M5 9v12"/><path d="m15 9-3-3 3-3"/><path d="M12 6h5a2 2 0 0 1 2 2v3"/><path d="M19 15v6"/><path d="M22 18h-6"/>`)

// GitPullRequestCreate is [git-pull-request-create](https://lucide.dev/icons/git-pull-request-create).
var GitPullRequestCreate = icon.Lucide(viewBox, paint, `<circle cx="6" cy="6" r="3"/><path d="M6 9v12"/><path d="M13 6h3a2 2 0 0 1 2 2v3"/><path d="M18 15v6"/><path d="M21 18h-6"/>`)

// GitPullRequestDraft is [git-pull-request-draft](https://lucide.dev/icons/git-pull-request-draft).
var GitPullRequestDraft = icon.Lucide(viewBox, paint, `<circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M18 6V5"/><path d="M18 11v-1"/><line x1="6" x2="6" y1="9" y2="21"/>`)

// GitPullRequest is [git-pull-request](https://lucide.dev/icons/git-pull-request).
var GitPullRequest = icon.Lucide(viewBox, paint, `<circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M13 6h3a2 2 0 0 1 2 2v7"/><line x1="6" x2="6" y1="9" y2="21"/>`)

// GlassWater is [glass-water](https://lucide.dev/icons/glass-water).
var GlassWater = icon.Lucide(viewBox, paint, `<path d="M5.116 4.104A1 1 0 0 1 6.11 3h11.78a1 1 0 0 1 .994 1.105L17.19 20.21A2 2 0 0 1 15.2 22H8.8a2 2 0 0 1-2-1.79z"/><path d="M6 12a5 5 0 0 1 6 0 5 5 0 0 0 6 0"/>`)

// Glasses is [glasses](https://lucide.dev/icons/glasses).
var Glasses = icon.Lucide(viewBox, paint, `<circle cx="6" cy="15" r="4"/><circle cx="18" cy="15" r="4"/><path d="M14 15a2 2 0 0 0-2-2 2 2 0 0 0-2 2"/><path d="M2.5 13 5 7c.7-1.3 1.4-2 3-2"/><path d="M21.5 13 19 7c-.7-1.3-1.5-2-3-2"/>`)

// GlobeCheck is [globe-check](https://lucide.dev/icons/globe-check).
var GlobeCheck = icon.Lucide(viewBox, paint, `<path d="m15 6 2 2 4-4"/><path d="M2 12h20A10 10 0 1 1 12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 4-10"/>`)

// GlobeCode is [globe-code](https://lucide.dev/icons/globe-code).
var GlobeCode = icon.Lucide(viewBox, paint, `<path d="M15.5 10 13 7.5 15.5 5"/><path d="M15.861 14A14.5 14.5 0 0112 22a14.48 14.48 0 010-20 10 10 0 109.888 11.5"/><path d="M19.5 5 22 7.5 19.5 10"/><path d="M2 12h8.5"/>`)

// GlobeLock is [globe-lock](https://lucide.dev/icons/globe-lock).
var GlobeLock = icon.Lucide(viewBox, paint, `<path d="M15.686 15A14.5 14.5 0 0 1 12 22a14.5 14.5 0 0 1 0-20 10 10 0 1 0 9.542 13"/><path d="M2 12h8.5"/><path d="M20 6V4a2 2 0 1 0-4 0v2"/><rect width="8" height="5" x="14" y="6" rx="1"/>`)

// GlobeOff is [globe-off](https://lucide.dev/icons/globe-off).
var GlobeOff = icon.Lucide(viewBox, paint, `<path d="M10.114 4.462A14.5 14.5 0 0 1 12 2a10 10 0 0 1 9.313 13.643"/><path d="M15.557 15.556A14.5 14.5 0 0 1 12 22 10 10 0 0 1 4.929 4.929"/><path d="M15.892 10.234A14.5 14.5 0 0 0 12 2a10 10 0 0 0-3.643.687"/><path d="M17.656 12H22"/><path d="M19.071 19.071A10 10 0 0 1 12 22 14.5 14.5 0 0 1 8.44 8.45"/><path d="M2 12h10"/><path d="m2 2 20 20"/>`)

// GlobeX is [globe-x](https://lucide.dev/icons/globe-x).
var GlobeX = icon.Lucide(viewBox, paint, `<path d="m16 3 5 5"/><path d="M2 12h20A10 10 0 1 1 12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 4-10"/><path d="m21 3-5 5"/>`)

// Globe is [globe](https://lucide.dev/icons/globe).
var Globe = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20"/><path d="M2 12h20"/>`)

// Goal is [goal](https://lucide.dev/icons/goal).
var Goal = icon.Lucide(viewBox, paint, `<path d="M12 13V2l8 4-8 4"/><path d="M20.561 10.222a9 9 0 1 1-12.55-5.29"/><path d="M8.002 9.997a5 5 0 1 0 8.9 2.02"/>`)

// Gpu is [gpu](https://lucide.dev/icons/gpu).
var Gpu = icon.Lucide(viewBox, paint, `<path d="M2 17h18a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H2"/><path d="M2 21V3"/><path d="M7 17v3a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-3"/><circle cx="16" cy="11" r="2"/><circle cx="8" cy="11" r="2"/>`)

// GraduationCap is [graduation-cap](https://lucide.dev/icons/graduation-cap).
var GraduationCap = icon.Lucide(viewBox, paint, `<path d="M21.42 10.922a1 1 0 0 0-.019-1.838L12.83 5.18a2 2 0 0 0-1.66 0L2.6 9.08a1 1 0 0 0 0 1.832l8.57 3.908a2 2 0 0 0 1.66 0z"/><path d="M22 10v6"/><path d="M6 12.5V16a6 3 0 0 0 12 0v-3.5"/>`)

// Grape is [grape](https://lucide.dev/icons/grape).
var Grape = icon.Lucide(viewBox, paint, `<path d="M22 5V2l-5.89 5.89"/><circle cx="16.6" cy="15.89" r="3"/><circle cx="8.11" cy="7.4" r="3"/><circle cx="12.35" cy="11.65" r="3"/><circle cx="13.91" cy="5.85" r="3"/><circle cx="18.15" cy="10.09" r="3"/><circle cx="6.56" cy="13.2" r="3"/><circle cx="10.8" cy="17.44" r="3"/><circle cx="5" cy="19" r="3"/>`)

// GraphHub is [graph-hub](https://lucide.dev/icons/graph-hub).
var GraphHub = icon.Lucide(viewBox, paint, `<circle cx="12" cy="7" r="4"/><circle cx="4" cy="19" r="2"/><circle cx="12" cy="19" r="2"/><circle cx="20" cy="19" r="2"/><path d="M8 7a4 4 0 0 0-4 4v6"/><path d="M12 11v6"/><path d="M16 7a4 4 0 0 1 4 4v6"/>`)

// GraphStar is [graph-star](https://lucide.dev/icons/graph-star).
var GraphStar = icon.Lucide(viewBox, paint, `<path d="m12.342 9.316 1.764-3.527"/><path d="m12.59 14.545 2.35 3.759"/><path d="m13.982 12.332 4.03.447"/><path d="M5.414 17.586 8.88 14.12"/><path d="M9.025 9.742 5.317 5.505"/><circle cx="11" cy="12" r="3"/><circle cx="15" cy="4" r="2"/><circle cx="16" cy="20" r="2"/><circle cx="20" cy="13" r="2"/><circle cx="4" cy="19" r="2"/><circle cx="4" cy="4" r="2"/>`)

// Grid2x2Check is [grid-2x2-check](https://lucide.dev/icons/grid-2x2-check).
var Grid2x2Check = icon.Lucide(viewBox, paint, `<path d="M12 3v17a1 1 0 0 1-1 1H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6a1 1 0 0 1-1 1H3"/><path d="m16 19 2 2 4-4"/>`)

// Grid2x2Plus is [grid-2x2-plus](https://lucide.dev/icons/grid-2x2-plus).
var Grid2x2Plus = icon.Lucide(viewBox, paint, `<path d="M12 3v17a1 1 0 0 1-1 1H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6a1 1 0 0 1-1 1H3"/><path d="M16 19h6"/><path d="M19 22v-6"/>`)

// Grid2x2X is [grid-2x2-x](https://lucide.dev/icons/grid-2x2-x).
var Grid2x2X = icon.Lucide(viewBox, paint, `<path d="M12 3v17a1 1 0 0 1-1 1H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v6a1 1 0 0 1-1 1H3"/><path d="m16.5 16.5 5 5"/><path d="m16.5 21.5 5-5"/>`)

// Grid2x2 is [grid-2x2](https://lucide.dev/icons/grid-2x2).
var Grid2x2 = icon.Lucide(viewBox, paint, `<path d="M12 3v18"/><path d="M3 12h18"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// Grid3x2 is [grid-3x2](https://lucide.dev/icons/grid-3x2).
var Grid3x2 = icon.Lucide(viewBox, paint, `<path d="M15 3v18"/><path d="M3 12h18"/><path d="M9 3v18"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// Grid3x3 is [grid-3x3](https://lucide.dev/icons/grid-3x3).
var Grid3x3 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18"/><path d="M3 15h18"/><path d="M9 3v18"/><path d="M15 3v18"/>`)

// GripHorizontal is [grip-horizontal](https://lucide.dev/icons/grip-horizontal).
var GripHorizontal = icon.Lucide(viewBox, paint, `<circle cx="12" cy="9" r="1"/><circle cx="19" cy="9" r="1"/><circle cx="5" cy="9" r="1"/><circle cx="12" cy="15" r="1"/><circle cx="19" cy="15" r="1"/><circle cx="5" cy="15" r="1"/>`)

// GripVertical is [grip-vertical](https://lucide.dev/icons/grip-vertical).
var GripVertical = icon.Lucide(viewBox, paint, `<circle cx="9" cy="12" r="1"/><circle cx="9" cy="5" r="1"/><circle cx="9" cy="19" r="1"/><circle cx="15" cy="12" r="1"/><circle cx="15" cy="5" r="1"/><circle cx="15" cy="19" r="1"/>`)

// Grip is [grip](https://lucide.dev/icons/grip).
var Grip = icon.Lucide(viewBox, paint, `<circle cx="12" cy="5" r="1"/><circle cx="19" cy="5" r="1"/><circle cx="5" cy="5" r="1"/><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/><circle cx="12" cy="19" r="1"/><circle cx="19" cy="19" r="1"/><circle cx="5" cy="19" r="1"/>`)

// Groceries is [groceries](https://lucide.dev/icons/groceries).
var Groceries = icon.Lucide(viewBox, paint, `<path d="M10 15h4"/><path d="M19.424 11.095a2.5 2.5 0 00.471-2.309 2 2 0 00.424-3.286 2 2 0 00-2.143-3.322 2 2 0 00-.676.502 2 2 0 00-3.287.424 2.5 2.5 0 00-3.189 2.06A3 3 0 0012 11l4-4"/><path d="M4 12.006A1 1 0 014.994 11H19a1 1 0 011 1v7a2 2 0 01-2 2H6a2 2 0 01-2-2z"/><path d="M7 11a4 4 0 01-4-4V5a1 1 0 011-1h2a4 4 0 013.584 2.222"/>`)

// Group is [group](https://lucide.dev/icons/group).
var Group = icon.Lucide(viewBox, paint, `<path d="M3 7V5c0-1.1.9-2 2-2h2"/><path d="M17 3h2c1.1 0 2 .9 2 2v2"/><path d="M21 17v2c0 1.1-.9 2-2 2h-2"/><path d="M7 21H5c-1.1 0-2-.9-2-2v-2"/><rect width="7" height="5" x="7" y="7" rx="1"/><rect width="7" height="5" x="10" y="12" rx="1"/>`)

// GuitarElectric is [guitar-electric](https://lucide.dev/icons/guitar-electric).
var GuitarElectric = icon.Lucide(viewBox, paint, `<path d="m12.172 11.828 4.242-4.242"/><path d="M16.414 7.586A2 2 0 0017 6.172V5a2 2 0 011.177-1.823L20.59 2.09a1 1 0 011.32 1.32l-1.088 2.412A2 2 0 0119 7h-1a2 2 0 00-1.586.586"/><path d="m6 16 2 2"/><path d="M8.23 9.85a3 3 0 011.28-1.454 1 1 0 011.496.815 4 4 0 003.783 3.783 1 1 0 01.815 1.496 3 3 0 01-1.454 1.28l-.92.38A2 2 0 0012 18a4 4 0 01-4 4 6 6 0 01-6-6 4 4 0 014-4 2 2 0 001.85-1.23z"/>`)

// Guitar is [guitar](https://lucide.dev/icons/guitar).
var Guitar = icon.Lucide(viewBox, paint, `<path d="m11.9 12.1 4.514-4.514"/><path d="M20.1 2.3a1 1 0 0 0-1.4 0l-1.114 1.114A2 2 0 0 0 17 4.828v1.344a2 2 0 0 1-.586 1.414A2 2 0 0 1 17.828 7h1.344a2 2 0 0 0 1.414-.586L21.7 5.3a1 1 0 0 0 0-1.4z"/><path d="m6 16 2 2"/><path d="M8.23 9.85A3 3 0 0 1 11 8a5 5 0 0 1 5 5 3 3 0 0 1-1.85 2.77l-.92.38A2 2 0 0 0 12 18a4 4 0 0 1-4 4 6 6 0 0 1-6-6 4 4 0 0 1 4-4 2 2 0 0 0 1.85-1.23z"/>`)

// HalfNote is [half-note](https://lucide.dev/icons/half-note).
var HalfNote = icon.Lucide(viewBox, paint, `<circle cx="12" cy="18" r="4"/><path d="M16 18V2"/>`)

// Ham is [ham](https://lucide.dev/icons/ham).
var Ham = icon.Lucide(viewBox, paint, `<path d="M13.144 21.144A7.274 10.445 45 1 0 2.856 10.856"/><path d="M13.144 21.144A7.274 4.365 45 0 0 2.856 10.856a7.274 4.365 45 0 0 10.288 10.288"/><path d="M16.565 10.435 18.6 8.4a2.501 2.501 0 1 0 1.65-4.65 2.5 2.5 0 1 0-4.66 1.66l-2.024 2.025"/><path d="m8.5 16.5-1-1"/>`)

// Hamburger is [hamburger](https://lucide.dev/icons/hamburger).
var Hamburger = icon.Lucide(viewBox, paint, `<path d="M12 16H4a2 2 0 1 1 0-4h16a2 2 0 1 1 0 4h-4.25"/><path d="M5 12a2 2 0 0 1-2-2 9 7 0 0 1 18 0 2 2 0 0 1-2 2"/><path d="M5 16a2 2 0 0 0-2 2 3 3 0 0 0 3 3h12a3 3 0 0 0 3-3 2 2 0 0 0-2-2q0 0 0 0"/><path d="m6.67 12 6.13 4.6a2 2 0 0 0 2.8-.4l3.15-4.2"/>`)

// Hammer is [hammer](https://lucide.dev/icons/hammer).
var Hammer = icon.Lucide(viewBox, paint, `<path d="m15 12-9.373 9.373a1 1 0 0 1-3.001-3L12 9"/><path d="m18 15 4-4"/><path d="m21.5 11.5-1.914-1.914A2 2 0 0 1 19 8.172v-.344a2 2 0 0 0-.586-1.414l-1.657-1.657A6 6 0 0 0 12.516 3H9l1.243 1.243A6 6 0 0 1 12 8.485V10l2 2h1.172a2 2 0 0 1 1.414.586L18.5 14.5"/>`)

// HandCoins is [hand-coins](https://lucide.dev/icons/hand-coins).
var HandCoins = icon.Lucide(viewBox, paint, `<path d="M11 15h2a2 2 0 1 0 0-4h-3c-.6 0-1.1.2-1.4.6L3 17"/><path d="m7 21 1.6-1.4c.3-.4.8-.6 1.4-.6h4c1.1 0 2.1-.4 2.8-1.2l4.6-4.4a2 2 0 0 0-2.75-2.91l-4.2 3.9"/><path d="m2 16 6 6"/><circle cx="16" cy="9" r="2.9"/><circle cx="6" cy="5" r="3"/>`)

// HandFist is [hand-fist](https://lucide.dev/icons/hand-fist).
var HandFist = icon.Lucide(viewBox, paint, `<path d="M12.035 17.012a3 3 0 0 0-3-3l-.311-.002a.72.72 0 0 1-.505-1.229l1.195-1.195A2 2 0 0 1 10.828 11H12a2 2 0 0 0 0-4H9.243a3 3 0 0 0-2.122.879l-2.707 2.707A4.83 4.83 0 0 0 3 14a8 8 0 0 0 8 8h2a8 8 0 0 0 8-8V7a2 2 0 1 0-4 0v2a2 2 0 1 0 4 0"/><path d="M13.888 9.662A2 2 0 0 0 17 8V5A2 2 0 1 0 13 5"/><path d="M9 5A2 2 0 1 0 5 5V10"/><path d="M9 7V4A2 2 0 1 1 13 4V7.268"/>`)

// HandGrab is [hand-grab](https://lucide.dev/icons/hand-grab).
var HandGrab = icon.Lucide(viewBox, paint, `<path d="M18 11.5V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v1.4"/><path d="M14 10V8a2 2 0 0 0-2-2a2 2 0 0 0-2 2v2"/><path d="M10 9.9V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v5"/><path d="M6 14a2 2 0 0 0-2-2a2 2 0 0 0-2 2"/><path d="M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-4a8 8 0 0 1-8-8 2 2 0 1 1 4 0"/>`)

// HandHeart is [hand-heart](https://lucide.dev/icons/hand-heart).
var HandHeart = icon.Lucide(viewBox, paint, `<path d="M11 14h2a2 2 0 0 0 0-4h-3c-.6 0-1.1.2-1.4.6L3 16"/><path d="m14.45 13.39 5.05-4.694C20.196 8 21 6.85 21 5.75a2.75 2.75 0 0 0-4.797-1.837.276.276 0 0 1-.406 0A2.75 2.75 0 0 0 11 5.75c0 1.2.802 2.248 1.5 2.946L16 11.95"/><path d="m2 15 6 6"/><path d="m7 20 1.6-1.4c.3-.4.8-.6 1.4-.6h4c1.1 0 2.1-.4 2.8-1.2l4.6-4.4a1 1 0 0 0-2.75-2.91"/>`)

// HandHelping is [hand-helping](https://lucide.dev/icons/hand-helping).
var HandHelping = icon.Lucide(viewBox, paint, `<path d="M11 12h2a2 2 0 1 0 0-4h-3c-.6 0-1.1.2-1.4.6L3 14"/><path d="m7 18 1.6-1.4c.3-.4.8-.6 1.4-.6h4c1.1 0 2.1-.4 2.8-1.2l4.6-4.4a2 2 0 0 0-2.75-2.91l-4.2 3.9"/><path d="m2 13 6 6"/>`)

// HandMetal is [hand-metal](https://lucide.dev/icons/hand-metal).
var HandMetal = icon.Lucide(viewBox, paint, `<path d="M18 12.5V10a2 2 0 0 0-2-2a2 2 0 0 0-2 2v1.4"/><path d="M14 11V9a2 2 0 1 0-4 0v2"/><path d="M10 10.5V5a2 2 0 1 0-4 0v9"/><path d="m7 15-1.76-1.76a2 2 0 0 0-2.83 2.82l3.6 3.6C7.5 21.14 9.2 22 12 22h2a8 8 0 0 0 8-8V7a2 2 0 1 0-4 0v5"/>`)

// HandPlatter is [hand-platter](https://lucide.dev/icons/hand-platter).
var HandPlatter = icon.Lucide(viewBox, paint, `<path d="M12 3V2"/><path d="m15.4 17.4 3.2-2.8a2 2 0 1 1 2.8 2.9l-3.6 3.3c-.7.8-1.7 1.2-2.8 1.2h-4c-1.1 0-2.1-.4-2.8-1.2l-1.302-1.464A1 1 0 0 0 6.151 19H5"/><path d="M2 14h12a2 2 0 0 1 0 4h-2"/><path d="M4 10h16"/><path d="M5 10a7 7 0 0 1 14 0"/><path d="M5 14v6a1 1 0 0 1-1 1H2"/>`)

// Hand is [hand](https://lucide.dev/icons/hand).
var Hand = icon.Lucide(viewBox, paint, `<path d="M18 11V6a2 2 0 0 0-2-2a2 2 0 0 0-2 2"/><path d="M14 10V4a2 2 0 0 0-2-2a2 2 0 0 0-2 2v2"/><path d="M10 10.5V6a2 2 0 0 0-2-2a2 2 0 0 0-2 2v8"/><path d="M18 8a2 2 0 1 1 4 0v6a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15"/>`)

// Handbag is [handbag](https://lucide.dev/icons/handbag).
var Handbag = icon.Lucide(viewBox, paint, `<path d="M2.048 18.566A2 2 0 0 0 4 21h16a2 2 0 0 0 1.952-2.434l-2-9A2 2 0 0 0 18 8H6a2 2 0 0 0-1.952 1.566z"/><path d="M8 11V6a4 4 0 0 1 8 0v5"/>`)

// HandleBottomRight is [handle-bottom-right](https://lucide.dev/icons/handle-bottom-right).
var HandleBottomRight = icon.Lucide(viewBox, paint, `<path d="M21 8 8 21"/><path d="m21 15-6 6"/>`)

// Handshake is [handshake](https://lucide.dev/icons/handshake).
var Handshake = icon.Lucide(viewBox, paint, `<path d="m11 17 2 2a1 1 0 1 0 3-3"/><path d="m14 14 2.5 2.5a1 1 0 1 0 3-3l-3.88-3.88a3 3 0 0 0-4.24 0l-.88.88a1 1 0 1 1-3-3l2.81-2.81a5.79 5.79 0 0 1 7.06-.87l.47.28a2 2 0 0 0 1.42.25L21 4"/><path d="m21 3 1 11h-2"/><path d="M3 3 2 14l6.5 6.5a1 1 0 1 0 3-3"/><path d="M3 4h8"/>`)

// HardDriveDownload is [hard-drive-download](https://lucide.dev/icons/hard-drive-download).
var HardDriveDownload = icon.Lucide(viewBox, paint, `<path d="M12 2v8"/><path d="m16 6-4 4-4-4"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h.01"/><path d="M10 18h.01"/>`)

// HardDriveUpload is [hard-drive-upload](https://lucide.dev/icons/hard-drive-upload).
var HardDriveUpload = icon.Lucide(viewBox, paint, `<path d="m16 6-4-4-4 4"/><path d="M12 2v8"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h.01"/><path d="M10 18h.01"/>`)

// HardDrive is [hard-drive](https://lucide.dev/icons/hard-drive).
var HardDrive = icon.Lucide(viewBox, paint, `<path d="M10 16h.01"/><path d="M2.212 11.577a2 2 0 0 0-.212.896V18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5.527a2 2 0 0 0-.212-.896L18.55 5.11A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"/><path d="M21.946 12.013H2.054"/><path d="M6 16h.01"/>`)

// HardHat is [hard-hat](https://lucide.dev/icons/hard-hat).
var HardHat = icon.Lucide(viewBox, paint, `<path d="M10 10V5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v5"/><path d="M14 6a6 6 0 0 1 6 6v3"/><path d="M4 15v-3a6 6 0 0 1 6-6"/><rect x="2" y="15" width="20" height="4" rx="1"/>`)

// Hash is [hash](https://lucide.dev/icons/hash).
var Hash = icon.Lucide(viewBox, paint, `<line x1="4" x2="20" y1="9" y2="9"/><line x1="4" x2="20" y1="15" y2="15"/><line x1="10" x2="8" y1="3" y2="21"/><line x1="16" x2="14" y1="3" y2="21"/>`)

// HatGlasses is [hat-glasses](https://lucide.dev/icons/hat-glasses).
var HatGlasses = icon.Lucide(viewBox, paint, `<path d="M14 18a2 2 0 0 0-4 0"/><path d="m19 11-2.11-6.657a2 2 0 0 0-2.752-1.148l-1.276.61A2 2 0 0 1 12 4H8.5a2 2 0 0 0-1.925 1.456L5 11"/><path d="M2 11h20"/><circle cx="17" cy="18" r="3"/><circle cx="7" cy="18" r="3"/>`)

// Haze is [haze](https://lucide.dev/icons/haze).
var Haze = icon.Lucide(viewBox, paint, `<path d="m5.2 6.2 1.4 1.4"/><path d="M2 13h2"/><path d="M20 13h2"/><path d="m17.4 7.6 1.4-1.4"/><path d="M22 17H2"/><path d="M22 21H2"/><path d="M16 13a4 4 0 0 0-8 0"/><path d="M12 5V2.5"/>`)

// Hd is [hd](https://lucide.dev/icons/hd).
var Hd = icon.Lucide(viewBox, paint, `<path d="M10 12H6"/><path d="M10 15V9"/><path d="M14 14.5a.5.5 0 0 0 .5.5h1a2.5 2.5 0 0 0 2.5-2.5v-1A2.5 2.5 0 0 0 15.5 9h-1a.5.5 0 0 0-.5.5z"/><path d="M6 15V9"/><rect x="2" y="5" width="20" height="14" rx="2"/>`)

// HdmiPort is [hdmi-port](https://lucide.dev/icons/hdmi-port).
var HdmiPort = icon.Lucide(viewBox, paint, `<path d="M22 9a1 1 0 00-1-1H3a1 1 0 00-1 1v4a1 1 0 001 1h.5a2 2 0 011.6.8l.3.4A2 2 0 007 16h10a2 2 0 001.6-.8l.3-.4a2 2 0 011.6-.8h.5a1 1 0 001-1z"/><path d="M8 12h8"/>`)

// Heading1 is [heading-1](https://lucide.dev/icons/heading-1).
var Heading1 = icon.Lucide(viewBox, paint, `<path d="M4 12h8"/><path d="M4 18V6"/><path d="M12 18V6"/><path d="m17 12 3-2v8"/>`)

// Heading2 is [heading-2](https://lucide.dev/icons/heading-2).
var Heading2 = icon.Lucide(viewBox, paint, `<path d="M4 12h8"/><path d="M4 18V6"/><path d="M12 18V6"/><path d="M21 18h-4c0-4 4-3 4-6 0-1.5-2-2.5-4-1"/>`)

// Heading3 is [heading-3](https://lucide.dev/icons/heading-3).
var Heading3 = icon.Lucide(viewBox, paint, `<path d="M4 12h8"/><path d="M4 18V6"/><path d="M12 18V6"/><path d="M17.5 10.5c1.7-1 3.5 0 3.5 1.5a2 2 0 0 1-2 2"/><path d="M17 17.5c2 1.5 4 .3 4-1.5a2 2 0 0 0-2-2"/>`)

// Heading4 is [heading-4](https://lucide.dev/icons/heading-4).
var Heading4 = icon.Lucide(viewBox, paint, `<path d="M12 18V6"/><path d="M17 10v3a1 1 0 0 0 1 1h3"/><path d="M21 10v8"/><path d="M4 12h8"/><path d="M4 18V6"/>`)

// Heading5 is [heading-5](https://lucide.dev/icons/heading-5).
var Heading5 = icon.Lucide(viewBox, paint, `<path d="M4 12h8"/><path d="M4 18V6"/><path d="M12 18V6"/><path d="M17 13v-3h4"/><path d="M17 17.7c.4.2.8.3 1.3.3 1.5 0 2.7-1.1 2.7-2.5S19.8 13 18.3 13H17"/>`)

// Heading6 is [heading-6](https://lucide.dev/icons/heading-6).
var Heading6 = icon.Lucide(viewBox, paint, `<path d="M4 12h8"/><path d="M4 18V6"/><path d="M12 18V6"/><circle cx="19" cy="16" r="2"/><path d="M20 10c-2 2-3 3.5-3 6"/>`)

// Heading is [heading](https://lucide.dev/icons/heading).
var Heading = icon.Lucide(viewBox, paint, `<path d="M6 12h12"/><path d="M6 20V4"/><path d="M18 20V4"/>`)

// HeadphoneOff is [headphone-off](https://lucide.dev/icons/headphone-off).
var HeadphoneOff = icon.Lucide(viewBox, paint, `<path d="M21 14h-1.343"/><path d="M9.128 3.47A9 9 0 0 1 21 12v3.343"/><path d="m2 2 20 20"/><path d="M20.414 20.414A2 2 0 0 1 19 21h-1a2 2 0 0 1-2-2v-3"/><path d="M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 2.636-6.364"/>`)

// Headphones is [headphones](https://lucide.dev/icons/headphones).
var Headphones = icon.Lucide(viewBox, paint, `<path d="M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 18 0v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3"/>`)

// Headset is [headset](https://lucide.dev/icons/headset).
var Headset = icon.Lucide(viewBox, paint, `<path d="M3 11h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5Zm0 0a9 9 0 1 1 18 0m0 0v5a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3Z"/><path d="M21 16v2a4 4 0 0 1-4 4h-5"/>`)

// HeartCrack is [heart-crack](https://lucide.dev/icons/heart-crack).
var HeartCrack = icon.Lucide(viewBox, paint, `<path d="M12.409 5.824c-.702.792-1.15 1.496-1.415 2.166l2.153 2.156a.5.5 0 0 1 0 .707l-2.293 2.293a.5.5 0 0 0 0 .707L12 15"/><path d="M13.508 20.313a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5a5.5 5.5 0 0 1 9.591-3.677.6.6 0 0 0 .818.001A5.5 5.5 0 0 1 22 9.5c0 2.29-1.5 4-3 5.5z"/>`)

// HeartHandshake is [heart-handshake](https://lucide.dev/icons/heart-handshake).
var HeartHandshake = icon.Lucide(viewBox, paint, `<path d="M19.414 14.414C21 12.828 22 11.5 22 9.5a5.5 5.5 0 0 0-9.591-3.676.6.6 0 0 1-.818.001A5.5 5.5 0 0 0 2 9.5c0 2.3 1.5 4 3 5.5l5.535 5.362a2 2 0 0 0 2.879.052 2.12 2.12 0 0 0-.004-3 2.124 2.124 0 1 0 3-3 2.124 2.124 0 0 0 3.004 0 2 2 0 0 0 0-2.828l-1.881-1.882a2.41 2.41 0 0 0-3.409 0l-1.71 1.71a2 2 0 0 1-2.828 0 2 2 0 0 1 0-2.828l2.823-2.762"/>`)

// HeartMinus is [heart-minus](https://lucide.dev/icons/heart-minus).
var HeartMinus = icon.Lucide(viewBox, paint, `<path d="m14.876 18.99-1.368 1.323a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5a5.5 5.5 0 0 1 9.591-3.676.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5a5.2 5.2 0 0 1-.244 1.572"/><path d="M15 15h6"/>`)

// HeartOff is [heart-off](https://lucide.dev/icons/heart-off).
var HeartOff = icon.Lucide(viewBox, paint, `<path d="M10.5 4.893a5.5 5.5 0 0 1 1.091.931.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5c0 1.872-1.002 3.356-2.187 4.655"/><path d="m16.967 16.967-3.459 3.346a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5a5.5 5.5 0 0 1 2.747-4.761"/><path d="m2 2 20 20"/>`)

// HeartPlus is [heart-plus](https://lucide.dev/icons/heart-plus).
var HeartPlus = icon.Lucide(viewBox, paint, `<path d="m14.479 19.374-.971.939a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5a5.5 5.5 0 0 1 9.591-3.676.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5a5.2 5.2 0 0 1-.219 1.49"/><path d="M15 15h6"/><path d="M18 12v6"/>`)

// HeartPulse is [heart-pulse](https://lucide.dev/icons/heart-pulse).
var HeartPulse = icon.Lucide(viewBox, paint, `<path d="M2 9.5a5.5 5.5 0 0 1 9.591-3.676.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5c0 2.29-1.5 4-3 5.5l-5.492 5.313a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5"/><path d="M3.22 13H9.5l.5-1 2 4.5 2-7 1.5 3.5h5.27"/>`)

// HeartX is [heart-x](https://lucide.dev/icons/heart-x).
var HeartX = icon.Lucide(viewBox, paint, `<path d="m15.5 12.5 5 5"/><path d="m20.5 12.5-5 5"/><path d="M21.955 8.774a5.5 5.5 0 0 0-9.546-2.95.6.6 0 0 1-.818 0A5.5 5.5 0 0 0 2 9.5c0 2.3 1.5 4 3 5.5l5.508 5.332a2 2 0 0 0 2.57.352"/>`)

// Heart is [heart](https://lucide.dev/icons/heart).
var Heart = icon.Lucide(viewBox, paint, `<path d="M2 9.5a5.5 5.5 0 0 1 9.591-3.676.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5c0 2.29-1.5 4-3 5.5l-5.492 5.313a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5"/>`)

// Heater is [heater](https://lucide.dev/icons/heater).
var Heater = icon.Lucide(viewBox, paint, `<path d="M11 8c2-3-2-3 0-6"/><path d="M15.5 8c2-3-2-3 0-6"/><path d="M6 10h.01"/><path d="M6 14h.01"/><path d="M10 16v-4"/><path d="M14 16v-4"/><path d="M18 16v-4"/><path d="M20 6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3"/><path d="M5 20v2"/><path d="M19 20v2"/>`)

// Helicopter is [helicopter](https://lucide.dev/icons/helicopter).
var Helicopter = icon.Lucide(viewBox, paint, `<path d="M11 17v4"/><path d="M14 3v8a2 2 0 0 0 2 2h5.865"/><path d="M17 17v4"/><path d="M18 17a4 4 0 0 0 4-4 8 6 0 0 0-8-6 6 5 0 0 0-6 5v3a2 2 0 0 0 2 2z"/><path d="M2 10v5"/><path d="M6 3h16"/><path d="M7 21h14"/><path d="M8 13H2"/>`)

// Hexagon is [hexagon](https://lucide.dev/icons/hexagon).
var Hexagon = icon.Lucide(viewBox, paint, `<path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>`)

// Highlighter is [highlighter](https://lucide.dev/icons/highlighter).
var Highlighter = icon.Lucide(viewBox, paint, `<path d="m9 11-6 6v3h9l3-3"/><path d="m22 12-4.6 4.6a2 2 0 0 1-2.8 0l-5.2-5.2a2 2 0 0 1 0-2.8L14 4"/>`)

// HikingStick is [hiking-stick](https://lucide.dev/icons/hiking-stick).
var HikingStick = icon.Lucide(viewBox, paint, `<path d="M13.5 10.5 2 22"/><path d="M16.352 11.648a1.205 1.205 0 01-1.704 0l-2.296-2.296a1.205 1.205 0 010-1.704l.72-.72a2 2 0 011.022-.546l.599-.12a2 2 0 001.569-1.57l.12-.598a2 2 0 01.546-1.022l.72-.72a1.205 1.205 0 011.704 0l2.296 2.296a1.205 1.205 0 010 1.704l-6.02 6.02a1 1 0 103 3l.201-.201A7.4 7.4 0 0021 9.93V7"/><path d="m6 21-3-3"/>`)

// HopOff is [hop-off](https://lucide.dev/icons/hop-off).
var HopOff = icon.Lucide(viewBox, paint, `<path d="M10.82 16.12c1.69.6 3.91.79 5.18.85.28.01.53-.09.7-.27"/><path d="M11.14 20.57c.52.24 2.44 1.12 4.08 1.37.46.06.86-.25.9-.71.12-1.52-.3-3.43-.5-4.28"/><path d="M16.13 21.05c1.65.63 3.68.84 4.87.91a.9.9 0 0 0 .7-.26"/><path d="M17.99 5.52a20.83 20.83 0 0 1 3.15 4.5.8.8 0 0 1-.68 1.13c-1.17.1-2.5.02-3.9-.25"/><path d="M20.57 11.14c.24.52 1.12 2.44 1.37 4.08.04.3-.08.59-.31.75"/><path d="M4.93 4.93a10 10 0 0 0-.67 13.4c.35.43.96.4 1.17-.12.69-1.71 1.07-5.07 1.07-6.71 1.34.45 3.1.9 4.88.62a.85.85 0 0 0 .48-.24"/><path d="M5.52 17.99c1.05.95 2.91 2.42 4.5 3.15a.8.8 0 0 0 1.13-.68c.2-2.34-.33-5.3-1.57-8.28"/><path d="M8.35 2.68a10 10 0 0 1 9.98 1.58c.43.35.4.96-.12 1.17-1.5.6-4.3.98-6.07 1.05"/><path d="m2 2 20 20"/>`)

// Hop is [hop](https://lucide.dev/icons/hop).
var Hop = icon.Lucide(viewBox, paint, `<path d="M10.82 16.12c1.69.6 3.91.79 5.18.85.55.03 1-.42.97-.97-.06-1.27-.26-3.5-.85-5.18"/><path d="M11.5 6.5c1.64 0 5-.38 6.71-1.07.52-.2.55-.82.12-1.17A10 10 0 0 0 4.26 18.33c.35.43.96.4 1.17-.12.69-1.71 1.07-5.07 1.07-6.71 1.34.45 3.1.9 4.88.62a.88.88 0 0 0 .73-.74c.3-2.14-.15-3.5-.61-4.88"/><path d="M15.62 16.95c.2.85.62 2.76.5 4.28a.77.77 0 0 1-.9.7 16.64 16.64 0 0 1-4.08-1.36"/><path d="M16.13 21.05c1.65.63 3.68.84 4.87.91a.9.9 0 0 0 .96-.96 17.68 17.68 0 0 0-.9-4.87"/><path d="M16.94 15.62c.86.2 2.77.62 4.29.5a.77.77 0 0 0 .7-.9 16.64 16.64 0 0 0-1.36-4.08"/><path d="M17.99 5.52a20.82 20.82 0 0 1 3.15 4.5.8.8 0 0 1-.68 1.13c-2.33.2-5.3-.32-8.27-1.57"/><path d="M4.93 4.93 3 3a.7.7 0 0 1 0-1"/><path d="M9.58 12.18c1.24 2.98 1.77 5.95 1.57 8.28a.8.8 0 0 1-1.13.68 20.82 20.82 0 0 1-4.5-3.15"/>`)

// Hospital is [hospital](https://lucide.dev/icons/hospital).
var Hospital = icon.Lucide(viewBox, paint, `<path d="M12 7v4"/><path d="M14 21v-3a2 2 0 0 0-4 0v3"/><path d="M14 9h-4"/><path d="M18 11h2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h2"/><path d="M18 21V5a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16"/>`)

// Hotel is [hotel](https://lucide.dev/icons/hotel).
var Hotel = icon.Lucide(viewBox, paint, `<path d="M10 22v-6.57"/><path d="M12 11h.01"/><path d="M12 7h.01"/><path d="M14 15.43V22"/><path d="M15 16a5 5 0 0 0-6 0"/><path d="M16 11h.01"/><path d="M16 7h.01"/><path d="M8 11h.01"/><path d="M8 7h.01"/><rect x="4" y="2" width="16" height="20" rx="2"/>`)

// HourglassCog is [hourglass-cog](https://lucide.dev/icons/hourglass-cog).
var HourglassCog = icon.Lucide(viewBox, paint, `<path d="m14.305 19.53.923-.382"/><path d="m15.228 16.852-.923-.383"/><path d="m16.852 15.228-.383-.923"/><path d="m16.852 20.772-.383.924"/><path d="M17 2v4.172a2 2 0 0 1-.586 1.414l-8.828 8.828A2 2 0 0 0 7 17.828V22"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.696-.382-.924"/><path d="m20.772 16.852.924-.383"/><path d="m20.772 19.148.924.383"/><path d="M5 22h6.159"/><path d="M5 2h14"/><path d="M7 2v4.172a2 2 0 0 0 .586 1.414l5.188 5.188"/><circle cx="18" cy="18" r="3"/>`)

// Hourglass is [hourglass](https://lucide.dev/icons/hourglass).
var Hourglass = icon.Lucide(viewBox, paint, `<path d="M5 22h14"/><path d="M5 2h14"/><path d="M17 22v-4.172a2 2 0 0 0-.586-1.414L12 12l-4.414 4.414A2 2 0 0 0 7 17.828V22"/><path d="M7 2v4.172a2 2 0 0 0 .586 1.414L12 12l4.414-4.414A2 2 0 0 0 17 6.172V2"/>`)

// HouseCog is [house-cog](https://lucide.dev/icons/house-cog).
var HouseCog = icon.Lucide(viewBox, paint, `<path d="M10.584 21H5a2 2 0 01-2-2v-9a2 2 0 01.709-1.527l7-6a2 2 0 012.582 0l7 6A2 2 0 0121 10.001v.583"/><path d="M14 12H10a1 1 0 00-1 1v8"/><path d="m14.305 19.53.923-.382"/><path d="m15.229 16.852-.924-.383"/><path d="m16.852 15.228-.383-.923"/><path d="m16.852 20.773-.383.924"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.697-.382-.924"/><path d="m20.773 16.852.922-.383"/><path d="m20.773 19.148.922.383"/><circle cx="18" cy="18" r="3"/>`)

// HouseHeart is [house-heart](https://lucide.dev/icons/house-heart).
var HouseHeart = icon.Lucide(viewBox, paint, `<path d="M8.62 13.8A2.25 2.25 0 1 1 12 10.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"/><path d="M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>`)

// HousePlug is [house-plug](https://lucide.dev/icons/house-plug).
var HousePlug = icon.Lucide(viewBox, paint, `<path d="M10 12V8.964"/><path d="M14 12V8.964"/><path d="M15 12a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-2a1 1 0 0 1 1-1z"/><path d="M8.5 21H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2v-2"/>`)

// HousePlus is [house-plus](https://lucide.dev/icons/house-plus).
var HousePlus = icon.Lucide(viewBox, paint, `<path d="M12.35 21H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 .71-1.53l7-6a2 2 0 0 1 2.58 0l7 6A2 2 0 0 1 21 10v2.35"/><path d="M14.8 12.4A1 1 0 0 0 14 12h-4a1 1 0 0 0-1 1v8"/><path d="M15 18h6"/><path d="M18 15v6"/>`)

// HouseWifi is [house-wifi](https://lucide.dev/icons/house-wifi).
var HouseWifi = icon.Lucide(viewBox, paint, `<path d="M9.5 13.866a4 4 0 0 1 5 .01"/><path d="M12 17h.01"/><path d="M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M7 10.754a8 8 0 0 1 10 0"/>`)

// House is [house](https://lucide.dev/icons/house).
var House = icon.Lucide(viewBox, paint, `<path d="M15 21v-8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v8"/><path d="M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>`)

// Houses is [houses](https://lucide.dev/icons/houses).
var Houses = icon.Lucide(viewBox, paint, `<path d="m12.681 4.24.834-.715a1.45 1.45 0 011.88 0l5.09 4.364A1.45 1.45 0 0121 9v6.546a1.45 1.45 0 01-1 1.381"/><path d="M15.485 11.889A1.45 1.45 0 0116 13v6.546A1.454 1.454 0 0114.546 21H4.364a1.454 1.454 0 01-1.454-1.454V13a1.45 1.45 0 01.515-1.111l5.09-4.364a1.45 1.45 0 011.88 0z"/><path d="M7.41 20.546v-4a1 1 0 011-1h2a1 1 0 011 1v4"/>`)

// IceCreamBowl is [ice-cream-bowl](https://lucide.dev/icons/ice-cream-bowl).
var IceCreamBowl = icon.Lucide(viewBox, paint, `<path d="M12 17c5 0 8-2.69 8-6H4c0 3.31 3 6 8 6m-4 4h8m-4-3v3M5.14 11a3.5 3.5 0 1 1 6.71 0"/><path d="M12.14 11a3.5 3.5 0 1 1 6.71 0"/><path d="M15.5 6.5a3.5 3.5 0 1 0-7 0"/>`)

// IceCreamCone is [ice-cream-cone](https://lucide.dev/icons/ice-cream-cone).
var IceCreamCone = icon.Lucide(viewBox, paint, `<path d="m7 11 4.08 10.35a1 1 0 0 0 1.84 0L17 11"/><path d="M17 7A5 5 0 0 0 7 7"/><path d="M17 7a2 2 0 0 1 0 4H7a2 2 0 0 1 0-4"/>`)

// IdCardLanyard is [id-card-lanyard](https://lucide.dev/icons/id-card-lanyard).
var IdCardLanyard = icon.Lucide(viewBox, paint, `<path d="M13.5 8h-3"/><path d="m15 2-1 2h3a2 2 0 012 2v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6a2 2 0 012-2h3"/><path d="M16 22a4 4 0 00-8 0"/><path d="m9 2 3 6"/><circle cx="12" cy="15" r="3"/>`)

// IdCard is [id-card](https://lucide.dev/icons/id-card).
var IdCard = icon.Lucide(viewBox, paint, `<path d="M13 19a4 4 0 00-8 0"/><path d="M16 10h2"/><path d="M16 14h2"/><circle cx="9" cy="12" r="3"/><rect x="2" y="5" width="20" height="14" rx="2"/>`)

// ImageDown is [image-down](https://lucide.dev/icons/image-down).
var ImageDown = icon.Lucide(viewBox, paint, `<path d="M10.3 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10l-3.1-3.1a2 2 0 0 0-2.814.014L6 21"/><path d="m14 19 3 3v-5.5"/><path d="m17 22 3-3"/><circle cx="9" cy="9" r="2"/>`)

// ImageMinus is [image-minus](https://lucide.dev/icons/image-minus).
var ImageMinus = icon.Lucide(viewBox, paint, `<path d="M21 9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7"/><line x1="16" x2="22" y1="5" y2="5"/><circle cx="9" cy="9" r="2"/><path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/>`)

// ImageOff is [image-off](https://lucide.dev/icons/image-off).
var ImageOff = icon.Lucide(viewBox, paint, `<line x1="2" x2="22" y1="2" y2="22"/><path d="M10.41 10.41a2 2 0 1 1-2.83-2.83"/><line x1="13.5" x2="6" y1="13.5" y2="21"/><line x1="18" x2="21" y1="12" y2="15"/><path d="M3.59 3.59A1.99 1.99 0 0 0 3 5v14a2 2 0 0 0 2 2h14c.55 0 1.052-.22 1.41-.59"/><path d="M21 15V5a2 2 0 0 0-2-2H9"/>`)

// ImagePlay is [image-play](https://lucide.dev/icons/image-play).
var ImagePlay = icon.Lucide(viewBox, paint, `<path d="M15 15.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z"/><path d="M21 12.17V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"/><path d="m6 21 5-5"/><circle cx="9" cy="9" r="2"/>`)

// ImagePlus is [image-plus](https://lucide.dev/icons/image-plus).
var ImagePlus = icon.Lucide(viewBox, paint, `<path d="M16 5h6"/><path d="M19 2v6"/><path d="M21 11.5V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7.5"/><path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/><circle cx="9" cy="9" r="2"/>`)

// ImageUp is [image-up](https://lucide.dev/icons/image-up).
var ImageUp = icon.Lucide(viewBox, paint, `<path d="M10.3 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10l-3.1-3.1a2 2 0 0 0-2.814.014L6 21"/><path d="m14 19.5 3-3 3 3"/><path d="M17 22v-5.5"/><circle cx="9" cy="9" r="2"/>`)

// ImageUpscale is [image-upscale](https://lucide.dev/icons/image-upscale).
var ImageUpscale = icon.Lucide(viewBox, paint, `<path d="M16 3h5v5"/><path d="M17 21h2a2 2 0 0 0 2-2"/><path d="M21 12v3"/><path d="m21 3-5 5"/><path d="M3 7V5a2 2 0 0 1 2-2"/><path d="m5 21 4.144-4.144a1.21 1.21 0 0 1 1.712 0L13 19"/><path d="M9 3h3"/><rect x="3" y="11" width="10" height="10" rx="1"/>`)

// Image is [image](https://lucide.dev/icons/image).
var Image = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><circle cx="9" cy="9" r="2"/><path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/>`)

// Images is [images](https://lucide.dev/icons/images).
var Images = icon.Lucide(viewBox, paint, `<path d="m22 11-1.296-1.296a2.4 2.4 0 0 0-3.408 0L11 16"/><path d="M4 8a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2"/><circle cx="13" cy="7" r="1" fill="currentColor"/><rect x="8" y="2" width="14" height="14" rx="2"/>`)

// Import is [import](https://lucide.dev/icons/import).
var Import = icon.Lucide(viewBox, paint, `<path d="M12 3v12"/><path d="m8 11 4 4 4-4"/><path d="M8 5H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-4"/>`)

// Inbox is [inbox](https://lucide.dev/icons/inbox).
var Inbox = icon.Lucide(viewBox, paint, `<polyline points="22 12 16 12 14 15 10 15 8 12 2 12"/><path d="M5.45 5.11 2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"/>`)

// IndianRupee is [indian-rupee](https://lucide.dev/icons/indian-rupee).
var IndianRupee = icon.Lucide(viewBox, paint, `<path d="M6 3h12"/><path d="M6 8h12"/><path d="m6 13 8.5 8"/><path d="M6 13h3"/><path d="M9 13c6.667 0 6.667-10 0-10"/>`)

// Infinity is [infinity](https://lucide.dev/icons/infinity).
var Infinity = icon.Lucide(viewBox, paint, `<path d="M6 16c5 0 7-8 12-8a4 4 0 0 1 0 8c-5 0-7-8-12-8a4 4 0 1 0 0 8"/>`)

// Info is [info](https://lucide.dev/icons/info).
var Info = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/>`)

// InspectionPanel is [inspection-panel](https://lucide.dev/icons/inspection-panel).
var InspectionPanel = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 7h.01"/><path d="M17 7h.01"/><path d="M7 17h.01"/><path d="M17 17h.01"/>`)

// Italic is [italic](https://lucide.dev/icons/italic).
var Italic = icon.Lucide(viewBox, paint, `<line x1="19" x2="10" y1="4" y2="4"/><line x1="14" x2="5" y1="20" y2="20"/><line x1="15" x2="9" y1="4" y2="20"/>`)

// IterationCcw is [iteration-ccw](https://lucide.dev/icons/iteration-ccw).
var IterationCcw = icon.Lucide(viewBox, paint, `<path d="m16 14 4 4-4 4"/><path d="M20 10a8 8 0 1 0-8 8h8"/>`)

// IterationCw is [iteration-cw](https://lucide.dev/icons/iteration-cw).
var IterationCw = icon.Lucide(viewBox, paint, `<path d="M4 10a8 8 0 1 1 8 8H4"/><path d="m8 22-4-4 4-4"/>`)

// IvBag is [iv-bag](https://lucide.dev/icons/iv-bag).
var IvBag = icon.Lucide(viewBox, paint, `<path d="M12 18v2a2 2 0 002 2h6"/><path d="M6 11c.72.5 1.44 1 3 1 3 0 3-2 6-2 1.56 0 2.28.5 3 1"/><path d="M9.293 3c.453 0 .887-.18 1.207-.5s.754-.5 1.207-.5h.586c.453 0 .887.18 1.207.5s.754.5 1.207.5H16a2 2 0 012 2v11a2 2 0 01-2 2H8a2 2 0 01-2-2V5a2 2 0 012-2z"/>`)

// JapaneseYen is [japanese-yen](https://lucide.dev/icons/japanese-yen).
var JapaneseYen = icon.Lucide(viewBox, paint, `<path d="M12 9.5V21m0-11.5L6 3m6 6.5L18 3"/><path d="M6 15h12"/><path d="M6 11h12"/>`)

// Joystick is [joystick](https://lucide.dev/icons/joystick).
var Joystick = icon.Lucide(viewBox, paint, `<path d="M21 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2Z"/><path d="M6 15v-2"/><path d="M12 15V9"/><circle cx="12" cy="6" r="3"/>`)

// Kanban is [kanban](https://lucide.dev/icons/kanban).
var Kanban = icon.Lucide(viewBox, paint, `<path d="M5 3v14"/><path d="M12 3v8"/><path d="M19 3v18"/>`)

// Kayak is [kayak](https://lucide.dev/icons/kayak).
var Kayak = icon.Lucide(viewBox, paint, `<path d="M18 17a1 1 0 0 0-1 1v1a2 2 0 1 0 2-2z"/><path d="M20.97 3.61a.45.45 0 0 0-.58-.58C10.2 6.6 6.6 10.2 3.03 20.39a.45.45 0 0 0 .58.58C13.8 17.4 17.4 13.8 20.97 3.61"/><path d="m6.707 6.707 10.586 10.586"/><path d="M7 5a2 2 0 1 0-2 2h1a1 1 0 0 0 1-1z"/>`)

// KazakhTenge is [kazakh-tenge](https://lucide.dev/icons/kazakh-tenge).
var KazakhTenge = icon.Lucide(viewBox, paint, `<path d="M12 8v12"/><path d="M6 4h12"/><path d="M6 8h12"/>`)

// KeyRound is [key-round](https://lucide.dev/icons/key-round).
var KeyRound = icon.Lucide(viewBox, paint, `<path d="M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z"/><circle cx="16.5" cy="7.5" r=".5" fill="currentColor"/>`)

// KeySquare is [key-square](https://lucide.dev/icons/key-square).
var KeySquare = icon.Lucide(viewBox, paint, `<path d="M12.4 2.7a2.5 2.5 0 0 1 3.4 0l5.5 5.5a2.5 2.5 0 0 1 0 3.4l-3.7 3.7a2.5 2.5 0 0 1-3.4 0L8.7 9.8a2.5 2.5 0 0 1 0-3.4z"/><path d="m14 7 3 3"/><path d="m9.4 10.6-6.814 6.814A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814"/>`)

// Key is [key](https://lucide.dev/icons/key).
var Key = icon.Lucide(viewBox, paint, `<path d="m2 21 9.6-9.6"/><path d="m7.5 15.5 2.3 2.3a1 1 0 0 1 0 1.4l-2.1 2.1a1 1 0 0 1-1.4 0L4 19"/><circle cx="15.5" cy="7.5" r="5.5"/>`)

// KeyboardMusic is [keyboard-music](https://lucide.dev/icons/keyboard-music).
var KeyboardMusic = icon.Lucide(viewBox, paint, `<rect width="20" height="16" x="2" y="4" rx="2"/><path d="M6 8h4"/><path d="M14 8h.01"/><path d="M18 8h.01"/><path d="M2 12h20"/><path d="M6 12v4"/><path d="M10 12v4"/><path d="M14 12v4"/><path d="M18 12v4"/>`)

// KeyboardOff is [keyboard-off](https://lucide.dev/icons/keyboard-off).
var KeyboardOff = icon.Lucide(viewBox, paint, `<path d="M 20 4 A2 2 0 0 1 22 6"/><path d="M 22 6 L 22 16.41"/><path d="M 7 16 L 16 16"/><path d="M 9.69 4 L 20 4"/><path d="M14 8h.01"/><path d="M18 8h.01"/><path d="m2 2 20 20"/><path d="M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2"/><path d="M6 8h.01"/><path d="M8 12h.01"/>`)

// Keyboard is [keyboard](https://lucide.dev/icons/keyboard).
var Keyboard = icon.Lucide(viewBox, paint, `<path d="M10 8h.01"/><path d="M12 12h.01"/><path d="M14 8h.01"/><path d="M16 12h.01"/><path d="M18 8h.01"/><path d="M6 8h.01"/><path d="M7 16h10"/><path d="M8 12h.01"/><rect width="20" height="16" x="2" y="4" rx="2"/>`)

// Lambda is [lambda](https://lucide.dev/icons/lambda).
var Lambda = icon.Lucide(viewBox, paint, `<path d="M11.38 10 5 20"/><path d="M19 18a2 2 0 01-2 2c-4.87-.003-5.052-16-10-16a2 2 0 00-2 2"/>`)

// LampCeiling is [lamp-ceiling](https://lucide.dev/icons/lamp-ceiling).
var LampCeiling = icon.Lucide(viewBox, paint, `<path d="M12 2v5"/><path d="M14.829 15.998a3 3 0 1 1-5.658 0"/><path d="M20.92 14.606A1 1 0 0 1 20 16H4a1 1 0 0 1-.92-1.394l3-7A1 1 0 0 1 7 7h10a1 1 0 0 1 .92.606z"/>`)

// LampDesk is [lamp-desk](https://lucide.dev/icons/lamp-desk).
var LampDesk = icon.Lucide(viewBox, paint, `<path d="M10.293 2.293a1 1 0 0 1 1.414 0l2.5 2.5 5.994 1.227a1 1 0 0 1 .506 1.687l-7 7a1 1 0 0 1-1.687-.506l-1.227-5.994-2.5-2.5a1 1 0 0 1 0-1.414z"/><path d="m14.207 4.793-3.414 3.414"/><path d="M3 20a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/><path d="m9.086 6.5-4.793 4.793a1 1 0 0 0-.18 1.17L7 18"/>`)

// LampFloor is [lamp-floor](https://lucide.dev/icons/lamp-floor).
var LampFloor = icon.Lucide(viewBox, paint, `<path d="M12 10v12"/><path d="M17.929 7.629A1 1 0 0 1 17 9H7a1 1 0 0 1-.928-1.371l2-5A1 1 0 0 1 9 2h6a1 1 0 0 1 .928.629z"/><path d="M9 22h6"/>`)

// LampWallDown is [lamp-wall-down](https://lucide.dev/icons/lamp-wall-down).
var LampWallDown = icon.Lucide(viewBox, paint, `<path d="M19.929 18.629A1 1 0 0 1 19 20H9a1 1 0 0 1-.928-1.371l2-5A1 1 0 0 1 11 13h6a1 1 0 0 1 .928.629z"/><path d="M6 3a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1z"/><path d="M8 6h4a2 2 0 0 1 2 2v5"/>`)

// LampWallUp is [lamp-wall-up](https://lucide.dev/icons/lamp-wall-up).
var LampWallUp = icon.Lucide(viewBox, paint, `<path d="M19.929 9.629A1 1 0 0 1 19 11H9a1 1 0 0 1-.928-1.371l2-5A1 1 0 0 1 11 4h6a1 1 0 0 1 .928.629z"/><path d="M6 15a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"/><path d="M8 18h4a2 2 0 0 0 2-2v-5"/>`)

// Lamp is [lamp](https://lucide.dev/icons/lamp).
var Lamp = icon.Lucide(viewBox, paint, `<path d="M12 12v6"/><path d="M4.077 10.615A1 1 0 0 0 5 12h14a1 1 0 0 0 .923-1.385l-3.077-7.384A2 2 0 0 0 15 2H9a2 2 0 0 0-1.846 1.23Z"/><path d="M8 20a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1z"/>`)

// LandPlot is [land-plot](https://lucide.dev/icons/land-plot).
var LandPlot = icon.Lucide(viewBox, paint, `<path d="m12 8 6-3-6-3v10"/><path d="m8 11.99-5.5 3.14a1 1 0 0 0 0 1.74l8.5 4.86a2 2 0 0 0 2 0l8.5-4.86a1 1 0 0 0 0-1.74L16 12"/><path d="m6.49 12.85 11.02 6.3"/><path d="M17.51 12.85 6.5 19.15"/>`)

// Landmark is [landmark](https://lucide.dev/icons/landmark).
var Landmark = icon.Lucide(viewBox, paint, `<path d="M10 18v-7"/><path d="M11.119 2.205a2 2 0 0 1 1.762 0l7.84 3.846A.5.5 0 0 1 20.5 7h-17a.5.5 0 0 1-.22-.949z"/><path d="M14 18v-7"/><path d="M18 18v-7"/><path d="M3 22h18"/><path d="M6 18v-7"/>`)

// Languages is [languages](https://lucide.dev/icons/languages).
var Languages = icon.Lucide(viewBox, paint, `<path d="m5 8 6 6"/><path d="m4 14 6-6 2-3"/><path d="M2 5h12"/><path d="M7 2h1"/><path d="m22 22-5-10-5 10"/><path d="M14 18h6"/>`)

// LaptopMinimalCheck is [laptop-minimal-check](https://lucide.dev/icons/laptop-minimal-check).
var LaptopMinimalCheck = icon.Lucide(viewBox, paint, `<path d="M2 20h20"/><path d="m9 10 2 2 4-4"/><rect x="3" y="4" width="18" height="12" rx="2"/>`)

// LaptopMinimal is [laptop-minimal](https://lucide.dev/icons/laptop-minimal).
var LaptopMinimal = icon.Lucide(viewBox, paint, `<rect width="18" height="12" x="3" y="4" rx="2" ry="2"/><line x1="2" x2="22" y1="20" y2="20"/>`)

// Laptop is [laptop](https://lucide.dev/icons/laptop).
var Laptop = icon.Lucide(viewBox, paint, `<path d="M18 5a2 2 0 0 1 2 2v8.526a2 2 0 0 0 .212.897l1.068 2.127a1 1 0 0 1-.9 1.45H3.62a1 1 0 0 1-.9-1.45l1.068-2.127A2 2 0 0 0 4 15.526V7a2 2 0 0 1 2-2z"/><path d="M20.054 15.987H3.946"/>`)

// LassoSelect is [lasso-select](https://lucide.dev/icons/lasso-select).
var LassoSelect = icon.Lucide(viewBox, paint, `<path d="M7 22a5 5 0 0 1-2-4"/><path d="M7 16.93c.96.43 1.96.74 2.99.91"/><path d="M3.34 14A6.8 6.8 0 0 1 2 10c0-4.42 4.48-8 10-8s10 3.58 10 8a7.19 7.19 0 0 1-.33 2"/><path d="M5 18a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"/><path d="M14.33 22h-.09a.35.35 0 0 1-.24-.32v-10a.34.34 0 0 1 .33-.34c.08 0 .15.03.21.08l7.34 6a.33.33 0 0 1-.21.59h-4.49l-2.57 3.85a.35.35 0 0 1-.28.14z"/>`)

// Lasso is [lasso](https://lucide.dev/icons/lasso).
var Lasso = icon.Lucide(viewBox, paint, `<path d="M3.704 14.467a10 8 0 1 1 3.115 2.375"/><path d="M7 22a5 5 0 0 1-2-3.994"/><circle cx="5" cy="16" r="2"/>`)

// LayerArrowDown is [layer-arrow-down](https://lucide.dev/icons/layer-arrow-down).
var LayerArrowDown = icon.Lucide(viewBox, paint, `<path d="M12 10v10"/><path d="M22 10a1 1 0 01-.59.92l-5.077 2.308"/><path d="M22.017 10.005a1 1 0 00-.597-.916l-8.59-3.91a2 2 0 00-1.66.001L2.6 9.08a1 1 0 00-.02 1.831l5.093 2.316"/><path d="m9 17 3 3 3-3"/>`)

// LayerArrowUp is [layer-arrow-up](https://lucide.dev/icons/layer-arrow-up).
var LayerArrowUp = icon.Lucide(viewBox, paint, `<path d="M12 14V4"/><path d="M7.674 10.774 2.58 13.09a1 1 0 000 1.822l8.6 3.91a2 2 0 001.65 0l8.58-3.9a1 1 0 00.59-.92 1 1 0 00-.59-.922l-5.078-2.308"/><path d="m9 7 3-3 3 3"/>`)

// Layers2 is [layers-2](https://lucide.dev/icons/layers-2).
var Layers2 = icon.Lucide(viewBox, paint, `<path d="M13 13.74a2 2 0 0 1-2 0L2.5 8.87a1 1 0 0 1 0-1.74L11 2.26a2 2 0 0 1 2 0l8.5 4.87a1 1 0 0 1 0 1.74z"/><path d="m20 14.285 1.5.845a1 1 0 0 1 0 1.74L13 21.74a2 2 0 0 1-2 0l-8.5-4.87a1 1 0 0 1 0-1.74l1.5-.845"/>`)

// LayersArrowDown is [layers-arrow-down](https://lucide.dev/icons/layers-arrow-down).
var LayersArrowDown = icon.Lucide(viewBox, paint, `<path d="M12 7v15"/><path d="M2 12a1 1 0 00.58.91l5.093 2.316"/><path d="M22 12a1 1 0 01-.59.92l-5.077 2.308"/><path d="M8 10.37 2.6 7.91a1 1 0 010-1.831l8.57-3.9a2 2 0 011.66.001l8.59 3.91a1 1 0 010 1.831l-5.392 2.45"/><path d="m9 19 3 3 3-3"/>`)

// LayersArrowUp is [layers-arrow-up](https://lucide.dev/icons/layers-arrow-up).
var LayersArrowUp = icon.Lucide(viewBox, paint, `<path d="M12 12V2"/><path d="M2 17.002a1 1 0 00.58.91l8.6 3.91a2 2 0 001.65 0l8.58-3.9a1 1 0 00.59-.92"/><path d="M7.674 8.774 2.58 11.09a1 1 0 000 1.822l8.6 3.91a2 2 0 001.65 0l8.58-3.9a1 1 0 00.59-.92 1 1 0 00-.59-.922l-5.078-2.308"/><path d="m9 5 3-3 3 3"/>`)

// LayersMinus is [layers-minus](https://lucide.dev/icons/layers-minus).
var LayersMinus = icon.Lucide(viewBox, paint, `<path d="M12.83 2.18a2 2 0 0 0-1.66 0L2.6 6.08a1 1 0 0 0 0 1.83l8.58 3.91a2 2 0 0 0 .83.18 2 2 0 0 0 .83-.18l8.58-3.9a1 1 0 0 0 0-1.832z"/><path d="M16 17h6"/><path d="M2.003 11.995a1 1 0 0 0 .597.915l8.58 3.91a2 2 0 0 0 .83.18"/><path d="M2.003 16.995a1 1 0 0 0 .597.915l8.58 3.91a2 2 0 0 0 .83.18 2 2 0 0 0 .83-.18l2.11-.96"/><path d="M22.018 12.004a1 1 0 0 1-.598.916l-.177.08"/>`)

// LayersPlus is [layers-plus](https://lucide.dev/icons/layers-plus).
var LayersPlus = icon.Lucide(viewBox, paint, `<path d="M12.83 2.18a2 2 0 0 0-1.66 0L2.6 6.08a1 1 0 0 0 0 1.83l8.58 3.91a2 2 0 0 0 .83.18 2 2 0 0 0 .83-.18l8.58-3.9a1 1 0 0 0 0-1.831z"/><path d="M16 17h6"/><path d="M19 14v6"/><path d="M2 12a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 .825.178"/><path d="M2 17a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 1.65 0l2.116-.962"/>`)

// Layers is [layers](https://lucide.dev/icons/layers).
var Layers = icon.Lucide(viewBox, paint, `<path d="M12.83 2.18a2 2 0 0 0-1.66 0L2.6 6.08a1 1 0 0 0 0 1.83l8.58 3.91a2 2 0 0 0 1.66 0l8.58-3.9a1 1 0 0 0 0-1.83z"/><path d="M2 12a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9A1 1 0 0 0 22 12"/><path d="M2 17a1 1 0 0 0 .58.91l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9A1 1 0 0 0 22 17"/>`)

// LayoutArrowDown is [layout-arrow-down](https://lucide.dev/icons/layout-arrow-down).
var LayoutArrowDown = icon.Lucide(viewBox, paint, `<rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><path d="M18 3v18"/><path d="m21 18-3 3-3-3"/>`)

// LayoutArrowRight is [layout-arrow-right](https://lucide.dev/icons/layout-arrow-right).
var LayoutArrowRight = icon.Lucide(viewBox, paint, `<rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><path d="M3 18h18"/><path d="m18 21 3-3-3-3"/>`)

// LayoutDashboard is [layout-dashboard](https://lucide.dev/icons/layout-dashboard).
var LayoutDashboard = icon.Lucide(viewBox, paint, `<rect width="7" height="9" x="3" y="3" rx="1"/><rect width="7" height="5" x="14" y="3" rx="1"/><rect width="7" height="9" x="14" y="12" rx="1"/><rect width="7" height="5" x="3" y="16" rx="1"/>`)

// LayoutFreeform is [layout-freeform](https://lucide.dev/icons/layout-freeform).
var LayoutFreeform = icon.Lucide(viewBox, paint, `<rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="4" rx="1"/><rect width="7" height="7" x="4" y="14" rx="1"/>`)

// LayoutGridCircles is [layout-grid-circles](https://lucide.dev/icons/layout-grid-circles).
var LayoutGridCircles = icon.Lucide(viewBox, paint, `<circle cx="17.5" cy="17.5" r="3.5"/><circle cx="17.5" cy="6.5" r="3.5"/><circle cx="6.5" cy="17.5" r="3.5"/><circle cx="6.5" cy="6.5" r="3.5"/>`)

// LayoutGrid is [layout-grid](https://lucide.dev/icons/layout-grid).
var LayoutGrid = icon.Lucide(viewBox, paint, `<rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/>`)

// LayoutList is [layout-list](https://lucide.dev/icons/layout-list).
var LayoutList = icon.Lucide(viewBox, paint, `<rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><path d="M14 4h7"/><path d="M14 9h7"/><path d="M14 15h7"/><path d="M14 20h7"/>`)

// LayoutPanelLeft is [layout-panel-left](https://lucide.dev/icons/layout-panel-left).
var LayoutPanelLeft = icon.Lucide(viewBox, paint, `<rect width="7" height="18" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/>`)

// LayoutPanelTop is [layout-panel-top](https://lucide.dev/icons/layout-panel-top).
var LayoutPanelTop = icon.Lucide(viewBox, paint, `<rect width="18" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/>`)

// LayoutTemplate is [layout-template](https://lucide.dev/icons/layout-template).
var LayoutTemplate = icon.Lucide(viewBox, paint, `<rect width="18" height="7" x="3" y="3" rx="1"/><rect width="9" height="7" x="3" y="14" rx="1"/><rect width="5" height="7" x="16" y="14" rx="1"/>`)

// Leaf is [leaf](https://lucide.dev/icons/leaf).
var Leaf = icon.Lucide(viewBox, paint, `<path d="M11 20a10 10 0 0010-10 25.9 25.9 0 00-1.04-7.281 1 1 0 00-1.755-.325C15.833 5.5 13 5.5 9.8 6.1A7 7 0 0011 20"/><path d="M2 21a5 5 0 012.911-4.544C7.613 15.212 8.351 15.24 11 13"/>`)

// LeafyGreen is [leafy-green](https://lucide.dev/icons/leafy-green).
var LeafyGreen = icon.Lucide(viewBox, paint, `<path d="M2 22c1.25-.987 2.27-1.975 3.9-2.2a5.56 5.56 0 0 1 3.8 1.5 4 4 0 0 0 6.187-2.353 3.5 3.5 0 0 0 3.69-5.116A3.5 3.5 0 0 0 20.95 8 3.5 3.5 0 1 0 16 3.05a3.5 3.5 0 0 0-5.831 1.373 3.5 3.5 0 0 0-5.116 3.69 4 4 0 0 0-2.348 6.155C3.499 15.42 4.409 16.712 4.2 18.1 3.926 19.743 3.014 20.732 2 22"/><path d="M2 22 17 7"/>`)

// Lectern is [lectern](https://lucide.dev/icons/lectern).
var Lectern = icon.Lucide(viewBox, paint, `<path d="M15 13h4a2 2 0 001.901-1.38l1.057-4.333A1 1 0 0021 6H3a1 1 0 00-.958 1.287L3.1 11.621A2 2 0 005.001 13h4"/><path d="M15 22V11a1 1 0 00-1-1h-4a1 1 0 00-1 1v11"/><path d="M18 22H6"/><path d="M18 6V3a1 1 0 00-1-1h-3"/>`)

// LensConcave is [lens-concave](https://lucide.dev/icons/lens-concave).
var LensConcave = icon.Lucide(viewBox, paint, `<path d="M7 2a1 1 0 0 0-.8 1.6 14 14 0 0 1 0 16.8A1 1 0 0 0 7 22h10a1 1 0 0 0 .8-1.6 14 14 0 0 1 0-16.8A1 1 0 0 0 17 2z"/>`)

// LensConvex is [lens-convex](https://lucide.dev/icons/lens-convex).
var LensConvex = icon.Lucide(viewBox, paint, `<path d="M13.433 2a1 1 0 0 1 .824.448 18 18 0 0 1 0 19.104 1 1 0 0 1-.824.448h-2.866a1 1 0 0 1-.824-.448 18 18 0 0 1 0-19.104A1 1 0 0 1 10.567 2z"/>`)

// Lens is [lens](https://lucide.dev/icons/lens).
var Lens = icon.Lucide(viewBox, paint, `<path d="M15.725 7.296h.01"/><path d="M6 12a6 6 0 016-6"/><circle cx="12" cy="12" r="10"/>`)

// LetterSpacing is [letter-spacing](https://lucide.dev/icons/letter-spacing).
var LetterSpacing = icon.Lucide(viewBox, paint, `<path d="M15 14H9"/><path d="M21 3v18"/><path d="M3 3v18"/><path d="m8 16 3.453-7.648a.6.6 0 011.094.001L16 16"/>`)

// Letters is [letters](https://lucide.dev/icons/letters).
var Letters = icon.Lucide(viewBox, paint, `<path d="M15 8H9"/><path d="M21 15.354a4 4 0 100 5.292"/><path d="M3 18h4a2 2 0 010 4H3.5a.5.5 0 01-.5-.5v-7a.5.5 0 01.5-.5H6a2 2 0 010 4"/><path d="m8 10 3.453-7.648a.6.6 0 011.094 0L16 10"/>`)

// LibraryBig is [library-big](https://lucide.dev/icons/library-big).
var LibraryBig = icon.Lucide(viewBox, paint, `<rect width="8" height="18" x="3" y="3" rx="1"/><path d="M7 3v18"/><path d="M20.4 18.9c.2.5-.1 1.1-.6 1.3l-1.9.7c-.5.2-1.1-.1-1.3-.6L11.1 5.1c-.2-.5.1-1.1.6-1.3l1.9-.7c.5-.2 1.1.1 1.3.6Z"/>`)

// Library is [library](https://lucide.dev/icons/library).
var Library = icon.Lucide(viewBox, paint, `<path d="m16 6 4 14"/><path d="M12 6v14"/><path d="M8 8v12"/><path d="M4 4v16"/>`)

// LifeBuoy is [life-buoy](https://lucide.dev/icons/life-buoy).
var LifeBuoy = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><path d="m4.93 4.93 4.24 4.24"/><path d="m14.83 9.17 4.24-4.24"/><path d="m14.83 14.83 4.24 4.24"/><path d="m9.17 14.83-4.24 4.24"/><circle cx="12" cy="12" r="4"/>`)

// Ligature is [ligature](https://lucide.dev/icons/ligature).
var Ligature = icon.Lucide(viewBox, paint, `<path d="M14 12h2v8"/><path d="M14 20h4"/><path d="M6 12h4"/><path d="M6 20h4"/><path d="M8 20V8a4 4 0 0 1 7.464-2"/>`)

// LightbulbOff is [lightbulb-off](https://lucide.dev/icons/lightbulb-off).
var LightbulbOff = icon.Lucide(viewBox, paint, `<path d="M16.8 11.2c.8-.9 1.2-2 1.2-3.2a6 6 0 0 0-9.3-5"/><path d="m2 2 20 20"/><path d="M6.3 6.3a4.67 4.67 0 0 0 1.2 5.2c.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/>`)

// Lightbulb is [lightbulb](https://lucide.dev/icons/lightbulb).
var Lightbulb = icon.Lucide(viewBox, paint, `<path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"/><path d="M9 18h6"/><path d="M10 22h4"/>`)

// Lighthouse is [lighthouse](https://lucide.dev/icons/lighthouse).
var Lighthouse = icon.Lucide(viewBox, paint, `<path d="M12 3V2"/><path d="M16.066 16.865 7 22l2-11V6a3 3 0 016 0v5l2 11"/><path d="m19.792 4.5.866-.5"/><path d="m19.797 13.5.866.5"/><path d="M21 9h1"/><path d="M3 9H2"/><path d="m4.203 13.5-.866.5"/><path d="M4.208 4.5 3.342 4"/><path d="M5.5 22h13"/><path d="m7.932 16.875 7.377-4.178"/><path d="M8 11h8"/><path d="M8 7h8"/>`)

// LineDotBottomVertical is [line-dot-bottom-vertical](https://lucide.dev/icons/line-dot-bottom-vertical).
var LineDotBottomVertical = icon.Lucide(viewBox, paint, `<path d="M12 3v12"/><circle cx="12" cy="18" r="3"/>`)

// LineDotLeftHorizontal is [line-dot-left-horizontal](https://lucide.dev/icons/line-dot-left-horizontal).
var LineDotLeftHorizontal = icon.Lucide(viewBox, paint, `<path d="M9 12h12"/><circle cx="6" cy="12" r="3"/>`)

// LineDotRightHorizontal is [line-dot-right-horizontal](https://lucide.dev/icons/line-dot-right-horizontal).
var LineDotRightHorizontal = icon.Lucide(viewBox, paint, `<path d="M 3 12 L 15 12"/><circle cx="18" cy="12" r="3"/>`)

// LineDotTopVertical is [line-dot-top-vertical](https://lucide.dev/icons/line-dot-top-vertical).
var LineDotTopVertical = icon.Lucide(viewBox, paint, `<path d="M12 9v12"/><circle cx="12" cy="6" r="3"/>`)

// LineHeight is [line-height](https://lucide.dev/icons/line-height).
var LineHeight = icon.Lucide(viewBox, paint, `<path d="M15 14H9"/><path d="M21 21H3"/><path d="M21 3H3"/><path d="m8 16 3.453-7.648a.6.6 0 011.094.001L16 16"/>`)

// LineSquiggle is [line-squiggle](https://lucide.dev/icons/line-squiggle).
var LineSquiggle = icon.Lucide(viewBox, paint, `<path d="M7 3.5c5-2 7 2.5 3 4C1.5 10 2 15 5 16c5 2 9-10 14-7s.5 13.5-4 12c-5-2.5.5-11 6-2"/>`)

// LineStyle is [line-style](https://lucide.dev/icons/line-style).
var LineStyle = icon.Lucide(viewBox, paint, `<path d="M11 5h2"/><path d="M15 12h6"/><path d="M19 5h2"/><path d="M3 12h6"/><path d="M3 19h18"/><path d="M3 5h2"/>`)

// Link2Off is [link-2-off](https://lucide.dev/icons/link-2-off).
var Link2Off = icon.Lucide(viewBox, paint, `<path d="M9 17H7A5 5 0 0 1 7 7"/><path d="M15 7h2a5 5 0 0 1 4 8"/><line x1="8" x2="12" y1="12" y2="12"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Link2 is [link-2](https://lucide.dev/icons/link-2).
var Link2 = icon.Lucide(viewBox, paint, `<path d="M9 17H7A5 5 0 0 1 7 7h2"/><path d="M15 7h2a5 5 0 1 1 0 10h-2"/><line x1="8" x2="16" y1="12" y2="12"/>`)

// Link is [link](https://lucide.dev/icons/link).
var Link = icon.Lucide(viewBox, paint, `<path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>`)

// ListCheck is [list-check](https://lucide.dev/icons/list-check).
var ListCheck = icon.Lucide(viewBox, paint, `<path d="M16 5H3"/><path d="M16 12H3"/><path d="M11 19H3"/><path d="m15 18 2 2 4-4"/>`)

// ListChecks is [list-checks](https://lucide.dev/icons/list-checks).
var ListChecks = icon.Lucide(viewBox, paint, `<path d="M13 5h8"/><path d="M13 12h8"/><path d="M13 19h8"/><path d="m3 17 2 2 4-4"/><path d="m3 7 2 2 4-4"/>`)

// ListChevronsDownUp is [list-chevrons-down-up](https://lucide.dev/icons/list-chevrons-down-up).
var ListChevronsDownUp = icon.Lucide(viewBox, paint, `<path d="M3 5h8"/><path d="M3 12h8"/><path d="M3 19h8"/><path d="m15 5 3 3 3-3"/><path d="m15 19 3-3 3 3"/>`)

// ListChevronsUpDown is [list-chevrons-up-down](https://lucide.dev/icons/list-chevrons-up-down).
var ListChevronsUpDown = icon.Lucide(viewBox, paint, `<path d="M3 5h8"/><path d="M3 12h8"/><path d="M3 19h8"/><path d="m15 8 3-3 3 3"/><path d="m15 16 3 3 3-3"/>`)

// ListClock is [list-clock](https://lucide.dev/icons/list-clock).
var ListClock = icon.Lucide(viewBox, paint, `<path d="M16 13v2.2l1.6 1"/><path d="M3 12h3.458"/><path d="M3 19h3.832"/><path d="M3 5h18"/><circle cx="16" cy="15" r="6"/>`)

// ListCollapse is [list-collapse](https://lucide.dev/icons/list-collapse).
var ListCollapse = icon.Lucide(viewBox, paint, `<path d="M10 5h11"/><path d="M10 12h11"/><path d="M10 19h11"/><path d="m3 10 3-3-3-3"/><path d="m3 20 3-3-3-3"/>`)

// ListEnd is [list-end](https://lucide.dev/icons/list-end).
var ListEnd = icon.Lucide(viewBox, paint, `<path d="M16 5H3"/><path d="M16 12H3"/><path d="M9 19H3"/><path d="m16 16-3 3 3 3"/><path d="M21 5v12a2 2 0 0 1-2 2h-6"/>`)

// ListFilterPlus is [list-filter-plus](https://lucide.dev/icons/list-filter-plus).
var ListFilterPlus = icon.Lucide(viewBox, paint, `<path d="M12 5H2"/><path d="M6 12h12"/><path d="M9 19h6"/><path d="M16 5h6"/><path d="M19 8V2"/>`)

// ListFilter is [list-filter](https://lucide.dev/icons/list-filter).
var ListFilter = icon.Lucide(viewBox, paint, `<path d="M2 5h20"/><path d="M6 12h12"/><path d="M9 19h6"/>`)

// ListIndentDecrease is [list-indent-decrease](https://lucide.dev/icons/list-indent-decrease).
var ListIndentDecrease = icon.Lucide(viewBox, paint, `<path d="M21 5H11"/><path d="M21 12H11"/><path d="M21 19H11"/><path d="m7 8-4 4 4 4"/>`)

// ListIndentIncrease is [list-indent-increase](https://lucide.dev/icons/list-indent-increase).
var ListIndentIncrease = icon.Lucide(viewBox, paint, `<path d="M21 5H11"/><path d="M21 12H11"/><path d="M21 19H11"/><path d="m3 8 4 4-4 4"/>`)

// ListMinus is [list-minus](https://lucide.dev/icons/list-minus).
var ListMinus = icon.Lucide(viewBox, paint, `<path d="M16 5H3"/><path d="M11 12H3"/><path d="M16 19H3"/><path d="M21 12h-6"/>`)

// ListMusic is [list-music](https://lucide.dev/icons/list-music).
var ListMusic = icon.Lucide(viewBox, paint, `<path d="M16 5H3"/><path d="M11 12H3"/><path d="M11 19H3"/><path d="M21 16V5"/><circle cx="18" cy="16" r="3"/>`)

// ListOrdered is [list-ordered](https://lucide.dev/icons/list-ordered).
var ListOrdered = icon.Lucide(viewBox, paint, `<path d="M11 5h10"/><path d="M11 12h10"/><path d="M11 19h10"/><path d="M4 4h1v5"/><path d="M4 9h2"/><path d="M6.5 20H3.4c0-1 2.6-1.925 2.6-3.5a1.5 1.5 0 0 0-2.6-1.02"/>`)

// ListPlus is [list-plus](https://lucide.dev/icons/list-plus).
var ListPlus = icon.Lucide(viewBox, paint, `<path d="M16 5H3"/><path d="M11 12H3"/><path d="M16 19H3"/><path d="M18 9v6"/><path d="M21 12h-6"/>`)

// ListRestart is [list-restart](https://lucide.dev/icons/list-restart).
var ListRestart = icon.Lucide(viewBox, paint, `<path d="M21 5H3"/><path d="M7 12H3"/><path d="M7 19H3"/><path d="M12 18a5 5 0 0 0 9-3 4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L11 14"/><path d="M11 10v4h4"/>`)

// ListSortAscending is [list-sort-ascending](https://lucide.dev/icons/list-sort-ascending).
var ListSortAscending = icon.Lucide(viewBox, paint, `<path d="M3 19h18"/><path d="M15 12H3"/><path d="M9 5H3"/>`)

// ListSortDescending is [list-sort-descending](https://lucide.dev/icons/list-sort-descending).
var ListSortDescending = icon.Lucide(viewBox, paint, `<path d="M15 12H3"/><path d="M3 5h18"/><path d="M9 19H3"/>`)

// ListStart is [list-start](https://lucide.dev/icons/list-start).
var ListStart = icon.Lucide(viewBox, paint, `<path d="M3 5h6"/><path d="M3 12h13"/><path d="M3 19h13"/><path d="m16 8-3-3 3-3"/><path d="M21 19V7a2 2 0 0 0-2-2h-6"/>`)

// ListTodo is [list-todo](https://lucide.dev/icons/list-todo).
var ListTodo = icon.Lucide(viewBox, paint, `<path d="M13 5h8"/><path d="M13 12h8"/><path d="M13 19h8"/><path d="m3 17 2 2 4-4"/><rect x="3" y="4" width="6" height="6" rx="1"/>`)

// ListTree is [list-tree](https://lucide.dev/icons/list-tree).
var ListTree = icon.Lucide(viewBox, paint, `<path d="M8 5h13"/><path d="M13 12h8"/><path d="M13 19h8"/><path d="M3 10a2 2 0 0 0 2 2h3"/><path d="M3 5v12a2 2 0 0 0 2 2h3"/>`)

// ListVideo is [list-video](https://lucide.dev/icons/list-video).
var ListVideo = icon.Lucide(viewBox, paint, `<path d="M21 5H3"/><path d="M10 12H3"/><path d="M10 19H3"/><path d="M15 12.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z"/>`)

// ListX is [list-x](https://lucide.dev/icons/list-x).
var ListX = icon.Lucide(viewBox, paint, `<path d="M16 5H3"/><path d="M11 12H3"/><path d="M16 19H3"/><path d="m15.5 9.5 5 5"/><path d="m20.5 9.5-5 5"/>`)

// List is [list](https://lucide.dev/icons/list).
var List = icon.Lucide(viewBox, paint, `<path d="M3 5h.01"/><path d="M3 12h.01"/><path d="M3 19h.01"/><path d="M8 5h13"/><path d="M8 12h13"/><path d="M8 19h13"/>`)

// LoaderCircle is [loader-circle](https://lucide.dev/icons/loader-circle).
var LoaderCircle = icon.Lucide(viewBox, paint, `<path d="M21 12a9 9 0 1 1-6.219-8.56"/>`)

// LoaderPinwheel is [loader-pinwheel](https://lucide.dev/icons/loader-pinwheel).
var LoaderPinwheel = icon.Lucide(viewBox, paint, `<path d="M22 12a1 1 0 0 1-10 0 1 1 0 0 0-10 0"/><path d="M7 20.7a1 1 0 1 1 5-8.7 1 1 0 1 0 5-8.6"/><path d="M7 3.3a1 1 0 1 1 5 8.6 1 1 0 1 0 5 8.6"/><circle cx="12" cy="12" r="10"/>`)

// Loader is [loader](https://lucide.dev/icons/loader).
var Loader = icon.Lucide(viewBox, paint, `<path d="M12 2v4"/><path d="m16.2 7.8 2.9-2.9"/><path d="M18 12h4"/><path d="m16.2 16.2 2.9 2.9"/><path d="M12 18v4"/><path d="m4.9 19.1 2.9-2.9"/><path d="M2 12h4"/><path d="m4.9 4.9 2.9 2.9"/>`)

// LocateFixed is [locate-fixed](https://lucide.dev/icons/locate-fixed).
var LocateFixed = icon.Lucide(viewBox, paint, `<line x1="2" x2="5" y1="12" y2="12"/><line x1="19" x2="22" y1="12" y2="12"/><line x1="12" x2="12" y1="2" y2="5"/><line x1="12" x2="12" y1="19" y2="22"/><circle cx="12" cy="12" r="7"/><circle cx="12" cy="12" r="3"/>`)

// LocateOff is [locate-off](https://lucide.dev/icons/locate-off).
var LocateOff = icon.Lucide(viewBox, paint, `<path d="M12 19v3"/><path d="M12 2v3"/><path d="M18.89 13.24a7 7 0 0 0-8.13-8.13"/><path d="M19 12h3"/><path d="M2 12h3"/><path d="m2 2 20 20"/><path d="M7.05 7.05a7 7 0 0 0 9.9 9.9"/>`)

// Locate is [locate](https://lucide.dev/icons/locate).
var Locate = icon.Lucide(viewBox, paint, `<line x1="2" x2="5" y1="12" y2="12"/><line x1="19" x2="22" y1="12" y2="12"/><line x1="12" x2="12" y1="2" y2="5"/><line x1="12" x2="12" y1="19" y2="22"/><circle cx="12" cy="12" r="7"/>`)

// LockKeyholeOpen is [lock-keyhole-open](https://lucide.dev/icons/lock-keyhole-open).
var LockKeyholeOpen = icon.Lucide(viewBox, paint, `<circle cx="12" cy="16" r="1"/><rect width="18" height="12" x="3" y="10" rx="2"/><path d="M7 10V7a5 5 0 0 1 9.33-2.5"/>`)

// LockKeyhole is [lock-keyhole](https://lucide.dev/icons/lock-keyhole).
var LockKeyhole = icon.Lucide(viewBox, paint, `<circle cx="12" cy="16" r="1"/><rect x="3" y="10" width="18" height="12" rx="2"/><path d="M7 10V7a5 5 0 0 1 10 0v3"/>`)

// LockOpen is [lock-open](https://lucide.dev/icons/lock-open).
var LockOpen = icon.Lucide(viewBox, paint, `<rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 9.9-1"/>`)

// Lock is [lock](https://lucide.dev/icons/lock).
var Lock = icon.Lucide(viewBox, paint, `<rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/>`)

// LogIn is [log-in](https://lucide.dev/icons/log-in).
var LogIn = icon.Lucide(viewBox, paint, `<path d="m10 17 5-5-5-5"/><path d="M15 12H3"/><path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"/>`)

// LogOut is [log-out](https://lucide.dev/icons/log-out).
var LogOut = icon.Lucide(viewBox, paint, `<path d="m16 17 5-5-5-5"/><path d="M21 12H9"/><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>`)

// Logs is [logs](https://lucide.dev/icons/logs).
var Logs = icon.Lucide(viewBox, paint, `<path d="M3 5h1"/><path d="M3 12h1"/><path d="M3 19h1"/><path d="M8 5h1"/><path d="M8 12h1"/><path d="M8 19h1"/><path d="M13 5h8"/><path d="M13 12h8"/><path d="M13 19h8"/>`)

// Lollipop is [lollipop](https://lucide.dev/icons/lollipop).
var Lollipop = icon.Lucide(viewBox, paint, `<circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/><path d="M11 11a2 2 0 0 0 4 0 4 4 0 0 0-8 0 6 6 0 0 0 12 0"/>`)

// Luggage is [luggage](https://lucide.dev/icons/luggage).
var Luggage = icon.Lucide(viewBox, paint, `<path d="M6 20a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2"/><path d="M8 18V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v14"/><path d="M10 20h4"/><circle cx="16" cy="20" r="2"/><circle cx="8" cy="20" r="2"/>`)

// Magnet is [magnet](https://lucide.dev/icons/magnet).
var Magnet = icon.Lucide(viewBox, paint, `<path d="m12 15 4 4"/><path d="M2.352 10.648a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l6.029-6.029a1 1 0 1 1 3 3l-6.029 6.029a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l6.365-6.367A1 1 0 0 0 8.716 4.282z"/><path d="m5 8 4 4"/>`)

// MailBadge is [mail-badge](https://lucide.dev/icons/mail-badge).
var MailBadge = icon.Lucide(viewBox, paint, `<path d="M22 7.7V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8.25"/><path d="M12 12.996a1.94 1.94 0 0 1-1.03-.296L2 7"/><path d="m20.69 16.479 1.29 4.88a.5.5 0 0 1-.698.591l-1.843-.849a1 1 0 0 0-.879.001l-1.846.85a.5.5 0 0 1-.692-.593l1.29-4.88"/><circle cx="19" cy="14" r="3"/>`)

// MailCheck is [mail-check](https://lucide.dev/icons/mail-check).
var MailCheck = icon.Lucide(viewBox, paint, `<path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/><path d="m16 19 2 2 4-4"/>`)

// MailClock is [mail-clock](https://lucide.dev/icons/mail-clock).
var MailClock = icon.Lucide(viewBox, paint, `<path d="M16 14v2.2l1.6 1"/><path d="m22 7-.759.484"/><path d="M6.835 20H4a2 2 0 01-2-2V6a2 2 0 012-2h16a2 2 0 012 2v2"/><path d="M7.605 10.567 2 7"/><circle cx="16" cy="16" r="6"/>`)

// MailDot is [mail-dot](https://lucide.dev/icons/mail-dot).
var MailDot = icon.Lucide(viewBox, paint, `<path d="M12.674 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4.675"/><path d="m2 7 9 5.701a2 2 0 0 0 2 0l1.084-.686"/><circle cx="19" cy="7" r="3"/>`)

// MailMinus is [mail-minus](https://lucide.dev/icons/mail-minus).
var MailMinus = icon.Lucide(viewBox, paint, `<path d="M22 15V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/><path d="M16 19h6"/>`)

// MailOpen is [mail-open](https://lucide.dev/icons/mail-open).
var MailOpen = icon.Lucide(viewBox, paint, `<path d="M21.2 8.4c.5.38.8.97.8 1.6v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V10a2 2 0 0 1 .8-1.6l8-6a2 2 0 0 1 2.4 0l8 6Z"/><path d="m22 10-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 10"/>`)

// MailPen is [mail-pen](https://lucide.dev/icons/mail-pen).
var MailPen = icon.Lucide(viewBox, paint, `<path d="M15.363 17.634a2 2 0 00-.506.854l-.837 2.87a.5.5 0 00.62.62l2.87-.837a2 2 0 00.854-.506l3.013-3.009a1 1 0 10-3.004-3.004z"/><path d="M22 10.38V6a2 2 0 00-2-2H4a2 2 0 00-2 2v12a2 2 0 002 2h6.25"/><path d="m22 7-8.991 5.727a2 2 0 01-2.009 0L2 7"/>`)

// MailPlus is [mail-plus](https://lucide.dev/icons/mail-plus).
var MailPlus = icon.Lucide(viewBox, paint, `<path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/><path d="M19 16v6"/><path d="M16 19h6"/>`)

// MailQuestionMark is [mail-question-mark](https://lucide.dev/icons/mail-question-mark).
var MailQuestionMark = icon.Lucide(viewBox, paint, `<path d="M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/><path d="M18 15.28c.2-.4.5-.8.9-1a2.1 2.1 0 0 1 2.6.4c.3.4.5.8.5 1.3 0 1.3-2 2-2 2"/><path d="M20 22v.01"/>`)

// MailSearch is [mail-search](https://lucide.dev/icons/mail-search).
var MailSearch = icon.Lucide(viewBox, paint, `<path d="M22 12.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h7.5"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/><path d="M18 21a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z"/><circle cx="18" cy="18" r="3"/><path d="m22 22-1.5-1.5"/>`)

// MailWarning is [mail-warning](https://lucide.dev/icons/mail-warning).
var MailWarning = icon.Lucide(viewBox, paint, `<path d="M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"/><path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/><path d="M20 14v4"/><path d="M20 22v.01"/>`)

// MailX is [mail-x](https://lucide.dev/icons/mail-x).
var MailX = icon.Lucide(viewBox, paint, `<path d="M22 12.532V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8.792"/><path d="m22 7-8.991 5.727a2 2 0 0 1-2.009 0L2 7"/><path d="m16.5 16.5 5 5"/><path d="m21.5 16.5-5 5"/>`)

// Mail is [mail](https://lucide.dev/icons/mail).
var Mail = icon.Lucide(viewBox, paint, `<path d="m22 7-8.991 5.727a2 2 0 0 1-2.009 0L2 7"/><rect x="2" y="4" width="20" height="16" rx="2"/>`)

// Mailbox is [mailbox](https://lucide.dev/icons/mailbox).
var Mailbox = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9.5C2 7 4 5 6.5 5H18c2.2 0 4 1.8 4 4v8Z"/><polyline points="15,9 18,9 18,11"/><path d="M6.5 5C9 5 11 7 11 9.5V17a2 2 0 0 1-2 2"/><line x1="6" x2="7" y1="10" y2="10"/>`)

// Mails is [mails](https://lucide.dev/icons/mails).
var Mails = icon.Lucide(viewBox, paint, `<path d="M17 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 1-1.732"/><path d="m22 5.5-6.419 4.179a2 2 0 0 1-2.162 0L7 5.5"/><rect x="7" y="3" width="15" height="12" rx="2"/>`)

// MapMinus is [map-minus](https://lucide.dev/icons/map-minus).
var MapMinus = icon.Lucide(viewBox, paint, `<path d="m11 19-1.106-.552a2 2 0 0 0-1.788 0l-3.659 1.83A1 1 0 0 1 3 19.381V6.618a1 1 0 0 1 .553-.894l4.553-2.277a2 2 0 0 1 1.788 0l4.212 2.106a2 2 0 0 0 1.788 0l3.659-1.83A1 1 0 0 1 21 4.619V14"/><path d="M15 5.764V14"/><path d="M21 18h-6"/><path d="M9 3.236v15"/>`)

// MapPinCheckInside is [map-pin-check-inside](https://lucide.dev/icons/map-pin-check-inside).
var MapPinCheckInside = icon.Lucide(viewBox, paint, `<path d="M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"/><path d="m9 10 2 2 4-4"/>`)

// MapPinCheck is [map-pin-check](https://lucide.dev/icons/map-pin-check).
var MapPinCheck = icon.Lucide(viewBox, paint, `<path d="M19.43 12.935c.357-.967.57-1.955.57-2.935a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32.197 32.197 0 0 0 .813-.728"/><circle cx="12" cy="10" r="3"/><path d="m16 18 2 2 4-4"/>`)

// MapPinHouse is [map-pin-house](https://lucide.dev/icons/map-pin-house).
var MapPinHouse = icon.Lucide(viewBox, paint, `<path d="M15 22a1 1 0 0 1-1-1v-4a1 1 0 0 1 .445-.832l3-2a1 1 0 0 1 1.11 0l3 2A1 1 0 0 1 22 17v4a1 1 0 0 1-1 1z"/><path d="M18 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 .601.2"/><path d="M18 22v-3"/><circle cx="10" cy="10" r="3"/>`)

// MapPinMinusInside is [map-pin-minus-inside](https://lucide.dev/icons/map-pin-minus-inside).
var MapPinMinusInside = icon.Lucide(viewBox, paint, `<path d="M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"/><path d="M9 10h6"/>`)

// MapPinMinus is [map-pin-minus](https://lucide.dev/icons/map-pin-minus).
var MapPinMinus = icon.Lucide(viewBox, paint, `<path d="M18.977 14C19.6 12.701 20 11.343 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32 32 0 0 0 .824-.738"/><circle cx="12" cy="10" r="3"/><path d="M16 18h6"/>`)

// MapPinOff is [map-pin-off](https://lucide.dev/icons/map-pin-off).
var MapPinOff = icon.Lucide(viewBox, paint, `<path d="M12.75 7.09a3 3 0 0 1 2.16 2.16"/><path d="M17.072 17.072c-1.634 2.17-3.527 3.912-4.471 4.727a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 1.432-4.568"/><path d="m2 2 20 20"/><path d="M8.475 2.818A8 8 0 0 1 20 10c0 1.183-.31 2.377-.81 3.533"/><path d="M9.13 9.13a3 3 0 0 0 3.74 3.74"/>`)

// MapPinPen is [map-pin-pen](https://lucide.dev/icons/map-pin-pen).
var MapPinPen = icon.Lucide(viewBox, paint, `<path d="M17.97 9.304A8 8 0 0 0 2 10c0 4.69 4.887 9.562 7.022 11.468"/><path d="M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/><circle cx="10" cy="10" r="3"/>`)

// MapPinPlusInside is [map-pin-plus-inside](https://lucide.dev/icons/map-pin-plus-inside).
var MapPinPlusInside = icon.Lucide(viewBox, paint, `<path d="M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"/><path d="M12 7v6"/><path d="M9 10h6"/>`)

// MapPinPlus is [map-pin-plus](https://lucide.dev/icons/map-pin-plus).
var MapPinPlus = icon.Lucide(viewBox, paint, `<path d="M19.914 11.105A7.298 7.298 0 0 0 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32 32 0 0 0 .824-.738"/><circle cx="12" cy="10" r="3"/><path d="M16 18h6"/><path d="M19 15v6"/>`)

// MapPinSearch is [map-pin-search](https://lucide.dev/icons/map-pin-search).
var MapPinSearch = icon.Lucide(viewBox, paint, `<path d="M 12.248 21.969 a 1 1 0 0 1 -0.849 -0.17 C 9.539 20.193 4 14.993 4 10 a 8 8 0 0 1 16 0 C 20 10.42 19.961 10.841 19.888 11.262"/><path d="m22 22-1.88-1.88"/><circle cx="12" cy="10" r="3"/><circle cx="18" cy="18" r="3"/>`)

// MapPinXInside is [map-pin-x-inside](https://lucide.dev/icons/map-pin-x-inside).
var MapPinXInside = icon.Lucide(viewBox, paint, `<path d="M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"/><path d="m14.5 7.5-5 5"/><path d="m9.5 7.5 5 5"/>`)

// MapPinX is [map-pin-x](https://lucide.dev/icons/map-pin-x).
var MapPinX = icon.Lucide(viewBox, paint, `<path d="M19.752 11.901A7.78 7.78 0 0 0 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 19 19 0 0 0 .09-.077"/><circle cx="12" cy="10" r="3"/><path d="m21.5 15.5-5 5"/><path d="m21.5 20.5-5-5"/>`)

// MapPin is [map-pin](https://lucide.dev/icons/map-pin).
var MapPin = icon.Lucide(viewBox, paint, `<path d="M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"/><circle cx="12" cy="10" r="3"/>`)

// MapPinned is [map-pinned](https://lucide.dev/icons/map-pinned).
var MapPinned = icon.Lucide(viewBox, paint, `<path d="M18 8c0 3.613-3.869 7.429-5.393 8.795a1 1 0 01-1.214 0C9.87 15.429 6 11.613 6 8a6 6 0 0112 0"/><path d="M4.474 15h-.197a1 1 0 00-.969.753l-1.097 4.35a1.5 1.5 0 001.444 1.898L20.344 22a1.5 1.5 0 001.446-1.897l-1.098-4.35a1 1 0 00-.969-.753h-.197"/><circle cx="12" cy="8" r="2"/>`)

// MapPlus is [map-plus](https://lucide.dev/icons/map-plus).
var MapPlus = icon.Lucide(viewBox, paint, `<path d="m11 19-1.106-.552a2 2 0 0 0-1.788 0l-3.659 1.83A1 1 0 0 1 3 19.381V6.618a1 1 0 0 1 .553-.894l4.553-2.277a2 2 0 0 1 1.788 0l4.212 2.106a2 2 0 0 0 1.788 0l3.659-1.83A1 1 0 0 1 21 4.619V12"/><path d="M15 5.764V12"/><path d="M18 15v6"/><path d="M21 18h-6"/><path d="M9 3.236v15"/>`)

// Map is [map](https://lucide.dev/icons/map).
var Map = icon.Lucide(viewBox, paint, `<path d="M14.106 5.553a2 2 0 0 0 1.788 0l3.659-1.83A1 1 0 0 1 21 4.619v12.764a1 1 0 0 1-.553.894l-4.553 2.277a2 2 0 0 1-1.788 0l-4.212-2.106a2 2 0 0 0-1.788 0l-3.659 1.83A1 1 0 0 1 3 19.381V6.618a1 1 0 0 1 .553-.894l4.553-2.277a2 2 0 0 1 1.788 0z"/><path d="M15 5.764v15"/><path d="M9 3.236v15"/>`)

// MarsStroke is [mars-stroke](https://lucide.dev/icons/mars-stroke).
var MarsStroke = icon.Lucide(viewBox, paint, `<path d="m14 6 4 4"/><path d="M17 3h4v4"/><path d="m21 3-7.75 7.75"/><circle cx="9" cy="15" r="6"/>`)

// Mars is [mars](https://lucide.dev/icons/mars).
var Mars = icon.Lucide(viewBox, paint, `<path d="M16 3h5v5"/><path d="m21 3-6.75 6.75"/><circle cx="10" cy="14" r="6"/>`)

// Martini is [martini](https://lucide.dev/icons/martini).
var Martini = icon.Lucide(viewBox, paint, `<path d="M12 12 4.207 4.207A.707.707 0 0 1 4.707 3h14.586a.707.707 0 0 1 .5 1.207z"/><path d="M12 12v10"/><path d="M7 22h10"/>`)

// Maximize2 is [maximize-2](https://lucide.dev/icons/maximize-2).
var Maximize2 = icon.Lucide(viewBox, paint, `<path d="M15 3h6v6"/><path d="m21 3-7 7"/><path d="m3 21 7-7"/><path d="M9 21H3v-6"/>`)

// Maximize is [maximize](https://lucide.dev/icons/maximize).
var Maximize = icon.Lucide(viewBox, paint, `<path d="M8 3H5a2 2 0 0 0-2 2v3"/><path d="M21 8V5a2 2 0 0 0-2-2h-3"/><path d="M3 16v3a2 2 0 0 0 2 2h3"/><path d="M16 21h3a2 2 0 0 0 2-2v-3"/>`)

// Medal is [medal](https://lucide.dev/icons/medal).
var Medal = icon.Lucide(viewBox, paint, `<path d="M7.21 15 2.66 7.14a2 2 0 0 1 .13-2.2L4.4 2.8A2 2 0 0 1 6 2h12a2 2 0 0 1 1.6.8l1.6 2.14a2 2 0 0 1 .14 2.2L16.79 15"/><path d="M11 12 5.12 2.2"/><path d="m13 12 5.88-9.8"/><path d="M8 7h8"/><circle cx="12" cy="17" r="5"/><path d="M12 18v-2h-.5"/>`)

// MegaphoneOff is [megaphone-off](https://lucide.dev/icons/megaphone-off).
var MegaphoneOff = icon.Lucide(viewBox, paint, `<path d="M11.636 6A13 13 0 0 0 19.4 3.2 1 1 0 0 1 21 4v11.344"/><path d="M14.378 14.357A13 13 0 0 0 11 14H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1"/><path d="m2 2 20 20"/><path d="M6 14a12 12 0 0 0 2.4 7.2 2 2 0 0 0 3.2-2.4A8 8 0 0 1 10 14"/><path d="M8 8v6"/>`)

// Megaphone is [megaphone](https://lucide.dev/icons/megaphone).
var Megaphone = icon.Lucide(viewBox, paint, `<path d="M11 6a13 13 0 0 0 8.4-2.8A1 1 0 0 1 21 4v12a1 1 0 0 1-1.6.8A13 13 0 0 0 11 14H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2z"/><path d="M6 14a12 12 0 0 0 2.4 7.2 2 2 0 0 0 3.2-2.4A8 8 0 0 1 10 14"/><path d="M8 6v8"/>`)

// MemoryStick is [memory-stick](https://lucide.dev/icons/memory-stick).
var MemoryStick = icon.Lucide(viewBox, paint, `<path d="M12 12v-2"/><path d="M12 18v-2"/><path d="M16 12v-2"/><path d="M16 18v-2"/><path d="M2 11h1.5"/><path d="M20 18v-2"/><path d="M20.5 11H22"/><path d="M4 18v-2"/><path d="M8 12v-2"/><path d="M8 18v-2"/><rect x="2" y="6" width="20" height="10" rx="2"/>`)

// Menu is [menu](https://lucide.dev/icons/menu).
var Menu = icon.Lucide(viewBox, paint, `<path d="M4 5h16"/><path d="M4 12h16"/><path d="M4 19h16"/>`)

// Merge is [merge](https://lucide.dev/icons/merge).
var Merge = icon.Lucide(viewBox, paint, `<path d="m8 6 4-4 4 4"/><path d="M12 2v10.3a4 4 0 0 1-1.172 2.872L4 22"/><path d="m20 22-5-5"/>`)

// MessageCircleCheck is [message-circle-check](https://lucide.dev/icons/message-circle-check).
var MessageCircleCheck = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="m16 9-5.5 5.5L8 12"/>`)

// MessageCircleCode is [message-circle-code](https://lucide.dev/icons/message-circle-code).
var MessageCircleCode = icon.Lucide(viewBox, paint, `<path d="m10 9-3 3 3 3"/><path d="m14 15 3-3-3-3"/><path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/>`)

// MessageCircleDashedCheck is [message-circle-dashed-check](https://lucide.dev/icons/message-circle-dashed-check).
var MessageCircleDashedCheck = icon.Lucide(viewBox, paint, `<path d="M10.1 2.182a10 10 0 013.8 0"/><path d="M13.9 21.818a10 10 0 01-3.8 0"/><path d="M17.609 3.72a10 10 0 012.69 2.7"/><path d="M2.182 13.9a10 10 0 010-3.8"/><path d="M20.28 17.61a10 10 0 01-2.7 2.69"/><path d="M21.818 10.1a10 10 0 010 3.8"/><path d="M3.721 6.391a10 10 0 012.7-2.69"/><path d="m6.163 21.117-2.906.85a1 1 0 01-1.236-1.169l.965-2.98"/><path d="m16 9-5.5 5.5L8 12"/>`)

// MessageCircleDashed is [message-circle-dashed](https://lucide.dev/icons/message-circle-dashed).
var MessageCircleDashed = icon.Lucide(viewBox, paint, `<path d="M10.1 2.182a10 10 0 0 1 3.8 0"/><path d="M13.9 21.818a10 10 0 0 1-3.8 0"/><path d="M17.609 3.72a10 10 0 0 1 2.69 2.7"/><path d="M2.182 13.9a10 10 0 0 1 0-3.8"/><path d="M20.28 17.61a10 10 0 0 1-2.7 2.69"/><path d="M21.818 10.1a10 10 0 0 1 0 3.8"/><path d="M3.721 6.391a10 10 0 0 1 2.7-2.69"/><path d="m6.163 21.117-2.906.85a1 1 0 0 1-1.236-1.169l.965-2.98"/>`)

// MessageCircleHeart is [message-circle-heart](https://lucide.dev/icons/message-circle-heart).
var MessageCircleHeart = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="M7.828 13.07A3 3 0 0 1 12 8.764a3 3 0 0 1 5.004 2.224 3 3 0 0 1-.832 2.083l-3.447 3.62a1 1 0 0 1-1.45-.001z"/>`)

// MessageCircleMore is [message-circle-more](https://lucide.dev/icons/message-circle-more).
var MessageCircleMore = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="M8 12h.01"/><path d="M12 12h.01"/><path d="M16 12h.01"/>`)

// MessageCircleOff is [message-circle-off](https://lucide.dev/icons/message-circle-off).
var MessageCircleOff = icon.Lucide(viewBox, paint, `<path d="m2 2 20 20"/><path d="M4.93 4.929a10 10 0 0 0-1.938 11.412 2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 0 0 11.302-1.989"/><path d="M8.35 2.69A10 10 0 0 1 21.3 15.65"/>`)

// MessageCirclePlus is [message-circle-plus](https://lucide.dev/icons/message-circle-plus).
var MessageCirclePlus = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="M8 12h8"/><path d="M12 8v8"/>`)

// MessageCircleQuestionMark is [message-circle-question-mark](https://lucide.dev/icons/message-circle-question-mark).
var MessageCircleQuestionMark = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><path d="M12 17h.01"/>`)

// MessageCircleReply is [message-circle-reply](https://lucide.dev/icons/message-circle-reply).
var MessageCircleReply = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="m10 15-3-3 3-3"/><path d="M7 12h8a2 2 0 0 1 2 2v1"/>`)

// MessageCircleText is [message-circle-text](https://lucide.dev/icons/message-circle-text).
var MessageCircleText = icon.Lucide(viewBox, paint, `<path d="M3.086 17.509a2 2 0 00-.094-1.167 10 10 0 114.777 4.719 2 2 0 00-1.099-.092l-3.413.998a1 1 0 01-1.236-1.168z"/><path d="M8 12h9"/><path d="M8 16h5"/><path d="M8 8h7"/>`)

// MessageCircleWarning is [message-circle-warning](https://lucide.dev/icons/message-circle-warning).
var MessageCircleWarning = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="M12 8v4"/><path d="M12 16h.01"/>`)

// MessageCircleX is [message-circle-x](https://lucide.dev/icons/message-circle-x).
var MessageCircleX = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/><path d="m15 9-6 6"/><path d="m9 9 6 6"/>`)

// MessageCircle is [message-circle](https://lucide.dev/icons/message-circle).
var MessageCircle = icon.Lucide(viewBox, paint, `<path d="M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"/>`)

// MessageSquareCheck is [message-square-check](https://lucide.dev/icons/message-square-check).
var MessageSquareCheck = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.7.7 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="m9 11 2 2 4-4"/>`)

// MessageSquareCode is [message-square-code](https://lucide.dev/icons/message-square-code).
var MessageSquareCode = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="m10 8-3 3 3 3"/><path d="m14 14 3-3-3-3"/>`)

// MessageSquareDashed is [message-square-dashed](https://lucide.dev/icons/message-square-dashed).
var MessageSquareDashed = icon.Lucide(viewBox, paint, `<path d="M14 3h2"/><path d="M16 19h-2"/><path d="M2 12v-2"/><path d="M2 16v5.286a.71.71 0 0 0 1.212.502l1.149-1.149"/><path d="M20 19a2 2 0 0 0 2-2v-1"/><path d="M22 10v2"/><path d="M22 6V5a2 2 0 0 0-2-2"/><path d="M4 3a2 2 0 0 0-2 2v1"/><path d="M8 19h2"/><path d="M8 3h2"/>`)

// MessageSquareDiff is [message-square-diff](https://lucide.dev/icons/message-square-diff).
var MessageSquareDiff = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="M10 15h4"/><path d="M10 9h4"/><path d="M12 7v4"/>`)

// MessageSquareDot is [message-square-dot](https://lucide.dev/icons/message-square-dot).
var MessageSquareDot = icon.Lucide(viewBox, paint, `<path d="M12.7 3H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H20a2 2 0 0 0 2-2v-4.7"/><circle cx="19" cy="6" r="3"/>`)

// MessageSquareHeart is [message-square-heart](https://lucide.dev/icons/message-square-heart).
var MessageSquareHeart = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="M7.5 9.5c0 .687.265 1.383.697 1.844l3.009 3.264a1.14 1.14 0 0 0 .407.314 1 1 0 0 0 .783-.004 1.14 1.14 0 0 0 .398-.31l3.008-3.264A2.77 2.77 0 0 0 16.5 9.5 2.5 2.5 0 0 0 12 8a2.5 2.5 0 0 0-4.5 1.5"/>`)

// MessageSquareLock is [message-square-lock](https://lucide.dev/icons/message-square-lock).
var MessageSquareLock = icon.Lucide(viewBox, paint, `<path d="M22 8.5V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H10"/><path d="M20 15v-2a2 2 0 0 0-4 0v2"/><rect x="14" y="15" width="8" height="5" rx="1"/>`)

// MessageSquareMore is [message-square-more](https://lucide.dev/icons/message-square-more).
var MessageSquareMore = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="M12 11h.01"/><path d="M16 11h.01"/><path d="M8 11h.01"/>`)

// MessageSquareOff is [message-square-off](https://lucide.dev/icons/message-square-off).
var MessageSquareOff = icon.Lucide(viewBox, paint, `<path d="M19 19H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.7.7 0 0 1 2 21.286V5a2 2 0 0 1 1.184-1.826"/><path d="m2 2 20 20"/><path d="M8.656 3H20a2 2 0 0 1 2 2v11.344"/>`)

// MessageSquarePlus is [message-square-plus](https://lucide.dev/icons/message-square-plus).
var MessageSquarePlus = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="M12 8v6"/><path d="M9 11h6"/>`)

// MessageSquareQuote is [message-square-quote](https://lucide.dev/icons/message-square-quote).
var MessageSquareQuote = icon.Lucide(viewBox, paint, `<path d="M14 14a2 2 0 0 0 2-2V8h-2"/><path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="M8 14a2 2 0 0 0 2-2V8H8"/>`)

// MessageSquareReply is [message-square-reply](https://lucide.dev/icons/message-square-reply).
var MessageSquareReply = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="m10 8-3 3 3 3"/><path d="M17 14v-1a2 2 0 0 0-2-2H7"/>`)

// MessageSquareShare is [message-square-share](https://lucide.dev/icons/message-square-share).
var MessageSquareShare = icon.Lucide(viewBox, paint, `<path d="M12 3H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H20a2 2 0 0 0 2-2v-4"/><path d="M16 3h6v6"/><path d="m16 9 6-6"/>`)

// MessageSquareText is [message-square-text](https://lucide.dev/icons/message-square-text).
var MessageSquareText = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="M7 11h10"/><path d="M7 15h6"/><path d="M7 7h8"/>`)

// MessageSquareWarning is [message-square-warning](https://lucide.dev/icons/message-square-warning).
var MessageSquareWarning = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="M12 15h.01"/><path d="M12 7v4"/>`)

// MessageSquareX is [message-square-x](https://lucide.dev/icons/message-square-x).
var MessageSquareX = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/><path d="m14.5 8.5-5 5"/><path d="m9.5 8.5 5 5"/>`)

// MessageSquare is [message-square](https://lucide.dev/icons/message-square).
var MessageSquare = icon.Lucide(viewBox, paint, `<path d="M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"/>`)

// MessagesCircle is [messages-circle](https://lucide.dev/icons/messages-circle).
var MessagesCircle = icon.Lucide(viewBox, paint, `<path d="M19.95 10.05a7 7 0 011.412 7.872 1 1 0 00-.058.787l.675 2.089a1 1 0 01-1.236 1.168l-2.155-.631a1 1 0 00-.745.06 7 7 0 01-7.793-1.445"/><path d="M2.696 12.708a1 1 0 00-.058-.785 7 7 0 113.518 3.473 1 1 0 00-.744-.061l-2.155.63a1 1 0 01-1.236-1.167z"/>`)

// MessagesSquare is [messages-square](https://lucide.dev/icons/messages-square).
var MessagesSquare = icon.Lucide(viewBox, paint, `<path d="M16 10a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 14.286V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/><path d="M20 9a2 2 0 0 1 2 2v10.286a.71.71 0 0 1-1.212.502l-2.202-2.202A2 2 0 0 0 17.172 19H10a2 2 0 0 1-2-2v-1"/>`)

// Metronome is [metronome](https://lucide.dev/icons/metronome).
var Metronome = icon.Lucide(viewBox, paint, `<path d="M12 11.4V9.1"/><path d="m12 17 6.59-6.59"/><path d="m15.05 5.7-.218-.691a3 3 0 0 0-5.663 0L4.418 19.695A1 1 0 0 0 5.37 21h13.253a1 1 0 0 0 .951-1.31L18.45 16.2"/><circle cx="20" cy="9" r="2"/>`)

// MicAudioLines is [mic-audio-lines](https://lucide.dev/icons/mic-audio-lines).
var MicAudioLines = icon.Lucide(viewBox, paint, `<path d="M10 3v2.341"/><path d="M12 17v4"/><path d="M14 5v.341"/><path d="M18 5v13"/><path d="M2 10v3"/><path d="M22 10v3"/><path d="M6 6v11"/><path d="M9 21h6"/><rect width="4" height="8" x="10" y="9" rx="2"/>`)

// MicOff is [mic-off](https://lucide.dev/icons/mic-off).
var MicOff = icon.Lucide(viewBox, paint, `<path d="M12 19v3"/><path d="M15 9.34V5a3 3 0 0 0-5.68-1.33"/><path d="M16.95 16.95A7 7 0 0 1 5 12v-2"/><path d="M18.89 13.23A7 7 0 0 0 19 12v-2"/><path d="m2 2 20 20"/><path d="M9 9v3a3 3 0 0 0 5.12 2.12"/>`)

// MicSignal is [mic-signal](https://lucide.dev/icons/mic-signal).
var MicSignal = icon.Lucide(viewBox, paint, `<path d="M12 17v4"/><path d="M18 11a6 6 0 00-3-5.197"/><path d="M2 11a10 10 0 015-8.662"/><path d="M22 11a10 10 0 00-5-8.662"/><path d="M6 11a6 6 0 013-5.197"/><path d="M9 21h6"/><rect x="10" y="9" width="4" height="8" rx="2"/>`)

// MicVocal is [mic-vocal](https://lucide.dev/icons/mic-vocal).
var MicVocal = icon.Lucide(viewBox, paint, `<path d="m11 7.601-5.994 8.19a1 1 0 0 0 .1 1.298l.817.818a1 1 0 0 0 1.314.087L15.09 12"/><path d="M16.5 21.174C15.5 20.5 14.372 20 13 20c-2.058 0-3.928 2.356-6 2-2.072-.356-2.775-3.369-1.5-4.5"/><circle cx="16" cy="7" r="5"/>`)

// Mic is [mic](https://lucide.dev/icons/mic).
var Mic = icon.Lucide(viewBox, paint, `<path d="M12 19v3"/><path d="M19 10v2a7 7 0 0 1-14 0v-2"/><rect x="9" y="2" width="6" height="13" rx="3"/>`)

// Microchip is [microchip](https://lucide.dev/icons/microchip).
var Microchip = icon.Lucide(viewBox, paint, `<path d="M10 12h4"/><path d="M10 17h4"/><path d="M10 7h4"/><path d="M18 12h2"/><path d="M18 18h2"/><path d="M18 6h2"/><path d="M4 12h2"/><path d="M4 18h2"/><path d="M4 6h2"/><rect x="6" y="2" width="12" height="20" rx="2"/>`)

// Microscope is [microscope](https://lucide.dev/icons/microscope).
var Microscope = icon.Lucide(viewBox, paint, `<path d="M6 18h8"/><path d="M3 22h18"/><path d="M14 22a7 7 0 1 0 0-14h-1"/><path d="M9 14h2"/><path d="M9 12a2 2 0 0 1-2-2V6h6v4a2 2 0 0 1-2 2Z"/><path d="M12 6V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3"/>`)

// Microwave is [microwave](https://lucide.dev/icons/microwave).
var Microwave = icon.Lucide(viewBox, paint, `<rect width="20" height="15" x="2" y="4" rx="2"/><rect width="8" height="7" x="6" y="8" rx="1"/><path d="M18 8v7"/><path d="M6 19v2"/><path d="M18 19v2"/>`)

// MidiPort is [midi-port](https://lucide.dev/icons/midi-port).
var MidiPort = icon.Lucide(viewBox, paint, `<path d="M12 18h.01"/><path d="M15 2.458V5a1 1 0 01-1 1h-4a1 1 0 01-1-1V2.458"/><path d="M16 16h.01"/><path d="M18 12h.01"/><path d="M6 12h.01"/><path d="M8 16h.01"/><circle cx="12" cy="12" r="10"/>`)

// Milestone is [milestone](https://lucide.dev/icons/milestone).
var Milestone = icon.Lucide(viewBox, paint, `<path d="M12 13v8"/><path d="M12 3v3"/><path d="M18.172 6a2 2 0 0 1 1.414.586l2.06 2.06a1.207 1.207 0 0 1 0 1.708l-2.06 2.06a2 2 0 0 1-1.414.586H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1z"/>`)

// MilkOff is [milk-off](https://lucide.dev/icons/milk-off).
var MilkOff = icon.Lucide(viewBox, paint, `<path d="M8 2h8"/><path d="M9 2v1.343M15 2v2.789a4 4 0 0 0 .672 2.219l.656.984a4 4 0 0 1 .672 2.22v1.131M7.8 7.8l-.128.192A4 4 0 0 0 7 10.212V20a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-3"/><path d="M7 15a6.47 6.47 0 0 1 5 0 6.472 6.472 0 0 0 3.435.435"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Milk is [milk](https://lucide.dev/icons/milk).
var Milk = icon.Lucide(viewBox, paint, `<path d="M8 2h8"/><path d="M9 2v2.789a4 4 0 0 1-.672 2.219l-.656.984A4 4 0 0 0 7 10.212V20a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-9.789a4 4 0 0 0-.672-2.219l-.656-.984A4 4 0 0 1 15 4.788V2"/><path d="M7 15a6.472 6.472 0 0 1 5 0 6.47 6.47 0 0 0 5 0"/>`)

// Minimize2 is [minimize-2](https://lucide.dev/icons/minimize-2).
var Minimize2 = icon.Lucide(viewBox, paint, `<path d="m14 10 7-7"/><path d="M20 10h-6V4"/><path d="m3 21 7-7"/><path d="M4 14h6v6"/>`)

// Minimize is [minimize](https://lucide.dev/icons/minimize).
var Minimize = icon.Lucide(viewBox, paint, `<path d="M8 3v3a2 2 0 0 1-2 2H3"/><path d="M21 8h-3a2 2 0 0 1-2-2V3"/><path d="M3 16h3a2 2 0 0 1 2 2v3"/><path d="M16 21v-3a2 2 0 0 1 2-2h3"/>`)

// Minus is [minus](https://lucide.dev/icons/minus).
var Minus = icon.Lucide(viewBox, paint, `<path d="M5 12h14"/>`)

// MirrorRectangular is [mirror-rectangular](https://lucide.dev/icons/mirror-rectangular).
var MirrorRectangular = icon.Lucide(viewBox, paint, `<path d="M11 6 8 9"/><path d="m16 7-8 8"/><rect x="4" y="2" width="16" height="20" rx="2"/>`)

// MirrorRound is [mirror-round](https://lucide.dev/icons/mirror-round).
var MirrorRound = icon.Lucide(viewBox, paint, `<path d="M10 6.6 8.6 8"/><path d="M12 18v4"/><path d="M15 7.5 9.5 13"/><path d="M7 22h10"/><circle cx="12" cy="10" r="8"/>`)

// MonitorCheck is [monitor-check](https://lucide.dev/icons/monitor-check).
var MonitorCheck = icon.Lucide(viewBox, paint, `<path d="m9 10 2 2 4-4"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4"/><path d="M8 21h8"/>`)

// MonitorCloud is [monitor-cloud](https://lucide.dev/icons/monitor-cloud).
var MonitorCloud = icon.Lucide(viewBox, paint, `<path d="M11 13a3 3 0 1 1 2.83-4H14a2 2 0 0 1 0 4z"/><path d="M12 17v4"/><path d="M8 21h8"/><rect x="2" y="3" width="20" height="14" rx="2"/>`)

// MonitorCog is [monitor-cog](https://lucide.dev/icons/monitor-cog).
var MonitorCog = icon.Lucide(viewBox, paint, `<path d="M12 17v4"/><path d="m14.305 7.53.923-.382"/><path d="m15.228 4.852-.923-.383"/><path d="m16.852 3.228-.383-.924"/><path d="m16.852 8.772-.383.923"/><path d="m19.148 3.228.383-.924"/><path d="m19.53 9.696-.382-.924"/><path d="m20.772 4.852.924-.383"/><path d="m20.772 7.148.924.383"/><path d="M22 13v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7"/><path d="M8 21h8"/><circle cx="18" cy="6" r="3"/>`)

// MonitorDot is [monitor-dot](https://lucide.dev/icons/monitor-dot).
var MonitorDot = icon.Lucide(viewBox, paint, `<path d="M12 17v4"/><path d="M22 12.307V15a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h8.693"/><path d="M8 21h8"/><circle cx="19" cy="6" r="3"/>`)

// MonitorDown is [monitor-down](https://lucide.dev/icons/monitor-down).
var MonitorDown = icon.Lucide(viewBox, paint, `<path d="M12 13V7"/><path d="m15 10-3 3-3-3"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4"/><path d="M8 21h8"/>`)

// MonitorOff is [monitor-off](https://lucide.dev/icons/monitor-off).
var MonitorOff = icon.Lucide(viewBox, paint, `<path d="M12 17v4"/><path d="M17 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 1.184-1.826"/><path d="m2 2 20 20"/><path d="M8 21h8"/><path d="M8.656 3H20a2 2 0 0 1 2 2v10a2 2 0 0 1-.293 1.042"/>`)

// MonitorPause is [monitor-pause](https://lucide.dev/icons/monitor-pause).
var MonitorPause = icon.Lucide(viewBox, paint, `<path d="M10 13V7"/><path d="M14 13V7"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4"/><path d="M8 21h8"/>`)

// MonitorPc is [monitor-pc](https://lucide.dev/icons/monitor-pc).
var MonitorPc = icon.Lucide(viewBox, paint, `<path d="M10 15H4a2 2 0 01-2-2V7a2 2 0 012-2h6"/><path d="M10 19H5"/><path d="M14 11h8"/><path d="M14 7h8"/><path d="M18 17h.01"/><path d="M9 19v-4"/><rect x="14" y="3" width="8" height="18" rx="1"/>`)

// MonitorPlay is [monitor-play](https://lucide.dev/icons/monitor-play).
var MonitorPlay = icon.Lucide(viewBox, paint, `<path d="M15.033 9.44a.647.647 0 0 1 0 1.12l-4.065 2.352a.645.645 0 0 1-.968-.56V7.648a.645.645 0 0 1 .967-.56z"/><path d="M12 17v4"/><path d="M8 21h8"/><rect x="2" y="3" width="20" height="14" rx="2"/>`)

// MonitorSmartphone is [monitor-smartphone](https://lucide.dev/icons/monitor-smartphone).
var MonitorSmartphone = icon.Lucide(viewBox, paint, `<path d="M18 8V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8"/><path d="M10 19v-3.96 3.15"/><path d="M7 19h5"/><rect width="6" height="10" x="16" y="12" rx="2"/>`)

// MonitorSpeaker is [monitor-speaker](https://lucide.dev/icons/monitor-speaker).
var MonitorSpeaker = icon.Lucide(viewBox, paint, `<path d="M5.5 20H8"/><path d="M17 9h.01"/><rect width="10" height="16" x="12" y="4" rx="2"/><path d="M8 6H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h4"/><circle cx="17" cy="15" r="1"/>`)

// MonitorStop is [monitor-stop](https://lucide.dev/icons/monitor-stop).
var MonitorStop = icon.Lucide(viewBox, paint, `<path d="M12 17v4"/><path d="M8 21h8"/><rect x="2" y="3" width="20" height="14" rx="2"/><rect x="9" y="7" width="6" height="6" rx="1"/>`)

// MonitorUp is [monitor-up](https://lucide.dev/icons/monitor-up).
var MonitorUp = icon.Lucide(viewBox, paint, `<path d="m9 10 3-3 3 3"/><path d="M12 13V7"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4"/><path d="M8 21h8"/>`)

// MonitorX is [monitor-x](https://lucide.dev/icons/monitor-x).
var MonitorX = icon.Lucide(viewBox, paint, `<path d="m14.5 12.5-5-5"/><path d="m9.5 12.5 5-5"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4"/><path d="M8 21h8"/>`)

// Monitor is [monitor](https://lucide.dev/icons/monitor).
var Monitor = icon.Lucide(viewBox, paint, `<rect width="20" height="14" x="2" y="3" rx="2"/><line x1="8" x2="16" y1="21" y2="21"/><line x1="12" x2="12" y1="17" y2="21"/>`)

// MoonStar is [moon-star](https://lucide.dev/icons/moon-star).
var MoonStar = icon.Lucide(viewBox, paint, `<path d="M18 5h4"/><path d="M20 3v4"/><path d="M20.985 12.486a9 9 0 1 1-9.473-9.472c.405-.022.617.46.402.803a6 6 0 0 0 8.268 8.268c.344-.215.825-.004.803.401"/>`)

// Moon is [moon](https://lucide.dev/icons/moon).
var Moon = icon.Lucide(viewBox, paint, `<path d="M20.985 12.486a9 9 0 1 1-9.473-9.472c.405-.022.617.46.402.803a6 6 0 0 0 8.268 8.268c.344-.215.825-.004.803.401"/>`)

// MopSparkles is [mop-sparkles](https://lucide.dev/icons/mop-sparkles).
var MopSparkles = icon.Lucide(viewBox, paint, `<path d="M10 22a3 3 0 01-3-3"/><path d="M10 22c2.761 0 5-1.79 5-4-4.42 0-4.08-5-8.5-5a4.501 4.501 0 000 9z"/><path d="M10 3H8"/><path d="M12.5 11.5 22 2"/><path d="M20 13v4"/><path d="M22 15h-4"/><path d="M4 5v4"/><path d="M6 7H2"/><path d="m6.98 13.02 2.665-2.664a1.21 1.21 0 011.71 0l2.29 2.288a1.21 1.21 0 010 1.712l-2.088 2.087"/><path d="M9 2v2"/>`)

// Mop is [mop](https://lucide.dev/icons/mop).
var Mop = icon.Lucide(viewBox, paint, `<path d="M10 22c2.761 0 5-1.79 5-4-4.42 0-4.08-5-8.5-5a1 1 0 100 9za3 3 0 01-3-3"/><path d="M12.5 11.5 22 2"/><path d="m6.98 13.02 2.665-2.664a1.21 1.21 0 011.71 0l2.29 2.288a1.21 1.21 0 010 1.712l-2.088 2.087"/>`)

// Mosque is [mosque](https://lucide.dev/icons/mosque).
var Mosque = icon.Lucide(viewBox, paint, `<path d="M12.268 2a2 2 0 003.465 2"/><path d="M14 5 L14 8"/><path d="M16 22v-3a2 2 0 00-4 0v3"/><path d="M21 13c-.662-1.497-1.666-2.753-2.9-3.63C16.825 8.47 15.422 8 14 8s-2.826.47-4.1 1.37C8.668 10.248 7.663 11.504 7 13z"/><path d="M3 9h4"/><path d="M7 22V6a5 5 0 00-2-4 5 5 0 00-2 4v14a2 2 0 002 2h14a2 2 0 002-2v-7"/>`)

// Motorbike is [motorbike](https://lucide.dev/icons/motorbike).
var Motorbike = icon.Lucide(viewBox, paint, `<path d="m18 14-1-3"/><path d="m3 9 6 2a2 2 0 0 1 2-2h2a2 2 0 0 1 1.99 1.81"/><path d="M8 17h3a1 1 0 0 0 1-1 6 6 0 0 1 6-6 1 1 0 0 0 1-1v-.75A5 5 0 0 0 17 5"/><circle cx="19" cy="17" r="3"/><circle cx="5" cy="17" r="3"/>`)

// MountainSnow is [mountain-snow](https://lucide.dev/icons/mountain-snow).
var MountainSnow = icon.Lucide(viewBox, paint, `<path d="m8 3 4 8 5-5 5 15H2L8 3z"/><path d="M4.14 15.08c2.62-1.57 5.24-1.43 7.86.42 2.74 1.94 5.49 2 8.23.19"/>`)

// Mountain is [mountain](https://lucide.dev/icons/mountain).
var Mountain = icon.Lucide(viewBox, paint, `<path d="m8 3 4 8 5-5 5 15H2L8 3z"/>`)

// MouseLeft is [mouse-left](https://lucide.dev/icons/mouse-left).
var MouseLeft = icon.Lucide(viewBox, paint, `<path d="M12 7.318V10"/><path d="M5 10v5a7 7 0 0 0 14 0V9c0-3.527-2.608-6.515-6-7"/><circle cx="7" cy="4" r="2"/>`)

// MouseOff is [mouse-off](https://lucide.dev/icons/mouse-off).
var MouseOff = icon.Lucide(viewBox, paint, `<path d="M12 6v.343"/><path d="M18.218 18.218A7 7 0 0 1 5 15V9a7 7 0 0 1 .782-3.218"/><path d="M19 13.343V9A7 7 0 0 0 8.56 2.902"/><path d="M22 22 2 2"/>`)

// MousePointer2Off is [mouse-pointer-2-off](https://lucide.dev/icons/mouse-pointer-2-off).
var MousePointer2Off = icon.Lucide(viewBox, paint, `<path d="m15.55 8.45 5.138 2.087a.5.5 0 0 1-.063.947l-6.124 1.58a2 2 0 0 0-1.438 1.435l-1.579 6.126a.5.5 0 0 1-.947.063L8.45 15.551"/><path d="M22 2 2 22"/><path d="m6.816 11.528-2.779-6.84a.495.495 0 0 1 .651-.651l6.84 2.779"/>`)

// MousePointer2 is [mouse-pointer-2](https://lucide.dev/icons/mouse-pointer-2).
var MousePointer2 = icon.Lucide(viewBox, paint, `<path d="M4.037 4.688a.495.495 0 0 1 .651-.651l16 6.5a.5.5 0 0 1-.063.947l-6.124 1.58a2 2 0 0 0-1.438 1.435l-1.579 6.126a.5.5 0 0 1-.947.063z"/>`)

// MousePointerBan is [mouse-pointer-ban](https://lucide.dev/icons/mouse-pointer-ban).
var MousePointerBan = icon.Lucide(viewBox, paint, `<path d="M2.034 2.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.944L8.204 7.545a1 1 0 0 0-.66.66l-1.066 3.443a.5.5 0 0 1-.944.033z"/><circle cx="16" cy="16" r="6"/><path d="m11.8 11.8 8.4 8.4"/>`)

// MousePointerClick is [mouse-pointer-click](https://lucide.dev/icons/mouse-pointer-click).
var MousePointerClick = icon.Lucide(viewBox, paint, `<path d="M14 4.1 12 6"/><path d="m5.1 8-2.9-.8"/><path d="m6 12-1.9 2"/><path d="M7.2 2.2 8 5.1"/><path d="M9.037 9.69a.498.498 0 0 1 .653-.653l11 4.5a.5.5 0 0 1-.074.949l-4.349 1.041a1 1 0 0 0-.74.739l-1.04 4.35a.5.5 0 0 1-.95.074z"/>`)

// MousePointer is [mouse-pointer](https://lucide.dev/icons/mouse-pointer).
var MousePointer = icon.Lucide(viewBox, paint, `<path d="M12.586 12.586 19 19"/><path d="M3.688 3.037a.497.497 0 0 0-.651.651l6.5 15.999a.501.501 0 0 0 .947-.062l1.569-6.083a2 2 0 0 1 1.448-1.479l6.124-1.579a.5.5 0 0 0 .063-.947z"/>`)

// MouseRight is [mouse-right](https://lucide.dev/icons/mouse-right).
var MouseRight = icon.Lucide(viewBox, paint, `<path d="M12 7.318V10"/><path d="M19 10v5a7 7 0 0 1-14 0V9c0-3.527 2.608-6.515 6-7"/><circle cx="17" cy="4" r="2"/>`)

// Mouse is [mouse](https://lucide.dev/icons/mouse).
var Mouse = icon.Lucide(viewBox, paint, `<rect x="5" y="2" width="14" height="20" rx="7"/><path d="M12 6v4"/>`)

// MouthOff is [mouth-off](https://lucide.dev/icons/mouth-off).
var MouthOff = icon.Lucide(viewBox, paint, `<path d="M13.074 7.417a2.6 2.6 0 012.989.099c1.829 1.36 3.053 2.076 5.479 3.644a1 1 0 01.308 1.368 11.6 11.6 0 01-1.617 2.05"/><path d="M2 12a50.5 50.5 0 0010.99.99"/><path d="m2 2 20 20"/><path d="M21 11a1 1 0 011 1 51 51 0 01-3.734.61"/><path d="M7.695 7.695c-1.7 1.247-2.92 1.967-5.238 3.464a1 1 0 00-.307 1.369 11.6 11.6 0 0014.766 4.388"/>`)

// Mouth is [mouth](https://lucide.dev/icons/mouth).
var Mouth = icon.Lucide(viewBox, paint, `<path d="M2 12a50.5 50.5 0 0020 0 1 1 0 00-1-1"/><path d="M2.457 11.159a1 1 0 00-.307 1.369 11.59 11.59 0 0019.7 0 1 1 0 00-.308-1.368c-2.426-1.568-3.65-2.284-5.479-3.644a2.6 2.6 0 00-3.373.208 1 1 0 01-1.38 0 2.62 2.62 0 00-3.373-.208c-1.83 1.36-3.053 2.076-5.48 3.643"/>`)

// Move3d is [move-3d](https://lucide.dev/icons/move-3d).
var Move3d = icon.Lucide(viewBox, paint, `<path d="M5 3v16h16"/><path d="m5 19 6-6"/><path d="m2 6 3-3 3 3"/><path d="m18 16 3 3-3 3"/>`)

// MoveDiagonal2 is [move-diagonal-2](https://lucide.dev/icons/move-diagonal-2).
var MoveDiagonal2 = icon.Lucide(viewBox, paint, `<path d="M19 13v6h-6"/><path d="M5 11V5h6"/><path d="m5 5 14 14"/>`)

// MoveDiagonal is [move-diagonal](https://lucide.dev/icons/move-diagonal).
var MoveDiagonal = icon.Lucide(viewBox, paint, `<path d="M11 19H5v-6"/><path d="M13 5h6v6"/><path d="M19 5 5 19"/>`)

// MoveDownLeft is [move-down-left](https://lucide.dev/icons/move-down-left).
var MoveDownLeft = icon.Lucide(viewBox, paint, `<path d="M11 19H5V13"/><path d="M19 5L5 19"/>`)

// MoveDownRight is [move-down-right](https://lucide.dev/icons/move-down-right).
var MoveDownRight = icon.Lucide(viewBox, paint, `<path d="M19 13V19H13"/><path d="M5 5L19 19"/>`)

// MoveDown is [move-down](https://lucide.dev/icons/move-down).
var MoveDown = icon.Lucide(viewBox, paint, `<path d="M8 18L12 22L16 18"/><path d="M12 2V22"/>`)

// MoveHorizontal is [move-horizontal](https://lucide.dev/icons/move-horizontal).
var MoveHorizontal = icon.Lucide(viewBox, paint, `<path d="m18 8 4 4-4 4"/><path d="M2 12h20"/><path d="m6 8-4 4 4 4"/>`)

// MoveLeft is [move-left](https://lucide.dev/icons/move-left).
var MoveLeft = icon.Lucide(viewBox, paint, `<path d="M6 8L2 12L6 16"/><path d="M2 12H22"/>`)

// MoveRight is [move-right](https://lucide.dev/icons/move-right).
var MoveRight = icon.Lucide(viewBox, paint, `<path d="M18 8L22 12L18 16"/><path d="M2 12H22"/>`)

// MoveUpLeft is [move-up-left](https://lucide.dev/icons/move-up-left).
var MoveUpLeft = icon.Lucide(viewBox, paint, `<path d="M5 11V5H11"/><path d="M5 5L19 19"/>`)

// MoveUpRight is [move-up-right](https://lucide.dev/icons/move-up-right).
var MoveUpRight = icon.Lucide(viewBox, paint, `<path d="M13 5H19V11"/><path d="M19 5L5 19"/>`)

// MoveUp is [move-up](https://lucide.dev/icons/move-up).
var MoveUp = icon.Lucide(viewBox, paint, `<path d="M8 6L12 2L16 6"/><path d="M12 2V22"/>`)

// MoveVertical is [move-vertical](https://lucide.dev/icons/move-vertical).
var MoveVertical = icon.Lucide(viewBox, paint, `<path d="M12 2v20"/><path d="m8 18 4 4 4-4"/><path d="m8 6 4-4 4 4"/>`)

// Move is [move](https://lucide.dev/icons/move).
var Move = icon.Lucide(viewBox, paint, `<path d="M12 2v20"/><path d="m15 19-3 3-3-3"/><path d="m19 9 3 3-3 3"/><path d="M2 12h20"/><path d="m5 9-3 3 3 3"/><path d="m9 5 3-3 3 3"/>`)

// Navigation2Off is [navigation-2-off](https://lucide.dev/icons/navigation-2-off).
var Navigation2Off = icon.Lucide(viewBox, paint, `<path d="M9.31 9.31 5 21l7-4 7 4-1.17-3.17"/><path d="M14.53 8.88 12 2l-1.17 3.17"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Navigation2 is [navigation-2](https://lucide.dev/icons/navigation-2).
var Navigation2 = icon.Lucide(viewBox, paint, `<polygon points="12 2 19 21 12 17 5 21 12 2"/>`)

// NavigationOff is [navigation-off](https://lucide.dev/icons/navigation-off).
var NavigationOff = icon.Lucide(viewBox, paint, `<path d="M8.43 8.43 3 11l8 2 2 8 2.57-5.43"/><path d="M17.39 11.73 22 2l-9.73 4.61"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Navigation is [navigation](https://lucide.dev/icons/navigation).
var Navigation = icon.Lucide(viewBox, paint, `<polygon points="3 11 22 2 13 21 11 13 3 11"/>`)

// NepaliRupee is [nepali-rupee](https://lucide.dev/icons/nepali-rupee).
var NepaliRupee = icon.Lucide(viewBox, paint, `<path d="M18 16.173 A4.74 4.74 0 0 0 13.496 8.005"/><path d="M4 3 L20 3"/><path d="M5 13 L13.5 21"/><path d="M5 13 L9 13"/><path d="M8 13 C15.5 13 14.667 3 8 3"/>`)

// Network is [network](https://lucide.dev/icons/network).
var Network = icon.Lucide(viewBox, paint, `<rect x="16" y="16" width="6" height="6" rx="1"/><rect x="2" y="16" width="6" height="6" rx="1"/><rect x="9" y="2" width="6" height="6" rx="1"/><path d="M5 16v-3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3"/><path d="M12 12V8"/>`)

// Newspaper is [newspaper](https://lucide.dev/icons/newspaper).
var Newspaper = icon.Lucide(viewBox, paint, `<path d="M15 18h-5"/><path d="M18 14h-8"/><path d="M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16a2 2 0 0 1-4 0v-9a2 2 0 0 1 2-2h2"/><rect width="8" height="4" x="10" y="6" rx="1"/>`)

// Nfc is [nfc](https://lucide.dev/icons/nfc).
var Nfc = icon.Lucide(viewBox, paint, `<path d="M6 8.32a7.43 7.43 0 0 1 0 7.36"/><path d="M9.46 6.21a11.76 11.76 0 0 1 0 11.58"/><path d="M12.91 4.1a15.91 15.91 0 0 1 .01 15.8"/><path d="M16.37 2a20.16 20.16 0 0 1 0 20"/>`)

// NonBinary is [non-binary](https://lucide.dev/icons/non-binary).
var NonBinary = icon.Lucide(viewBox, paint, `<path d="M12 2v10"/><path d="m8.5 4 7 4"/><path d="m8.5 8 7-4"/><circle cx="12" cy="17" r="5"/>`)

// NotebookDot is [notebook-dot](https://lucide.dev/icons/notebook-dot).
var NotebookDot = icon.Lucide(viewBox, paint, `<path d="M16 11.75V22"/><path d="M2 10h4"/><path d="M2 14h4"/><path d="M2 18h4"/><path d="M2 6h4"/><path d="M20 11.75V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h5.65"/><circle cx="18" cy="5" r="3"/>`)

// NotebookPen is [notebook-pen](https://lucide.dev/icons/notebook-pen).
var NotebookPen = icon.Lucide(viewBox, paint, `<path d="M13.4 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7.4"/><path d="M2 6h4"/><path d="M2 10h4"/><path d="M2 14h4"/><path d="M2 18h4"/><path d="M21.378 5.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/>`)

// NotebookTabs is [notebook-tabs](https://lucide.dev/icons/notebook-tabs).
var NotebookTabs = icon.Lucide(viewBox, paint, `<path d="M2 6h4"/><path d="M2 10h4"/><path d="M2 14h4"/><path d="M2 18h4"/><rect width="16" height="20" x="4" y="2" rx="2"/><path d="M15 2v20"/><path d="M15 7h5"/><path d="M15 12h5"/><path d="M15 17h5"/>`)

// NotebookText is [notebook-text](https://lucide.dev/icons/notebook-text).
var NotebookText = icon.Lucide(viewBox, paint, `<path d="M2 6h4"/><path d="M2 10h4"/><path d="M2 14h4"/><path d="M2 18h4"/><rect width="16" height="20" x="4" y="2" rx="2"/><path d="M9.5 8h5"/><path d="M9.5 12H16"/><path d="M9.5 16H14"/>`)

// Notebook is [notebook](https://lucide.dev/icons/notebook).
var Notebook = icon.Lucide(viewBox, paint, `<path d="M2 6h4"/><path d="M2 10h4"/><path d="M2 14h4"/><path d="M2 18h4"/><rect width="16" height="20" x="4" y="2" rx="2"/><path d="M16 2v20"/>`)

// NotepadTextDashed is [notepad-text-dashed](https://lucide.dev/icons/notepad-text-dashed).
var NotepadTextDashed = icon.Lucide(viewBox, paint, `<path d="M8 2v4"/><path d="M12 2v4"/><path d="M16 2v4"/><path d="M16 4h2a2 2 0 0 1 2 2v2"/><path d="M20 12v2"/><path d="M20 18v2a2 2 0 0 1-2 2h-1"/><path d="M13 22h-2"/><path d="M7 22H6a2 2 0 0 1-2-2v-2"/><path d="M4 14v-2"/><path d="M4 8V6a2 2 0 0 1 2-2h2"/><path d="M8 10h6"/><path d="M8 14h8"/><path d="M8 18h5"/>`)

// NotepadText is [notepad-text](https://lucide.dev/icons/notepad-text).
var NotepadText = icon.Lucide(viewBox, paint, `<path d="M8 2v4"/><path d="M12 2v4"/><path d="M16 2v4"/><rect width="16" height="18" x="4" y="4" rx="2"/><path d="M8 10h6"/><path d="M8 14h8"/><path d="M8 18h5"/>`)

// NutOff is [nut-off](https://lucide.dev/icons/nut-off).
var NutOff = icon.Lucide(viewBox, paint, `<path d="M11.868 11.868a.88.88 0 01-.488.252c-1.78.28-3.54-.17-4.88-.62 0 1.272-.229 3.578-.653 5.347a10 10 0 01-.417 1.363c-.21.52-.82.55-1.17.12a10 10 0 01.677-13.393"/><path d="M12.14 6.485a27.4 27.4 0 004.707-.638L20 9a7.23 7.23 0 011.706 7.05"/><path d="m2 2 20 20"/><path d="M20.707 20.707A1 1 0 0120 21h-1c-1.069 0-1.648.242-2.485.552A7.2 7.2 0 019.002 20l-3.155-3.153"/><path d="M8.356 2.7a10 10 0 019.974 1.56c.43.35.4.97-.12 1.17a10 10 0 01-1.363.417"/>`)

// Nut is [nut](https://lucide.dev/icons/nut).
var Nut = icon.Lucide(viewBox, paint, `<path d="M16.847 5.847 20 9a7.23 7.23 0 011.551 7.516C21.241 17.352 21 17.932 21 19v1a1 1 0 01-1 1h-1c-1.069 0-1.648.242-2.485.552A7.2 7.2 0 019.002 20l-3.155-3.153"/><path d="M18.21 5.43c-1.71.69-5.07 1.07-6.71 1.07.46 1.38.91 2.74.61 4.88a.88.88 0 01-.73.74c-1.78.28-3.54-.17-4.88-.62 0 1.64-.38 5-1.07 6.71-.21.52-.82.55-1.17.12A10 10 0 0118.33 4.26c.43.35.4.97-.12 1.17"/><path d="M4.93 4.93 3 3a.7.7 0 010-1"/>`)

// OctagonAlert is [octagon-alert](https://lucide.dev/icons/octagon-alert).
var OctagonAlert = icon.Lucide(viewBox, paint, `<path d="M12 16h.01"/><path d="M12 8v4"/><path d="M15.312 2a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586l-4.688-4.688A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2z"/>`)

// OctagonMinus is [octagon-minus](https://lucide.dev/icons/octagon-minus).
var OctagonMinus = icon.Lucide(viewBox, paint, `<path d="M2.586 16.726A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2h6.624a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586z"/><path d="M8 12h8"/>`)

// OctagonPause is [octagon-pause](https://lucide.dev/icons/octagon-pause).
var OctagonPause = icon.Lucide(viewBox, paint, `<path d="M10 15V9"/><path d="M14 15V9"/><path d="M2.586 16.726A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2h6.624a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586z"/>`)

// OctagonX is [octagon-x](https://lucide.dev/icons/octagon-x).
var OctagonX = icon.Lucide(viewBox, paint, `<path d="m15 9-6 6"/><path d="M2.586 16.726A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2h6.624a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586z"/><path d="m9 9 6 6"/>`)

// Octagon is [octagon](https://lucide.dev/icons/octagon).
var Octagon = icon.Lucide(viewBox, paint, `<path d="M2.586 16.726A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2h6.624a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586z"/>`)

// Omega is [omega](https://lucide.dev/icons/omega).
var Omega = icon.Lucide(viewBox, paint, `<path d="M3 20h4.5a.5.5 0 0 0 .5-.5v-.282a.52.52 0 0 0-.247-.437 8 8 0 1 1 8.494-.001.52.52 0 0 0-.247.438v.282a.5.5 0 0 0 .5.5H21"/>`)

// Option is [option](https://lucide.dev/icons/option).
var Option = icon.Lucide(viewBox, paint, `<path d="M14 3h7"/><path d="M3 3h5.28a1 1 0 0 1 .948.684l5.544 16.632a1 1 0 0 0 .949.684H21"/>`)

// Orbit is [orbit](https://lucide.dev/icons/orbit).
var Orbit = icon.Lucide(viewBox, paint, `<path d="M20.341 6.484A10 10 0 0 1 10.266 21.85"/><path d="M3.659 17.516A10 10 0 0 1 13.74 2.152"/><circle cx="12" cy="12" r="3"/><circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/>`)

// Origami is [origami](https://lucide.dev/icons/origami).
var Origami = icon.Lucide(viewBox, paint, `<path d="M12 12V4a1 1 0 0 1 1-1h6.297a1 1 0 0 1 .651 1.759l-4.696 4.025"/><path d="m12 21-7.414-7.414A2 2 0 0 1 4 12.172V6.415a1.002 1.002 0 0 1 1.707-.707L20 20.009"/><path d="m12.214 3.381 8.414 14.966a1 1 0 0 1-.167 1.199l-1.168 1.163a1 1 0 0 1-.706.291H6.351a1 1 0 0 1-.625-.219L3.25 18.8a1 1 0 0 1 .631-1.781l4.165.027"/>`)

// Package2 is [package-2](https://lucide.dev/icons/package-2).
var Package2 = icon.Lucide(viewBox, paint, `<path d="M12 3v6"/><path d="M16.76 3a2 2 0 0 1 1.8 1.1l2.23 4.479a2 2 0 0 1 .21.891V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9.472a2 2 0 0 1 .211-.894L5.45 4.1A2 2 0 0 1 7.24 3z"/><path d="M3.054 9.013h17.893"/>`)

// PackageCheck is [package-check](https://lucide.dev/icons/package-check).
var PackageCheck = icon.Lucide(viewBox, paint, `<path d="M12 22V12"/><path d="m16 17 2 2 4-4"/><path d="M21 11.127V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.729l7 4a2 2 0 0 0 2 .001l1.32-.753"/><path d="M3.29 7 12 12l8.71-5"/><path d="m7.5 4.27 8.997 5.148"/>`)

// PackageMinus is [package-minus](https://lucide.dev/icons/package-minus).
var PackageMinus = icon.Lucide(viewBox, paint, `<path d="M12 22V12"/><path d="M16 17h6"/><path d="M21 13V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.729l7 4a2 2 0 0 0 2 .001l1.675-.955"/><path d="M3.29 7 12 12l8.71-5"/><path d="m7.5 4.27 8.997 5.148"/>`)

// PackageOpen is [package-open](https://lucide.dev/icons/package-open).
var PackageOpen = icon.Lucide(viewBox, paint, `<path d="M12 22v-9"/><path d="M15.17 2.21a1.67 1.67 0 0 1 1.63 0L21 4.57a1.93 1.93 0 0 1 0 3.36L8.82 14.79a1.655 1.655 0 0 1-1.64 0L3 12.43a1.93 1.93 0 0 1 0-3.36z"/><path d="M20 13v3.87a2.06 2.06 0 0 1-1.11 1.83l-6 3.08a1.93 1.93 0 0 1-1.78 0l-6-3.08A2.06 2.06 0 0 1 4 16.87V13"/><path d="M21 12.43a1.93 1.93 0 0 0 0-3.36L8.83 2.2a1.64 1.64 0 0 0-1.63 0L3 4.57a1.93 1.93 0 0 0 0 3.36l12.18 6.86a1.636 1.636 0 0 0 1.63 0z"/>`)

// PackagePlus is [package-plus](https://lucide.dev/icons/package-plus).
var PackagePlus = icon.Lucide(viewBox, paint, `<path d="M12 22V12"/><path d="M16 17h6"/><path d="M19 14v6"/><path d="M21 10.535V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.729l7 4a2 2 0 0 0 2 .001l1.675-.955"/><path d="M3.29 7 12 12l8.71-5"/><path d="m7.5 4.27 8.997 5.148"/>`)

// PackageSearch is [package-search](https://lucide.dev/icons/package-search).
var PackageSearch = icon.Lucide(viewBox, paint, `<path d="M12 22V12"/><path d="M20.27 18.27 22 20"/><path d="M21 10.498V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.729l7 4a2 2 0 0 0 2 .001l.98-.559"/><path d="M3.29 7 12 12l8.71-5"/><path d="m7.5 4.27 8.997 5.148"/><circle cx="18.5" cy="16.5" r="2.5"/>`)

// PackageX is [package-x](https://lucide.dev/icons/package-x).
var PackageX = icon.Lucide(viewBox, paint, `<path d="M12 22V12"/><path d="m16.5 14.5 5 5"/><path d="m16.5 19.5 5-5"/><path d="M21 10.5V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.729l7 4a2 2 0 0 0 2 .001l.13-.074"/><path d="M3.29 7 12 12l8.71-5"/><path d="m7.5 4.27 8.997 5.148"/>`)

// Package is [package](https://lucide.dev/icons/package).
var Package = icon.Lucide(viewBox, paint, `<path d="M11 21.73a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73z"/><path d="M12 22V12"/><polyline points="3.29 7 12 12 20.71 7"/><path d="m7.5 4.27 9 5.15"/>`)

// PaintBucket is [paint-bucket](https://lucide.dev/icons/paint-bucket).
var PaintBucket = icon.Lucide(viewBox, paint, `<path d="M11 7 6 2"/><path d="M18.992 12H2.041"/><path d="M21.145 18.38A3.34 3.34 0 0 1 20 16.5a3.3 3.3 0 0 1-1.145 1.88c-.575.46-.855 1.02-.855 1.595A2 2 0 0 0 20 22a2 2 0 0 0 2-2.025c0-.58-.285-1.13-.855-1.595"/><path d="m8.5 4.5 2.148-2.148a1.205 1.205 0 0 1 1.704 0l7.296 7.296a1.205 1.205 0 0 1 0 1.704l-7.592 7.592a3.615 3.615 0 0 1-5.112 0l-3.888-3.888a3.615 3.615 0 0 1 0-5.112L5.67 7.33"/>`)

// PaintRoller is [paint-roller](https://lucide.dev/icons/paint-roller).
var PaintRoller = icon.Lucide(viewBox, paint, `<rect width="16" height="6" x="2" y="2" rx="2"/><path d="M10 16v-2a2 2 0 0 1 2-2h8a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2"/><rect width="4" height="6" x="8" y="16" rx="1"/>`)

// PaintbrushVertical is [paintbrush-vertical](https://lucide.dev/icons/paintbrush-vertical).
var PaintbrushVertical = icon.Lucide(viewBox, paint, `<path d="M10 2v2"/><path d="M14 2v4"/><path d="M17 2a1 1 0 0 1 1 1v9H6V3a1 1 0 0 1 1-1z"/><path d="M6 12a1 1 0 0 0-1 1v1a2 2 0 0 0 2 2h2a1 1 0 0 1 1 1v2.9a2 2 0 1 0 4 0V17a1 1 0 0 1 1-1h2a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1"/>`)

// Paintbrush is [paintbrush](https://lucide.dev/icons/paintbrush).
var Paintbrush = icon.Lucide(viewBox, paint, `<path d="m14.622 17.897-10.68-2.913"/><path d="M18.376 2.622a1 1 0 1 1 3.002 3.002L17.36 9.643a.5.5 0 0 0 0 .707l.944.944a2.41 2.41 0 0 1 0 3.408l-.944.944a.5.5 0 0 1-.707 0L8.354 7.348a.5.5 0 0 1 0-.707l.944-.944a2.41 2.41 0 0 1 3.408 0l.944.944a.5.5 0 0 0 .707 0z"/><path d="M9 8c-1.804 2.71-3.97 3.46-6.583 3.948a.507.507 0 0 0-.302.819l7.32 8.883a1 1 0 0 0 1.185.204C12.735 20.405 16 16.792 16 15"/>`)

// Palette is [palette](https://lucide.dev/icons/palette).
var Palette = icon.Lucide(viewBox, paint, `<path d="M12 22a1 1 0 0 1 0-20 10 9 0 0 1 10 9 5 5 0 0 1-5 5h-2.25a1.75 1.75 0 0 0-1.4 2.8l.3.4a1.75 1.75 0 0 1-1.4 2.8z"/><circle cx="13.5" cy="6.5" r=".5" fill="currentColor"/><circle cx="17.5" cy="10.5" r=".5" fill="currentColor"/><circle cx="6.5" cy="12.5" r=".5" fill="currentColor"/><circle cx="8.5" cy="7.5" r=".5" fill="currentColor"/>`)

// Panda is [panda](https://lucide.dev/icons/panda).
var Panda = icon.Lucide(viewBox, paint, `<path d="M11.25 17.25h1.5L12 18z"/><path d="m15 12 2 2"/><path d="M17.902 6.599a8 8 0 0 0-.5-.5"/><path d="M2 14.5C2 19.47 6.48 22 12 22s10-2.53 10-7.5a10 10 0 0 0-1.3-4.83 4.5 4.5 0 1 0-7.05-5.5 8 8 0 0 0-3.3 0 4.5 4.5 0 1 0-7.04 5.5A10 10 0 0 0 2 14.5"/><path d="M6.099 6.599a8 8 0 0 1 .5-.5"/><path d="m9 12-2 2"/>`)

// PanelBottomClose is [panel-bottom-close](https://lucide.dev/icons/panel-bottom-close).
var PanelBottomClose = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 15h18"/><path d="m15 8-3 3-3-3"/>`)

// PanelBottomDashed is [panel-bottom-dashed](https://lucide.dev/icons/panel-bottom-dashed).
var PanelBottomDashed = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M14 15h1"/><path d="M19 15h2"/><path d="M3 15h2"/><path d="M9 15h1"/>`)

// PanelBottomOpen is [panel-bottom-open](https://lucide.dev/icons/panel-bottom-open).
var PanelBottomOpen = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 15h18"/><path d="m9 10 3-3 3 3"/>`)

// PanelBottom is [panel-bottom](https://lucide.dev/icons/panel-bottom).
var PanelBottom = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 15h18"/>`)

// PanelLeftClose is [panel-left-close](https://lucide.dev/icons/panel-left-close).
var PanelLeftClose = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="m16 15-3-3 3-3"/>`)

// PanelLeftDashed is [panel-left-dashed](https://lucide.dev/icons/panel-left-dashed).
var PanelLeftDashed = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 14v1"/><path d="M9 19v2"/><path d="M9 3v2"/><path d="M9 9v1"/>`)

// PanelLeftOpen is [panel-left-open](https://lucide.dev/icons/panel-left-open).
var PanelLeftOpen = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="m14 9 3 3-3 3"/>`)

// PanelLeftRightDashed is [panel-left-right-dashed](https://lucide.dev/icons/panel-left-right-dashed).
var PanelLeftRightDashed = icon.Lucide(viewBox, paint, `<path d="M15 10V9"/><path d="M15 15v-1"/><path d="M15 21v-2"/><path d="M15 5V3"/><path d="M9 10V9"/><path d="M9 15v-1"/><path d="M9 21v-2"/><path d="M9 5V3"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// PanelLeft is [panel-left](https://lucide.dev/icons/panel-left).
var PanelLeft = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/>`)

// PanelRightClose is [panel-right-close](https://lucide.dev/icons/panel-right-close).
var PanelRightClose = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M15 3v18"/><path d="m8 9 3 3-3 3"/>`)

// PanelRightDashed is [panel-right-dashed](https://lucide.dev/icons/panel-right-dashed).
var PanelRightDashed = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M15 14v1"/><path d="M15 19v2"/><path d="M15 3v2"/><path d="M15 9v1"/>`)

// PanelRightOpen is [panel-right-open](https://lucide.dev/icons/panel-right-open).
var PanelRightOpen = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M15 3v18"/><path d="m10 15-3-3 3-3"/>`)

// PanelRight is [panel-right](https://lucide.dev/icons/panel-right).
var PanelRight = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M15 3v18"/>`)

// PanelTopBottomDashed is [panel-top-bottom-dashed](https://lucide.dev/icons/panel-top-bottom-dashed).
var PanelTopBottomDashed = icon.Lucide(viewBox, paint, `<path d="M14 15h1"/><path d="M14 9h1"/><path d="M19 15h2"/><path d="M19 9h2"/><path d="M3 15h2"/><path d="M3 9h2"/><path d="M9 15h1"/><path d="M9 9h1"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// PanelTopClose is [panel-top-close](https://lucide.dev/icons/panel-top-close).
var PanelTopClose = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18"/><path d="m9 16 3-3 3 3"/>`)

// PanelTopDashed is [panel-top-dashed](https://lucide.dev/icons/panel-top-dashed).
var PanelTopDashed = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M14 9h1"/><path d="M19 9h2"/><path d="M3 9h2"/><path d="M9 9h1"/>`)

// PanelTopOpen is [panel-top-open](https://lucide.dev/icons/panel-top-open).
var PanelTopOpen = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18"/><path d="m15 14-3 3-3-3"/>`)

// PanelTop is [panel-top](https://lucide.dev/icons/panel-top).
var PanelTop = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18"/>`)

// PanelsLeftBottom is [panels-left-bottom](https://lucide.dev/icons/panels-left-bottom).
var PanelsLeftBottom = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="M9 15h12"/>`)

// PanelsRightBottom is [panels-right-bottom](https://lucide.dev/icons/panels-right-bottom).
var PanelsRightBottom = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 15h12"/><path d="M15 3v18"/>`)

// PanelsTopLeft is [panels-top-left](https://lucide.dev/icons/panels-top-left).
var PanelsTopLeft = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18"/><path d="M9 21V9"/>`)

// PaperBag is [paper-bag](https://lucide.dev/icons/paper-bag).
var PaperBag = icon.Lucide(viewBox, paint, `<path d="M5.364 3.848C4 6 3 9.652 3 12.652V19a2 2 0 002 2h14a2 2 0 002-2v-5c0-2.334-1.816-4.668-2.622-7.002"/><path d="M7 3h11.379a2 2 0 011.789 1.106l.723 1.447A1 1 0 0119.997 7h-8.525a2 2 0 01-1.789-1.106L8.79 4.105a2 2 0 10-3.579 1.789l2.261 4.522A5 5 0 018 12.652V21"/>`)

// Paperclip is [paperclip](https://lucide.dev/icons/paperclip).
var Paperclip = icon.Lucide(viewBox, paint, `<path d="m16 6-8.414 8.586a2 2 0 0 0 2.829 2.829l8.414-8.586a4 4 0 1 0-5.657-5.657l-8.379 8.551a6 6 0 1 0 8.485 8.485l8.379-8.551"/>`)

// Parasol is [parasol](https://lucide.dev/icons/parasol).
var Parasol = icon.Lucide(viewBox, paint, `<path d="M12.5 11.134 18.196 21"/><path d="M20.425 5.299a10 10 0 0 0-16.941 9.78c.183.563.843.774 1.355.478L20.16 6.711c.512-.296.66-.973.264-1.413"/><path d="M21 21H3"/>`)

// Parentheses is [parentheses](https://lucide.dev/icons/parentheses).
var Parentheses = icon.Lucide(viewBox, paint, `<path d="M8 21s-4-3-4-9 4-9 4-9"/><path d="M16 3s4 3 4 9-4 9-4 9"/>`)

// Park is [park](https://lucide.dev/icons/park).
var Park = icon.Lucide(viewBox, paint, `<path d="M12 18h10"/><path d="M13.248 9.998A4.5 4.5 0 0 0 11.75 8.6V8a1 1 0 0 0-7.5 0 4.9 4.9 0 0 0 2.25 9H8"/><path d="m15 14-2 6"/><path d="m19 14 2 6"/><path d="M21 14h-8"/><path d="M8 20v-5.922a2 2 0 0 0-.586-1.414L6.5 11.75"/><path d="M9.205 12.795 8 14"/><circle cx="19" cy="6" r="2"/>`)

// ParkingMeter is [parking-meter](https://lucide.dev/icons/parking-meter).
var ParkingMeter = icon.Lucide(viewBox, paint, `<path d="M11 15h2"/><path d="M12 12v3"/><path d="M12 19v3"/><path d="M15.282 19a1 1 0 0 0 .948-.68l2.37-6.988a7 7 0 1 0-13.2 0l2.37 6.988a1 1 0 0 0 .948.68z"/><path d="M9 9a3 3 0 1 1 6 0"/>`)

// PartyPopper is [party-popper](https://lucide.dev/icons/party-popper).
var PartyPopper = icon.Lucide(viewBox, paint, `<path d="M5.8 11.3 2 22l10.7-3.79"/><path d="M4 3h.01"/><path d="M22 8h.01"/><path d="M15 2h.01"/><path d="M22 20h.01"/><path d="m22 2-2.24.75a2.9 2.9 0 0 0-1.96 3.12c.1.86-.57 1.63-1.45 1.63h-.38c-.86 0-1.6.6-1.76 1.44L14 10"/><path d="m22 13-.82-.33c-.86-.34-1.82.2-1.98 1.11c-.11.7-.72 1.22-1.43 1.22H17"/><path d="m11 2 .33.82c.34.86-.2 1.82-1.11 1.98C9.52 4.9 9 5.52 9 6.23V7"/><path d="M11 13c1.93 1.93 2.83 4.17 2 5-.83.83-3.07-.07-5-2-1.93-1.93-2.83-4.17-2-5 .83-.83 3.07.07 5 2Z"/>`)

// Pause is [pause](https://lucide.dev/icons/pause).
var Pause = icon.Lucide(viewBox, paint, `<rect x="14" y="3" width="5" height="18" rx="1"/><rect x="5" y="3" width="5" height="18" rx="1"/>`)

// PawPrint is [paw-print](https://lucide.dev/icons/paw-print).
var PawPrint = icon.Lucide(viewBox, paint, `<circle cx="11" cy="4" r="2"/><circle cx="18" cy="8" r="2"/><circle cx="20" cy="16" r="2"/><path d="M9 10a5 5 0 0 1 5 5v3.5a3.5 3.5 0 0 1-6.84 1.045Q6.52 17.48 4.46 16.84A3.5 3.5 0 0 1 5.5 10Z"/>`)

// PcCase is [pc-case](https://lucide.dev/icons/pc-case).
var PcCase = icon.Lucide(viewBox, paint, `<rect width="14" height="20" x="5" y="2" rx="2"/><path d="M15 14h.01"/><path d="M9 6h6"/><path d="M9 10h6"/>`)

// PenLine is [pen-line](https://lucide.dev/icons/pen-line).
var PenLine = icon.Lucide(viewBox, paint, `<path d="M13 21h8"/><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"/>`)

// PenOff is [pen-off](https://lucide.dev/icons/pen-off).
var PenOff = icon.Lucide(viewBox, paint, `<path d="m10 10-6.157 6.162a2 2 0 0 0-.5.833l-1.322 4.36a.5.5 0 0 0 .622.624l4.358-1.323a2 2 0 0 0 .83-.5L14 13.982"/><path d="m12.829 7.172 4.359-4.346a1 1 0 1 1 3.986 3.986l-4.353 4.353"/><path d="m2 2 20 20"/>`)

// PenTool is [pen-tool](https://lucide.dev/icons/pen-tool).
var PenTool = icon.Lucide(viewBox, paint, `<path d="m18 13-.654-4.736a4 4 0 00-2.756-3.266l-9.163-2.9a2 2 0 00-2.017.493l-.82.819a2 2 0 00-.492 2.017l2.9 9.163a4 4 0 003.266 2.756L13 18"/><path d="M18.648 12.352a1.205 1.205 0 011.704 0l1.296 1.296a1.205 1.205 0 010 1.704l-6.296 6.296a1.205 1.205 0 01-1.704 0l-1.296-1.296a1.205 1.205 0 010-1.704z"/><path d="m3 3 6.586 6.586"/><circle cx="11" cy="11" r="2"/>`)

// Pen is [pen](https://lucide.dev/icons/pen).
var Pen = icon.Lucide(viewBox, paint, `<path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"/>`)

// PencilLine is [pencil-line](https://lucide.dev/icons/pencil-line).
var PencilLine = icon.Lucide(viewBox, paint, `<path d="M13 21h8"/><path d="m15 5 4 4"/><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"/>`)

// PencilOff is [pencil-off](https://lucide.dev/icons/pencil-off).
var PencilOff = icon.Lucide(viewBox, paint, `<path d="m10 10-6.157 6.162a2 2 0 0 0-.5.833l-1.322 4.36a.5.5 0 0 0 .622.624l4.358-1.323a2 2 0 0 0 .83-.5L14 13.982"/><path d="m12.829 7.172 4.359-4.346a1 1 0 1 1 3.986 3.986l-4.353 4.353"/><path d="m15 5 4 4"/><path d="m2 2 20 20"/>`)

// PencilRuler is [pencil-ruler](https://lucide.dev/icons/pencil-ruler).
var PencilRuler = icon.Lucide(viewBox, paint, `<path d="M13 7 8.7 2.7a2.41 2.41 0 0 0-3.4 0L2.7 5.3a2.41 2.41 0 0 0 0 3.4L7 13"/><path d="m8 6 2-2"/><path d="m18 16 2-2"/><path d="m17 11 4.3 4.3c.94.94.94 2.46 0 3.4l-2.6 2.6c-.94.94-2.46.94-3.4 0L11 17"/><path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"/><path d="m15 5 4 4"/>`)

// PencilSparkles is [pencil-sparkles](https://lucide.dev/icons/pencil-sparkles).
var PencilSparkles = icon.Lucide(viewBox, paint, `<path d="M10 3H8"/><path d="m15.007 5.008 3.987 3.986"/><path d="M20 15v4"/><path d="M21.174 6.813a2.82 2.82 0 0 0-3.986-3.987L3.842 16.175a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"/><path d="M22 17h-4"/><path d="M4 5v4"/><path d="M6 7H2"/><path d="M9 2v2"/>`)

// Pencil is [pencil](https://lucide.dev/icons/pencil).
var Pencil = icon.Lucide(viewBox, paint, `<path d="M21.174 6.812a1 1 0 0 0-3.986-3.987L3.842 16.174a2 2 0 0 0-.5.83l-1.321 4.352a.5.5 0 0 0 .623.622l4.353-1.32a2 2 0 0 0 .83-.497z"/><path d="m15 5 4 4"/>`)

// Pentagon is [pentagon](https://lucide.dev/icons/pentagon).
var Pentagon = icon.Lucide(viewBox, paint, `<path d="M10.83 2.38a2 2 0 0 1 2.34 0l8 5.74a2 2 0 0 1 .73 2.25l-3.04 9.26a2 2 0 0 1-1.9 1.37H7.04a2 2 0 0 1-1.9-1.37L2.1 10.37a2 2 0 0 1 .73-2.25z"/>`)

// Percent is [percent](https://lucide.dev/icons/percent).
var Percent = icon.Lucide(viewBox, paint, `<line x1="19" x2="5" y1="5" y2="19"/><circle cx="6.5" cy="6.5" r="2.5"/><circle cx="17.5" cy="17.5" r="2.5"/>`)

// PerfumeBottle is [perfume-bottle](https://lucide.dev/icons/perfume-bottle).
var PerfumeBottle = icon.Lucide(viewBox, paint, `<path d="M 5 22 H 11.072"/><path d="M15 5 C 23 4 23.328 12.676 20 14.536"/><path d="M15.802 10.308 C 14.776 9.465 14 8.496 14 7"/><path d="M2 2h.01"/><path d="M2 7h.01"/><path d="M5 4.5h.01"/><path d="M8.127 22 A 8 8 0 0 1 4 15 3 3 0 0 1 4.879 12.879C 6.879 10.879 10 10 10 7"/><circle cx="18" cy="18" r="4"/><rect x="9" y="2" width="6" height="5" rx="1"/>`)

// PersonStanding is [person-standing](https://lucide.dev/icons/person-standing).
var PersonStanding = icon.Lucide(viewBox, paint, `<circle cx="12" cy="5" r="1"/><path d="m9 20 3-6 3 6"/><path d="m6 8 6 2 6-2"/><path d="M12 10v4"/>`)

// Phi is [phi](https://lucide.dev/icons/phi).
var Phi = icon.Lucide(viewBox, paint, `<path d="M12 2v20"/><circle cx="12" cy="12" r="7"/>`)

// PhilippinePeso is [philippine-peso](https://lucide.dev/icons/philippine-peso).
var PhilippinePeso = icon.Lucide(viewBox, paint, `<path d="M20 11H4"/><path d="M20 7H4"/><path d="M7 21V4a1 1 0 0 1 1-1h4a1 1 0 0 1 0 12H7"/>`)

// PhoneCall is [phone-call](https://lucide.dev/icons/phone-call).
var PhoneCall = icon.Lucide(viewBox, paint, `<path d="M13 2a9 9 0 0 1 9 9"/><path d="M13 6a5 5 0 0 1 5 5"/><path d="M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"/>`)

// PhoneForwarded is [phone-forwarded](https://lucide.dev/icons/phone-forwarded).
var PhoneForwarded = icon.Lucide(viewBox, paint, `<path d="M14 6h8"/><path d="m18 2 4 4-4 4"/><path d="M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"/>`)

// PhoneIncoming is [phone-incoming](https://lucide.dev/icons/phone-incoming).
var PhoneIncoming = icon.Lucide(viewBox, paint, `<path d="M16 2v6h6"/><path d="m22 2-6 6"/><path d="M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"/>`)

// PhoneLog is [phone-log](https://lucide.dev/icons/phone-log).
var PhoneLog = icon.Lucide(viewBox, paint, `<path d="M13.832 16.568a1 1 0 001.213-.303l.355-.465A2 2 0 0117 15h3a2 2 0 012 2v3a2 2 0 01-2 2A18 18 0 012 4a2 2 0 012-2h3a2 2 0 012 2v3a2 2 0 01-.8 1.6l-.468.351a1 1 0 00-.292 1.233 14 14 0 006.392 6.384"/><path d="M14.396 4.171h7.171"/><path d="M17.343 8.17h4.224"/>`)

// PhoneMissed is [phone-missed](https://lucide.dev/icons/phone-missed).
var PhoneMissed = icon.Lucide(viewBox, paint, `<path d="m16 2 6 6"/><path d="m22 2-6 6"/><path d="M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"/>`)

// PhoneOff is [phone-off](https://lucide.dev/icons/phone-off).
var PhoneOff = icon.Lucide(viewBox, paint, `<path d="M10.1 13.9a14 14 0 0 0 3.732 2.668 1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2 18 18 0 0 1-12.728-5.272"/><path d="M22 2 2 22"/><path d="M4.76 13.582A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 .244.473"/>`)

// PhoneOutgoing is [phone-outgoing](https://lucide.dev/icons/phone-outgoing).
var PhoneOutgoing = icon.Lucide(viewBox, paint, `<path d="m16 8 6-6"/><path d="M22 8V2h-6"/><path d="M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"/>`)

// Phone is [phone](https://lucide.dev/icons/phone).
var Phone = icon.Lucide(viewBox, paint, `<path d="M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"/>`)

// Pi is [pi](https://lucide.dev/icons/pi).
var Pi = icon.Lucide(viewBox, paint, `<line x1="9" x2="9" y1="4" y2="20"/><path d="M4 7c0-1.7 1.3-3 3-3h13"/><path d="M18 20c-1.7 0-3-1.3-3-3V4"/>`)

// Piano is [piano](https://lucide.dev/icons/piano).
var Piano = icon.Lucide(viewBox, paint, `<path d="M10 13v4"/><path d="M14 13v4"/><path d="M18 13v4"/><path d="M2 13h20"/><path d="M22 11.5A3.5 3.5 0 0018.5 8a3.52 3.52 0 01-3.173-2A7 7 0 002 9v10a2 2 0 002 2h16a2 2 0 002-2z"/><path d="M6 13v4"/>`)

// Pickaxe is [pickaxe](https://lucide.dev/icons/pickaxe).
var Pickaxe = icon.Lucide(viewBox, paint, `<path d="m14 13-8.381 8.38a1 1 0 0 1-3.001-3L11 9.999"/><path d="M15.973 4.027A13 13 0 0 0 5.902 2.373c-1.398.342-1.092 2.158.277 2.601a19.9 19.9 0 0 1 5.822 3.024"/><path d="M16.001 11.999a19.9 19.9 0 0 1 3.024 5.824c.444 1.369 2.26 1.676 2.603.278A13 13 0 0 0 20 8.069"/><path d="M18.352 3.352a1.205 1.205 0 0 0-1.704 0l-5.296 5.296a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l5.296-5.296a1.205 1.205 0 0 0 0-1.704z"/>`)

// PictureInPicture2 is [picture-in-picture-2](https://lucide.dev/icons/picture-in-picture-2).
var PictureInPicture2 = icon.Lucide(viewBox, paint, `<path d="M21 9V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10c0 1.1.9 2 2 2h4"/><rect width="10" height="7" x="12" y="13" rx="2"/>`)

// PictureInPicture is [picture-in-picture](https://lucide.dev/icons/picture-in-picture).
var PictureInPicture = icon.Lucide(viewBox, paint, `<path d="M2 10h6V4"/><path d="m2 4 6 6"/><path d="M21 10V7a2 2 0 0 0-2-2h-7"/><path d="M3 14v2a2 2 0 0 0 2 2h3"/><rect x="12" y="14" width="10" height="7" rx="1"/>`)

// PiggyBank is [piggy-bank](https://lucide.dev/icons/piggy-bank).
var PiggyBank = icon.Lucide(viewBox, paint, `<path d="M11 17h3v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3a3.16 3.16 0 0 0 2-2h1a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-1a5 5 0 0 0-2-4V3a4 4 0 0 0-3.2 1.6l-.3.4H11a6 6 0 0 0-6 6v1a5 5 0 0 0 2 4v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1z"/><path d="M16 10h.01"/><path d="M2 8v1a2 2 0 0 0 2 2h1"/>`)

// PilcrowLeft is [pilcrow-left](https://lucide.dev/icons/pilcrow-left).
var PilcrowLeft = icon.Lucide(viewBox, paint, `<path d="M14 3v11"/><path d="M14 9h-3a3 3 0 0 1 0-6h9"/><path d="M18 3v11"/><path d="M22 18H2l4-4"/><path d="m6 22-4-4"/>`)

// PilcrowRight is [pilcrow-right](https://lucide.dev/icons/pilcrow-right).
var PilcrowRight = icon.Lucide(viewBox, paint, `<path d="M10 3v11"/><path d="M10 9H7a1 1 0 0 1 0-6h8"/><path d="M14 3v11"/><path d="m18 14 4 4H2"/><path d="m22 18-4 4"/>`)

// Pilcrow is [pilcrow](https://lucide.dev/icons/pilcrow).
var Pilcrow = icon.Lucide(viewBox, paint, `<path d="M13 4v16"/><path d="M17 4v16"/><path d="M19 4H9.5a4.5 4.5 0 0 0 0 9H13"/>`)

// PillBottle is [pill-bottle](https://lucide.dev/icons/pill-bottle).
var PillBottle = icon.Lucide(viewBox, paint, `<path d="M18 11h-4a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h4"/><path d="M6 7v13a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7"/><rect width="16" height="5" x="4" y="2" rx="1"/>`)

// Pill is [pill](https://lucide.dev/icons/pill).
var Pill = icon.Lucide(viewBox, paint, `<path d="m10.5 20.5 10-10a4.95 4.95 0 1 0-7-7l-10 10a4.95 4.95 0 1 0 7 7Z"/><path d="m8.5 8.5 7 7"/>`)

// PinOff is [pin-off](https://lucide.dev/icons/pin-off).
var PinOff = icon.Lucide(viewBox, paint, `<path d="M12 17v5"/><path d="M15 9.34V7a1 1 0 0 1 1-1 2 2 0 0 0 0-4H7.89"/><path d="m2 2 20 20"/><path d="M9 9v1.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V16a1 1 0 0 0 1 1h11"/>`)

// Pin is [pin](https://lucide.dev/icons/pin).
var Pin = icon.Lucide(viewBox, paint, `<path d="M12 17v5"/><path d="M9 10.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-.76a2 2 0 0 0-1.11-1.79l-1.78-.9A2 2 0 0 1 15 10.76V7a1 1 0 0 1 1-1 2 2 0 0 0 0-4H8a2 2 0 0 0 0 4 1 1 0 0 1 1 1z"/>`)

// Pipette is [pipette](https://lucide.dev/icons/pipette).
var Pipette = icon.Lucide(viewBox, paint, `<path d="m12 9-8.414 8.414A2 2 0 0 0 3 18.828v1.344a2 2 0 0 1-.586 1.414A2 2 0 0 1 3.828 21h1.344a2 2 0 0 0 1.414-.586L15 12"/><path d="m18 9 .4.4a1 1 0 1 1-3 3l-3.8-3.8a1 1 0 1 1 3-3l.4.4 3.4-3.4a1 1 0 1 1 3 3z"/><path d="m2 22 .414-.414"/>`)

// Pizza is [pizza](https://lucide.dev/icons/pizza).
var Pizza = icon.Lucide(viewBox, paint, `<path d="m12 14-1 1"/><path d="m13.75 18.25-1.25 1.42"/><path d="M17.775 5.654a15.68 15.68 0 0 0-12.121 12.12"/><path d="M18.8 9.3a1 1 0 0 0 2.1 7.7"/><path d="M21.964 20.732a1 1 0 0 1-1.232 1.232l-18-5a1 1 0 0 1-.695-1.232A19.68 19.68 0 0 1 15.732 2.037a1 1 0 0 1 1.232.695z"/>`)

// PlaneLanding is [plane-landing](https://lucide.dev/icons/plane-landing).
var PlaneLanding = icon.Lucide(viewBox, paint, `<path d="M2 22h20"/><path d="M3.77 10.77 2 9l2-4.5 1.1.55c.55.28.9.84.9 1.45s.35 1.17.9 1.45L8 8.5l3-6 1.05.53a2 2 0 0 1 1.09 1.52l.72 5.4a2 2 0 0 0 1.09 1.52l4.4 2.2c.42.22.78.55 1.01.96l.6 1.03c.49.88-.06 1.98-1.06 2.1l-1.18.15c-.47.06-.95-.02-1.37-.24L4.29 11.15a2 2 0 0 1-.52-.38Z"/>`)

// PlaneTakeoff is [plane-takeoff](https://lucide.dev/icons/plane-takeoff).
var PlaneTakeoff = icon.Lucide(viewBox, paint, `<path d="M2 22h20"/><path d="M6.36 17.4 4 17l-2-4 1.1-.55a2 2 0 0 1 1.8 0l.17.1a2 2 0 0 0 1.8 0L8 12 5 6l.9-.45a2 2 0 0 1 2.09.2l4.02 3a2 2 0 0 0 2.1.2l4.19-2.06a2.41 2.41 0 0 1 1.73-.17L21 7a1.4 1.4 0 0 1 .87 1.99l-.38.76c-.23.46-.6.84-1.07 1.08L7.58 17.2a2 2 0 0 1-1.22.18Z"/>`)

// Plane is [plane](https://lucide.dev/icons/plane).
var Plane = icon.Lucide(viewBox, paint, `<path d="M17.8 19.2 16 11l3.5-3.5C21 6 21.5 4 21 3c-1-.5-3 0-4.5 1.5L13 8 4.8 6.2c-.5-.1-.9.1-1.1.5l-.3.5c-.2.5-.1 1 .3 1.3L9 12l-2 3H4l-1 1 3 2 2 3 1-1v-3l3-2 3.5 5.3c.3.4.8.5 1.3.3l.5-.2c.4-.3.6-.7.5-1.2z"/>`)

// PlantPot is [plant-pot](https://lucide.dev/icons/plant-pot).
var PlantPot = icon.Lucide(viewBox, paint, `<path d="M14 8.536V6a4 4 0 014-4h1.5a.5.5 0 01.5.5V4a4 4 0 01-4 4 4 4 0 00-4 4 5 5 0 01-8-4 5 5 0 018 4c0 2 1 3 1 5"/><path d="m18 17-1.085 3.58A2 2 0 0115 22H9.002a2 2 0 01-1.913-1.418L6 17"/><path d="M5 17h14"/>`)

// PlayOff is [play-off](https://lucide.dev/icons/play-off).
var PlayOff = icon.Lucide(viewBox, paint, `<path d="m10.215 4.56 9.79 5.71a2 2 0 0 1 .003 3.458l-.393.23"/><path d="m16.042 16.042-8.034 4.686A2 2 0 0 1 5 19V5"/><path d="m2 2 20 20"/>`)

// Play is [play](https://lucide.dev/icons/play).
var Play = icon.Lucide(viewBox, paint, `<path d="M5 5a2 2 0 0 1 3.008-1.728l11.997 6.998a2 2 0 0 1 .003 3.458l-12 7A2 2 0 0 1 5 19z"/>`)

// PlayingCard is [playing-card](https://lucide.dev/icons/playing-card).
var PlayingCard = icon.Lucide(viewBox, paint, `<path d="M12.832 8.445a1 1 0 00-1.589-.098l-2.075 3.098a1 1 0 000 1.11l2 3a1 1 0 001.664 0l2-3a1 1 0 000-1.11z"/><rect x="5" y="2" width="14" height="20" rx="2"/>`)

// PlayingCardsFan is [playing-cards-fan](https://lucide.dev/icons/playing-cards-fan).
var PlayingCardsFan = icon.Lucide(viewBox, paint, `<path d="M12.65 7.65a2 2 0 012.629-1.046l5.51 2.374a2 2 0 011.046 2.628l-3.957 9.184a2 2 0 01-2.628 1.046l-5.51-2.374a2 2 0 01-1.046-2.628z"/><path d="M18 7.777V4a2 2 0 00-2-2h-6a2 2 0 00-2 2v10a2 2 0 001.137 1.805"/><path d="m8 4.389-4.364.809a2 2 0 00-1.602 2.33l1.822 9.833a2 2 0 002.331 1.602l2.542-.47"/>`)

// PlayingCards is [playing-cards](https://lucide.dev/icons/playing-cards).
var PlayingCards = icon.Lucide(viewBox, paint, `<path d="M14.832 8.445a1 1 0 00-1.589-.098l-2.075 3.098a1 1 0 000 1.11l2 3a1 1 0 001.664 0l2-3a1 1 0 000-1.11z"/><path d="m7.18 20.827-5-11a2 2 0 01.993-2.647L7 5.44"/><rect x="7" y="2" width="14" height="20" rx="2"/>`)

// Plug2 is [plug-2](https://lucide.dev/icons/plug-2).
var Plug2 = icon.Lucide(viewBox, paint, `<path d="M9 2v6"/><path d="M15 2v6"/><path d="M12 17v5"/><path d="M5 8h14"/><path d="M6 11V8h12v3a6 6 0 1 1-12 0Z"/>`)

// PlugZap is [plug-zap](https://lucide.dev/icons/plug-zap).
var PlugZap = icon.Lucide(viewBox, paint, `<path d="M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6-2.3 2.3a2.4 2.4 0 0 0 0 3.4Z"/><path d="m2 22 3-3"/><path d="M7.5 13.5 10 11"/><path d="M10.5 16.5 13 14"/><path d="m18 3-4 4h6l-4 4"/>`)

// Plug is [plug](https://lucide.dev/icons/plug).
var Plug = icon.Lucide(viewBox, paint, `<path d="M12 22v-5"/><path d="M15 8V2"/><path d="M17 8a1 1 0 0 1 1 1v4a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4V9a1 1 0 0 1 1-1z"/><path d="M9 8V2"/>`)

// Plus is [plus](https://lucide.dev/icons/plus).
var Plus = icon.Lucide(viewBox, paint, `<path d="M5 12h14"/><path d="M12 5v14"/>`)

// PocketKnife is [pocket-knife](https://lucide.dev/icons/pocket-knife).
var PocketKnife = icon.Lucide(viewBox, paint, `<path d="M3 2v1c0 1 2 1 2 2S3 6 3 7s2 1 2 2-2 1-2 2 2 1 2 2"/><path d="M18 6h.01"/><path d="M6 18h.01"/><path d="M20.83 8.83a4 4 0 0 0-5.66-5.66l-12 12a4 4 0 1 0 5.66 5.66Z"/><path d="M18 11.66V22a4 4 0 0 0 4-4V6"/>`)

// Podium is [podium](https://lucide.dev/icons/podium).
var Podium = icon.Lucide(viewBox, paint, `<path d="M12 6V2h-1"/><path d="M9 15a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1"/><path d="M9 21V11a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v10"/>`)

// PointerOff is [pointer-off](https://lucide.dev/icons/pointer-off).
var PointerOff = icon.Lucide(viewBox, paint, `<path d="M10 4.5V4a2 2 0 0 0-2.41-1.957"/><path d="M13.9 8.4a2 2 0 0 0-1.26-1.295"/><path d="M21.7 16.2A8 8 0 0 0 22 14v-3a2 2 0 1 0-4 0v-1a2 2 0 0 0-3.63-1.158"/><path d="m7 15-1.8-1.8a2 2 0 0 0-2.79 2.86L6 19.7a7.74 7.74 0 0 0 6 2.3h2a8 8 0 0 0 5.657-2.343"/><path d="M6 6v8"/><path d="m2 2 20 20"/>`)

// Pointer is [pointer](https://lucide.dev/icons/pointer).
var Pointer = icon.Lucide(viewBox, paint, `<path d="M22 14a8 8 0 0 1-8 8"/><path d="M18 11v-1a2 2 0 0 0-2-2a2 2 0 0 0-2 2"/><path d="M14 10V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v1"/><path d="M10 9.5V4a2 2 0 0 0-2-2a2 2 0 0 0-2 2v10"/><path d="M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15"/>`)

// Popcorn is [popcorn](https://lucide.dev/icons/popcorn).
var Popcorn = icon.Lucide(viewBox, paint, `<path d="M18 8a2 2 0 0 0 0-4 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0 0 4"/><path d="M10 22 9 8"/><path d="m14 22 1-14"/><path d="M20 8c.5 0 .9.4.8 1l-2.6 12c-.1.5-.7 1-1.2 1H7c-.6 0-1.1-.4-1.2-1L3.2 9c-.1-.6.3-1 .8-1Z"/>`)

// Popsicle is [popsicle](https://lucide.dev/icons/popsicle).
var Popsicle = icon.Lucide(viewBox, paint, `<path d="M18.6 14.4c.8-.8.8-2 0-2.8l-8.1-8.1a4.95 4.95 0 1 0-7.1 7.1l8.1 8.1c.9.7 2.1.7 2.9-.1Z"/><path d="m22 22-5.5-5.5"/>`)

// PoundSterling is [pound-sterling](https://lucide.dev/icons/pound-sterling).
var PoundSterling = icon.Lucide(viewBox, paint, `<path d="M18 7c0-5.333-8-5.333-8 0"/><path d="M10 7v14"/><path d="M6 21h12"/><path d="M6 13h10"/>`)

// PowerOff is [power-off](https://lucide.dev/icons/power-off).
var PowerOff = icon.Lucide(viewBox, paint, `<path d="M18.36 6.64A9 9 0 0 1 20.77 15"/><path d="M6.16 6.16a9 9 0 1 0 12.68 12.68"/><path d="M12 2v4"/><path d="m2 2 20 20"/>`)

// Power is [power](https://lucide.dev/icons/power).
var Power = icon.Lucide(viewBox, paint, `<path d="M12 2v10"/><path d="M18.4 6.6a9 9 0 1 1-12.77.04"/>`)

// Presentation is [presentation](https://lucide.dev/icons/presentation).
var Presentation = icon.Lucide(viewBox, paint, `<path d="M2 3h20"/><path d="M21 3v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3"/><path d="m7 21 5-5 5 5"/>`)

// Printer3d is [printer-3d](https://lucide.dev/icons/printer-3d).
var Printer3d = icon.Lucide(viewBox, paint, `<path d="M10 11v1"/><path d="M12 8h8"/><path d="M15 20v-3a1 1 0 00-1-1H9a1 1 0 00-1 1v3"/><path d="M4 20h16"/><path d="M4 22V4a2 2 0 012-2h12a2 2 0 012 2v18"/><path d="M4 8h4"/><path d="M8.635 10.093A2 2 0 018 8.631V7a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 01-.293.707l-1 1a1 1 0 01-1.414 0z"/>`)

// PrinterCheck is [printer-check](https://lucide.dev/icons/printer-check).
var PrinterCheck = icon.Lucide(viewBox, paint, `<path d="M13.5 22H7a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v.5"/><path d="m16 19 2 2 4-4"/><path d="M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2"/><path d="M6 9V3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v6"/>`)

// PrinterX is [printer-x](https://lucide.dev/icons/printer-x).
var PrinterX = icon.Lucide(viewBox, paint, `<path d="M12.531 22H7a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h6.377"/><path d="m16.5 16.5 5 5"/><path d="m16.5 21.5 5-5"/><path d="M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1.5"/><path d="M6 9V3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v6"/>`)

// Printer is [printer](https://lucide.dev/icons/printer).
var Printer = icon.Lucide(viewBox, paint, `<path d="M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"/><path d="M6 9V3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v6"/><rect x="6" y="14" width="12" height="8" rx="1"/>`)

// Projector is [projector](https://lucide.dev/icons/projector).
var Projector = icon.Lucide(viewBox, paint, `<path d="M5 7 3 5"/><path d="M9 6V3"/><path d="m13 7 2-2"/><circle cx="9" cy="13" r="3"/><path d="M11.83 12H20a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2.17"/><path d="M16 16h2"/>`)

// Proportions is [proportions](https://lucide.dev/icons/proportions).
var Proportions = icon.Lucide(viewBox, paint, `<rect width="20" height="16" x="2" y="4" rx="2"/><path d="M12 9v11"/><path d="M2 9h13a2 2 0 0 1 2 2v9"/>`)

// Puzzle is [puzzle](https://lucide.dev/icons/puzzle).
var Puzzle = icon.Lucide(viewBox, paint, `<path d="M15.39 4.39a1 1 0 0 0 1.68-.474 2.5 2.5 0 1 1 3.014 3.015 1 1 0 0 0-.474 1.68l1.683 1.682a2.414 2.414 0 0 1 0 3.414L19.61 15.39a1 1 0 0 1-1.68-.474 2.5 2.5 0 1 0-3.014 3.015 1 1 0 0 1 .474 1.68l-1.683 1.682a2.414 2.414 0 0 1-3.414 0L8.61 19.61a1 1 0 0 0-1.68.474 2.5 2.5 0 1 1-3.014-3.015 1 1 0 0 0 .474-1.68l-1.683-1.682a2.414 2.414 0 0 1 0-3.414L4.39 8.61a1 1 0 0 1 1.68.474 2.5 2.5 0 1 0 3.014-3.015 1 1 0 0 1-.474-1.68l1.683-1.682a2.414 2.414 0 0 1 3.414 0z"/>`)

// Pyramid is [pyramid](https://lucide.dev/icons/pyramid).
var Pyramid = icon.Lucide(viewBox, paint, `<path d="M2.5 16.88a1 1 0 0 1-.32-1.43l9-13.02a1 1 0 0 1 1.64 0l9 13.01a1 1 0 0 1-.32 1.44l-8.51 4.86a2 2 0 0 1-1.98 0Z"/><path d="M12 2v20"/>`)

// QrCode is [qr-code](https://lucide.dev/icons/qr-code).
var QrCode = icon.Lucide(viewBox, paint, `<rect width="5" height="5" x="3" y="3" rx="1"/><rect width="5" height="5" x="16" y="3" rx="1"/><rect width="5" height="5" x="3" y="16" rx="1"/><path d="M21 16h-3a2 2 0 0 0-2 2v3"/><path d="M21 21v.01"/><path d="M12 7v3a2 2 0 0 1-2 2H7"/><path d="M3 12h.01"/><path d="M12 3h.01"/><path d="M12 16v.01"/><path d="M16 12h1"/><path d="M21 12v.01"/><path d="M12 21v-1"/>`)

// Quote is [quote](https://lucide.dev/icons/quote).
var Quote = icon.Lucide(viewBox, paint, `<path d="M16 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"/><path d="M5 3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2 1 1 0 0 1 1 1v1a2 2 0 0 1-2 2 1 1 0 0 0-1 1v2a1 1 0 0 0 1 1 6 6 0 0 0 6-6V5a2 2 0 0 0-2-2z"/>`)

// Rabbit is [rabbit](https://lucide.dev/icons/rabbit).
var Rabbit = icon.Lucide(viewBox, paint, `<path d="M13 16a3 3 0 0 1 2.24 5"/><path d="M18 12h.01"/><path d="M18 21h-8a4 4 0 0 1-4-4 7 7 0 0 1 7-7h.2L9.6 6.4a1 1 0 1 1 2.8-2.8L15.8 7h.2c3.3 0 6 2.7 6 6v1a2 2 0 0 1-2 2h-1a3 3 0 0 0-3 3"/><path d="M20 8.54V4a2 2 0 1 0-4 0v3"/><path d="M7.612 12.524a3 3 0 1 0-1.6 4.3"/>`)

// Radar is [radar](https://lucide.dev/icons/radar).
var Radar = icon.Lucide(viewBox, paint, `<path d="M19.07 4.93A10 10 0 0 0 6.99 3.34"/><path d="M4 6h.01"/><path d="M2.29 9.62A10 10 0 1 0 21.31 8.35"/><path d="M16.24 7.76A6 6 0 1 0 8.23 16.67"/><path d="M12 18h.01"/><path d="M17.99 11.66A6 6 0 0 1 15.77 16.67"/><circle cx="12" cy="12" r="2"/><path d="m13.41 10.59 5.66-5.66"/>`)

// Radiation is [radiation](https://lucide.dev/icons/radiation).
var Radiation = icon.Lucide(viewBox, paint, `<path d="M12 12h.01"/><path d="M14 15.4641a4 4 0 0 1-4 0L7.52786 19.74597 A 1 1 0 0 0 7.99303 21.16211 10 10 0 0 0 16.00697 21.16211 1 1 0 0 0 16.47214 19.74597z"/><path d="M16 12a4 4 0 0 0-2-3.464l2.472-4.282a1 1 0 0 1 1.46-.305 10 10 0 0 1 4.006 6.94A1 1 0 0 1 21 12z"/><path d="M8 12a4 4 0 0 1 2-3.464L7.528 4.254a1 1 0 0 0-1.46-.305 10 10 0 0 0-4.006 6.94A1 1 0 0 0 3 12z"/>`)

// Radical is [radical](https://lucide.dev/icons/radical).
var Radical = icon.Lucide(viewBox, paint, `<path d="M3 12h3.28a1 1 0 0 1 .948.684l2.298 7.934a.5.5 0 0 0 .96-.044L13.82 4.771A1 1 0 0 1 14.792 4H21"/>`)

// RadioOff is [radio-off](https://lucide.dev/icons/radio-off).
var RadioOff = icon.Lucide(viewBox, paint, `<path d="M13.414 13.414a2 2 0 1 1-2.828-2.828"/><path d="M16.247 7.761a6 6 0 0 1 1.744 4.572"/><path d="M19.075 4.933a10 10 0 0 1 2.234 10.72"/><path d="m2 2 20 20"/><path d="M4.925 19.067a10 10 0 0 1 0-14.134"/><path d="M7.753 16.239a6 6 0 0 1 0-8.478"/>`)

// RadioReceiver is [radio-receiver](https://lucide.dev/icons/radio-receiver).
var RadioReceiver = icon.Lucide(viewBox, paint, `<path d="M5 16v2"/><path d="M19 16v2"/><rect width="20" height="8" x="2" y="8" rx="2"/><path d="M18 12h.01"/>`)

// RadioTower is [radio-tower](https://lucide.dev/icons/radio-tower).
var RadioTower = icon.Lucide(viewBox, paint, `<path d="M4.9 16.1C1 12.2 1 5.8 4.9 1.9"/><path d="M7.8 4.7a6.14 6.14 0 0 0-.8 7.5"/><circle cx="12" cy="9" r="2"/><path d="M16.2 4.8c2 2 2.26 5.11.8 7.47"/><path d="M19.1 1.9a9.96 9.96 0 0 1 0 14.1"/><path d="M9.5 18h5"/><path d="m8 22 4-11 4 11"/>`)

// Radio is [radio](https://lucide.dev/icons/radio).
var Radio = icon.Lucide(viewBox, paint, `<path d="M16.247 7.761a6 6 0 0 1 0 8.478"/><path d="M19.075 4.933a10 10 0 0 1 0 14.134"/><path d="M4.925 19.067a10 10 0 0 1 0-14.134"/><path d="M7.753 16.239a6 6 0 0 1 0-8.478"/><circle cx="12" cy="12" r="2"/>`)

// Radius is [radius](https://lucide.dev/icons/radius).
var Radius = icon.Lucide(viewBox, paint, `<path d="M20.34 17.52a10 10 0 1 0-2.82 2.82"/><circle cx="19" cy="19" r="2"/><path d="m13.41 13.41 4.18 4.18"/><circle cx="12" cy="12" r="2"/>`)

// Rainbow is [rainbow](https://lucide.dev/icons/rainbow).
var Rainbow = icon.Lucide(viewBox, paint, `<path d="M22 17a10 10 0 0 0-20 0"/><path d="M6 17a6 6 0 0 1 12 0"/><path d="M10 17a2 2 0 0 1 4 0"/>`)

// Rat is [rat](https://lucide.dev/icons/rat).
var Rat = icon.Lucide(viewBox, paint, `<path d="M13 22H4a2 2 0 0 1 0-4h12"/><path d="M13.236 18a3 3 0 0 0-2.2-5"/><path d="M16 9h.01"/><path d="M16.82 3.94a3 3 0 1 1 3.237 4.868l1.815 2.587a1.5 1.5 0 0 1-1.5 2.1l-2.872-.453a3 3 0 0 0-3.5 3"/><path d="M17 4.988a3 3 0 1 0-5.2 2.052A7 7 0 0 0 4 14.015 4 4 0 0 0 8 18"/>`)

// Ratio is [ratio](https://lucide.dev/icons/ratio).
var Ratio = icon.Lucide(viewBox, paint, `<rect width="12" height="20" x="6" y="2" rx="2"/><rect width="20" height="12" x="2" y="6" rx="2"/>`)

// ReceiptCent is [receipt-cent](https://lucide.dev/icons/receipt-cent).
var ReceiptCent = icon.Lucide(viewBox, paint, `<path d="M12 7v10"/><path d="M14.828 14.829a4 4 0 0 1-5.656 0 4 4 0 0 1 0-5.657 4 4 0 0 1 5.656 0"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/>`)

// ReceiptEuro is [receipt-euro](https://lucide.dev/icons/receipt-euro).
var ReceiptEuro = icon.Lucide(viewBox, paint, `<path d="M15.828 14.829a4 4 0 0 1-5.656 0 4 4 0 0 1 0-5.657 4 4 0 0 1 5.656 0"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/><path d="M8 12h5"/>`)

// ReceiptIndianRupee is [receipt-indian-rupee](https://lucide.dev/icons/receipt-indian-rupee).
var ReceiptIndianRupee = icon.Lucide(viewBox, paint, `<path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/><path d="M8 11h8"/><path d="M8 7h8"/><path d="M9 7a4 4 0 0 1 0 8H8l3 2"/>`)

// ReceiptJapaneseYen is [receipt-japanese-yen](https://lucide.dev/icons/receipt-japanese-yen).
var ReceiptJapaneseYen = icon.Lucide(viewBox, paint, `<path d="m12 10 3-3"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/><path d="M9 11h6"/><path d="M9 15h6"/><path d="m9 7 3 3v7"/>`)

// ReceiptPoundSterling is [receipt-pound-sterling](https://lucide.dev/icons/receipt-pound-sterling).
var ReceiptPoundSterling = icon.Lucide(viewBox, paint, `<path d="M10 17V9.5a1 1 0 0 1 5 0"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/><path d="M8 13h5"/><path d="M8 17h7"/>`)

// ReceiptRussianRuble is [receipt-russian-ruble](https://lucide.dev/icons/receipt-russian-ruble).
var ReceiptRussianRuble = icon.Lucide(viewBox, paint, `<path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/><path d="M8 11h5a2 2 0 0 0 0-4h-3v10"/><path d="M8 15h5"/>`)

// ReceiptSwissFranc is [receipt-swiss-franc](https://lucide.dev/icons/receipt-swiss-franc).
var ReceiptSwissFranc = icon.Lucide(viewBox, paint, `<path d="M10 11h4"/><path d="M10 17V7h5"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/><path d="M8 15h5"/>`)

// ReceiptText is [receipt-text](https://lucide.dev/icons/receipt-text).
var ReceiptText = icon.Lucide(viewBox, paint, `<path d="M13 16H8"/><path d="M14 8H8"/><path d="M16 12H8"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/>`)

// ReceiptTurkishLira is [receipt-turkish-lira](https://lucide.dev/icons/receipt-turkish-lira).
var ReceiptTurkishLira = icon.Lucide(viewBox, paint, `<path d="M10 7v10a5 5 0 0 0 5-5"/><path d="m14 8-6 3"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/>`)

// Receipt is [receipt](https://lucide.dev/icons/receipt).
var Receipt = icon.Lucide(viewBox, paint, `<path d="M12 17V7"/><path d="M16 8h-6a2 2 0 0 0 0 4h4a2 2 0 0 1 0 4H8"/><path d="M4 3a1 1 0 0 1 1-1 1.3 1.3 0 0 1 .7.2l.933.6a1.3 1.3 0 0 0 1.4 0l.934-.6a1.3 1.3 0 0 1 1.4 0l.933.6a1.3 1.3 0 0 0 1.4 0l.933-.6a1.3 1.3 0 0 1 1.4 0l.934.6a1.3 1.3 0 0 0 1.4 0l.933-.6A1.3 1.3 0 0 1 19 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1 1.3 1.3 0 0 1-.7-.2l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.934.6a1.3 1.3 0 0 1-1.4 0l-.933-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-1.4 0l-.934-.6a1.3 1.3 0 0 0-1.4 0l-.933.6a1.3 1.3 0 0 1-.7.2 1 1 0 0 1-1-1z"/>`)

// RectangleCircle is [rectangle-circle](https://lucide.dev/icons/rectangle-circle).
var RectangleCircle = icon.Lucide(viewBox, paint, `<path d="M14 4v16H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1z"/><circle cx="14" cy="12" r="8"/>`)

// RectangleEllipsis is [rectangle-ellipsis](https://lucide.dev/icons/rectangle-ellipsis).
var RectangleEllipsis = icon.Lucide(viewBox, paint, `<rect width="20" height="12" x="2" y="6" rx="2"/><path d="M12 12h.01"/><path d="M17 12h.01"/><path d="M7 12h.01"/>`)

// RectangleGoggles is [rectangle-goggles](https://lucide.dev/icons/rectangle-goggles).
var RectangleGoggles = icon.Lucide(viewBox, paint, `<path d="M20 6a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4a2 2 0 0 1-1.6-.8l-1.6-2.13a1 1 0 0 0-1.6 0L9.6 17.2A2 2 0 0 1 8 18H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2z"/>`)

// RectangleHorizontal is [rectangle-horizontal](https://lucide.dev/icons/rectangle-horizontal).
var RectangleHorizontal = icon.Lucide(viewBox, paint, `<rect width="20" height="12" x="2" y="6" rx="2"/>`)

// RectangleVertical is [rectangle-vertical](https://lucide.dev/icons/rectangle-vertical).
var RectangleVertical = icon.Lucide(viewBox, paint, `<rect width="12" height="20" x="6" y="2" rx="2"/>`)

// Recycle is [recycle](https://lucide.dev/icons/recycle).
var Recycle = icon.Lucide(viewBox, paint, `<path d="M7 19H4.815a1.83 1.83 0 0 1-1.57-.881 1.785 1.785 0 0 1-.004-1.784L7.196 9.5"/><path d="M11 19h8.203a1.83 1.83 0 0 0 1.556-.89 1.784 1.784 0 0 0 0-1.775l-1.226-2.12"/><path d="m14 16-3 3 3 3"/><path d="M8.293 13.596 7.196 9.5 3.1 10.598"/><path d="m9.344 5.811 1.093-1.892A1.83 1.83 0 0 1 11.985 3a1.784 1.784 0 0 1 1.546.888l3.943 6.843"/><path d="m13.378 9.633 4.096 1.098 1.097-4.096"/>`)

// Redo2 is [redo-2](https://lucide.dev/icons/redo-2).
var Redo2 = icon.Lucide(viewBox, paint, `<path d="m15 14 5-5-5-5"/><path d="M20 9H9.5A5.5 5.5 0 0 0 4 14.5A5.5 5.5 0 0 0 9.5 20H13"/>`)

// RedoDot is [redo-dot](https://lucide.dev/icons/redo-dot).
var RedoDot = icon.Lucide(viewBox, paint, `<circle cx="12" cy="17" r="1"/><path d="M21 7v6h-6"/><path d="M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7"/>`)

// Redo is [redo](https://lucide.dev/icons/redo).
var Redo = icon.Lucide(viewBox, paint, `<path d="M21 7v6h-6"/><path d="M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7"/>`)

// RefreshCcwDot is [refresh-ccw-dot](https://lucide.dev/icons/refresh-ccw-dot).
var RefreshCcwDot = icon.Lucide(viewBox, paint, `<path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5"/><path d="M3 12a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"/><path d="M16 16h5v5"/><circle cx="12" cy="12" r="1"/>`)

// RefreshCcw is [refresh-ccw](https://lucide.dev/icons/refresh-ccw).
var RefreshCcw = icon.Lucide(viewBox, paint, `<path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5"/><path d="M3 12a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"/><path d="M16 16h5v5"/>`)

// RefreshCwOff is [refresh-cw-off](https://lucide.dev/icons/refresh-cw-off).
var RefreshCwOff = icon.Lucide(viewBox, paint, `<path d="M21 8L18.74 5.74A9.75 9.75 0 0 0 12 3C11 3 10.03 3.16 9.13 3.47"/><path d="M8 16H3v5"/><path d="M3 12C3 9.51 4 7.26 5.64 5.64"/><path d="m3 16 2.26 2.26A9.75 9.75 0 0 0 12 21c2.49 0 4.74-1 6.36-2.64"/><path d="M21 12c0 1-.16 1.97-.47 2.87"/><path d="M21 3v5h-5"/><path d="M22 22 2 2"/>`)

// RefreshCw is [refresh-cw](https://lucide.dev/icons/refresh-cw).
var RefreshCw = icon.Lucide(viewBox, paint, `<path d="M3 12a9 9 0 0 1 9-9 9.75 9.75 0 0 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/><path d="M21 12a9 9 0 0 1-9 9 9.75 9.75 0 0 1-6.74-2.74L3 16"/><path d="M8 16H3v5"/>`)

// Refrigerator is [refrigerator](https://lucide.dev/icons/refrigerator).
var Refrigerator = icon.Lucide(viewBox, paint, `<path d="M5 6a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6Z"/><path d="M5 10h14"/><path d="M15 7v6"/>`)

// Regex is [regex](https://lucide.dev/icons/regex).
var Regex = icon.Lucide(viewBox, paint, `<path d="M17 3v10"/><path d="m12.67 5.5 8.66 5"/><path d="m12.67 10.5 8.66-5"/><path d="M9 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2z"/>`)

// RemoveFormatting is [remove-formatting](https://lucide.dev/icons/remove-formatting).
var RemoveFormatting = icon.Lucide(viewBox, paint, `<path d="M4 7V4h16v3"/><path d="M5 20h6"/><path d="M13 4 8 20"/><path d="m15 15 5 5"/><path d="m20 15-5 5"/>`)

// Repeat1 is [repeat-1](https://lucide.dev/icons/repeat-1).
var Repeat1 = icon.Lucide(viewBox, paint, `<path d="m17 2 4 4-4 4"/><path d="M3 11v-1a4 4 0 0 1 4-4h14"/><path d="m7 22-4-4 4-4"/><path d="M21 13v1a4 4 0 0 1-4 4H3"/><path d="M11 10h1v4"/>`)

// Repeat2 is [repeat-2](https://lucide.dev/icons/repeat-2).
var Repeat2 = icon.Lucide(viewBox, paint, `<path d="m2 9 3-3 3 3"/><path d="M13 18H7a2 2 0 0 1-2-2V6"/><path d="m22 15-3 3-3-3"/><path d="M11 6h6a2 2 0 0 1 2 2v10"/>`)

// RepeatOff is [repeat-off](https://lucide.dev/icons/repeat-off).
var RepeatOff = icon.Lucide(viewBox, paint, `<path d="M11.656 6H21l-4-4"/><path d="M17.898 17.898A4 4 0 0 1 17 18H3l4-4"/><path d="m2 2 20 20"/><path d="M21 13v1a4 4 0 0 1-.171 1.159"/><path d="m21 6-4 4"/><path d="M3 11v-1a4 4 0 0 1 3.102-3.898"/><path d="m7 22-4-4"/>`)

// Repeat is [repeat](https://lucide.dev/icons/repeat).
var Repeat = icon.Lucide(viewBox, paint, `<path d="m17 2 4 4-4 4"/><path d="M3 11v-1a4 4 0 0 1 4-4h14"/><path d="m7 22-4-4 4-4"/><path d="M21 13v1a4 4 0 0 1-4 4H3"/>`)

// ReplaceAll is [replace-all](https://lucide.dev/icons/replace-all).
var ReplaceAll = icon.Lucide(viewBox, paint, `<path d="M14 14a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1"/><path d="M14 4a1 1 0 0 1 1-1"/><path d="M15 10a1 1 0 0 1-1-1"/><path d="M19 14a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1"/><path d="M21 4a1 1 0 0 0-1-1"/><path d="M21 9a1 1 0 0 1-1 1"/><path d="m3 7 3 3 3-3"/><path d="M6 10V5a2 2 0 0 1 2-2h2"/><rect x="3" y="14" width="7" height="7" rx="1"/>`)

// Replace is [replace](https://lucide.dev/icons/replace).
var Replace = icon.Lucide(viewBox, paint, `<path d="M14 4a1 1 0 0 1 1-1"/><path d="M15 10a1 1 0 0 1-1-1"/><path d="M21 4a1 1 0 0 0-1-1"/><path d="M21 9a1 1 0 0 1-1 1"/><path d="m3 7 3 3 3-3"/><path d="M6 10V5a2 2 0 0 1 2-2h2"/><rect x="3" y="14" width="7" height="7" rx="1"/>`)

// ReplyAll is [reply-all](https://lucide.dev/icons/reply-all).
var ReplyAll = icon.Lucide(viewBox, paint, `<path d="m12 17-5-5 5-5"/><path d="M22 18v-2a4 4 0 0 0-4-4H7"/><path d="m7 17-5-5 5-5"/>`)

// Reply is [reply](https://lucide.dev/icons/reply).
var Reply = icon.Lucide(viewBox, paint, `<path d="M20 18v-2a4 4 0 0 0-4-4H4"/><path d="m9 17-5-5 5-5"/>`)

// Rewind is [rewind](https://lucide.dev/icons/rewind).
var Rewind = icon.Lucide(viewBox, paint, `<path d="M12 6a2 2 0 0 0-3.414-1.414l-6 6a2 2 0 0 0 0 2.828l6 6A2 2 0 0 0 12 18z"/><path d="M22 6a2 2 0 0 0-3.414-1.414l-6 6a2 2 0 0 0 0 2.828l6 6A2 2 0 0 0 22 18z"/>`)

// Ribbon is [ribbon](https://lucide.dev/icons/ribbon).
var Ribbon = icon.Lucide(viewBox, paint, `<path d="M12 11.22C11 9.997 10 9 10 8a2 2 0 0 1 4 0c0 1-.998 2.002-2.01 3.22"/><path d="m12 18 2.57-3.5"/><path d="M6.243 9.016a7 7 0 0 1 11.507-.009"/><path d="M9.35 14.53 12 11.22"/><path d="M9.35 14.53C7.728 12.246 6 10.221 6 7a6 5 0 0 1 12 0c-.005 3.22-1.778 5.235-3.43 7.5l3.557 4.527a1 1 0 0 1-.203 1.43l-1.894 1.36a1 1 0 0 1-1.384-.215L12 18l-2.679 3.593a1 1 0 0 1-1.39.213l-1.865-1.353a1 1 0 0 1-.203-1.422z"/>`)

// Road is [road](https://lucide.dev/icons/road).
var Road = icon.Lucide(viewBox, paint, `<path d="M12 17v4"/><path d="M12 5V3"/><path d="M12 9v3"/><path d="M2.077 18.449A2 2 0 0 0 4 21h16a2 2 0 0 0 1.924-2.55l-4-14A2 2 0 0 0 16 3H8a2 2 0 0 0-1.924 1.45z"/>`)

// RobotArm is [robot-arm](https://lucide.dev/icons/robot-arm).
var RobotArm = icon.Lucide(viewBox, paint, `<path d="M12 21 7.5 8.322"/><path d="m14 7 1.75-3.767a.5.5 0 0 1 .662-.172L20 5.005"/><path d="m20 8.998-3.588 1.944a.5.5 0 0 1-.662-.172L14 7H8"/><path d="M3.486 21h10"/><path d="M5 21V8.732"/><circle cx="6" cy="7" r="2"/>`)

// RobotVacuum is [robot-vacuum](https://lucide.dev/icons/robot-vacuum).
var RobotVacuum = icon.Lucide(viewBox, paint, `<path d="M11 17h2"/><path d="M12 12h.01"/><path d="M17 12a5 5 0 00-10 0"/><path d="M19 2v2.8"/><path d="M2 5h2.8"/><path d="M22 5h-2.8"/><path d="M5 2v2.8"/><circle cx="12" cy="12" r="10"/>`)

// Rocket is [rocket](https://lucide.dev/icons/rocket).
var Rocket = icon.Lucide(viewBox, paint, `<path d="M12 15v5s3.03-.55 4-2c1.08-1.62 0-5 0-5"/><path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09"/><path d="M9 12a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.4 22.4 0 0 1-4 2z"/><path d="M9 12H4s.55-3.03 2-4c1.62-1.08 5 .05 5 .05"/>`)

// RockingChair is [rocking-chair](https://lucide.dev/icons/rocking-chair).
var RockingChair = icon.Lucide(viewBox, paint, `<path d="m15 13 3.708 7.416"/><path d="M3 19a15 15 0 0 0 18 0"/><path d="m3 2 3.21 9.633A2 2 0 0 0 8.109 13H18"/><path d="m9 13-3.708 7.416"/>`)

// RollerCoaster is [roller-coaster](https://lucide.dev/icons/roller-coaster).
var RollerCoaster = icon.Lucide(viewBox, paint, `<path d="M6 19V5"/><path d="M10 19V6.8"/><path d="M14 19v-7.8"/><path d="M18 5v4"/><path d="M18 19v-6"/><path d="M22 19V9"/><path d="M2 19V9a4 4 0 0 1 4-4c2 0 4 1.33 6 4s4 4 6 4a4 4 0 1 0-3-6.65"/>`)

// Rose is [rose](https://lucide.dev/icons/rose).
var Rose = icon.Lucide(viewBox, paint, `<path d="M17 10h-1a4 4 0 1 1 4-4v.534"/><path d="M17 6h1a4 4 0 0 1 1.42 7.74l-2.29.87a6 6 0 0 1-5.339-10.68l2.069-1.31"/><path d="M4.5 17c2.8-.5 4.4 0 5.5.8s1.8 2.2 2.3 3.7c-2 .4-3.5.4-4.8-.3-1.2-.6-2.3-1.9-3-4.2"/><path d="M9.77 12C4 15 2 22 2 22"/><circle cx="17" cy="8" r="2"/>`)

// Rotate3d is [rotate-3d](https://lucide.dev/icons/rotate-3d).
var Rotate3d = icon.Lucide(viewBox, paint, `<path d="m15.194 13.707 3.814 1.86-1.86 3.814"/><path d="M16.47214 7.52786 A 5 10 0 1 0 13 21.79796"/><path d="M21.79796 11 A 10 5 0 1 0 19 15.57071"/>`)

// RotateCcwClock is [rotate-ccw-clock](https://lucide.dev/icons/rotate-ccw-clock).
var RotateCcwClock = icon.Lucide(viewBox, paint, `<path d="M3 12a9 9 0 1 0 9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5"/><path d="M12 7v5l4 2"/>`)

// RotateCcwKey is [rotate-ccw-key](https://lucide.dev/icons/rotate-ccw-key).
var RotateCcwKey = icon.Lucide(viewBox, paint, `<path d="M12 7v6"/><path d="M12 9h2"/><path d="M3 12a9 9 0 1 0 9-9 9.74 9.74 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5"/><circle cx="12" cy="15" r="2"/>`)

// RotateCcwSquare is [rotate-ccw-square](https://lucide.dev/icons/rotate-ccw-square).
var RotateCcwSquare = icon.Lucide(viewBox, paint, `<path d="M20 9V7a2 2 0 0 0-2-2h-6"/><path d="m15 2-3 3 3 3"/><path d="M20 13v5a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2"/>`)

// RotateCcw is [rotate-ccw](https://lucide.dev/icons/rotate-ccw).
var RotateCcw = icon.Lucide(viewBox, paint, `<path d="M3 12a9 9 0 1 0 9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5"/>`)

// RotateCwClock is [rotate-cw-clock](https://lucide.dev/icons/rotate-cw-clock).
var RotateCwClock = icon.Lucide(viewBox, paint, `<path d="M12 7v5l4 2"/><path d="M16 8h5V3"/><path d="m21 8-2.3-2.3A9.7 9.7 0 0012 3a9 9 0 109 9"/>`)

// RotateCwFadingClock is [rotate-cw-fading-clock](https://lucide.dev/icons/rotate-cw-fading-clock).
var RotateCwFadingClock = icon.Lucide(viewBox, paint, `<path d="M12 3a9.75 9.75 0 0 1 6.74 2.74"/><path d="M18.74 5.74 21 8"/><path d="M21 8V3"/><path d="M7.5 19.794c-6-3.464-6-12.124 0-15.588"/><path d="M7.5 4.206A9 9 0 0 1 12 3"/><path d="M12 7v5l4 2"/><path d="M14 20.775A9 9 0 0 1 12 21"/><path d="M19 17.656a9 9 0 0 1-1.5 1.456"/><path d="M21 12a9 9 0 0 1-.228 2"/><path d="M21 8h-5"/>`)

// RotateCwSquare is [rotate-cw-square](https://lucide.dev/icons/rotate-cw-square).
var RotateCwSquare = icon.Lucide(viewBox, paint, `<path d="M12 5H6a2 2 0 0 0-2 2v3"/><path d="m9 8 3-3-3-3"/><path d="M4 14v4a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2"/>`)

// RotateCw is [rotate-cw](https://lucide.dev/icons/rotate-cw).
var RotateCw = icon.Lucide(viewBox, paint, `<path d="M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/>`)

// RouteOff is [route-off](https://lucide.dev/icons/route-off).
var RouteOff = icon.Lucide(viewBox, paint, `<circle cx="6" cy="19" r="3"/><path d="M9 19h8.5c.4 0 .9-.1 1.3-.2"/><path d="M5.2 5.2A3.5 3.53 0 0 0 6.5 12H12"/><path d="m2 2 20 20"/><path d="M21 15.3a3.5 3.5 0 0 0-3.3-3.3"/><path d="M15 5h-4.3"/><circle cx="18" cy="5" r="3"/>`)

// Route is [route](https://lucide.dev/icons/route).
var Route = icon.Lucide(viewBox, paint, `<circle cx="6" cy="19" r="3"/><path d="M9 19h8.5a3.5 3.5 0 0 0 0-7h-11a3.5 3.5 0 0 1 0-7H15"/><circle cx="18" cy="5" r="3"/>`)

// Router is [router](https://lucide.dev/icons/router).
var Router = icon.Lucide(viewBox, paint, `<rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6.01 18H6"/><path d="M10.01 18H10"/><path d="M15 10v4"/><path d="M17.84 7.17a4 4 0 0 0-5.66 0"/><path d="M20.66 4.34a8 8 0 0 0-11.31 0"/>`)

// Rows2 is [rows-2](https://lucide.dev/icons/rows-2).
var Rows2 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 12h18"/>`)

// Rows3 is [rows-3](https://lucide.dev/icons/rows-3).
var Rows3 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M21 9H3"/><path d="M21 15H3"/>`)

// Rows4 is [rows-4](https://lucide.dev/icons/rows-4).
var Rows4 = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M21 7.5H3"/><path d="M21 12H3"/><path d="M21 16.5H3"/>`)

// Rss is [rss](https://lucide.dev/icons/rss).
var Rss = icon.Lucide(viewBox, paint, `<path d="M4 11a9 9 0 0 1 9 9"/><path d="M4 4a16 16 0 0 1 16 16"/><circle cx="5" cy="19" r="1"/>`)

// RugbyBall is [rugby-ball](https://lucide.dev/icons/rugby-ball).
var RugbyBall = icon.Lucide(viewBox, paint, `<path d="m10 10 4 4"/><path d="m13 7 4 4"/><path d="M15.34 2.138A15 15 0 002.138 15.34c-.357 2.94.004 4.919.805 5.717.798.8 2.778 1.162 5.718.805A15 15 0 0021.862 8.661c.357-2.94-.004-4.92-.805-5.718-.798-.8-2.778-1.162-5.717-.805"/><path d="M17 7 7 17"/><path d="m7 13 4 4"/>`)

// RulerDimensionLine is [ruler-dimension-line](https://lucide.dev/icons/ruler-dimension-line).
var RulerDimensionLine = icon.Lucide(viewBox, paint, `<path d="M10 15v-3"/><path d="M14 15v-3"/><path d="M18 15v-3"/><path d="M2 8V4"/><path d="M22 6H2"/><path d="M22 8V4"/><path d="M6 15v-3"/><rect x="2" y="12" width="20" height="8" rx="2"/>`)

// Ruler is [ruler](https://lucide.dev/icons/ruler).
var Ruler = icon.Lucide(viewBox, paint, `<path d="M21.3 15.3a2.4 2.4 0 0 1 0 3.4l-2.6 2.6a2.4 2.4 0 0 1-3.4 0L2.7 8.7a2.41 2.41 0 0 1 0-3.4l2.6-2.6a2.41 2.41 0 0 1 3.4 0Z"/><path d="m14.5 12.5 2-2"/><path d="m11.5 9.5 2-2"/><path d="m8.5 6.5 2-2"/><path d="m17.5 15.5 2-2"/>`)

// RussianRuble is [russian-ruble](https://lucide.dev/icons/russian-ruble).
var RussianRuble = icon.Lucide(viewBox, paint, `<path d="M6 11h8a4 4 0 0 0 0-8H9v18"/><path d="M6 15h8"/>`)

// Sailboat is [sailboat](https://lucide.dev/icons/sailboat).
var Sailboat = icon.Lucide(viewBox, paint, `<path d="M10 2v15"/><path d="M7 22a4 4 0 0 1-4-4 1 1 0 0 1 1-1h16a1 1 0 0 1 1 1 4 4 0 0 1-4 4z"/><path d="M9.159 2.46a1 1 0 0 1 1.521-.193l9.977 8.98A1 1 0 0 1 20 13H4a1 1 0 0 1-.824-1.567z"/>`)

// Salad is [salad](https://lucide.dev/icons/salad).
var Salad = icon.Lucide(viewBox, paint, `<path d="M19.496 12a2.5 2.5 0 00.399-2.214A2 2 0 0020.32 6.5 2 2 0 0019 3a2 2 0 00-1.5.68 2 2 0 00-3.287.424 2.5 2.5 0 00-3.189 2.06A3 3 0 0012 12l4-4"/><path d="M4 12a1 1 0 00-.99 1.133A9 9 0 0012 21a9 9 0 008.99-7.867A1 1 0 0020 12z"/><path d="M7 21h10"/><path d="M9.85 6.907A3.5 3.5 0 005.05 12"/>`)

// Sandwich is [sandwich](https://lucide.dev/icons/sandwich).
var Sandwich = icon.Lucide(viewBox, paint, `<path d="m2.37 11.223 8.372-6.777a2 2 0 0 1 2.516 0l8.371 6.777"/><path d="M21 15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-5.25"/><path d="M3 15a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h9"/><path d="m6.67 15 6.13 4.6a2 2 0 0 0 2.8-.4l3.15-4.2"/><rect width="20" height="4" x="2" y="11" rx="1"/>`)

// SatelliteDish is [satellite-dish](https://lucide.dev/icons/satellite-dish).
var SatelliteDish = icon.Lucide(viewBox, paint, `<path d="M18 12a6 6 0 00-6-6"/><path d="M2.824 10.459a8 8 0 0010.717 10.717c.558-.276.623-1.012.183-1.452l-9.448-9.448c-.44-.44-1.176-.375-1.452.183"/><path d="M22 12A10 10 0 0012 2"/><path d="m9 15 4-4"/>`)

// Satellite is [satellite](https://lucide.dev/icons/satellite).
var Satellite = icon.Lucide(viewBox, paint, `<path d="m13.5 6.5-3.148-3.148a1.205 1.205 0 0 0-1.704 0L6.352 5.648a1.205 1.205 0 0 0 0 1.704L9.5 10.5"/><path d="M16.5 7.5 19 5"/><path d="m17.5 10.5 3.148 3.148a1.205 1.205 0 0 1 0 1.704l-2.296 2.296a1.205 1.205 0 0 1-1.704 0L13.5 14.5"/><path d="M9 21a6 6 0 0 0-6-6"/><path d="M9.352 10.648a1.205 1.205 0 0 0 0 1.704l2.296 2.296a1.205 1.205 0 0 0 1.704 0l4.296-4.296a1.205 1.205 0 0 0 0-1.704l-2.296-2.296a1.205 1.205 0 0 0-1.704 0z"/>`)

// SaudiRiyal is [saudi-riyal](https://lucide.dev/icons/saudi-riyal).
var SaudiRiyal = icon.Lucide(viewBox, paint, `<path d="m20 19.5-5.5 1.2"/><path d="M14.5 4v11.22a1 1 0 0 0 1.242.97L20 15.2"/><path d="m2.978 19.351 5.549-1.363A2 2 0 0 0 10 16V2"/><path d="M20 10 4 13.5"/>`)

// SaveAll is [save-all](https://lucide.dev/icons/save-all).
var SaveAll = icon.Lucide(viewBox, paint, `<path d="M10 2v3a1 1 0 0 0 1 1h5"/><path d="M18 18v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6"/><path d="M18 22H4a2 2 0 0 1-2-2V6"/><path d="M8 18a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9.172a2 2 0 0 1 1.414.586l2.828 2.828A2 2 0 0 1 22 6.828V16a2 2 0 0 1-2.01 2z"/>`)

// SaveCheck is [save-check](https://lucide.dev/icons/save-check).
var SaveCheck = icon.Lucide(viewBox, paint, `<path d="M12.5 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h10.2a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4v4.35"/><path d="m16 19 2 2 4-4"/><path d="M17 15.13V14a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7"/><path d="M7 3v4a1 1 0 0 0 1 1h7"/>`)

// SaveOff is [save-off](https://lucide.dev/icons/save-off).
var SaveOff = icon.Lucide(viewBox, paint, `<path d="M13 13H8a1 1 0 0 0-1 1v7"/><path d="M14 8h1"/><path d="M17 21v-4"/><path d="m2 2 20 20"/><path d="M20.41 20.41A2 2 0 0 1 19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 .59-1.41"/><path d="M9 3h6.2a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V15"/>`)

// SavePen is [save-pen](https://lucide.dev/icons/save-pen).
var SavePen = icon.Lucide(viewBox, paint, `<path d="M13.33 13H8a1 1 0 00-1 1v7"/><path d="M14.363 17.634a2 2 0 00-.506.854l-.837 2.87a.5.5 0 00.62.62l2.87-.837a2 2 0 00.854-.506l4.013-4.009a1 1 0 10-3.004-3.004z"/><path d="M7 3v4a1 1 0 001 1h7"/><path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h10.2a2 2 0 011.4.6l3.8 3.8a2 2 0 01.6 1.4v.3"/>`)

// SavePlus is [save-plus](https://lucide.dev/icons/save-plus).
var SavePlus = icon.Lucide(viewBox, paint, `<path d="M12.5 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h10.2a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V12"/><path d="M16 13H8a1 1 0 0 0-1 1v7"/><path d="M19 22v-6"/><path d="M22 19h-6"/><path d="M7 3v4a1 1 0 0 0 1 1h7"/>`)

// Save is [save](https://lucide.dev/icons/save).
var Save = icon.Lucide(viewBox, paint, `<path d="M15.2 3a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z"/><path d="M17 21v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7"/><path d="M7 3v4a1 1 0 0 0 1 1h7"/>`)

// Scale3d is [scale-3d](https://lucide.dev/icons/scale-3d).
var Scale3d = icon.Lucide(viewBox, paint, `<path d="M5 7v11a1 1 0 0 0 1 1h11"/><path d="M5.293 18.707 11 13"/><circle cx="19" cy="19" r="2"/><circle cx="5" cy="5" r="2"/>`)

// Scale is [scale](https://lucide.dev/icons/scale).
var Scale = icon.Lucide(viewBox, paint, `<path d="M12 3v18"/><path d="m19 8 3 8a5 5 0 0 1-6 0zV7"/><path d="M3 7h1a17 17 0 0 0 8-2 17 17 0 0 0 8 2h1"/><path d="m5 8 3 8a5 5 0 0 1-6 0zV7"/><path d="M7 21h10"/>`)

// Scaling is [scaling](https://lucide.dev/icons/scaling).
var Scaling = icon.Lucide(viewBox, paint, `<path d="M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M14 15H9v-5"/><path d="M16 3h5v5"/><path d="M21 3 9 15"/>`)

// ScanBarcode is [scan-barcode](https://lucide.dev/icons/scan-barcode).
var ScanBarcode = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><path d="M8 7v10"/><path d="M12 7v10"/><path d="M17 7v10"/>`)

// ScanBox is [scan-box](https://lucide.dev/icons/scan-box).
var ScanBox = icon.Lucide(viewBox, paint, `<path d="M12 12v5.5"/><path d="M17 3h2a2 2 0 012 2v2"/><path d="M21 17v2a2 2 0 01-2 2h-2"/><path d="M3 7V5a2 2 0 012-2h2"/><path d="M7 21H5a2 2 0 01-2-2v-2"/><path d="M7.264 9.252 12 12l4.737-2.748"/><path d="M7.995 8.514A2 2 0 007 10.244v3.516a2 2 0 00.996 1.73l3 1.74a2 2 0 002.008 0l3-1.74A2 2 0 0017 13.76v-3.517a2 2 0 00-.995-1.73l-3-1.742a2 2 0 00-1.892-.064z"/>`)

// ScanEye is [scan-eye](https://lucide.dev/icons/scan-eye).
var ScanEye = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><circle cx="12" cy="12" r="1"/><path d="M18.944 12.33a1 1 0 0 0 0-.66 7.5 7.5 0 0 0-13.888 0 1 1 0 0 0 0 .66 7.5 7.5 0 0 0 13.888 0"/>`)

// ScanFace is [scan-face](https://lucide.dev/icons/scan-face).
var ScanFace = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><path d="M8 14s1.5 2 4 2 4-2 4-2"/><path d="M9 9h.01"/><path d="M15 9h.01"/>`)

// ScanHeart is [scan-heart](https://lucide.dev/icons/scan-heart).
var ScanHeart = icon.Lucide(viewBox, paint, `<path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><path d="M7.828 13.07A3 3 0 0 1 12 8.764a3 3 0 0 1 4.172 4.306l-3.447 3.62a1 1 0 0 1-1.449 0z"/>`)

// ScanLine is [scan-line](https://lucide.dev/icons/scan-line).
var ScanLine = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><path d="M7 12h10"/>`)

// ScanQrCode is [scan-qr-code](https://lucide.dev/icons/scan-qr-code).
var ScanQrCode = icon.Lucide(viewBox, paint, `<path d="M17 12v4a1 1 0 0 1-1 1h-4"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M17 8V7"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M7 17h.01"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><rect x="7" y="7" width="5" height="5" rx="1"/>`)

// ScanSearch is [scan-search](https://lucide.dev/icons/scan-search).
var ScanSearch = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><circle cx="12" cy="12" r="3"/><path d="m16 16-1.9-1.9"/>`)

// ScanSquare is [scan-square](https://lucide.dev/icons/scan-square).
var ScanSquare = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><rect width="8" height="8" x="8" y="8" rx="1"/>`)

// ScanText is [scan-text](https://lucide.dev/icons/scan-text).
var ScanText = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/><path d="M7 8h8"/><path d="M7 12h10"/><path d="M7 16h6"/>`)

// Scan is [scan](https://lucide.dev/icons/scan).
var Scan = icon.Lucide(viewBox, paint, `<path d="M3 7V5a2 2 0 0 1 2-2h2"/><path d="M17 3h2a2 2 0 0 1 2 2v2"/><path d="M21 17v2a2 2 0 0 1-2 2h-2"/><path d="M7 21H5a2 2 0 0 1-2-2v-2"/>`)

// School is [school](https://lucide.dev/icons/school).
var School = icon.Lucide(viewBox, paint, `<path d="M14 21v-3a2 2 0 0 0-4 0v3"/><path d="M18 4.933V21"/><path d="m4 6 7.106-3.79a2 2 0 0 1 1.788 0L20 6"/><path d="m6 11-3.52 2.147a1 1 0 0 0-.48.854V19a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5a1 1 0 0 0-.48-.853L18 11"/><path d="M6 4.933V21"/><circle cx="12" cy="9" r="2"/>`)

// ScissorsLineDashed is [scissors-line-dashed](https://lucide.dev/icons/scissors-line-dashed).
var ScissorsLineDashed = icon.Lucide(viewBox, paint, `<path d="M5.42 9.42 8 12"/><circle cx="4" cy="8" r="2"/><path d="m14 6-8.58 8.58"/><circle cx="4" cy="16" r="2"/><path d="M10.8 14.8 14 18"/><path d="M16 12h-2"/><path d="M22 12h-2"/>`)

// Scissors is [scissors](https://lucide.dev/icons/scissors).
var Scissors = icon.Lucide(viewBox, paint, `<circle cx="6" cy="6" r="3"/><path d="M8.12 8.12 12 12"/><path d="M20 4 8.12 15.88"/><circle cx="6" cy="18" r="3"/><path d="M14.8 14.8 20 20"/>`)

// Scooter is [scooter](https://lucide.dev/icons/scooter).
var Scooter = icon.Lucide(viewBox, paint, `<path d="M21 4h-3.5l2 11.05"/><path d="M6.95 17h5.142c.523 0 .95-.406 1.063-.916a6.5 6.5 0 0 1 5.345-5.009"/><circle cx="19.5" cy="17.5" r="2.5"/><circle cx="4.5" cy="17.5" r="2.5"/>`)

// ScratchBlocks is [scratch-blocks](https://lucide.dev/icons/scratch-blocks).
var ScratchBlocks = icon.Lucide(viewBox, paint, `<path d="M19 5a2 2 0 012 2v11a2 2 0 01-2 2h-5.586a1 1 0 00-.707.293l-1.414 1.414a1 1 0 01-.707.293H8.414a1 1 0 01-.707-.293l-1.414-1.414A1 1 0 005.586 20H5a2 2 0 01-2-2V5.286c0-.394.11-.785.36-1.09a6 6 0 019.168-.133C12.996 4.6 13.637 5 14.35 5z"/><path d="M21 12h-7.586a1 1 0 00-.707.293l-1.414 1.414a1 1 0 01-.707.293H8.414a1 1 0 01-.707-.293l-1.414-1.414A1 1 0 005.586 12H3"/>`)

// ScreenShareOff is [screen-share-off](https://lucide.dev/icons/screen-share-off).
var ScreenShareOff = icon.Lucide(viewBox, paint, `<path d="M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3"/><path d="M8 21h8"/><path d="M12 17v4"/><path d="m22 3-5 5"/><path d="m17 3 5 5"/>`)

// ScreenShare is [screen-share](https://lucide.dev/icons/screen-share).
var ScreenShare = icon.Lucide(viewBox, paint, `<path d="M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3"/><path d="M8 21h8"/><path d="M12 17v4"/><path d="m17 8 5-5"/><path d="M17 3h5v5"/>`)

// Screw is [screw](https://lucide.dev/icons/screw).
var Screw = icon.Lucide(viewBox, paint, `<path d="M10 9v10"/><path d="M14 13v2"/><path d="M16.5 10.5 5.706 21.294A2.4 2.4 0 014.002 22H2.5a.5.5 0 01-.5-.5v-1.502a2.4 2.4 0 01.706-1.704L13.5 7.5"/><path d="M20.819 11.575a6 6 0 00.969-5.162.6.6 0 00-.981-.219l-.455.454a1.2 1.2 0 01-1.704 0l-1.296-1.296a1.2 1.2 0 010-1.704l.454-.453a.595.595 0 00-.219-.98 6 6 0 00-5.162.967 2 2 0 00-.222 3.019l5.596 5.597a2 2 0 003.02-.222"/><path d="M6 13v8"/>`)

// ScrollText is [scroll-text](https://lucide.dev/icons/scroll-text).
var ScrollText = icon.Lucide(viewBox, paint, `<path d="M15 12h-5"/><path d="M15 8h-5"/><path d="M19 17V5a2 2 0 0 0-2-2H4"/><path d="M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3"/>`)

// Scroll is [scroll](https://lucide.dev/icons/scroll).
var Scroll = icon.Lucide(viewBox, paint, `<path d="M19 17V5a2 2 0 0 0-2-2H4"/><path d="M8 21h12a2 2 0 0 0 2-2v-1a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v1a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v2a1 1 0 0 0 1 1h3"/>`)

// SearchAlert is [search-alert](https://lucide.dev/icons/search-alert).
var SearchAlert = icon.Lucide(viewBox, paint, `<circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/><path d="M11 7v4"/><path d="M11 15h.01"/>`)

// SearchCheck is [search-check](https://lucide.dev/icons/search-check).
var SearchCheck = icon.Lucide(viewBox, paint, `<path d="m8 11 2 2 4-4"/><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/>`)

// SearchCode is [search-code](https://lucide.dev/icons/search-code).
var SearchCode = icon.Lucide(viewBox, paint, `<path d="m13 13.5 2-2.5-2-2.5"/><path d="m21 21-4.3-4.3"/><path d="M9 8.5 7 11l2 2.5"/><circle cx="11" cy="11" r="8"/>`)

// SearchSlash is [search-slash](https://lucide.dev/icons/search-slash).
var SearchSlash = icon.Lucide(viewBox, paint, `<path d="m13.5 8.5-5 5"/><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/>`)

// SearchX is [search-x](https://lucide.dev/icons/search-x).
var SearchX = icon.Lucide(viewBox, paint, `<path d="m13.5 8.5-5 5"/><path d="m8.5 8.5 5 5"/><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/>`)

// Search is [search](https://lucide.dev/icons/search).
var Search = icon.Lucide(viewBox, paint, `<path d="m21 21-4.34-4.34"/><circle cx="11" cy="11" r="8"/>`)

// Section is [section](https://lucide.dev/icons/section).
var Section = icon.Lucide(viewBox, paint, `<path d="M16 5a4 3 0 0 0-8 0c0 4 8 3 8 7a4 3 0 0 1-8 0"/><path d="M8 19a4 3 0 0 0 8 0c0-4-8-3-8-7a4 3 0 0 1 8 0"/>`)

// SendHorizontal is [send-horizontal](https://lucide.dev/icons/send-horizontal).
var SendHorizontal = icon.Lucide(viewBox, paint, `<path d="M3.714 3.048a.498.498 0 0 0-.683.627l2.843 7.627a2 2 0 0 1 0 1.396l-2.842 7.627a.498.498 0 0 0 .682.627l18-8.5a.5.5 0 0 0 0-.904z"/><path d="M6 12h16"/>`)

// SendToBack is [send-to-back](https://lucide.dev/icons/send-to-back).
var SendToBack = icon.Lucide(viewBox, paint, `<rect x="14" y="14" width="8" height="8" rx="2"/><rect x="2" y="2" width="8" height="8" rx="2"/><path d="M7 14v1a2 2 0 0 0 2 2h1"/><path d="M14 7h1a2 2 0 0 1 2 2v1"/>`)

// Send is [send](https://lucide.dev/icons/send).
var Send = icon.Lucide(viewBox, paint, `<path d="M14.536 21.686a.5.5 0 0 0 .937-.024l6.5-19a.496.496 0 0 0-.635-.635l-19 6.5a.5.5 0 0 0-.024.937l7.93 3.18a2 2 0 0 1 1.112 1.11z"/><path d="m21.854 2.147-10.94 10.939"/>`)

// SeparatorHorizontal is [separator-horizontal](https://lucide.dev/icons/separator-horizontal).
var SeparatorHorizontal = icon.Lucide(viewBox, paint, `<path d="m16 16-4 4-4-4"/><path d="M3 12h18"/><path d="m8 8 4-4 4 4"/>`)

// SeparatorVertical is [separator-vertical](https://lucide.dev/icons/separator-vertical).
var SeparatorVertical = icon.Lucide(viewBox, paint, `<path d="M12 3v18"/><path d="m16 16 4-4-4-4"/><path d="m8 8-4 4 4 4"/>`)

// ServerCog is [server-cog](https://lucide.dev/icons/server-cog).
var ServerCog = icon.Lucide(viewBox, paint, `<path d="m10.852 14.772-.383.923"/><path d="M13.148 14.772a3 3 0 1 0-2.296-5.544l-.383-.923"/><path d="m13.148 9.228.383-.923"/><path d="m13.53 15.696-.382-.924a3 3 0 1 1-2.296-5.544"/><path d="m14.772 10.852.923-.383"/><path d="m14.772 13.148.923.383"/><path d="M4.5 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-.5"/><path d="M4.5 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-.5"/><path d="M6 18h.01"/><path d="M6 6h.01"/><path d="m9.228 10.852-.923-.383"/><path d="m9.228 13.148-.923.383"/>`)

// ServerCrash is [server-crash](https://lucide.dev/icons/server-crash).
var ServerCrash = icon.Lucide(viewBox, paint, `<path d="M6 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2"/><path d="M6 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2"/><path d="M6 6h.01"/><path d="M6 18h.01"/><path d="m13 6-4 6h6l-4 6"/>`)

// ServerNas is [server-nas](https://lucide.dev/icons/server-nas).
var ServerNas = icon.Lucide(viewBox, paint, `<path d="M10 16V8"/><path d="M14 16V8"/><path d="M18 16V8"/><path d="M6 12h.01"/><path d="M6 16h.01"/><rect x="2" y="4" width="20" height="16" rx="2"/>`)

// ServerOff is [server-off](https://lucide.dev/icons/server-off).
var ServerOff = icon.Lucide(viewBox, paint, `<path d="M7 2h13a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-5"/><path d="M10 10 2.5 2.5C2 2 2 2.5 2 5v3a2 2 0 0 0 2 2h6z"/><path d="M22 17v-1a2 2 0 0 0-2-2h-1"/><path d="M4 14a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16.5l1-.5.5.5-8-8H4z"/><path d="M6 18h.01"/><path d="m2 2 20 20"/>`)

// ServerPlus is [server-plus](https://lucide.dev/icons/server-plus).
var ServerPlus = icon.Lucide(viewBox, paint, `<path d="M12.5 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2"/><path d="M16 12h6"/><path d="M19 9v6"/><path d="M22 18v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h8.5"/><path d="M6 18h.01"/><path d="M6 6h.01"/>`)

// Server is [server](https://lucide.dev/icons/server).
var Server = icon.Lucide(viewBox, paint, `<rect width="20" height="8" x="2" y="2" rx="2" ry="2"/><rect width="20" height="8" x="2" y="14" rx="2" ry="2"/><line x1="6" x2="6.01" y1="6" y2="6"/><line x1="6" x2="6.01" y1="18" y2="18"/>`)

// Settings2 is [settings-2](https://lucide.dev/icons/settings-2).
var Settings2 = icon.Lucide(viewBox, paint, `<path d="M14 17H5"/><path d="M19 7h-9"/><circle cx="17" cy="17" r="3"/><circle cx="7" cy="7" r="3"/>`)

// Settings is [settings](https://lucide.dev/icons/settings).
var Settings = icon.Lucide(viewBox, paint, `<path d="M9.671 4.136a2.34 2.34 0 0 1 4.659 0 2.34 2.34 0 0 0 3.319 1.915 2.34 2.34 0 0 1 2.33 4.033 2.34 2.34 0 0 0 0 3.831 2.34 2.34 0 0 1-2.33 4.033 2.34 2.34 0 0 0-3.319 1.915 2.34 2.34 0 0 1-4.659 0 2.34 2.34 0 0 0-3.32-1.915 2.34 2.34 0 0 1-2.33-4.033 2.34 2.34 0 0 0 0-3.831A2.34 2.34 0 0 1 6.35 6.051a2.34 2.34 0 0 0 3.319-1.915"/><circle cx="12" cy="12" r="3"/>`)

// Shapes is [shapes](https://lucide.dev/icons/shapes).
var Shapes = icon.Lucide(viewBox, paint, `<path d="M8.3 10a.7.7 0 0 1-.626-1.079L11.4 3a.7.7 0 0 1 1.198-.043L16.3 8.9a.7.7 0 0 1-.572 1.1Z"/><rect x="3" y="14" width="7" height="7" rx="1"/><circle cx="17.5" cy="17.5" r="3.5"/>`)

// Share2 is [share-2](https://lucide.dev/icons/share-2).
var Share2 = icon.Lucide(viewBox, paint, `<circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/><line x1="8.59" x2="15.42" y1="13.51" y2="17.49"/><line x1="15.41" x2="8.59" y1="6.51" y2="10.49"/>`)

// Share is [share](https://lucide.dev/icons/share).
var Share = icon.Lucide(viewBox, paint, `<path d="M12 2v13"/><path d="m16 6-4-4-4 4"/><path d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/>`)

// Sheet is [sheet](https://lucide.dev/icons/sheet).
var Sheet = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><line x1="3" x2="21" y1="9" y2="9"/><line x1="3" x2="21" y1="15" y2="15"/><line x1="9" x2="9" y1="9" y2="21"/><line x1="15" x2="15" y1="9" y2="21"/>`)

// Shell is [shell](https://lucide.dev/icons/shell).
var Shell = icon.Lucide(viewBox, paint, `<path d="M14 11a2 2 0 1 1-4 0 4 4 0 0 1 8 0 6 6 0 0 1-12 0 8 8 0 0 1 16 0 10 10 0 1 1-20 0 11.93 11.93 0 0 1 2.42-7.22 2 2 0 1 1 3.16 2.44"/>`)

// ShelvingUnit is [shelving-unit](https://lucide.dev/icons/shelving-unit).
var ShelvingUnit = icon.Lucide(viewBox, paint, `<path d="M12 12V9a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3"/><path d="M16 20v-3a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3"/><path d="M20 22V2"/><path d="M4 12h16"/><path d="M4 20h16"/><path d="M4 2v20"/><path d="M4 4h16"/>`)

// ShieldAlert is [shield-alert](https://lucide.dev/icons/shield-alert).
var ShieldAlert = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="M12 8v4"/><path d="M12 16h.01"/>`)

// ShieldBan is [shield-ban](https://lucide.dev/icons/shield-ban).
var ShieldBan = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="m4.243 5.21 14.39 12.472"/>`)

// ShieldCheck is [shield-check](https://lucide.dev/icons/shield-check).
var ShieldCheck = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="m9 12 2 2 4-4"/>`)

// ShieldCogCorner is [shield-cog-corner](https://lucide.dev/icons/shield-cog-corner).
var ShieldCogCorner = icon.Lucide(viewBox, paint, `<path d="M11 22c-3.806-1.45-7-3.966-7-9V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1v4"/><path d="M14.923 16.547 14 16.164"/><path d="m14.923 18.843-.923.383"/><path d="M16.547 14.923 16.164 14"/><path d="m16.547 20.467-.383.924"/><path d="m18.843 14.923.383-.923"/><path d="m19.225 21.391-.382-.924"/><path d="m20.467 16.547.923-.383"/><path d="m20.467 18.843.923.383"/><circle cx="17.695" cy="17.695" r="3"/>`)

// ShieldCog is [shield-cog](https://lucide.dev/icons/shield-cog).
var ShieldCog = icon.Lucide(viewBox, paint, `<path d="m10.929 14.467-.383.924"/><path d="M10.929 8.923 10.546 8"/><path d="M13.225 8.923 13.608 8"/><path d="m13.607 15.391-.382-.924"/><path d="m14.849 10.547.923-.383"/><path d="m14.849 12.843.923.383"/><path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="m9.305 10.547-.923-.383"/><path d="m9.305 12.843-.923.383"/><circle cx="12.077" cy="11.695" r="3"/>`)

// ShieldEllipsis is [shield-ellipsis](https://lucide.dev/icons/shield-ellipsis).
var ShieldEllipsis = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="M8 12h.01"/><path d="M12 12h.01"/><path d="M16 12h.01"/>`)

// ShieldHalf is [shield-half](https://lucide.dev/icons/shield-half).
var ShieldHalf = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="M12 22V2"/>`)

// ShieldHouse is [shield-house](https://lucide.dev/icons/shield-house).
var ShieldHouse = icon.Lucide(viewBox, paint, `<path d="M15 22a1 1 0 01-1-1v-4a1 1 0 01.445-.832l3-2a1 1 0 011.11 0l3 2A1 1 0 0122 17v4a1 1 0 01-1 1z"/><path d="M18 22v-3"/><path d="M20 10.418V6a1 1 0 00-1-1c-2 0-4.49-1.19-6.24-2.72a1.17 1.17 0 00-1.52 0C9.5 3.8 7 5 5 5a1 1 0 00-1 1v7c0 4.602 2.67 7.1 6.037 8.603"/>`)

// ShieldKeyhole is [shield-keyhole](https://lucide.dev/icons/shield-keyhole).
var ShieldKeyhole = icon.Lucide(viewBox, paint, `<path d="M12 13v3"/><path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 01-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 011-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 011.52 0C14.51 3.81 17 5 19 5a1 1 0 011 1z"/><circle cx="12" cy="11" r="2"/>`)

// ShieldLock is [shield-lock](https://lucide.dev/icons/shield-lock).
var ShieldLock = icon.Lucide(viewBox, paint, `<path d="M20 9.807V6a1 1 0 00-1-1c-2 0-4.49-1.19-6.24-2.72a1.17 1.17 0 00-1.52 0C9.5 3.8 7 5 5 5a1 1 0 00-1 1v7c0 3.88 2.107 6.254 5 7.796"/><path d="M19 17v-2a2 2 0 00-4 0v2"/><rect x="13" y="17" width="8" height="5" rx="1"/>`)

// ShieldMinus is [shield-minus](https://lucide.dev/icons/shield-minus).
var ShieldMinus = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="M9 12h6"/>`)

// ShieldOff is [shield-off](https://lucide.dev/icons/shield-off).
var ShieldOff = icon.Lucide(viewBox, paint, `<path d="m2 2 20 20"/><path d="M5 5a1 1 0 0 0-1 1v7c0 5 3.5 7.5 7.67 8.94a1 1 0 0 0 .67.01c2.35-.82 4.48-1.97 5.9-3.71"/><path d="M9.309 3.652A12.252 12.252 0 0 0 11.24 2.28a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1v7a9.784 9.784 0 0 1-.08 1.264"/>`)

// ShieldPlus is [shield-plus](https://lucide.dev/icons/shield-plus).
var ShieldPlus = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="M9 12h6"/><path d="M12 9v6"/>`)

// ShieldQuestionMark is [shield-question-mark](https://lucide.dev/icons/shield-question-mark).
var ShieldQuestionMark = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="M9.1 9a3 3 0 0 1 5.82 1c0 2-3 3-3 3"/><path d="M12 17h.01"/>`)

// ShieldUser is [shield-user](https://lucide.dev/icons/shield-user).
var ShieldUser = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="M6.376 18.91a6 6 0 0 1 11.249.003"/><circle cx="12" cy="11" r="4"/>`)

// ShieldX is [shield-x](https://lucide.dev/icons/shield-x).
var ShieldX = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/><path d="m14.5 9.5-5 5"/><path d="m9.5 9.5 5 5"/>`)

// Shield is [shield](https://lucide.dev/icons/shield).
var Shield = icon.Lucide(viewBox, paint, `<path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"/>`)

// ShipCargo is [ship-cargo](https://lucide.dev/icons/ship-cargo).
var ShipCargo = icon.Lucide(viewBox, paint, `<path d="M12 15v-3"/><path d="M12 2v2"/><path d="M16.5 12V9a1 1 0 011-1h1a1 1 0 001-1V5a1 1 0 00-1-1h-13a1 1 0 00-1 1v2a1 1 0 001 1h1a1 1 0 011 1v3"/><path d="M19.38 19c1.076-1.815 1.636-4.89 1.628-6.008a1 1 0 00-1-.992H3.984a1 1 0 00-1 .984c-.03 1.86.97 5.621 2.826 7.776"/><path d="M2 20c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1s1.2 1 2.5 1c2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"/>`)

// ShipWheel is [ship-wheel](https://lucide.dev/icons/ship-wheel).
var ShipWheel = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="8"/><path d="M12 2v7.5"/><path d="m19 5-5.23 5.23"/><path d="M22 12h-7.5"/><path d="m19 19-5.23-5.23"/><path d="M12 14.5V22"/><path d="M10.23 13.77 5 19"/><path d="M9.5 12H2"/><path d="M10.23 10.23 5 5"/><circle cx="12" cy="12" r="2.5"/>`)

// Ship is [ship](https://lucide.dev/icons/ship).
var Ship = icon.Lucide(viewBox, paint, `<path d="M12 2v2"/><path d="M12 9.189V13"/><path d="M19 12V6a2 2 0 00-2-2H7a2 2 0 00-2 2v6"/><path d="M19.38 19A11.6 11.6 0 0021 13l-8.188-3.639a2 2 0 00-1.624 0L3 13.001a11.6 11.6 0 002.81 7.76"/><path d="M2 20c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1s1.2 1 2.5 1c2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"/>`)

// Shirt is [shirt](https://lucide.dev/icons/shirt).
var Shirt = icon.Lucide(viewBox, paint, `<path d="M20.38 3.46 16 2a4 4 0 0 1-8 0L3.62 3.46a2 2 0 0 0-1.34 2.23l.58 3.47a1 1 0 0 0 .99.84H6v10c0 1.1.9 2 2 2h8a2 2 0 0 0 2-2V10h2.15a1 1 0 0 0 .99-.84l.58-3.47a2 2 0 0 0-1.34-2.23z"/>`)

// ShoppingBag is [shopping-bag](https://lucide.dev/icons/shopping-bag).
var ShoppingBag = icon.Lucide(viewBox, paint, `<path d="M16 10a4 4 0 0 1-8 0"/><path d="M3.103 6.034h17.794"/><path d="M3.4 5.467a2 2 0 0 0-.4 1.2V20a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6.667a2 2 0 0 0-.4-1.2l-2-2.667A2 2 0 0 0 17 2H7a2 2 0 0 0-1.6.8z"/>`)

// ShoppingBasket is [shopping-basket](https://lucide.dev/icons/shopping-basket).
var ShoppingBasket = icon.Lucide(viewBox, paint, `<path d="m15 11-1 9"/><path d="m19 11-4-7"/><path d="M2 11h20"/><path d="m3.5 11 1.6 7.4a2 2 0 0 0 2 1.6h9.8a2 2 0 0 0 2-1.6l1.7-7.4"/><path d="M4.5 15.5h15"/><path d="m5 11 4-7"/><path d="m9 11 1 9"/>`)

// ShoppingCartMinus is [shopping-cart-minus](https://lucide.dev/icons/shopping-cart-minus).
var ShoppingCartMinus = icon.Lucide(viewBox, paint, `<path d="M16 5h6"/><path d="m2.05 2.05 1.099-.028a1 1 0 011.008.815l2.69 14.347A1 1 0 007.83 18H18"/><path d="M4.564 5H12"/><path d="M6.25 14h12.712a2 2 0 001.991-1.57l.514-3.113"/><circle cx="18" cy="20" r="2"/><circle cx="8" cy="20" r="2"/>`)

// ShoppingCartPlus is [shopping-cart-plus](https://lucide.dev/icons/shopping-cart-plus).
var ShoppingCartPlus = icon.Lucide(viewBox, paint, `<path d="M16 5h6"/><path d="M19 2v6"/><path d="m2.05 2.05 1.099-.028a1 1 0 011.008.815l2.69 14.347A1 1 0 007.83 18H18"/><path d="M4.564 5H12"/><path d="M6.25 14h12.712a2 2 0 001.991-1.57l.172-1.041"/><circle cx="18" cy="20" r="2"/><circle cx="8" cy="20" r="2"/>`)

// ShoppingCart is [shopping-cart](https://lucide.dev/icons/shopping-cart).
var ShoppingCart = icon.Lucide(viewBox, paint, `<path d="m2.05 2.05 1.099-.028a1 1 0 0 1 1.008.815l2.69 14.347A1 1 0 0 0 7.83 18H18"/><path d="M4.563 5h16.435a1 1 0 0 1 .981 1.204l-1.026 6.226A2 2 0 0 1 18.962 14H6.25"/><circle cx="18" cy="20" r="2"/><circle cx="8" cy="20" r="2"/>`)

// Shovel is [shovel](https://lucide.dev/icons/shovel).
var Shovel = icon.Lucide(viewBox, paint, `<path d="M21.56 4.56a1.5 1.5 0 0 1 0 2.122l-.47.47a3 3 0 0 1-4.212-.03 3 3 0 0 1 0-4.243l.44-.44a1.5 1.5 0 0 1 2.121 0z"/><path d="M3 22a1 1 0 0 1-1-1v-3.586a1 1 0 0 1 .293-.707l3.355-3.355a1.205 1.205 0 0 1 1.704 0l3.296 3.296a1.205 1.205 0 0 1 0 1.704l-3.355 3.355a1 1 0 0 1-.707.293z"/><path d="m9 15 7.879-7.878"/>`)

// ShowerHead is [shower-head](https://lucide.dev/icons/shower-head).
var ShowerHead = icon.Lucide(viewBox, paint, `<path d="m4 4 2.5 2.5"/><path d="M13.5 6.5a4.95 4.95 0 0 0-7 7"/><path d="M15 5 5 15"/><path d="M14 17v.01"/><path d="M10 16v.01"/><path d="M13 13v.01"/><path d="M16 10v.01"/><path d="M11 20v.01"/><path d="M17 14v.01"/><path d="M20 11v.01"/>`)

// Shredder is [shredder](https://lucide.dev/icons/shredder).
var Shredder = icon.Lucide(viewBox, paint, `<path d="M4 13V4a2 2 0 0 1 2-2h8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 20 8v5"/><path d="M14 2v5a1 1 0 0 0 1 1h5"/><path d="M10 22v-5"/><path d="M14 19v-2"/><path d="M18 20v-3"/><path d="M2 13h20"/><path d="M6 20v-3"/>`)

// ShrimpOff is [shrimp-off](https://lucide.dev/icons/shrimp-off).
var ShrimpOff = icon.Lucide(viewBox, paint, `<path d="M10 2a3.28 3.28 0 003.227 1.798l6.17-.561A1 1 0 1119.614 8H13.5"/><path d="M11 20c-.5.5-1.12 1-2.5 1a1 1 0 010-5H12a7 7 0 003.283-.817"/><path d="M11 22c-.5-.5-1.12-1-2.5-1a6.5 6.5 0 01-5.63-3.25 6.44 6.44 0 015.236-9.744"/><path d="M18.04 12.54A7 7 0 0019 9V8"/><path d="m2 2 20 20"/><path d="M8 16c-2 0-4.5-4-4-6"/><path d="M9.43 9.33A8.5 8.5 0 0010 16"/>`)

// Shrimp is [shrimp](https://lucide.dev/icons/shrimp).
var Shrimp = icon.Lucide(viewBox, paint, `<path d="M10 2a3.28 3.28 0 003.227 1.798l6.17-.561A1 1 0 1119.614 8H8.5a6.44 6.44 0 00-5.63 9.75A6.5 6.5 0 008.5 21c1.38 0 2-.5 2.5-1"/><path d="M10 8a8.5 8.5 0 000 8"/><path d="M11 22c-.5-.5-1.12-1-2.5-1a1 1 0 010-5H12a7 7 0 007-7V8"/><path d="M13 12h.01"/><path d="M8 16c-2 0-4.5-4-4-6"/>`)

// Shrink is [shrink](https://lucide.dev/icons/shrink).
var Shrink = icon.Lucide(viewBox, paint, `<path d="m15 15 6 6m-6-6v4.8m0-4.8h4.8"/><path d="M9 19.8V15m0 0H4.2M9 15l-6 6"/><path d="M15 4.2V9m0 0h4.8M15 9l6-6"/><path d="M9 4.2V9m0 0H4.2M9 9 3 3"/>`)

// Shrub is [shrub](https://lucide.dev/icons/shrub).
var Shrub = icon.Lucide(viewBox, paint, `<path d="M12 22v-5.172a2 2 0 0 0-.586-1.414L9.5 13.5"/><path d="M14.5 14.5 12 17"/><path d="M17 8.8A6 6 0 0 1 13.8 20H10A6.5 6.5 0 0 1 7 8a5 5 0 0 1 10 0z"/>`)

// Shuffle is [shuffle](https://lucide.dev/icons/shuffle).
var Shuffle = icon.Lucide(viewBox, paint, `<path d="m18 14 4 4-4 4"/><path d="m18 2 4 4-4 4"/><path d="M2 18h1.973a4 4 0 0 0 3.3-1.7l5.454-8.6a4 4 0 0 1 3.3-1.7H22"/><path d="M2 6h1.972a4 4 0 0 1 3.6 2.2"/><path d="M22 18h-6.041a4 4 0 0 1-3.3-1.8l-.359-.45"/>`)

// Sigma is [sigma](https://lucide.dev/icons/sigma).
var Sigma = icon.Lucide(viewBox, paint, `<path d="M18 7V5a1 1 0 0 0-1-1H6.5a.5.5 0 0 0-.4.8l4.5 6a2 2 0 0 1 0 2.4l-4.5 6a.5.5 0 0 0 .4.8H17a1 1 0 0 0 1-1v-2"/>`)

// SignalHigh is [signal-high](https://lucide.dev/icons/signal-high).
var SignalHigh = icon.Lucide(viewBox, paint, `<path d="M2 20h.01"/><path d="M7 20v-4"/><path d="M12 20v-8"/><path d="M17 20V8"/>`)

// SignalLow is [signal-low](https://lucide.dev/icons/signal-low).
var SignalLow = icon.Lucide(viewBox, paint, `<path d="M2 20h.01"/><path d="M7 20v-4"/>`)

// SignalMedium is [signal-medium](https://lucide.dev/icons/signal-medium).
var SignalMedium = icon.Lucide(viewBox, paint, `<path d="M2 20h.01"/><path d="M7 20v-4"/><path d="M12 20v-8"/>`)

// SignalZero is [signal-zero](https://lucide.dev/icons/signal-zero).
var SignalZero = icon.Lucide(viewBox, paint, `<path d="M2 20h.01"/>`)

// Signal is [signal](https://lucide.dev/icons/signal).
var Signal = icon.Lucide(viewBox, paint, `<path d="M2 20h.01"/><path d="M7 20v-4"/><path d="M12 20v-8"/><path d="M17 20V8"/><path d="M22 4v16"/>`)

// Signature is [signature](https://lucide.dev/icons/signature).
var Signature = icon.Lucide(viewBox, paint, `<path d="m21 17-2.156-1.868A.5.5 0 0 0 18 15.5v.5a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1c0-2.545-3.991-3.97-8.5-4a1 1 0 0 0 0 5c4.153 0 4.745-11.295 5.708-13.5a2.5 2.5 0 1 1 3.31 3.284"/><path d="M3 21h18"/>`)

// SignpostBig is [signpost-big](https://lucide.dev/icons/signpost-big).
var SignpostBig = icon.Lucide(viewBox, paint, `<path d="M10 9H4L2 7l2-2h6"/><path d="M14 5h6l2 2-2 2h-6"/><path d="M10 22V4a2 2 0 1 1 4 0v18"/><path d="M8 22h8"/>`)

// Signpost is [signpost](https://lucide.dev/icons/signpost).
var Signpost = icon.Lucide(viewBox, paint, `<path d="M12 13v8"/><path d="M12 3v3"/><path d="M2.354 10.354a1.207 1.207 0 0 1 0-1.708l2.06-2.06A2 2 0 0 1 5.828 6h12.344a2 2 0 0 1 1.414.586l2.06 2.06a1.207 1.207 0 0 1 0 1.708l-2.06 2.06a2 2 0 0 1-1.414.586H5.828a2 2 0 0 1-1.414-.586z"/>`)

// Siren is [siren](https://lucide.dev/icons/siren).
var Siren = icon.Lucide(viewBox, paint, `<path d="M7 18v-6a5 5 0 1 1 10 0v6"/><path d="M5 21a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2z"/><path d="M21 12h1"/><path d="M18.5 4.5 18 5"/><path d="M2 12h1"/><path d="M12 2v1"/><path d="m4.929 4.929.707.707"/><path d="M12 12v6"/>`)

// SixteenthNotes is [sixteenth-notes](https://lucide.dev/icons/sixteenth-notes).
var SixteenthNotes = icon.Lucide(viewBox, paint, `<path d="M21 16V3a1 1 0 0 0-1.28-.96l-10 2.917a1 1 0 0 0-.72.96V19"/><path d="M21 6 9 9.5"/><circle cx="18" cy="16" r="3"/><circle cx="6" cy="19" r="3"/>`)

// SkipBack is [skip-back](https://lucide.dev/icons/skip-back).
var SkipBack = icon.Lucide(viewBox, paint, `<path d="M17.971 4.285A2 2 0 0 1 21 6v12a2 2 0 0 1-3.029 1.715l-9.997-5.998a2 2 0 0 1-.003-3.432z"/><path d="M3 20V4"/>`)

// SkipForward is [skip-forward](https://lucide.dev/icons/skip-forward).
var SkipForward = icon.Lucide(viewBox, paint, `<path d="M21 4v16"/><path d="M6.029 4.285A2 2 0 0 0 3 6v12a2 2 0 0 0 3.029 1.715l9.997-5.998a2 2 0 0 0 .003-3.432z"/>`)

// Skull is [skull](https://lucide.dev/icons/skull).
var Skull = icon.Lucide(viewBox, paint, `<path d="m12.5 17-.5-1-.5 1h1z"/><path d="M15 22a1 1 0 0 0 1-1v-1a2 2 0 0 0 1.56-3.25 8 8 0 1 0-11.12 0A2 2 0 0 0 8 20v1a1 1 0 0 0 1 1z"/><circle cx="15" cy="12" r="1"/><circle cx="9" cy="12" r="1"/>`)

// Slash is [slash](https://lucide.dev/icons/slash).
var Slash = icon.Lucide(viewBox, paint, `<path d="M22 2 2 22"/>`)

// Slice is [slice](https://lucide.dev/icons/slice).
var Slice = icon.Lucide(viewBox, paint, `<path d="M11 16.586V19a1 1 0 0 1-1 1H2L18.37 3.63a1 1 0 1 1 3 3l-9.663 9.663a1 1 0 0 1-1.414 0L8 14"/>`)

// SlidersHorizontal is [sliders-horizontal](https://lucide.dev/icons/sliders-horizontal).
var SlidersHorizontal = icon.Lucide(viewBox, paint, `<path d="M10 5H3"/><path d="M12 19H3"/><path d="M14 3v4"/><path d="M16 17v4"/><path d="M21 12h-9"/><path d="M21 19h-5"/><path d="M21 5h-7"/><path d="M8 10v4"/><path d="M8 12H3"/>`)

// SlidersVertical is [sliders-vertical](https://lucide.dev/icons/sliders-vertical).
var SlidersVertical = icon.Lucide(viewBox, paint, `<path d="M10 8h4"/><path d="M12 21v-9"/><path d="M12 8V3"/><path d="M17 16h4"/><path d="M19 12V3"/><path d="M19 21v-5"/><path d="M3 14h4"/><path d="M5 10V3"/><path d="M5 21v-7"/>`)

// SmartphoneCharging is [smartphone-charging](https://lucide.dev/icons/smartphone-charging).
var SmartphoneCharging = icon.Lucide(viewBox, paint, `<rect width="14" height="20" x="5" y="2" rx="2" ry="2"/><path d="M12.667 8 10 12h4l-2.667 4"/>`)

// SmartphoneNfc is [smartphone-nfc](https://lucide.dev/icons/smartphone-nfc).
var SmartphoneNfc = icon.Lucide(viewBox, paint, `<rect width="7" height="12" x="2" y="6" rx="1"/><path d="M13 8.32a7.43 7.43 0 0 1 0 7.36"/><path d="M16.46 6.21a11.76 11.76 0 0 1 0 11.58"/><path d="M19.91 4.1a15.91 15.91 0 0 1 .01 15.8"/>`)

// SmartphoneUserRound is [smartphone-user-round](https://lucide.dev/icons/smartphone-user-round).
var SmartphoneUserRound = icon.Lucide(viewBox, paint, `<path d="M10 22H6a2 2 0 01-2-2V4a2 2 0 012-2h10a2 2 0 012 2v4"/><path d="M11 18h.01"/><path d="M22 22a4 4 0 00-8 0"/><circle cx="18" cy="15" r="3"/>`)

// Smartphone is [smartphone](https://lucide.dev/icons/smartphone).
var Smartphone = icon.Lucide(viewBox, paint, `<rect width="14" height="20" x="5" y="2" rx="2" ry="2"/><path d="M12 18h.01"/>`)

// Snail is [snail](https://lucide.dev/icons/snail).
var Snail = icon.Lucide(viewBox, paint, `<path d="M2 13a6 6 0 1 0 12 0 4 4 0 1 0-8 0 2 2 0 0 0 4 0"/><circle cx="10" cy="13" r="8"/><path d="M2 21h12c4.4 0 8-3.6 8-8V7a2 2 0 1 0-4 0v6"/><path d="M18 3 19.1 5.2"/><path d="M22 3 20.9 5.2"/>`)

// Snowflake is [snowflake](https://lucide.dev/icons/snowflake).
var Snowflake = icon.Lucide(viewBox, paint, `<path d="m10 20-1.25-2.5L6 18"/><path d="M10 4 8.75 6.5 6 6"/><path d="m14 20 1.25-2.5L18 18"/><path d="m14 4 1.25 2.5L18 6"/><path d="m17 21-3-6h-4"/><path d="m17 3-3 6 1.5 3"/><path d="M2 12h6.5L10 9"/><path d="m20 10-1.5 2 1.5 2"/><path d="M22 12h-6.5L14 15"/><path d="m4 10 1.5 2L4 14"/><path d="m7 21 3-6-1.5-3"/><path d="m7 3 3 6h4"/>`)

// SoapDispenserDroplet is [soap-dispenser-droplet](https://lucide.dev/icons/soap-dispenser-droplet).
var SoapDispenserDroplet = icon.Lucide(viewBox, paint, `<path d="M10.5 2v4"/><path d="M14 2H7a2 2 0 0 0-2 2"/><path d="M19.29 14.76A6.67 6.67 0 0 1 17 11a6.6 6.6 0 0 1-2.29 3.76c-1.15.92-1.71 2.04-1.71 3.19 0 2.22 1.8 4.05 4 4.05s4-1.83 4-4.05c0-1.16-.57-2.26-1.71-3.19"/><path d="M9.607 21H6a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h7V7a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3"/>`)

// Sofa is [sofa](https://lucide.dev/icons/sofa).
var Sofa = icon.Lucide(viewBox, paint, `<path d="M20 9V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v3"/><path d="M2 16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v1.5a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5V11a2 2 0 0 0-4 0z"/><path d="M4 18v2"/><path d="M20 18v2"/><path d="M12 4v9"/>`)

// SolarPanel is [solar-panel](https://lucide.dev/icons/solar-panel).
var SolarPanel = icon.Lucide(viewBox, paint, `<path d="M11 2h2"/><path d="m14.28 14-4.56 8"/><path d="m21 22-1.558-4H4.558"/><path d="M3 10v2"/><path d="M6.245 15.04A2 2 0 0 1 8 14h12a1 1 0 0 1 .864 1.505l-3.11 5.457A2 2 0 0 1 16 22H4a1 1 0 0 1-.863-1.506z"/><path d="M7 2a4 4 0 0 1-4 4"/><path d="m8.66 7.66 1.41 1.41"/>`)

// Soup is [soup](https://lucide.dev/icons/soup).
var Soup = icon.Lucide(viewBox, paint, `<path d="M11.248 3c.273.1.808.53.747 1.36-.05.83-.938 1.2-.989 2.02-.06.78.333 1.24.727 1.62"/><path d="M16.252 3c.268.1.794.53.745 1.36-.06.83-.923 1.2-.993 2.02-.05.78.338 1.24.725 1.62"/><path d="M19.5 12 22 6"/><path d="M4 12a1 1 0 00-.99 1.133A9 9 0 0012 21a9 9 0 008.99-7.867A1 1 0 0020 12z"/><path d="M6.252 3c.268.1.794.53.745 1.36-.06.83-.923 1.2-.993 2.02-.05.78.338 1.24.735 1.62"/><path d="M7 21h10"/>`)

// Space is [space](https://lucide.dev/icons/space).
var Space = icon.Lucide(viewBox, paint, `<path d="M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"/>`)

// Spade is [spade](https://lucide.dev/icons/spade).
var Spade = icon.Lucide(viewBox, paint, `<path d="M12 18v4"/><path d="M2 14.499a5.5 5.5 0 0 0 9.591 3.675.6.6 0 0 1 .818.001A5.5 5.5 0 0 0 22 14.5c0-2.29-1.5-4-3-5.5l-5.492-5.312a2 2 0 0 0-3-.02L5 8.999c-1.5 1.5-3 3.2-3 5.5"/>`)

// Sparkle is [sparkle](https://lucide.dev/icons/sparkle).
var Sparkle = icon.Lucide(viewBox, paint, `<path d="M11.017 2.814a1 1 0 0 1 1.966 0l1.051 5.558a2 2 0 0 0 1.594 1.594l5.558 1.051a1 1 0 0 1 0 1.966l-5.558 1.051a2 2 0 0 0-1.594 1.594l-1.051 5.558a1 1 0 0 1-1.966 0l-1.051-5.558a2 2 0 0 0-1.594-1.594l-5.558-1.051a1 1 0 0 1 0-1.966l5.558-1.051a2 2 0 0 0 1.594-1.594z"/>`)

// Sparkles is [sparkles](https://lucide.dev/icons/sparkles).
var Sparkles = icon.Lucide(viewBox, paint, `<path d="M11.017 2.814a1 1 0 0 1 1.966 0l1.051 5.558a2 2 0 0 0 1.594 1.594l5.558 1.051a1 1 0 0 1 0 1.966l-5.558 1.051a2 2 0 0 0-1.594 1.594l-1.051 5.558a1 1 0 0 1-1.966 0l-1.051-5.558a2 2 0 0 0-1.594-1.594l-5.558-1.051a1 1 0 0 1 0-1.966l5.558-1.051a2 2 0 0 0 1.594-1.594z"/><path d="M20 2v4"/><path d="M22 4h-4"/><circle cx="4" cy="20" r="2"/>`)

// Speaker is [speaker](https://lucide.dev/icons/speaker).
var Speaker = icon.Lucide(viewBox, paint, `<rect width="16" height="20" x="4" y="2" rx="2"/><path d="M12 6h.01"/><circle cx="12" cy="14" r="4"/><path d="M12 14h.01"/>`)

// Speech is [speech](https://lucide.dev/icons/speech).
var Speech = icon.Lucide(viewBox, paint, `<path d="M8.8 20v-4.1l1.9.2a2.3 2.3 0 0 0 2.164-2.1V8.3A5.37 5.37 0 0 0 2 8.25c0 2.8.656 3.054 1 4.55a5.77 5.77 0 0 1 .029 2.758L2 20"/><path d="M19.8 17.8a7.5 7.5 0 0 0 .003-10.603"/><path d="M17 15a3.5 3.5 0 0 0-.025-4.975"/>`)

// SpellCheck2 is [spell-check-2](https://lucide.dev/icons/spell-check-2).
var SpellCheck2 = icon.Lucide(viewBox, paint, `<path d="m6 16 6-12 6 12"/><path d="M8 12h8"/><path d="M4 21c1.1 0 1.1-1 2.3-1s1.1 1 2.3 1c1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1 1.1 0 1.1 1 2.3 1 1.1 0 1.1-1 2.3-1"/>`)

// SpellCheck is [spell-check](https://lucide.dev/icons/spell-check).
var SpellCheck = icon.Lucide(viewBox, paint, `<path d="m20 15-5.5 5.5L12 18"/><path d="m4 16 6-12 5.115 10.23"/><path d="M6 12h8"/>`)

// SplinePointer is [spline-pointer](https://lucide.dev/icons/spline-pointer).
var SplinePointer = icon.Lucide(viewBox, paint, `<path d="M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z"/><path d="M5 17A12 12 0 0 1 17 5"/><circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/>`)

// Spline is [spline](https://lucide.dev/icons/spline).
var Spline = icon.Lucide(viewBox, paint, `<circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/><path d="M5 17A12 12 0 0 1 17 5"/>`)

// Split is [split](https://lucide.dev/icons/split).
var Split = icon.Lucide(viewBox, paint, `<path d="M16 3h5v5"/><path d="M8 3H3v5"/><path d="M12 22v-8.3a4 4 0 0 0-1.172-2.872L3 3"/><path d="m15 9 6-6"/>`)

// Spool is [spool](https://lucide.dev/icons/spool).
var Spool = icon.Lucide(viewBox, paint, `<path d="M17 13.44 4.442 17.082A2 2 0 0 0 4.982 21H19a2 2 0 0 0 .558-3.921l-1.115-.32A2 2 0 0 1 17 14.837V7.66"/><path d="m7 10.56 12.558-3.642A2 2 0 0 0 19.018 3H5a2 2 0 0 0-.558 3.921l1.115.32A2 2 0 0 1 7 9.163v7.178"/>`)

// SportShoe is [sport-shoe](https://lucide.dev/icons/sport-shoe).
var SportShoe = icon.Lucide(viewBox, paint, `<path d="m15 10.42 4.8-5.07"/><path d="M19 18h3"/><path d="M9.5 22 21.414 9.415A2 2 0 0 0 21.2 6.4l-5.61-4.208A1 1 0 0 0 14 3v2a2 2 0 0 1-1.394 1.906L8.677 8.053A1 1 0 0 0 8 9c-.155 6.393-2.082 9-4 9a2 2 0 0 0 0 4h14"/>`)

// Spotlight is [spotlight](https://lucide.dev/icons/spotlight).
var Spotlight = icon.Lucide(viewBox, paint, `<path d="M15.295 19.562 16 22"/><path d="m17 16 3.758 2.098"/><path d="m19 12.5 3.026-.598"/><path d="M7.61 6.3a3 3 0 0 0-3.92 1.3l-1.38 2.79a3 3 0 0 0 1.3 3.91l6.89 3.597a1 1 0 0 0 1.342-.447l3.106-6.211a1 1 0 0 0-.447-1.341z"/><path d="M8 9V2"/>`)

// SprayCan is [spray-can](https://lucide.dev/icons/spray-can).
var SprayCan = icon.Lucide(viewBox, paint, `<path d="M3 3h.01"/><path d="M7 5h.01"/><path d="M11 7h.01"/><path d="M3 7h.01"/><path d="M7 9h.01"/><path d="M3 11h.01"/><rect width="4" height="4" x="15" y="5"/><path d="m19 9 2 2v10c0 .6-.4 1-1 1h-6c-.6 0-1-.4-1-1V11l2-2"/><path d="m13 14 8-2"/><path d="m13 19 8-2"/>`)

// Sprout is [sprout](https://lucide.dev/icons/sprout).
var Sprout = icon.Lucide(viewBox, paint, `<path d="M14 9.536V7a4 4 0 0 1 4-4h1.5a.5.5 0 0 1 .5.5V5a4 4 0 0 1-4 4 4 4 0 0 0-4 4c0 2 1 3 1 5a5 5 0 0 1-1 3"/><path d="M4 9a5 5 0 0 1 8 4 5 5 0 0 1-8-4"/><path d="M5 21h14"/>`)

// SquareActivity is [square-activity](https://lucide.dev/icons/square-activity).
var SquareActivity = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M17 12h-2l-2 5-2-10-2 5H7"/>`)

// SquareArrowDownLeft is [square-arrow-down-left](https://lucide.dev/icons/square-arrow-down-left).
var SquareArrowDownLeft = icon.Lucide(viewBox, paint, `<path d="M15 15H9l6-6"/><path d="M9 15V9"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareArrowDownRight is [square-arrow-down-right](https://lucide.dev/icons/square-arrow-down-right).
var SquareArrowDownRight = icon.Lucide(viewBox, paint, `<path d="M15 15 9 9"/><path d="M9 15h6V9"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareArrowDown is [square-arrow-down](https://lucide.dev/icons/square-arrow-down).
var SquareArrowDown = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 8v8"/><path d="m8 12 4 4 4-4"/>`)

// SquareArrowLeft is [square-arrow-left](https://lucide.dev/icons/square-arrow-left).
var SquareArrowLeft = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m12 8-4 4 4 4"/><path d="M16 12H8"/>`)

// SquareArrowOutDownLeft is [square-arrow-out-down-left](https://lucide.dev/icons/square-arrow-out-down-left).
var SquareArrowOutDownLeft = icon.Lucide(viewBox, paint, `<path d="M13 21h6a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v6"/><path d="m3 21 9-9"/><path d="M9 21H3v-6"/>`)

// SquareArrowOutDownRight is [square-arrow-out-down-right](https://lucide.dev/icons/square-arrow-out-down-right).
var SquareArrowOutDownRight = icon.Lucide(viewBox, paint, `<path d="M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"/><path d="m21 21-9-9"/><path d="M21 15v6h-6"/>`)

// SquareArrowOutUpLeft is [square-arrow-out-up-left](https://lucide.dev/icons/square-arrow-out-up-left).
var SquareArrowOutUpLeft = icon.Lucide(viewBox, paint, `<path d="M13 3h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-6"/><path d="m3 3 9 9"/><path d="M3 9V3h6"/>`)

// SquareArrowOutUpRight is [square-arrow-out-up-right](https://lucide.dev/icons/square-arrow-out-up-right).
var SquareArrowOutUpRight = icon.Lucide(viewBox, paint, `<path d="M21 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h6"/><path d="m21 3-9 9"/><path d="M15 3h6v6"/>`)

// SquareArrowRightEnter is [square-arrow-right-enter](https://lucide.dev/icons/square-arrow-right-enter).
var SquareArrowRightEnter = icon.Lucide(viewBox, paint, `<path d="m10 16 4-4-4-4"/><path d="M3 12h11"/><path d="M3 8V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3"/>`)

// SquareArrowRightExit is [square-arrow-right-exit](https://lucide.dev/icons/square-arrow-right-exit).
var SquareArrowRightExit = icon.Lucide(viewBox, paint, `<path d="M10 12h11"/><path d="m17 16 4-4-4-4"/><path d="M21 6.344V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-1.344"/>`)

// SquareArrowRight is [square-arrow-right](https://lucide.dev/icons/square-arrow-right).
var SquareArrowRight = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8"/><path d="m12 16 4-4-4-4"/>`)

// SquareArrowUpLeft is [square-arrow-up-left](https://lucide.dev/icons/square-arrow-up-left).
var SquareArrowUpLeft = icon.Lucide(viewBox, paint, `<path d="M15 15 9 9"/><path d="M9 15V9h6"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareArrowUpRight is [square-arrow-up-right](https://lucide.dev/icons/square-arrow-up-right).
var SquareArrowUpRight = icon.Lucide(viewBox, paint, `<path d="M15 15V9H9"/><path d="m9 15 6-6"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareArrowUp is [square-arrow-up](https://lucide.dev/icons/square-arrow-up).
var SquareArrowUp = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 12-4-4-4 4"/><path d="M12 16V8"/>`)

// SquareAsterisk is [square-asterisk](https://lucide.dev/icons/square-asterisk).
var SquareAsterisk = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 8v8"/><path d="m8.5 14 7-4"/><path d="m8.5 10 7 4"/>`)

// SquareBookmark is [square-bookmark](https://lucide.dev/icons/square-bookmark).
var SquareBookmark = icon.Lucide(viewBox, paint, `<path d="M11 3v7.751a.25.25 0 00.407.195l2.28-1.834a.5.5 0 01.627 0l2.28 1.834a.25.25 0 00.406-.195V3"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareBottomDashedScissors is [square-bottom-dashed-scissors](https://lucide.dev/icons/square-bottom-dashed-scissors).
var SquareBottomDashedScissors = icon.Lucide(viewBox, paint, `<path d="M14 21h1"/><path d="m17 17-2.18-2.18"/><path d="M5 21a2 2 0 01-2-2V5a2 2 0 012-2h14a2 2 0 012 2v14a2 2 0 01-2 2"/><path d="M9 21h1"/><path d="M9.56 14.44 17 7"/><path d="M9.56 9.56 12 12"/><circle cx="8.5" cy="15.5" r="1.5"/><circle cx="8.5" cy="8.5" r="1.5"/>`)

// SquareCenterlineDashedHorizontal is [square-centerline-dashed-horizontal](https://lucide.dev/icons/square-centerline-dashed-horizontal).
var SquareCenterlineDashedHorizontal = icon.Lucide(viewBox, paint, `<path d="M8 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h3"/><path d="M16 3h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-3"/><path d="M12 20v2"/><path d="M12 14v2"/><path d="M12 8v2"/><path d="M12 2v2"/>`)

// SquareCenterlineDashedVertical is [square-centerline-dashed-vertical](https://lucide.dev/icons/square-centerline-dashed-vertical).
var SquareCenterlineDashedVertical = icon.Lucide(viewBox, paint, `<path d="M21 8V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3"/><path d="M21 16v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3"/><path d="M4 12H2"/><path d="M10 12H8"/><path d="M16 12h-2"/><path d="M22 12h-2"/>`)

// SquareChartGantt is [square-chart-gantt](https://lucide.dev/icons/square-chart-gantt).
var SquareChartGantt = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 8h7"/><path d="M8 12h6"/><path d="M11 16h5"/>`)

// SquareCheckBig is [square-check-big](https://lucide.dev/icons/square-check-big).
var SquareCheckBig = icon.Lucide(viewBox, paint, `<path d="M21 10.656V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h12.344"/><path d="m9 11 3 3L22 4"/>`)

// SquareCheck is [square-check](https://lucide.dev/icons/square-check).
var SquareCheck = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 9-5.5 5.5L8 12"/>`)

// SquareChevronDown is [square-chevron-down](https://lucide.dev/icons/square-chevron-down).
var SquareChevronDown = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 10-4 4-4-4"/>`)

// SquareChevronLeft is [square-chevron-left](https://lucide.dev/icons/square-chevron-left).
var SquareChevronLeft = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m14 16-4-4 4-4"/>`)

// SquareChevronRight is [square-chevron-right](https://lucide.dev/icons/square-chevron-right).
var SquareChevronRight = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m10 8 4 4-4 4"/>`)

// SquareChevronUp is [square-chevron-up](https://lucide.dev/icons/square-chevron-up).
var SquareChevronUp = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m8 14 4-4 4 4"/>`)

// SquareCode is [square-code](https://lucide.dev/icons/square-code).
var SquareCode = icon.Lucide(viewBox, paint, `<path d="m10 9-3 3 3 3"/><path d="m14 15 3-3-3-3"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareDashedBottomCode is [square-dashed-bottom-code](https://lucide.dev/icons/square-dashed-bottom-code).
var SquareDashedBottomCode = icon.Lucide(viewBox, paint, `<path d="M10 9.5 8 12l2 2.5"/><path d="M14 21h1"/><path d="m14 9.5 2 2.5-2 2.5"/><path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2"/><path d="M9 21h1"/>`)

// SquareDashedBottom is [square-dashed-bottom](https://lucide.dev/icons/square-dashed-bottom).
var SquareDashedBottom = icon.Lucide(viewBox, paint, `<path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2"/><path d="M9 21h1"/><path d="M14 21h1"/>`)

// SquareDashedKanban is [square-dashed-kanban](https://lucide.dev/icons/square-dashed-kanban).
var SquareDashedKanban = icon.Lucide(viewBox, paint, `<path d="M8 7v7"/><path d="M12 7v4"/><path d="M16 7v9"/><path d="M5 3a2 2 0 0 0-2 2"/><path d="M9 3h1"/><path d="M14 3h1"/><path d="M19 3a2 2 0 0 1 2 2"/><path d="M21 9v1"/><path d="M21 14v1"/><path d="M21 19a2 2 0 0 1-2 2"/><path d="M14 21h1"/><path d="M9 21h1"/><path d="M5 21a2 2 0 0 1-2-2"/><path d="M3 14v1"/><path d="M3 9v1"/>`)

// SquareDashedMousePointer is [square-dashed-mouse-pointer](https://lucide.dev/icons/square-dashed-mouse-pointer).
var SquareDashedMousePointer = icon.Lucide(viewBox, paint, `<path d="M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z"/><path d="M5 3a2 2 0 0 0-2 2"/><path d="M19 3a2 2 0 0 1 2 2"/><path d="M5 21a2 2 0 0 1-2-2"/><path d="M9 3h1"/><path d="M9 21h2"/><path d="M14 3h1"/><path d="M3 9v1"/><path d="M21 9v2"/><path d="M3 14v1"/>`)

// SquareDashedPlus is [square-dashed-plus](https://lucide.dev/icons/square-dashed-plus).
var SquareDashedPlus = icon.Lucide(viewBox, paint, `<path d="M5 3a2 2 0 0 0-2 2"/><path d="M19 3a2 2 0 0 1 2 2"/><path d="M21 19a2 2 0 0 1-2 2"/><path d="M5 21a2 2 0 0 1-2-2"/><path d="M9 3h1"/><path d="M9 21h1"/><path d="M14 3h1"/><path d="M14 21h1"/><path d="M3 9v1"/><path d="M21 9v1"/><path d="M3 14v1"/><path d="M21 14v1"/><path d="M8 12h8"/><path d="M12 8v8"/>`)

// SquareDashedText is [square-dashed-text](https://lucide.dev/icons/square-dashed-text).
var SquareDashedText = icon.Lucide(viewBox, paint, `<path d="M14 21h1"/><path d="M14 3h1"/><path d="M19 3a2 2 0 0 1 2 2"/><path d="M21 14v1"/><path d="M21 19a2 2 0 0 1-2 2"/><path d="M21 9v1"/><path d="M3 14v1"/><path d="M3 9v1"/><path d="M5 21a2 2 0 0 1-2-2"/><path d="M5 3a2 2 0 0 0-2 2"/><path d="M7 12h10"/><path d="M7 16h6"/><path d="M7 8h8"/><path d="M9 21h1"/><path d="M9 3h1"/>`)

// SquareDashedTopSolid is [square-dashed-top-solid](https://lucide.dev/icons/square-dashed-top-solid).
var SquareDashedTopSolid = icon.Lucide(viewBox, paint, `<path d="M14 21h1"/><path d="M21 14v1"/><path d="M21 19a2 2 0 0 1-2 2"/><path d="M21 9v1"/><path d="M3 14v1"/><path d="M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2"/><path d="M3 9v1"/><path d="M5 21a2 2 0 0 1-2-2"/><path d="M9 21h1"/>`)

// SquareDashedXCorner is [square-dashed-x-corner](https://lucide.dev/icons/square-dashed-x-corner).
var SquareDashedXCorner = icon.Lucide(viewBox, paint, `<path d="M14 3h1"/><path d="m16 16 5 5"/><path d="M19 3a2 2 0 012 2"/><path d="m21 16-5 5"/><path d="M21 9v1"/><path d="M3 14v1"/><path d="M3 9v1"/><path d="M5 21a2 2 0 01-2-2"/><path d="M5 3a2 2 0 00-2 2"/><path d="M9 21h1"/><path d="M9 3h1"/>`)

// SquareDashedX is [square-dashed-x](https://lucide.dev/icons/square-dashed-x).
var SquareDashedX = icon.Lucide(viewBox, paint, `<path d="M14 21h1"/><path d="M14 3h1"/><path d="M19 3a2 2 0 012 2"/><path d="M21 14v1"/><path d="M21 19a2 2 0 01-2 2"/><path d="M21 9v1"/><path d="M3 14v1"/><path d="M3 9v1"/><path d="M5 21a2 2 0 01-2-2"/><path d="M5 3a2 2 0 00-2 2"/><path d="m9 15 6-6"/><path d="M9 21h1"/><path d="M9 3h1"/><path d="m9 9 6 6"/>`)

// SquareDashed is [square-dashed](https://lucide.dev/icons/square-dashed).
var SquareDashed = icon.Lucide(viewBox, paint, `<path d="M5 3a2 2 0 0 0-2 2"/><path d="M19 3a2 2 0 0 1 2 2"/><path d="M21 19a2 2 0 0 1-2 2"/><path d="M5 21a2 2 0 0 1-2-2"/><path d="M9 3h1"/><path d="M9 21h1"/><path d="M14 3h1"/><path d="M14 21h1"/><path d="M3 9v1"/><path d="M21 9v1"/><path d="M3 14v1"/><path d="M21 14v1"/>`)

// SquareDimensions is [square-dimensions](https://lucide.dev/icons/square-dimensions).
var SquareDimensions = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 7H7v5"/><path d="M12 17h5v-5"/>`)

// SquareDivide is [square-divide](https://lucide.dev/icons/square-divide).
var SquareDivide = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><line x1="8" x2="16" y1="12" y2="12"/><line x1="12" x2="12" y1="16" y2="16"/><line x1="12" x2="12" y1="8" y2="8"/>`)

// SquareDot is [square-dot](https://lucide.dev/icons/square-dot).
var SquareDot = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="12" cy="12" r="1"/>`)

// SquareEqual is [square-equal](https://lucide.dev/icons/square-equal).
var SquareEqual = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 10h10"/><path d="M7 14h10"/>`)

// SquareExclamationPoint is [square-exclamation-point](https://lucide.dev/icons/square-exclamation-point).
var SquareExclamationPoint = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><line x1="12" x2="12" y1="8" y2="12"/><line x1="12" x2="12.01" y1="16" y2="16"/>`)

// SquareFunction is [square-function](https://lucide.dev/icons/square-function).
var SquareFunction = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 17c2 0 2.8-1 2.8-2.8V10c0-2 1-3.3 3.2-3"/><path d="M9 11.2h5.7"/>`)

// SquareKanban is [square-kanban](https://lucide.dev/icons/square-kanban).
var SquareKanban = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 7v7"/><path d="M12 7v4"/><path d="M16 7v9"/>`)

// SquareLibrary is [square-library](https://lucide.dev/icons/square-library).
var SquareLibrary = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 7v10"/><path d="M11 7v10"/><path d="m15 7 2 10"/>`)

// SquareM is [square-m](https://lucide.dev/icons/square-m).
var SquareM = icon.Lucide(viewBox, paint, `<path d="M8 16V8.5a.5.5 0 0 1 .9-.3l2.7 3.599a.5.5 0 0 0 .8 0l2.7-3.6a.5.5 0 0 1 .9.3V16"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareMenu is [square-menu](https://lucide.dev/icons/square-menu).
var SquareMenu = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 8h10"/><path d="M7 12h10"/><path d="M7 16h10"/>`)

// SquareMinus is [square-minus](https://lucide.dev/icons/square-minus).
var SquareMinus = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8"/>`)

// SquareMousePointer is [square-mouse-pointer](https://lucide.dev/icons/square-mouse-pointer).
var SquareMousePointer = icon.Lucide(viewBox, paint, `<path d="M12.034 12.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.943l-3.444 1.068a1 1 0 0 0-.66.66l-1.067 3.443a.5.5 0 0 1-.943.033z"/><path d="M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"/>`)

// SquareOff is [square-off](https://lucide.dev/icons/square-off).
var SquareOff = icon.Lucide(viewBox, paint, `<path d="M20.4 20.4a2 2 0 01-1.4.6H5a2 2 0 01-2-2V5a2 2 0 01.59-1.41"/><path d="M21 15.3V5a2 2 0 00-2-2H8.7"/><path d="M22 22 2 2"/>`)

// SquareParkingOff is [square-parking-off](https://lucide.dev/icons/square-parking-off).
var SquareParkingOff = icon.Lucide(viewBox, paint, `<path d="M3.6 3.6A2 2 0 0 1 5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-.59 1.41"/><path d="M3 8.7V19a2 2 0 0 0 2 2h10.3"/><path d="m2 2 20 20"/><path d="M13 13a3 3 0 1 0 0-6H9v2"/><path d="M9 17v-2.3"/>`)

// SquareParking is [square-parking](https://lucide.dev/icons/square-parking).
var SquareParking = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 17V7h4a3 3 0 0 1 0 6H9"/>`)

// SquarePause is [square-pause](https://lucide.dev/icons/square-pause).
var SquarePause = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><line x1="10" x2="10" y1="15" y2="9"/><line x1="14" x2="14" y1="15" y2="9"/>`)

// SquarePen is [square-pen](https://lucide.dev/icons/square-pen).
var SquarePen = icon.Lucide(viewBox, paint, `<path d="M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.375 2.625a1 1 0 0 1 3 3l-9.013 9.014a2 2 0 0 1-.853.505l-2.873.84a.5.5 0 0 1-.62-.62l.84-2.873a2 2 0 0 1 .506-.852z"/>`)

// SquarePercent is [square-percent](https://lucide.dev/icons/square-percent).
var SquarePercent = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="m15 9-6 6"/><path d="M9 9h.01"/><path d="M15 15h.01"/>`)

// SquarePi is [square-pi](https://lucide.dev/icons/square-pi).
var SquarePi = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 7h10"/><path d="M10 7v10"/><path d="M16 17a2 2 0 0 1-2-2V7"/>`)

// SquarePilcrow is [square-pilcrow](https://lucide.dev/icons/square-pilcrow).
var SquarePilcrow = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 12H9.5a2.5 2.5 0 0 1 0-5H17"/><path d="M12 7v10"/><path d="M16 7v10"/>`)

// SquarePlay is [square-play](https://lucide.dev/icons/square-play).
var SquarePlay = icon.Lucide(viewBox, paint, `<rect x="3" y="3" width="18" height="18" rx="2"/><path d="M9 9.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997A1 1 0 0 1 9 14.996z"/>`)

// SquarePlus is [square-plus](https://lucide.dev/icons/square-plus).
var SquarePlus = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8"/><path d="M12 8v8"/>`)

// SquarePower is [square-power](https://lucide.dev/icons/square-power).
var SquarePower = icon.Lucide(viewBox, paint, `<path d="M12 7v4"/><path d="M7.998 9.003a5 5 0 1 0 8-.005"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareRadical is [square-radical](https://lucide.dev/icons/square-radical).
var SquareRadical = icon.Lucide(viewBox, paint, `<path d="M7 12h2l2 5 2-10h4"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareRoundCorner is [square-round-corner](https://lucide.dev/icons/square-round-corner).
var SquareRoundCorner = icon.Lucide(viewBox, paint, `<path d="M21 11a8 8 0 0 0-8-8"/><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>`)

// SquareScissors is [square-scissors](https://lucide.dev/icons/square-scissors).
var SquareScissors = icon.Lucide(viewBox, paint, `<path d="m17 17-2.18-2.18"/><path d="M9.56 14.44 17 7"/><path d="M9.56 9.56 12 12"/><circle cx="8.5" cy="15.5" r="1.5"/><circle cx="8.5" cy="8.5" r="1.5"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareSigma is [square-sigma](https://lucide.dev/icons/square-sigma).
var SquareSigma = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M16 8.9V7H8l4 5-4 5h8v-1.9"/>`)

// SquareSlash is [square-slash](https://lucide.dev/icons/square-slash).
var SquareSlash = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><line x1="9" x2="15" y1="15" y2="9"/>`)

// SquareSparkles is [square-sparkles](https://lucide.dev/icons/square-sparkles).
var SquareSparkles = icon.Lucide(viewBox, paint, `<path d="M11 15H7"/><path d="M15.41 2.49a.6.6 0 011.18 0l.63 3.334a1.2 1.2 0 00.956.955l3.334.631a.6.6 0 010 1.18l-3.334.63a1.2 1.2 0 00-.955.956l-.631 3.334a.6.6 0 01-1.18 0l-.63-3.334a1.2 1.2 0 00-.956-.955L10.49 8.59a.6.6 0 010-1.18l3.334-.63a1.2 1.2 0 00.955-.956z"/><path d="M21 13v6a2 2 0 01-2 2H5a2 2 0 01-2-2V5a2 2 0 012-2h6"/><path d="M9 13v4"/>`)

// SquareSplitHorizontal is [square-split-horizontal](https://lucide.dev/icons/square-split-horizontal).
var SquareSplitHorizontal = icon.Lucide(viewBox, paint, `<path d="M12 2v20"/><path d="M16 3h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-3"/><path d="M8 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3"/>`)

// SquareSplitVertical is [square-split-vertical](https://lucide.dev/icons/square-split-vertical).
var SquareSplitVertical = icon.Lucide(viewBox, paint, `<path d="M2 12h20"/><path d="M21 16v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3"/><path d="M3 8V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3"/>`)

// SquareSquare is [square-square](https://lucide.dev/icons/square-square).
var SquareSquare = icon.Lucide(viewBox, paint, `<rect x="3" y="3" width="18" height="18" rx="2"/><rect x="8" y="8" width="8" height="8" rx="1"/>`)

// SquareStack is [square-stack](https://lucide.dev/icons/square-stack).
var SquareStack = icon.Lucide(viewBox, paint, `<path d="M4 10c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2"/><path d="M10 16c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2"/><rect width="8" height="8" x="14" y="14" rx="2"/>`)

// SquareStar is [square-star](https://lucide.dev/icons/square-star).
var SquareStar = icon.Lucide(viewBox, paint, `<path d="M11.035 7.69a1 1 0 0 1 1.909.024l.737 1.452a1 1 0 0 0 .737.535l1.634.256a1 1 0 0 1 .588 1.806l-1.172 1.168a1 1 0 0 0-.282.866l.259 1.613a1 1 0 0 1-1.541 1.134l-1.465-.75a1 1 0 0 0-.912 0l-1.465.75a1 1 0 0 1-1.539-1.133l.258-1.613a1 1 0 0 0-.282-.866l-1.156-1.153a1 1 0 0 1 .572-1.822l1.633-.256a1 1 0 0 0 .737-.535z"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// SquareStop is [square-stop](https://lucide.dev/icons/square-stop).
var SquareStop = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><rect x="9" y="9" width="6" height="6" rx="1"/>`)

// SquareTerminal is [square-terminal](https://lucide.dev/icons/square-terminal).
var SquareTerminal = icon.Lucide(viewBox, paint, `<path d="m7 11 2-2-2-2"/><path d="M11 13h4"/><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/>`)

// SquareText is [square-text](https://lucide.dev/icons/square-text).
var SquareText = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 8h8"/><path d="M7 12h10"/><path d="M7 16h6"/>`)

// SquareUserRound is [square-user-round](https://lucide.dev/icons/square-user-round).
var SquareUserRound = icon.Lucide(viewBox, paint, `<path d="M18 21a6 6 0 0 0-12 0"/><circle cx="12" cy="11" r="4"/><rect width="18" height="18" x="3" y="3" rx="2"/>`)

// SquareUser is [square-user](https://lucide.dev/icons/square-user).
var SquareUser = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="12" cy="10" r="3"/><path d="M7 21v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"/>`)

// SquareX is [square-x](https://lucide.dev/icons/square-x).
var SquareX = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="m15 9-6 6"/><path d="m9 9 6 6"/>`)

// Square is [square](https://lucide.dev/icons/square).
var Square = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/>`)

// SquaresExclude is [squares-exclude](https://lucide.dev/icons/squares-exclude).
var SquaresExclude = icon.Lucide(viewBox, paint, `<path d="M16 12v2a2 2 0 0 1-2 2H9a1 1 0 0 0-1 1v3a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2h0"/><path d="M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v3a1 1 0 0 1-1 1h-5a2 2 0 0 0-2 2v2"/>`)

// SquaresIntersect is [squares-intersect](https://lucide.dev/icons/squares-intersect).
var SquaresIntersect = icon.Lucide(viewBox, paint, `<path d="M10 22a2 2 0 0 1-2-2"/><path d="M14 2a2 2 0 0 1 2 2"/><path d="M16 22h-2"/><path d="M2 10V8"/><path d="M2 4a2 2 0 0 1 2-2"/><path d="M20 8a2 2 0 0 1 2 2"/><path d="M22 14v2"/><path d="M22 20a2 2 0 0 1-2 2"/><path d="M4 16a2 2 0 0 1-2-2"/><path d="M8 10a2 2 0 0 1 2-2h5a1 1 0 0 1 1 1v5a2 2 0 0 1-2 2H9a1 1 0 0 1-1-1z"/><path d="M8 2h2"/>`)

// SquaresSubtract is [squares-subtract](https://lucide.dev/icons/squares-subtract).
var SquaresSubtract = icon.Lucide(viewBox, paint, `<path d="M10 22a2 2 0 0 1-2-2"/><path d="M16 22h-2"/><path d="M16 4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3a1 1 0 0 0 1-1v-5a2 2 0 0 1 2-2h5a1 1 0 0 0 1-1z"/><path d="M20 8a2 2 0 0 1 2 2"/><path d="M22 14v2"/><path d="M22 20a2 2 0 0 1-2 2"/>`)

// SquaresUnite is [squares-unite](https://lucide.dev/icons/squares-unite).
var SquaresUnite = icon.Lucide(viewBox, paint, `<path d="M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v3a1 1 0 0 0 1 1h3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2v-3a1 1 0 0 0-1-1z"/>`)

// SquircleDashed is [squircle-dashed](https://lucide.dev/icons/squircle-dashed).
var SquircleDashed = icon.Lucide(viewBox, paint, `<path d="M13.77 3.043a34 34 0 0 0-3.54 0"/><path d="M13.771 20.956a33 33 0 0 1-3.541.001"/><path d="M20.18 17.74c-.51 1.15-1.29 1.93-2.439 2.44"/><path d="M20.18 6.259c-.51-1.148-1.291-1.929-2.44-2.438"/><path d="M20.957 10.23a33 33 0 0 1 0 3.54"/><path d="M3.043 10.23a34 34 0 0 0 .001 3.541"/><path d="M6.26 20.179c-1.15-.508-1.93-1.29-2.44-2.438"/><path d="M6.26 3.82c-1.149.51-1.93 1.291-2.44 2.44"/>`)

// Squircle is [squircle](https://lucide.dev/icons/squircle).
var Squircle = icon.Lucide(viewBox, paint, `<path d="M12 3c7.2 0 9 1.8 9 9s-1.8 9-9 9-9-1.8-9-9 1.8-9 9-9"/>`)

// Squirrel is [squirrel](https://lucide.dev/icons/squirrel).
var Squirrel = icon.Lucide(viewBox, paint, `<path d="M15.236 22a3 3 0 0 0-2.2-5"/><path d="M16 20a3 3 0 0 1 3-3h1a2 2 0 0 0 2-2v-2a4 4 0 0 0-4-4V4"/><path d="M18 13h.01"/><path d="M18 6a4 4 0 0 0-4 4 7 7 0 0 0-7 7c0-5 4-5 4-10.5a4.5 4.5 0 1 0-9 0 2.5 2.5 0 0 0 5 0C7 10 3 11 3 17c0 2.8 2.2 5 5 5h10"/>`)

// Stamp is [stamp](https://lucide.dev/icons/stamp).
var Stamp = icon.Lucide(viewBox, paint, `<path d="M14 13V8.5C14 7 15 7 15 5a3 3 0 0 0-6 0c0 2 1 2 1 3.5V13"/><path d="M20 15.5a2.5 2.5 0 0 0-2.5-2.5h-11A2.5 2.5 0 0 0 4 15.5V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1z"/><path d="M5 22h14"/>`)

// StarCheck is [star-check](https://lucide.dev/icons/star-check).
var StarCheck = icon.Lucide(viewBox, paint, `<path d="m19.06 12.501 2.78-2.707a.53.53 0 0 0-.294-.905l-5.166-.755a2.1 2.1 0 0 1-1.595-1.16l-2.31-4.68a.53.53 0 0 0-.95.001L9.216 6.974a2.1 2.1 0 0 1-1.597 1.16l-5.165.755a.53.53 0 0 0-.294.906l3.736 3.637a2.1 2.1 0 0 1 .611 1.879l-.88 5.139a.53.53 0 0 0 .769.56l4.617-2.428.027-.014"/><path d="m15 18 2 2 4-4"/>`)

// StarHalf is [star-half](https://lucide.dev/icons/star-half).
var StarHalf = icon.Lucide(viewBox, paint, `<path d="M12 18.338a2.1 2.1 0 0 0-.987.244L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.12 2.12 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.12 2.12 0 0 0 1.597-1.16l2.309-4.679A.53.53 0 0 1 12 2"/>`)

// StarMinus is [star-minus](https://lucide.dev/icons/star-minus).
var StarMinus = icon.Lucide(viewBox, paint, `<path d="M15 18h6"/><path d="M17.688 14a2.1 2.1 0 0 1 .416-.568l3.736-3.638a.53.53 0 0 0-.294-.905l-5.166-.755a2.1 2.1 0 0 1-1.595-1.16l-2.31-4.68a.53.53 0 0 0-.95.001L9.216 6.974a2.1 2.1 0 0 1-1.597 1.16l-5.165.755a.53.53 0 0 0-.294.906l3.736 3.637a2.1 2.1 0 0 1 .611 1.879l-.88 5.139a.53.53 0 0 0 .769.56l4.617-2.428.027-.014"/>`)

// StarOff is [star-off](https://lucide.dev/icons/star-off).
var StarOff = icon.Lucide(viewBox, paint, `<path d="m10.344 4.688 1.181-2.393a.53.53 0 0 1 .95 0l2.31 4.679a2.12 2.12 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904l-3.237 3.152"/><path d="m17.945 17.945.43 2.505a.53.53 0 0 1-.771.56l-4.618-2.428a2.12 2.12 0 0 0-1.973 0L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.12 2.12 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a8 8 0 0 0 .4-.099"/><path d="m2 2 20 20"/>`)

// StarPlus is [star-plus](https://lucide.dev/icons/star-plus).
var StarPlus = icon.Lucide(viewBox, paint, `<path d="M11.013 18.582 6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.12 2.12 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.12 2.12 0 0 0 1.597-1.16l2.309-4.679a.53.53 0 0 1 .95 0l2.31 4.679a2.12 2.12 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904L20 11.5"/><path d="M15 18h6"/><path d="M18 15v6"/>`)

// StarX is [star-x](https://lucide.dev/icons/star-x).
var StarX = icon.Lucide(viewBox, paint, `<path d="m15.5 15.5 5 5"/><path d="m20.063 11.525 1.777-1.731a.53.53 0 0 0-.294-.905l-5.166-.755a2.1 2.1 0 0 1-1.595-1.16l-2.31-4.68a.53.53 0 0 0-.95.001L9.216 6.974a2.1 2.1 0 0 1-1.597 1.16l-5.165.755a.53.53 0 0 0-.294.906l3.736 3.637a2.1 2.1 0 0 1 .611 1.879l-.88 5.139a.53.53 0 0 0 .769.56l4.617-2.428a2.1 2.1 0 0 1 .987-.243 2 2 0 0 1 .132.004"/><path d="m20.5 15.5-5 5"/>`)

// Star is [star](https://lucide.dev/icons/star).
var Star = icon.Lucide(viewBox, paint, `<path d="M11.525 2.295a.53.53 0 0 1 .95 0l2.31 4.679a2.123 2.123 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904l-3.736 3.638a2.123 2.123 0 0 0-.611 1.878l.882 5.14a.53.53 0 0 1-.771.56l-4.618-2.428a2.122 2.122 0 0 0-1.973 0L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.122 2.122 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.122 2.122 0 0 0 1.597-1.16z"/>`)

// StepBack is [step-back](https://lucide.dev/icons/step-back).
var StepBack = icon.Lucide(viewBox, paint, `<path d="M13.971 4.285A2 2 0 0 1 17 6v12a2 2 0 0 1-3.029 1.715l-9.997-5.998a2 2 0 0 1-.003-3.432z"/><path d="M21 20V4"/>`)

// StepForward is [step-forward](https://lucide.dev/icons/step-forward).
var StepForward = icon.Lucide(viewBox, paint, `<path d="M10.029 4.285A2 2 0 0 0 7 6v12a2 2 0 0 0 3.029 1.715l9.997-5.998a2 2 0 0 0 .003-3.432z"/><path d="M3 4v16"/>`)

// Stethoscope is [stethoscope](https://lucide.dev/icons/stethoscope).
var Stethoscope = icon.Lucide(viewBox, paint, `<path d="M11 2v2"/><path d="M5 2v2"/><path d="M5 3H4a2 2 0 0 0-2 2v4a6 6 0 0 0 12 0V5a2 2 0 0 0-2-2h-1"/><path d="M8 15a6 6 0 0 0 12 0v-3"/><circle cx="20" cy="10" r="2"/>`)

// Sticker is [sticker](https://lucide.dev/icons/sticker).
var Sticker = icon.Lucide(viewBox, paint, `<path d="M21 9a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 15 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2z"/><path d="M15 3v5a1 1 0 0 0 1 1h5"/><path d="M8 13h.01"/><path d="M16 13h.01"/><path d="M10 16s.8 1 2 1c1.3 0 2-1 2-1"/>`)

// StickyNoteCheck is [sticky-note-check](https://lucide.dev/icons/sticky-note-check).
var StickyNoteCheck = icon.Lucide(viewBox, paint, `<path d="m15 19 2 2 4-4"/><path d="M15 3v5a1 1 0 0 0 1 1h5"/><path d="M21 13V9a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 15 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6.5"/>`)

// StickyNoteMinus is [sticky-note-minus](https://lucide.dev/icons/sticky-note-minus).
var StickyNoteMinus = icon.Lucide(viewBox, paint, `<path d="M15 3v5a1 1 0 0 0 1 1h5"/><path d="M21 14V9a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 15 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7.35"/><path d="M21 18h-6"/>`)

// StickyNoteOff is [sticky-note-off](https://lucide.dev/icons/sticky-note-off).
var StickyNoteOff = icon.Lucide(viewBox, paint, `<path d="M15 3v5a1 1 0 0 0 1 1h5"/><path d="m2 2 20 20"/><path d="M3.586 3.586A2 2 0 0 0 3 5v14a2 2 0 0 0 2 2h14a2 2 0 0 0 1.414-.586"/><path d="M8.656 3H15a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 21 9v6.344"/>`)

// StickyNotePlus is [sticky-note-plus](https://lucide.dev/icons/sticky-note-plus).
var StickyNotePlus = icon.Lucide(viewBox, paint, `<path d="M15 3v5a1 1 0 0 0 1 1h5"/><path d="M18 15v6"/><path d="M21 12.356V9a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 15 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7.355"/><path d="M21 18h-6"/>`)

// StickyNoteX is [sticky-note-x](https://lucide.dev/icons/sticky-note-x).
var StickyNoteX = icon.Lucide(viewBox, paint, `<path d="M15 3v5a1 1 0 0 0 1 1h5"/><path d="m16 16 5 5"/><path d="M21 12V9a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 15 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7"/><path d="m21 16-5 5"/>`)

// StickyNote is [sticky-note](https://lucide.dev/icons/sticky-note).
var StickyNote = icon.Lucide(viewBox, paint, `<path d="M21 9a2.4 2.4 0 0 0-.706-1.706l-3.588-3.588A2.4 2.4 0 0 0 15 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2z"/><path d="M15 3v5a1 1 0 0 0 1 1h5"/>`)

// StickyNotes is [sticky-notes](https://lucide.dev/icons/sticky-notes).
var StickyNotes = icon.Lucide(viewBox, paint, `<path d="M10 8a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 16 14v6a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2z"/><path d="M10 8v5a1 1 0 0 0 1 1h5"/><path d="M8 4a2 2 0 0 1 2-2h6a2.4 2.4 0 0 1 1.706.706l3.588 3.588A2.4 2.4 0 0 1 22 8v6a2 2 0 0 1-2 2"/><path d="M16 2v5a1 1 0 0 0 1 1h5"/>`)

// Stone is [stone](https://lucide.dev/icons/stone).
var Stone = icon.Lucide(viewBox, paint, `<path d="M11.264 2.205A4 4 0 0 0 6.42 4.211l-4 8a4 4 0 0 0 1.359 5.117l6 4a4 4 0 0 0 4.438 0l6-4a4 4 0 0 0 1.576-4.592l-2-6a4 4 0 0 0-2.53-2.53z"/><path d="M11.99 22 14 12l7.822 3.184"/><path d="M14 12 8.47 2.302"/>`)

// Store is [store](https://lucide.dev/icons/store).
var Store = icon.Lucide(viewBox, paint, `<path d="M15 21v-5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v5"/><path d="M17.774 10.31a1.12 1.12 0 0 0-1.549 0 2.5 2.5 0 0 1-3.451 0 1.12 1.12 0 0 0-1.548 0 2.5 2.5 0 0 1-3.452 0 1.12 1.12 0 0 0-1.549 0 2.5 2.5 0 0 1-3.77-3.248l2.889-4.184A2 2 0 0 1 7 2h10a2 2 0 0 1 1.653.873l2.895 4.192a2.5 2.5 0 0 1-3.774 3.244"/><path d="M4 10.95V19a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8.05"/>`)

// StretchHorizontal is [stretch-horizontal](https://lucide.dev/icons/stretch-horizontal).
var StretchHorizontal = icon.Lucide(viewBox, paint, `<rect width="20" height="6" x="2" y="4" rx="2"/><rect width="20" height="6" x="2" y="14" rx="2"/>`)

// StretchVertical is [stretch-vertical](https://lucide.dev/icons/stretch-vertical).
var StretchVertical = icon.Lucide(viewBox, paint, `<rect width="6" height="20" x="4" y="2" rx="2"/><rect width="6" height="20" x="14" y="2" rx="2"/>`)

// Strikethrough is [strikethrough](https://lucide.dev/icons/strikethrough).
var Strikethrough = icon.Lucide(viewBox, paint, `<path d="M16 4H9a3 3 0 0 0-2.83 4"/><path d="M14 12a4 4 0 0 1 0 8H6"/><line x1="4" x2="20" y1="12" y2="12"/>`)

// Subscript is [subscript](https://lucide.dev/icons/subscript).
var Subscript = icon.Lucide(viewBox, paint, `<path d="m4 5 8 8"/><path d="m12 5-8 8"/><path d="M20 19h-4c0-1.5.44-2 1.5-2.5S20 15.33 20 14c0-.47-.17-.93-.48-1.29a2.11 2.11 0 0 0-2.62-.44c-.42.24-.74.62-.9 1.07"/>`)

// Summary is [summary](https://lucide.dev/icons/summary).
var Summary = icon.Lucide(viewBox, paint, `<path d="M15 4H7"/><path d="m18 16 3 3-3 3"/><path d="M3 4v13a2 2 0 0 0 2 2h16"/><path d="M7 14h7"/><path d="M7 9h12"/>`)

// SunDim is [sun-dim](https://lucide.dev/icons/sun-dim).
var SunDim = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="4"/><path d="M12 4h.01"/><path d="M20 12h.01"/><path d="M12 20h.01"/><path d="M4 12h.01"/><path d="M17.657 6.343h.01"/><path d="M17.657 17.657h.01"/><path d="M6.343 17.657h.01"/><path d="M6.343 6.343h.01"/>`)

// SunMedium is [sun-medium](https://lucide.dev/icons/sun-medium).
var SunMedium = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="4"/><path d="M12 3v1"/><path d="M12 20v1"/><path d="M3 12h1"/><path d="M20 12h1"/><path d="m18.364 5.636-.707.707"/><path d="m6.343 17.657-.707.707"/><path d="m5.636 5.636.707.707"/><path d="m17.657 17.657.707.707"/>`)

// SunMoon is [sun-moon](https://lucide.dev/icons/sun-moon).
var SunMoon = icon.Lucide(viewBox, paint, `<path d="M12 2v2"/><path d="M14.837 16.385a6 6 0 1 1-7.223-7.222c.624-.147.97.66.715 1.248a4 4 0 0 0 5.26 5.259c.589-.255 1.396.09 1.248.715"/><path d="M16 12a4 4 0 0 0-4-4"/><path d="m19 5-1.256 1.256"/><path d="M20 12h2"/>`)

// SunSnow is [sun-snow](https://lucide.dev/icons/sun-snow).
var SunSnow = icon.Lucide(viewBox, paint, `<path d="M10 21v-1"/><path d="M10 4V3"/><path d="M10 9a3 3 0 0 0 0 6"/><path d="m14 20 1.25-2.5L18 18"/><path d="m14 4 1.25 2.5L18 6"/><path d="m17 21-3-6 1.5-3H22"/><path d="m17 3-3 6 1.5 3"/><path d="M2 12h1"/><path d="m20 10-1.5 2 1.5 2"/><path d="m3.64 18.36.7-.7"/><path d="m4.34 6.34-.7-.7"/>`)

// Sun is [sun](https://lucide.dev/icons/sun).
var Sun = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="4"/><path d="M12 2v2"/><path d="M12 20v2"/><path d="m4.93 4.93 1.41 1.41"/><path d="m17.66 17.66 1.41 1.41"/><path d="M2 12h2"/><path d="M20 12h2"/><path d="m6.34 17.66-1.41 1.41"/><path d="m19.07 4.93-1.41 1.41"/>`)

// Sunrise is [sunrise](https://lucide.dev/icons/sunrise).
var Sunrise = icon.Lucide(viewBox, paint, `<path d="M12 2v8"/><path d="m4.93 10.93 1.41 1.41"/><path d="M2 18h2"/><path d="M20 18h2"/><path d="m19.07 10.93-1.41 1.41"/><path d="M22 22H2"/><path d="m8 6 4-4 4 4"/><path d="M16 18a4 4 0 0 0-8 0"/>`)

// Sunset is [sunset](https://lucide.dev/icons/sunset).
var Sunset = icon.Lucide(viewBox, paint, `<path d="M12 10V2"/><path d="m4.93 10.93 1.41 1.41"/><path d="M2 18h2"/><path d="M20 18h2"/><path d="m19.07 10.93-1.41 1.41"/><path d="M22 22H2"/><path d="m16 6-4 4-4-4"/><path d="M16 18a4 4 0 0 0-8 0"/>`)

// Superscript is [superscript](https://lucide.dev/icons/superscript).
var Superscript = icon.Lucide(viewBox, paint, `<path d="m4 19 8-8"/><path d="m12 19-8-8"/><path d="M20 12h-4c0-1.5.442-2 1.5-2.5S20 8.334 20 7.002c0-.472-.17-.93-.484-1.29a2.105 2.105 0 0 0-2.617-.436c-.42.239-.738.614-.899 1.06"/>`)

// SwatchBook is [swatch-book](https://lucide.dev/icons/swatch-book).
var SwatchBook = icon.Lucide(viewBox, paint, `<path d="M11 17a4 4 0 0 1-8 0V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2Z"/><path d="M16.7 13H19a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H7"/><path d="M 7 17h.01"/><path d="m11 8 2.3-2.3a2.4 2.4 0 0 1 3.404.004L18.6 7.6a2.4 2.4 0 0 1 .026 3.434L9.9 19.8"/>`)

// SwissFranc is [swiss-franc](https://lucide.dev/icons/swiss-franc).
var SwissFranc = icon.Lucide(viewBox, paint, `<path d="M10 21V3h8"/><path d="M6 16h9"/><path d="M10 9.5h7"/>`)

// SwitchCamera is [switch-camera](https://lucide.dev/icons/switch-camera).
var SwitchCamera = icon.Lucide(viewBox, paint, `<path d="M11 19H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h5"/><path d="M13 5h7a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-5"/><circle cx="12" cy="12" r="3"/><path d="m18 22-3-3 3-3"/><path d="m6 2 3 3-3 3"/>`)

// Sword is [sword](https://lucide.dev/icons/sword).
var Sword = icon.Lucide(viewBox, paint, `<path d="m11 19-6-6"/><path d="m5 21-2-2"/><path d="m8 16-4 4"/><path d="M9.5 17.5 20.414 6.586A2 2 0 0021 5.172V3h-2.172a2 2 0 00-1.414.586L6.5 14.5"/>`)

// Swords is [swords](https://lucide.dev/icons/swords).
var Swords = icon.Lucide(viewBox, paint, `<path d="m13 19 6-6"/><path d="M14.5 17.5 3.586 6.586A2 2 0 013 5.172V3h2.172a2 2 0 011.414.586L17.5 14.5"/><path d="m14.828 6.172 2.586-2.586A2 2 0 0118.828 3H21v2.172a2 2 0 01-.586 1.414l-2.586 2.586"/><path d="m16 16 4 4"/><path d="m19 21 2-2"/><path d="m5 14 4 4"/><path d="m5 21-2-2"/><path d="M7.5 16.5 4 20"/>`)

// Syringe is [syringe](https://lucide.dev/icons/syringe).
var Syringe = icon.Lucide(viewBox, paint, `<path d="m18 2 4 4"/><path d="m17 7 3-3"/><path d="M19 9 8.7 19.3c-1 1-2.5 1-3.4 0l-.6-.6c-1-1-1-2.5 0-3.4L15 5"/><path d="m9 11 4 4"/><path d="m5 19-3 3"/><path d="m14 4 6 6"/>`)

// Table2 is [table-2](https://lucide.dev/icons/table-2).
var Table2 = icon.Lucide(viewBox, paint, `<path d="M3 9h18"/><path d="M9 3v18"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// TableCellsMerge is [table-cells-merge](https://lucide.dev/icons/table-cells-merge).
var TableCellsMerge = icon.Lucide(viewBox, paint, `<path d="M12 21v-6"/><path d="M12 9V3"/><path d="M3 15h18"/><path d="M3 9h18"/><rect width="18" height="18" x="3" y="3" rx="2"/>`)

// TableCellsSplit is [table-cells-split](https://lucide.dev/icons/table-cells-split).
var TableCellsSplit = icon.Lucide(viewBox, paint, `<path d="M12 15V9"/><path d="M3 15h18"/><path d="M3 9h18"/><rect width="18" height="18" x="3" y="3" rx="2"/>`)

// TableColumnsSplit is [table-columns-split](https://lucide.dev/icons/table-columns-split).
var TableColumnsSplit = icon.Lucide(viewBox, paint, `<path d="M14 14v2"/><path d="M14 20v2"/><path d="M14 2v2"/><path d="M14 8v2"/><path d="M2 15h8"/><path d="M2 3h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H2"/><path d="M2 9h8"/><path d="M22 15h-4"/><path d="M22 3h-2a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h2"/><path d="M22 9h-4"/><path d="M5 3v18"/>`)

// TableOfContents is [table-of-contents](https://lucide.dev/icons/table-of-contents).
var TableOfContents = icon.Lucide(viewBox, paint, `<path d="M16 5H3"/><path d="M16 12H3"/><path d="M16 19H3"/><path d="M21 5h.01"/><path d="M21 12h.01"/><path d="M21 19h.01"/>`)

// TableProperties is [table-properties](https://lucide.dev/icons/table-properties).
var TableProperties = icon.Lucide(viewBox, paint, `<path d="M15 3v18"/><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M21 9H3"/><path d="M21 15H3"/>`)

// TableRowsSplit is [table-rows-split](https://lucide.dev/icons/table-rows-split).
var TableRowsSplit = icon.Lucide(viewBox, paint, `<path d="M14 10h2"/><path d="M15 22v-8"/><path d="M15 2v4"/><path d="M2 10h2"/><path d="M20 10h2"/><path d="M3 19h18"/><path d="M3 22v-6a2 2 135 0 1 2-2h14a2 2 45 0 1 2 2v6"/><path d="M3 2v2a2 2 45 0 0 2 2h14a2 2 135 0 0 2-2V2"/><path d="M8 10h2"/><path d="M9 22v-8"/><path d="M9 2v4"/>`)

// Table is [table](https://lucide.dev/icons/table).
var Table = icon.Lucide(viewBox, paint, `<path d="M12 3v18"/><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18"/><path d="M3 15h18"/>`)

// TabletSmartphone is [tablet-smartphone](https://lucide.dev/icons/tablet-smartphone).
var TabletSmartphone = icon.Lucide(viewBox, paint, `<rect width="10" height="14" x="3" y="8" rx="2"/><path d="M5 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-2.4"/><path d="M8 18h.01"/>`)

// Tablet is [tablet](https://lucide.dev/icons/tablet).
var Tablet = icon.Lucide(viewBox, paint, `<rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><line x1="12" x2="12.01" y1="18" y2="18"/>`)

// Tablets is [tablets](https://lucide.dev/icons/tablets).
var Tablets = icon.Lucide(viewBox, paint, `<circle cx="7" cy="7" r="5"/><circle cx="17" cy="17" r="5"/><path d="M12 17h10"/><path d="m3.46 10.54 7.08-7.08"/>`)

// TagPlus is [tag-plus](https://lucide.dev/icons/tag-plus).
var TagPlus = icon.Lucide(viewBox, paint, `<path d="M16 13h6"/><path d="m16.5 6.5-3.914-3.914A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l1.79-1.79"/><path d="M19 10v6"/><circle cx="7.5" cy="7.5" r=".5" fill="currentColor"/>`)

// TagX is [tag-x](https://lucide.dev/icons/tag-x).
var TagX = icon.Lucide(viewBox, paint, `<path d="m16.5 6.5-3.914-3.914A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.43 2.43 0 0 0 3.42 0l1.79-1.79"/><path d="m16.5 10.5 5 5"/><path d="m21.5 10.5-5 5"/><circle cx="7.5" cy="7.5" r=".5" fill="currentColor"/>`)

// Tag is [tag](https://lucide.dev/icons/tag).
var Tag = icon.Lucide(viewBox, paint, `<path d="M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8.704 8.704a2.426 2.426 0 0 0 3.42 0l6.58-6.58a2.426 2.426 0 0 0 0-3.42z"/><circle cx="7.5" cy="7.5" r=".5" fill="currentColor"/>`)

// Tags is [tags](https://lucide.dev/icons/tags).
var Tags = icon.Lucide(viewBox, paint, `<path d="M13.172 2a2 2 0 0 1 1.414.586l6.71 6.71a2.4 2.4 0 0 1 0 3.408l-4.592 4.592a2.4 2.4 0 0 1-3.408 0l-6.71-6.71A2 2 0 0 1 6 9.172V3a1 1 0 0 1 1-1z"/><path d="M2 7v6.172a2 2 0 0 0 .586 1.414l6.71 6.71a2.4 2.4 0 0 0 3.191.193"/><circle cx="10.5" cy="6.5" r=".5" fill="currentColor"/>`)

// Tally1 is [tally-1](https://lucide.dev/icons/tally-1).
var Tally1 = icon.Lucide(viewBox, paint, `<path d="M4 4v16"/>`)

// Tally2 is [tally-2](https://lucide.dev/icons/tally-2).
var Tally2 = icon.Lucide(viewBox, paint, `<path d="M4 4v16"/><path d="M9 4v16"/>`)

// Tally3 is [tally-3](https://lucide.dev/icons/tally-3).
var Tally3 = icon.Lucide(viewBox, paint, `<path d="M4 4v16"/><path d="M9 4v16"/><path d="M14 4v16"/>`)

// Tally4 is [tally-4](https://lucide.dev/icons/tally-4).
var Tally4 = icon.Lucide(viewBox, paint, `<path d="M4 4v16"/><path d="M9 4v16"/><path d="M14 4v16"/><path d="M19 4v16"/>`)

// Tally5 is [tally-5](https://lucide.dev/icons/tally-5).
var Tally5 = icon.Lucide(viewBox, paint, `<path d="M4 4v16"/><path d="M9 4v16"/><path d="M14 4v16"/><path d="M19 4v16"/><path d="M22 6 2 18"/>`)

// Tangent is [tangent](https://lucide.dev/icons/tangent).
var Tangent = icon.Lucide(viewBox, paint, `<circle cx="17" cy="4" r="2"/><path d="M15.59 5.41 5.41 15.59"/><circle cx="4" cy="17" r="2"/><path d="M12 22s-4-9-1.5-11.5S22 12 22 12"/>`)

// Target is [target](https://lucide.dev/icons/target).
var Target = icon.Lucide(viewBox, paint, `<circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="6"/><circle cx="12" cy="12" r="2"/>`)

// Telescope is [telescope](https://lucide.dev/icons/telescope).
var Telescope = icon.Lucide(viewBox, paint, `<path d="m10.065 12.493-6.18 1.318a.934.934 0 0 1-1.108-.702l-.537-2.15a1.07 1.07 0 0 1 .691-1.265l13.504-4.44"/><path d="m13.56 11.747 4.332-.924"/><path d="m16 21-3.105-6.21"/><path d="M16.485 5.94a2 2 0 0 1 1.455-2.425l1.09-.272a1 1 0 0 1 1.212.727l1.515 6.06a1 1 0 0 1-.727 1.213l-1.09.272a2 2 0 0 1-2.425-1.455z"/><path d="m6.158 8.633 1.114 4.456"/><path d="m8 21 3.105-6.21"/><circle cx="12" cy="13" r="2"/>`)

// TentTree is [tent-tree](https://lucide.dev/icons/tent-tree).
var TentTree = icon.Lucide(viewBox, paint, `<circle cx="4" cy="4" r="2"/><path d="m14 5 3-3 3 3"/><path d="m14 10 3-3 3 3"/><path d="M17 14V2"/><path d="M17 14H7l-5 8h20Z"/><path d="M8 14v8"/><path d="m9 14 5 8"/>`)

// Tent is [tent](https://lucide.dev/icons/tent).
var Tent = icon.Lucide(viewBox, paint, `<path d="M3.5 21 14 3"/><path d="M20.5 21 10 3"/><path d="M15.5 21 12 15l-3.5 6"/><path d="M2 21h20"/>`)

// Terminal is [terminal](https://lucide.dev/icons/terminal).
var Terminal = icon.Lucide(viewBox, paint, `<path d="M12 19h8"/><path d="m4 17 6-6-6-6"/>`)

// TestTubeDiagonal is [test-tube-diagonal](https://lucide.dev/icons/test-tube-diagonal).
var TestTubeDiagonal = icon.Lucide(viewBox, paint, `<path d="M21 7 6.82 21.18a2.83 2.83 0 0 1-3.99-.01a2.83 2.83 0 0 1 0-4L17 3"/><path d="m16 2 6 6"/><path d="M12 16H4"/>`)

// TestTube is [test-tube](https://lucide.dev/icons/test-tube).
var TestTube = icon.Lucide(viewBox, paint, `<path d="M14.5 2v17.5c0 1.4-1.1 2.5-2.5 2.5c-1.4 0-2.5-1.1-2.5-2.5V2"/><path d="M8.5 2h7"/><path d="M14.5 16h-5"/>`)

// TestTubes is [test-tubes](https://lucide.dev/icons/test-tubes).
var TestTubes = icon.Lucide(viewBox, paint, `<path d="M9 2v17.5A2.5 2.5 0 0 1 6.5 22A2.5 2.5 0 0 1 4 19.5V2"/><path d="M20 2v17.5a2.5 2.5 0 0 1-2.5 2.5a2.5 2.5 0 0 1-2.5-2.5V2"/><path d="M3 2h7"/><path d="M14 2h7"/><path d="M9 16H4"/><path d="M20 16h-5"/>`)

// TextAlignCenter is [text-align-center](https://lucide.dev/icons/text-align-center).
var TextAlignCenter = icon.Lucide(viewBox, paint, `<path d="M21 5H3"/><path d="M17 12H7"/><path d="M19 19H5"/>`)

// TextAlignEnd is [text-align-end](https://lucide.dev/icons/text-align-end).
var TextAlignEnd = icon.Lucide(viewBox, paint, `<path d="M21 5H3"/><path d="M21 12H9"/><path d="M21 19H7"/>`)

// TextAlignJustifyCenter is [text-align-justify-center](https://lucide.dev/icons/text-align-justify-center).
var TextAlignJustifyCenter = icon.Lucide(viewBox, paint, `<path d="M3 5h18"/><path d="M3 12h18"/><path d="M7 19h10"/>`)

// TextAlignJustifyEnd is [text-align-justify-end](https://lucide.dev/icons/text-align-justify-end).
var TextAlignJustifyEnd = icon.Lucide(viewBox, paint, `<path d="M3 5h18"/><path d="M3 12h18"/><path d="M11 19h10"/>`)

// TextAlignJustifyStart is [text-align-justify-start](https://lucide.dev/icons/text-align-justify-start).
var TextAlignJustifyStart = icon.Lucide(viewBox, paint, `<path d="M3 5h18"/><path d="M3 12h18"/><path d="M3 19h10"/>`)

// TextAlignJustify is [text-align-justify](https://lucide.dev/icons/text-align-justify).
var TextAlignJustify = icon.Lucide(viewBox, paint, `<path d="M3 5h18"/><path d="M3 12h18"/><path d="M3 19h18"/>`)

// TextAlignStart is [text-align-start](https://lucide.dev/icons/text-align-start).
var TextAlignStart = icon.Lucide(viewBox, paint, `<path d="M21 5H3"/><path d="M15 12H3"/><path d="M17 19H3"/>`)

// TextCursorInput is [text-cursor-input](https://lucide.dev/icons/text-cursor-input).
var TextCursorInput = icon.Lucide(viewBox, paint, `<path d="M12 20h-1a2 2 0 0 1-2-2 2 2 0 0 1-2 2H6"/><path d="M13 8h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-7"/><path d="M5 16H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h1"/><path d="M6 4h1a2 2 0 0 1 2 2 2 2 0 0 1 2-2h1"/><path d="M9 6v12"/>`)

// TextCursor is [text-cursor](https://lucide.dev/icons/text-cursor).
var TextCursor = icon.Lucide(viewBox, paint, `<path d="M17 22h-1a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h1"/><path d="M7 22h1a4 4 0 0 0 4-4"/><path d="M7 2h1a4 4 0 0 1 4 4"/>`)

// TextInitial is [text-initial](https://lucide.dev/icons/text-initial).
var TextInitial = icon.Lucide(viewBox, paint, `<path d="M15 5h6"/><path d="M15 12h6"/><path d="M3 19h18"/><path d="m3 12 3.553-7.724a.5.5 0 0 1 .894 0L11 12"/><path d="M3.92 10h6.16"/>`)

// TextQuote is [text-quote](https://lucide.dev/icons/text-quote).
var TextQuote = icon.Lucide(viewBox, paint, `<path d="M17 5H3"/><path d="M21 12H8"/><path d="M21 19H8"/><path d="M3 12v7"/>`)

// TextSearch is [text-search](https://lucide.dev/icons/text-search).
var TextSearch = icon.Lucide(viewBox, paint, `<path d="M21 5H3"/><path d="M10 12H3"/><path d="M10 19H3"/><circle cx="17" cy="15" r="3"/><path d="m21 19-1.9-1.9"/>`)

// TextWrap is [text-wrap](https://lucide.dev/icons/text-wrap).
var TextWrap = icon.Lucide(viewBox, paint, `<path d="m16 16-3 3 3 3"/><path d="M3 12h14.5a1 1 0 0 1 0 7H13"/><path d="M3 19h6"/><path d="M3 5h18"/>`)

// Theater is [theater](https://lucide.dev/icons/theater).
var Theater = icon.Lucide(viewBox, paint, `<path d="M2 10s3-3 3-8"/><path d="M22 10s-3-3-3-8"/><path d="M10 2c0 4.4-3.6 8-8 8"/><path d="M14 2c0 4.4 3.6 8 8 8"/><path d="M2 10s2 2 2 5"/><path d="M22 10s-2 2-2 5"/><path d="M8 15h8"/><path d="M2 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1"/><path d="M14 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1"/>`)

// ThermometerSnowflake is [thermometer-snowflake](https://lucide.dev/icons/thermometer-snowflake).
var ThermometerSnowflake = icon.Lucide(viewBox, paint, `<path d="m10 20-1.25-2.5L6 18"/><path d="M10 4 8.75 6.5 6 6"/><path d="M10.585 15H10"/><path d="M2 12h6.5L10 9"/><path d="M20 14.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0z"/><path d="m4 10 1.5 2L4 14"/><path d="m7 21 3-6-1.5-3"/><path d="m7 3 3 6h2"/>`)

// ThermometerSun is [thermometer-sun](https://lucide.dev/icons/thermometer-sun).
var ThermometerSun = icon.Lucide(viewBox, paint, `<path d="M12 2v2"/><path d="M12 8a4 4 0 0 0-1.645 7.647"/><path d="M2 12h2"/><path d="M20 14.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0z"/><path d="m4.93 4.93 1.41 1.41"/><path d="m6.34 17.66-1.41 1.41"/>`)

// Thermometer is [thermometer](https://lucide.dev/icons/thermometer).
var Thermometer = icon.Lucide(viewBox, paint, `<path d="M14 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z"/>`)

// ThumbsDown is [thumbs-down](https://lucide.dev/icons/thumbs-down).
var ThumbsDown = icon.Lucide(viewBox, paint, `<path d="M9 18.12 10 14H4.17a2 2 0 0 1-1.92-2.56l2.33-8A2 2 0 0 1 6.5 2H20a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2.76a2 2 0 0 0-1.79 1.11L12 22a3.13 3.13 0 0 1-3-3.88Z"/><path d="M17 14V2"/>`)

// ThumbsUp is [thumbs-up](https://lucide.dev/icons/thumbs-up).
var ThumbsUp = icon.Lucide(viewBox, paint, `<path d="M15 5.88 14 10h5.83a2 2 0 0 1 1.92 2.56l-2.33 8A2 2 0 0 1 17.5 22H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h2.76a2 2 0 0 0 1.79-1.11L12 2a3.13 3.13 0 0 1 3 3.88Z"/><path d="M7 10v12"/>`)

// TicTacToe is [tic-tac-toe](https://lucide.dev/icons/tic-tac-toe).
var TicTacToe = icon.Lucide(viewBox, paint, `<path d="M12 2v20"/><path d="m21 16-5 5"/><path d="m21 21-5-5"/><path d="M22 12H2"/><path d="M8 3 3 8"/><path d="M8 8 3 3"/><circle cx="18.5" cy="5.5" r="2.5"/><circle cx="5.5" cy="18.5" r="2.5"/>`)

// TicketCheck is [ticket-check](https://lucide.dev/icons/ticket-check).
var TicketCheck = icon.Lucide(viewBox, paint, `<path d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"/><path d="m9 12 2 2 4-4"/>`)

// TicketMinus is [ticket-minus](https://lucide.dev/icons/ticket-minus).
var TicketMinus = icon.Lucide(viewBox, paint, `<path d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"/><path d="M9 12h6"/>`)

// TicketPercent is [ticket-percent](https://lucide.dev/icons/ticket-percent).
var TicketPercent = icon.Lucide(viewBox, paint, `<path d="M2 9a3 3 0 1 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 1 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"/><path d="M9 9h.01"/><path d="m15 9-6 6"/><path d="M15 15h.01"/>`)

// TicketPlus is [ticket-plus](https://lucide.dev/icons/ticket-plus).
var TicketPlus = icon.Lucide(viewBox, paint, `<path d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"/><path d="M9 12h6"/><path d="M12 9v6"/>`)

// TicketSlash is [ticket-slash](https://lucide.dev/icons/ticket-slash).
var TicketSlash = icon.Lucide(viewBox, paint, `<path d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"/><path d="m9.5 14.5 5-5"/>`)

// TicketX is [ticket-x](https://lucide.dev/icons/ticket-x).
var TicketX = icon.Lucide(viewBox, paint, `<path d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"/><path d="m9.5 14.5 5-5"/><path d="m9.5 9.5 5 5"/>`)

// Ticket is [ticket](https://lucide.dev/icons/ticket).
var Ticket = icon.Lucide(viewBox, paint, `<path d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"/><path d="M13 5v2"/><path d="M13 17v2"/><path d="M13 11v2"/>`)

// TicketsPlane is [tickets-plane](https://lucide.dev/icons/tickets-plane).
var TicketsPlane = icon.Lucide(viewBox, paint, `<path d="M10.5 17h1.227a2 2 0 0 0 1.345-.52L18 12"/><path d="m12 13.5 3.794.506"/><path d="m3.173 8.18 11-5a2 2 0 0 1 2.647.993L18.56 8"/><path d="M6 10V8"/><path d="M6 14v1"/><path d="M6 19v2"/><rect x="2" y="8" width="20" height="13" rx="2"/>`)

// Tickets is [tickets](https://lucide.dev/icons/tickets).
var Tickets = icon.Lucide(viewBox, paint, `<path d="m3.173 8.18 11-5a2 2 0 0 1 2.647.993L18.56 8"/><path d="M6 10V8"/><path d="M6 14v1"/><path d="M6 19v2"/><rect x="2" y="8" width="20" height="13" rx="2"/>`)

// Timeline is [timeline](https://lucide.dev/icons/timeline).
var Timeline = icon.Lucide(viewBox, paint, `<path d="M4 12h.01"/><path d="M4 16h.01"/><path d="M4 20h.01"/><path d="M4 4h.01"/><path d="M4 8h.01"/><path d="M9.414 13.414a2 2 0 0 0 1.414.586H19a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-8.172a2 2 0 0 0-1.414.586L8 12z"/><path d="M9.414 21.414a2 2 0 0 0 1.414.586H19a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-8.172a2 2 0 0 0-1.414.586L8 20z"/><path d="M9.414 5.414A2 2 0 0 0 10.828 6H19a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1h-8.172a2 2 0 0 0-1.414.586L8 4z"/>`)

// TimerOff is [timer-off](https://lucide.dev/icons/timer-off).
var TimerOff = icon.Lucide(viewBox, paint, `<path d="M10 2h4"/><path d="M4.6 11a8 8 0 0 0 1.7 8.7 8 8 0 0 0 8.7 1.7"/><path d="M7.4 7.4a8 8 0 0 1 10.3 1 8 8 0 0 1 .9 10.2"/><path d="m2 2 20 20"/><path d="M12 12v-2"/>`)

// TimerReset is [timer-reset](https://lucide.dev/icons/timer-reset).
var TimerReset = icon.Lucide(viewBox, paint, `<path d="M10 2h4"/><path d="M12 14v-4"/><path d="M4 13a8 8 0 0 1 8-7 8 8 0 1 1-5.3 14L4 17.6"/><path d="M9 17H4v5"/>`)

// Timer is [timer](https://lucide.dev/icons/timer).
var Timer = icon.Lucide(viewBox, paint, `<line x1="10" x2="14" y1="2" y2="2"/><line x1="12" x2="15" y1="14" y2="11"/><circle cx="12" cy="14" r="8"/>`)

// ToggleLeft is [toggle-left](https://lucide.dev/icons/toggle-left).
var ToggleLeft = icon.Lucide(viewBox, paint, `<circle cx="9" cy="12" r="3"/><rect width="20" height="14" x="2" y="5" rx="7"/>`)

// ToggleRight is [toggle-right](https://lucide.dev/icons/toggle-right).
var ToggleRight = icon.Lucide(viewBox, paint, `<circle cx="15" cy="12" r="3"/><rect width="20" height="14" x="2" y="5" rx="7"/>`)

// Toilet is [toilet](https://lucide.dev/icons/toilet).
var Toilet = icon.Lucide(viewBox, paint, `<path d="M7 12h13a1 1 0 0 1 1 1 5 5 0 0 1-5 5h-.598a.5.5 0 0 0-.424.765l1.544 2.47a.5.5 0 0 1-.424.765H5.402a.5.5 0 0 1-.424-.765L7 18"/><path d="M8 18a5 5 0 0 1-5-5V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8"/>`)

// Tombstone is [tombstone](https://lucide.dev/icons/tombstone).
var Tombstone = icon.Lucide(viewBox, paint, `<path d="M11 10h2"/><path d="M19 21V10a7 7 0 00-14 0v11"/><path d="M2 21h20"/><path d="M9 14h6"/>`)

// ToolCase is [tool-case](https://lucide.dev/icons/tool-case).
var ToolCase = icon.Lucide(viewBox, paint, `<path d="M10 15h4"/><path d="m14.5 4.058.571-1.429a1 1 0 011.3-.558l5 2a1 1 0 01.549.534.88.88 0 01.01.765l-2.332 5.83"/><path d="M15.634 11a2 2 0 00-.186-.34l-1.17-1.757 1.31-1.656a2 2 0 00-2.109-3.165l-2.032.57-1.17-1.758a2 2 0 00-3.664 1.028l-.086 2.11-2.032.568a2 2 0 00-.155 3.8l1.617.6"/><path d="M4 12.006A1 1 0 014.994 11H19a1 1 0 011 1v7a2 2 0 01-2 2H6a2 2 0 01-2-2z"/>`)

// Toolbox is [toolbox](https://lucide.dev/icons/toolbox).
var Toolbox = icon.Lucide(viewBox, paint, `<path d="M16 12v4"/><path d="M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2"/><path d="M17 6a2 2 0 011.414.586l3 3A2 2 0 0122 11v8a2 2 0 01-2 2H4a2 2 0 01-2-2v-8a2 2 0 01.586-1.414l3-3A2 2 0 017 6z"/><path d="M2 14h20"/><path d="M8 12v4"/>`)

// ToothbrushSparkles is [toothbrush-sparkles](https://lucide.dev/icons/toothbrush-sparkles).
var ToothbrushSparkles = icon.Lucide(viewBox, paint, `<path d="M10 3H8"/><path d="M14.586 11.414 22 4"/><path d="M15 11a8 8 0 01-.429.4"/><path d="m2 22 7-7c1.857-1.857 3.714-1.99 5.571-3.6l-1.985-1.986A2 2 0 0114 6a2 2 0 012-2 2 2 0 013.262-1.552l2.152 2.138"/><path d="M20 15v4"/><path d="M22 17h-4"/><path d="M4 5v4"/><path d="M6 7H2"/><path d="M9 2v2"/>`)

// Toothbrush is [toothbrush](https://lucide.dev/icons/toothbrush).
var Toothbrush = icon.Lucide(viewBox, paint, `<path d="M15 11c-2 2-4 2-6 4l-7 7"/><path d="m22 4-7.414 7.414-2-2A2 2 0 0114 6c0-.512.196-1.024.586-1.414A2 2 0 0116 4a2 2 0 013.262-1.552l2.152 2.138"/>`)

// Tornado is [tornado](https://lucide.dev/icons/tornado).
var Tornado = icon.Lucide(viewBox, paint, `<path d="M21 4H3"/><path d="M18 8H6"/><path d="M19 12H9"/><path d="M16 16h-6"/><path d="M11 20H9"/>`)

// Torus is [torus](https://lucide.dev/icons/torus).
var Torus = icon.Lucide(viewBox, paint, `<ellipse cx="12" cy="11" rx="3" ry="2"/><ellipse cx="12" cy="12.5" rx="10" ry="8.5"/>`)

// TouchpadOff is [touchpad-off](https://lucide.dev/icons/touchpad-off).
var TouchpadOff = icon.Lucide(viewBox, paint, `<path d="M12 20v-6"/><path d="M19.656 14H22"/><path d="M2 14h12"/><path d="m2 2 20 20"/><path d="M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2"/><path d="M9.656 4H20a2 2 0 0 1 2 2v10.344"/>`)

// Touchpad is [touchpad](https://lucide.dev/icons/touchpad).
var Touchpad = icon.Lucide(viewBox, paint, `<rect width="20" height="16" x="2" y="4" rx="2"/><path d="M2 14h20"/><path d="M12 20v-6"/>`)

// TowelRack is [towel-rack](https://lucide.dev/icons/towel-rack).
var TowelRack = icon.Lucide(viewBox, paint, `<path d="M22 7h-2"/><path d="M6.5 3h11A2.5 2.5 0 0 1 20 5.5V20a1 1 0 0 1-1 1h-9a1 1 0 0 1-1-1V5.5a1 1 0 0 0-5 0V17a1 1 0 0 0 1 1h4"/><path d="M9 7H2"/>`)

// TowerControl is [tower-control](https://lucide.dev/icons/tower-control).
var TowerControl = icon.Lucide(viewBox, paint, `<path d="M18.2 12.27 20 6H4l1.8 6.27a1 1 0 0 0 .95.73h10.5a1 1 0 0 0 .96-.73Z"/><path d="M8 13v9"/><path d="M16 22v-9"/><path d="m9 6 1 7"/><path d="m15 6-1 7"/><path d="M12 6V2"/><path d="M13 2h-2"/>`)

// ToyBrick is [toy-brick](https://lucide.dev/icons/toy-brick).
var ToyBrick = icon.Lucide(viewBox, paint, `<rect width="18" height="12" x="3" y="8" rx="1"/><path d="M10 8V5c0-.6-.4-1-1-1H6a1 1 0 0 0-1 1v3"/><path d="M19 8V5c0-.6-.4-1-1-1h-3a1 1 0 0 0-1 1v3"/>`)

// Tractor is [tractor](https://lucide.dev/icons/tractor).
var Tractor = icon.Lucide(viewBox, paint, `<path d="m10 11 11 .9a1 1 0 0 1 .8 1.1l-.665 4.158a1 1 0 0 1-.988.842H20"/><path d="M16 18h-5"/><path d="M18 5a1 1 0 0 0-1 1v5.573"/><path d="M3 4h8.129a1 1 0 0 1 .99.863L13 11.246"/><path d="M4 11V4"/><path d="M7 15h.01"/><path d="M8 10.1V4"/><circle cx="18" cy="18" r="2"/><circle cx="7" cy="15" r="5"/>`)

// TrafficCone is [traffic-cone](https://lucide.dev/icons/traffic-cone).
var TrafficCone = icon.Lucide(viewBox, paint, `<path d="M16.05 10.966a5 2.5 0 0 1-8.1 0"/><path d="m16.923 14.049 4.48 2.04a1 1 0 0 1 .001 1.831l-8.574 3.9a2 2 0 0 1-1.66 0l-8.574-3.91a1 1 0 0 1 0-1.83l4.484-2.04"/><path d="M16.949 14.14a5 2.5 0 1 1-9.9 0L10.063 3.5a2 2 0 0 1 3.874 0z"/><path d="M9.194 6.57a5 2.5 0 0 0 5.61 0"/>`)

// Trailer is [trailer](https://lucide.dev/icons/trailer).
var Trailer = icon.Lucide(viewBox, paint, `<path d="M10 11.341V10"/><path d="M14 13v-3"/><path d="M18 17V8a2 2 0 00-2-2H4a2 2 0 00-2 2v7a2 2 0 002 2h2"/><path d="M22 15v1a1 1 0 01-1 1H10"/><path d="M6 11.341V10"/><circle cx="8" cy="17" r="2"/>`)

// TrainFrontTunnel is [train-front-tunnel](https://lucide.dev/icons/train-front-tunnel).
var TrainFrontTunnel = icon.Lucide(viewBox, paint, `<path d="M2 22V12a10 10 0 1 1 20 0v10"/><path d="M15 6.8v1.4a3 2.8 0 1 1-6 0V6.8"/><path d="M10 15h.01"/><path d="M14 15h.01"/><path d="M10 19a4 4 0 0 1-4-4v-3a6 6 0 1 1 12 0v3a4 4 0 0 1-4 4Z"/><path d="m9 19-2 3"/><path d="m15 19 2 3"/>`)

// TrainFront is [train-front](https://lucide.dev/icons/train-front).
var TrainFront = icon.Lucide(viewBox, paint, `<path d="M8 3.1V7a4 4 0 0 0 8 0V3.1"/><path d="m9 15-1-1"/><path d="m15 15 1-1"/><path d="M9 19c-2.8 0-5-2.2-5-5v-4a8 8 0 0 1 16 0v4c0 2.8-2.2 5-5 5Z"/><path d="m8 19-2 3"/><path d="m16 19 2 3"/>`)

// TrainTrack is [train-track](https://lucide.dev/icons/train-track).
var TrainTrack = icon.Lucide(viewBox, paint, `<path d="M2 17 17 2"/><path d="m2 14 8 8"/><path d="m5 11 8 8"/><path d="m8 8 8 8"/><path d="m11 5 8 8"/><path d="m14 2 8 8"/><path d="M7 22 22 7"/>`)

// TramFront is [tram-front](https://lucide.dev/icons/tram-front).
var TramFront = icon.Lucide(viewBox, paint, `<rect width="16" height="16" x="4" y="3" rx="2"/><path d="M4 11h16"/><path d="M12 3v8"/><path d="m8 19-2 3"/><path d="m18 22-2-3"/><path d="M8 15h.01"/><path d="M16 15h.01"/>`)

// Transgender is [transgender](https://lucide.dev/icons/transgender).
var Transgender = icon.Lucide(viewBox, paint, `<path d="M12 16v6"/><path d="M14 20h-4"/><path d="M18 2h4v4"/><path d="m2 2 7.17 7.17"/><path d="M2 5.355V2h3.357"/><path d="m22 2-7.17 7.17"/><path d="M8 5 5 8"/><circle cx="12" cy="12" r="4"/>`)

// TrashOff is [trash-off](https://lucide.dev/icons/trash-off).
var TrashOff = icon.Lucide(viewBox, paint, `<path d="M10 11v6"/><path d="M14 17v-3"/><path d="M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-1.576.768"/><path d="M19 6v7.344"/><path d="m2 2 20 20"/><path d="M21 6h-9.344"/><path d="M3 6h3"/><path d="M5 6v14a2 2 0 002 2h10a2 2 0 002-2v-1"/>`)

// Trash is [trash](https://lucide.dev/icons/trash).
var Trash = icon.Lucide(viewBox, paint, `<path d="M10 11v6"/><path d="M14 11v6"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"/><path d="M3 6h18"/><path d="M8 6V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>`)

// TreeDeciduous is [tree-deciduous](https://lucide.dev/icons/tree-deciduous).
var TreeDeciduous = icon.Lucide(viewBox, paint, `<path d="M8 19a4 4 0 0 1-2.24-7.32A3.5 3.5 0 0 1 9 6.03V6a3 3 0 1 1 6 0v.04a3.5 3.5 0 0 1 3.24 5.65A4 4 0 0 1 16 19Z"/><path d="M12 19v3"/>`)

// TreePalm is [tree-palm](https://lucide.dev/icons/tree-palm).
var TreePalm = icon.Lucide(viewBox, paint, `<path d="M13 8c0-2.76-2.46-5-5.5-5S2 5.24 2 8h2l1-1 1 1h4"/><path d="M13 7.14A5.82 5.82 0 0 1 16.5 6c3.04 0 5.5 2.24 5.5 5h-3l-1-1-1 1h-3"/><path d="M5.89 9.71c-2.15 2.15-2.3 5.47-.35 7.43l4.24-4.25.7-.7.71-.71 2.12-2.12c-1.95-1.96-5.27-1.8-7.42.35"/><path d="M11 15.5c.5 2.5-.17 4.5-1 6.5h4c2-5.5-.5-12-1-14"/>`)

// TreePine is [tree-pine](https://lucide.dev/icons/tree-pine).
var TreePine = icon.Lucide(viewBox, paint, `<path d="m17 14 3 3.3a1 1 0 0 1-.7 1.7H4.7a1 1 0 0 1-.7-1.7L7 14h-.3a1 1 0 0 1-.7-1.7L9 9h-.2A1 1 0 0 1 8 7.3L12 3l4 4.3a1 1 0 0 1-.8 1.7H15l3 3.3a1 1 0 0 1-.7 1.7H17Z"/><path d="M12 22v-3"/>`)

// Trees is [trees](https://lucide.dev/icons/trees).
var Trees = icon.Lucide(viewBox, paint, `<path d="M10 10v.2A3 3 0 0 1 8.9 16H5a3 3 0 0 1-1-5.8V10a3 3 0 0 1 6 0Z"/><path d="M7 16v6"/><path d="M13 19v3"/><path d="M12 19h8.3a1 1 0 0 0 .7-1.7L18 14h.3a1 1 0 0 0 .7-1.7L16 9h.2a1 1 0 0 0 .8-1.7L13 3l-1.4 1.5"/>`)

// TrendingDown is [trending-down](https://lucide.dev/icons/trending-down).
var TrendingDown = icon.Lucide(viewBox, paint, `<path d="M16 17h6v-6"/><path d="m22 17-8.5-8.5-5 5L2 7"/>`)

// TrendingUpDown is [trending-up-down](https://lucide.dev/icons/trending-up-down).
var TrendingUpDown = icon.Lucide(viewBox, paint, `<path d="M14.828 14.828 21 21"/><path d="M21 16v5h-5"/><path d="m21 3-9 9-4-4-6 6"/><path d="M21 8V3h-5"/>`)

// TrendingUp is [trending-up](https://lucide.dev/icons/trending-up).
var TrendingUp = icon.Lucide(viewBox, paint, `<path d="M16 7h6v6"/><path d="m22 7-8.5 8.5-5-5L2 17"/>`)

// TriangleAlert is [triangle-alert](https://lucide.dev/icons/triangle-alert).
var TriangleAlert = icon.Lucide(viewBox, paint, `<path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3"/><path d="M12 9v4"/><path d="M12 17h.01"/>`)

// TriangleDashed is [triangle-dashed](https://lucide.dev/icons/triangle-dashed).
var TriangleDashed = icon.Lucide(viewBox, paint, `<path d="M10.17 4.193a2 2 0 0 1 3.666.013"/><path d="M14 21h2"/><path d="m15.874 7.743 1 1.732"/><path d="m18.849 12.952 1 1.732"/><path d="M21.824 18.18a2 2 0 0 1-1.835 2.824"/><path d="M4.024 21a2 2 0 0 1-1.839-2.839"/><path d="m5.136 12.952-1 1.732"/><path d="M8 21h2"/><path d="m8.102 7.743-1 1.732"/>`)

// TriangleRight is [triangle-right](https://lucide.dev/icons/triangle-right).
var TriangleRight = icon.Lucide(viewBox, paint, `<path d="M22 18a2 2 0 0 1-2 2H3c-1.1 0-1.3-.6-.4-1.3L20.4 4.3c.9-.7 1.6-.4 1.6.7Z"/>`)

// Triangle is [triangle](https://lucide.dev/icons/triangle).
var Triangle = icon.Lucide(viewBox, paint, `<path d="M13.73 4a2 2 0 0 0-3.46 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z"/>`)

// TrianglesCenterlineDashedHorizontal is [triangles-centerline-dashed-horizontal](https://lucide.dev/icons/triangles-centerline-dashed-horizontal).
var TrianglesCenterlineDashedHorizontal = icon.Lucide(viewBox, paint, `<path d="M10 12H8"/><path d="M16 12h-2"/><path d="M22 12h-2"/><path d="M4 12H2"/><path d="M7.298 20.288A1 1 0 008 22h8a1 1 0 00.703-1.712l-3.991-3.99a1 1 0 00-1.424-.001z"/><path d="M7.298 3.712A1 1 0 018 2h8a1 1 0 01.703 1.712l-3.991 3.99a1 1 0 01-1.424.001z"/>`)

// TrianglesCenterlineDashedVertical is [triangles-centerline-dashed-vertical](https://lucide.dev/icons/triangles-centerline-dashed-vertical).
var TrianglesCenterlineDashedVertical = icon.Lucide(viewBox, paint, `<path d="M12 14v2"/><path d="M12 20v2"/><path d="M12 2v2"/><path d="M12 8v2"/><path d="M20.288 16.703A1 1 0 0022 16V8a1 1 0 00-1.712-.703l-3.99 3.991a1 1 0 00-.001 1.424z"/><path d="M3.712 16.703A1 1 0 012 16V8a1 1 0 011.712-.703l3.99 3.991a1 1 0 01.001 1.424z"/>`)

// Trophy is [trophy](https://lucide.dev/icons/trophy).
var Trophy = icon.Lucide(viewBox, paint, `<path d="M10 14.66V17a1 1 0 0 1-1 1 2 2 0 0 0-2 2v2"/><path d="M14 14.66V17a1 1 0 0 0 1 1 2 2 0 0 1 2 2v2"/><path d="M17.916 10H19.5A2.5 2.5 0 0 0 22 7.5V5a1 1 0 0 0-1-1h-3"/><path d="M4 22h16"/><path d="M6 9a6 6 0 0 0 12 0V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1z"/><path d="M6.084 10H4.5A2.5 2.5 0 0 1 2 7.5V5a1 1 0 0 1 1-1h3"/>`)

// TruckElectric is [truck-electric](https://lucide.dev/icons/truck-electric).
var TruckElectric = icon.Lucide(viewBox, paint, `<path d="M14 19V7a2 2 0 0 0-2-2H9"/><path d="M15 19H9"/><path d="M19 19h2a1 1 0 0 0 1-1v-3.65a1 1 0 0 0-.22-.62L18.3 9.38a1 1 0 0 0-.78-.38H14"/><path d="M2 13v5a1 1 0 0 0 1 1h2"/><path d="M4 3 2.15 5.15a.495.495 0 0 0 .35.86h2.15a.47.47 0 0 1 .35.86L3 9.02"/><circle cx="17" cy="19" r="2"/><circle cx="7" cy="19" r="2"/>`)

// Truck is [truck](https://lucide.dev/icons/truck).
var Truck = icon.Lucide(viewBox, paint, `<path d="M14 18V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v11a1 1 0 0 0 1 1h2"/><path d="M15 18H9"/><path d="M19 18h2a1 1 0 0 0 1-1v-3.65a1 1 0 0 0-.22-.624l-3.48-4.35A1 1 0 0 0 17.52 8H14"/><circle cx="17" cy="18" r="2"/><circle cx="7" cy="18" r="2"/>`)

// TubeLotion is [tube-lotion](https://lucide.dev/icons/tube-lotion).
var TubeLotion = icon.Lucide(viewBox, paint, `<path d="M15 18v3a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-3"/><path d="M17 2a2 2 0 0 1 1.6 3.2A8 8 0 0 0 17 10v6a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-6a8 8 0 0 0-1.6-4.8A2 2 0 0 1 7 2z"/><path d="M7 10a6.47 6.47 0 0 1 5 0 6.47 6.47 0 0 0 5 0"/>`)

// TurkishLira is [turkish-lira](https://lucide.dev/icons/turkish-lira).
var TurkishLira = icon.Lucide(viewBox, paint, `<path d="M15 4 5 9"/><path d="m15 8.5-10 5"/><path d="M18 12a9 9 0 0 1-9 9V3"/>`)

// Turntable is [turntable](https://lucide.dev/icons/turntable).
var Turntable = icon.Lucide(viewBox, paint, `<path d="M10 12.01h.01"/><path d="M18 8v4a8 8 0 0 1-1.07 4"/><circle cx="10" cy="12" r="4"/><rect x="2" y="4" width="20" height="16" rx="2"/>`)

// Turtle is [turtle](https://lucide.dev/icons/turtle).
var Turtle = icon.Lucide(viewBox, paint, `<path d="m12 10 2 4v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3a8 8 0 1 0-16 0v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3l2-4h4Z"/><path d="M4.82 7.9 8 10"/><path d="M15.18 7.9 12 10"/><path d="M16.93 10H20a2 2 0 0 1 0 4H2"/>`)

// TvMinimalPlay is [tv-minimal-play](https://lucide.dev/icons/tv-minimal-play).
var TvMinimalPlay = icon.Lucide(viewBox, paint, `<path d="M15.033 9.44a.647.647 0 0 1 0 1.12l-4.065 2.352a.645.645 0 0 1-.968-.56V7.648a.645.645 0 0 1 .967-.56z"/><path d="M7 21h10"/><rect width="20" height="14" x="2" y="3" rx="2"/>`)

// TvMinimal is [tv-minimal](https://lucide.dev/icons/tv-minimal).
var TvMinimal = icon.Lucide(viewBox, paint, `<path d="M7 21h10"/><rect width="20" height="14" x="2" y="3" rx="2"/>`)

// Tv is [tv](https://lucide.dev/icons/tv).
var Tv = icon.Lucide(viewBox, paint, `<path d="m17 2-5 5-5-5"/><rect width="20" height="15" x="2" y="7" rx="2"/>`)

// TypeOutline is [type-outline](https://lucide.dev/icons/type-outline).
var TypeOutline = icon.Lucide(viewBox, paint, `<path d="M14 16.5a.5.5 0 0 0 .5.5h.5a2 2 0 0 1 0 4H9a2 2 0 0 1 0-4h.5a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5V8a2 2 0 0 1-4 0V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v3a2 2 0 0 1-4 0v-.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5Z"/>`)

// Type is [type](https://lucide.dev/icons/type).
var Type = icon.Lucide(viewBox, paint, `<path d="M12 4v16"/><path d="M4 7V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2"/><path d="M9 20h6"/>`)

// UmbrellaOff is [umbrella-off](https://lucide.dev/icons/umbrella-off).
var UmbrellaOff = icon.Lucide(viewBox, paint, `<path d="M12 13v7a2 2 0 0 0 4 0"/><path d="M12 2v2"/><path d="M18.656 13h2.336a1 1 0 0 0 .97-1.274 10.284 10.284 0 0 0-12.07-7.51"/><path d="m2 2 20 20"/><path d="M5.961 5.957a10.28 10.28 0 0 0-3.922 5.769A1 1 0 0 0 3 13h10"/>`)

// Umbrella is [umbrella](https://lucide.dev/icons/umbrella).
var Umbrella = icon.Lucide(viewBox, paint, `<path d="M12 13v7a2 2 0 0 0 4 0"/><path d="M12 2v2"/><path d="M20.992 13a1 1 0 0 0 .97-1.274 10.284 10.284 0 0 0-19.923 0A1 1 0 0 0 3 13z"/>`)

// Underline is [underline](https://lucide.dev/icons/underline).
var Underline = icon.Lucide(viewBox, paint, `<path d="M6 4v6a6 6 0 0 0 12 0V4"/><line x1="4" x2="20" y1="20" y2="20"/>`)

// Undo2 is [undo-2](https://lucide.dev/icons/undo-2).
var Undo2 = icon.Lucide(viewBox, paint, `<path d="M9 14 4 9l5-5"/><path d="M4 9h10.5a5.5 5.5 0 0 1 5.5 5.5a5.5 5.5 0 0 1-5.5 5.5H11"/>`)

// UndoDot is [undo-dot](https://lucide.dev/icons/undo-dot).
var UndoDot = icon.Lucide(viewBox, paint, `<path d="M21 17a9 9 0 0 0-15-6.7L3 13"/><path d="M3 7v6h6"/><circle cx="12" cy="17" r="1"/>`)

// Undo is [undo](https://lucide.dev/icons/undo).
var Undo = icon.Lucide(viewBox, paint, `<path d="M3 7v6h6"/><path d="M21 17a9 9 0 0 0-9-9 9 9 0 0 0-6 2.3L3 13"/>`)

// UnfoldHorizontal is [unfold-horizontal](https://lucide.dev/icons/unfold-horizontal).
var UnfoldHorizontal = icon.Lucide(viewBox, paint, `<path d="M16 12h6"/><path d="M8 12H2"/><path d="M12 2v2"/><path d="M12 8v2"/><path d="M12 14v2"/><path d="M12 20v2"/><path d="m19 15 3-3-3-3"/><path d="m5 9-3 3 3 3"/>`)

// UnfoldVertical is [unfold-vertical](https://lucide.dev/icons/unfold-vertical).
var UnfoldVertical = icon.Lucide(viewBox, paint, `<path d="M12 22v-6"/><path d="M12 8V2"/><path d="M4 12H2"/><path d="M10 12H8"/><path d="M16 12h-2"/><path d="M22 12h-2"/><path d="m15 19-3 3-3-3"/><path d="m15 5-3-3-3 3"/>`)

// Ungroup is [ungroup](https://lucide.dev/icons/ungroup).
var Ungroup = icon.Lucide(viewBox, paint, `<rect x="11" y="14" width="10" height="7" rx="2"/><rect x="3" y="3" width="10" height="7" rx="2"/>`)

// University is [university](https://lucide.dev/icons/university).
var University = icon.Lucide(viewBox, paint, `<path d="M14 21v-3a2 2 0 0 0-4 0v3"/><path d="M18 12h.01"/><path d="M18 16h.01"/><path d="M22 7a1 1 0 0 0-1-1h-2a2 2 0 0 1-1.143-.359L13.143 2.36a2 2 0 0 0-2.286-.001L6.143 5.64A2 2 0 0 1 5 6H3a1 1 0 0 0-1 1v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2z"/><path d="M6 12h.01"/><path d="M6 16h.01"/><circle cx="12" cy="10" r="2"/>`)

// Unlink2 is [unlink-2](https://lucide.dev/icons/unlink-2).
var Unlink2 = icon.Lucide(viewBox, paint, `<path d="M15 7h2a5 5 0 0 1 0 10h-2m-6 0H7A5 5 0 0 1 7 7h2"/>`)

// Unlink is [unlink](https://lucide.dev/icons/unlink).
var Unlink = icon.Lucide(viewBox, paint, `<path d="m18.84 12.25 1.72-1.71h-.02a5.004 5.004 0 0 0-.12-7.07 5.006 5.006 0 0 0-6.95 0l-1.72 1.71"/><path d="m5.17 11.75-1.71 1.71a5.004 5.004 0 0 0 .12 7.07 5.006 5.006 0 0 0 6.95 0l1.71-1.71"/><line x1="8" x2="8" y1="2" y2="5"/><line x1="2" x2="5" y1="8" y2="8"/><line x1="16" x2="16" y1="19" y2="22"/><line x1="19" x2="22" y1="16" y2="16"/>`)

// Unplug is [unplug](https://lucide.dev/icons/unplug).
var Unplug = icon.Lucide(viewBox, paint, `<path d="m19 5 3-3"/><path d="m2 22 3-3"/><path d="M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6-2.3 2.3a2.4 2.4 0 0 0 0 3.4Z"/><path d="M7.5 13.5 10 11"/><path d="M10.5 16.5 13 14"/><path d="m12 6 6 6 2.3-2.3a2.4 2.4 0 0 0 0-3.4l-2.6-2.6a2.4 2.4 0 0 0-3.4 0Z"/>`)

// Upload is [upload](https://lucide.dev/icons/upload).
var Upload = icon.Lucide(viewBox, paint, `<path d="M12 3v12"/><path d="m17 8-5-5-5 5"/><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>`)

// UsbCPort is [usb-c-port](https://lucide.dev/icons/usb-c-port).
var UsbCPort = icon.Lucide(viewBox, paint, `<path d="M6 12h12"/><rect x="2" y="8" width="20" height="8" rx="4"/>`)

// Usb is [usb](https://lucide.dev/icons/usb).
var Usb = icon.Lucide(viewBox, paint, `<circle cx="10" cy="7" r="1"/><circle cx="4" cy="20" r="1"/><path d="M4.7 19.3 19 5"/><path d="m21 3-3 1 2 2Z"/><path d="M9.26 7.68 5 12l2 5"/><path d="m10 14 5 2 3.5-3.5"/><path d="m18 12 1-1 1 1-1 1Z"/>`)

// UserCheck is [user-check](https://lucide.dev/icons/user-check).
var UserCheck = icon.Lucide(viewBox, paint, `<path d="m16 11 2 2 4-4"/><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/>`)

// UserCog is [user-cog](https://lucide.dev/icons/user-cog).
var UserCog = icon.Lucide(viewBox, paint, `<path d="M10 15H6a4 4 0 0 0-4 4v2"/><path d="m14.305 16.53.923-.382"/><path d="m15.228 13.852-.923-.383"/><path d="m16.852 12.228-.383-.923"/><path d="m16.852 17.772-.383.924"/><path d="m19.148 12.228.383-.923"/><path d="m19.53 18.696-.382-.924"/><path d="m20.772 13.852.924-.383"/><path d="m20.772 16.148.924.383"/><circle cx="18" cy="15" r="3"/><circle cx="9" cy="7" r="4"/>`)

// UserGroup is [user-group](https://lucide.dev/icons/user-group).
var UserGroup = icon.Lucide(viewBox, paint, `<path d="M17 21v-1a2 2 0 00-2-2H9a2 2 0 00-2 2v1"/><path d="M19 10h1a2 2 0 012 2v1"/><path d="M5 10H4a2 2 0 00-2 2v1"/><circle cx="12" cy="11" r="3"/><circle cx="18" cy="4" r="2"/><circle cx="6" cy="4" r="2"/>`)

// UserKey is [user-key](https://lucide.dev/icons/user-key).
var UserKey = icon.Lucide(viewBox, paint, `<path d="M20 11v6"/><path d="M20 13h2"/><path d="M3 21v-2a4 4 0 0 1 4-4h6a4 4 0 0 1 2.072.578"/><circle cx="10" cy="7" r="4"/><circle cx="20" cy="19" r="2"/>`)

// UserLock is [user-lock](https://lucide.dev/icons/user-lock).
var UserLock = icon.Lucide(viewBox, paint, `<path d="M19 16v-2a2 2 0 0 0-4 0v2"/><path d="M9.5 15H7a4 4 0 0 0-4 4v2"/><circle cx="10" cy="7" r="4"/><rect x="13" y="16" width="8" height="5" rx=".899"/>`)

// UserMinus is [user-minus](https://lucide.dev/icons/user-minus).
var UserMinus = icon.Lucide(viewBox, paint, `<path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><line x1="22" x2="16" y1="11" y2="11"/>`)

// UserPen is [user-pen](https://lucide.dev/icons/user-pen).
var UserPen = icon.Lucide(viewBox, paint, `<path d="M11.5 15H7a4 4 0 0 0-4 4v2"/><path d="M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/><circle cx="10" cy="7" r="4"/>`)

// UserPlus is [user-plus](https://lucide.dev/icons/user-plus).
var UserPlus = icon.Lucide(viewBox, paint, `<path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><line x1="19" x2="19" y1="8" y2="14"/><line x1="22" x2="16" y1="11" y2="11"/>`)

// UserRoundArrowLeft is [user-round-arrow-left](https://lucide.dev/icons/user-round-arrow-left).
var UserRoundArrowLeft = icon.Lucide(viewBox, paint, `<path d="m19 16-3 3"/><path d="M2 21a8 8 0 0 1 12.664-6.5"/><path d="M22 19h-6l3 3"/><circle cx="10" cy="8" r="5"/>`)

// UserRoundCheck is [user-round-check](https://lucide.dev/icons/user-round-check).
var UserRoundCheck = icon.Lucide(viewBox, paint, `<path d="M2 21a8 8 0 0 1 13.292-6"/><circle cx="10" cy="8" r="5"/><path d="m16 19 2 2 4-4"/>`)

// UserRoundCog is [user-round-cog](https://lucide.dev/icons/user-round-cog).
var UserRoundCog = icon.Lucide(viewBox, paint, `<path d="m14.305 19.53.923-.382"/><path d="m15.228 16.852-.923-.383"/><path d="m16.852 15.228-.383-.923"/><path d="m16.852 20.772-.383.924"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.696-.382-.924"/><path d="M2 21a8 8 0 0 1 10.434-7.62"/><path d="m20.772 16.852.924-.383"/><path d="m20.772 19.148.924.383"/><circle cx="10" cy="8" r="5"/><circle cx="18" cy="18" r="3"/>`)

// UserRoundGroup is [user-round-group](https://lucide.dev/icons/user-round-group).
var UserRoundGroup = icon.Lucide(viewBox, paint, `<path d="M17 21a5 5 0 00-10 0"/><path d="M22 10.5a3.5 3.5 0 00-5.507-2.868"/><path d="M7.507 7.632A3.5 3.5 0 002 10.5"/><circle cx="12" cy="13" r="3"/><circle cx="18.5" cy="4.5" r="2.5"/><circle cx="5.5" cy="4.5" r="2.5"/>`)

// UserRoundKey is [user-round-key](https://lucide.dev/icons/user-round-key).
var UserRoundKey = icon.Lucide(viewBox, paint, `<path d="M19 11v6"/><path d="M19 13h2"/><path d="M2 21a8 8 0 0 1 12.868-6.349"/><circle cx="10" cy="8" r="5"/><circle cx="19" cy="19" r="2"/>`)

// UserRoundMinus is [user-round-minus](https://lucide.dev/icons/user-round-minus).
var UserRoundMinus = icon.Lucide(viewBox, paint, `<path d="M2 21a8 8 0 0 1 13.292-6"/><circle cx="10" cy="8" r="5"/><path d="M22 19h-6"/>`)

// UserRoundPen is [user-round-pen](https://lucide.dev/icons/user-round-pen).
var UserRoundPen = icon.Lucide(viewBox, paint, `<path d="M2 21a8 8 0 0 1 10.821-7.487"/><path d="M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/><circle cx="10" cy="8" r="5"/>`)

// UserRoundPlus is [user-round-plus](https://lucide.dev/icons/user-round-plus).
var UserRoundPlus = icon.Lucide(viewBox, paint, `<path d="M2 21a8 8 0 0 1 13.292-6"/><circle cx="10" cy="8" r="5"/><path d="M19 16v6"/><path d="M22 19h-6"/>`)

// UserRoundSearch is [user-round-search](https://lucide.dev/icons/user-round-search).
var UserRoundSearch = icon.Lucide(viewBox, paint, `<circle cx="10" cy="8" r="5"/><path d="M2 21a8 8 0 0 1 10.434-7.62"/><circle cx="18" cy="18" r="3"/><path d="m22 22-1.9-1.9"/>`)

// UserRoundStar is [user-round-star](https://lucide.dev/icons/user-round-star).
var UserRoundStar = icon.Lucide(viewBox, paint, `<path d="M16.051 12.616a1 1 0 0 1 1.909.024l.737 1.452a1 1 0 0 0 .737.535l1.634.256a1 1 0 0 1 .588 1.806l-1.172 1.168a1 1 0 0 0-.282.866l.259 1.613a1 1 0 0 1-1.541 1.134l-1.465-.75a1 1 0 0 0-.912 0l-1.465.75a1 1 0 0 1-1.539-1.133l.258-1.613a1 1 0 0 0-.282-.867l-1.156-1.152a1 1 0 0 1 .572-1.822l1.633-.256a1 1 0 0 0 .737-.535z"/><path d="M2 21a8 8 0 0 1 6.778-7.905"/><path d="M8.927 12.885a5 5 0 1 1 6.061-4.535"/>`)

// UserRoundX is [user-round-x](https://lucide.dev/icons/user-round-x).
var UserRoundX = icon.Lucide(viewBox, paint, `<path d="m16.5 16.5 5 5"/><path d="M2 21a8 8 0 0 1 11.531-7.18"/><path d="m21.5 16.5-5 5"/><circle cx="10" cy="8" r="5"/>`)

// UserRound is [user-round](https://lucide.dev/icons/user-round).
var UserRound = icon.Lucide(viewBox, paint, `<circle cx="12" cy="8" r="5"/><path d="M20 21a8 8 0 0 0-16 0"/>`)

// UserSearch is [user-search](https://lucide.dev/icons/user-search).
var UserSearch = icon.Lucide(viewBox, paint, `<circle cx="10" cy="7" r="4"/><path d="M10.3 15H7a4 4 0 0 0-4 4v2"/><circle cx="17" cy="17" r="3"/><path d="m21 21-1.9-1.9"/>`)

// UserShield is [user-shield](https://lucide.dev/icons/user-shield).
var UserShield = icon.Lucide(viewBox, paint, `<path d="M10 15H6a4 4 0 0 0-4 4v2"/><path d="M22 17.5c0 2.499-1.75 3.749-3.83 4.474a.5.5 0 0 1-.335-.005c-2.085-.72-3.835-1.97-3.835-4.47V14a.5.5 0 0 1 .5-.499c1 0 2.25-.6 3.12-1.36a.6.6 0 0 1 .76-.001c.875.765 2.12 1.36 3.12 1.36a.5.5 0 0 1 .5.5z"/><circle cx="9" cy="7" r="4"/>`)

// UserStar is [user-star](https://lucide.dev/icons/user-star).
var UserStar = icon.Lucide(viewBox, paint, `<path d="M16.051 12.616a1 1 0 0 1 1.909.024l.737 1.452a1 1 0 0 0 .737.535l1.634.256a1 1 0 0 1 .588 1.806l-1.172 1.168a1 1 0 0 0-.282.866l.259 1.613a1 1 0 0 1-1.541 1.134l-1.465-.75a1 1 0 0 0-.912 0l-1.465.75a1 1 0 0 1-1.539-1.133l.258-1.613a1 1 0 0 0-.282-.866l-1.156-1.153a1 1 0 0 1 .572-1.822l1.633-.256a1 1 0 0 0 .737-.535z"/><path d="M8 15H7a4 4 0 0 0-4 4v2"/><circle cx="10" cy="7" r="4"/>`)

// UserX is [user-x](https://lucide.dev/icons/user-x).
var UserX = icon.Lucide(viewBox, paint, `<path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><line x1="17" x2="22" y1="8" y2="13"/><line x1="22" x2="17" y1="8" y2="13"/>`)

// User is [user](https://lucide.dev/icons/user).
var User = icon.Lucide(viewBox, paint, `<path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/>`)

// UsersCollaboration is [users-collaboration](https://lucide.dev/icons/users-collaboration).
var UsersCollaboration = icon.Lucide(viewBox, paint, `<path d="M12.925 3.047a9 9 0 017.616 6.11"/><path d="M13 22a2 2 0 012-2h4a2 2 0 012 2"/><path d="M3 12a2 2 0 012-2h4a2 2 0 012 2"/><path d="M3.889 15.9a9 9 0 005.247 4.632"/><circle cx="17" cy="14" r="2"/><circle cx="7" cy="4" r="2"/>`)

// UsersRoundCollaboration is [users-round-collaboration](https://lucide.dev/icons/users-round-collaboration).
var UsersRoundCollaboration = icon.Lucide(viewBox, paint, `<path d="M20.453 8.911a9 9 0 00-6.696-5.738"/><path d="M21 22a4 4 0 00-8 0"/><path d="M3 12a4 4 0 017.963-.543"/><path d="M3.889 15.9a9 9 0 005.247 4.632"/><circle cx="17" cy="15" r="3"/><circle cx="7" cy="5" r="3"/>`)

// UsersRound is [users-round](https://lucide.dev/icons/users-round).
var UsersRound = icon.Lucide(viewBox, paint, `<path d="M18 21a8 8 0 0 0-16 0"/><circle cx="10" cy="8" r="5"/><path d="M22 20c0-3.37-2-6.5-4-8a5 5 0 0 0-.45-8.3"/>`)

// Users is [users](https://lucide.dev/icons/users).
var Users = icon.Lucide(viewBox, paint, `<path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><path d="M16 3.128a4 4 0 0 1 0 7.744"/><path d="M22 21v-2a4 4 0 0 0-3-3.87"/><circle cx="9" cy="7" r="4"/>`)

// UtensilsCrossed is [utensils-crossed](https://lucide.dev/icons/utensils-crossed).
var UtensilsCrossed = icon.Lucide(viewBox, paint, `<path d="m16 2-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8"/><path d="M15 15 3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0L15 15Zm0 0 7 7"/><path d="m2.1 21.8 6.4-6.3"/><path d="m19 5-7 7"/>`)

// Utensils is [utensils](https://lucide.dev/icons/utensils).
var Utensils = icon.Lucide(viewBox, paint, `<path d="M3 2v7c0 1.1.9 2 2 2h4a2 2 0 0 0 2-2V2"/><path d="M7 2v20"/><path d="M21 15V2a5 5 0 0 0-5 5v6c0 1.1.9 2 2 2h3Zm0 0v7"/>`)

// UtilityPole is [utility-pole](https://lucide.dev/icons/utility-pole).
var UtilityPole = icon.Lucide(viewBox, paint, `<path d="M12 2v20"/><path d="M2 5h20"/><path d="M3 3v2"/><path d="M7 3v2"/><path d="M17 3v2"/><path d="M21 3v2"/><path d="m19 5-7 7-7-7"/>`)

// Van is [van](https://lucide.dev/icons/van).
var Van = icon.Lucide(viewBox, paint, `<path d="M13 6v5a1 1 0 0 0 1 1h6.102a1 1 0 0 1 .712.298l.898.91a1 1 0 0 1 .288.702V17a1 1 0 0 1-1 1h-3"/><path d="M5 18H3a1 1 0 0 1-1-1V8a2 2 0 0 1 2-2h12c1.1 0 2.1.8 2.4 1.8l1.176 4.2"/><path d="M9 18h5"/><circle cx="16" cy="18" r="2"/><circle cx="7" cy="18" r="2"/>`)

// Variable is [variable](https://lucide.dev/icons/variable).
var Variable = icon.Lucide(viewBox, paint, `<path d="M8 21s-4-3-4-9 4-9 4-9"/><path d="M16 3s4 3 4 9-4 9-4 9"/><line x1="15" x2="9" y1="9" y2="15"/><line x1="9" x2="15" y1="9" y2="15"/>`)

// Vault is [vault](https://lucide.dev/icons/vault).
var Vault = icon.Lucide(viewBox, paint, `<rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="7.5" cy="7.5" r=".5" fill="currentColor"/><path d="m7.9 7.9 2.7 2.7"/><circle cx="16.5" cy="7.5" r=".5" fill="currentColor"/><path d="m13.4 10.6 2.7-2.7"/><circle cx="7.5" cy="16.5" r=".5" fill="currentColor"/><path d="m7.9 16.1 2.7-2.7"/><circle cx="16.5" cy="16.5" r=".5" fill="currentColor"/><path d="m13.4 13.4 2.7 2.7"/><circle cx="12" cy="12" r="2"/>`)

// VectorPolygon is [vector-polygon](https://lucide.dev/icons/vector-polygon).
var VectorPolygon = icon.Lucide(viewBox, paint, `<path d="m12.828 4.813 5.344 2.375"/><path d="m15.769 18.153 3.461-8.306"/><path d="m5.687 14.074 7.625 4.852"/><path d="M9.772 5.579 5.228 11.42"/><circle cx="11" cy="4" r="2"/><circle cx="15" cy="20" r="2"/><circle cx="20" cy="8" r="2"/><circle cx="4" cy="13" r="2"/>`)

// VectorSquare is [vector-square](https://lucide.dev/icons/vector-square).
var VectorSquare = icon.Lucide(viewBox, paint, `<path d="M17.055 4.533a24 24 0 00-10.11 0"/><path d="M19.467 17.055a24 24 0 000-10.11"/><path d="M4.533 6.945a24 24 0 000 10.11"/><path d="M6.945 19.467a24 24 0 0010.11 0"/><circle cx="19" cy="19" r="2"/><circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/><circle cx="5" cy="5" r="2"/>`)

// Vegan is [vegan](https://lucide.dev/icons/vegan).
var Vegan = icon.Lucide(viewBox, paint, `<path d="M16 8q6 0 6-6-6 0-6 6"/><path d="M17.41 3.59a10 10 0 1 0 3 3"/><path d="M2 2a26.6 26.6 0 0 1 10 20c.9-6.82 1.5-9.5 4-14"/>`)

// VenetianMask is [venetian-mask](https://lucide.dev/icons/venetian-mask).
var VenetianMask = icon.Lucide(viewBox, paint, `<path d="M18 11c-1.5 0-2.5.5-3 2"/><path d="M4 6a2 2 0 0 0-2 2v4a5 5 0 0 0 5 5 8 8 0 0 1 5 2 8 8 0 0 1 5-2 5 5 0 0 0 5-5V8a2 2 0 0 0-2-2h-3a8 8 0 0 0-5 2 8 8 0 0 0-5-2z"/><path d="M6 11c1.5 0 2.5.5 3 2"/>`)

// VenusAndMars is [venus-and-mars](https://lucide.dev/icons/venus-and-mars).
var VenusAndMars = icon.Lucide(viewBox, paint, `<path d="M10 20h4"/><path d="M12 16v6"/><path d="M17 2h4v4"/><path d="m21 2-5.46 5.46"/><circle cx="12" cy="11" r="5"/>`)

// Venus is [venus](https://lucide.dev/icons/venus).
var Venus = icon.Lucide(viewBox, paint, `<path d="M12 15v7"/><path d="M9 19h6"/><circle cx="12" cy="9" r="6"/>`)

// VibrateOff is [vibrate-off](https://lucide.dev/icons/vibrate-off).
var VibrateOff = icon.Lucide(viewBox, paint, `<path d="m2 8 2 2-2 2 2 2-2 2"/><path d="m22 8-2 2 2 2-2 2 2 2"/><path d="M8 8v10c0 .55.45 1 1 1h6c.55 0 1-.45 1-1v-2"/><path d="M16 10.34V6c0-.55-.45-1-1-1h-4.34"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Vibrate is [vibrate](https://lucide.dev/icons/vibrate).
var Vibrate = icon.Lucide(viewBox, paint, `<path d="m2 8 2 2-2 2 2 2-2 2"/><path d="m22 8-2 2 2 2-2 2 2 2"/><rect width="8" height="14" x="8" y="5" rx="1"/>`)

// VideoOff is [video-off](https://lucide.dev/icons/video-off).
var VideoOff = icon.Lucide(viewBox, paint, `<path d="M10.66 6H14a2 2 0 0 1 2 2v2.5l5.248-3.062A.5.5 0 0 1 22 7.87v8.196"/><path d="M16 16a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2"/><path d="m2 2 20 20"/>`)

// Video is [video](https://lucide.dev/icons/video).
var Video = icon.Lucide(viewBox, paint, `<path d="m16 13 5.223 3.482a.5.5 0 0 0 .777-.416V7.87a.5.5 0 0 0-.752-.432L16 10.5"/><rect x="2" y="6" width="14" height="12" rx="2"/>`)

// Videotape is [videotape](https://lucide.dev/icons/videotape).
var Videotape = icon.Lucide(viewBox, paint, `<rect width="20" height="16" x="2" y="4" rx="2"/><path d="M2 8h20"/><circle cx="8" cy="14" r="2"/><path d="M8 12h8"/><circle cx="16" cy="14" r="2"/>`)

// View is [view](https://lucide.dev/icons/view).
var View = icon.Lucide(viewBox, paint, `<path d="M21 17v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2"/><path d="M21 7V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2"/><circle cx="12" cy="12" r="1"/><path d="M18.944 12.33a1 1 0 0 0 0-.66 7.5 7.5 0 0 0-13.888 0 1 1 0 0 0 0 .66 7.5 7.5 0 0 0 13.888 0"/>`)

// VirusOff is [virus-off](https://lucide.dev/icons/virus-off).
var VirusOff = icon.Lucide(viewBox, paint, `<path d="M10.01 10h.01"/><path d="M12 14.991h.01"/><path d="M12 22v-3"/><path d="M12 2v3"/><path d="M13 22h-2"/><path d="M13 2h-2"/><path d="m16.5 19.794-1-1.733"/><path d="m16.5 4.205-1 1.732"/><path d="M18.891 13.235a7 7 0 00-8.126-8.126"/><path d="m19.794 7.5-1.732 1"/><path d="M2 12h3"/><path d="M2 13v-2"/><path d="m2 2 20 20"/><path d="M22 12h-3"/><path d="M22 13v-2"/><path d="m4.206 16.5 1.732-1"/><path d="m4.206 7.5 1.732 1"/><path d="M7.05 7.05a7 7 0 009.9 9.9"/><path d="m7.5 19.794 1-1.733"/><path d="M9 12h.01"/>`)

// Virus is [virus](https://lucide.dev/icons/virus).
var Virus = icon.Lucide(viewBox, paint, `<path d="M12 14.991h.01"/><path d="M12 22v-3"/><path d="M12 2v3"/><path d="M13 22h-2"/><path d="M13 2h-2"/><path d="M13.99 10H14"/><path d="m16.5 19.794-1-1.733"/><path d="m16.5 4.205-1 1.732"/><path d="m19.794 16.5-1.732-1"/><path d="m19.794 7.5-1.732 1"/><path d="M2 12h3"/><path d="M2 13v-2"/><path d="M22 12h-3"/><path d="M22 13v-2"/><path d="m4.206 16.5 1.732-1"/><path d="m4.206 7.5 1.732 1"/><path d="m7.5 19.794 1-1.733"/><path d="m7.5 4.205 1 1.732"/><path d="M9 12h.01"/><circle cx="12" cy="12" r="7"/>`)

// Voicemail is [voicemail](https://lucide.dev/icons/voicemail).
var Voicemail = icon.Lucide(viewBox, paint, `<circle cx="6" cy="12" r="4"/><circle cx="18" cy="12" r="4"/><line x1="6" x2="18" y1="16" y2="16"/>`)

// Volleyball is [volleyball](https://lucide.dev/icons/volleyball).
var Volleyball = icon.Lucide(viewBox, paint, `<path d="M11 7a16 16 20 0 1 10.98 4.362"/><path d="M12 12a13 13 0 0 1-8.66 5"/><path d="M16.83 13.634a16 16 0 0 1-9.267 7.328"/><path d="M20.66 17A13 13 0 0 0 12 12a13 13 0 0 1 0-10"/><path d="M8.17 15.366a16 16 0 0 1-1.713-11.69"/><circle cx="12" cy="12" r="10"/>`)

// Volume1 is [volume-1](https://lucide.dev/icons/volume-1).
var Volume1 = icon.Lucide(viewBox, paint, `<path d="M11 4.702a.705.705 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298z"/><path d="M16 9a5 5 0 0 1 0 6"/>`)

// Volume2 is [volume-2](https://lucide.dev/icons/volume-2).
var Volume2 = icon.Lucide(viewBox, paint, `<path d="M11 4.702a.705.705 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298z"/><path d="M16 9a5 5 0 0 1 0 6"/><path d="M19.364 18.364a9 9 0 0 0 0-12.728"/>`)

// VolumeOff is [volume-off](https://lucide.dev/icons/volume-off).
var VolumeOff = icon.Lucide(viewBox, paint, `<path d="M16 9a5 5 0 0 1 .95 2.293"/><path d="M19.364 5.636a9 9 0 0 1 1.889 9.96"/><path d="m2 2 20 20"/><path d="m7 7-.587.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298V11"/><path d="M9.828 4.172A.686.686 0 0 1 11 4.657v.686"/>`)

// VolumeX is [volume-x](https://lucide.dev/icons/volume-x).
var VolumeX = icon.Lucide(viewBox, paint, `<path d="M11 4.702a.7.7 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.7.7 0 0 0 11 19.298z"/><path d="m16.5 14.5 5-5"/><path d="m16.5 9.5 5 5"/>`)

// Volume is [volume](https://lucide.dev/icons/volume).
var Volume = icon.Lucide(viewBox, paint, `<path d="M11 4.702a.705.705 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298z"/>`)

// Vote is [vote](https://lucide.dev/icons/vote).
var Vote = icon.Lucide(viewBox, paint, `<path d="m9 12 2 2 4-4"/><path d="M5 7c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2v12H5V7Z"/><path d="M22 19H2"/>`)

// WalletCards is [wallet-cards](https://lucide.dev/icons/wallet-cards).
var WalletCards = icon.Lucide(viewBox, paint, `<path d="M3 11h3.75a2 2 0 0 1 1.6.8l.45.6a4 4 0 0 0 6.4 0l.45-.6a2 2 0 0 1 1.6-.8H21"/><path d="M3 7h18"/><rect x="3" y="3" width="18" height="18" rx="2"/>`)

// WalletMinimal is [wallet-minimal](https://lucide.dev/icons/wallet-minimal).
var WalletMinimal = icon.Lucide(viewBox, paint, `<path d="M17 14h.01"/><path d="M7 7h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14"/>`)

// Wallet is [wallet](https://lucide.dev/icons/wallet).
var Wallet = icon.Lucide(viewBox, paint, `<path d="M19 7V4a1 1 0 0 0-1-1H5a2 2 0 0 0 0 4h15a1 1 0 0 1 1 1v4h-3a2 2 0 0 0 0 4h3a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1"/><path d="M3 5v14a2 2 0 0 0 2 2h15a1 1 0 0 0 1-1v-4"/>`)

// Wallpaper is [wallpaper](https://lucide.dev/icons/wallpaper).
var Wallpaper = icon.Lucide(viewBox, paint, `<path d="M12 17v4"/><path d="M8 21h8"/><path d="m9 17 6.1-6.1a2 2 0 0 1 2.81.01L22 15"/><circle cx="8" cy="9" r="2"/><rect x="2" y="3" width="20" height="14" rx="2"/>`)

// WandSparkles is [wand-sparkles](https://lucide.dev/icons/wand-sparkles).
var WandSparkles = icon.Lucide(viewBox, paint, `<path d="m21.64 3.64-1.28-1.28a1.21 1.21 0 0 0-1.72 0L2.36 18.64a1.21 1.21 0 0 0 0 1.72l1.28 1.28a1.2 1.2 0 0 0 1.72 0L21.64 5.36a1.2 1.2 0 0 0 0-1.72"/><path d="m14 7 3 3"/><path d="M5 6v4"/><path d="M19 14v4"/><path d="M10 2v2"/><path d="M7 8H3"/><path d="M21 16h-4"/><path d="M11 3H9"/>`)

// Wand is [wand](https://lucide.dev/icons/wand).
var Wand = icon.Lucide(viewBox, paint, `<path d="M15 4V2"/><path d="M15 16v-2"/><path d="M8 9h2"/><path d="M20 9h2"/><path d="M17.8 11.8 19 13"/><path d="M15 9h.01"/><path d="M17.8 6.2 19 5"/><path d="m3 21 9-9"/><path d="M12.2 6.2 11 5"/>`)

// Warehouse is [warehouse](https://lucide.dev/icons/warehouse).
var Warehouse = icon.Lucide(viewBox, paint, `<path d="M18 21V10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v11"/><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 1.132-1.803l7.95-3.974a2 2 0 0 1 1.837 0l7.948 3.974A2 2 0 0 1 22 8z"/><path d="M6 13h12"/><path d="M6 17h12"/>`)

// WashingMachine is [washing-machine](https://lucide.dev/icons/washing-machine).
var WashingMachine = icon.Lucide(viewBox, paint, `<path d="M3 6h3"/><path d="M17 6h.01"/><rect width="18" height="20" x="3" y="2" rx="2"/><circle cx="12" cy="13" r="5"/><path d="M12 18a2.5 2.5 0 0 0 0-5 2.5 2.5 0 0 1 0-5"/>`)

// Watch is [watch](https://lucide.dev/icons/watch).
var Watch = icon.Lucide(viewBox, paint, `<path d="M12 10v2.2l1.6 1"/><path d="m16.13 7.66-.81-4.05a2 2 0 0 0-2-1.61h-2.68a2 2 0 0 0-2 1.61l-.78 4.05"/><path d="m7.88 16.36.8 4a2 2 0 0 0 2 1.61h2.72a2 2 0 0 0 2-1.61l.81-4.05"/><circle cx="12" cy="12" r="6"/>`)

// WavesArrowDown is [waves-arrow-down](https://lucide.dev/icons/waves-arrow-down).
var WavesArrowDown = icon.Lucide(viewBox, paint, `<path d="M12 10L12 2"/><path d="M16 6L12 10L8 6"/><path d="M2 15C2.6 15.5 3.2 16 4.5 16C7 16 7 14 9.5 14C12.1 14 11.9 16 14.5 16C17 16 17 14 19.5 14C20.8 14 21.4 14.5 22 15"/><path d="M2 21C2.6 21.5 3.2 22 4.5 22C7 22 7 20 9.5 20C12.1 20 11.9 22 14.5 22C17 22 17 20 19.5 20C20.8 20 21.4 20.5 22 21"/>`)

// WavesArrowUp is [waves-arrow-up](https://lucide.dev/icons/waves-arrow-up).
var WavesArrowUp = icon.Lucide(viewBox, paint, `<path d="M12 2v8"/><path d="M2 15c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"/><path d="M2 21c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"/><path d="m8 6 4-4 4 4"/>`)

// WavesHorizontal is [waves-horizontal](https://lucide.dev/icons/waves-horizontal).
var WavesHorizontal = icon.Lucide(viewBox, paint, `<path d="M2 12q2.5 2 5 0t5 0 5 0 5 0"/><path d="M2 19q2.5 2 5 0t5 0 5 0 5 0"/><path d="M2 5q2.5 2 5 0t5 0 5 0 5 0"/>`)

// WavesLadder is [waves-ladder](https://lucide.dev/icons/waves-ladder).
var WavesLadder = icon.Lucide(viewBox, paint, `<path d="M19 5a2 2 0 0 0-2 2v11"/><path d="M2 18c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1"/><path d="M7 13h10"/><path d="M7 9h10"/><path d="M9 5a2 2 0 0 0-2 2v11"/>`)

// WavesVertical is [waves-vertical](https://lucide.dev/icons/waves-vertical).
var WavesVertical = icon.Lucide(viewBox, paint, `<path d="M12 2q2 2.5 0 5t0 5 0 5 0 5"/><path d="M19 2q2 2.5 0 5t0 5 0 5 0 5"/><path d="M5 2q2 2.5 0 5t0 5 0 5 0 5"/>`)

// Waypoints is [waypoints](https://lucide.dev/icons/waypoints).
var Waypoints = icon.Lucide(viewBox, paint, `<path d="m10.586 5.414-5.172 5.172"/><path d="m18.586 13.414-5.172 5.172"/><path d="M6 12h12"/><circle cx="12" cy="20" r="2"/><circle cx="12" cy="4" r="2"/><circle cx="20" cy="12" r="2"/><circle cx="4" cy="12" r="2"/>`)

// WebcamOff is [webcam-off](https://lucide.dev/icons/webcam-off).
var WebcamOff = icon.Lucide(viewBox, paint, `<path d="M12 22v-4"/><path d="M12.754 7.096a3 3 0 0 1 2.15 2.15"/><path d="M12.863 12.873a3 3 0 0 1-3.736-3.735"/><path d="M16.566 16.57A8 8 0 0 1 5.43 5.433"/><path d="m2 2 20 20"/><path d="M7 22h10"/><path d="M8.478 2.817a8 8 0 0 1 10.705 10.705"/>`)

// Webcam is [webcam](https://lucide.dev/icons/webcam).
var Webcam = icon.Lucide(viewBox, paint, `<circle cx="12" cy="10" r="8"/><circle cx="12" cy="10" r="3"/><path d="M7 22h10"/><path d="M12 22v-4"/>`)

// WebhookOff is [webhook-off](https://lucide.dev/icons/webhook-off).
var WebhookOff = icon.Lucide(viewBox, paint, `<path d="M17 17h-5c-1.09-.02-1.94.92-2.5 1.9A3 3 0 1 1 2.57 15"/><path d="M9 3.4a4 4 0 0 1 6.52.66"/><path d="m6 17 3.1-5.8a2.5 2.5 0 0 0 .057-2.05"/><path d="M20.3 20.3a4 4 0 0 1-2.3.7"/><path d="M18.6 13a4 4 0 0 1 3.357 3.414"/><path d="m12 6 .6 1"/><path d="m2 2 20 20"/>`)

// Webhook is [webhook](https://lucide.dev/icons/webhook).
var Webhook = icon.Lucide(viewBox, paint, `<path d="M18 16.98h-5.99c-1.1 0-1.95.94-2.48 1.9A4 4 0 0 1 2 17c.01-.7.2-1.4.57-2"/><path d="m6 17 3.13-5.78c.53-.97.1-2.18-.5-3.1a4 4 0 1 1 6.89-4.06"/><path d="m12 6 3.13 5.73C15.66 12.7 16.9 13 18 13a4 4 0 0 1 0 8"/>`)

// WeightTilde is [weight-tilde](https://lucide.dev/icons/weight-tilde).
var WeightTilde = icon.Lucide(viewBox, paint, `<path d="M6.5 8a2 2 0 0 0-1.906 1.46L2.1 18.5A2 2 0 0 0 4 21h16a2 2 0 0 0 1.925-2.54L19.4 9.5A2 2 0 0 0 17.48 8z"/><path d="M7.999 15a2.5 2.5 0 0 1 4 0 2.5 2.5 0 0 0 4 0"/><circle cx="12" cy="5" r="3"/>`)

// Weight is [weight](https://lucide.dev/icons/weight).
var Weight = icon.Lucide(viewBox, paint, `<circle cx="12" cy="5" r="3"/><path d="M6.5 8a2 2 0 0 0-1.905 1.46L2.1 18.5A2 2 0 0 0 4 21h16a2 2 0 0 0 1.925-2.54L19.4 9.5A2 2 0 0 0 17.48 8Z"/>`)

// WheatOff is [wheat-off](https://lucide.dev/icons/wheat-off).
var WheatOff = icon.Lucide(viewBox, paint, `<path d="m2 22 10-10"/><path d="m16 8-1.17 1.17"/><path d="M3.47 12.53 5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"/><path d="m8 8-.53.53a3.5 3.5 0 0 0 0 4.94L9 15l1.53-1.53c.55-.55.88-1.25.98-1.97"/><path d="M10.91 5.26c.15-.26.34-.51.56-.73L13 3l1.53 1.53a3.5 3.5 0 0 1 .28 4.62"/><path d="M20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"/><path d="M11.47 17.47 13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"/><path d="m16 16-.53.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.49 3.49 0 0 1 1.97-.98"/><path d="M18.74 13.09c.26-.15.51-.34.73-.56L21 11l-1.53-1.53a3.5 3.5 0 0 0-4.62-.28"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Wheat is [wheat](https://lucide.dev/icons/wheat).
var Wheat = icon.Lucide(viewBox, paint, `<path d="M2 22 16 8"/><path d="M3.47 12.53 5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"/><path d="M7.47 8.53 9 7l1.53 1.53a3.5 3.5 0 0 1 0 4.94L9 15l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"/><path d="M11.47 4.53 13 3l1.53 1.53a3.5 3.5 0 0 1 0 4.94L13 11l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Z"/><path d="M20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"/><path d="M11.47 17.47 13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"/><path d="M15.47 13.47 17 15l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"/><path d="M19.47 9.47 21 11l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L13 11l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"/>`)

// Whistle is [whistle](https://lucide.dev/icons/whistle).
var Whistle = icon.Lucide(viewBox, paint, `<path d="M10 6v4"/><path d="M21 6a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-5.675A7 7 0 1 1 9 6z"/>`)

// WholeWord is [whole-word](https://lucide.dev/icons/whole-word).
var WholeWord = icon.Lucide(viewBox, paint, `<circle cx="7" cy="12" r="3"/><path d="M10 9v6"/><circle cx="17" cy="12" r="3"/><path d="M14 7v8"/><path d="M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"/>`)

// WifiCog is [wifi-cog](https://lucide.dev/icons/wifi-cog).
var WifiCog = icon.Lucide(viewBox, paint, `<path d="m14.305 19.53.923-.382"/><path d="m15.228 16.852-.923-.383"/><path d="m16.852 15.228-.383-.923"/><path d="m16.852 20.772-.383.924"/><path d="m19.148 15.228.383-.923"/><path d="m19.53 21.696-.382-.924"/><path d="M2 8.82a15 15 0 0120 0"/><path d="m20.772 16.852.924-.383"/><path d="m20.772 19.148.924.383"/><path d="M5 12.859a10 10 0 0110.18-2.342"/><path d="M8.5 16.429a5 5 0 011.998-1.2"/><circle cx="18" cy="18" r="3"/>`)

// WifiHigh is [wifi-high](https://lucide.dev/icons/wifi-high).
var WifiHigh = icon.Lucide(viewBox, paint, `<path d="M12 20h.01"/><path d="M5 12.859a10 10 0 0 1 14 0"/><path d="M8.5 16.429a5 5 0 0 1 7 0"/>`)

// WifiLock is [wifi-lock](https://lucide.dev/icons/wifi-lock).
var WifiLock = icon.Lucide(viewBox, paint, `<path d="M2 8.82a15 15 0 0120 0"/><path d="M20 16v-2a2 2 0 00-4 0v2"/><path d="M5 12.859a10 10 0 018.436-2.756"/><path d="M8.5 16.429a5 5 0 011.794-1.13"/><rect x="14" y="16" width="8" height="5" rx="1"/>`)

// WifiLow is [wifi-low](https://lucide.dev/icons/wifi-low).
var WifiLow = icon.Lucide(viewBox, paint, `<path d="M12 20h.01"/><path d="M8.5 16.429a5 5 0 0 1 7 0"/>`)

// WifiOff is [wifi-off](https://lucide.dev/icons/wifi-off).
var WifiOff = icon.Lucide(viewBox, paint, `<path d="M12 20h.01"/><path d="M8.5 16.429a5 5 0 0 1 7 0"/><path d="M5 12.859a10 10 0 0 1 5.17-2.69"/><path d="M19 12.859a10 10 0 0 0-2.007-1.523"/><path d="M2 8.82a15 15 0 0 1 4.177-2.643"/><path d="M22 8.82a15 15 0 0 0-11.288-3.764"/><path d="m2 2 20 20"/>`)

// WifiPen is [wifi-pen](https://lucide.dev/icons/wifi-pen).
var WifiPen = icon.Lucide(viewBox, paint, `<path d="M2 8.82a15 15 0 0 1 20 0"/><path d="M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/><path d="M5 12.859a10 10 0 0 1 10.5-2.222"/><path d="M8.5 16.429a5 5 0 0 1 3-1.406"/>`)

// WifiSync is [wifi-sync](https://lucide.dev/icons/wifi-sync).
var WifiSync = icon.Lucide(viewBox, paint, `<path d="M11.965 10.105v4L13.5 12.5a5 5 0 0 1 8 1.5"/><path d="M11.965 14.105h4"/><path d="M17.965 18.105h4L20.43 19.71a5 5 0 0 1-8-1.5"/><path d="M2 8.82a15 15 0 0 1 20 0"/><path d="M21.965 22.105v-4"/><path d="M5 12.86a10 10 0 0 1 3-2.032"/><path d="M8.5 16.429h.01"/>`)

// WifiZero is [wifi-zero](https://lucide.dev/icons/wifi-zero).
var WifiZero = icon.Lucide(viewBox, paint, `<path d="M12 20h.01"/>`)

// Wifi is [wifi](https://lucide.dev/icons/wifi).
var Wifi = icon.Lucide(viewBox, paint, `<path d="M12 20h.01"/><path d="M2 8.82a15 15 0 0 1 20 0"/><path d="M5 12.859a10 10 0 0 1 14 0"/><path d="M8.5 16.429a5 5 0 0 1 7 0"/>`)

// WindArrowDown is [wind-arrow-down](https://lucide.dev/icons/wind-arrow-down).
var WindArrowDown = icon.Lucide(viewBox, paint, `<path d="M10 2v8"/><path d="M12.8 21.6A2 2 0 1 0 14 18H2"/><path d="M17.5 10a2.5 2.5 0 1 1 2 4H2"/><path d="m6 6 4 4 4-4"/>`)

// WindArrowUp is [wind-arrow-up](https://lucide.dev/icons/wind-arrow-up).
var WindArrowUp = icon.Lucide(viewBox, paint, `<path d="M10 2v8"/><path d="M12.8 21.6A2 2 0 1 0 14 18H2"/><path d="M17.5 10a2.5 2.5 0 1 1 2 4H2"/><path d="m6 6 4 -4 4 4"/>`)

// Wind is [wind](https://lucide.dev/icons/wind).
var Wind = icon.Lucide(viewBox, paint, `<path d="M12.8 19.6A2 2 0 1 0 14 16H2"/><path d="M17.5 8a2.5 2.5 0 1 1 2 4H2"/><path d="M9.8 4.4A2 2 0 1 1 11 8H2"/>`)

// WineOff is [wine-off](https://lucide.dev/icons/wine-off).
var WineOff = icon.Lucide(viewBox, paint, `<path d="M8 22h8"/><path d="M7 10h3m7 0h-1.343"/><path d="M12 15v7"/><path d="M7.307 7.307A12.33 12.33 0 0 0 7 10a5 5 0 0 0 7.391 4.391M8.638 2.981C8.75 2.668 8.872 2.34 9 2h6c1.5 4 2 6 2 8 0 .407-.05.809-.145 1.198"/><line x1="2" x2="22" y1="2" y2="22"/>`)

// Wine is [wine](https://lucide.dev/icons/wine).
var Wine = icon.Lucide(viewBox, paint, `<path d="M8 22h8"/><path d="M7 10h10"/><path d="M12 15v7"/><path d="M12 15a5 5 0 0 0 5-5c0-2-.5-4-2-8H9c-1.5 4-2 6-2 8a5 5 0 0 0 5 5Z"/>`)

// Workflow is [workflow](https://lucide.dev/icons/workflow).
var Workflow = icon.Lucide(viewBox, paint, `<rect width="8" height="8" x="3" y="3" rx="2"/><path d="M7 11v4a2 2 0 0 0 2 2h4"/><rect width="8" height="8" x="13" y="13" rx="2"/>`)

// Worm is [worm](https://lucide.dev/icons/worm).
var Worm = icon.Lucide(viewBox, paint, `<path d="m19 12-1.5 3"/><path d="M19.63 18.81 22 20"/><path d="M6.47 8.23a1.68 1.68 0 0 1 2.44 1.93l-.64 2.08a6.76 6.76 0 0 0 10.16 7.67l.42-.27a1 1 0 1 0-2.73-4.21l-.42.27a1.76 1.76 0 0 1-2.63-1.99l.64-2.08A6.66 6.66 0 0 0 3.94 3.9l-.7.4a1 1 0 1 0 2.55 4.34z"/>`)

// WrenchOff is [wrench-off](https://lucide.dev/icons/wrench-off).
var WrenchOff = icon.Lucide(viewBox, paint, `<path d="M10.747 5.093a6 6 0 0 1 6.841-2.882c.438.12.54.662.219.984L14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.106-3.105c.32-.322.863-.22.983.218a6 6 0 0 1-2.882 6.842"/><path d="m13.5 13.5-7.88 7.88a1 1 0 0 1-2.999-3l7.88-7.88"/><path d="m2 2 20 20"/>`)

// Wrench is [wrench](https://lucide.dev/icons/wrench).
var Wrench = icon.Lucide(viewBox, paint, `<path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.106-3.105c.32-.322.863-.22.983.218a6 6 0 0 1-8.259 7.057l-7.91 7.91a1 1 0 0 1-2.999-3l7.91-7.91a6 6 0 0 1 7.057-8.259c.438.12.54.662.219.984z"/>`)

// XLineTop is [x-line-top](https://lucide.dev/icons/x-line-top).
var XLineTop = icon.Lucide(viewBox, paint, `<path d="M18 4H6"/><path d="M18 8 6 20"/><path d="m6 8 12 12"/>`)

// X is [x](https://lucide.dev/icons/x).
var X = icon.Lucide(viewBox, paint, `<path d="M18 6 6 18"/><path d="m6 6 12 12"/>`)

// ZapOff is [zap-off](https://lucide.dev/icons/zap-off).
var ZapOff = icon.Lucide(viewBox, paint, `<path d="M10.768 5.111 13.44 2.44a1.5 1.5 0 012.474 1.561l-1.633 4.625"/><path d="m18.889 13.232.672-.672A1.5 1.5 0 0018.5 10h-2.844"/><path d="m2 2 20 20"/><path d="m7.94 7.94-3.5 3.499A1.5 1.5 0 005.5 14h4.002a.5.5 0 01.471.666L8.086 20a1.5 1.5 0 002.475 1.56l5.5-5.5"/>`)

// Zap is [zap](https://lucide.dev/icons/zap).
var Zap = icon.Lucide(viewBox, paint, `<path d="M15.914 4a1.5 1.5 0 00-2.474-1.561l-9 9A1.5 1.5 0 005.5 14h4.002a.5.5 0 01.471.666L8.086 20a1.5 1.5 0 002.475 1.56l9-9A1.5 1.5 0 0018.5 10h-3.997a.5.5 0 01-.472-.667z"/>`)

// ZodiacAquarius is [zodiac-aquarius](https://lucide.dev/icons/zodiac-aquarius).
var ZodiacAquarius = icon.Lucide(viewBox, paint, `<path d="m2 10 2.456-3.684a.7.7 0 0 1 1.106-.013l2.39 3.413a.7.7 0 0 0 1.096-.001l2.402-3.432a.7.7 0 0 1 1.098 0l2.402 3.432a.7.7 0 0 0 1.098 0l2.389-3.413a.7.7 0 0 1 1.106.013L22 10"/><path d="m2 18.002 2.456-3.684a.7.7 0 0 1 1.106-.013l2.39 3.413a.7.7 0 0 0 1.097 0l2.402-3.432a.7.7 0 0 1 1.098 0l2.402 3.432a.7.7 0 0 0 1.098 0l2.389-3.413a.7.7 0 0 1 1.106.013L22 18.002"/>`)

// ZodiacAries is [zodiac-aries](https://lucide.dev/icons/zodiac-aries).
var ZodiacAries = icon.Lucide(viewBox, paint, `<path d="M12 7.5a4.5 4.5 0 1 1 5 4.5"/><path d="M7 12a4.5 4.5 0 1 1 5-4.5V21"/>`)

// ZodiacCancer is [zodiac-cancer](https://lucide.dev/icons/zodiac-cancer).
var ZodiacCancer = icon.Lucide(viewBox, paint, `<path d="M21 14.5A9 6.5 0 0 1 5.5 19"/><path d="M3 9.5A9 6.5 0 0 1 18.5 5"/><circle cx="17.5" cy="14.5" r="3.5"/><circle cx="6.5" cy="9.5" r="3.5"/>`)

// ZodiacCapricorn is [zodiac-capricorn](https://lucide.dev/icons/zodiac-capricorn).
var ZodiacCapricorn = icon.Lucide(viewBox, paint, `<path d="M11 21a3 3 0 0 0 3-3V6.5a1 1 0 0 0-7 0"/><path d="M7 19V6a3 3 0 0 0-3-3h0"/><circle cx="17" cy="17" r="3"/>`)

// ZodiacGemini is [zodiac-gemini](https://lucide.dev/icons/zodiac-gemini).
var ZodiacGemini = icon.Lucide(viewBox, paint, `<path d="M16 4.525v14.948"/><path d="M20 3A17 17 0 0 1 4 3"/><path d="M4 21a17 17 0 0 1 16 0"/><path d="M8 4.525v14.948"/>`)

// ZodiacLeo is [zodiac-leo](https://lucide.dev/icons/zodiac-leo).
var ZodiacLeo = icon.Lucide(viewBox, paint, `<path d="M10 16c0-4-3-4.5-3-8a5 5 0 0 1 10 0c0 3.466-3 6.196-3 10a3 3 0 0 0 6 0"/><circle cx="7" cy="16" r="3"/>`)

// ZodiacLibra is [zodiac-libra](https://lucide.dev/icons/zodiac-libra).
var ZodiacLibra = icon.Lucide(viewBox, paint, `<path d="M3 16h6.857c.162-.012.19-.323.038-.38a6 6 0 1 1 4.212 0c-.153.057-.125.368.038.38H21"/><path d="M3 20h18"/>`)

// ZodiacOphiuchus is [zodiac-ophiuchus](https://lucide.dev/icons/zodiac-ophiuchus).
var ZodiacOphiuchus = icon.Lucide(viewBox, paint, `<path d="M3 10A6.06 6.06 0 0 1 12 10 A6.06 6.06 0 0 0 21 10"/><path d="M6 3v12a6 6 0 0 0 12 0V3"/>`)

// ZodiacPisces is [zodiac-pisces](https://lucide.dev/icons/zodiac-pisces).
var ZodiacPisces = icon.Lucide(viewBox, paint, `<path d="M19 21a15 15 0 0 1 0-18"/><path d="M20 12H4"/><path d="M5 3a15 15 0 0 1 0 18"/>`)

// ZodiacSagittarius is [zodiac-sagittarius](https://lucide.dev/icons/zodiac-sagittarius).
var ZodiacSagittarius = icon.Lucide(viewBox, paint, `<path d="M15 3h6v6"/><path d="M21 3 3 21"/><path d="m9 9 6 6"/>`)

// ZodiacScorpio is [zodiac-scorpio](https://lucide.dev/icons/zodiac-scorpio).
var ZodiacScorpio = icon.Lucide(viewBox, paint, `<path d="M10 19V5.5a1 1 0 0 1 5 0V17a2 2 0 0 0 2 2h5l-3-3"/><path d="m22 19-3 3"/><path d="M5 19V5.5a1 1 0 0 1 5 0"/><path d="M5 5.5A2.5 2.5 0 0 0 2.5 3"/>`)

// ZodiacTaurus is [zodiac-taurus](https://lucide.dev/icons/zodiac-taurus).
var ZodiacTaurus = icon.Lucide(viewBox, paint, `<circle cx="12" cy="15" r="6"/><path d="M18 3A6 6 0 0 1 6 3"/>`)

// ZodiacVirgo is [zodiac-virgo](https://lucide.dev/icons/zodiac-virgo).
var ZodiacVirgo = icon.Lucide(viewBox, paint, `<path d="M11 5.5a1 1 0 0 1 5 0V16a5 5 0 0 0 5 5"/><path d="M16 11.5a1 1 0 0 1 5 0V16a5 5 0 0 1-5 5"/><path d="M6 19V6a3 3 0 0 0-3-3h0"/><path d="M6 5.5a1 1 0 0 1 5 0V19"/>`)

// ZoomIn is [zoom-in](https://lucide.dev/icons/zoom-in).
var ZoomIn = icon.Lucide(viewBox, paint, `<circle cx="11" cy="11" r="8"/><line x1="21" x2="16.65" y1="21" y2="16.65"/><line x1="11" x2="11" y1="8" y2="14"/><line x1="8" x2="14" y1="11" y2="11"/>`)

// ZoomOut is [zoom-out](https://lucide.dev/icons/zoom-out).
var ZoomOut = icon.Lucide(viewBox, paint, `<circle cx="11" cy="11" r="8"/><line x1="21" x2="16.65" y1="21" y2="16.65"/><line x1="8" x2="14" y1="11" y2="11"/>`)
