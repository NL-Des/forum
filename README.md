# FORUM

# Objectifs :

Ce projet à pour objectif de créer un forum web qui autorise :
- La communication entre les utilisateurs.
- L'association de catégories aux sujets.
- Pourvoir liker, disliker les sujets et les messages.
- Pouvoir filtrer les sujets.

## Database :
- Pour stocker les données dans notre forum, il nous est demandé d'utiliser SQLite.
- Pour bien visualiser notre base de données, nous devons réaliser un diagramme relationnel : https://www.smartdraw.com/entity-relationship-diagram/
- Nous devons utiliser au moins une fois les requêtes SELECT, CREATE et INSERT

## Authentification :
- L'utilisateur doit être capable de créer un nouveau compte en enregistrant ces informations. (Action Register)
    * email : attention, un email ne peut-être utilisé deux fois.
    * nom d'utilisateur : même problématique.
    * mot de passe.
    * dotation des droits de base de l'utilisateur.
    
- L'utilisateur doit se connecter pour accéder au forum, créer des sujets et faire des messages. (Action Login Session)
- Nous devrons utiliser les cookies pour restreindre chaque utilisateur à ne pouvoir ouvrir qu'une seule session.
- Chaque session ouverte doit avoir une date d'expiration, pour éviter d'être ouverte en permanence. 

## Communication dans le serveur :
- Les utilisateurs non enregistrés et enregistrés peuvent :
    * voir les sujets.
    * voir les réponses des sujets.
    * voir le nombre de likes.
    * voir le nombre de dislike.
    * voir les catégories des sujets du forum.
    * utiliser le filtre pour les catégories

- Seuls les utilisateurs enregistrés peuvent :
    * créer des sujets.
    * créer des messages.
    * associer une ou plusieurs catégories a un sujet (à sa création ?).
    * mettre un like ou un dislike sur un sujet ou un message. Impossibilité de faire les deux sur un même sujet ou un même message
    * voir les sujets créés par chaque utilisateur.
    * voir les sujets likés par les utilisateurs.
    * utiliser le filtre pour les catégories, les messages créés et les likés par la personne utilisant le filtre.

## Docker :
- Nous devons utiliser Docker pour le déploiement du forum.

## Catégories :
- Nous devons choisir un thème pour le forum.

## Packages autorisés :
- GO
- squlite3 : https://github.com/mattn/go-sqlite3
- bcrypt : https://pkg.go.dev/golang.org/x/crypto/bcrypt
- gofrs/uuid (https://github.com/gofrs/uuid) ou google/uuid (https://github.com/google/uuid)

## Bonus du projet de base :
- script de création de l'image et des conteneurs.


# OPTIONS :

## Fonctionnalités avancées :
- L'utilisateur d'un sujet est :
    * notifié d'un like.
    * notifié d'un dislike.
    * notifié d'un message sur son sujet.

- L'utilisateur d'un message est :
    * notifié d'un like.
    * notifié d'un dislike.

- L'utilisateur doit disposer d'une page personnelle suivant son activité :
    * les sujets créés.
    * les messages créés : l'utilisateur doit voir son message en son entier, mais aussi le sujet d'où provient le message.
    * les likes postés.
    * les dislikes postés.
    * /!\ L'utilisateur doit pouvoir éditer / supprimer sujets et messages dont il est le créateur. Depuis cette page personnelle.

## Authentification :
L'utilisateur doit pouvoir s'authentifier sur le forum par au minimum deux moyens extérieurs (Github et Google).

## Poster des images :
L'utilisateur doit pouvoir mettre des images dans ces messages. 
    - Les formats d'images minimum attendus à pouvoir gérer sont : JPEG, PNG, GIF.
    - La taille des images ne doit pas dépasser 20 Mo.

## Modération :
Un modérateur doit pouvoir, suivant le niveau d'accès de l'utilisateur, approuver les messages qui seront visible pour tout le monde. Il classera les messages retenu en quatre catégories :
    - hors-sujet.
    - obscène.
    - illégal.
    - insultant.

Les utilisateurs auront quatre statuts d'accréditations dans les droits d'utilisations du forum :
- invité : utilisateur non-enregistré.
    * ils ne peuvent que voir les sujets, messages, likes et dislikes.

- utilisateur :
  * créer des sujets.
  * créer des messages.
  * affilier des catégories.
  * liker.
  * disliker.

- modérateur : en plus des droits d'utilisateur, ils surveillent le contenu des utilisateurs par :
    * la supression du sujet ou message.
    * le signalement à l'administrateur. 
    * /!\ Seul l'administrateur peut élever un utilisateur au statut de modérateur.

- administrateur : 
    * élever le statut d'un utilisateur à celui de modérateur, et inversement.
    * réception des rapports de signalement des modérateurs.
    * réponse aux modérateurs de leurs rapports de signalement.
    * suppression des sujets et des messages.
    * gestion des catégories, création et supression.

## Sécurité :
- implémenter un protocole HTTPS : https://developer.mozilla.org/en-US/docs/Glossary/HTTPS
- générer un certificat de sécurité SSL. Nous pouvons créer le nôtre ou utiliser le "Certificate Authorities"(CA's)
- intégrer une limite de requêtes http possibles par utilisateur : https://en.wikipedia.org/wiki/Rate_limiting
- chiffrer les mots de passes des utilisateurs dans la base de donnée.
- créer un mot de passe pour accéder à la base de donnée.
- chaque cookie doit être unique et stocké sur le serveur, pour limiter la lecture et la falsification par un attaquant de la session.
- Indice pour aider à la sécurisation : 
    * https://en.wikipedia.org/wiki/Cipher_suite
    * manuel d'utilisation openssl
    * pour les cookies : https://en.wikipedia.org/wiki/Universally_unique_identifier