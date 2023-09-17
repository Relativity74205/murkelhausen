# Prefect

## Install

on bi-automate

```bash
pipx install prefect
pipx inject prefect prefect-shell
```

## Setup

after supervisorctl setup:

```bash
prefect cloud login
```

## Test

- e.g. from `/home/beowulf/prefect`:
  ```bash
  python backup.py
  ```


## Deployment (from beowulf)

- create deployment
  ```bash
  prefect deployment build ./backup.py:beowulf -n backup_beowulf -q beowulf --cron "0 2 * * *" --timezone Europe/Berlin
  ```

- apply deployment
  ```bash
  prefect deployment apply beowulf-deployment.yaml
  ```
  
- run deployment
  ```bash
  prefect deployment run "beowulf backup and monitoring/backup_beowulf"
  ```



