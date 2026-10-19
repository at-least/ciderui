// The checks' bookkeeping: each check passes or fails with what it compared,
// and a test file ends with its tally, exiting non-zero on any failure.

const failures = [];
let passed = 0;

// a computed value against the expected one, lengths within 1/32px
export function check(name, got, want, { tolerance = 1 / 32 } = {}) {
  const ok = same(got, want, tolerance);
  if (ok) passed++;
  else failures.push(`${name}: ${got} (want ${want})`);
  return ok;
}

const same = (a, b, tol) => {
  if (a === b) return true;
  const na = String(a).match(/^(-?[\d.]+)px$/), nb = String(b).match(/^(-?[\d.]+)px$/);
  return Boolean(na && nb && Math.abs(Number(na[1]) - Number(nb[1])) <= tol);
};

export function done(label) {
  for (const f of failures) console.log(`FAIL ${f}`);
  console.log(`${label}: ${passed} passed, ${failures.length} failed`);
  process.exit(failures.length ? 1 : 0);
}
