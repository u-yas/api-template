-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS backnumbers (
  id SERIAL PRIMARY KEY,
  uuid uuid NOT NULL UNIQUE,
  works_id integer NOT NULL,
  user_id integer NOT NULL,
  title text NOT NULL,
  description text NOT NULL,
  thumbnail_url varchar(256) NOT NULL,
  view_count Integer DEFAULT 0 NOT NULL,
  public boolean DEFAULT true NOT NULL,
  backnumber_url varchar(256) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
  FOREIGN KEY (works_id) REFERENCES works(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

CREATE INDEX idx_backnumbers_title ON backnumbers (title);
CREATE INDEX idx_backnumbers_updated_at ON backnumbers (updated_at);
CREATE INDEX idx_backnumbers_created_at ON backnumbers (created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE backnumbers;
-- +goose StatementEnd
