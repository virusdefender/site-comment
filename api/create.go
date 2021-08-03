package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func insertDB(comment *Comment) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(comment.ArticleID))
		if err != nil {
			return err
		}
		id, _ := bucket.NextSequence()
		comment.ID = int(id)
		val, _ := json.Marshal(comment)
		return bucket.Put(itob(comment.ID), val)
	})
	return err
}

func CreateComment(c echo.Context) error {
	referer := c.Request().Header.Get("referer")
	if referer == "" {
		return errors.New("referer 为空")
	}
	u, err := url.Parse(referer)
	if err != nil {
		return err
	}
	for _, item := range refererAllowDomains {
		if item == u.Hostname() {
			goto refererOK
		}
	}
	return fmt.Errorf("%s 域名不在允许范围内", u.Hostname())

refererOK:
	comment := &Comment{}
	err = c.Bind(comment)
	if err != nil {
		return err
	}
	comment.IP = c.RealIP()
	comment.TimeStamp = time.Now().Unix()
	if !articleIDRegex.MatchString(comment.ArticleID) {
		return errors.New("article id 格式错误")
	}
	if strings.TrimSpace(comment.Username) == "" || strings.TrimSpace(comment.Content) == "" {
		return errors.New("用户名或者评论为空")
	}
	err = insertDB(comment)
	if err != nil {
		return err
	}
	err = dingdingPush(comment)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, respBody(comment))
}
