#!/bin/bash
echo "#update git"
go get -u github.com/shudiwsh2009/reservation_thzy_go
go get -u github.com/shudiwsh2009/reservation_thzy_go_reminder
go get -u github.com/shudiwsh2009/reservation_thzy_go_timetable

echo "#deploy website"
supervisorctl stop reservation_thzy_go
sleep 5
cd $GOPATH/src/github.com/shudiwsh2009/reservation_thzy_go
go build
supervisorctl start reservation_thzy_go

echo "#deploy reminder"
cd $GOPATH/src/github.com/shudiwsh2009/reservation_thzy_go_reminder
go build
echo "#deploy timetable"
cd $GOPATH/src/github.com/shudiwsh2009/reservation_thzy_go_timetable
go build

echo "restart cron"
service cron restart

echo "###Deployment Completed"
