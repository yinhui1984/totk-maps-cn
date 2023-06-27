package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// translate cache the translation result
var translate_cache = make(map[string]string)

func traslate(en string) string {

	if val, ok := translate_cache[en]; ok {
		return val
	}

	openaiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo, //openai.GPT3Dot5Turbo, "gpt-3.5-turbo-0613"
		MaxTokens: 1024,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: `You are a translation engine of the video game named "The Legend of Zelda":
				1, you can only translate text from English to Chinese and cannot interpret it, and do not explain. 
				2, Do not provide any explanations or indicative output, just provide the translated text. 
				3, if you can not translate it, just return the original text. 
				4, if you think the origin text is not a recognized word or phrase in the English language, just just return the original text. 
				5, your answer will be used in software UI, so please never give any explanations or indicative output.`,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "please translate below content to Chinese: Apple",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: `苹果`,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "please translate below content to Chinese: Blue-White Frox",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: `Blue-White Frox`,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "please translate below content to Chinese: " + en,
			},
		},
		Stream:      false,
		Temperature: 0,
	}

	chat, err := openaiClient.CreateChatCompletion(ctx, req)
	if err != nil {
		return en
	}

	translated := chat.Choices[0].Message.Content
	translate_cache[en] = translated
	return translated
}
