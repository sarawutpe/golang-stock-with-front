setup project

setup module
- go mod init main

install packages
- go get github.com/gin-gonic/gin
- go get gorm.io/driver/sqlite
- go get gorm.io/gorm
- go get github.com/dgrijalva/jwt-go
- go get golang.org/x/crypto/bcrypt

deploy
- go build
- go build -ldflags "-s -w"
- using env:   export GIN_MODE=release
- using code:  gin.SetMode(gin.ReleaseMode)

start project (dev mode)
backend
- go run server.go

frontend
- cd webserver
- go run webserver.go