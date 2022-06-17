package router

import (
	"riddy_instagram/internal/app/riddy_instagram/recorder"
	"riddy_instagram/internal/app/riddy_instagram/routes"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(recorder.Recorder())

	routes.NotFound(r)
	routes.NoMethods(r)

	r.StaticFile("/p/Ceqo9R8MFz4", "./web/static/boneralert.html")
	r.StaticFile("./favicon.ico", "./web/static/favicon.ico")

	return r
}
