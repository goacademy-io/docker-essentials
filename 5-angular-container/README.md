## create app & run locally

```bash
ng new hello-world
cd hello-world
ng serve
```

## Docker Build and Run

```bash
cd hello-world
docker build . -t ng-app
docker run -d -p 8080:80 --name webapp ng-app
```

## Test the app

```bash
docker run -d curlimages/curl http://127.0.0.1:8080
```
