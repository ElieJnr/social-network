# Social Network

Une plateforme de réseau social moderne construite avec **Go** pour le backend et **Next.js** pour le frontend. Cette application permet aux utilisateurs de se connecter, de partager du contenu, de former des groupes, d'organiser des événements et de communiquer en temps réel.

## 🌟 Fonctionnalités

### Authentification & Profils
- ✅ Inscription et connexion sécurisées
- ✅ Gestion de sessions
- ✅ Profils utilisateur (publics et privés)
- ✅ Édition du profil (email, prénom, nom, date de naissance, avatar, surnom, bio)
- ✅ Système de suivi (followers/following)
- ✅ Demandes de suivi pour utilisateurs privés

### Contenus & Interactions
- ✅ Création, édition et suppression de posts
- ✅ Système de commentaires sur les posts
- ✅ Likes et dislikes
- ✅ Notifications en temps réel
- ✅ Chat privé entre utilisateurs
- ✅ WebSocket pour communication instantanée

### Groupes & Événements
- ✅ Création et gestion de groupes
- ✅ Ajout de membres aux groupes
- ✅ Création d'événements dans les groupes
- ✅ Suggestions de groupes
- ✅ Affichage des groupes publics

### Communication
- ✅ Messages instantanés (WebSocket)
- ✅ Notifications push
- ✅ Système de notification pour les interactions

---

## 📋 Prerequis

- **Docker** & **Docker Compose** (pour utiliser les conteneurs)
- **Go 1.22+** (pour développement local du backend)
- **Node.js 18+** & **pnpm** (pour développement local du frontend)
- **SQLite 3** (inclus dans le Go)

---

## 🚀 Installation & Démarrage

### Option 1: Avec Docker Compose (Recommandé)

```bash
# Cloner le repository
git clone <repository-url>
cd social-network

# Démarrer les services
docker-compose up -d

# Accéder à l'application
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
```

### Option 2: Avec le script shell

```bash
bash run_docker.sh
```

### Option 3: Développement local

#### Backend (Go)

```bash
cd backend

# Installer les dépendances
go mod download

# Démarrer le serveur
go run main.go
# Le serveur tournera sur http://localhost:8080
```

#### Frontend (Next.js)

```bash
cd frontend

# Installer les dépendances
pnpm install

# Démarrer le serveur de développement
pnpm dev
# L'application tournera sur http://localhost:3000
```

---

## 📁 Structure du Projet

### Backend (`/backend`)

```
backend/
├── main.go                    # Point d'entrée de l'application
├── go.mod                     # Dépendances Go
├── config/                    # Configuration des services
├── handlers/                  # Contrôleurs HTTP
│   ├── authentication/        # Login, registration, logout
│   ├── posts, comments/       # Gestion des posts
│   ├── groups, events/        # Groupes et événements
│   ├── followers/             # Système de suivi
│   ├── notifications/         # Notifications
│   └── websocket.go           # WebSocket
├── middlewares/               # CORS, authentification
├── pkg/
│   ├── db/                    # Base de données SQLite
│   │   ├── migrations/        # Scripts de migrations
│   │   └── sqlite/            # Driver SQLite
│   ├── models/                # Modèles de données
│   └── services/              # Logique métier
└── routes/                    # Définition des routes API
```

### Frontend (`/frontend`)

```
frontend/
├── app/                       # Structure Next.js 13+ App Router
│   ├── layout.js              # Layout principal
│   ├── page.js                # Page d'accueil
│   ├── actions/               # Server actions
│   ├── auth/                  # Pages d'authentification
│   ├── groups/                # Pages des groupes
│   ├── profil/                # Pages de profil
│   └── store/                 # État global (Zustand)
├── components/                # Composants React réutilisables
│   ├── PostCard.jsx
│   ├── CommentCard.jsx
│   ├── Nav.jsx, NavBar.jsx
│   └── ui/                    # Composants UI (Radix UI)
├── lib/                       # Utilitaires
├── public/                    # Assets statiques
└── globals.css                # Styles globaux (Tailwind)
```

---

## 🔧 Configuration

### Variables d'environnement - Backend

Créer un fichier `.env` dans le dossier `backend/`:

```env
SERVER_PORT=:8080
DATABASE_URL=sqlite:///social_network.db
```

### Variables d'environnement - Frontend

Créer un fichier `.env.local` dans le dossier `frontend/`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

---

## 📦 Technologie Stack

### Backend
- **Go 1.22** - Langage de programmation
- **Gorilla WebSocket** - Communication en temps réel
- **SQLite** - Base de données
- **UUID** - Génération d'identifiants uniques
- **crypto** - Hachage des mots de passe

### Frontend
- **Next.js 14** - Framework React
- **React 18** - Librairie UI
- **Tailwind CSS** - Styling utilitaire
- **Radix UI** - Composants accessibles
- **React Hook Form** - Gestion des formulaires
- **Zustand** - Gestion d'état
- **SWR** - Récupération de données
- **Lucide Icons** - Icônes
- **Emoji Picker** - Sélecteur d'emojis

---

## 🔗 API Endpoints

### Authentification
- `POST /register` - Inscription
- `POST /login` - Connexion
- `POST /logout` - Déconnexion

### Utilisateurs
- `GET /users/:id` - Profil utilisateur
- `PUT /users/:id` - Mise à jour profil
- `GET /users/search` - Recherche d'utilisateurs

### Posts
- `GET /posts` - Récupérer les posts
- `POST /posts` - Créer un post
- `PUT /posts/:id` - Éditer un post
- `DELETE /posts/:id` - Supprimer un post

### Commentaires
- `GET /posts/:id/comments` - Récupérer les commentaires
- `POST /posts/:id/comments` - Ajouter un commentaire
- `DELETE /comments/:id` - Supprimer un commentaire

### Followers
- `POST /users/:id/follow` - Suivre un utilisateur
- `DELETE /users/:id/follow` - Arrêter de suivre
- `GET /users/:id/followers` - Lister les followers
- `GET /users/:id/following` - Lister les utilisateurs suivis

### Groupes
- `GET /groups` - Lister les groupes
- `POST /groups` - Créer un groupe
- `GET /groups/:id` - Détails du groupe
- `POST /groups/:id/members` - Ajouter un membre
- `DELETE /groups/:id/members/:userId` - Retirer un membre

### Événements
- `GET /events` - Lister les événements
- `POST /events` - Créer un événement
- `GET /events/:id` - Détails de l'événement

### Notifications
- `GET /notifications` - Récupérer les notifications
- `PUT /notifications/:id/read` - Marquer comme lue

### Chat (WebSocket)
- `WS /ws` - Connexion WebSocket pour chat en temps réel

---

## 🗄️ Base de Données

Le projet utilise **SQLite** avec un système de migrations automatique.

Les tables principales incluent:
- `users` - Profils utilisateurs
- `posts` - Posts
- `comments` - Commentaires
- `likes_dislikes` - Likes/Dislikes
- `followers` - Relations de suivi
- `groups`, `group_members` - Groupes
- `events` - Événements
- `notifications` - Notifications
- `chats` - Messages
- `sessions` - Sessions utilisateurs

---

## 📝 Scripts Disponibles

### Backend

```bash
cd backend

# Télécharger les dépendances
go mod download

# Exécuter le serveur
go run main.go

# Construire l'exécutable
go build -o social-network-api

# Nettoyer les dépendances
go mod tidy
```

### Frontend

```bash
cd frontend

# Installer les dépendances
pnpm install

# Démarrage en développement
pnpm dev

# Build pour production
pnpm build

# Lancer en production
pnpm start

# Linter
pnpm lint
```

---

## 🐳 Docker

### Build manuel des images

```bash
# Backend
docker build -t social-network-backend ./backend

# Frontend
docker build -t social-network-frontend ./frontend
```

### Run avec docker-compose

```bash
docker-compose up
docker-compose down
docker-compose logs -f
```

---

## 🔐 Sécurité

- Les mots de passe sont hachés avec `crypto`
- Les sessions sont gérées via cookies sécurisés
- CORS est configuré pour les origins autorisées
- Authentification requise sur les endpoints protégés
- Validation des données d'entrée

---

## 🛠️ Développement

### Ajouter une nouvelle fonctionnalité

1. **Backend**: Créer un handler dans `/backend/handlers/`
2. **Service**: Implémenter la logique dans `/backend/pkg/services/`
3. **Model**: Définir le modèle dans `/backend/pkg/models/`
4. **Routes**: Ajouter la route dans `/backend/routes/routes.go`
5. **Frontend**: Créer le composant React dans `/frontend/components/`
6. **Actions**: Implémenter l'appel API dans `/frontend/app/actions/`

### Conventions de code

- **Go**: Noms en camelCase, gestion d'erreurs explicite
- **JavaScript/React**: Composants fonctionnels avec Hooks, destructuration
- **Tailwind CSS**: Classes utilitaires, dark mode supporté
- **Commits**: Messages clairs et descriptifs

---

## 📚 Documentation supplémentaire

- Voir `audit.md` pour les critères d'acceptation
- Voir `task.todo` ou `taskAudit.todo` pour les tâches en cours

---

## 🤝 Contribution

Les contributions sont les bienvenues! Pour contribuer:

1. Fork le projet
2. Créer une branche (`git checkout -b feature/amazing-feature`)
3. Commit les changements (`git commit -m 'Add amazing feature'`)
4. Push vers la branche (`git push origin feature/amazing-feature`)
5. Ouvrir une Pull Request

---

## 📄 Licence

Ce projet est sous licence [À spécifier].

---

## 📞 Support

Pour toute question ou problème, veuillez ouvrir une issue sur le repository.

---

**Dernière mise à jour:** Avril 2026
