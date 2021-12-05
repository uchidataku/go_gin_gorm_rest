package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go_gin_gorm_rest/auth"
	"go_gin_gorm_rest/db"
	"go_gin_gorm_rest/entity"
	"log"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var payload LoginPayload
	var account entity.Account

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	db := db.GetDB()
	result := db.Where("email = ?", payload.Email).First(&account)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	//err = user.CheckPassword(payload.Password)
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(401, gin.H{
	//		"msg": "invalid user credentials",
	//	})
	//	c.Abort()
	//	return
	//}

	jwtWrapper := auth.JwtWrapper{
		SecretKey: "hogehoge",
		Subject:   account.ID,
	}

	signedToken, err := jwtWrapper.GenerateToken()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	c.JSON(200, tokenResponse)

	return
}
