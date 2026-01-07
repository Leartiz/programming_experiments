package n3

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Exec() {
	nowTimeStr := strconv.FormatInt(time.Now().Unix(), 10)
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"login": "admin",
			"stamp": nowTimeStr,
		})
	tokStr, err := tok.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		return // abort without decoding!
	} else {
		fmt.Println(tokStr)
	}

	// ***

	{
		dToc, err := jwt.Parse(tokStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}

			return []byte("secret"), nil
		})

		if err != nil {
			fmt.Println(err)
		} else {
			mapClaims, ok := dToc.Claims.(jwt.MapClaims)
			if !ok {
				fmt.Println("token claims are not of type *jwt.MapClaims")
			} else {
				login := mapClaims["login"]
				stamp := mapClaims["stamp"]

				fmt.Println("login:", login)
				fmt.Println("stamp:", stamp)
			}
		}
	}
}
