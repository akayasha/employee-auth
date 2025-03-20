package controllers

import (
	"employee-auth/service"
	"employee-auth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func RegisterUser(c *gin.Context) {
	var registerData struct {
		NIP      string `json:"nip"`
		Fullname string `json:"fullname"`
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Dob      string `json:"dob"`
	}

	// Parse JSON input
	if err := c.ShouldBindJSON(&registerData); err != nil {
		utils.RespondError(c, 400, "Invalid request data")
		fmt.Println("JSON binding error:", err)
		return
	}

	// Parse the date using format "YYYY-MM-DD"
	dob, err := time.Parse("2006-01-02", registerData.Dob)
	if err != nil {
		utils.RespondError(c, 400, "Invalid date format, use YYYY-MM-DD")
		fmt.Println("Date parsing error:", err)
		return
	}

	// Register user
	user, err := service.RegisterUser(
		registerData.NIP,
		registerData.UserName,
		registerData.Fullname,
		registerData.Email,
		registerData.Password,
		registerData.Role,
		dob,
	)
	if err != nil {
		utils.RespondError(c, 400, err.Error())
		fmt.Println("User registration error:", err)
		return
	}

	utils.Response(c, 201, "User registered successfully", user)
}

// Login User
func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		utils.RespondError(c, 400, "Invalid request data")
		return
	}

	response, err := service.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		utils.RespondError(c, 401, err.Error())
		return
	}

	utils.Response(c, 200, "Login successful", response)

}
