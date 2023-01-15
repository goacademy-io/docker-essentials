## How to build and run
1. Create a Docker network:
    ```shell script
    docker network create students-net
    ```
2. Start the DB:
    ```shell script
    docker run -d -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=students --name students-db --net=students-net postgres:14.5
    ```
3. Build the application image:
    ```shell script
    docker build -t students-app .
    ```
4. Start the application container:
    ```shell script
    docker run -d -p 8080:8080 -e DB_PASS='postgres' --net=students-net students-app
    ```
Access the application via http://localhost:8080
