
# IPs

-
- C430A GO (192.168.1.24; fixed)
- NAS (192.168.1.19; fixed)
- beowulf (192.168.1.69; fixed)
  - mqtt
  - superset <http://192.168.1.69:8088/>
  - kafka <http://192.168.1.69:9021/>
  - zigbeemqtt <http://192.168.1.69:8080/>
- rasp1 (192.168.1.18 (fixed)):
  - pi-hole: <http://rasp1.local/admin/>
- rasp2 (192.168.1.28 (fixed))
  - pi-hole: <http://rasp2.local/admin/>
  - unifi: <https://rasp2.local:8443/>
- shellies:
  - floodmeter:
- Smartmeter IR Leser:
  - Haushalt: 192.168.1.84; fixed
  - Wärmepumpe: 192.168.1.85; fixed
  
# TODO list

- add gohausen cron command at startup
  - alternative: supervisord script
- add kafka docker compose command at startup
  - change restart policy to always
- add upsert/ignore duplicates for kafka connect
  - why?
- [x] powerdata make PowerCurrentPX optional
- use tstamp columns in superset
- create deadletterqueue for postgres sink
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
- monitoring:
  - are processes and containers running?
  - are backups created?
  - are NAS NFS mounts online?

- [x] NAS mount (<https://michael-casey.com/2022/01/09/mount-synology-nas-to-raspberry-pi-using-nfs/> <https://eliaslundgaard.com/posts/mount-nas-to-pi/> )
  - [x] rasp1
  - [x] raps2
  - [x] NUK
- [x] backup unifi backups
- [x] backup pi-hole settings and pihole-FLT.db (<https://docs.pi-hole.net/database/ftl/?h=backup#backup-database>)
- pi-hole
  - [x] redundancy:
    - [x] <https://www.reddit.com/r/pihole/comments/692vaf/what_if_the_pi_goes_down/>
    - [x] second pi-hole in parallel
    - [x] keep in sync: <https://github.com/vmstan/gravity-sync>
  - [x] backup settings (<https://www.reddit.com/r/pihole/comments/ncjzxx/how_to_properly_backup_my_pihole_setup/>)
  - pihole update block lists
  - set clients in pi-hole
- force dns traffic
  - <https://scotthelme.co.uk/catching-naughty-devices-on-my-home-network/>
  - <https://fictionbecomesfact.com/usg-redirect-pihole>

- collectd to kafka (<https://github.com/collectd/collectd>)
  - rasp1
  - NUK
- domain names for internal ips?
