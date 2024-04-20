#!/bin/bash 
sudo wget -qO- https://get.docker.com/ | sh
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
git clone https://github.com/israelnp/trabalho-grafana-prometheus-go-terraform.git
cd /trabalho-grafana-prometheus-go-terraform/app
sudo docker-compose up -d