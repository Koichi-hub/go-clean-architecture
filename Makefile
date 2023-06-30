# -----> development <-----
build:
	go build -o main ./main.go

dev:
	DOTENV_PATH=".dev.env" go run main.go

# -----> docker <-----
docker-build:
	docker build --rm --tag todo .

# -----> tests <-----
tests: test-usecases test-repos test-controllers

test-usecases:
	@echo "########## test-usecases ##########"
	go test -v ./tests/usecases_tests/
	@echo

test-repos:
	@echo "########## test-repos ##########"
	go test -v ./tests/databases_tests/mysql/repos/
	@echo

test-controllers:
	@echo "########## test-controllers ##########"
	GIN_MODE=release go test -v ./tests/controllers_tests
	@echo