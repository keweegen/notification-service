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

generate:
	go generate ./...

install-binaries:
	# See: https://github.com/volatiletech/sqlboiler
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
