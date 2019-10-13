# TP - Déployer une application sur Heroku

## Pré-requis

- Assurez vous d'avoir [créer un compte sur Heroku.com](https://heroku.com), je vous invite à utiliser vos compte `GitHub Education` pour bénéficier de [différents avantages sur Heroku pour ce TP et sur vos projets persos](https://www.heroku.com/github-students).
- Installez la [ligne de commande heroku](https://devcenter.heroku.com/articles/heroku-cli) pour piloter vos applications (Vous pouvez vous passez de cette CLI en utilisant exclusivement la console Heroku si vous préférez).
- Téléchargez (ou clonez) ce repository sur votre machine, car nous allons utiliser des applications et exemples présents dans ce repo.
- [client MySQL installé sur votre machine](https://dev.mysql.com/doc/refman/8.0/en/mysql.html)

## L'exercice

Dans un premier temps, nous allons nous familiariser avec le déploiement d'une application sur la plateforme Heroku.

Pour ce faire [nous allons déployer l'application disponible ici, dans ce même repository](./app).
L'objectif de cette application est simple:
- Déployer un serveur node.js (avec `express.js`), avec deux endpoints:
  - `/`: affiche un message:
    ```json
    {
      "message": "hello"
    }
    ```
  - `/messages`: Se connecte à une base de donnée MySQL pour query une liste de `messages` (dans une table du même nom)

À travers ce TP, nous verrons:
- Comment déployer une application sur Heroku, TRÈS SIMPLEMENT (back-end dans un premier temps, puis nous déploierons un front-end pour intéragir avec le back).
- Comment utiliser le système `d'addon` proposé par Heroku pour des services externes (Ici, nous utiliserons l'addon `MySQL` pour gérer notre base de donnée).


## Déployer L'API back-end

Pour créer un projet Heroku:
- utilisez la commande `heroku login` pour vous connecter
- utilisez la commande `heroku create` pour créer une nouvelle application.

Je vous invite à utiliser l'option `--help` pour obtenir des informations sur les commandes Heroku.

Lorsque l'on créer une application, Heroku va créer un `repo` git sur un de leurs serveurs, il va nous assigner un lien vers ce repo (exemple: `https://git.heroku.com/intense-cove-02127.git`), il va donc vous falloir initialiser le repo git dans [votre dossier contenant l'application, cloné en local](./app).
```bash
git init
git remote add heroku <lien donné par la commande heroku create>
```
Ce repository git est important, car il va nous permettre de déployer l'application sur un simple push. Dans un cas réel, on utiliserai deux `git remote`: Un pour notre repo Github, Gitlab Bitbucket.... et l'autre pour la remote Heroku, nous permettant de déployer notre application sur un `git push heroku master`.

Pour déployer une application, Heroku a besoins de savoir quelle commande vous souhaitez exécuter pour démarrer votre application, pour cela, on va utiliser un [Procfile](https://devcenter.heroku.com/articles/procfile), à la racine de notre repo (de notre application).

Pour cette application, nous allons utiliser le paramètre `web` (processus de type web): Serveur HTTP sur lequel on va servir du contenu (API, site statique...) pour spécifier la commande à exécuter.

Si vous regardez le [fichier package.json](./app/package.json) vous verrez que la commande pour démarrer l'application est `npm run start`.

Une fois votre Procfile configurée, vous pouvez simplement déployer l'application grâce à git:
- git add
- git commit
- git push heroku master

Et hop, heroku est notifié que vous avez modifié du code, et que vous souhaitez le déployer en production. Une fois le déploiement terminé, Heroku va vous afficher `released vx, https://lien-de-votre-app.herokuapp.com`.

Pour vérifier que votre application a été bien déployée, rendez vous sur ce lien, et vérifiez que vous obtenez bien le résultat suivant:
```json
{
  "message": "Hello World"
}
```

Si vous regardez [le code du serveur node](./app/app.js), vous verrez que le `World` ne s'affiche que lorsque la variable d'environnement `NAME` est absente. C'est donc l'occasion pour nous de voir comment configurer notre application (et ici, la variable `NAME`).

Avec Heroku, on peut utiliser la commande `heroku config`. **À L'aide de la documentation de la commande `heroku config`, ajoutez une variable `NAME=tp-heroku`**.

Ensuite, rechargez votre navigateur, vous devriez obtenir le résultat suivant:
```json
{
  "message": "Hello tp-heroku"
}
```

Nous pourrons réutiliser ces variables d'environnement (`config` dans le jargon Heroku) pour paramétrer une configuration particulière pour notre environnement.

Maintenant, ouvrez votre application sur la route `/messages`, vous devriez obtenir le message d'erreur suivant:
```json
{
  error: {
    errno: "ECONNREFUSED",
    code: "ECONNREFUSED",
    syscall: "connect",
    address: "127.0.0.1",
    port: 3306,
    fatal: true
  },
  message: "Work is not done yet..."
}
```

Comme dit dans la section au-dessus, notre API se connecte à une base de donnée `MySQL` pour récupérer une liste de messages, présent dans la base. À l'heure actuelle, nous n'avons pas déployé de base de donnée, ni configuré notre API.

Nous allons voir comment utiliser les `addons` pour créer une base de donnée `MySQL` entièrement gérée par Heroku (y compris les sauvegardes).

## Déployer une base (Addons Heroku)

Maintenant, vous allez apprendre à utiliser les addons Heroku, pour héberger des services tels que `MySQL`.

**Il est important de savoir que lorsqu'on utilise un addon sur une application, des paramètres de configurations (variables d'environnements) définies par l'addon seront injectés dans votre application**.

Par exemple, si j'utilise une addon `Redis To Go` (Redis), une nouvelle variable d'environnement `REDISTOGO_URL` sera injectée dans mon application. Tout cela est très important, car il va nous permettre de développer nos applications en utilisant des variables d'environnement pour configurer ces apps. On va ensuite pouvoir simplement porter le code d'un environnement à un autre, en changeant seulement ces variables:
- Utiliser un serveur Redis et MySQL en local avec des variables d'environnements pointant sur ces services
- Utiliser un autre serveur Redis et MySQL sur environnement de production avec les addons Heroku par exemple, qui injecteront les paramètres de config.

**Ici, vous installerez une add-on MySQL**. Vous pouvez voir [dans le code de notre API]() que nous attendons une variable d'environnement à redéfinir, vous allez devoir trouver quel paramètre est injecté par JAWS DB, puis remplacer votre configuration:
```javascript
// ligne 22, renommez la clé "TO_DEFINE" avec la nom de la variable injectée par Jaws DB
const connection = await mysql.createConnection(process.env.TO_DEFINE);
```

Je vous invite à suivre [la documentation de l'addon JAWS DB MySQL](https://devcenter.heroku.com/articles/jawsdb) (Utilisez le plan Kitefin shared, gratuit) pour ajouter cet addon.

Une fois l'addon ajoutée, et le code modifié avec le bon paramètre de configuration, n'oubliez pas de redéployer votre application (git add, commit, push).

Pour la suite de l'exercice vous aurez besoins de créer votre base de donnée [à partir du fichier database.sql dans ce repo](./database.sql), qui contient la structure de la base: une table `messages` (`id`, `content`) avec 4 enregistrements.

Je vous invite à suivre [la documentation de l'addon](https://devcenter.heroku.com/articles/jawsdb) pour apprendre comment importer des données depuis un fichier `.sql` existant (vous aurez besoins d'avoir le client MySQL installé en local).

Une fois votre base de donnée importée, vous allez pouvoir vérifier que votre API a bien accès aux messages: Allez sur votre application sur la route `/messages`, vous devriez obtenir le résultat suivant:
```json
[
  {
    id: 1,
    content: "hello"
  },
  {
    id: 2,
    content: "world"
  },
  {
    id: 3,
    content: "docker"
  },
  {
    id: 4,
    content: "compose"
  }
]
```
Bien sur, si vous modifiez la base (supprimez ou ajouter un message par exemple), vous verrez que ce résultat changera, puisque notre API est à présent connectée à la base.

## Déployer un Front-end

Ici, nous allons voir comment déployer une [application front (React)](./client).

Vous avez déjà vu cette application dans le TP précédent sur AWS.

Elle va simplement se connecter à une API configurée par une variable d'environnement, et afficher une liste de messages exposée par cette API.

Nous allons donc connecter notre application FRONT avec l'API que nous avons déployée dans la partie précédente.

Vous allez donc créer une seconde [application Heroku](https://dashboard.heroku.com/apps), soit:
- depuis [le dashboard](https://dashboard.heroku.com/apps)
- Depuis votre ligne de commande avec la CLI heroku (heroku create)

[Celon l'article officiel du blog d'Heroku](https://blog.heroku.com/deploying-react-with-zero-configuration), il n'a jamais été aussi simple de déployer une application React sur leur service.

**Côté configuration, l'application utilise les variables d'environnement suivantes**:
- `REACT_APP_BACKEND_URL`: Lien HTTP vers l'API déployée (exemple `https://mon-application.heroku.com`).

Ces variables devront être configurées via le système de `config` proposé par Heroku:
- Soit avec la CLI: `heroku config:set -a <nom-de-lapp> VARIABLE=valeur`
- Soit depuis le dashboard Heroku de votre application, onglet 
`Settings` (`Config Vars`).

Une fois ces variables configurées, il vous faudra simplement déployée votre application (git add, commit, push heroku master).

Heorku va détecter que vous utilisez une application développée avec React, et va s'occuper de construire l'application, et de la servir.

Rendez-vous sur l'adresse de votre application pour vous assurer que tout fonctionne comme prévu.

## Ajouter un service (TODO)