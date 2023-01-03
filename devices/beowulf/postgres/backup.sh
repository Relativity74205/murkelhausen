#!/bin/bash

backup_path=/home/arkadius/backup/postgres
postgres_path=/home/arkadius/postgres
backup_date=$(date +%Y-%m-%d_%H_%M_%S)

# https://stackoverflow.com/questions/24718706/backup-restore-a-dockerized-postgresql-database
docker-compose -f ${postgres_path}/docker-compose.yml exec postgres pg_dumpall --globals-only -U postgres > ${backup_path}/globals_${backup_date}.sql
docker-compose -f ${postgres_path}/docker-compose.yml exec postgres pg_dump -F c -U postgres murkelhausen_datastore > ${backup_path}/murkelhausen_datastore_${backup_date}.dump

# https://serverfault.com/questions/196843/logrotate-rotating-non-log-files
# ls -1 /root/backup_* | sort -r | tail -n +6 | xargs rm > /dev/null 2>&1
