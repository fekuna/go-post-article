.PHONY: migrate migrate_down migrate_up migrate_version docker local test

# ==============================================================================
# Go migrate postgresql

migrate_create:
	migrate create -seq -ext=.sql -dir=./migrations ${file_name}

force:
	migrate -database "mysql://root:secret@tcp(localhost:3306)/post_article" -path migrations force 1

version:
	migrate -database "mysql://root:secret@tcp(localhost:3306)/post_article" -path migrations version

migrate_up:
	migrate -database "mysql://root:secret@tcp(localhost:3306)/post_article" -path migrations up 1

migrate_down:
	migrate -database "mysql://root:secret@tcp(localhost:3306)/post_article" -path migrations down 1

# ==============================================================================
# Docker compose commands

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build -d


# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)

clean-volume:
	sudo rm -rf ./.local