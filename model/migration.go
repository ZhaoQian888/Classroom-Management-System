package model

func migration() {
	// 自动迁移
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{})
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&ClassRoom{})
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&ClassRoomOrder{})
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&RoomStatus{})
}
