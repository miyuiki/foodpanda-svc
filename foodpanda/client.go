package foodpanda

import (
	"bytes"
	"encoding/json"
	"errors"
	"foodpanda-svc/model"
	"io"
	"net/http"
	"net/url"
)

type FoodPandaClient struct {
	LoginURL   string
	APIBaseURL string
}

func NewFoodPandaClient() *FoodPandaClient {
	return &FoodPandaClient{
		LoginURL:   "https://www.foodpanda.com.tw/login/new/api/login",
		APIBaseURL: "https://tw.fd-api.com/api/v5",
	}
}

func (c *FoodPandaClient) Login(username string, password string, xdevice string) (model.LoginInfo, model.LoginCookie, int, error) {
	body := map[string]string{
		"username": username,
		"password": password,
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"x-device":     xdevice,
		"x-otp-method": "EMAIL",
		"Accept":       "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", c.LoginURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return model.LoginInfo{}, model.LoginCookie{}, 0, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.LoginInfo{}, model.LoginCookie{}, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.LoginInfo{}, model.LoginCookie{}, resp.StatusCode, errors.New("login failed")
	}

	cookies := resp.Cookies()
	loginCookie := model.LoginCookie{}
	for _, cookie := range cookies {
		if cookie.Name == "device_token" {
			loginCookie.DeviceToken = cookie.Value
		}
		if cookie.Name == "refresh_token" {
			loginCookie.RefreshToken = cookie.Value
		}
		if cookie.Name == "token" {
			loginCookie.Token = cookie.Value
		}
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.LoginInfo{}, model.LoginCookie{}, 0, err
	}
	var loginInfo model.LoginInfo
	if err := json.Unmarshal(respBody, &loginInfo); err != nil {
		return model.LoginInfo{}, model.LoginCookie{}, 0, err
	}
	return loginInfo, loginCookie, resp.StatusCode, nil
}

func (c *FoodPandaClient) GetOrderHistory(token string) (model.OrderHistory, int, error) {
	params := url.Values{}
	params.Add("include", "order_products,order_details")
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
		"X-Fp-Api-Key":  "volo",
		"Accept":        "application/json",
		"Connection":    "keep-alive",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	}
	req, err := http.NewRequest("GET", c.APIBaseURL+"/orders/order_history?"+params.Encode(), nil)
	if err != nil {
		return model.OrderHistory{}, 0, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.OrderHistory{}, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.OrderHistory{}, resp.StatusCode, errors.New("failed to get order history")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.OrderHistory{}, 0, err
	}
	var orderHistory model.OrderHistory
	if err := json.Unmarshal(body, &orderHistory); err != nil {
		return model.OrderHistory{}, 0, err
	}
	return orderHistory, resp.StatusCode, nil
}

func (c *FoodPandaClient) GetOrder(token string, code string) (model.Order, int, error) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
		"X-Fp-Api-Key":  "volo",
		"Accept":        "application/json",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	}
	req, err := http.NewRequest("GET", c.APIBaseURL+"/tracking/orders/"+code, nil)
	if err != nil {
		return model.Order{}, 0, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Order{}, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Order{}, resp.StatusCode, errors.New("failed to get order by code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Order{}, 0, err
	}
	var order model.Order
	if err := json.Unmarshal(body, &order); err != nil {
		return model.Order{}, 0, err
	}
	return order, resp.StatusCode, nil
}
