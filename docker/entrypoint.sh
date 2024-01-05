#!/bin/sh

echo "Running a $UNCHAINED_NODE_TYPE node."

if [ $UNCHAINED_NODE_TYPE = "full" ]; then
  unchained postgres migrate conf.yaml
  retVal=$?
  if [ $retVal -ne 0 ]; then
    exit $retVal
  fi
fi

unchained start conf.yaml --generate
