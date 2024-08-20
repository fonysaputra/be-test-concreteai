# Backend Test CencreteAi With Go Gin and NodeJs fastify

This repository contains a Go application and NodeJs fastify that interacts with Supabase to manage user authentication, retrieve transaction and job scheduler for the auto debit/recurring simulate payments will automatically be processed at specified intervals. This application is set up to run with Docker and Docker Compose.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) installed on your machine.
- [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine.

## Environment Variables

Create a `.env` file in the root directory of the project with the following variables:

```env
SUPABASE_URL=https://bitvvdykgckoiudndxpi.supabase.co
SUPABASE_API_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJpdHZ2ZHlrZ2Nrb2l1ZG5keHBpIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjQwNzk4MDEsImV4cCI6MjAzOTY1NTgwMX0.VrUYF0gxtraeBbTUFqKz3_b8tq4ZxYFd4xwFTqsooUU
DSN_POSTGRES=user=postgres.bitvvdykgckoiudndxpi password=P@ssw0rd2024TestBE host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 dbname=postgres
PORT=8080
```

## How To Run This Project
- **Build and Start the Services**: docker-compose up --build
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **Stopping the Services**: docker-compose down

## Testing and Troubleshooting
- **Open Swagger UI**: Visit http://localhost:8080/swagger/index.html in your web browser.
- **Login or Register** :
    - You can interact with the API directly from the Swagger UI by filling in required parameters and hitting the "Execute" button.
- **Authorize API** :
    - Copy on the access token from response login or register enter your Bearer token in the Swagger UI and hitting the "Execute button".

### Explanation:

- **Environment Variables**: The `.env` file contains sensitive information like Supabase URL And API key.
- **Dockerfile**: Builds the Go application and prepares a minimal image to run it.
- **docker-compose.yml**: Configures services for your Go application, NodeJs Fastify and a PostgreSQL database.
- **Swagger UI**: Instructions for accessing and using Swagger UI to test the API.
- **Running the Application**: Instructions on how to build and run the application with Docker Compose.
- **Testing and Troubleshooting**: Basic guidelines for testing the API and troubleshooting common issues.


This `README.md` provides a comprehensive guide to setting up and running your Go application with Docker Compose, including environment configuration and troubleshooting tips. Let me know if you need any further adjustments or additions!
