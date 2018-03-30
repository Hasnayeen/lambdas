package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Payload struct {
		References struct {
			Post map[string]Postdata
		} `json:"references"`
	} `json:"payload"`
}

type Postdata struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	FirstPublishedAt int64  `json:"firstPublishedAt"`
	UniqueSlug       string `json:"uniqueSlug"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

	resp, err := http.Get("https://medium.com/@searching.nehal/latest?format=json")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	data := string(body)
	convertedData := strings.Replace(data, "])}while(1);</x>", "", 1)

	var r Response

	json.Unmarshal([]byte(convertedData), &r)

	response, _ := json.Marshal(r.Payload.References.Post)

	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
