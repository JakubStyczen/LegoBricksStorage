# LegoBricksStorage

### TODO:

- CRUD by brick - number, name, size, shape, quantity
- CRUD by LEGO set
- check if your brick database allows you to build set by passing bricks list

# Project my_project

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```

## SQL

To migrate data to sql db in _internal/sql/schema_:

```
goose postgres URL up/down
```

To generate sql functions for table run in _internal_ folder:

```
sqlc generate
```
