package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
	"net/http"
)

func ListComment(c echo.Context) error {
	articleID := c.QueryParam("article_id")
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	comments := make([]*Comment, 0, 5)
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(articleID))
		if bucket == nil {
			return nil
		}
		cur := bucket.Cursor()
		count := 0
		for k, v := cur.Last(); k != nil; k, v = cur.Prev() {
			// 目前没有分页，就限制最多100
			count++
			if count > 100 {
				break
			}
			comment := Comment{}
			err = json.Unmarshal(v, &comment)
			if err != nil {
				return err
			}
			comment.Email = maskEmail(comment.Email)
			comments = append(comments, &comment)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, respBody(comments))
}
