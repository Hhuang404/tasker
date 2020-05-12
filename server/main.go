package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"tasker/server/common"
	"tasker/server/router"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRoute(r)
	port := viper.GetString("server.port")
	InitWeb(r)
	if port != "" {
		panic(r.Run(":" + port))
	}
}

func InitWeb(r *gin.Engine) {
	r.LoadHTMLGlob("dist/*.html")              // 添加入口index.html
	r.LoadHTMLFiles("static/*/*")              // 添加资源路径
	r.Static("/static", "./dist/static")       // 添加资源路径
	r.StaticFile("/hello/", "dist/index.html") //前端接口
}

//读取配置文件
func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("db")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	readErr := viper.ReadInConfig()
	if readErr != nil {
		panic(readErr)
	}
}
