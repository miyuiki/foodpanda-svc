package service

import (
	"foodpanda-svc/foodpanda"
	"foodpanda-svc/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckNewOrder(c *gin.Context) {
	client := foodpanda.NewFoodPandaClient()
	token, err := foodpanda.LoadToken()

	if err != nil {
		log.Fatalf("[error] loading token: %v", err)
		c.String(http.StatusInternalServerError, "error")
	}
	retryCnt := 0
	log.Println("[info] request API with token: ", token)
	orderHistory, statusCode, err := retryCheckNewOrder(client, token)
	for err != nil && retryCnt < 3 {
		token, _ := foodpanda.LoadToken()
		orderHistory, statusCode, err = retryCheckNewOrder(client, token)
		retryCnt++
	}
	if err != nil {
		log.Printf("[error] %v, statusCode: %v", err, statusCode)
		c.String(statusCode, "error")
	}
	var orderTrackingList []string
	for _, order := range orderHistory.Data.Items {
		if order.CurrentStatus.Code != 16 && order.CurrentStatus.Code != 13 {
			// code 16 is completed.
			// code 13 is canceld.
			orderTrackingList = append(orderTrackingList, order.OrderCode)
		}
	}
	if len(orderTrackingList) > 0 {
		c.String(http.StatusOK, orderTrackingList[0])
	} else {
		c.String(http.StatusOK, "no new order")
	}
}

func GetOrderByCode(c *gin.Context) {
	client := foodpanda.NewFoodPandaClient()
	code := c.Param("code")
	token, err := foodpanda.LoadToken()

	if err != nil {
		log.Fatalf("[error] loading token: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
	}
	// retryCnt := 0
	log.Println("[info] request API with token: ", token)
	order, statusCode, err := client.GetOrder(token, code)
	if err != nil {
		c.String(statusCode, err.Error())
		return
	}

	c.JSON(statusCode, order)
}

func retryCheckNewOrder(client *foodpanda.FoodPandaClient, token string) (model.OrderHistory, int, error) {
	orderHistory, statusCode, err := client.GetOrderHistory(token)
	if err != nil {
		log.Printf("[warning] %v, statusCode: %v", err, statusCode)
		log.Printf("[info] try to refresh token...")
		foodpanda.RefreshToken(client)
		time.Sleep(5 * time.Second)
	}
	return orderHistory, statusCode, err
}
