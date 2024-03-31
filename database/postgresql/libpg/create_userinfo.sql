CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username CHARACTER VARYING(100) NOT NULL,
    departname CHARACTER VARYING(500) NOT NULL,
    created DATE,
    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);