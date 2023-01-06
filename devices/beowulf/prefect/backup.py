"""
prefect deployment build ./backup.py:beowulf -n backup_beowulf -q beowulf --cron "0 2 * * *"
"""

from datetime import datetime
from pathlib import Path
from dateutil import relativedelta

from prefect import flow, task, get_run_logger
from prefect_shell import shell_run_command
from prefect.task_runners import ConcurrentTaskRunner

POSTGRES_BACKUP_PATH = "/home/arkadius/backup/postgres"
POSTGRES_PATH = "/home/arkadius/postgres"
POSTGRES_DATABASE_PREFIX = "murkelhausen_datastore"
POSTGRES_GLOBALS_PREFIX = "globals"
POSTGRES_BACKUP_LAST_COUNT = 2


def get_months_between_dates(date1, date2) -> int:
    r = relativedelta.relativedelta(date1, date2)
    months = (r.years * 12) + r.months
    return months


@task
def postgres_backup_cleanup():
    logger = get_run_logger()
    # TODO add error handling
    files = [
        ele
        for ele in Path(POSTGRES_BACKUP_PATH).glob(f"*{POSTGRES_DATABASE_PREFIX}.dump")
        if ele.is_file()
    ]
    files_with_dates = [
        (
            file,
            datetime.strptime(file.stem.split("__")[0], "%Y-%m-%dT%H_%M_%S"),
        )
        for file in files
    ]

    files_with_dates.sort(key=lambda x: x[1], reverse=True)

    files_to_delete = [
        (file, file_date)
        for i, (file, file_date) in enumerate(files_with_dates)
        if not (
            # keep the last x files
            i < POSTGRES_BACKUP_LAST_COUNT
            # all files from beginning of the month shall be kept
            or file_date.day == 1
            # all files from sunday shall be kept, however only for the last x months
            or file_date.weekday() == 6
            and get_months_between_dates(datetime.today(), file_date) <= 3
        )
    ]
    logger.info(f"{files_to_delete=}")
    for file, file_date in files_to_delete:
        globals_file = (
            file.parent
            / f"{file_date.strftime('%Y-%m-%dT%H_%M_%S')}__{POSTGRES_GLOBALS_PREFIX}.sql"
        )
        file.unlink()
        globals_file.unlink()

        logger.info(f"Deleted {file} and {globals_file}.")


@task
def backup_kafka():
    """
    https://docs.confluent.io/platform/current/kafka-rest/api.html

    broker:
    curl http://localhost:8082/v3/clusters | jq
    curl http://localhost:8082/v3/clusters/3x4LP0wLSdm1jZGXUxfYZw/brokers/1/configs | jq

    schema registry:
    curl http://localhost:8081/subjects | jq
    curl http://localhost:8081/subjects/power_data_v2-value/versions/1 | jq

    """
    pass


@task
def backup_mosquitto():
    pass


@task
def backup_zigbee2mqtt():
    pass


@task
def monitor_docker_processes():
    pass


@task
def monitor_supervisor_processes():
    pass


@flow(name="beowulf backup and monitoring", task_runner=ConcurrentTaskRunner())
def beowulf():
    logger = get_run_logger()
    postgres_backup = shell_run_command.with_options(name="postgres_backup").submit(
        command="/home/arkadius/postgres/backup.sh",
        return_all=False,
        env={
            "POSTGRES_BACKUP_PATH": POSTGRES_BACKUP_PATH,
            "POSTGRES_PATH": POSTGRES_PATH,
        },
    )
    logger.info("Performed backup of postgres database.")
    postgres_backup_cleanup.submit(wait_for=[postgres_backup])
    backup_kafka.submit()
    backup_mosquitto.submit()
    backup_zigbee2mqtt.submit()
    monitor_docker_processes.submit()
    monitor_supervisor_processes.submit()


if __name__ == "__main__":
    beowulf()
