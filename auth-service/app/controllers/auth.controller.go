package controllers

import (
	"net/http"
	"strings"
	"time"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
)

const requestTimeout = time.Second * 5

var logger = utils.NewLoggerService("auth-controller")

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login godoc
// @Summary Login
// @Description This request is used to login a user and get a token in cookies
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Param req body dto.LoginDto true "Login Details"
// @Success	200  {string} 	message
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	401  {number} 	http.StatusUnauthorized
// @Failure	404  {number} 	http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /auth/login [POST]
func (ac AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user dto.LoginDto

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		foundUser, err := ac.authService.GetUserByEmail(user.Email)

		if err != nil {
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Message: err.Message,
				Data:    nil,
				Status:  err.Code,
			})
			return
		}

		isValidPassword, err := utils.VerifyPassword(user.Password, foundUser.Password)

		if !isValidPassword {
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Message: err.Message,
				Data:    nil,
				Status:  err.Code,
			})
			return
		}
		token, err := utils.CreateToken(foundUser.Id, foundUser.Email, foundUser.Name, foundUser.IsAdmin)
		if err != nil {
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Message: err.Message,
				Data:    nil,
				Status:  err.Code,
			})
			return
		}
		token = "Bearer " + token
		c.SetCookie("token", token, 3600, "/", "", false, true)
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Message: "logged in",
			Data:    nil,
			Status:  http.StatusOK,
		})
	}
}

// Logout godoc
// @Summary Logout
// @Description This request is used to logout a user
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	message
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /auth/logout [POST]
func (AuthController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(200, dto.ResponseDTO{
			Message: "logged out",
			Data:    nil,
			Status:  http.StatusOK,
		})
	}
}

// VerifyToken godoc
// @Summary VerifyToken
// @Description This request is used to verify a token internally or by frontend
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} utils.SignedDetails
// @Failure	401  {number} http.StatusUnauthorized
// @Router /auth/verify-token [POST]
func (AuthController) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseDTO{
				Message: "token not found",
				Data:    nil,
				Status:  http.StatusUnauthorized,
			})
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		isValid, _ := utils.ValidateToken(token)
		if isValid != true {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseDTO{
				Message: "Invalid token",
				Data:    nil,
				Status:  http.StatusUnauthorized,
			})
			return
		}
		claims, err_ := utils.GetClaimsFromToken(token)
		if err_ != nil {
			c.AbortWithStatusJSON(err_.Code, dto.ResponseDTO{
				Message: err_.Message,
				Data:    nil,
				Status:  err_.Code,
			})
			return
		}
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Message: "token verified",
			Data:    claims,
			Status:  http.StatusOK,
		})
	}
}
