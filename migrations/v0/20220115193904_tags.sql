-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tags (
  id SERIAL PRIMARY KEY,
  uuid uuid NOT NULL UNIQUE,
  name varchar(256) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS backnumber_tag (
  id SERIAL PRIMARY KEY,
  uuid uuid NOT NULL UNIQUE,
  tag_id Integer NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  FOREIGN KEY (tag_id) REFERENCES tags(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

CREATE INDEX idx_backnumber_tag_created_at ON backnumber_tag (created_at);
CREATE INDEX idx_backnumber_tag_updated_at ON backnumber_tag (updated_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tags, backnumber_tag;
-- +goose StatementEnd
