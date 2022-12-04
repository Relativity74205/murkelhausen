
# IPs

- superset
- kafka
- zigbeemqtt
- unifi
- C430A GO
- NAS
- mqtt
- pi-hole
- shellies:
  - floodmeter:
  -
- Smartmeter IR Leser:
  - Haushalt: 192.168.1.80
  - Wärmepumpe: 192.168.1.84
  
# TODO list

- add gohausen cron command at startup
- add kafka docker compose command at startup
- add upsert/ignore duplicates for kafka connect
- powerdata make PowerCurrentPX optional
- use tstamp columns in superset
- [x] create start/stop script
- create deploy script
- create github pipeline
- create versioning system
- change postgres connect name
- create connect deploy script
- look where to increase security in
  - zigbeetomqtt
  - mosquitto
  - kafka connect
  - postgres
  - superset
- add weather api data
- add google traffic forecast data
- create scheduling system
- add sonos control consumer

- NAS mount (<https://michael-casey.com/2022/01/09/mount-synology-nas-to-raspberry-pi-using-nfs/> <https://eliaslundgaard.com/posts/mount-nas-to-pi/> )
  - rasp1
  - raps2
  - NUK
- backup unifi backups
- backup pi-hole settings and pihole-FLT.db
  - <https://docs.pi-hole.net/database/ftl/?h=backup#backup-database>
- pi-hole
  - redundancy:
    - [x] <https://www.reddit.com/r/pihole/comments/692vaf/what_if_the_pi_goes_down/>
    - [x] second pi-hole in parallel
    - keep in sync: <https://github.com/vmstan/gravity-sync>
  - backup settings (<https://www.reddit.com/r/pihole/comments/ncjzxx/how_to_properly_backup_my_pihole_setup/>)
- force dns traffic
  - <https://scotthelme.co.uk/catching-naughty-devices-on-my-home-network/>
  - <https://fictionbecomesfact.com/usg-redirect-pihole>

- collectd to kafka (<https://github.com/collectd/collectd>)
  - rasp1
  - NUK
- domain names for internal ips?
