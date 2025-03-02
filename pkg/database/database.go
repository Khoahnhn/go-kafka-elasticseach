package database

import (
	"fmt"
	"github.com/Khoahnhn/go-kafka-elastichsearch/settings/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func InitDatabase() {
	// Load biến môi trường
	env.LoadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.GetEnv("DB_USER", "root"),
		env.GetEnv("DB_PASSWORD", "password"),
		env.GetEnv("DB_HOST", "127.0.0.1"),
		env.GetEnv("DB_PORT", "3306"),
		env.GetEnv("DB_NAME", "testdb"),
	)

	// Ket noi database
	//DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal("Không thể kết nối database: ", err)
	//}

	var err error
	for i := 0; i < 5; i++ { // Thử kết nối lại 5 lần
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			//Logger: logger.Default.LogMode(logger.Silent),
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err == nil {
			break
		}
		log.Println("⏳ MySQL chưa sẵn sàng, thử lại sau 2 giây...")
		time.Sleep(2 * time.Second)
	}

	// Chạy migration tự động
	//if err := DB.AutoMigrate(&user.User{}, &product.Product{}); err != nil {
	//	log.Fatal("Migration thất bại:", err)
	//}

	//// Chạy migration
	//if err := versions.MigrateCreateUsersTable(DB); err != nil {
	//	log.Fatalf("❌ Migration thất bại: %v", err)
	//}

	log.Println("Kết nối MySQL thành công!")

	// Kiểm tra kết nối bằng câu lệnh đơn giản
	var result int
	if err := DB.Raw("SELECT 1").Scan(&result).Error; err != nil {
		log.Fatal("Kiểm tra kết nối thất bại: ", err)
	} else {
		log.Println("Kết nối MySQL hoạt động bình thường!")
	}
}
