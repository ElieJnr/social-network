#### Fonctionnel

###### L'exigence pour les paquets autorisés a-t-elle été respectée?

##### Ouvrir le projet

###### En examinant le système de fichiers du backend, avez-vous trouvé une structure bien organisée, similaire à l'exemple fourni dans le sujet, avec une séparation claire des paquets et des dossiers de migration?

###### Le système de fichiers pour le frontend est-il bien organisé?

#### Retour

###### Le backend inclut-il une séparation claire des responsabilités entre ses trois parties principales - Serveur, Application et Base de données?

###### Y a-t-il un serveur qui reçoit efficacement les demandes entrantes et sert de point d'entrée pour toutes les demandes à l'application?

###### L'application (App) qui s'exécute sur le serveur écoute-t-elle efficacement les demandes, récupère-t-elle les informations de la base de données et envoie-t-elle des réponses?

###### La logique de base du réseau social est-elle implémentée dans le composant App, y compris la logique de traitement de différents types de requêtes basées sur HTTP ou d'autres protocoles?

#### Base de données

###### SQLite est-il utilisé dans le projet comme base de données?

###### Les clients peuvent-ils demander des informations stockées dans la base de données et peuvent-ils soumettre des données à y ajouter sans rencontrer d'erreurs ou de problèmes?

###### L'application implémente-t-elle un système de migration?

###### Ce système de fichiers de migration est-il bien organisé? (comme l'exemple du sujet)

##### Démarrez l'application du réseau social, puis entrez dans la base de données à l'aide de la commande #"sqlite3 <database_name.db>.

###### Les migrations sont-elles appliquées par le système de migration?

#### Authentification

###### L'application implémente-t-elle des sessions pour l'authentification des utilisateurs?

###### Les éléments de formulaire corrects sont-ils utilisés dans l'enregistrement? (Email, Mot de passe, Prénom, Nom, Date de naissance, Avatar/Image (Facultatif), Surnom (Facultatif), À propos de moi (Facultatif))

##### Essayez d'enregistrer un utilisateur.

###### Lors de l'enregistrement, lors de la tentative d'enregistrement d'un utilisateur, l'application a-t-elle correctement enregistré l'utilisateur enregistré dans la base de données sans aucune erreur?

##### Essayez de vous connecter avec l'utilisateur que vous venez d'enregistrer.

###### Lorsque vous tentez de vous connecter avec l'utilisateur que vous venez d'enregistrer, le processus de connexion a-t-il fonctionné sans aucun problème?

##### Essayez de vous connecter avec l'utilisateur que vous avez créé, mais avec un mot de passe ou un email incorrect.

###### L'application a-t-elle correctement détecté et répondu aux informations de connexion incorrectes?

##### Essayez d'enregistrer le même utilisateur que vous avez déjà enregistré.

###### L'application a-t-elle détecté si l'e-mail/utilisateur est déjà présent dans la base de données?

##### Ouvrez deux navigateurs (ex: Chrome et Firefox), connectez-vous à l'un et actualisez les autres navigateurs.

###### Pouvez-vous confirmer que le navigateur non connecté reste non enregistré?

##### En utilisant les deux navigateurs, connectez-vous avec des utilisateurs différents dans chacun. Puis actualisez les deux navigateurs.

###### Pouvez-vous confirmer que les deux navigateurs continuent avec les bons utilisateurs?

#### Abonnés

##### Essayez de suivre un utilisateur privé.

###### Pouvez-vous envoyer une demande suivante à l'utilisateur privé?

##### Essayez de suivre un utilisateur public.

###### Êtes-vous en mesure de suivre l'utilisateur public sans avoir besoin d'envoyer une demande suivante?

##### Ouvrez deux navigateurs (ex: Chrome et Firefox), connectez-vous en tant que deux utilisateurs privés différents et avec l'un d'eux essayez de suivre l'autre.

###### L'utilisateur qui a reçu la demande est-il en mesure d'accepter ou de refuser la demande suivante?

##### Après avoir suivi un autre utilisateur avec succès essayer de ne pas le suivre.

###### Avez-vous pu le faire?

##### Profil

##### Essayez d'ouvrir votre propre profil.

###### Le profil affiche-t-il toutes les informations demandées dans le formulaire d'inscription, en dehors du mot de passe?

##### Essayez d'ouvrir votre propre profil.

###### Le profil affiche-t-il chaque message créé par l'utilisateur?

##### Essayez d'ouvrir votre propre profil.

###### Le profil affiche-t-il les utilisateurs que vous suivez et ceux qui vous suivent?

##### Essayez d'ouvrir votre propre profil.

###### Êtes-vous capable de changer entre profil privé et profil public?

##### Ouvrez deux navigateurs et connectez-vous avec différents utilisateurs sur eux, avec l'un des utilisateurs ayant un profil privé et suivez avec succès cet utilisateur.

###### Êtes-vous en mesure de voir un profil privé utilisateur suivi?

##### En utilisant les deux navigateurs avec les mêmes utilisateurs, avec l'un des utilisateurs ayant un profil privé et assurez-vous de ne pas le suivre.

###### Avez-vous été empêché de voir un profil privé utilisateur non suivi?

##### En utilisant les deux navigateurs avec les utilisateurs, avec l'un des utilisateurs ayant un profil public et assurez-vous de ne pas le suivre.

###### Êtes-vous capable de voir un profil public d'utilisateur non suivi?

##### En utilisant les deux navigateurs avec les utilisateurs, avec l'un des utilisateurs ayant un profil public et suivre avec succès cet utilisateur.

###### Êtes-vous en mesure de voir un profil public d'utilisateur suivi?

#### Messages

###### Êtes-vous en mesure de créer un post et de commenter les messages déjà existants après vous être connecté?

##### Essayez de créer un post.

###### Êtes-vous en mesure d'inclure une image (JPG ou PNG) ou un GIF dessus?

##### Essayez de créer un commentaire.

###### Êtes-vous en mesure d'inclure une image (JPG ou PNG) ou un GIF dessus?

##### Essayez de créer un post.

###### Pouvez-vous spécifier le type de confidentialité du message (privé, public, presque privé)?

###### Si vous choisissez l'option de confidentialité presque privée, pouvez-vous spécifier les utilisateurs autorisés à voir le message?

##### Groupes

##### Essayez de créer un groupe.

###### Avez-vous pu inviter un de vos abonnés à rejoindre le groupe?

##### Ouvrez deux navigateurs, connectez-vous avec différents utilisateurs sur chaque navigateur, suivez-vous et avec l'un des utilisateurs créez un groupe et invitez l'autre utilisateur.

###### L'autre utilisateur a-t-il reçu une invitation de groupe qu'il/elle peut refuser/accepter?

##### En utilisant les mêmes navigateurs et les mêmes utilisateurs, avec l'un des utilisateurs créer un groupe et avec l'autre essayer de faire une demande d'entrée de groupe.

###### Le propriétaire du groupe a-t-il reçu une demande qu'il/elle peut refuser/accepter?

###### Un utilisateur peut-il faire des invitations de groupe, après avoir fait partie du groupe (étant l'utilisateur différent du créateur du groupe)?

###### Un utilisateur peut-il faire une demande de saisie de groupe (une demande de saisie de groupe)?

###### Après avoir fait partie d'un groupe, l'utilisateur peut-il créer des messages et commenter des messages déjà créés?

##### Essayez de créer un événement dans un groupe.

###### On vous a demandé un titre, une description, un jour/heure et au moins deux options (aller, ne pas aller)?

##### En utilisant les mêmes navigateurs et les mêmes utilisateurs, une fois que les deux font partie du même groupe, créez un événement avec l'un d'eux.

###### L'autre utilisateur est-il capable de voir l'événement et de voter dans quelle option il veut?

#### Chat

##### Essayez d'ouvrir deux navigateurs (ex: Chrome et Firefox), connectez-vous avec différents utilisateurs dans chacun d'eux. Ensuite, avec un des utilisateurs, essayez d'envoyer un message privé à l'autre utilisateur.

###### L'autre utilisateur a-t-il reçu le message en temps réel?

##### Essayez d'ouvrir deux navigateurs (ex: Chrome et Firefox), connectez-vous avec différents utilisateurs qui ne se suivent pas du tout. Ensuite, avec un des utilisateurs, essayez d'envoyer un message privé à l'autre utilisateur.

###### Pouvez-vous confirmer qu'il n'a pas été possible de créer un chat entre ces deux utilisateurs?

##### En utilisant les deux navigateurs avec les utilisateurs, commencez une conversation entre les deux.

###### Le chat entre les utilisateurs s'est-il bien passé? (ne pas planter le serveur)

##### Essayez d'ouvrir trois navigateurs (ex: Chrome et Firefox ou un navigateur privé), connectez-vous avec différents utilisateurs dans chacun d'eux. Ensuite, avec un des utilisateurs, essayez d'envoyer un message privé à l'un des autres utilisateurs.

###### Seul l'utilisateur ciblé a-t-il reçu le message?

##### En utilisant les trois navigateurs avec les utilisateurs, entrez avec chaque utilisateur un groupe commun. Ensuite, commencez à envoyer des messages à la salle de discussion commune en utilisant l'un des utilisateurs.

###### Tous les utilisateurs communs au groupe ont-ils reçu le message en temps réel?

##### En utilisant les trois navigateurs avec les utilisateurs, continuez à discuter entre les utilisateurs du groupe.

###### Le chat entre les utilisateurs s'est-il bien passé? (ne pas planter le serveur)

###### Pouvez-vous confirmer qu'il est possible d'envoyer des emojis par chat à d'autres utilisateurs?

#### Notifications

###### Pouvez-vous vérifier les notifications sur chaque page du projet?

##### Ouvrez deux navigateurs, connectez-vous en tant que deux utilisateurs privés différents et avec l'un d'eux essayez de suivre l'autre.

###### L'autre utilisateur a-t-il reçu une notification concernant la demande suivante?

##### Ouvrez deux navigateurs, connectez-vous avec différents utilisateurs sur chaque navigateur, suivez-vous et avec l'un des utilisateurs créez un groupe et invitez l'autre utilisateur.

###### L'utilisateur invité a-t-il reçu une notification concernant la demande d'invitation de groupe?

##### Ouvrez deux navigateurs, connectez-vous avec des utilisateurs différents sur chaque navigateur, créez un groupe avec l'un d'eux et avec l'autre envoyez une demande d'entrée de groupe.

###### L'autre utilisateur a-t-il reçu une notification concernant la demande d'entrée de groupe?

##### Ouvrez deux navigateurs, connectez-vous avec différents utilisateurs sur chaque navigateur, faites partie du même groupe avec les deux utilisateurs et avec l'un des utilisateurs créez un événement.

###### L'autre utilisateur a-t-il reçu une notification concernant la création de l'événement?

#### Docker

##### Essayez d'exécuter l'application et d'utiliser la commande docker #"docker ps -a"

###### Pouvez-vous confirmer qu'il y a deux conteneurs (backend et frontend), et que les deux conteneurs ont des tailles non nulles indiquant qu'ils ne sont pas vides?

##### Essayez d'accéder à l'application de réseau social via votre navigateur Web.

###### Avez-vous pu accéder à l'application de réseau social via votre navigateur Web après avoir exécuté les conteneurs docker, confirmant que les conteneurs fonctionnent et servent l'application comme prévu?

#### Bonus

###### +Pouvez-vous vous connecter en utilisant Github ou un autre type d'Authenticator externe (standard ouvert pour la délégation d'accès)?

###### +L'étudiant a-t-il créé une migration pour remplir la base de données?

###### +Si vous ne suivez pas un utilisateur, obtenez-vous une fenêtre contextuelle de confirmation?

###### +Si vous changez votre profil de public en privé (ou vice versa), obtenez-vous une fenêtre contextuelle de confirmation?

###### +Y a-t-il une autre notification en dehors de celles explicites sur le sujet?

###### +Le projet présente-t-il un script pour construire les images et les conteneurs? (en utilisant un script pour simplifier la construction)

###### +Pensez-vous en général que ce projet est bien fait?