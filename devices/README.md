# Devices

## Backup

### Synology NAS

Shared folders on Synology NAS:

- rasp1_backup
- rasp2_backup
- beowulf_backup

with following settings:

- Permissions:
  - admin: Read/Write
  - Arkadius: Read/Write
  - murkelhausen_admin: Read/Write
- NFS Permissions:
  - HostName or IP: depending on device IP
  - Privilege: Read/Write
  - Squash: Map all users to admin
  - Security: sys

### Settings on devices

Add the following line to /etc/fstab

For PIs:

```bash
192.168.1.19:/volume1/backup_rasp1 /home/pi/backup nfs defaults 0 0
192.168.1.19:/volume1/backup_rasp2 /home/pi/backup nfs defaults 0 0
```

For NUK:

```bash
sudo apt install nfs-common
192.168.1.19:/volume1/backup_beowulf /home/arkadius/backup nfs defaults 0 0
```

#### cron

rasp1:

```cron
0 1 * * * /usr/local/bin/pihole -a -t "/home/pi/backup/pihole-$(hostname)-teleporter-$(date -I).tar.gz"
```

rasp2:

```cron
0 1 * * * /usr/local/bin/pihole -a -t "/home/pi/backup/pihole/pihole-$(hostname)-teleporter-$(date -I).tar.gz"
0 3 * * * rsync -uv --dirs --delete "/home/pi/unifi_config/data/backup/autobackup/" "/home/pi/backup/unifi"
```
