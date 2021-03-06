package tail

import (
	"fmt"
	"os"
	"strings"
)

type simpleLogWriter struct {
	Type string
}

func (w simpleLogWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	s = strings.Trim(s, "\n")

	if w.Type == "stdout" {
		fmt.Fprintf(os.Stdout, s)
	} else {
		fmt.Fprintf(os.Stderr, s)
	}

	return len(p), nil
}
