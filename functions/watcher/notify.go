package watcher

import "github.com/slack-go/slack"

func Notify(message string) error {
	// アクセストークンを使用してクライアントを生成する
	c := slack.New(mustGetenv("SLACK_TOKEN"))

	_, _, err := c.PostMessage(mustGetenv("SALCK_CHANNEL"), slack.MsgOptionText(message, true))
	if err != nil {
		return err
	}
	return nil
}
