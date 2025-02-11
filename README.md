# Task Service

This project is a Go-based backend service that implements a simple task management API. It allows you to create tasks with a title, description, and automatically generated creation timestamp. Additionally, it integrates with Prometheus to expose metrics (total tasks created and task creation duration) and provides a Docker Compose setup to run the backend service along with PostgreSQL, Prometheus, and Grafana.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)

## Features

- **Task API**:  
  - `POST /tasks`: Create a new task with a title and description.
  - Automatically stores the task with an auto-generated `id` and `created_at` timestamp.
- **Database Integration**:  
  - Uses PostgreSQL for data persistence.
  - Automatic migration using GORM.
- **Metrics**:  
  - Exposes Prometheus metrics:
    - `tasks_created_total`: Counter for created tasks.
    - `task_creation_duration_seconds`: Histogram for task creation processing time.
- **Docker Compose Setup**:  
  - Runs the backend service, PostgreSQL, Prometheus, and Grafana.
- **Clean Architecture**:  
  - Organized into domain, repository, service, and delivery layers for maintainability.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- (Optional) [Go 1.20+](https://golang.org/dl/) for local development