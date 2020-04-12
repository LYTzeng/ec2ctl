package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// https://developers.mattermost.com/integrate/slash-commands/
type mattermostBody struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
	Username     string `json:"username"`
}

func printRemoteCMD(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, _ := url.ParseQuery(event.Body) //x-www-form-urlencoded body
	// body is a map, containing the following keys:
	// channel_id
	// channel_name
	// command
	// response_url
	// team_domain
	// team_id
	// text
	// token
	// trigger_id
	// user_id
	// user_name
	fmt.Println(body)
	// Atuhorization
	token := os.Getenv("MATTERMOST_TOKEN")
	if body["token"][0] != token {
		return toUnauthorized()
	}

	argString := body["text"][0]

	usage := "Usage:\n\t" +
		"ec2ctl version\n\t" +
		"ec2ctl (stop | start) -i=<id>\n" +
		"Options:\n\t" +
		"-i=<id>\tSpecify an instance ID."

	if argString == "" {
		return toResponse(usage, body["text"][0], body["user_name"][0])
	}
	args := strings.Fields(argString)
	osArgs := []string{"ec2ctl"}
	os.Args = append(osArgs, args...)
	fmt.Println(os.Args)

	var responseString string

	cmds := []Runner{
		ShowVersion(),
		StopEC2(),
		StartEC2(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			responseString = cmd.Run() // Here the program takes action, eg. Stop or Start instances.
			return toResponse(responseString, body["text"][0], body["user_name"][0])
		}
	}

	return toResponse(usage, body["text"][0], body["user_name"][0])
}

func toResponse(responseString, args, user string) (events.APIGatewayProxyResponse, error) {
	bytesBody, err := json.Marshal(&mattermostBody{
		ResponseType: "in_channel",
		Text: "```\n〉ec2ctl " + args + "\n" +
			responseString + "\n```",
		Username: user,
	})

	mattermostBodyString := string(bytesBody[:])
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: mattermostBodyString,
	}

	return response, err
}

func toUnauthorized() (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 401,
		Headers: map[string]string{
			"content-type": "text/plain",
		},
		Body: "Unauthorized. What's your problem??",
	}
	return response, nil
}

func main() {
	lambda.Start(printRemoteCMD)
}
