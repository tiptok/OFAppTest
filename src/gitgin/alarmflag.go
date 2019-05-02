package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

//127.0.0.1:8081/alarmflag/update|save|delete
func AlarmFlagSave(g *gin.Context) {
	var alarm AlarmFlag
	err := g.ShouldBind(&alarm)
	rspMsg := ""
	if err == nil {
		log.Printf("recv data:%v", alarm)
		var created AlarmFlag
		db.Table("bas_AlarmFlag").FirstOrCreate(&created, alarm)
		if alarm.AlarmFlag != 0 {
			rspMsg = fmt.Sprintf("success add new row-> AlarmFlag:%d AlarmName:%v", created.AlarmFlag, created.AlarmName)
			ServerSuccess(g, rspMsg)
			return
		}
	} else {
		g.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	ServerSuccess(g, "success")
}

func AlarmFlagDelete(g *gin.Context) {
	var alarm AlarmFlag
	err := g.ShouldBind(&alarm)
	rspMsg := ""
	if err == nil {
		log.Printf("recv data:%v", alarm)
		var del AlarmFlag
		db.Table("bas_AlarmFlag").Where("AlarmFlag=?", alarm.AlarmFlag).Delete(&del)

		if alarm.AlarmFlag == del.AlarmFlag {
			rspMsg = fmt.Sprintf("success delete AlarmFlag:%d AlarmName:%s", alarm.AlarmFlag, alarm.AlarmName)
		} else {
			//rspMsg = fmt.Sprintf("success delete AlarmFlag:%d AlarmName:%s", del.AlarmFlag, del.AlarmName)
			rspMsg = fmt.Sprintf("success delete AlarmFlag:%d AlarmName:%s", alarm.AlarmFlag, alarm.AlarmName)
		}
		ServerSuccess(g, rspMsg)
		return
	}
	rspMsg = err.Error()
	ServerSuccess(g, rspMsg)
}

//AlarmFlagUpdate  update table bas_AlarmFlag
func AlarmFlagUpdate(g *gin.Context) {
	var alarm AlarmFlag
	err := g.ShouldBind(&alarm)
	rspMsg := ""
	if err == nil {
		db.Table("bas_AlarmFlag").Where("AlarmFlag=?", alarm.AlarmFlag).Update(alarm)
		var first AlarmFlag
		db.Table("bas_AlarmFlag").Where("AlarmFlag=?", alarm.AlarmFlag).First(&first)
		if first.AlarmFlag != 0 {
			rspMsg = fmt.Sprintf("success update AlarmFlag:%d AlarmName:%s", first.AlarmFlag, first.AlarmName)
		}
		ServerSuccess(g, rspMsg)
		return
	}
	rspMsg = err.Error()
	ServerSuccess(g, rspMsg)
}

func ServerSuccess(g *gin.Context, msg string) {
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": msg,
	})
}
