Pour récupérer le contenu de la branche `master` alors que vous êtes dans une autre branche, vous pouvez utiliser la commande suivante :

```bash
git pull origin master
```

ou, si vous voulez fusionner directement la branche `master` dans la branche courante :

```bash
git merge master
```

Et si vous voulez simplement récupérer des fichiers spécifiques depuis la branche `master`, utilisez :

```bash
git checkout master -- chemin/vers/fichier
```

Chaque commande a une utilisation spécifique :
- `git pull origin master` met à jour votre branche `master` locale et la fusionne avec la branche actuelle.
- `git merge master` fusionne le contenu de la branche `master` dans la branche courante.
- `git checkout master -- chemin/vers/fichier` permet de récupérer un fichier spécifique de la branche `master` sans fusionner toute la branche.