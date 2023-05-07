# go-mangosteen

## sqlc & golang-migrate

```bash
brew install sqlc
brew install golang-migrate
```

## migrate database

```hash
migrate create -ext sql -dir config/migrations -seq add_phone_for_users
```
## run migration file

```bash
migrate -database "postgres://mangosteen:123456@localhost:5432/mangosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" up
```
