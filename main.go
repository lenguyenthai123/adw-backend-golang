package main

import (
	"backend-golang/api"
	"backend-golang/config"
	"backend-golang/pkgs/transport/http/server"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	KeycloakURL   = "http://localhost:8080/auth"       // URL của Keycloak
	Realm         = "master"                           // Tên Realm
	ClientID      = "admin-client"                     // Client ID
	ClientSecret  = "nVS1tk1zrmPhstmpfx6gJ49KmT4Jih5w" // Client Secret
	UsernameAdmin = "admin-client"
	PasswordAdmin = "123456"
)

var gocloakClient *gocloak.GoCloak

func init() {
	// Khởi tạo GoCloak client
	gocloakClient = gocloak.NewClient(KeycloakURL)
}

// Đăng ký người dùng mới
func RegisterUserHandler(c *gin.Context) {
	var req struct {
		Username  string `json:"username" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Email     string `json:"email" binding:"required"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tạo user trong Keycloak
	user := gocloak.User{
		Username:  gocloak.StringP(req.Username),
		Email:     gocloak.StringP(req.Email),
		FirstName: gocloak.StringP(req.FirstName),
		LastName:  gocloak.StringP(req.LastName),
		Enabled:   gocloak.BoolP(true),
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP(req.Password),
				Temporary: gocloak.BoolP(false),
			},
		},
	}

	token, err := gocloakClient.LoginAdmin(c, UsernameAdmin, PasswordAdmin, Realm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log in as admin"})
		return
	}

	_, err = gocloakClient.CreateUser(c, token.AccessToken, Realm, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Đăng nhập và lấy token
func LoginUserHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gửi yêu cầu tới Keycloak để lấy Access Token
	token, err := gocloakClient.Login(c, ClientID, ClientSecret, Realm, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"expires_in":    token.ExpiresIn,
	})
}

// Lấy thông tin người dùng từ Access Token
func GetUserInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	tokenStr := authHeader[len("Bearer "):]
	rpt, err := gocloakClient.RetrospectToken(c, tokenStr, ClientID, ClientSecret, Realm)
	if err != nil || !*rpt.Active {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	userInfo, err := gocloakClient.GetUserInfo(c, tokenStr, Realm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_info": userInfo,
	})
}

func main() {

	// Tải JWKS khi ứng dụng khởi chạy
	config.LoadJWKS()

	server.MustRun(api.NewServer())

}
