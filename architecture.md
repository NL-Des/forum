```
forum/
├── cmd/
│	└── server/
│		└── main.go			# Point d'entrée
│
├── docker/
│  	├── docker-compose.yml
│	└── Dockerfile
│
├── migrations/ 			# Définition du contenu structurel de la base de donnée (BdD)
│	├── 001_create_tables.sql
│	├── 002_seed.sql
│  	└── 003_indexes.sql
│
├── internal/
│	│
│  	├── config/				# Fichiers de configurations. Pour l'instant juste celui qui lance la BdD et ses migrations
│  	│	└── initdb.go
│	│
│  	├── domain/				# Définition des Entités (les struct/interface)
│  	│	├── user_type.go
│  	│	└── topic-post_type.go
│	│
│  	├── repositories/       # Accès DB (requêtes SQL pour champs 'user')
│  	│	├── runmigrations.go
│  	│	├── user_repository.go
│  	│	└── topic-post_repository.go
│	│
│  	├── services/           # Logique métier: actions sur chaque "objet"; pour user: validation, hash/cryptage de son mot de passe,
│  	│	├── user_service.go
│  	│	└── topic-post_service.go
│  	│
│  	├── handlers/			# HTTP handlers (routage, controlleur)
│  	│	├── init_handlers.go
│  	│	├── router.go		# fichier dédié à l'ensemble des routes
│  	│	├── user_handler.go
│  	│	├── topic-post_handler.go
│   │	└── home_handler.go
│  	│
│  	└── templates/          # Templates FrontEnd HTML/CSS/JS
│  		├── assets/
│  		├── home.html
│  		├── register.html
│  		└── topic.html
│
├── go.mod
└── go.sum
```