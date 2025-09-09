# wash-directory-core

Core backend for "The Wash Directory" – a car wash management and directory service.

## Features

- **Car Wash Management**: Start a new car wash and track its status.
- **User & Washer Profiles**: Model washer profiles, including service details and availability, linked to user accounts.
- **REST API**: Built using [Gin](https://github.com/gin-gonic/gin) for handling HTTP requests.
- **Templating & Static Assets**: Templating for UI and Tailwind CSS for styling.
- **Database Support**: PostgreSQL connection for persistent storage (see `atlas.hcl` for setup).
- **Live Development**: Integrated with [Air](https://github.com/cosmtrek/air) for Go live reloading, [Templ](https://templ.guide/) for HTML templating, and Tailwind CSS for styles.

## Quickstart

### Prerequisites

- Go 1.21+
- PostgreSQL
- Node.js & npm (for Tailwind CSS)
- [Air](https://github.com/cosmtrek/air) (for live reload)
- [Templ](https://templ.guide/) (for template watching)

### Development Workflow

```sh
# Install dependencies
go mod tidy

# Set up and migrate your Postgres DB (adjust connection string as needed)
atlas migrate apply -u 'postgres://user:password@localhost:5432/directory_core?sslmode=disable'

# Start all dev tasks (Air, Templ, Tailwind) in parallel
make dev
