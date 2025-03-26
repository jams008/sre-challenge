# Runbook: Docker Compose for Virtual Pet Application Development

## Overview
This runbook provides instructions for setting up and running the Virtual Pet application using Docker Compose. The setup includes the main application, MongoDB, Prometheus, and Alertmanager.

## Prerequisites
Before running the setup, ensure that you have the following installed:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Configuration
### Environment Variables
Create a `.env` file in the same directory as the `docker-compose.yaml` file and define the following variables:

```ini
SERVER_PORT=8080
MONGODB_PORT=27017
MONGODB_DATABASE=virtual_pet_db
MONGODB_USERNAME=admin
MONGODB_PASSWORD=securepassword
```

## Running the Application
### Start the Services
To start all services, run the following command:
```sh
docker-compose up -d
```

### Stop the Services
To stop all running services, use:
```sh
docker-compose down
```

### Restart the Services
If you need to restart the services, run:
```sh
docker-compose restart
```

## Logs and Debugging
### Viewing Logs
To view logs for a specific service, use:
```sh
docker-compose logs -f <service_name>
```
Example:
```sh
docker-compose logs -f app
```

### Checking Running Containers
To list running containers:
```sh
docker ps
```

### Accessing a Running Container
To open a shell inside a running container:
```sh
docker exec -it <container_id> /bin/sh
```
Example:
```sh
docker exec -it virtual-pet-app /bin/sh
```

## Cleanup
To remove all containers, networks, and volumes:
```sh
docker-compose down -v
```

## References
- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
