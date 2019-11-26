package config

import (
	"GoPratice/model"
	"github.com/joho/godotenv"
	"os"
)

// Init 初始化配置
func Init() error {
	//Load will read your env file(s) and load them into ENV for this process.
	godotenv.Load()
	// 读取翻译文件
	// if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	// 	panic(err)
	// }
	//链接数据库
	err := model.Database(os.Getenv("MYSQL_DSN"))
	if err != nil {
		panic(err)
	}
	return nil

}
