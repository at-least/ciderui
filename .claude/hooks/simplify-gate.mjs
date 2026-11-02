// One /simplify unlocks one commit: the commit consumes the flag.
import { execFileSync } from 'node:child_process';
import { existsSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { resolve } from 'node:path';

const input = JSON.parse(readFileSync(0, 'utf8'));
const git = (...args) => execFileSync('git', args, { cwd: input.cwd, encoding: 'utf8' }).trim();
const flag = () => resolve(input.cwd, git('rev-parse', '--git-path', 'simplify-done'));

// a /simplify typed by the user expands without a Skill call
if (input.hook_event_name === 'UserPromptSubmit') {
  if (/^\/simplify\b/.test(input.prompt)) writeFileSync(flag(), '');
  process.exit(0);
}
if (input.tool_name === 'Skill') {
  if (input.tool_input.skill === 'simplify') writeFileSync(flag(), '');
  process.exit(0);
}

if (!/\bgit\s+commit\b/.test(input.tool_input.command)) process.exit(0);
const changed = (git('diff', 'HEAD', '--name-only') + '\n' + git('ls-files', '--others', '--exclude-standard'))
  .split('\n')
  .filter(f => /\.(css|go|mjs|js|html|sh|swift)$/.test(f));
if (changed.length === 0) process.exit(0);
if (existsSync(flag())) {
  rmSync(flag());
  process.exit(0);
}
console.log(JSON.stringify({
  hookSpecificOutput: {
    hookEventName: 'PreToolUse',
    permissionDecision: 'deny',
    permissionDecisionReason: `Code changed since the last commit (${changed.join(', ')}). Run /simplify on it first (CLAUDE.md, Commits).`,
  },
}));
