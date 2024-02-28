TAG=gin-dubrava

all: run

build: .PHONY
	@echo "--- Сборка проекта"
	@go build -o bin/app .

run: build .PHONY
	@echo "--- Запуск проекта"
	@./bin/app

docker: .PHONY
	@echo "--- Сборка образа"
	@docker build . --tag $(TAG)

.PHONY: