const failures = [];
let passed = 0;

export function check(name, got, want, { tolerance = 1 / 32 } = {}) {
  return hold(same(got, want, tolerance), `${name}: ${got} (want ${want})`);
}

export function hold(ok, failure) {
  if (ok) passed++;
  else failures.push(failure);
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
