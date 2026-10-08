# bucket inventory

A REST API for storing files in S3 with their metadata kept in PostgreSQL. Uploaded files go to an S3 bucket, while
names, sizes and other metadata are saved in Postgres, so files can be listed and looked up without going to S3.
Downloads don't go through the API: each file comes back with a presigned S3 URL that the client uses to fetch the
object directly from the bucket.

## Features

- **Upload** files (`multipart/form-data`) to S3 and record their metadata
- **List** all stored files
- **Get** a single file by ID
- **Delete** a file from both S3 and the database
- **Presigned links** for downloading files directly from S3

## Tech stack

- **Language:** Go 1.26
- **HTTP:** Echo v5
- **Database:** PostgreSQL, go-jet (type-safe SQL), goose (migrations)
- **Storage:** AWS SDK for Go v2, LocalStack for local S3
- **App wiring & logging:** fx, slog
- **Testing:** testify, mockery
- **Tooling:** golangci-lint, GitHub Actions

## API

| Method   | Path         | Description                                   |
|----------|--------------|-----------------------------------------------|
| `POST`   | `/files`     | Upload a file (form field `file`)             |
| `GET`    | `/files`     | List files with presigned download links      |
| `GET`    | `/files/:id` | Get a file by ID with a presigned download link |
| `DELETE` | `/files/:id` | Delete a file                                 |

## Getting started

```sh
cp .env.example .env
docker compose up -d   # PostgreSQL + LocalStack
task migrate-up        # apply migrations and regenerate go-jet models
go run .
```

The server listens on `localhost:1111` by default (`HTTP_HOST` / `HTTP_PORT`).

Upload a file:

```sh
curl -F "file=@./photo.png" http://localhost:1111/upload
```

## Development

```sh
task test   # run tests with race detector and coverage
task lint   # run golangci-lint
task mock   # regenerate mocks with mockery
```
