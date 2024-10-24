package commands

import (
	"fmt"
	"strings"
)

// TODO : make it parsed with COMMAND and ARGS
func ParseMessage(message string) error {
	splitMsg := strings.Split(message, "\r\n")
	if len(splitMsg) < 2 {
		return fmt.Errorf("invalid input")
	}
	return nil
}
