# -----> development <-----
build:
	go build -o main ./main.go

dev:
	DOTENV_PATH=".dev.env" go run main.go

# -----> docker <-----
docker-build:
	docker build --rm --tag go-clean-architecture .

docker-compose-up:
	docker compose -f docker-compose.yml up

docker-compose-down:
	docker compose -f docker-compose.yml down

# -----> tests <-----
tests: test-repos test-usecases test-controllers test-middlewares

test-repos:
	@echo "########## test-repos ##########"
	go test -v ./tests/databases_tests/mysql/repos/
	@echo

test-usecases:
	@echo "########## test-usecases ##########"
	go test -v ./tests/usecases_tests/
	@echo

test-controllers:
	@echo "########## test-controllers ##########"
	GIN_MODE=release go test -v ./tests/controllers_tests
	@echo

test-middlewares:
	@echo "########## test-middlewares ##########"
	go test -v ./tests/controllers_tests/middlewares_tests
	@echo