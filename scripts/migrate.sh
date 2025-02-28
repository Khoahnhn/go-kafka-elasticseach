##!/bin/bash
#
## Tìm đường dẫn của migrate
#if ! command -v migrate &> /dev/null; then
#    echo "❌ ERROR: 'migrate' không được tìm thấy! Hãy kiểm tra lại GOPATH/bin."
#    exit 1
#fi
#
## Cấu hình database (thay đổi tùy theo môi trường của bạn)
## Load biến môi trường từ file .env nếu có
#if [ -f .env ]; then
#  # shellcheck disable=SC2046
#  export $(grep -v '^#' .env | xargs)
#fi
#
## Kiểm tra nếu biến DB_HOST chưa được thiết lập
#if [ -z "$DB_HOST" ]; then
#  echo "ERROR: Biến môi trường chưa được load! Hãy kiểm tra file .env"
#  exit 1
#fi
#
## Tạo connection string cho Golang Migrate
#DB_URL="mysql://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME"
#MIGRATIONS_DIR=".db/migration/version"
#MIGRATE_BIN="$HOME/go/bin/migrate"
#
## Nếu đang chạy trên Windows (Git Bash, WSL), cập nhật đường dẫn
#if [[ "$(uname -s)" =~ MINGW|MSYS|CYGWIN ]]; then
#  MIGRATE_BIN="C:/Users/Admin/go/bin/migrate.exe"
#fi
#
## Kiểm tra tham số đầu vào
#if [ "$#" -eq 0 ]; then
#  echo "Usage: $0 {up|down|force|goto <version>|version|status}"
#  exit 1
#fi
#
## Chạy lệnh migration theo tham số
#case "$1" in
#  up)
#    echo "🔼 Running all pending migrations..."
#    $(command -v migrate) -database "$DB_URL" -path "$MIGRATIONS_DIR" up
#    ;;
#  down)
#    echo "🔽 Rolling back last migration..."
#    "$MIGRATE_BIN" -database "$DB_URL" -path "$MIGRATIONS_DIR" down 1
#    ;;
#  force)
#    if [ -z "$2" ]; then
#      echo "⚠️  Missing version number. Usage: $0 force <version>"
#      exit 1
#    fi
#    echo "⚠️  Forcing migration version to $2..."
#    "$MIGRATE_BIN" -database "$DB_URL" -path "$MIGRATIONS_DIR" force "$2"
#    ;;
#  goto)
#    if [ -z "$2" ]; then
#      echo "⚠️  Missing version number. Usage: $0 goto <version>"
#      exit 1
#    fi
#    echo "➡️  Migrating to version $2..."
#    "$MIGRATE_BIN" -database "$DB_URL" -path "$MIGRATIONS_DIR" goto "$2"
#    ;;
#  version)
#    echo "📜 Current migration version:"
#    "$MIGRATE_BIN" -database "$DB_URL" -path "$MIGRATIONS_DIR" version
#    ;;
#  status)
#    echo "🔍 Checking migration status..."
#    "$MIGRATE_BIN" -database "$DB_URL" -path "$MIGRATIONS_DIR" version
#    ;;
#  *)
#    echo "❌ Invalid command. Usage: $0 {up|down|force|goto <version>|version|status}"
#    exit 1
#    ;;
#esac