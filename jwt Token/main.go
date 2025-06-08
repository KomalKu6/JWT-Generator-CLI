package main // main package , every go app starts from main package
import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)
// "fmt" : built-in Go package for printing and reading input
// "time":  allows you to manage time
// "github.com/golang-jwt/jwt/v5" : external library used to create and sign jwt tokens

var jwtkey = []byte("my_secret_key")
// secret key for siging the token
 func main(){
	var username, role string // declares two str variable
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter role (admin/user): ")
	fmt.Scanln(&role)
	claims := jwt.MapClaims{
		"username": username,
		"role": role,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println("Error generation token:", err)
		return
	}
	fmt.Println("\nGenerated JWT token:")
	fmt.Println(tokenString)
 }

