BINARY_NAME=doodocs

GO=go

BUILD_FLAGS=-ldflags "-s -w"

SRC_DIR=./cmd/doodocsApp

build:
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_NAME) $(SRC_DIR)

deps:
	$(GO) get github.com/BurntSushi/toml
	$(GO) get github.com/ajg/form 
	$(GO) get github.com/fatih/color
	$(GO) get github.com/go-chi/chi
	$(GO) get github.com/go-chi/chi/v5
	$(GO) get github.com/go-chi/render
	$(GO) get github.com/ilyakaznacheev/cleanenv
	$(GO) get github.com/joho/godotenv
	$(GO) get github.com/jordan-wright/email
	$(GO) get github.com/mattn/go-colorable
	$(GO) get github.com/mattn/go-isatty
	$(GO) get github.com/mattn/go-sqlite3
	$(GO) get golang.org/x/exp
	$(GO) get golang.org/x/sys 
	$(GO) get gopkg.in/yaml.v3
	$(GO) get olympos.io/encoding/edn


clean:
	$(GO) clean
	rm -f $(BINARY_NAME)

run:
	$(GO) run $(SRC_DIR)

help:
	@echo "Использование:"
	@echo "  make build    - Сборка исполняемого файла"
	@echo "  make deps     - Установка зависимостей"
	@echo "  make clean    - Очистка временных файлов"
	@echo "  make test     - Запуск тестов"
	@echo "  make run      - Запуск программы"
	@echo "  make help     - Это сообщение"

default: help
