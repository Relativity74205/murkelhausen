#!/bin/bash

if [[ -z "${POSTGRES_BACKUP_PATH}" ]]; then
  postgres_backup_path=/home/arkadius/backup/postgres
else
  postgres_backup_path=${POSTGRES_BACKUP_PATH}
fi

if [[ -z "${POSTGRES_PATH}" ]]; then
  postgres_path=/home/arkadius/postgres
else
  postgres_path=${POSTGRES_PATH}
fi

backup_datetime=$(date +%Y-%m-%dT%H_%M_%S)

# https://stackoverflow.com/questions/24718706/backup-restore-a-dockerized-postgresql-database
docker-compose -f ${postgres_path}/docker-compose.yml exec -T postgres pg_dumpall --globals-only -U postgres > ${postgres_backup_path}/${backup_datetime}__globals.sql
docker-compose -f ${postgres_path}/docker-compose.yml exec -T postgres pg_dump -F c -U postgres murkelhausen_datastore > ${postgres_backup_path}/${backup_datetime}__murkelhausen_datastore.dump
