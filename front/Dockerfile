# Build stage
FROM node:18.17-alpine as develop-stage
WORKDIR /app
COPY package*.json ./
RUN yarn global add @quasar/cli
COPY . .

# Build stage
FROM develop-stage as build-stage
RUN yarn
RUN quasar build

# Production stage
FROM caddy:latest as production-stage
COPY --from=build-stage /app/dist/spa /srv
EXPOSE 80
CMD ["caddy", "file-server", "--root", "/srv", "--listen", "0.0.0.0:443"]
