
-- +migrate Up
CREATE TABLE IF NOT EXISTS expenses(
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	user_id BIGINT NOT NULL,
	paid_at DATE NOT NULL,
	amount INT NOT NULL,
	category INT NOT NULL,
	description TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS expenses;
