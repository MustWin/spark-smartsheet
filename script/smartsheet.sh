#!/bin/bash

AUTH="${AUTH:-Authorization: Bearer ${SHEET_API_TOKEN:-invalidtoken}}"
URL="${URL:-https://api.smartsheet.com/2.0/${FRAGMENT:-home}}"

. "$(dirname "$0")/.curl.sh"
