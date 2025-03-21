---
sidebar_position: 2
title: Local Setup
slug: /local-setup
---

## Perquisites

Before setting up the local development environment, ensure you have the following prerequisites installed on your system:

1. Docker - The project is fully containerized, requiring Docker to run all components.
    - [Download Docker](https://www.docker.com/get-started/)
    - minimum required version: 20.10.0+

2. Docker Compose - Used to orchestrate the multi-container application.
    - Usually included with Docker Desktop
    - Minimum required version: 2.0.0+

3. Git - Required to clone the repository.
    - [Download Git](https://git-scm.com/downloads)

## Getting Started

1. Clone the Repository

    ```bash
    git clone https://github.com/gabrielg2020/portfolio.git
    cd portfolio
    ```

2. Starting the Development Environment
  
    The project contains three main containers:

    - `frontend-dev`: The Svelte frontend development server with hot-reloading
    - `backend-dev`: The Go API server with auto-recompilation
    - `doc-dev`: Documentation server

    To start all development container simultaneously run:

    ```bash
    docker compose up frontend-dev backend-dev doc-dev
    ```

    This command will:

    1. Build all necessary Docker images if they don't exist
    2. Start the three development containers
    3. Mount local source code into the containers for live editing
    4. Configure port forwarding to your local machine

## Accessing Development Servers

Once the containers are running, you access:

- Frontend: [http://localhost:5173](http://localhost:5173)
- Backend: [http://localhost:8080](http://localhost:8080)
- Documentation: [http://localhost:3000](http://localhost:3000)

## Development Workflow

### Code Changes

The development containers are configured with volume mounts to the local source code:

- Frontend (Svelte) changes will automatically trigger hot-reloading
- Backend (Go) changes will trigger automatic recompilation
- Documentation changes will refresh the documentation site

### Viewing Logs

To view logs for a specific service:

```bash
docker compose logs -f [service-name]
```

For example:

```bash
docker compose logs -f backend-dev
```

### Stopping the Environment

To stop all running containers:

```bash
docker compose down
```

### Rebuilding Containers

If you make changes to the Dockerfiles or need to rebuild the containers

```bash
docker compose build [service-name]
```

Then restart the services

## Troubleshooting

### Common Issues

1. Port Conflicts: If you already have services running on ports 3000, 8080, or 8081, modify the port mappings in the docker-compose.yml file.
2. Docker Resource Allocation: Ensure Docker has enough resources allocated (memory, CPU) in Docker Desktop settings.
3. Permission Issues: On Linux systems, you might encounter permission issues with mounted volumes. Run the following command from the project root:

    ```bash
    sudo chown -R $USER:$USER .
    ```
