FROM node:lts-alpine

# instala um servidor http simples para servir conteúdo estático
RUN npm install -g http-server

WORKDIR /src
COPY package.json ./
COPY yarn.lock ./
COPY . .

# instala dependências do projeto
RUN yarn
RUN quasar build -m pwa
COPY . .

EXPOSE 80
CMD [ "http-server", "dist" ]