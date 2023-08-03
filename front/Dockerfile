FROM node:lts-alpine as global-deps-stage
RUN npm i --location=global @quasar/cli@latest

FROM global-deps-stage as develop-stage
WORKDIR /src
COPY package.json ./
COPY yarn.lock ./
COPY . .

FROM develop-stage as local-deps-stage
RUN yarn

FROM local-deps-stage as build-stage
RUN quasar build -m pwa

FROM build-stage
WORKDIR /app
EXPOSE 80
CMD ["node", "index.js"]
