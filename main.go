package main

import (
	"fmt"
	"log"
	"net/http"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("./db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// wait to close the database after opening
	////////////////////////////////////////////////////

	////////////////////////////////////////////////////

	// gin code here

	r := gin.Default() // default config for router

	////////////////////////////////////////////////////
	// endpoints below

	type TagMsg struct {
		Tag string `json:"tag"`
		Msg string `json:"msg"`
	}

	r.GET("/tags/:tag", func(c *gin.Context) {
		tag := c.Param("tag")
		err := db.View(func(txn *badger.Txn) error {
			item, err := txn.Get([]byte(tag))
			if err != nil {
				return err
			}
			item.Value(func(val []byte) error {
				c.String(http.StatusOK, "%s", val)
				return nil
			})
			return nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
	})

	r.POST("/tags", func(c *gin.Context) {
		var tagMsg TagMsg
		c.BindJSON(&tagMsg)
		err := db.Update(func(txn *badger.Txn) error {
			return txn.Set([]byte(tagMsg.Tag), []byte(tagMsg.Msg))
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, gin.H{"created_tag": tagMsg.Tag, "created_msg": tagMsg.Msg})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	////////////////////////////////////////////////////
	fmt.Println("http://localhost:8080")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
