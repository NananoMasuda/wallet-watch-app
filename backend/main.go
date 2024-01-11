package main

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// DB接続
	models.ConnectDataBase()

	// Ginルーター初期化
	router := gin.Default()

	public := router.Group("/api")

	// ユーザー登録
	public.POST("/register", controllers.Register)

	// ログイン（ユーザーの認証・トークン生成）
	public.POST("/login", controllers.Login)

	// トークン認証
	protected := router.Group("/api/admin")

	tokenMiddleware := middlewares.JwtAuthMiddleware()
	protected.Use(tokenMiddleware)

	// ユーザー取得
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":8000")
}
