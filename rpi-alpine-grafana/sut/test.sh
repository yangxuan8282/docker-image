#!/bin/bash

GRAFANA_HOST=${GRAFANA_HOST:-grafana}
GRAFANA_USER=admin
GRAFANA_PASS=AxwayPassword*

echo -n "test grafana presence...   "
i=0
r=1
while [[ $r -ne 0 ]]; do
  nslookup grafana >/dev/null 2>&1
  r=$?
  if [[ $i -gt 40 ]]; then break; fi
  ((i++))
  sleep 1
done
if [[ $r -ne 0 ]]; then
  echo
  echo "failed"
  nslookup grafana
  exit 1
fi
printf "%-16s[OK]\n" "resolves(${i}s)"

echo -n "test grafana api...        "
i=0
r=0
while [[ $r -lt 1 ]]; do
  ((i++))
  sleep 1
  r=$(curl -u $GRAFANA_USER:$GRAFANA_PASS $GRAFANA_HOST:3000/api/admin/stats 2>/dev/null | jq -r '.user_count')
  if [[ $i -gt 40 ]]; then break; fi
done
if [[ $r -lt 1 ]]; then
  echo
  echo "failed"
  curl -u $GRAFANA_USER:$GRAFANA_PASS $GRAFANA_HOST:3000/api/admin/stats
  ci=$(docker ps | grep /influxdb | awk '{print $1}')
  echo "logs from influxdb $ci:"
  docker logs $ci
  cg=$(docker ps | grep /grafana | awk '{print $1}')
  echo "logs from grafana $cg:"
  docker logs $cg
  exit 1
fi
printf "%-16s[OK]\n" " (${i}s)"

echo -n "test grafana auth...       "
org=$(curl -u $GRAFANA_USER:$GRAFANA_PASS $GRAFANA_HOST:3000/api/org 2>/dev/null | jq -r '.name')
r=$?
if [[ $r -ne 0 || -z "$org" || "x$org" = "xnull" ]]; then
  echo
  echo "auth failed"
  curl -u $GRAFANA_USER:$GRAFANA_PASS $GRAFANA_HOST:3000/api/org 2>/dev/null
  exit 1
fi
printf "%-16s[OK]\n" "($org)"

echo -n "test grafana datasource... "
ds=$(curl -u $GRAFANA_USER:$GRAFANA_PASS $GRAFANA_HOST:3000/api/datasources/name/telegraf 2>/dev/null | jq -r '.name')
r=$?
if [[ $r -ne 0 || "x$ds" != "xtelegraf" ]]; then
  echo
  echo "failed to found datasource"
  exit 1
fi
printf "%-16s[OK]\n" "($ds)"

echo -n "test grafana dashboards... "
n=$(curl -u $GRAFANA_USER:$GRAFANA_PASS $GRAFANA_HOST:3000/api/search?query=AMP%20%Swarm%20Health 2>/dev/null | jq -r 'map(select(.type=="dash-db"))| length')
r=$?
if [[ $r -ne 0 || -z "$n" || "x$n" = "xnull" || $n -lt 1 ]]; then
  echo
  echo "failed to found dashboards"
  exit 1
fi
printf "%-16s[OK]\n" "($n)"
echo "all tests passed successfully"
