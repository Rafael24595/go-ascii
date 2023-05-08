CREATE DATABASE logdb;

\c logdb;

CREATE TABLE IF NOT EXISTS session (
  session_id varchar(50) NOT NULL,
  timestamp TIMESTAMP,
  PRIMARY KEY (session_id)
);

CREATE TABLE IF NOT EXISTS register (
  register_id SERIAL NOT NULL,
  session_id_fk varchar(50) NOT NULL,
  category varchar(30) NOT NULL,
  family varchar(30) NOT NULL,
  message varchar(500) NOT NULL,
  timestamp BIGINT,
  PRIMARY KEY (register_id),
  CONSTRAINT session_id_fk
      FOREIGN KEY(session_id_fk) 
	  REFERENCES session(session_id)
);