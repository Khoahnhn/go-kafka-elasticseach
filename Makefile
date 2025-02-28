.PHONY: run build clean

run:
	@echo "Starting Go server..."
	@for /f "usebackq delims=" %%x in (`type .env`) do set %%x && go run cmd/main.go
	#@env $$(cat .env | xargs) go run cmd/main.go

build:
	@echo "Building project..."
	@go build -o app cmd/main.go

clean:
	@echo "Cleaning up..."
	@rm -f app #Linux/macOS
