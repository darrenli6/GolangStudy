package routers

import (
	"github.com/darrenli6/blog-server/global"
	"github.com/darrenli6/blog-server/internal/middleware"
	"github.com/darrenli6/blog-server/internal/routers/api"
	v1 "github.com/darrenli6/blog-server/internal/routers/v1"
	"github.com/darrenli6/blog-server/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBucket(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		// 日志
		r.Use(gin.Logger())
		// 异常处理
		r.Use(gin.Recovery())
	} else {

		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	// 限流
	r.Use(middleware.RateLimiter(methodLimiters))
	// 统一超时控制
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout * time.Second))
	// 国际化
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := v1.NewArticle()

	tag := v1.NewTag()

	// jaeger
	r.Use(middleware.Trancing())

	// 权限
	r.GET("/auth", api.GetAuth)
	// 上传文件
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	apiv1.Use(middleware.Recovery())

	{

		apiv1.POST("tags", tag.Create)
		apiv1.DELETE("tags/:id", tag.Delete)
		apiv1.PUT("tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("articles", article.Create)
		apiv1.DELETE("articles/:id", article.Delete)
		apiv1.PUT("articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
	}

	return r
}
