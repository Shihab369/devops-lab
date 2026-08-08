package main

import (
	"github.com/Shihab369/devops-thinking-lab/collectors/cpu"
	"github.com/Shihab369/devops-thinking-lab/collectors/disk"
	"github.com/Shihab369/devops-thinking-lab/collectors/load"
	"github.com/Shihab369/devops-thinking-lab/collectors/memory"
	"github.com/Shihab369/devops-thinking-lab/collectors/process"
	"github.com/Shihab369/devops-thinking-lab/collectors/system"
	"github.com/Shihab369/devops-thinking-lab/collectors/uptime"
	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/output"
)

func main() {
	runner := core.NewRunner(
		cpu.CPUCollector{},
		disk.DiskCollector{},
		load.LoadCollector{},
		memory.MemoryCollector{},
		process.ProcessCollector{},
		system.SystemCollector{},
		uptime.UptimeCollector{},
	)

	results := runner.Run()
	output.PrintResults(results)

}
