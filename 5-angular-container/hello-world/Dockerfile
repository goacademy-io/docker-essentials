## Stage-1 Build angular with node
FROM node:12.18.3 as BUILDER
RUN npm set registry https://registry.npmjs.org/
RUN npm install
WORKDIR goacademy
COPY . .
RUN ./node_modules/.bin/ng build --prod

## Stage-2 Copy the packaged build to nginx
FROM nginx:1.16.0-alpine
LABEL maintainer="info@goacademy.io"
COPY nginx/default.conf /etc/nginx/conf.d/
COPY --from=BUILDER /goacademy/hello-world /usr/share/nginx/html/
EXPOSE 8080 8443
CMD ["/bin/sh",  "-c",  "exec nginx -g 'daemon off;'"]
