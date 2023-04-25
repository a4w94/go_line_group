package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//richmenu.Build_RichMenu()

	router := gin.Default()
	// router.Use(gin.BasicAuth(gin.Accounts{
	// 	"admin": "123456",
	// }))
	router.LoadHTMLGlob("web/templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	//router.Any("/callback", v1.LineReply)

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.

	return router
}
