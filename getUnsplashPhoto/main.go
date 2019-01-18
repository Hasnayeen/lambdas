package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type image []struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Description string `json:"description"`
	User        struct {
		Name         string `json:"name"`
		PortfolioURL string `json:"portfolio_url"`
	} `json:"user"`
	Urls struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

	resp, err := http.Get("https://api.unsplash.com/users/nhasnayeen/likes?client_id=your_unsplash_api_client_id")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
