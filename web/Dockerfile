FROM node:12.7-alpine AS front
LABEL Omekov Azamat <umekovazamat@gmail.com>
WORKDIR /app
COPY . .
RUN npm install && npm run build

# stage 2

FROM nginx:1.17.1-alpine
COPY --from=front /app/dist/frontend-angular /usr/share/nginx/html
EXPOSE 80
