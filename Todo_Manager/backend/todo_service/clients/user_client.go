package clients

import (
	"fmt"
	"log"
	"github.com/go-resty/resty/v2"
)

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}


var apiClient = createRestClient()

func createRestClient() *resty.Client {
	client := resty.New()

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		log.Printf("🌐 [RESTY Запрос] -> %s %s", req.Method, req.URL)
		return nil
	})


	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		log.Printf("[RESTY Ответ] <- Статус: %d | Время: %v | URL: %s", 
			resp.StatusCode(), 
			resp.Time(), 
			resp.Request.URL,
		)
		
		if resp.IsError() {
			log.Printf("[RESTY Ошибка] Тело: %s", resp.String())
		}
		
		return nil
	})

	return client
}


func GetUserByID(userID uint) (*UserResponse, error) {
	var userResp UserResponse

	// URL нашего AuthService, который работает на 8081 порту
	url := fmt.Sprintf("http://auth_service:8081/users/%d", userID)
	
	resp, err := apiClient.R().
		SetResult(&userResp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("ошибка сети: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("AuthService вернул ошибку: %d", resp.StatusCode())
	}

	return &userResp, nil
}