package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/Shihab369/devops-thinking-lab/collectors/cpu"
	"github.com/Shihab369/devops-thinking-lab/collectors/disk"
	"github.com/Shihab369/devops-thinking-lab/collectors/load"
	"github.com/Shihab369/devops-thinking-lab/collectors/memory"
	"github.com/Shihab369/devops-thinking-lab/collectors/process"
	"github.com/Shihab369/devops-thinking-lab/collectors/system"
	"github.com/Shihab369/devops-thinking-lab/collectors/uptime"
	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
	"github.com/Shihab369/devops-thinking-lab/internal/output"
)

func main() {
	jsonOutput := flag.Bool("json", false, "output results as JSON")
	flag.Parse()

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

	if *jsonOutput {
		printJSON(results)
		return
	}

	output.PrintResults(results)
}

func printJSON(results []models.Result) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(results); err != nil {
		fmt.Fprintf(os.Stderr, "error encoding results: %v\n", err)
		os.Exit(1)
	}
}
