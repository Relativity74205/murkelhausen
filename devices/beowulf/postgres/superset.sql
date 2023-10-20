CREATE DATABASE murkelhausen_app;
CREATE USER murkelhausen_app WITH PASSWORD '';
ALTER DATABASE murkelhausen_app_dev OWNER TO murkelhausen_app;

TRUNCATE TABLE statements_commerzbankauszug;
SELECT * FROM statements_commerzbankauszug;