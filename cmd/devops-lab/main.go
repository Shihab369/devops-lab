package main

import (
	"os"

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

	if len(os.Args) > 1 && os.Args[1] == "--json" {
		output.PrintJSON(results)
		return
	}
	output.PrintResults(results)

}
