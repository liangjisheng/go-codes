CREATE TABLE IF NOT EXISTS "user" (
    "id" serial NOT NULL PRIMARY KEY,
    "username" text NOT NULL DEFAULT '' ,
    "age" integer NOT NULL DEFAULT 0 ,
    "sex" text NOT NULL DEFAULT '' ,
    "mobile" text NOT NULL DEFAULT '' ,
    "password" text NOT NULL DEFAULT '' ,
    "email" text NOT NULL DEFAULT '' 
)
WITH (OIDS=FALSE);