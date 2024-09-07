#!/bin/bash

install_docker() {
  echo "Installation de Docker..."

  sudo apt-get update
  sudo apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

  sudo mkdir -p /etc/apt/keyrings
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

  echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

  sudo apt-get update
  sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

  echo "Installation de Docker Compose..."
  sudo curl -L "https://github.com/docker/compose/releases/download/v2.12.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose

  echo "Docker et Docker Compose ont été installés."
}

run_docker() {
  echo "Construction et démarrage des conteneurs Docker..."
  docker-compose up --build -d
  echo "Vérification de l'état des conteneurs..."
  docker ps -a
  echo "Les conteneurs backend et frontend sont en cours d'exécution."
  echo "Ouverture de l'application dans le navigateur..."
  xdg-open http://localhost:3000
}

if [ "$1" == "-install" ]; then
  install_docker
fi

run_docker
