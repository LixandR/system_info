package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"sys_info/redis"
	"sys_info/util"
	"time"
)
func init() {
	redis.RedisInit()
}

func main() {
	println("服务器性能监控任务开启...")
	var logdate string

	// CPU使用率
	for {
		//检测mysql的内存占用
		//infoByPName := getInfoByPName("mysqld")

		//限制检测时间
		nowdatetime := time.Now()
		hour := nowdatetime.Hour()
		if hour >= 22 || hour < 7 { //设置空置时间
			println("进入睡眠。。。")
			time.Sleep(time.Minute)
			continue
		}
		datas := make(map[string]interface{})
		//获取ip
		ip := util.GetIp()
		datas["ip"] = ip
		//获取内存使用率 同时定时
		percent, _ := cpu.Percent(time.Second*14, false) //设置间隔时间
		nowtime := nowdatetime.Format("2006-01-02 15:04:05") //当前时间

		nowdate := nowdatetime.Format("2006-01-02")
		if logdate == "" || nowdate != logdate {
			logdate = nowdate
		}
		datas["time"] = nowtime
		unix := nowdatetime.Unix()

		memdata := util.GetMemInfo()
		datas["memory"] = memdata

		hostdata := util.GetHostInfo()
		datas["host"] = hostdata

		cpudata := util.GetCpuInfo(fmt.Sprintf("%.2f", percent[0]))
		datas["cpu"] = cpudata
		//datas["serverName"]=serverName
		//写入文件
		jsonStr, err := json.Marshal(datas)
		if err != nil {
			fmt.Println("MapToJsonDemo err: ", err)
		}
		redis.SaveLogToRedis(string(jsonStr),unix,ip)
	}
}
