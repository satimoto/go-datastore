# go-datastore
Satimoto datastore using golang

## Dependencies
```bash
brew install sqlc
brew install golang-migrate
```

## On schema changes
Whenever the SQL schema files change the following should be run

### Generate the SQL models and queries
Generates the database models and queries from sql files
```bash
sqlc generate
```

## Run Migration
Data migration can be run by either creating a .env or passing migration parameters to the command
```bash
go run ./cmd/data-migrate
```