PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
	users_id INTEGER PRIMARY KEY AUTOINCREMENT,
	users_username TEXT NOT NULL VARCHAR(50) UNIQUE,
	users_password TEXT NOT NULL,
	users_email TEXT NOT NULL UNIQUE,
	users_created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
	categories_id INTEGER PRIMARY KEY AUTOINCREMENT,
	categories_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS topics (
	topics_id INTEGER PRIMARY KEY AUTOINCREMENT,
	topics_title TEXT NOT NULL,
	topics_content TEXT NOT NULL,
	topics_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	topics_updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	topics_category_id INTEGER NOT NULL,
	topics_user_id INTEGER NOT NULL,
	FOREIGN KEY (topics_id) REFERENCES categories(categories_id),
	FOREIGN KEY (topics_user_id) REFERENCES users(users_id)
);

CREATE TABLE IF NOT EXISTS messages (
	messages_id INTEGER PRIMARY KEY AUTOINCREMENT,
	messages_content TEXT NOT NULL,
	messages_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	messages_updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	messages_topic_id INTEGER NOT NULL,
	messages_user_id INTEGER NOT NULL,
	FOREIGN KEY (messages_id) REFERENCES topics(topics_id),
	FOREIGN KEY (messages_id) REFERENCES users(users_id)
);

CREATE TABLE IF NOT EXISTS reactions (
	reactions_id INTEGER PRIMARY KEY AUTOINCREMENT,
	reactions_user_id INTEGER NOT NULL,
	reactions_value INTEGER NOT NULL CHECK(value IN (-1, 1)), /* -1 = dislike, 1 = like */
	reactions_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	reactions_updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	reactions_target_type TEXT NOT NULL CHECK(reactions_target_type IN ('topic','messages')), /*liste des "cibles": objets sur lesquels mettre des réactions*/
	reactions_target_id INTEGER NOT NULL,
	UNIQUE(reactions_target_type, reactions_target_id, reactions_user_id), /* un unique vote par utilisateur par cible */
	FOREIGN KEY (reactions_user_id) REFERENCES users(users_id)
);

/* Gestion des cookies */
/* CREATE TABLE sessions (
	sessions_id TEXT PRIMARY KEY,
	sessions_user_id INTEGER,
	sessions_created_at DATETIME DEFAULT CURRENT_TIMESTAMP
); */