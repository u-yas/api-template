BEGIN;

CREATE TABLE IF NOT EXISTS user_favorites_backnumber (
  id SERIAL PRIMARY KEY,
  uuid varchar(256) NOT NULL UNIQUE,
  user_id integer NOT NULL,
  backnumber_id integer NOT NULL, 
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
  FOREIGN KEY (backnumber_id) REFERENCES backnumbers(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

CREATE INDEX idx_user_favorites_backnumber_created_at ON user_favorites_backnumber (created_at);
CREATE INDEX idx_user_favorites_backnumber_updated_at ON user_favorites_backnumber (updated_at);

CREATE TABLE IF NOT EXISTS user_favorites_works (
  id SERIAL PRIMARY KEY,
  uuid varchar(256) NOT NULL UNIQUE,
  user_id integer NOT NULL,
  works_id integer NOT NULL, 
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
  FOREIGN KEY (works_id) REFERENCES works(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

CREATE INDEX idx_user_favorites_works_created_at ON user_favorites_works (created_at);
CREATE INDEX idx_user_favorites_works_updated_at ON user_favorites_works (updated_at);

COMMIT;