package middlewire

/**
 * @Description
 * @Author yingtie
 * @Date 2022/10/8 17:35
 **/

import (
	"gin-blog/conf"
	model "gin-blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(conf.DsnGibBlog), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}) //禁止自动外键
	if err != nil {
		//ssf.Logger.WithField("err", err).Errorln("mysql init error")
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	//debug模式
	//db = db.Debug()
	DB = db
}
