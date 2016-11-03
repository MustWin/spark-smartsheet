#!/bin/bash

HEADERS="-H '${AUTH:-Authorization: Bearer ${AUTH_TOKEN:-invalidtoken}}' -H 'Content-Type: ${CONTENT_TYPE:-application/json}'"
OPTIONS="-X${METHOD:-GET}"

VERBOSE=
DEBUG=

while getopts vd:h? opt; do
  case $opt in
    v)
      [ -z ${DEBUG} ] || DEBUG="--trace /dev/stderr -D /dev/stderr"
      [ -z ${DEBUG} ] && DEBUG="-D /dev/stderr"
      VERBOSE=1
      ;;
    d)
      OPTIONS="-X${METHOD:-POST} -d@${OPTARG}"
      ;;
    *)
      echo
      echo "Usage: $0 [-v] [-d <file>]"
      echo
      echo "  -v for verbose logging of request/response"
      echo "  -d to send a data file in the request"
      echo
      exit
      ;;
  esac
done

CMD="curl ${DEBUG} ${HEADERS} ${OPTIONS} '${URL:-http://localhost/}'"

[ -z ${VERBOSE} ] || echo -e "Executing:\n${CMD}\n"

eval ${CMD} | jq .
