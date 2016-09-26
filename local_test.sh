#!/usr/bin/env bash
# vim: filetype=sh:

export WERCKER_VERSIONING_POM=pom_example.xml
export WERCKER_VERSIONING_OUTFILE=version_output
export WERCKER_MAIN_PIPELINE_STARTED=123456789

echo -e "Running build"
go build main.go

if [ "$?" -ne 0 ]; then
  echo "Failed to build main"
  exit 1
fi

./main 

cat version_output
