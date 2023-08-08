#!/usr/bin/env bash

# Extrai a tag atual
TG=$(cat docker-compose.yml | grep otavio27/app_joinbus: | awk -F':' '{print $3}')

# Divide a tag nas partes inteira e decimal
IFS='.' read -ra TAG_PARTS <<<"$TG"
major=${TAG_PARTS[0]}
minor=${TAG_PARTS[1]}
patch=${TAG_PARTS[2]}

# Incrementa a tag
patch=$((patch + 1))

if [ $patch -eq 10 ]; then
    patch=0
    minor=$((minor + 1))

    if [ $minor -eq 10 ]; then
        minor=0
        major=$((major + 1))
    fi
fi

# Atualiza a tag no arquivo docker-compose.yml
new_tag="$major.$minor.$patch"

sed -i "s|otavio27/app_joinbus:${TG}|otavio27/app_joinbus:${new_tag}|g" docker-compose.yml

docker ps >/dev/null 2>&1
if [ $? -eq 1 ]; then
 #O comando sudo setfacl --modify user:otavio:rw /var/run/docker.sock
    #atribui permissões ACL (Access Control List) ao arquivo /var/run/docker.sock
    #para permitir que o usuário otavio tenha permissões de leitura e gravação (read-write) sobre ele.
    #Isso permite que o usuário otavio acesse o socket do Docker
    #e execute comandos Docker sem a necessidade de privilégios de superusuário.
    sudo setfacl --modify user:otavio:rw /var/run/docker.sock
    docker build -t otavio27/app_joinbus:${new_tag} . && clear
    docker push otavio27/app_joinbus:${new_tag} && sleep 2s && clear
else
    docker build -t otavio27/app_joinbus:${new_tag} . && clear
    docker push otavio27/app_joinbus:${new_tag} && sleep 2s && clear
fi
