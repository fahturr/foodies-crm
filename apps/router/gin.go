package router

import (
	"foodies-crm/apps/commons/middlewares"
	"foodies-crm/pkg/database"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	port       string
	router     *gin.Engine
	middleware *middlewares.GinMiddleware
	db         *database.Database
}

func NewGinRouter(port string, db *database.Database) *Gin {
	router := gin.Default()
	middleware := middlewares.NewGinMiddleware()

	return &Gin{
		port:       port,
		router:     router,
		middleware: middleware,
		db:         db,
	}
}

func (r *Gin) BuildRoutes() {
	r.router.Use(CORS)
}

func (r *Gin) Run() error {
	err := r.router.Run(r.port)

	return err
}

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
