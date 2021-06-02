package util

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"net"
	"strconv"
	"strings"
)

//最大内存占用预警
var maxMemRate float64 = 80
func GetIp() string {
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip
}

func GetInfoByPName(name string) map[string]interface{} {
	result := make(map[string]interface{})
	pidByName := GetPidByName(name)
	newProcess, err := process.NewProcess(pidByName)
	if err!=nil {
		println("获取pid失败")
	}
	//当前cup
	percent, err := newProcess.CPUPercent()
	println("cpu: ",fmt.Sprintf("%.2f", percent))
	//当前进程运行状态
	status, err := newProcess.Status()
	if err!=nil {
		println(err.Error())
	}
	println("status: ",status)
	//进程是否正在运行
	running, err := newProcess.IsRunning()
	println("running ",running)
	result["cpu"]=percent
	return result
}

//获取内存信息
func GetMemInfo() map[string]string {
	memdata := make(map[string]string)

	v, _ := mem.VirtualMemory()

	total := handerUnit(v.Total, NUM_GB, "G")
	available := handerUnit(v.Available, NUM_GB, "G")
	used := handerUnit(v.Used, NUM_GB, "G")
	free := handerUnit(v.Free, NUM_GB, "G")
	percent := v.UsedPercent
	//短信预警
	if percent >maxMemRate {
		SendRemindSms("内存占用过高"+GetIp()+": "+ strconv.FormatFloat(percent, 'E', -1, 64))
	}
	userPercent := fmt.Sprintf("%.2f", percent)
	memdata["总量"] = total
	memdata["可用"] = available
	memdata["已使用"] = used
	memdata["空闲"] = free
	memdata["使用率"] = userPercent + "%"

	return memdata
}

//获取CPU信息
func GetCpuInfo(percent string) []map[string]string {
	cpudatas := []map[string]string{}

	infos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range infos {
		cpudata := make(map[string]string)
		//cpudata["型号"] = ci.ModelName
		cpudata["数量"] = fmt.Sprint(ci.Cores)
		cpudata["使用率"] = percent + "%"

		cpudatas = append(cpudatas, cpudata)
	}
	return cpudatas
}

//获取主机信息
func GetHostInfo() map[string]string {
	hostdata := make(map[string]string)

	hInfo, _ := host.Info()
	hostdata["主机名称"] = hInfo.Hostname
	hostdata["系统"] = hInfo.OS
	//hostdata["平台"] = hInfo.Platform + "-" + hInfo.PlatformVersion + " " + hInfo.PlatformFamily
	hostdata["平台"] = hInfo.PlatformFamily
	hostdata["内核"] = hInfo.KernelArch

	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
	return hostdata
}


const (
	NUM_KB  = 1000.0000
	NUM_MIB = 1000000.0000
	NUM_GB  = 1000000000.0000
)

func handerUnit(num uint64, numtype float64, typename string) (newnum string) {
	f := fmt.Sprintf("%.2f", float64(num)/numtype)
	return f + typename
}
