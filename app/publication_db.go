package app

import "github.com/jinzhu/gorm"

const ConnectionString = "root:root@tcp(mysql:3306)/rs_db?parseTime=true&loc=UTC&charset=utf8"

var Db *gorm.DB
