package v1

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MessageParams struct {
	Bot     *linebot.Client
	Event   *linebot.Event
	Message *linebot.TextMessage
}

var bot *linebot.Client
var err error

func init() {
	//獲取 .env檔
	bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

//line bot 回應邏輯
func LineReply(c *gin.Context) {

	events, err := bot.ParseRequest(c.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {

			//針對接受到的文字訊息進行解析回應
			case *linebot.TextMessage:
				if message.Text == "yoyo" {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("check it out")).Do()
				}
				params := MessageParams{
					Bot:     bot,
					Event:   event,
					Message: message,
				}
				params.HelperReply()

			}
		}
	}
}
