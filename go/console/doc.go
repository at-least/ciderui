// Package console is ciderui's console set (components/console/), App Store
// Connect's newer look's parts, as gomponents nodes: the markup the
// cui-console-* classes expect. Only the components whose markup is more
// than a class on one element are here: a console button is a class,
// written on the element with the attributes the page gives it
// (<a class="cui-console-btn" data-size="sm">).
//
// Each component is a struct whose every field a call writes: the module's
// exhaustruct check (go tool exhaustruct) fails a call that leaves one out,
// as a site's own gate should, so a field added here stops every call that
// does not set it.
//
// A type's name is its class without cui-, in CamelCase (cui-console-badge,
// [ConsoleBadge]), the Console prefix kept although the package names it:
// one rule across every set, so one grep finds the class, the type and a
// site's uses (decided for the Go port).
//
// A choice among a few values (a [ConsoleColor], a [ConsoleBadgeVariant])
// is a struct type with one unexported field and its values as package
// variables, so a call cannot write a value of its own; its zero value is
// the choice that writes no attribute, or, for a choice with no such
// default ([ConsoleTabsVariant]), panics when rendered.
package console
