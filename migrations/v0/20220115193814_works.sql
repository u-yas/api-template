-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS works (
  id SERIAL PRIMARY KEY, 
  uuid uuid NOT NULL UNIQUE,
  user_id integer NOT NULL,
  title varchar(256) UNIQUE NOT NULL,
  description varchar(256) NOT NULL,
  bookmarks_count int default 0 NOT NULL,
  thumbnail_url varchar(256) NOT NULL,
  public boolean default true NOT NULL,
  works_url varchar(256) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX idx_works_title ON works (title);
CREATE INDEX idx_works_updated_at ON works (updated_at);
CREATE INDEX idx_works_created_at ON works (created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE works;
-- +goose StatementEnd
