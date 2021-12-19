BEGIN;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  uuid varchar(256) NOT NULL UNIQUE,
  name varchar(256) NOT NULL,
  email varchar(256) UNIQUE NOT NULL,
  photo_url text NOT NULL,
  suspend bool NOT NULL DEFAULT false,

  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_updated_at ON users (updated_at);
CREATE INDEX idx_users_created_at ON users (created_at);


CREATE TABLE IF NOT EXISTS user_followers (
  id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  follower_id INTEGER NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
  FOREIGN KEY (follower_id) REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

COMMIT;