## How to build and run
1. Create a Docker network:
    ```shell script
    docker network create students-net
    ```
2. Start the DB:
    ```shell script
    docker run -e POSTGRES_USER=go -e POSTGRES_PASSWORD=your-strong-pass -e POSTGRES_DB=go --name students-db --net=students-net postgres:11.5
    ```
3. Build the application image:
    ```shell script
    docker build -t students-app .
    ```
4. Start the application container:
    ```shell script
    docker run -p 8080:8080 -e DB_PASS='somepass' --net=students-net students-app
    ```
Access the application via http://localhost:8080
