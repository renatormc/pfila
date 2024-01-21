#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
"$SCRIPT_DIR/pfila_runner" --port  "$1" --proc-id "$2" > "$3" 2>&1 &