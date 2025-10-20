--🌱🌱🌱
/*Le SEED (ensemencement): peupler la base avec des données initiales.
Consiste à insérer des données de départ dans la base, souvent utiles pour :
Tester l'application
Préparer les catégories, rôles, ou utilisateurs par défaut
Simuler des contenus (topics, messages…)*/
--🌱🌱🌱

INSERT INTO categories (name) VALUES 
	("B. par époques"), 
	("les métiers de B."), 
	("Les accessoires de B."), 
	("K.");

INSERT OR IGNORE INTO users (username, password, email) VALUES 
	("fadadu13", "$2a$10$TE2ClmvgHd6p7I.IM.kyNO.MYUy3RNfpzz4R3UDE54yAX1MMufaXS", "fafa@wanadoo.com"),
	("kékédu76", "$2a$10$HB9ILjInaP/X2P9rHMgWmOBe9joo2I89z87J5IA7FgPhphz7r0fZ.", "stevenpq@aol.fr");

INSERT OR IGNORE INTO topics (title, content, user_id) VALUES 
	("B. vintage 50s", "j'adore le regard sournois des poupées originales!", 1),
	("B. vétérinaire", "j'ai adoré cette gamme, mais j'aurais aimé d'autres animaux que chiens et chats à soigner...", 1),
	("Le cabriolet", "j'adore le fait que la voiture n'aie pas de moteur, car il faut de l'espace pour les jambes de B. qui ne se plient pas XD!!", 1),
	("B. fluo 90s", "j'adore le fluo, je me suis sentie trop comblée quand cette gamme est sortie", 2);

INSERT OR IGNORE INTO posts (content, topic_id, user_id) VALUES 
	("moi aussi, elle plissent un peu les yeux comme si elles te jugeaient avec mépris. Trop mignon!", 1, 2),
	("clairement ils sont où les furets, iguanes, et tarantules domestiques?", 1, 2),
	("Elle avance grâce au pouvoir de l'amour et de l'imagination...", 2, 2),
	("XD", 3, 1);

INSERT OR IGNORE INTO topic_categories (topic_id, category_id) VALUES
	(1,1),
	(2,2),
	(3,3),
	(4,1);