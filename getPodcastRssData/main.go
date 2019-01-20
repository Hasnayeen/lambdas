package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type APIResponse struct {
	Channel struct {
		Items []struct {
			Title     string    `xml:"title" json:"title"`
			Duration  string    `xml:"duration" json:"duration"`
			PubDate   string    `xml:"pubDate" json:"pubDate"`
			EpisodeNo string    `xml:"episode" json:"episode"`
			AudioLink UrlString `xml:"enclosure" json:"audio"`
		} `xml:"item" json:"items"`
	} `xml:"channel" json:"channel"`
}

type UrlString struct {
	URL string `xml:"url,attr" json:"url"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

	param := request.QueryStringParameters

	resp, err := http.Get(param["url"])
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	}

	response := &APIResponse{}
	xml.Unmarshal(body, response)

	result, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{Body: string(result), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
