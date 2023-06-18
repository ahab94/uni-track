CREATE KEYSPACE streaming_service WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
    };


CREATE TABLE streaming_service.packets
(
    user_id         text,
    timestamp       timestamp,
    channel_id      text,
    sequence_number bigint,
    PRIMARY KEY ((user_id), timestamp, channel_id, sequence_number)
) WITH CLUSTERING ORDER BY (timestamp DESC);

CREATE TABLE streaming_service.hourly_usage
(
    hour       timestamp,
    user_id    text,
    channel_id text,
    pkt_count  counter,
    PRIMARY KEY ((user_id), hour, channel_id)
) WITH CLUSTERING ORDER BY (hour DESC);

CREATE TABLE streaming_service.users
(
    user_id   text,
    timestamp timestamp,
    PRIMARY KEY (user_id, timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);
