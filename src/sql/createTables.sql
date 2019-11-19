CREATE TABLE IF NOT EXISTS users(
  id SERIAL PRIMARY KEY,
  username varchar(50) DEFAULT '',
  password TEXT DEFAULT ''
);

INSERT INTO users VALUES(DEFAULT, 'brody', 'test');

CREATE TABLE IF NOT EXISTS posts(
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id),
  mood VARCHAR(20) NOT NULL,
  genre VARCHAR(20) NOT NULL,
  post TEXT NOT NULL,
  time_stamp timestamp DEFAULT 'now'
);

INSERT INTO posts VALUES(DEFAULT, 1, 'ok', 'server', 'first message');

 
