package route

import (
	v1 "go_line_group/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	//載入模板
	router.LoadHTMLGlob("web/templates/*.tmpl.html")
	router.Static("/static", "static")
	//app路徑測試是否正常開啟
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	//linebot 路徑
	router.Any("/callback", v1.LineReply)

	return router
}
