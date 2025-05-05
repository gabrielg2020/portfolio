---
layout: default
title: Local Setup
parent: Development Guide
nav_order: 1
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

    ```bash
    docker compose up frontend-dev backend-dev
    ```

## Accessing Development Server

Once the container is running, you access:

- API: [http://localhost:8080](http://localhost:8080)
- Frontend: [http://localhost:5173](http://localhost:5173)

