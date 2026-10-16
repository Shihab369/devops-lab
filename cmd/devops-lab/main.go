package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/Shihab369/devops-lab/collectors/cpu"
	"github.com/Shihab369/devops-lab/collectors/disk"
	"github.com/Shihab369/devops-lab/collectors/load"
	"github.com/Shihab369/devops-lab/collectors/memory"
	"github.com/Shihab369/devops-lab/collectors/network"
	"github.com/Shihab369/devops-lab/collectors/process"
	"github.com/Shihab369/devops-lab/collectors/system"
	"github.com/Shihab369/devops-lab/collectors/uptime"
	"github.com/Shihab369/devops-lab/internal/check"
	"github.com/Shihab369/devops-lab/internal/core"
	"github.com/Shihab369/devops-lab/internal/models"
	"github.com/Shihab369/devops-lab/internal/output"
)

func main() {
	jsonOutput := flag.Bool("json", false, "output results as JSON")
	checkMode := flag.Bool("check", false, "evaluate health thresholds and exit non-zero on failures")
	cpuMax := flag.Float64("cpu-max", 80, "maximum CPU usage percent for check mode")
	memMax := flag.Float64("mem-max", 85, "maximum memory usage percent for check mode")
	diskMax := flag.Float64("disk-max", 90, "maximum disk usage percent for check mode")
	loadMax := flag.Float64("load-max", float64(runtime.NumCPU()), "maximum 1m load average for check mode")
	flag.Parse()

	runner := core.NewRunner(
		cpu.CPUCollector{},
		disk.DiskCollector{},
		load.LoadCollector{},
		memory.MemoryCollector{},
		network.NetworkCollector{},
		process.ProcessCollector{},
		system.SystemCollector{},
		uptime.UptimeCollector{},
	)

	results := runner.Run()

	if *checkMode {
		summary := check.Evaluate(results, check.Thresholds{
			CPUMaxPercent:    *cpuMax,
			MemoryMaxPercent: *memMax,
			DiskMaxPercent:   *diskMax,
			Load1MMax:        *loadMax,
		})

		if *jsonOutput {
			checkResult := check.ToResult(summary)
			printJSON(append(results, checkResult))
		} else {
			output.PrintResults(results)
			printCheckSummary(summary)
		}

		if summary.Status != "healthy" {
			os.Exit(1)
		}

		return
	}

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

func printCheckSummary(summary check.Summary) {
	fmt.Println("[check]")
	fmt.Printf("status: %s\n", summary.Status)
	fmt.Printf("passed: %d\n", summary.Passed)
	fmt.Printf("failed: %d\n", summary.Failed)

	for _, item := range summary.Items {
		fmt.Printf("- %s: %s (value: %.2f%s, threshold: %.2f%s, %s)\n",
			item.Name,
			item.Status,
			item.Value,
			item.Unit,
			item.Threshold,
			item.Unit,
			item.Message,
		)
	}
}
