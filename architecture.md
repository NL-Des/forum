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
│  	├── domain/				# Définition des Entités (les struct/interface)
│  	│	└── user_handler.go
│	│
│  	├── repositories/       # Accès DB (requêtes SQL pour champs 'user')
│  	│	└── user_repository.go
│	│
│  	├── services/           # Logique métier: actions sur chaque "objet"; pour user: validation, hash/cryptage de son mot de passe,
│  	│	└── user_service.go
│  	│
│  	├── handlers/			# HTTP handlers (routage, controlleur)
│   │	└── user_handler.go
│  	│
│  	└── templates/          # Templates FrontEnd HTML/CSS/JS
│  		├── index.html
│  		└── register.html
│
├── go.mod
└── go.sum
```