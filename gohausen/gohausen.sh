#!/bin/bash

nohup ./gohausen &> logs_gohausen.log &
echo $! > pid_gohausen

kill $(cat pid_gohausen)