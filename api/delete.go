package api

import (
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
	"net/http"
	"strconv"
)

func DeleteComment(c echo.Context) error {
	articleID := c.QueryParam("article_id")
	idStr := c.QueryParam("id")
	token := c.QueryParam("token")
	if token != mgmtToken {
		return c.JSON(http.StatusUnauthorized, respBody(nil))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		// -1 等于全删
		if id == -1 {
			return tx.DeleteBucket([]byte(articleID))
		}
		bucket := tx.Bucket([]byte(articleID))
		if bucket == nil {
			return nil
		}
		return bucket.Delete(itob(id))
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, respBody(nil))
}
