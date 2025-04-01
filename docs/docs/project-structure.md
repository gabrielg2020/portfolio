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
├── backend/            # Go backend application
├── docs/               # Documentation (Docusaurus)
├── frontend/           # Svelte frontend application
├── .dockerignore       # Specifies files to exclude from Docker builds
├── docker-compose.yaml # Docker Compose configuration
├── Dockerfile.build    # Docker configuration for building
├── go.mod
└── go.sum
```

## Backend Structure

```md
backend
├── database
│   ├── data          # Data patches
│   ├── portfolio.db  # SQLite Database
│   └── rebuild_db.sh # Database rebuilding script
├── handlers          # HTTP request handlers
├── logger            # Where logger is kept
├── queries           # Query builders
├── routes            # API route definitions
├── static            # Static files served by the backend (When `npm run build` is ran)
├── utils             # Utility scripts
└── main.go           # Application entry point
```

## Frontend Structure

The frontend is built with Svelte and uses Vite as the build tool:

```md
frontend/
├── public/            # Static assets
├── src/               # Source code
│   ├── assets/        # Frontend assets
│   ├── scripts/       # Holds scripts for API request and theme state
│   ├── components/    # Reusable Svelte components
│   ├── pages/         # Main sections of frontend
│   ├── App.svelte     # Root Svelte component
│   ├── main.ts        # Application entry point
├── index.html         # HTML entry point
├── svelte.config.ts   # Svelte configuration
└── vite.config.ts     # Vite build configuration
```

## Documentation Structure

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

1. Dockerfile.build - Defines the final build contianer
2. docker-compose.yaml - Orchestrates the containers and defines:
    - Service configurations
    - Volume mappings
    - Network settings
    - Environment variables

This containerized approach ensures consistency across development and production environments, making the application easy to set up and run regardless of the system.
