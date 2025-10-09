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

INSERT INTO users (username, password, email) VALUES 
	("fadadu13", "hashedpassword13", "fafa@wanadoo.com"),
	("kékédu76", "76hashedpassword", "stevenpq@aol.fr");


INSERT INTO topics (title, content, category_id, user_id) VALUES 
	("B. vintage 50s", "j'adore le regard sournois des poupées originales!", 1, 1),
	("B. vétérinaire", "j'ai adoré cette gamme, mais j'aurais aimé d'autres animaux que chiens et chats à soigner...", 2, 1),
	("Le cabriolet", "j'adore le fait que la voiture n'aie pas de moteur, car il faut de l'espace pour les jambes de B. qui ne se plient pas XD!!", 3, 1),
	("B. fluo 90s", "j'adore le fluo, je me suis sentie trop comblée quand cette gamme est sortie", 1, 2);

INSERT INTO messages (content, topic_id, user_id) VALUES 
	("moi aussi, elle plissent un peu les yeux comme si elles te jugeaient avec mépris. Trop mignon!", 1, 2),
	("clairement ils sont où les furets, iguanes, et tarantules domestiques?", 2, 2),
	("Elle avance grâce au pouvoir de l'amour et l'imagination...", 3, 2),
	("XD", 3, 1);
