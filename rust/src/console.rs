//! The console set: App Store Connect's newer look's parts, the
//! `cui-console-*` classes (`components/console/`).

mod console_badge;
mod console_banner;
mod console_card;
mod console_checkbox;
mod console_chip;
mod console_dialog;
mod console_empty;
mod console_field;
mod console_hint;
mod console_menu;
mod console_pager;
mod console_sidebar;
mod console_status;
mod console_tabs;

pub use console_badge::{ConsoleBadge, ConsoleBadgeVariant, ConsoleColor};
pub use console_banner::ConsoleBanner;
pub use console_card::ConsoleMetric;
pub use console_checkbox::ConsoleCheckboxBox;
pub use console_chip::ConsoleChip;
pub use console_dialog::ConsoleDialog;
pub use console_empty::{ConsoleEmpty, ConsoleEmptySize};
pub use console_field::{
    ConsoleError, ConsoleFormField, ConsoleHelp, ConsoleSelect, console_error_id, console_help_id,
};
pub use console_hint::ConsoleHint;
pub use console_menu::{ConsoleMenu, ConsoleMenuGroup, ConsoleMenuItem};
pub use console_pager::ConsolePager;
pub use console_sidebar::{ConsoleSidebar, ConsoleSidebarItem};
pub use console_status::ConsoleStatus;
pub use console_tabs::{ConsoleTab, ConsoleTabs, ConsoleTabsVariant};
