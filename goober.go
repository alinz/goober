package goober

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Goober is a simple interface
type Goober interface {
	Yum(args ...interface{}) error
}

type peanut struct {
	command string
}

// Yum this method is thread and can accept optionals arguments
// and run the command
func (p *peanut) Yum(args ...interface{}) error {
	var args2 []interface{}
	var mapVar map[string]string

	// we are copying the command string
	// to prevent corruption in multiple go routines.
	command := p.command

	for _, arg := range args {
		switch v := arg.(type) {
		case map[string]string:
			if mapVar != nil {
				return fmt.Errorf("multiple maps given")
			}
			mapVar = v
		default:
			args2 = append(args2, arg)
		}
	}

	if mapVar != nil {
		for key, value := range mapVar {
			index := strings.Index(key, "$")
			if index == -1 {
				key = "$" + key
			} else if index != 0 {
				return fmt.Errorf("$ must be the first char for '%s'", key)
			}
			command = strings.Replace(command, key, value, -1)
		}
	}

	execCommand := fmt.Sprintf(command, args2...)

	// fmt.Println(execCommand)

	cmd := exec.Command("bash", "-c", execCommand)
	cmd.Env = os.Environ()

	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf(out.String())
		//return err
	}

	return nil
}

// New accepts commands with format and returns a Goober object
func New(format string) Goober {
	// removed all the line breaks and make the string flat
	command := strings.Join(strings.Fields(format), " ")
	return &peanut{
		command,
	}
}

// Logger is a simple type which has the similar signature as Printf
type Logger func(fmt string, args ...interface{})

// NewLogger creates a Logger object and cleanup function
// once a logger is created, upon on termination, cleanup fucntion must be called.
// the best practice is that as soon as logger created, you have to defer the cleanup
func NewLogger() (logger Logger, cleanup func()) {
	prevStrLen := 80
	pipe := make(chan string, 10)

	go func() {
		for value := range pipe {
			fmt.Printf("\r%s", value)
		}
	}()

	logger = func(format string, values ...interface{}) {
		pipe <- strings.Repeat(" ", prevStrLen)
		value := fmt.Sprintf(format, values...)
		prevStrLen = len(value)
		pipe <- value
	}

	cleanup = func() {
		logger("")
		time.Sleep(100 * time.Microsecond)
		close(pipe)
	}

	return
}
