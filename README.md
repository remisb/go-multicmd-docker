# go-multicmd-docker

Minimalistic golang project with multiple executables in the cmd folder with use of internal package. Project with docker compose and traefik.

## Docker compose 

Docker compose instructions makes bellow provided Docker sercion documentation obsolete. It left for time beeing without changes but will be updated in future.

Currently to start docker compose managed service run bellow provided commands.

```shell
# start all docker composed managed services
docker-compose up

# stop all docker compose managed services
docker-compose down 

# compose file defines 4 services named as: reverse-proxy, whoami, server, admin.
# to start only one of the specified services
docker-compose up -d server
docker-compose up -d admin
docker-compose up -d whoami

# to stop only one of the specified services
docker-compose down -d server
docker-compose down -d admin
docker-compose down -d whoami
```

HTTP routes available after all containers has been started

```http
# traefik dashboard (read only)
GET http://localhost:8000

GET http://whoami.docker.localhost/
GET http://server.docker.localhost/server
GET http://server.docker.localhost/server/health
GET http://admin.docker.localhost/admin
GET http://admin.docker.localhost/admin/health

# BTW all url is still served
GET http://localhost:8080/server
GET http://localhost:8080/server/health
GET http://localhost:8081/admin
GET http://localhost:8081/admin/health
```

Traefik related documentation can be found at https://doc.traefik.io/traefik/routing/entrypoints/ .


## Docker

### Traefik

TBD

TODO
- INPROCESS add traefik container
- INPROCESS add compose file 


`docker compose up -d reverse-proxy`

### server

To build server in a docker container from shell without `docker compose`, execute `docker build -t multicmd-server -f cmd/server/Dockerfile .`.

To start docker container with the server from shell withou `docker compose`, execute  `docker run --rm --name multi-server -d -p 8080:8080 multicmd-server`. 

After successful server startup you should see following output in the terminal `Server started and accessible at localhost: :8080`. Now you can access following url addresses served from the server in the docker container: `http://localhost:8080/server`, `http://localhost:8080/server/health`.

#### Configurable options

Server application in the docker container can be configured with an environment variable `SERVER_ADDRESS`. Default SERVER_ADDRESS value is set to `:8080`.

### admin

To build admin in a docker container from shell without `docker compose`, execute `docker build -t multicmd-admin -f cmd/admin/Dockerfile .`.

To start docker container with the server from shell withou `docker compose`, execute  `docker run --rm --name multi-admin -d -p 8081:8081 multicmd-admin`. 

After successful admin startup you should see following output in the terminal `Admin started and accessible at localhost: :8081`. Now you can access following url addresses served from the server in the docker container: `http://localhost:8080/server`, `http://localhost:8080/server/health`.

#### Configurable options

Admin application in the docker container can be configured with an environment variable `ADMIN_ADDRESS`. Default ADMIN_ADDRESS value is set to `:8081`.