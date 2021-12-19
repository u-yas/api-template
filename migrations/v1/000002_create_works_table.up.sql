BEGIN;

CREATE TABLE IF NOT EXISTS works (
  id SERIAL PRIMARY KEY, 
  uuid varchar(256) NOT NULL UNIQUE,
  user_id integer NOT NULL,
  title varchar(256) UNIQUE NOT NULL,
  description varchar(256) NOT NULL,
  tags text[],
  bookmarks_count int default 0 NOT NULL,
  thumbnail_url varchar(256) NOT NULL,
  public boolean default true NOT NULL,
  gcs_path varchar(256) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX idx_works_title ON works (title);
CREATE INDEX idx_works_updated_at ON works (updated_at);
CREATE INDEX idx_works_created_at ON works (created_at);

CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  uuid varchar(256) NOT NULL UNIQUE,
  name varchar(256) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS works_category (
  id SERIAL PRIMARY KEY,
  uuid varchar(256) NOT NULL UNIQUE,
  category_id INTEGER NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

COMMIT;