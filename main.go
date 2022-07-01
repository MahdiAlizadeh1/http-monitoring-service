package main

import (
	"time"

	"github.com/MahdiAlizadeh1/http-monitoring-service/db"
	"github.com/MahdiAlizadeh1/http-monitoring-service/db_manager"
	"github.com/MahdiAlizadeh1/http-monitoring-service/handler"
	"github.com/MahdiAlizadeh1/http-monitoring-service/service"
)

func main() {
	d := db.Setup("http-monitor.db")
	dm := db_manager.NewDbManager(d)
	mnt := service.NewMonitor(dm)
	sch, _ := service.NewScheduler(mnt)
	sch.WorkInIntervals(time.Minute)
	h := handler.NewHandler(dm, sch)
	h.Start(":8080")
}
