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