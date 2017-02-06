#!/bin/bash

TMPFILE=$(mktemp)
TESTDIR=$(dirname $0)
program=envtpl
cmd=$(dirname $0)/../$program

TAG_REGION=eu-west-1 TAG_DATACENTER=dc1 INTERVAL=2s OUTPUT_INFLUXDB_ENABLED=true HOSTNAME=localhost ${cmd} ${TESTDIR}/test.tpl > ${TMPFILE}
diff "${TESTDIR}/test.txt" "${TMPFILE}" >/dev/null
if [[ $? -ne 0 ]]; then
	echo "$program does not produce expected result"
	echo "expected result:"
	cat "${TESTDIR}/test.txt"
	echo "observed result:"
	cat "${TMPFILE}"
	rm "${TMPFILE}"
	return 1
fi
echo "Tests passed successfully"
rm "${TMPFILE}"
