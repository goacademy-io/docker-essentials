## Run docker image

```bash
 docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=root --name pg postgres:14.5
 docker exec -it pg /bin/bash
 psql -U postgres
```

## Building docker image

```bash
docker build . -t pg
```


