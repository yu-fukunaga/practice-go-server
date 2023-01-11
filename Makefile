.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.46.2

.PHONY: env
env:
	cp dev.envrc .envrc
	direnv allow .

.PHONY: mysql 
mysql: 
	docker-compose exec mysql mysql -u root --password=passwd sample_db

.PHONY: lint
lint:
	golangci-lint run

.PHONY: migration-exec
migration-exec:
	atlas migrate apply --dir file://ent/migrate/migrations --url mysql://root:passwd@localhost:3306/sample_db

.PHONY: migration-gen-%:
migration-gen-%:
	go run -mod=mod ent/migrate/main.go ${@:migration-gen-%=%}

.PHONY: clean
clean:
	docker-compose exec mysql mysql -u root --password=passwd -e "drop database if exists sample_db; create database sample_db;"
