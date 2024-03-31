create table t_user (
    id serial NOT NULL,
    username CHARACTER(50),
    password CHARACTER(50),
    age INTEGER,
    mobile CHARACTER(50),
    address CHARACTER(50),
    status INTEGER,
    role CHARACTER(50),
    sex CHARACTER(50),
    CONSTRAINT user_pkey PRIMARY KEY (id)
)
WITH (OIDS=FALSE);