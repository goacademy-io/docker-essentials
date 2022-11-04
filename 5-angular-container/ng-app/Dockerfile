## Stage-1 Build angular with node
FROM node:14 as BUILDER
WORKDIR /usr/src/app
COPY package.json package-lock.json ./
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
RUN npm set registry https://registry.npmjs.org/
RUN npm install
COPY . .
RUN ./node_modules/.bin/ng build --configuration production --aot --build-optimizer

## Stage-2 Copy the packaged build to nginx
FROM nginx:1.16.0-alpine
LABEL maintainer="info@goacademy.io"
RUN rm -rf /usr/share/nginx/html/*
COPY --from=BUILDER /usr/src/app/dist/hello-world /usr/share/nginx/html/
EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]
