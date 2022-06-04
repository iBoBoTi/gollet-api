package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/gollet-api/infrastructure/database"
	"github.com/iBoBoTi/gollet-api/internal/adapters/api/handler"
	"github.com/iBoBoTi/gollet-api/internal/adapters/api/helper"
	"github.com/iBoBoTi/gollet-api/internal/adapters/repository/psql"
	"github.com/iBoBoTi/gollet-api/internal/core/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ginEngine struct {
	router     *gin.Engine
	port       int
	ctxTimeout time.Duration
}

func newGinServer(port int, t time.Duration) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: t,
	}
}

func (g *ginEngine) setUpRouter() {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "test" {
		g.setAppHandlers()
		return
	}
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	g.router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	g.router.Use(gin.Recovery())
	// setup cors
	g.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	g.setAppHandlers()
}

func (g *ginEngine) setAppHandlers() {
	v1 := g.router.Group("/api/v1")

	tokenMaker, err := helper.NewJWTMaker(os.Getenv("TOKEN_SECRET"))
	if err != nil {
		log.Fatal("failed to create token maker")
	}

	db, err := database.NewDatabaseFactory(database.InstancePostgres)
	if err != nil {
		log.Fatal(err)
	}

	// User
	userRepo := psql.NewUserRepository(db.Postgres)
	userService := usecase.NewUserService(userRepo, tokenMaker)
	userHandler := handler.NewUserHandler(userService)

	userRouter := v1.Group("/users")
	userRouter.POST("/", userHandler.SignUpUser)
	userRouter.POST("/login", userHandler.LoginUser)
	g.router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func (g *ginEngine) Start() {
	g.setUpRouter()

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Error starting HTTP server", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed", err)
	}

	log.Println("Server down")
}
