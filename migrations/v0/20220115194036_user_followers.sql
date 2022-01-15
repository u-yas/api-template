-- +goose Up
-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_followers;
-- +goose StatementEnd
