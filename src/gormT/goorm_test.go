package gormT

import (
	"testing"

	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type biz_AuthLog struct {
// 	gorm.Model
// 	RCId       int
// 	UserId     string
// 	Level      int
// 	UpdateTime time.Time
// 	Operator   string
// }

var db *gorm.DB

func TestQuery(t *testing.T) {
	var err error
	db, err = gorm.Open("mssql", "{SQL Server};server=127.0.0.1;uid=sa;pwd=123456;database=MZC;")
	if err != nil {
		fmt.Println("Open DB Error:", err.Error())
	}
	t.Log("End")
}

type AlarmFlag struct {
	AlarmFlag  int    `gorm:"column:AlarmFlag"`
	AlarmName  string `gorm:"column:AlarmName"`
	AlarmType  int    `gorm:"column:AlarmType"`
	Remark     string `gorm:"column:Remark"`
	AlarmOrder int    `gorm:"column:AlarmOrder"`
}

func (t AlarmFlag) TableName() string {
	return "bas_AlarmFlag"
}

func TestMysqlQuery(t *testing.T) {
	var err error
	db, err = gorm.Open("mysql", "root:admin@(127.0.0.1:3306)/TopDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Open DB Error:", err.Error())
	} else {
		fmt.Println("db connect:", db)
	}
	fmt.Println("has table bas_AlarmFlag:", db.HasTable("bas_AlarmFlag"))
	alarmItem := AlarmFlag{}
	db.Table("bas_AlarmFlag").First(&alarmItem)
	fmt.Println("AlarmFlag Item:", alarmItem)
	t.Log("End")
}
