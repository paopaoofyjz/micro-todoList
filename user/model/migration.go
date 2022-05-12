package model

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf-8mb4").AutoMigrate(&User{})
}
