#!/bin/sh
# sh front.sh > results/front-safari.log &, on the Mac, killed as the run
# ends (README.md): every second until killed, the time (ms since 1970, the
# clock of run.mjs's claims), the front application, and the seconds since
# the last input (HIDIdleTime): what a Safari run is beside
while :; do
  t=$(perl -MTime::HiRes=time -e "printf q(%d), time*1000")
  app=$(lsappinfo info -only name "$(lsappinfo front)" | sed -n "1s/^\"\([^\"]*\)\".*/\1/p")
  idle=$(ioreg -c IOHIDSystem | awk "/HIDIdleTime/ {print int(\$NF/1000000000); exit}")
  echo "$t	$app	$idle"
  sleep 1
done
