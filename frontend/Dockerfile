# frontend/dockerfile
# Utiliser une image de base officielle de Node.js
FROM node:alpine

# Définir le répertoire de travail dans le conteneur

WORKDIR /app/frontend

# Copier les fichiers package.json et package-lock.json et installer les dépendances
COPY package.json package-lock.json ./

# Copier le reste des fichiers de l'application
COPY . /app/frontend

RUN npm install
# Construire l'application
RUN npm run build

# Exposer le port sur lequel l'application écoute
EXPOSE 3000

# Démarrer l'application
CMD ["npm", "start"]