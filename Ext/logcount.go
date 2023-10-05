package main

import (
	"fmt"
	"strings"
)

func main() {
	logReport := `
3. Warnings and Errors (Past 24h)
#TiDB WARN= 0 ERROR= 0
#Tikv WARN= 9 ERROR= 0
#PD WARN= 29 ERROR= 0
`

	// Split the log report by newline characters to get lines.
	lines := strings.Split(logReport, "\n")

	// Initialize variables to store the counts.
	tidbWarnCount := 0
	tidbErrorCount := 0
	tikvWarnCount := 0
	tikvErrorCount := 0
	pdWarnCount := 0
	pdErrorCount := 0

	// Process each line and extract counts.
	for _, line := range lines {
		// Remove leading and trailing whitespace.
		line = strings.TrimSpace(line)

		// Check for TiDB, TiKV, and PD components and extract counts.
		if strings.HasPrefix(line, "#TiDB") {
			_, err := fmt.Sscanf(line, "#TiDB\nWARN= %d\nERROR= %d", &tidbWarnCount, &tidbErrorCount)
			if err != nil {
				fmt.Println("Error parsing TiDB counts:", err)
			}
		} else if strings.HasPrefix(line, "#Tikv") {
			_, err := fmt.Sscanf(line, "#Tikv\nWARN= %d\nERROR= %d", &tikvWarnCount, &tikvErrorCount)
			if err != nil {
				fmt.Println("Error parsing TiKV counts:", err)
			}
		} else if strings.HasPrefix(line, "#PD") {
			_, err := fmt.Sscanf(line, "#PD\nWARN= %d\nERROR= %d", &pdWarnCount, &pdErrorCount)
			if err != nil {
				fmt.Println("Error parsing PD counts:", err)
			}
		}
	}

	// Display the extracted counts.
	fmt.Printf("TiDB: WARN=%d ERROR=%d\n", tidbWarnCount, tidbErrorCount)
	fmt.Printf("TiKV: WARN=%d ERROR=%d\n", tikvWarnCount, tikvErrorCount)
	fmt.Printf("PD: WARN=%d ERROR=%d\n", pdWarnCount, pdErrorCount)
}
