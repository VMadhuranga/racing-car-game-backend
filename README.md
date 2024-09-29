# Racing Car Game Backend

This repository holds backend code for [Racing Car Game](https://github.com/VMadhuranga/racing-car-game)

## Prerequisites

You need to have the following installed on your computer to run this program locally

- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

## Run Locally

- [Fork and clone](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo) the project.

- Go to the project directory

  ```bash
  cd racing-car-game-backend/
  ```

- Install dependencies

  ```bash
  go mod download
  ```

- Create `.env` file with following environment variables

  - PORT

    eg: `PORT=8080`

  - POSTGRES_URI

    > Note: Make sure to disable `sslmode` if you use local postgre db connection string

    eg: `POSTGRES_URI=protocol://username:password@host:port/database?sslmode=disable`

  - FRONTEND_BASE_URL

    eg: `FRONTEND_BASE_URL=http://localhost:5173`

  - ACCESS_TOKEN_SECRET
  - REFRESH_TOKEN_SECRET

    > Note: Use `openssl rand -base64 64` command to generate random string to use as token secrets

- Start the application

  ```
  go build -o rcg && ./rcg
  ```

  > Note: Make sure the frontend application is running
