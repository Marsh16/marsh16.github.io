package main

// go mod init vp_week11_echo
// GO111MODULE=on go get github.com/labstack/echo/v4
// go get github.com/tkanos/gonfig
// go get -u github.com/go-sql-driver/mysql
//go get golang.org/x/crypto/bcrypt
//go get github.com/dgrijalva/jwt-go
//go get github.com/labstack/echo/v4/middleware
//go get github.com/go-playground/validator
import (
	// "fmt"
	"log"
	"net/http"
	"archcalculator.github.io/db"
	// "archcalculator.github.io/helpers"
	"archcalculator.github.io/routes"
)

// const (
//     productionBaseURL = "https://marsh16.github.io/" // Production URL (hardcoded)
// )

func main() {
    db.Init()
    e := routes.Init()

    // Start the server using the production URL
    log.Fatal(http.ListenAndServe(":8080", e))
}