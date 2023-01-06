CREATE SCHEMA kafka;
CREATE USER kafka_connect WITH PASSWORD '';
ALTER SCHEMA kafka OWNER TO kafka_connect;
GRANT USAGE ON SCHEMA kafka TO kafka_connect;
GRANT CREATE ON SCHEMA kafka TO kafka_connect;
-- ALTER TABLE kafka.* OWNER TO kafka_connect;
