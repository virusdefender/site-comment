package api

import (
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
	"io/ioutil"
	"net/http"
)

func BackupDB(c echo.Context) error {
	token := c.QueryParam("token")
	if token != mgmtToken {
		return c.JSON(http.StatusUnauthorized, respBody(nil))
	}
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		f, err := ioutil.TempFile("", "comment_*.db")
		if err != nil {
			return err
		}
		_, err = tx.WriteTo(f)
		if err != nil {
			return err
		}
		c.Response().Header().Set("Content-Disposition", "attachment; filename=comment_backup.db")
		return c.File(f.Name())
	})
	return err
}
