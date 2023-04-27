package database

const schema = `
	create table cordle_test.users (
	    id     bigint          not null		primary key,
	    wins   int default 0   not null,
	    losses int default 0   not null,
	    draws  int default 0   not null,
	    elo    int default 500 not null,
	    constraint id
	        unique (id)
	);`
