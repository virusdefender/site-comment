package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DingDingLink struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	MessageURL string `json:"messageUrl"`
}

type DingDingMessage struct {
	MsgType string        `json:"msgtype"`
	Link    *DingDingLink `json:"link"`
}

type DingDingResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func dingdingPush(comment *Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	msg := DingDingMessage{
		MsgType: "link",
		Link: &DingDingLink{
			Text:       comment.Content,
			Title:      "爸爸，有新评论了",
			MessageURL: baseURL + comment.ArticleID,
		},
	}
	msgStr, _ := json.Marshal(msg)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		dingdingURL,
		bytes.NewReader(msgStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid dingding response code %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	data := DingDingResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	if data.ErrCode != 0 {
		return fmt.Errorf("dingding error: %s", data.ErrMsg)
	}
	return nil
}
