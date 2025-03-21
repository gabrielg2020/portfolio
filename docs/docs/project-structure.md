---
sidebar_position: 3
title: Project Structure
slug: /project-structure
---

This section provides an overview of the project's directory structure and the purpose of each component.

## Root Directory

The root directory contains configuration files and directories for the main components of the portfolio:

```md
.
├── backend/           # Go backend application
├── docs/              # Documentation (Docusaurus)
├── frontend/          # Svelte frontend application
├── .dockerignore      # Specifies files to exclude from Docker builds
├── docker-compose.yaml # Docker Compose configuration
│   # Docker configuration for front/backend
├── Dockerfile.backend 
├── Dockerfile.frontend
├── go.mod
└── go.sum
```

## Backend Structure

:::danger Not Finished

This section is not complete and is subject to change.

:::

The backend is written in Go and follows a modular structure:

```md
backend/
├── handlers/          # HTTP request handlers
│   └── handlers.go    
├── routes/            # API route definitions
│   └── routes.go      
├── static/            # Static files served by the backend
└── main.go            # Application entry point
```

### Key Backend Components

1. `handlers/` - Contains the business logic for processing API requests
2. `routes/` - Defines API endpoints and connects them to handlers
3. `main.go` - Initializes and starts the web server

## Frontend Structure

:::danger Not Finished

This section is not complete and is subject to change.

:::

The frontend is built with Svelte and uses Vite as the build tool:

```md
frontend/
├── public/            # Static assets
├── src/               # Source code
│   ├── assets/        # Frontend assets (images, fonts, etc.)
│   ├── lib/           # Reusable components and utilities
│   ├── App.svelte     # Root Svelte component
│   ├── main.js        # Application entry point
├── index.html         # HTML entry point
├── svelte.config.js   # Svelte configuration
└── vite.config.js     # Vite build configuration
```

### Key Frontend Components

1. `App.svelte` - The main application component
2. `main.js` - JavaScript entry point that mounts the Svelte app

## Documentation Structure

:::danger Not Finished

This section is not complete and is subject to change.

:::

The documentation is built using Docusaurus:

```md
docs/
├── docs/               # Documentation content
├── src/
│   └── pages/
│       └── index.tsx   # Homepage
├── docusaurus.config.ts # Docusaurus configuration
└── sidebars.ts         # Documentation sidebar configuration
```

## Docker Configuration

The project uses Docker for containerization:

1. Dockerfile.backend - Defines the container for the Go backend
2. Dockerfile.frontend - Defines the container for the Svelte frontend
3. docker-compose.yaml - Orchestrates the containers and defines:
    - Service configurations
    - Volume mappings
    - Network settings
    - Environment variables

This containerized approach ensures consistency across development and production environments, making the application easy to set up and run regardless of the system.
