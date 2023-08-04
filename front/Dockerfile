FROM node:lts-alpine as global-deps-stage
RUN npm i --location=global @quasar/cli@latest
RUN npm install -g http-server

WORKDIR /src
COPY package.json ./
COPY yarn.lock ./
COPY . .

RUN yarn

RUN quasar build -m pwa

CMD http-server -p80 /src/dist/pwa