---
layout: default
title: Project Structure
parent: Development Guide
nav_order: 2
---

# Project Structre

This section provides an overview of the project's directory structure.

## Root Directory

The root directory contains configuration files and directories for the main components of the portfolio:

```md
.
├── backend/            # Go backend application
├── docs/               # Documentation
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
├── data              # Where .json data is kept
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

## Docker Configuration

The project uses Docker for containerization:

1. Dockerfile.build - Defines the final build contianer
2. docker-compose.yaml - Orchestrates the containers and defines:
    - Service configurations
    - Volume mappings
    - Network settings
    - Environment variables

This containerized approach ensures consistency across development and production environments, making the application easy to set up and run regardless of the system.
