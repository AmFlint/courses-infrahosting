# Continuous Integration / Continuous Delivery|Deployment

Dans ce cours, nous apprendrons les bases de la CI/CD, pratique qui nous permet de livrer des applications avec aisance en continue et en assurant une bonne qualité du produit.

## Premier cas d'usage: L'intégration Continue (CI)

Dans ce premier exercice, je vous fournis une [API Back-End développée avec Node.JS](./server), dont l'objectif est d'afficher/créer des messages (cette application a déjà été vue au cours des thématiques précédentes):
![website finale](../tp-aws-ec2/assets/app.png)

Cette application est composées de la façon suivante:
- API Rest, accessible par appels HTTP aux endpoints suivants:
  - `GET /messages`: Affiche une liste de messages contenus dans la base de données
  - `POST /messages`: Créer une notification dans un bus de données (ici, RabbitMQ), afin de déléguer la création de messages en base de données à des Workers
- [Worker](./server/worker.js): Attends des notifications en écoutant sur un bus de données (ici, RabbitMQ), et va effectuer des tâches asynchrones (dans cet exercice, il sert uniquement à créer des messages en BDD, dans l'objectif de complexifier l'infrastructure en ajoutant un bus de données).
- Bus de données (RabbitMQ): Permet une communication sous forme d'évènements dans notre système, par exemple: `une API reçoit une image, la sauvegarde de manière temporaire, puis délègue la compression de l'image à un worker (car très coûteux en ressources)`.
- Base de données (MySQL ici) avec la structure suivante:
  - table `messages`:
    - `id`: int AUTO_INCREMENT, Primary Key
    - `content`: varchar(155), le contenu de notre message

**Le versioning de la base de données est géré avec le package npm `mysql-migrations`**: nous devrons donc lancer les migrations pour configurer notre BDD avec la commande:
```bash
npm migrate:up
# ou avec yarn
yarn migrate:up
```

Pour lancer l'application en local:
- Téléchargez le [dossier server](./server)
- Installer les dépendances (soit depuis votre ordinateur, soit depuis un container Docker)
- Un fichier [docker-compose de développement](./server/docker-compose.yml) est fournis, vous pouvez l'utiliser pour démarrer l'application en local:
  ```bash
  # Lancer les containers MySQL, PHPMYADMIN, RabbitMQ, API et WORKER:
  docker-compose up -d
  # N'oubliez pas de lancer les migrations de la base de données
  docker-compose exec server node migrations.js up
  ```
- Vérifiez que http://localhost:3000 affiche correctement la liste de messages

Dans cet exercice, je vous fournis également:
- Des `tests unitaires`, qui vont tester le bon fonctionnement du composant `services/messages` (récupérer les messages depuis la DB)
- Des `tests fonctionnels` pour tester l'API dans son intégralité:
  - Lister des messages sur `GET /messages`
  - Créer un message sur `POST /messages`
  - Relister les messages pour s'assurer que le message crée précédemment apparait

Pour lancer ces tests (vous aurez besoins de ces commandes pour l'exercice):
```bash
# lancer les tests unitaires
yarn unit

# lancer les tests fonctionnels, l'API doit être joignable sur http://localhost:3000
yarn e2e

# Lancer le linting du code
yarn lint
```

## L'exercice

**Vous aurez besoins de créer des repositories sur GitHub avec le code des différents exercices, ainsi que la configuration de l'outils de CI/CD (Github Actions/GitlabCI ou CircleCI)**.

Au cours de certaines partie (2 et 3), vous aurez besoins d'un serveur sur lequel déployer vos applications, c'est l'occasion de mettre en pratique ce que l'on a vu précédemment avec AWS, Terraform et Ansible pour configurer un serveur Linux sur AWS, et par la suite déployer nos applications dessus avec la `Continuous Delivery`.

### Première partie: CI

À l'aide de l'outils de votre choix (je vous recommande CircleCI ou Github Actions), des informations fournies au-dessus ainsi que le [code de l'application](./server), vous devrez créer et configurer une Pipeline d'Intégration Continue définie par les tâches suivantes:\
- `Linter`: Exécuter un Linting du code, pour s'assurer que les conventions de code sont respectées sur l'intégralité du projet
- `Tests unitaires`: Exécuter les tests unitaires du projet
- `Tests fonctionnels`: Exécuter les tests fonctionnels du projet. **Ici, vous devrez utiliser docker-compose pour lancer les tests sur une stack applicative locale**, n'oubliez pas d'exécuter les migrations de BDD (`docker-compose exec server node migrations.js up`).

Vous l'aurez compris, cette première partie est orientée sur la partie `qualité`: chaque fois qu'un développeur push du code, nous allons nous assurer que la fonctionalité fonctionne correctement, et que la codebase respecte les conventions définies dans le projet.

Dans la seconde partie, nous allons nous concentrer sur la partie Déploiement Continu de manière générale, puis nous l'appliquerons à cette application au cours de la troisième partie.

### Seconde Partie: CD de l'application Zelda (sans Docker, puis avec Docker)

Afin de voir les différents procédés de déploiement continu, nous allons utiliser le [Projet Zelda sur GitHub](https://github.com/AmFlint/hetic-w2-p2019-05), veillez donc à Télécharger le code, et vous créer votre propre Repository afin d'activer de la CD sur votre projet.

Ici, vous aurez besoins de votre propre serveur Linux pour déployer l'application, je ne vous donne pas d'indication là-dessus, vous avez eu l'occasion de le faire à plusieurs reprises.

Je vous conseille cependant d'utiliser le serveur Web `NGINX` pour la première partie, simple à configurer pour déployer l'application Zelda sans utiliser Docker.

Par la suite, il vous faudra installer Docker sur votre serveur, pour exécuter votre container (et le remettre à jours quand vous pousserai une nouvelle version de l'application).

**Les choses à savoir concernant le projet Zelda**:
- Le projet est développé avec Webpack, il faudra donc `builder` le code dans un premier temps: `npm run build | yarn build`.
- Vous aurez besoins d'utiliser une clé SSH liée à votre serveur pour vous y connecter à distance pendant le processus de déploiement (revoir les commandes `scp` et `ssh`).

#### Sans Docker

Vous devrez:
- Installer les dépendances NPM
- Builder l'application avec `npm run build` ou `yarn build`
- Uploader le dossier `build` (résultat de la compilation Webpack) sur votre serveur, derrière votre serveur Web.

#### Avec Docker

Vous devrez:
- Créer un fichier Dockerfile dans lequel il faudra `installer les dépendances npm`, `builder l'application`, puis `servir` les fichiers buildés avec un serveur web comme NGINX par exemple
- Lancer le build de votre image docker depuis la pipeline
- Stocker (`docker push`) votre image docker dans la registry Docker Hub (ou autre), pour cela il faudra utiliser le système de `secrets` pour stocker vos identifiants et les injecter dans votre tâche.
- Vous connectez à votre serveur (via `ssh`) pour lancer/relancer votre container Zelda avec la dernière version disponible.

### Trosième Partie: CD de l'application

Lors de la partie précédente, nous avons vu comment délivrer un site statique en passant par plusieurs étapes:
- Création d'artefacts (Fichier exécutables/utilisables pour le bon fonctionnement de notre application, ici les fichiers html, css, js, medias, fonts ...)
- Connexion à distance au serveur pour lancer/mettre à jours les artefacts présents.

Ici, nous rencontrons un cas d'usage un peu différent, puisqu'il s'agit d'une application comportant plusieurs services (et dont certains services dépendent d'autres):
- Base de données (Et la structure qui va avec, contrôlée par les migrations)
- Bus de données (RabbitMQ)
- API notre programme
- Worker notre second programme

Pour délivrer une nouvelle version de l'application, on va donc devoir:
- Créer un artefact pour notre application (Worker et API partagent le même code, le point d'entrée est différent: `server.js` pour l'API et `worker.js` pour le Worker).
- Récupérer les artefacts sur notre Serveur (Uploader du code via SSH/SCP, ou bien push puis pull d'une image Docker grâce à une Registry comme Docker Hub)
- Mettre à jours la structure de la base de données (Exécuter les migrations)
- Relancer les applications (Worker + API) avec l'artefact basé sur la nouvelle version
