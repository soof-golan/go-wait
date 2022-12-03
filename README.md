# Go Wait

A micro backend written in go that serves one thing,
and after a certain amount of time, it serves another.

This is used to create a waiting room web page.
where the clients have no access to the secret url
until the server serves it.

For obvious reasons, when clients are too early, they get served a rickroll url.

Is this the finest piece of code you've ever seen?
No, but it works.

# Usage

## Local

```bash
# ISO 8601 time format
export THRESHOLD_TIME=2022-12-03T18:42:30+0200
# The secret url
export SECRET_URL=<your secret url>

# Run the server on port 8080
go run main.go
```

## Docker

```bash
# Copy the .env file to the root of the project
cp .env.example .env
# Then edit .env

docker compose up --build
```

# Deployment

I used GCP's Cloud Run service, it wasn't much fun, so I won't recommend it.

# License

MIT License

Copyright (c) 2022, Soof Golan
