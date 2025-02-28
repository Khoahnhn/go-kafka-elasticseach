package env

import (
	"github.com/spf13/viper"
	"log"
)

//func LoadEnv() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
//	}
//}
//
//// GetEnv lấy giá trị của biến môi trường với giá trị mặc định
//func GetEnv(key, defaultValue string) string {
//	if value, exists := os.LookupEnv(key); exists {
//		return value
//	}
//	return defaultValue
//}

func LoadEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
	}
}

func GetEnv(key, defaultValue string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}
