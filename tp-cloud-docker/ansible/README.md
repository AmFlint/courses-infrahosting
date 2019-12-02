# Introduction à Ansible

## L'inventory

Ansible va appliquer une liste de tâche sur un groupe de machines. On spécifie ces machines dans un fichier appelé `inventory`, ou on va regrouper les machines par type (par exemple, serveur webs, bases de données...), [comme dans cet exemple](./inventory/inventory).

On peut également créer des `group_vars`: Variables qui seront appliqués en fonction du groupe de la machine. Pour ce faire, on va créer un dossier `inventory`, dans lequel on va créer un autre dossier `group_vars`, dans lequel on placera des fichier au formt `.yml`, contenant les variables. **Les fichiers doivent s'appeler `nom-du-groupe.yml`**. [Dans cet exemple, nous avons le groupe web](./inventory/group_vars), on créer donc un fichier `web.yml`, qui sera automatiquement appliqué uniquement sur les machines du group `web`.

## Les tasks

Les tâches sont des actions à effectuer sur un serveur distant, on peut faire une multitude d'actions, comme par exemple:
- copier des fichiers de notre ordinateur à un emplacement précis sur un serveur distant
- créer des fichiers ou des dossiers
- Installer des librairies système (avec apt par exemple)
- Influer sur l'état d'un service Linux (par exemple, redémarrer `nginx` ou arrêter `apache`).
- Utiliser des variables pour interpréter un `template`, puis uploader le contenu du fichier transformé directement sur un serveur distant
- Pleins d'autres choses...

Pour effectuer une action, on va se baser sur des `modules`, développés par Red Hat (Ansible) ou la communauté.

Par exemple dans le [role web](./roles/web/tasks/main.yml), on peut retrouver les modules `copy` (uploader des fichiers), `template` (interpréter puis copier un fichier), `apt` (installer des packets comme nginx), `service` (redémarrer nginx).

Pour en savoir plus sur les modules utilisables avec Ansible, vous pouvez vous référez à la documentation officielle.

## Les roles

Les roles contiennent une liste de tâches à effectuer, et ont un but bien précis. Dans l'exemple de ce dossier, [on utilise le role web pour déployer un serveur web nginx avec le site zelda](./roles/web).

Pour créer un role, on doit placer un dossier contenant notre role dans le dossier `roles`.

Un role contient la hiérarchie de fichier suivante:
- `tasks`: contient une liste de fichier au format yml, qui listent les tâches à effectuer dans le rôle, le fichier principale doit obligatoirement s'appelé `main.yml`
- `files`: Par convention (non obligatoire), on dépose les fichiers à uploader sur un serveur dans ce dossier, `tel quel`
- `templates`: Par convention (non obligatoire), on dépose nos fichiers de template à l'intérieur, [comme par exemple la configuration nginx du role web](./roles/web/templates/server.conf.j2). Vous remarquerez qu'on peut utiliser des variables ansible avec la syntaxe `{{ nom_de_la_variable }}` (format jinja 2)

## Les playbooks

Les playbooks regroupent une liste de rôle à effectuer dans un fichier, avec des règles précises comme par exemple sur quel groupe de machines veut-on exécuter les roles.
On parle ainsi d'exécuter un playbook.

Vous pouvez prendre exemple [sur ce playbook, pour installer nginx et déployer l'application zelda](./web.yml).

Pour exécuter un playbook, on utilise la commande suivante:
```bash
ansible-playbook -i <chemin-vers-fichier-inventory> <fichier-playbook> --key <chemin-clé-privée-ssh>

# par exemple pour utiliser le playbook web dans l'exemple
# Je spécifie le flag --user ubuntu, car sinon ansible va essayer de se connecter avec mon utiliser, "antoinemasselot"
ansible-playbook -i inventory/inventory web.yml --user ubuntu
```

Je vous invite à tester ce playbook, en prenant soin de [modifier le fichier d'inventory pour utiliser votre instance](./inventory/inventory).
