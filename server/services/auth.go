package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"progateway.ru/m/v2/models"
	"time"
)

// SignUp Регистрация
func SignUp(c *gin.Context) {
	var input models.User

	// Парсинг JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// Проверка на существующего пользователя
	if err := models.DB.Where("email=?", input.Email).First(&input).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"msg": "Электронная почта уже занята"})
		return
	}

	// Хэширования пароля с помощью bcrypt
	hash, err := hashing(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Ошибка сервера."})
		return
	}
	input.Password = hash
	input.Role = 0

	// Сохранение пользователя в БД
	models.DB.Create(&input)

	emailCheck := models.EmailCheck{ID: input.ID, UUID: uuid.New().String()}
	models.DB.Create(&emailCheck)

	SendEmailUUID(input.Email, emailCheck.UUID)

	c.JSON(http.StatusOK, gin.H{"msg": "Мы отправили письмо с подтверждением на вашу электронную почту. Вы сможете зайти на аккант только после подтверждения."})
}

// SignIn Аутентификация (вход)
func SignIn(c *gin.Context) {
	var input models.SignIn
	var user models.User

	// Парсинг JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// Поиск пользователя по Email
	if err := models.DB.Where("email=?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Неверная почта или пароль."})
		return
	}

	if user.Role == 0 {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Подтвердите адрес электронной почты"})
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Неверная почта или пароль."})
		return
	}

	var token models.Tokens

	token.UserID = user.ID
	token.Token = uuid.New().String()
	token.IP = c.ClientIP()
	token.Browser = c.Request.UserAgent()

	models.DB.Create(&token)

	// Установка refresh токена в файл cookie
	c.SetCookie("refresh_token", token.Token, 60*60*24*30, "/api", os.Getenv("COOKIE_DOMAIN"), secure(), true)
	c.JSON(http.StatusOK, gin.H{"access": createToken(user, 15)})
}

func Logout(c *gin.Context) {

	var token models.Tokens
	refresh, _ := c.Cookie("refresh_token")

	if err := models.DB.Where("token=?", refresh).First(&token).Error; err == nil {
		models.DB.Delete(&token)
	}

	c.SetCookie("refresh_token", "", -1, "/api", os.Getenv("COOKIE_DOMAIN"), secure(), true)
	c.JSON(http.StatusOK, gin.H{"msg": "Успешно"})
}

// hashing хэширование пароля с помощью bcrypt
func hashing(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// createToken Создание JWT
func createToken(user models.User, minutes int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Minute * time.Duration(minutes)).Unix(),
	})

	jwtToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return jwtToken
}

func secure() bool {
	if os.Getenv("COOKIE_SECURE") == "false" {
		return false
	}
	return true
}

func refresh(c *gin.Context) (uuid.UUID, error) {
	tokenRefresh, err := c.Cookie("refresh_token")
	if err != nil {
		return uuid.Nil, err
	}

	var token models.Tokens

	if err := models.DB.Where("token=?", tokenRefresh).First(&token).Error; err != nil {
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		return uuid.Nil, err
	}

	var user models.User

	models.DB.Where("id=?", token.UserID).First(&user)

	access := createToken(user, 15)
	token.UserID = user.ID
	token.Token = uuid.New().String()
	token.IP = c.ClientIP()
	token.Browser = c.Request.UserAgent()

	models.DB.Model(&token).Updates(token)

	c.SetCookie("refresh_token", token.Token, 60*60*24*30, "/api", os.Getenv("COOKIE_DOMAIN"), secure(), true)
	c.Header("access", access)
	return user.ID, nil
}

func CheckToken(c *gin.Context, role uint) (uuid.UUID, bool) {
	type MyCustomClaims struct {
		ID    uuid.UUID `json:"userid"`
		Email string    `json:"email"`
		Role  uint      `json:"role"`
		jwt.StandardClaims
	}

	tokenParse, err := jwt.ParseWithClaims(c.Request.Header.Get("access"), &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err == nil {
		if parsing, ok := tokenParse.Claims.(*MyCustomClaims); ok && tokenParse.Valid { // если access токен валидный
			var user models.User
			if err := models.DB.Where("id=?", parsing.ID).First(&user).Error; err != nil { // поиск роли юзера
				if user.Role != parsing.Role { // если роль не совпадает
					id, err := refresh(c) // пробуем рефрешить access токен
					if err != nil {       // если не получилось - выбиваем юзера
						return uuid.Nil, false
					}
					if user.Role < role { // если роль ниже необходимой - выбиваем юзера
						return uuid.Nil, false
					}
					return id, true // успех
				}
			}

		}
	}

	id, err := refresh(c) // если токен отсутствует, то производим рефреш
	if err != nil {       // при ошибке выбиваем юзера
		return uuid.Nil, false
	}
	return id, true // успех
}

func Activate(c *gin.Context) {
	emailCheck := models.EmailCheck{}
	if err := models.DB.Where("uuid=?", c.Param("uuid")).First(&emailCheck).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Подтверждение не существует или недействительно!"})
		return
	}

	user := models.User{}
	userUpdate := models.User{Role: 1}

	if err := models.DB.Where("id=?", emailCheck.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Аккаунт не найден!"})
		models.DB.Delete(&emailCheck)
		return
	}

	models.DB.Model(&user).Updates(userUpdate)
	models.DB.Delete(&emailCheck)

	c.JSON(http.StatusOK, gin.H{"msg": "Аккаунт успешно активирован!"})
}
