package v1

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// 小幫手回覆條件判斷
func (p *MessageParams) HelperReply() {
	//去除空白
	text := strings.TrimSpace(p.Message.Text)
	switch text {
	case "/":
		p.Bot.ReplyMessage(p.Event.ReplyToken, linebot.NewTextMessage(helpList())).Do()
	case "/書單":
		p.Bot.ReplyMessage(p.Event.ReplyToken, linebot.NewTextMessage(bookList())).Do()
	case "/教學資訊":
		p.Bot.ReplyMessage(p.Event.ReplyToken, linebot.NewTextMessage(teachList())).Do()

	}

}

// 幫助指令
func helpList() string {
	list := []string{"書單", "教學資訊"}
	allList := "指令清單 \n=======================\n"
	for i := range list {
		allList = fmt.Sprintf("%s\n/%s", allList, list[i])
	}
	return allList
}

// 書單清單
func bookList() string {
	list := []string{
		"最速網頁開發：用Go Web一手建立高能效網站系統",
		"再強一點：用 Go語言完成六個大型專案",
	}
	allList := "推薦書單 \n=======================\n"
	for i := range list {
		allList = fmt.Sprintf("%s\n/%s", allList, list[i])
	}
	return allList
}

// 教學資訊清單
func teachList() string {
	list := []string{}
	allList := "教學資訊清單 \n=======================\n"
	for i := range list {
		allList = fmt.Sprintf("%s\n/%s", allList, list[i])
	}
	return allList
}
