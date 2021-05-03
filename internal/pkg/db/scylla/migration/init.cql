CREATE KEYSPACE IF NOT EXISTS shortlink
	WITH REPLICATION = {
		'class' : 'SimpleStrategy',
		'replication_factor': 1
	};

CREATE TABLE IF NOT EXISTS shortlink.links (
	url text,
	hash text,
	ddd text,
	PRIMARY KEY(hash)
)
