PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	token TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS topics (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	category_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	UNIQUE(title, content, user_id),
	FOREIGN KEY (category_id) REFERENCES categories(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	topic_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	UNIQUE(content, topic_id, user_id),
	FOREIGN KEY (topic_id) REFERENCES topics(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS reactions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	value INTEGER NOT NULL CHECK(value IN (-1, 1)), /* -1 = dislike, 1 = like */
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	target_type TEXT NOT NULL CHECK(target_type IN ('topics','posts')), /*liste des "cibles": objets sur lesquels mettre des réactions*/
	target_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	UNIQUE(target_type, target_id, user_id), /* un unique vote par utilisateur par cible */
	FOREIGN KEY (user_id) REFERENCES users(id)
);