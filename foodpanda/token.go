package foodpanda

import (
	"fmt"
	"log"
	"os"
)

const (
	tokenPath = "./token.txt"
)

func RefreshToken(c *FoodPandaClient) {
	username := os.Getenv("FOODPANDA_USERNAME")
	password := os.Getenv("FOODPANDA_PASSWORD")
	xdevice := os.Getenv("FOODPANDA_XDEVICE")
	if username == "" || password == "" || xdevice == "" {
		log.Fatal("[error] FOODPANDA_USERNAME, FOODPANDA_PASSWORD, and FOODPANDA_XDEVICE must be set")
	}
	loginInfo, loginCookie, statusCode, err := c.Login(username, password, xdevice)
	if err != nil {
		log.Fatalf("[error] %v, statusCode: %v", err, statusCode)
	}
	msg := fmt.Sprintf("Login successful: UserID = %v, Token = %v, StatusCode = %v", loginInfo.UserID, loginCookie.Token, statusCode)
	log.Println(msg)
	if err := SaveToken(loginCookie.Token); err != nil {
		log.Fatalf("[error] %v", err)
	}
}

func LoadToken() (string, error) {
	token, err := os.ReadFile(tokenPath)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func SaveToken(token string) error {
	err := os.WriteFile(tokenPath, []byte(token), 0644)
	if err != nil {
		return err
	}
	return nil
}
