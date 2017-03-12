package printer

import "log"

// EnableLogf enables logf printer
var EnableLogf = false

// Logf prints log
func Logf(format string, value ...interface{}) {
	if EnableLogf {
		log.Printf(format, value...)
	}
}
