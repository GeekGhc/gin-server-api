package models

import (
	"fmt"
	"gin-server-api/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
	DeletedAt  int `json:"deleted_at"`
}

// init the database instance
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// 表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimesStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimesStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedAt` ,`UpdatedAt` when creating
func updateTimesStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if updateTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updateTimeField.IsBlank {
				updateTimeField.Set(nowTime)
			}
		}
	}
}

// deleteCallback will set `DeletedAt` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deleteOnField, hasDeletedOnField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPSATE %v SET %v%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimesStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:uodate_columb"); ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
