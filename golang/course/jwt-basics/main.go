package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	nowTimeStr := strconv.FormatInt(time.Now().Unix(), 10)
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"login": "admin",
			"stamp": nowTimeStr,
		})
	tokStr, err := tok.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tokStr)

	dToc, err := jwt.Parse(tokStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	mapClaims, ok := dToc.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("token claims are not of type jwt.MapClaims")
		return
	}
	fmt.Println("login:", mapClaims["login"])
	fmt.Println("stamp:", mapClaims["stamp"])
}
