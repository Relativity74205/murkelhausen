#!/bin/bash

start() {
  nohup ./gohausen &> logs_gohausen.log &
  echo $! > pid_gohausen
  echo "started gohausen and created pid file"
}

stop() {
  kill $(cat pid_gohausen)
  echo "Send kill command to gohausen, sleep for 3 seconds"
  sleep 3

  if pgrep -F pid_gohausen > /dev/null
  then
    echo "gohausen still running, unknown problem"
  else
    echo "gohausen stopped, removing pid file"
    rm pid_gohausen
  fi
}

command=$1

${subcommand} $@
