#!/bin/bash

AUTH="${AUTH:-Authorization: Bearer ${LOCAL_API_TOKEN:-invalidtoken}}"
URL="${URL:-http://localhost:8000/v1/${FRAGMENT:-}}"

. "$(dirname "$0")/.curl.sh"
