# Cours - Infrastructure et Hébergement

Dans ce cours, vous apprendrez les bases de l'hébergement et différentes offres sur le marché vous permettant d'atteindre la mise en production d'un produit web.

L'objectif ici sera d'exposer les étudiants à un contexte `Instances/Serveurs` où ils seront amenés à gérer l'intégralité de leur infrastructure: 
- Configuration réseau / Authentification
- installation et mise à jours des packets système
- Configuration des services
- Assurer une bonne démarche sécurité
- Packager et déployer du code sur des serveurs.

Ici, nous verrons ces fondamentaux sans mentionner les problématiques de scalabilité, car cela n'est pas nécessaire pour débuter dans la gestion d'une infrastructure permettant d'héberger une application WEB (notions + avancées). Ces différents fondamentaux seront dans un premier temps effectués manuellement (dans un contexte professionnel, nous utiliserons des outils pour automatiser la gestion de l'infrastructure et la configuration des serveurs, mais il faut d'abord connaitre la base).

Ce type d'infrastructure gérée en général par des DevOps: elle nécessite des compétences spécifiques en administration système, et parfois trop couteuse en `temps` et en `argent`.

Nous découvrirons par la suite différentes technologies et produits du marché plus appropriés à une audience de développeurs Web, comme par exemple:
- **Heroku**, pour déployer nos applications et gérer nos environnements et services très simplement, on verra même pour ajouter de l'intégration et déploiement continu
- **Firebase** pour la facilité d'utilisation, et la gamme assez large permettant le développement et déploiement d'une application de A à Z sans difficulté. Nous verrons peut être également les `Cloud Functions` proposées par Firebase, pour faire tourner des petites unités de logique sans complexité.

L'idée est de montrer aux étudiants différents outils, apportant chacun ses avantages et ses inconvénients, pour répondre à tout les besoins. Les étudiants seront en suite en mesure de sélectionner leur outils pour déployer et gérer leurs premières application à faible coùt dans le Cloud.

## Les TPs liés à ce cours:

- [Déployer une application sur des instances AWS EC2, from scratch](./tp-aws-ec2)
- [Déployer une application sur l'offre PAAS d'Heroku](./tp-heroku)

## Auteur

Antoine Masselot, DevOps Serverless et Intervenant à Hétic.
