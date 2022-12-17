package tasks

import (
	"fmt"
	"github.com/Relativity74205/murkelhausen/gohausen/internal/common"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	log "github.com/sirupsen/logrus"
	"time"
)

func getCPUUsageAverage(coreUsages []float64) float64 {
	var cpuUsageAvg float64
	for _, coreUsage := range coreUsages {
		cpuUsageAvg += coreUsage
	}

	return cpuUsageAvg / float64(len(coreUsages))
}

func getIOCounter(ioCounters []net.IOCountersStat, nameIOCounter string) (net.IOCountersStat, error) {
	for _, ioCounter := range ioCounters {
		if ioCounter.Name == nameIOCounter {
			return ioCounter, nil
		}
	}

	return net.IOCountersStat{}, fmt.Errorf("searched ioCounter not found")
}

func getStats(messageQueue chan common.ChannelPayload) {
	log.Info("Starting to collect system stats...")
	virtualMemory, _ := mem.VirtualMemory()
	cpuCores, _ := cpu.Counts(false)
	cpuLogical, _ := cpu.Counts(true)
	cpuAverageTime := 100 * time.Millisecond // TODO: move to config
	cpuLoad, _ := cpu.Percent(cpuAverageTime, true)
	// TODO add single cpuLoads
	hostInfo, _ := host.Info() // TODO: get OS, Platform, PlatformFamily, PlatformVersion, KernelVersion, KernelArch
	rootDiskUsage, _ := disk.Usage("/")
	loadValues, _ := load.Avg()
	ioCounters, _ := net.IOCounters(true)
	eth0ioCounter, _ := getIOCounter(ioCounters, "eth0") // TODO mainIoCounters, _ := disk.IOCounters("sda")
	processes, _ := process.Processes()
	// TODO more info about running processes
	// TODO min/max/current CPU frequency
	// intel current: cat /proc/cpuinfo | grep "MHz"   https://github.com/c9s/goprocinfo
	// intel min/max: lscpu | grep MHz
	// rasp:
	// sudo cat /sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_cur_freq
	// sudo cat /sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq
	// sudo cat /sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq
	// TODO system temperature (only raspberry? /usr/bin/vcgencmd measure_temp )
	// /usr/bin/vcgencmd measure_temp

	state := common.SystemState{
		Hostname:            hostInfo.Hostname,
		Uptime:              int64(hostInfo.Uptime),
		MemoryTotal:         int64(virtualMemory.Total),
		MemoryAvailable:     int64(virtualMemory.Available),
		MemoryUsed:          int64(virtualMemory.Used),
		MemoryUsedPercent:   virtualMemory.UsedPercent,
		MemoryFree:          int64(virtualMemory.Free),
		CpuCores:            cpuCores,
		CpuLogical:          cpuLogical,
		CpuUsageAvg:         getCPUUsageAverage(cpuLoad),
		RootDiskTotal:       int64(rootDiskUsage.Total),
		RootDiskFree:        int64(rootDiskUsage.Free),
		RootDiskUsed:        int64(rootDiskUsage.Used),
		RootDiskUsedPercent: rootDiskUsage.UsedPercent,
		Load01:              loadValues.Load1,
		Load05:              loadValues.Load5,
		Load15:              loadValues.Load15,
		NetworkBytesSent:    int64(eth0ioCounter.BytesSent),
		NetworkBytesRecv:    int64(eth0ioCounter.BytesRecv),
		ProcessCount:        len(processes),
	}

	log.WithField("SystemState", state).Debug("Stats collected. Sending...")

	channelPayload := common.ChannelPayload{
		Topic: "gohausenStates",
		Key:   state.Hostname,
		Value: state,
	}
	messageQueue <- channelPayload
	log.Info("Stats collected and send.")
}
