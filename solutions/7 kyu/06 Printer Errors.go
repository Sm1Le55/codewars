package kata

import "fmt"

func PrinterError(s string) string {
	errorCount := 0
	for _, c := range s {
		if c >= 110 {
			errorCount++
		}
	}
	return fmt.Sprintf("%v/%v", errorCount, len(s))
}
