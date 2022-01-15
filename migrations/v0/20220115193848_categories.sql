-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  uuid uuid NOT NULL UNIQUE,
  name varchar(256) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);
CREATE INDEX idx_categories_name ON categories (name);


CREATE TABLE IF NOT EXISTS works_category (
  id SERIAL PRIMARY KEY,
  uuid uuid NOT NULL UNIQUE,
  category_id INTEGER NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories,works_category;
-- +goose StatementEnd
