package main

import (
	"flag"
	"fmt"
	"github.com/shudiwsh2009/reservation_thzy_go/models"
	"github.com/shudiwsh2009/reservation_thzy_go/utils"
	"github.com/shudiwsh2009/reservation_thzy_go/workflow"
	"gopkg.in/mgo.v2"
	"time"
)

func main() {
	appEnv := flag.String("app-env", "STAGING", "app environment")
	smsUid := flag.String("sms-uid", "", "sms uid")
	smsKey := flag.String("sms-key", "", "sms key")
	flag.Parse()
	utils.APP_ENV = *appEnv
	utils.SMS_UID = *smsUid
	utils.SMS_KEY = *smsKey
	// 数据库连接
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Errorf("连接数据库失败：%v", err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	models.Mongo = session.DB("reservation_thzy")
	// 时区
	if utils.Location, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		fmt.Errorf("初始化时区失败：%v", err)
		return
	}
	// Reminder
	today := utils.GetToday()
	from := today.AddDate(0, 0, 1)
	to := today.AddDate(0, 0, 2)
	reservations, err := models.GetReservationsBetweenTime(from, to)
	if err != nil {
		fmt.Errorf("获取咨询列表失败：%v", err)
		return
	}
	for _, reservation := range reservations {
		if reservation.Status == models.RESERVATED {
			workflow.SendReminderSMS(reservation)
		}
	}
}
