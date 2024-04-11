# go-multicmd-docker

Minimalistic golang project with multiple executables in the cmd folder with use of internal package. Project with docker compose and traefik.

## Docker

### Traefik


### server

To build server in a docker container from shell without `docker compose`, execute `docker build -t multicmd-server -f cmd/server/Dockerfile`.

To start docker container with the server from shell withou `docker compose`, execute  `docker run -p 8080:8080  19bc`. 

After successful server startup you should see following output in the terminal `Server started and accessible at localhost: :8080`. Now you can access following url addresses served from the server in the docker container: `http://localhost:8080/server`, `http://localhost:8080/server/health`.

#### Configurable options

Server application in the docker container can be configured with an environment variable `SERVER_ADDRESS`. Default SERVER_ADDRESS value is set to `:8080`.

### admin

To build admin in a docker container from shell without `docker compose`, execute `docker build -t multicmd-admin -f cmd/admin/Dockerfile`.

To start docker container with the server from shell withou `docker compose`, execute  `docker run -p 8081:8081  19bc`. 

After successful admin startup you should see following output in the terminal `Admin started and accessible at localhost: :8081`. Now you can access following url addresses served from the server in the docker container: `http://localhost:8080/server`, `http://localhost:8080/server/health`.

#### Configurable options

Admin application in the docker container can be configured with an environment variable `ADMIN_ADDRESS`. Default ADMIN_ADDRESS value is set to `:8081`.