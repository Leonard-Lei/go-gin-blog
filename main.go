package main

/*
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8181") // listen and serve on 0.0.0.0:8080
}
*/
import (
	db "go-gin-blog/database"
 router "go-gin-blog/routers"
)

func main() {
   //数据库
   defer db.SqlDB.Close()

   //路由部分
   router:=router.InitRouter()

   //静态资源
   router.Static("/static", "./static")

   //运行的端口
   router.Run(":8000")

}
