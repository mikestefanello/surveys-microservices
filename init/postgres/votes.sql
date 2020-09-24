CREATE DATABASE voting;

\c voting

CREATE TABLE votes (
  id TEXT PRIMARY KEY,
  survey TEXT,
  question INT,
  created BIGINT
);

CREATE TABLE results (
  survey TEXT,
  question INT,
  votes INT,
  last_update BIGINT,
  PRIMARY KEY(survey, question)
);