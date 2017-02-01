#!/bin/bash

echo -n "Check system release...   "
sv=$(cat /etc/alpine-release)
echo "$sv" | egrep -q "^[0-9]\.[0-9]\.[0-9]$"
if [[ $? -ne 0 ]]; then
  echo "failed"
  echo "$sv"
  exit 1
fi
printf "%-20s[OK]\n" "($sv)"

echo -n "Test jq version...        "
jqv=$(jq --version)
if [[ $? -ne 0 || -z "$jqv" ]]; then
  echo
  echo "failed"
  exit 1
fi
printf "%-20s[OK]\n" "($jqv)"

echo -n "Test envtpl...            "
envtpl -h >/dev/null
if [[ $? -ne 0 ]]; then
  echo
  echo "failed"
  exit 1
fi
printf "%-20s[OK]\n" "ready"

echo -n "Test gosu                 "
test -x /usr/bin/gosu
if [[ $? -ne 0 ]]; then
  echo
  echo "failed"
  exit 1
fi
printf "%-20s[OK]\n" "ready"

echo "All tests passed successfully"
exit 0
