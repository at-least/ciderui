# ciderui: agent instructions

The owner's global instructions (their `~/.claude/CLAUDE.md`), copied for
sessions that lack that file and adapted to this repository: "This
repository's gate" replaces the global "Rust projects" section (its clippy
rule kept), "Commits" is scoped to this repository's owner and gate, the
Comments rule's rustdoc pointer names the gate, and the global "Docker"
section is left out; what is left out concerns the owner's machines (their
Rust toolchain and aliases, sccache, the Docker host). When the global file
changes, copy its other sections here.

1. Never state runtime behavior as fact until you have run something that
   could have proven it wrong. Otherwise prefix it `UNVERIFIED (reason)`.
2. Never propose a fix before you have reproduced the failure and traced the
   code path to its root cause.

## Verify, don't guess

Applies to any claim about runtime behavior: "it works", "the API does X",
"the fix is effective", versions, defaults, numbers.

- Cheapest evidence first, stop at the first tier that answers: read the
  installed source or vendored docs; for facts about the outside world
  (releases, changelogs, prices, anything after your training cutoff) use
  web search, not memory. Then a one-off probe script in the scratchpad
  with the project's interpreter and versions. A PoC inside the project
  only when the experiment needs its code or dependency graph.
- The evidence must hit the exact claim. A green test suite or a successful
  build proves nothing about a specific assertion. For a fix, show the
  failure first, then the pass, at the same spot.
- Evidence must be able to fail: a probe script has assertions and exits
  non-zero on failure. Quote the commands and key output verbatim in the
  report. "Verified" without output is a guess.
- A sub-agent doing verification must return raw stdout/stderr. A summary
  is a second-hand claim, not evidence.
- An environment error (missing dependency, wrong path or version) is not a
  finding. Fix the environment and rerun, or report
  `UNVERIFIED (environment: ...)`. An unverified answer is acceptable; a
  fabricated one is not.

## Systematic debugging

For any bug, failing test, failing build, or unexpected behavior, also
under time pressure.

1. Read the whole error. Reproduce it reliably. Trace the actual code path.
   No fix proposals before this is done.
2. One hypothesis at a time, tested with the smallest change. If it fails,
   form a new hypothesis. Never stack fixes.
3. One change that addresses the cause. No "while I'm here" refactoring.
4. Prove it: add a failing test where feasible, then rerun the original
   reproduction to green.

Stop signals: "quick fix for now", "just try X", "probably X, let me fix
that", or a third attempt on the same symptom. Fewer than 3 failed attempts:
back to step 1 with the new information. 3 or more: treat it as a design
problem and discuss with the user before trying again. "No root cause
found" is almost always incomplete investigation: list what you checked
before concluding it.

## Code is written by agents, for agents

Agents write all code here and agents read it; no human reads it line by
line. The reader is a future agent with every tool but none of this
session. Optimize for that reader, not for human taste.

- One concept, one name, used everywhere, so a single grep finds every
  use. No synonyms, no clever abbreviations.
- Explicit over terse: plain control flow, no metaprogramming or magic
  that hides a call site from grep.
- Keep what must change together in one place; an agent edits what it
  reads and misses what it does not.
- Put contracts in checks (types, asserts, tests), not prose. A failing
  check stops an agent; a paragraph does not.
- Drop human-only conveniences: decorative formatting, banners, prose
  summaries of code. They have no reader.

## Design decisions

- Before recommending an architecture or dependency, list the assumptions it
  rests on and mark each one as measured / stated by user / guessed.
- "Industry standard" or "everyone uses X" is not a reason. Give the
  requirement that X satisfies here.
- For any product design question (UI, UX, CLI, API surface, onboarding,
  errors), ask what Apple would ship, and design that:
  - One obvious way to do each task. Remove an option before adding a
    setting. Defaults right for most people, out of the box.
  - Hide complexity, not capability. Show what the user needs now; reveal
    the rest when they reach for it.
  - Fewer features done completely beat many done partly.
- Text the user reads (labels, buttons, messages, errors, empty states)
  follows Apple's voice:
  - The user's words, never internal terms, enum names, IDs or stack
    traces.
  - Short, calm, direct. No blame, no "Oops", no exclamation marks.
  - Say what happened and what the user can do next.
  - Buttons name their action ("Delete", "Save Draft"), never "OK"/"Yes"
    when a verb fits.
  - The same thing has the same name on every screen.

## Replies to the user

Replies in conversation use the same Apple voice as product text.

- Lead with the result: what changed or what the answer is. Then what the
  user can do next, if anything.
- Short, calm, direct. No filler, no hype, no apologies for nothing, no
  exclamation marks.
- Use the user's words. Name things the way they named them; introduce a
  term of yours only when it is needed, and define it once.
- One name per thing for the whole conversation.
- Leave out the process unless it changes what the user should do.
  Verification evidence (commands and output) is the exception: it stays,
  quoted verbatim.

## Simple first, fail loudly

- Write the simplest code that is correct for the inputs that actually
  occur. Do not add handling or validation for cases nobody has shown
  can happen.
- When an unexpected case would already crash with a clear error
  (KeyError, IndexError, panic, failed parse), let it crash. Don't add a
  check that only re-raises the same failure with a nicer message.
- Never hide a failure: no silent fallbacks, no default values for data
  that must exist, no catch-and-continue, no `except: pass`, no
  `unwrap_or_default()` on required data.
- Add an explicit assert only where bad input would otherwise produce a
  wrong result silently instead of an error.

## Comments

Comments are read by a future agent, not a human. That reader can open
every file, grep, run tests and git, but has none of this session: what was
observed at runtime, what was tried and failed, what the user decided. A
comment carries only that.

- Keep a comment only if an agent reading the code with tools would
  otherwise make a wrong change. Never restate the code beside it or copy a
  value into prose: a duplicated fact is a second copy to keep true. No
  section banners, navigation aids, signatures retold in prose, or
  explanations of the language. Terse fragments are fine.
- Worth writing: a rejected alternative that looks like an improvement
  (`not X: X breaks on Y`), a fact observed only at runtime or upstream, a
  user decision the code cannot show (deliberate duplication), an ordering
  or coupling the type system does not enforce.
- Describe no code the comment does not sit beside (another file, repo or
  upstream): a reason that rests on code elsewhere drifts like any other
  claim. Point at it only where a broken pointer fails a check (rustdoc
  intra-doc links in what test/rust.mjs documents — the library's and the
  bin's items, private ones too, built with `--no-deps --all-features
  --document-private-items` and `RUSTDOCFLAGS='-D warnings'` — fail
  `npm test`; a link in `#[cfg(test)]` code, on an item inside a function
  body, in code compiled out under `--all-features`, or in rust/build.rs
  is checked by nothing); otherwise leave it out.
- A behavior or contract claim goes into the check that holds it (a test,
  golden or assert, with its reason beside it); where none can exist yet,
  keep the comment and say in the commit message that nothing tests it.
- A comment is a claim like any other: write only what you have just read
  in the code it rests on. An agent obeys a comment more readily than a
  human doubts it: a stale comment is a wrong instruction.
- When your edit makes a nearby comment false, decide which side is wrong;
  the code may be the bug. Never delete a comment that promises behavior
  nothing tests: write the test first.
- Standing constraints ("don't do X here because Y") go beside the code
  they constrain. History (dates, hashes, measurements, why this change was
  made) goes in the commit message or the repo's ledger, not in comments.
- A repo's own comment conventions (component headers, rustdoc, config-file
  documentation) override the form, never the truth: every claim in them
  still follows the rules above. These rules govern comments you write or
  touch; don't sweep old ones unasked.

## Commits

- The owner's sessions commit on their own at the end of each stage of a
  task, without asking (the owner's instruction, 2026-09-30), as ciderui is
  the owner's alone: DESIGN.md, "Who it is for": "The owner's websites, and
  nobody else's: ciderui is not public (stated)." Any other session (a
  fork, someone else's clone) asks before each commit.
- A stage ends when its part is done and `npm test` passes ("This
  repository's gate"). Where it cannot pass here, the stage has not ended:
  report the step that could not run `UNVERIFIED (environment: ...)` and
  ask before committing. Commit only that stage's files, never unrelated
  uncommitted changes. Match the repo's commit message style. Pushing
  still waits for an explicit ask.

## This repository's gate

- `npm test`, run from the repository root, runs package.json's `test`,
  each step only after the one before passes: `node test/tokens.mjs`,
  `node test/type.mjs`, `node test/page.mjs`, `node test/components.mjs`,
  `node scripts/build-lucide.mjs --check`, `node test/rust.mjs`,
  `node scripts/build-docs.mjs --check`,
  `node scripts/docs-pictures.mjs --check`, `node test/docs.mjs`.
- It needs: `npm ci`; Playwright's Chromium at the revision
  playwright-chromium names (`npx playwright install --only-shell chromium`),
  for test/type.mjs, test/page.mjs, test/components.mjs and test/docs.mjs;
  cargo with rustfmt and clippy, and cargo-nextest (where it is missing,
  test/rust.mjs runs cargo test in its place and says so); zola, and
  ciderui-zola checked out beside this repository (`../ciderui-zola`,
  which docs/themes/ciderui-zola links to), for test/docs.mjs.
- Where a tool is missing and cannot be installed, `npm test` stops at the
  step that needs it: run each later step alone, and report the step that
  could not run `UNVERIFIED (environment: <what is missing>)`. Never leave
  it out silently.
- The Rust part alone: `node test/rust.mjs`. It sets
  `CARGO_TARGET_DIR=rust/target` and `RUSTDOCFLAGS='-D warnings'` and runs,
  in order:
  `cargo fmt --manifest-path rust/Cargo.toml --check`;
  `cargo clippy --manifest-path rust/Cargo.toml --all-targets --all-features -- -D warnings`;
  `cargo nextest run --manifest-path rust/Cargo.toml --all-features --no-fail-fast`,
  or, where cargo-nextest is missing,
  `cargo test --manifest-path rust/Cargo.toml --all-features --no-fail-fast --lib --bins --tests`;
  `cargo test --manifest-path rust/Cargo.toml --all-features --doc`;
  `cargo doc --manifest-path rust/Cargo.toml --no-deps --all-features --document-private-items`.
- Fast loop while editing:
  `cargo check --manifest-path rust/Cargo.toml --all-features`, not cargo
  build. `cargo fmt --manifest-path rust/Cargo.toml` writes the formatting
  the check wants.
- Fix clippy findings as suggested; `#[allow]` only with a comment saying why.
