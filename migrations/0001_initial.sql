CREATE SCHEMA IF NOT EXISTS uptime;

CREATE TABLE IF NOT EXISTS uptime.monitor (
    id SERIAL PRIMARY KEY,
    interval interval NOT NULL ,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS uptime.probe (
    id SERIAL PRIMARY KEY,
    definition jsonb,
    monitor_id INTEGER,
    FOREIGN KEY (monitor_id) REFERENCES uptime.monitor (id)
);

CREATE TABLE IF NOT EXISTS uptime.heartbeat (
    timestamp TIMESTAMPTZ NOT NULL,
    monitor_id INTEGER,
    success BOOLEAN,
    FOREIGN KEY (monitor_id) REFERENCES uptime.monitor(id)
);

SELECT create_hypertable('uptime.heartbeat', by_range('timestamp'));
INSERT INTO uptime.monitor(name, interval) VALUES ('My First Monitor', '30 seconds');
INSERT INTO uptime.probe(definition, monitor_id) VALUES ('{"url": "https://google.com"}', 1);