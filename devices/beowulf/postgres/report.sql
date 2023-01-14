CREATE SCHEMA report;
CREATE USER report WITH PASSWORD '';
ALTER SCHEMA report OWNER TO report;
GRANT USAGE ON SCHEMA kafka TO report;
GRANT SELECT ON ALL TABLES IN SCHEMA kafka to report;
ALTER DEFAULT PRIVILEGES IN SCHEMA kafka GRANT SELECT ON TABLES TO report;


-- view DDLs --
CREATE OR REPLACE VIEW report.aqara_sensor AS
    SELECT tstamp, message_tstamp, sensorname, battery, humidity, linkquality, power_outage_count, pressure, temperature, voltage FROM kafka.aqara_sensor;

CREATE OR REPLACE VIEW report.murkelhausen_states AS
    SELECT tstamp, message_tstamp, hostname, uptime, memory_total, memory_available, memory_used, memory_used_percent, memory_free, cpu_cores, cpu_logical, cpu_usage_avg, root_disk_total, root_disk_free, root_disk_used, root_disk_used_percent, load01, load05, load15, network_bytes_sent, network_bytes_recv, process_count
        FROM kafka.murkelhausen_states_v2
UNION
    SELECT tstamp, message_tstamp, hostname, uptime, memory_total, memory_available, memory_used, memory_used_percent, memory_free, cpu_cores, cpu_logical, cpu_usage_avg, root_disk_total, root_disk_free, root_disk_used, root_disk_used_percent, load01, load05, load15, network_bytes_sent, network_bytes_recv, process_count
        FROM kafka.murkelhausen_states_v1
UNION
    SELECT tstamp, message_tstamp, hostname, uptime, memory_total, memory_available, memory_used, memory_used_percent, memory_free, cpu_cores, cpu_logical, cpu_usage_avg, root_disk_total, root_disk_free, root_disk_used, root_disk_used_percent, load01, load05, load15, network_bytes_sent, network_bytes_recv, process_count
        FROM kafka.murkelhausen_states_v0
;

CREATE OR REPLACE VIEW report.power_data AS
    SELECT tstamp, message_tstamp, sensorname, power_current, power_total FROM kafka.power_data_v3
UNION
    SELECT tstamp, message_tstamp, sensorname, power_current, power_total FROM kafka.power_data_v2
UNION
    SELECT to_timestamp(tstamp, 'YYYY-MM-DDThh24:mi:ss')::timestamp AS tstamp, message_tstamp, sensorname, power_current, power_total FROM kafka.power_data_v1
UNION
    SELECT to_timestamp(tstamp, 'YYYY-MM-DDThh24:mi:ss')::timestamp AS tstamp, message_tstamp, sensorname, power_current, power_total FROM kafka.power_data_v0
;

CREATE OR REPLACE VIEW report.shelly_flood AS
    SELECT tstamp, message_tstamp, sensorname, temperature, flood, id FROM kafka.shelly_flood;

CREATE OR REPLACE VIEW report.shelly_ht_sensor AS
    SELECT tstamp, message_tstamp, sensorname, temperature, humidity, id FROM kafka.shelly_ht_sensor;

CREATE OR REPLACE VIEW report.xiaomi_mi_sensor AS
    SELECT tstamp, message_tstamp, sensorname, temperature, humidity, battery, linkquality, power_outage_count, voltage FROM kafka.xiaomi_mi_sensor;
