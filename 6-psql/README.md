## Run docker image

```bash
 docker run -d --name pg -p 5432:5432 -e POSTGRES_PASSWORD=root  postgres:14.5
 docker exec -it pg /bin/bash
 psql -U postgres
```

## Building docker image

```bash
docker build . -t pg
```


