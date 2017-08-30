-- Table: data

-- DROP TABLE data;

CREATE TABLE data
(
  article integer,
  title text,
  note text,
  sum integer,
  reject integer
)
WITH (
  OIDS=FALSE
);
ALTER TABLE data
  OWNER TO test;
