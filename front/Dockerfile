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
FROM nginx:1.25.1-alpine as production-stage
COPY --from=build-stage /app/dist/spa /usr/share/nginx/html
EXPOSE 8190
CMD ["nginx", "-g", "daemon off;"]
