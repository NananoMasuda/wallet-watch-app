package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/login", func(c *gin.Context) {
		var loginRequest struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// ログインの処理をここに実装
		if isValidUser(loginRequest.Username, loginRequest.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "ログイン成功"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "ログイン失敗"})
		}
	})

	r.Run(":8080")
}

func isValidUser(username, password string) bool {
	// ユーザーの認証処理
	// 例: ユーザー名が "admin" でパスワードが "password" の場合にtrueを返す
	return username == "admin" && password == "password"
}
