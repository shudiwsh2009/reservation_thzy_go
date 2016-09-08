#!/bin/bash
echo "#update git"
cd $GOPATH/src/github.com/shudiwsh2009/reservation_thzy_go
git reset --hard
git fetch $1
git checkout $1/$2

echo "#deploy website"
cd ~/thzyfzzx_go
go install github.com/shudiwsh2009/reservation_thzy_go/server
kill -9 $(lsof -t -i:8080)
sleep 5
now=$(date +"%Y_%m_%d_%T")
mv ~/thzyfzzx_go/log/server.log ~/thzyfzzx_go/log/server-${now}.log
cp -f $GOPATH/bin/server ./server/server.run
cp -rf $GOPATH/src/github.com/shudiwsh2009/reservation_thzy_go/templates ./templates
cp -rf $GOPATH/src/github.com/shudiwsh2009/reservation_thzy_go/assets ./assets
cd ~/thzyfzzx_go/server
chmod a+x ./server.run
nohup ./server.run --app-env="ONLINE" --mail-smtp="smtp.qq.com" --mail-username="thzyfzzx@qq.com" --mail-password="zztwhqsqpjeydche" > ~/thzyfzzx_go/log/server.log 2>&1 & echo $! > ~/thzyfzzx_go/run.pid &

echo "#deploy reminder"
cd ~/thzyfzzx_go
go install github.com/shudiwsh2009/reservation_thzy_go/reminder
cp -f $GOPATH/bin/reminder ./server/reminder.run
chmod a+x ./server/reminder.run
#0 20 * * * /root/thzyfzzx_go/server/reminder.run --app-env="ONLINE"
echo "restart cron"
service cron restart

echo "###Deployment Completed"
