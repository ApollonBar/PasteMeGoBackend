/*
@File: data.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-18 14:40  Lucien     1.0         Init
*/
package data

import (
	"fmt"
	"github.com/LucienShui/PasteMe/GoBackend/data/permanent"
	"github.com/LucienShui/PasteMe/GoBackend/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	username = "username"
	password = "password"
	network  = "tcp"
	server   = "mysql"
	port     = 3306
	database = "pasteme"
)

func format(
	username string,
	password string,
	network string,
	server string,
	port int,
	database string) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=True&loc=Local", username, password, network, server, port, database)
}

type Paste struct {
	Key      string `json:"key"`
	Lang     string `json:"lang"`
	Content  string `json:"content"`
	Password string `json:"password"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", format(username, password, network, server, port, database))
	if err != nil {
		panic(err)
	}
	// db = db.Debug() // DEBUG
	if !db.HasTable(&permanent.Permanent{}) {
		if err := db.Set(
			"gorm:table_options",
			"ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=100",
			).CreateTable(&permanent.Permanent{}).Error; err != nil {
			panic(err)
		}
	}
}

func Insert(paste Paste) (string, error) {
	var err error
	if paste.Key == "read_once" {
		// TODO
	} else if paste.Key == "" { // permanent
		if paste.Key, err = func() (string, error) {
			key, err := permanent.Insert(db, paste.Lang, paste.Content, paste.Password)
			if err != nil {
				return "", err
			}
			return util.Uint2string(key), nil
		}(); err != nil {
			return "", err
		}
	} else { // temporary
		// TODO
	}
	return paste.Key, err
}

func Query(key string) (Paste, error) {
	table, err := util.ValidChecker(key)
	paste := Paste{}
	if err != nil {
		return paste, err
	}
	if table == "permanent" {
		object, err := permanent.Query(db, util.String2uint(key))
		if err != nil {
			return paste, err
		}
		return Paste{
			Key: util.Uint2string(object.Key),
			Lang: object.Lang,
			Content: object.Content,
			Password: object.Password}, err
	} else { // temporary
		// TODO
	}
	return paste, err
}

func Delete(key string) error {
	table, err := util.ValidChecker(key)
	if err != nil {
		return err
	}
	if table == "permanent" {
		return permanent.Delete(db, util.String2uint(key))
	} else { // temporary
		// TODO
	}
	return err
}
