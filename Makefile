run:
	go run main.go

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