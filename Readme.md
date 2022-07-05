# MyApp is a trivial REST backend to use so I can have something to deploy (in order learn infra stuff)


### Build Docker image:
```bash docker build --tag my-app```


### View Docker images (my-app:latest should be there):
```bash docker build ls```


### To remove an image:
```bash docker image rm my-app:latest```


### Run the container, expose port 8080 on the container to 8080 on the host:
```bash docker run --publish 8080:8080 my-app```


### Ping the app
```bash curl http://localhost:8080/ping```


### Run container in detacted mode:
```bash docker run -d -p 8080:8080 docker-gs-ping```


### List running containers:
```bash docker ps```


### List all containers:
```bash docker ps -all```


### Restart a stopped container:
```bash docker restart a-stopped-container```


### Stop a container:
```bash docker stop a-running-container```


### Run container and give it a name:
```bash docker run -d -p 8080:8080 --name rest-server docker-gs-ping```


### Some examples to curl the endpoints:
```bash
curl http://localhost:8080/records

curl http://localhost:8080/records \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "5f26f150-3df3-443c-b7f6-5a9b2b611258","title": "Reckless","artist": "Todd Helder","price": 22.29}'

curl http://localhost:8080/records/2

curl http://localhost:8080/records/858c98db-10c3-4959-ad68-0c8c526dca07
```
