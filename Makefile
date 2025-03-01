.PHONY: run build clean chmod run-migration-up run-migration-down check-version-migration force-version goto-version

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

#chmod:
#	@chmod +x scripts/migrate.sh #Linux/MacOS
#
#run-migration-up:
#	@scripts/migrate.sh up
#
#run-migration-down:
#	@scripts/migrate.sh down
#
#check-version-migration:
#	@scripts/migrate.sh version
#
#force-version:
#	@scripts/migrate.sh force 2
#
#goto-version:
#	@scripts/migrate.sh goto 3