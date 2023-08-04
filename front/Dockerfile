# Build stage
FROM node:lts-alpine as develop-stage
WORKDIR /app
COPY package*.json ./
RUN yarn global add @quasar/cli
COPY . .

# Build stage
FROM develop-stage as build-stage
RUN yarn
RUN quasar build -m pwa

# Production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist/pwa /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
