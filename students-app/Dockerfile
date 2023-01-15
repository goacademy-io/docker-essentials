FROM node:12.11 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@8.3.12
WORKDIR /usr/src/app/
COPY webapp /usr/src/app/
RUN npm install && ng build --prod

FROM golang:1.13.1-alpine AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.10
WORKDIR app
COPY --from=ANGULAR_BUILD /usr/src/app/dist/webapp/* ./usr/src/app/dist/webapp/
COPY --from=GO_BUILD /go/bin/server ./
EXPOSE 8080
RUN ls
CMD ./server
