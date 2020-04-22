# Notes Infrastructure avec retours

- Pierre Turnbull (G5): 15
- Alexandre Delaloy (G12): 12,5
- Quentin Aimé (G1): 16
- Rainald Durand (G1): 16
- Amaury (G4): 16,5
- PHAM Than-Son Vincent (G7): 17
- AMAR BENSABER Hichem (G9): 16,5
- Corner Guillaume (G11): 15,5
- Paule Hermann (G2): 16,5
- Kevin Manssat (G10): 18
- Paartheepan Raveenthiran (G10): 18
- Keny Zachelin (G3): 16,5

# Equipe 5 - Pierre Turnbull

## Docker

- Veillez à n'installer que les dépendances de production dans vos Dockerfile (`npm insall --production`)
- Multi Stage build: Très bien
- Je ne vois pas de configuration concernant le Front: Comment indiquez-vous dynamique l'URL de l'API ?
- Build et push de l'image dans la pipeline de déploiement continu: Bien

4.5/5

## Terraform

- Je vous invite à vous pencher sur l'utilisation de variables Terraform (pour la clé SSH par exemple), afin de m'éviter de modifier vos fichiers à la main.
- Pensez à spécifier une version du provider AWS afin de vous éviter des breaking changes lors de futurs mises à jours automatiques

4/5

## Ansible

Pas de code Ansible, cependant je vois que vous avez fait le job avec des scripts Shell et des mécanismes d'intégration/déploiement Continu (TravisCI), bravo pour ça.

- Scripts shells pour installer l'environnement (docker, docker-compose) et remettre à jours le projet docker-compose: bravo.
- Utilisation d'outils de déploiement continu comme TravisCI: Bien

3.5/5

## Reproduction

- Terraform: OK 2.5/2.5
- Ansible/Scripts: Avec l'absence de documentation, difficile de déployer votre application (Quels secrets doit-je créer, quels scripts importer et dans quel ordre les exécuter ?): 0.5/2.5

3/5

Total: 15/20

---

# Equipe 1

# INFRA

## FRONT - Quentin Aimé

### Docker
- Front:
  - 2 steps build (build le code dans node, puis COPY dans nginx)
  - Je ne vois pas la configuration du endpoint de l'API, comment déploie-t-on différents environnements (dev, staging, prod) sans pouvoir configurer dynamique le endpoint back-end ?

4/5

### Terraform

- Variable Path SSH KEY Avec un défaut (OK)
- egress: 0-65535 -> 0-0 autorise tous les ports, protocol -1 => tous les protocols
- PLus intéressant de pin la version du provider AWS, autrement, on risque de prendre des breaking changes

4.5/5

### Ansible

- Inventory Statique, ansible_user=ubuntu, je recommande de ne pas spécifier le user à cet endroit mais plutôt en option de la commande `ansible-playbook --user ubuntu`
- Pour la suite, je vous recommande:
  - D'utiliser des roles ansible-galaxy open-source (geerlingguy.docker par exemple)
  - Des tags sur vos `tasks`, afin de pouvoir ré-exécuter uniquement certaines tâches d'un rôle (par exemple, lors d'une mise à jours de l'application, on voudra juste lancer les tâches liées au redéploiement et pas toute la configuration)
  - Vous n'utilisez aucune variable de configuration, ce qui est dommage pour le redéploiement des images docker, qui nous force à toujours publier des images `latest`, et écraser la version précédentes. Ce qui peut poser problème lorsqu'on souhaite "versionner" les déploiement pour facilement rollback sur un version antérieure.
  - Vous déployez un container docker seul, je recommande l'utilisation de docker-compose en production, car permet de scale plus efficacement les déploiement (en ajoutant de la configuration par exemple, et vous permet de relancer/arrêter/modifier la configuration directement lorsqu'un incident se produit).

Cela dit, vous avez compris la base d'Ansible, il faut maintenant entrer dans une utilisation + complète de l'outils avec:
- variables de configuration
- tags pour relancer des tâches par groupes
- Ansible-Vault pour stocker vos secrets
- Inventaire Dynamique pour générer vos inventaires sans avoir besoins de renseigner les IPs de vos hôtes à la main.

3/5

Reproduction:
Terraform OK 2.5/2.5
Ansible les hosts visés par les playbooks !== de ceux précisés dans l'inventory (`webserver` !== `webservers`), autrement tout fonctionne. 1.5/2.5
4/5

Total Front: 15.5

## Back - Rainald

### Docker
- Dockerfile OK
- Docker-compose utilise networks et volumes OK
- Pas de configuration injectée via variables d'environnement ou fichier dans le docker-compose, rend le déploiement trop statique (trop peu flexible).
4/5

### Ansible:
- Il est préférable d'utiliser des modules Ansible plutôt que des commandes Shells quand ils sont disponibles. Par exemple: `docker_service ou docker_compose` à la place de `shell: docker-compose up -d`.
- Ajouter de la configuration via Ansible variables (`group_vars`, `ansible_vault`).
- Utiliser le système de template pour déployer la configuration de l'API (docker-compose.yml)
- Nommage des playbooks/roles (docker_compose.yml)
- Utiliser des playbooks opensource (geerlinguy/ansible-role-docker)
- Utiliser des `handlers` (redémarrer la stack docker-compose si la configuration change)
- Utiliser les variables d'environnement dans le code côté API pour la configuration dynamique

Je ne vois pas de variables d'environnement côté Ansible et le docker-compose de votre application en production.
Faites attention à utiliser des variables Ansible avec le module `template` pour gérer la configuration de votre fichier docker-compose. Vous avez utilisé copy, avec des variables écrites en dur dans le fichier, ce qui entrave l'agilité/configuration du projet

3/5

### Terraform

- (-) Utiliser des variables terraform (pour la clé SSH et l'IP pour le SSH par exemple)
- egress: 0-65535 -> 0-0 autorise tous les ports, protocol -1 => tous les protocols
- (+) La version du provider AWS est mentionnée, ça permet d'éviter des breaking changes sur un même script Terraform

4/5

### Reproduction

Terraform OK 2.5/2.5
Ansible OK 2.5/2.5
5/5

16/20

---

# Equipe 4 - Amaury

## Docker

- Back OK
- Front OK, 2 stages build. Attention cependant, vous faites un `COPY serverfront.conf`, je ne parviens pas à trouver ce fichier, votre build va fail.
- Je ne vois pas de configuration dans votre fichier docker-compose, j'imagine que la configuration (connexion à la base de données) est écrite en dur dans le code. Attention à ce point là, ça casse la flexibilité des déploiements.

3.5/5


## Terraform
- SSH clé, nom de la variable + précis/descriptif. Variable moins contraignante (l'utilisateur passe le chemin complet)
- Penser à définir une version pour le provider AWS, autrement on peut se retrouver avec des breaking changes, et le script ne fonctionne plus

4.5/5

## Ansible
- + utilisation des tags
- + build en host local (possibilité de réutiliser le playbook dans la CI/CD)
- + Inventaire Dynamique
- Ansible User en ligne de commande
- Réimporter les variables (all.yml) dans les group_vars de chacun des roles
- déplacer updatedockerfile dans "application" (et renommer deploydockercompose), ajouter des tags, et `delegate_to` pour builder en local
- Scoper les variables par Role
- Les templates sont dans un fichier templates et non files
- top-level become
- ~/site => Chemin absolu potentiellement variabilisé
- Handler pour relancer les services sur docker-compose

5/5

## Reproduction

Terraform OK 2.5/2.5
Ansible:
- Il manque des modules Python à installer dans le playbook pour utiliser les modules docker et docker_compose:
  - Installer `python3-pip` avec le module apt
  - Installer les libraries `docker` et `docker-compose` avec le module pip
1/2.5

3.5/5

Total: 16.5/20

---

# Équipe 7 - PHAM Than-Son Vincent

## Docker

- Bonne utilisation de la Dockerfile de l'API (copy package.json ET yarn.lock), yarn install --production pour minimiser les dépendances inutiles
- Utilisation de Variable de configuration pour la Dockerfile Front-End, permet une bonne flexibilité pour déployer le front sur différents environnements
- Ré-écriture de la configuration NGINX pour rediriger sur le fichier index.html: OK
- Utilisation de variables d'environnements pour la configuration des applications dans le fichier docker-compose: Très bien
- Je ne vois pas de docker-networks, comment les containers back, prisma et mongo communiquent-ils ensemble ?

4/5

## Terraform

- Penser à spécifier une version pour le provider AWS, afin d'éviter les breaking changes entre deux nouvelles versions
- Le chemin vers la clé SSH est trop contraignant, vous devriez utiliser une variable qui pointe sur un fichier pour plus de flexibilité
- Security Groups OK

4.5/5

## Ansible

- Utilisation de variables pour la configuration des applications (docker-compose) ainsi que du module template
- Playbook complets, allant de l'installation des composants, build des images docker et exécution du projet avec docker-compose: bravo
- (-) Je vous conseille de vous renseigner sur le système de `Tags` Ansible, afin de pouvoir exécuter uniquement certaines tâches d'un rôle/playbook (par exemple build+push est un groupe de tâche à part entière, redéployer le docker-compose avec les nouvelles images constitue un autre groupe de tâches).
- Utilisation des `defaults`
- Utilisation de roles open-source comme geerlingguy.docker: très bien
- Pour la suite, je vous recommande de regarer le système d'inventaire dynamique proposé par Ansible

5/5

## Reproduction

Terraform OK 2.5/2.5
Ansible Le build des images ne passe pas. "Failed to Compile" côté back-office 1/2.5
3.5/5

17/20

---

# Équipe 9 - Hichem

## Docker
- La configuration de l'application est injectée par variable d'environnement dans le docker-compose: Très bien, ça permet une flexibilité lors du déploiement sur différents environnements
- Bonne utilisation des docker-networks: Uniquement API, Mongo-Express et Mongo sont mentionnés sur le réseau
- Dockerfile BACK OK
- Je ne vois pas le paramètre de configuration `API_URL` dans votre Dockerfile, alors que vous l'utilisez dans votre code front-end ?

4/5

## Terraform
- Chemin de clé SSH moins contraignant (~/.ssh => variable)
- Nom de la variable + précis (ssh_key_file, ssh_key_name, ssh_key_content)
- Version du provider

4.5/5

## Ansible
- (+) utilise le dynamique inventory
- (+) Requirements.txt pour les modules python nécessaires pour run le Playbook Ansible
- (+) Configuration de l'application
- (+) Utilisation des defaults pour configurer des variables par défaut.
- Utiliser Ansible-Vault pour les secrets
- Passer la clé privé en ligne de commande (ansible-playbook --key <chemin-clé>)
- Passer user en ligne de commande
- Installer python ne devrait pas être fait dans le roles pour docker
- Utiliser handler pour restart docker-compose plutôt que down puis up (la stack ne rebootera pas si erreur de configuration == downtime)
- Préfixer les variables par le nom du role
- Je vous recommande de regarder le système de Tags pour lancer des groupes de tâches.
- Je vous recommande de vous renseigner sur l'utilisation de roles open source avec Ansible Galaxy (installation de docker par exemple)

4.5/5

## Reproduction

- Terraform OK 2.5/2.5
- Ansible:
  - pip n'est pas installé sur la machine par défaut, je dois l'installer en plus sinon ça plante`
1/2.5

3.5/5

Total: 16.5/20

---

# Équipe 11

## Docker

- Utilisation de docker-compose en développement et production
- Configuration via variables d'environnements
- Pas besoins de network FRONT dans votre projet: Le front est téléchargé dans le navigateur de votre utilisateur, les requêtes se font depuis son ordinateur et non pas depuis le container front
- Bonne utilisation du network database
- Je ne trouve pas les Dockerfiles liées à vos projets, je ne peux donc pas les évaluer

3/5

## Terraform

- Description + Default sur les variables
- SSH Key name trop contraignant (l'utilisateur doit pouvoir passer son chemin complet)
- AWS Profile (default par défaut)
- Pensez à spécifier une version à votre provider AWS, afin d'éviter les breaking changes lors de futurs mises à jours

4.5/5

## Ansible

- + Pre-install pour configurer python/ansible
- + Inventory Dynanique (Faire attention aux tags, hosts ec2 === toutes les machines)
- + Requirements.txt pour les dépendances Python pour lancer les playbooks
- + Système de variables, configuration via Ansible (Possibilité d'utiliser Ansible Vault)
- Ansible_user en flag de CLI au lieu de en dur dans le playbook
- Scoper les variables par roles (préfixées par le nom du role)
- Oublie de variabiliser les images dans le docker-compose compose de l'application
- Ajouter des tags pour l'évolution du playbook et des roles
- Attention à utiliser become uniquement sur les tâches dont c'est nécessaire
- Vous devriez variabiliser le nom et tag des images Docker de votre application, je suis obligé de modifier votre template pour déployer une image venant de ma registry
- Vous devriez renommer votre role `python` de façon plus générique, en réalité il sert à installer des packages nécessaires à faire tourner votre application. Je l'aurai nommé `packages` ou `common` par exemple.

5/5

## Reproduction

- Terraform: OK 2.5/2.5
- Ansible: 0.5/2.5
  - J'ai été contraint de modifier le fichier `inventory/ec2.ini` pour empêcher l'utilisation d'Elasticache (mon token n'est scopé qu'aux instances EC2 pour limiter les permissions).
  - Ne trouvant pas les Dockerfiles du projet, je ne peux pas build et push les images dans ma registry, le playbook crash car je ne peux pas pull vos images.

3/5

Total: 15.5/20

---

# Equipe 8 - pas de note infra

---

# Equipe 6 - pas de note infra

---

# Equipe 12 - Alexandre Delaloy

## Docker

- `FROM` image de base taguée, très bien ça permet d'éviter les breakings changes lors des nouvelles mises à jours de latest ou alpine (qui est en réalité la latest sur la distribution alpine)
- 2 étapes de build pour le client, très bien, ça permet de conserver des petites images avec seulement les logiciels nécessaires.
- Je ne vois pas la configuration du Endpoint back-end dans le dockerfile Client, j'imagine qu'elle est inscrite en dur dans le code ?

Plusieurs problèmes sur le docker-compose à prendre en compte les prochaines fois:
- Pas de container de base de données, au vu du nom de l'image API, j'en déduis que vous utilisez un outils comme json-server pour vous générez une API Rest à partir d'un fichier JSON ?
- Pas de configuration dynamique de votre API (aucune variable d'environnement ni de volume de fichiers de configuration): Comment configurez vous l'accés à la base de données ?
- L'image est inscrite en dur (blyndusk/phrh-client et blyndusk/phrh-fake-server): Je ne peux pas utiliser votre docker-compose.yml en l'état si je souhaite stocker vos images dans une autre registry sous un autre nom. Je vous recommande d'utiliser des variables (ici, vous avez utiliser des variables d'environnements pour les tags des images: TRÈS BIEN, vous devriez faire de même pour les noms des images).

3/5

## Terraform

- Pensez à pin une version du provider `aws` pour éviter les breakings changes de vos scripts lors des prochaines mises à jours du provider AWS
- Utilisez des variables pour des données comme le contenu de la clé SSH Publique (et potentiellement le type de l'instance ainsi que l'image)
- Security Groups OK

4.5/5

## Ansible

Je ne trouve rien concernant Ansible, mais je vais appliquer une notation au script `bin/entrypoint.sh`.

- Scripts pour build+push les images docker de front et back: Bien
- Utilisation du script Shell avec SSH + SCP pour mettre à jours le fichier docker-compose (avec nouveau tag d'images) et les variables d'environnement (.env), il faudra penser à apporter + de flexibilité à la configuration (ici vous vous contentez de dupliquer le .env d'exemple et l'uploader sur le serveur), mais c'est un bon début.

Vous avez utilisé les Github Actions pour effectuer des tâches d'intégration continue sur votre projet: bien, dommage que vous n'ayez pas ajouté l'étape de déploiement (un simple scp de votre projet buildé aurai fait l'affaire pour le front). Ou encore, vous pourriez utiliser le commandes shells de votre script `bin/entrypoint.sh` pour build+push les images docker + une commande `ssh ubuntu@<IP> -t docker-compose up --force-recreate` pour redéployer vos nouvelles images sur un Merge par exemple.

**J'aurai aimé un script Shell pour l'installation de docker, c'est 4 commandes et ça aurai fait l'affaire**.s

2/5

## Reproduction

- Terraform: OK 2.5/2.5
- Ansible/scripts: 0.5/2.5
  - Utilisez des variables (arguments au script, ou variable d'environnement) pour la clé privée, voir même ne l'incluez pas, l'utilisateur (ici, moi) peut ajouter sa clé SSH à sa session avec la commande:
    ```bash
    ssh-add ~/.ssh/ma_cle
    # cela me permet de me connecter à mon instance sans devoir préciser le chemin vers ma clé à chaque fois
    # ci-dessous, plus besoins du flag "-i"
    ssh ubuntu@"$INSTANCE_IP"
    ```
  - Le script reste coincé sur les différentes étapes

3/5

Total: 12.5

---

# Equipe 2 - Paule

## Docker

- Application configurable dans le docker-compose via variables d'environnement
- Multi stages build pour le front: Très bien, mais veuillez nommer votre stage de build le `--from=0` peut porter confusion
- Je ne vois pas de variable de configuration pour l'URL de l'API dans le Dockerfile du front, comment configurez vous le nom de domaine/IP de l'API ?
- Dockerfile back: OK

4/5

## Terraform

- Il faut spécifier une version du provider AWS, afin d'évtier les breaking changes dans les nouvelles version du provider lors de futurs déploiements
- Vous devriez utiliser une variable pour la clé SSH Publique à déployer, afin de m'éviter de modifier votre fichier
- Au niveau des règles de Security Groups, vous pourriez utiliser une variable pour l'IP (plutôt que de mettre Hétic en dur)
- En général, on souhaite exposer nos Applications (API back-end, Front-end) sur tout l'Internet, avec un ingress '0.0.0.0', ici vous avez uniquement autorisé Hétic

4/5

## Ansible

- (+) Utilisation de roles opensource avec Ansible Galaxy
- (+) Configuration par variables (Inventory/Group vars)
- (+) Variables scopées par roles
- (+) Role Common pour l'installation de tous les packages etc
- (+) Utilisation de variables defaults
- Je vous invite à vous documenter sur les inventaires dynamique
- Pourquoi installer `nginx` dans votre role `common` ? Évitez d'installer des logiciels qui ne vous servent pas

5/5

## Reproduction

- Terraform: OK 2.5/2.5
- Ansible: 1/2.5
  - Erreur de syntaxe dans votre fichier `roles/webapp/tasks/main.yml` au niveau de l'indentation

3.5/5

16.5/20

---

# Équipe 10

## Docker

- Configuration des containers via variables d'environnements dans le docker-compose: Bien
- Utilisation de networks pour scoper les communications (Serveur - DB)
- Multi-stages build pour builder le code, et pour le servir: Très bien
- Dockerfile BACK OK

5/5


## Terraform

- Penser à spécifier la version du provider AWS pour éviter les breaking changes lors des futurs mises à jours
- Préférez l'utilisation de variables pour faciliter la configuration (SSH KEY) et éviter à vos utilisateurs de devoir modifier du code

4.5/5

## Ansible

- Configuration avec les variables ansible (scopées par role): Très bien
- Utilisation des templates: Bien
- Veuillez tout de même à utiliser des variables pour le nom de l'image, ainsi que le tag, ça me permettrai de déployer votre application depuis ma propre registry, ainsi que de pouvoir versionner l'application en production, et rollback sur une vesion précédente en cas de bugs.
- Je vous invite à vous renseigner sur le concept de tags Ansible, pour relancer un playbook mais n'appliquer que des tâches définies par un tag (pour redéployer sans tout ré-installer par exemple)
- Bonne séparation des tâches par rôle (Docker, pip et site), je recommande de renommer pip en `packages` ou `common`.
- Je vous invite à regarder les inventaires dynamiques pour gérer un parc de machine à plus grande échelle.
- Utilisation d'Ansible-Vault pour stocker/appliquer les secrets: Très bien

5/5

## Reproduction

- Terraform OK 2.5/2.5
- Ansible: 1/2.5 (Difficile à reproduire, vos images sont dans une registry privée, avec des certificats SSL invalides)
  - problème avec pip (utilisation de pip3 par défault, vous installez python2, le playbook crash avec ma version d'ansible (la plus à jours))

3/5

Total: 18

---

# Equipe 3 - Keny Zachelin

## Docker

- Configuration de l'application par variables d'environnements dans le fichier docker-compose: Très bien
- Build des images de production avec des build arguments pour le front: Très bien.
- Utilisation des networks: bien, cependant vous devriez limité l'accès à la base de données. Le container front n'a pas besoins d'accéder au networks_database car il ne tape pas dessus directement, c'est une faille de sécurité.
- Vous devriez supprimer les statements inutiles dans vos Dockerfiles (ARG foo1, ARG foo2, RUN echo "$foo1"...)
- Multi-stage build pour l'image de Front: Très bien
- Vous pourriez limité l'installation des dépendances NPM à la production (`npm install --production`)

4/5

## Terraform

- Il est préférable de fixer une version au provider AWS pour éviter les breaking changes lors des futurs mises à jours
- Utilisation de variables pour configurer le script: Très bien
- Exécution des scripts dans Terraform pour déployer votre infrastructure: Bien, vous permet de livrer dans les temps et dans des bonnes conditions même si vous ne respectez pas totalement les consignes.

5/5

## Ansible NA

Pas de code Ansible.

Cependant, je vous félicite pour vous être débrouillés avec les moyens du bord (script Shells exécutés dans Terraform pour configurer vos serveurs).

Au niveau de la configuration de votre application en production, vous devriez utiliser un système de variable d'environnement dans votre docker-compose, car votre configuration est écrite en claire dans votre repository Git.
Je vous recommande de vous renseigner sur les services Linux avec systemD, pour configurer un service avec votre configuration en environnement, qui lance/arrête votre projet avec docker-compose.

2.5/5

## Reproduction

- Terraform OK 2.5/2.5
- Ansible/Scripts: Très bien, rien à redire, tout fonctionne en une seule commande: 2.5/2.5

5/5

16.5/20
