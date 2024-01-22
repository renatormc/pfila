#!/bin/bash
total_args=$#
outputFile="${!#}"
"${@:1:$((total_args-1))}" > "$outputFile" 2>&1 && echo "PFila: Finish!" >> "$outputFile" &