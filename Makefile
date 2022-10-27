dc := docker compose --file=.docker/docker-compose.yml --env-file=.docker/.env --project-name=notification

setup:
	cp config.example.yml config.yml
	cp .docker/.env.example .docker/.env
	make install binaries
	make d-up

d-build:
	@$(dc) build

d-up:
	@$(dc) up -d

d-ps:
	@$(dc) ps

d-logs:
	@$(dc) logs -f $(service)

d-down:
	@$(dc) down --remove-orphans

http-server:
	go run . http

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm -rf coverage.out

gen-models:
	sqlboiler psql

gen-mock:
	mockgen -source=internal/messagetemplate/template.go -destination=internal/messagetemplate/mock/template.go
	mockgen -source=internal/channel/driver.go -destination=internal/channel/mock/driver.go
	mockgen -source=internal/repository/message.go -destination=internal/repository/mock/message.go
	mockgen -source=internal/repository/user.go -destination=internal/repository/mock/user.go
	mockgen -source=logger/logger.go -destination=logger/mock/logger.go

install-binaries:
	# https://github.com/volatiletech/sqlboiler
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
