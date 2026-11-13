// Package console renders components/console/ (App Store Connect's newer look)
// as gomponents nodes. A part with no type here is classes on the page's own
// elements, as its examples write them: <pre class="cui-console-pre"
// tabindex="0">, <div class="cui-console-card">, <dialog
// class="cui-console-dialog"> (opened by showModal() or a button's
// command="show-modal": an open attribute shows it with no backdrop, and
// Escape does not close it).
package console
