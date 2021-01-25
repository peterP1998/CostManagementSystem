package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/peterP1998/CostManagementSystem/models"
	"github.com/peterP1998/CostManagementSystem/views"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AccountService struct {
	UserServiceWired UserService
}

var jwtKey = []byte("my_secret_key")

func (account *AccountService) ParseToken(tokenString string) (string, bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return "", false, err
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.Username, claims.Admin, nil
	} else {
		return "", false, err
	}
}
func (account AccountService) CheckAuthBeforeOperate(r *http.Request, w http.ResponseWriter) *http.Cookie {
	token, err := checkForAuthorization(r)
	authorizationResponses(w, err, r)
	return token
}
func (account AccountService) Login(password string, username string, w http.ResponseWriter) {
	user, err := account.UserServiceWired.SelectUserByName(username)
	if err != nil {
		views.CreateView(w, "static/templates/accounts/index.html", map[string]interface{}{"messg": "Wrong username or password!"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		views.CreateView(w, "static/templates/accounts/index.html", map[string]interface{}{"messg": "Wrong username or password!"})
		return
	}
	err = createAndConfigureToken(user, w)
	if err != nil {
		views.CreateView(w, "static/templates/accounts/index.html", map[string]interface{}{"messg": "Something went wrong please try againg!"})
	}
}
func createAndConfigureToken(user models.User, w http.ResponseWriter) error {
	claims, expirationTime := configureToken(user)
	tokenString, err := createToken(claims)
	if err != nil {
		return err
	}
	setCookieToken(tokenString, expirationTime, w)
	return nil
}
func configureToken(user models.User) (*models.Claims, time.Time) {
	expirationTime := time.Now().Add(180 * time.Minute)
	claims := &models.Claims{
		Username: user.Name,
		Admin:    user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	return claims, expirationTime
}
func createToken(claims *models.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}
func setCookieToken(tokenString string, expirationTime time.Time, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
func checkForAuthorization(r *http.Request) (*http.Cookie, error) {
	toekn, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}
	return toekn, nil
}
func authorizationResponses(w http.ResponseWriter, err error, r *http.Request) {
	if err != nil {
		http.Redirect(w, r, "/api/login", http.StatusSeeOther)
	}
}
