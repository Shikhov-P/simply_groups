package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/nbutton23/zxcvbn-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	utils "../utils"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Groups []Group `gorm:"ForeignKey:AccountRefer";json:"user_groups"'`
	Token     string `json:"token";sql:"-"`
}

//TODO: display all errors if more than one are present
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !(utils.IsAlphaNumeric(account.FirstName) && utils.IsAlphaNumeric(account.LastName)){
		return utils.Message(false, "First and last names must contain letters only."), false
	}

	usernameRegexp := regexp.MustCompile(`^\w+$`)
	if !usernameRegexp.MatchString(account.Username){
		return utils.Message(false, "Userame must contain alphanumeric letters only."), false
	}

	emailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegexp.MatchString(account.Email){
		return utils.Message(false, "Wrong email format."), false
	}

	passwordCheck := zxcvbn.PasswordStrength(account.Password, []string{account.Email})
	if passwordCheck.Score < 4 {
		return utils.Message(false, "Use a stronger password."), false
	}

	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error."), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email is already in use."), false
	}

	return utils.Message(false, "Account is valid."), true
}

func (account *Account) Create() (map[string]interface{}) {
	if resp, isValidated := account.Validate(); !isValidated {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	account.Email = strings.ToLower(account.Email)

	GetDB().Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Account creation error. Please, retry.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""

	response := utils.Message(true, "Account has been created.")
	response["account"] = account
	return response
}


func Login(email, password string) (map[string]interface{}) {
	account := &Account{}
	email = strings.ToLower(email)
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email or password is incorrect")
		}
		return utils.Message(false, "Connection error.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil  {
		return utils.Message(false, "Email or password is incorrect")
	}

	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	response := utils.Message(true, "Logged in.")
	response["account"] = account
	return response
}

func GetUser(userId uint) *Account {
	account := &Account{}
	GetDB().Table("accounts").Where("id = ?", userId).First(account)
	if account.Email == "" {
		return nil
	}

	account.Password = ""
	return account
}