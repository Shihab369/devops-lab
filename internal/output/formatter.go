package output

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/Shihab369/devops-lab/internal/models"
)

func PrintResults(results []models.Result) {
	for _, result := range results {
		fmt.Printf("[%s]\n", result.Name)
		printData(result.Data, 1)
	}
}

func PrintJSON(results []models.Result) {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("failed to encode results:", err)
		return
	}

	fmt.Println(string(data))
}

func sortedKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))

	for key := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func printData(data map[string]interface{}, level int) {
	for _, key := range sortedKeys(data) {
		value := data[key]

		printIndent(level)
		fmt.Printf("%s: ", key)

		switch nested := value.(type) {
		case map[string]interface{}:
			fmt.Println()
			printData(nested, level+1)

		case map[string]string:
			fmt.Println()
			for key, value := range nested {
				printIndent(level + 1)
				fmt.Printf("%s: %s\n", key, value)
			}

		case map[string]uint64:
			fmt.Println()
			for key, value := range nested {
				printIndent(level + 1)
				fmt.Printf("%s: %d\n", key, value)
			}

		default:
			fmt.Printf("%v\n", value)
		}
	}
}

func printIndent(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
}
