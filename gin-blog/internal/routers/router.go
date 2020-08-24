package routers

import (
	"gin-blog/global"
	"gin-blog/internal/middleware"
	v1 "gin-blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "gin-blog/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//路由管理
func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//新增Translations的注册
	r.Use(middleware.Translations())

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := v1.NewUpload()


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))




	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags",tag.Create)
		apiv1.DELETE("/tags/:id",tag.Delete)
		apiv1.PUT("/tags/:id",tag.Update)
		apiv1.PATCH("/tags/:id/state",tag.Update)
		apiv1.GET("/tags",tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id",article.Delete)
		apiv1.PUT("/articles/:id",article.Update)
		apiv1.PATCH("/articles/:id/state",article.Update)
		apiv1.GET("/articles/:id",article.Get)
		apiv1.GET("/articles",article.List)
	}
	return r
}

