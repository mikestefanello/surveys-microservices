CREATE DATABASE voting;

\c voting

CREATE TABLE votes (
  id TEXT PRIMARY KEY,
  survey TEXT,
  question INT,
  created TIMESTAMP
);

CREATE TABLE results (
  survey TEXT,
  question INT,
  votes INT,
  last_update TIMESTAMP,
  PRIMARY KEY(survey, question)
);