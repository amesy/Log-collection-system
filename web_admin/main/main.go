package main

import (
	"fmt"
	"Log_Project/web_admin/model"
	_ "Log_Project/web_admin/router"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "mysqlcli:12345678@tcp(10.68.7.20:3306)/logadmin")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	model.InitDb(database)
	return
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.68.7.20:2379", "10.68.7.20:22379", "10.68.7.20:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	model.InitEtcd(cli)
	return
}

func main() {

	err := initDb()
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return
	}
	beego.Run()
}
