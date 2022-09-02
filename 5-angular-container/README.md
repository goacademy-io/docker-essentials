## create app & run locally

```bash
ng new hello-world
cd hello-world
ng serve
```

## build using dockerfile

copy nginx/default.conf to hello-world
copy Dockerfile to hello-world

```bash
docker build . -t ng-app
```

## access the app from ui

```bash
http://127.0.0.1:8080
```

## access using curl container

```bash
docker run -d curlimages/curl http://127.0.0.1:8080
```
