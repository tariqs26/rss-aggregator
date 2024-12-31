# RSS Aggregator

A simple RSS feed aggregator that allows users to:

- Add RSS feeds to be collected.
- Follow and un-follow RSS feeds that they have added.
- Fetch the latest posts from the RSS feeds they follow.

## API Routes

### General

- `GET /v1/ping`: health check endpoint.

### Users

- `POST /v1/users`: create a new user.
  - request body:
    ```json
    { "name": "string" }
    ```
- `GET /v1/users` (requires authentication): get current user details.
- `DELETE /v1/users` (requires authentication): delete current user.
- `GET /v1/users/posts` (requires authentication): get posts from the feeds the current user follows.

### Feeds

- `POST /v1/feeds`: add a new feed.
  - request body:
    ```json
    {
      "name": "string",
      "url": "string"
    }
    ```
- `GET /v1/feeds` (requires authentication): get all feeds.
- `DELETE /v1/feeds/{id}` (requires authentication): delete a feed by ID.
- `POST /v1/feeds/{id}/follow` (requires authentication): follow a feed by ID.

### Feed Follows

- `GET /v1/feed-follows` (requires authentication): get all followed feeds.
- `DELETE /v1/feed-follows/{id}` (requires authentication): unfollow a feed by ID.

## Setup and Installation

### Prerequisites

- [Go](https://golang.org/doc/install) 1.22.3 or later
- [PostgreSQL](https://www.postgresql.org/download/) 16.0 or later

### Clone the Repository

```bash
git clone https://github.com/tariqs26/rss-aggregator.git
cd rss-aggregator
```

### Install Dependencies

This project uses the following Go packages:

```bash
go get github.com/joho/godotenv@v1.5.1
go get github.com/go-chi/chi/v5@v5.0.12
go get github.com/google/uuid@v1.6.0
go get github.com/lib/pq@v1.10.9
go get github.com/rs/cors@v1.11.0
```

### Additional Tools

- [sqlc](https://github.com/sqlc-dev/sqlc) for generating type-safe code from SQL.

  ```bash
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  ```

- [goose](https://github.com/pressly/goose) for database migrations.

  ```bash
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```

### Environment Variables

Create a `.env` file in the root directory of your project and add the following environment variables:

```bash
PORT=8080
DATABASE_URL=postgres://username:password@localhost:5432/yourdbname?sslmode=disable
```

Replace `username`, `password`, `localhost:5432`, and `yourdbname` with your PostgreSQL credentials and database details.

## Usage

### Apply Migrations

Navigate to the directory containing the migration files and run goose to apply the migrations:

```bash
cd sql/schema
goose -dir . postgres "postgres://username:password@localhost:5432/yourdbname?sslmode=disable" up
```

### Generate SQLC Code

Generate the SQLC code for interacting with the database:

```bash
sqlc generate
```

### Build and Run the Server

Compile the server and run it:

```bash
go build -o rss-aggregator
./rss-aggregator
```

The server will start on the port specified in the `.env` file (default is `8080`).
