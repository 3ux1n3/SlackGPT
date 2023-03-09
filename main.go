package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strings"

	"github.com/shomali11/slacker"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	bot := slacker.NewClient(os.Getenv("SLACK_AUTH_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	chatContext := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful assistant living in our slack chat where we will ask you for assistance",
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "this test is just to set the context please don't reply : At PartnersWire, we are dedicated to being the premier provider of software development services to professional service businesses. Our mission is to help these businesses improve their online presence and streamline their internal processes, so they can focus on delivering value to their clients and growing their businesses. from now on you are a member of our slack team",
			},
		},
	}
	bot.DefaultCommand(func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
		if len(chatContext.Messages) >= 10 {
			chatContext.Messages = chatContext.Messages[0:8]
		}
		message := botCtx.Event().Text
		if strings.HasPrefix(message, "!bot") {
			fmt.Println("Received message:", message)
			message = strings.ReplaceAll(message, "!bot", "")
			chatContext.Messages = append(chatContext.Messages,
				openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				})

			fmt.Println("generating Response")
			resp, err := client.CreateChatCompletion(
				context.Background(),
				chatContext,
			)

			if err != nil {
				return
			}
			chatContext.Messages = append(chatContext.Messages,
				openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: resp.Choices[0].Message.Content,
				})
			err = response.Reply(resp.Choices[0].Message.Content)
			fmt.Println("response sent message:", resp.Choices[0].Message.Content)
			if err != nil {
				return
			}

		}

	})

	err = bot.Listen(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
