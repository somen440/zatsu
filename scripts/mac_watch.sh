#!/bin/bash

cmd=$1
sleep=$2

if [ "${sleep}" == "" ]; then
  sleep=2
fi

while :; do
  clear
  date
  $cmd
  sleep $sleep
done
