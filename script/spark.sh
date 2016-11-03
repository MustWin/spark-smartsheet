#!/bin/bash

ROOM="${ROOM:-Y2lzY29zcGFyazovL3VzL1JPT00vM2FlMmNhNjAtYTIwMC0xMWU2LWFmYzEtOTFlNWMyOWM4NjY1}"

AUTH="${AUTH:-Authorization: Bearer ${SPARK_API_TOKEN:-invalidtoken}}"
URL="${URL:-https://api.ciscospark.com/v1/${FRAGMENT:-messages}?roomId=${ROOM}}"

. "$(dirname "$0")/.curl.sh"
