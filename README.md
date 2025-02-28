
#lệnh tạo migration
migrate create -ext sql -dir migrations -seq create_users_table

#lệnh chạy migration up
migrate -database "mysql://user:password@tcp(localhost:3306)/dbname" -path migrations up

#lệnh chạy migration down
migrate -database "mysql://user:password@tcp(localhost:3306)/dbname" -path migrations down 1