package goober

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Option helps chnaging Options
type Option func(*Peanut) error

// OptSystemEnv is a optional operation which adds all the system env
// to the Peanut and each env variable can be accessed by $<name>. e.g. $GOPATH
func OptSystemEnv() Option {
	return func(peanut *Peanut) error {
		if peanut.env == nil {
			peanut.env = make(map[string]string)
		}

		systemEnvs := os.Environ()
		for _, systemEnv := range systemEnvs {
			keyValue := strings.Split(systemEnv, "=")
			if len(keyValue) == 2 {
				peanut.env[keyValue[0]] = keyValue[1]
			}
		}

		return nil
	}
}

// OptCustomEnv adds custom env to internal env which can be accessed by
// $<name> e.g. $GOPATH
func OptCustomEnv(env map[string]string) Option {
	return func(peanut *Peanut) error {
		if peanut.env == nil {
			peanut.env = make(map[string]string)
		}

		for key, value := range env {
			peanut.env[key] = value
		}

		return nil
	}
}

// Peanut is the minimum requires struct to build goober
type Peanut struct {
	env  map[string]string
	cmd  string
	args []string
}

// Yum eats commands and parsed it for later Burp
// it can be chained
func (p *Peanut) Yum(content string) *Peanut {
	lines := strings.Fields(content)

	for _, line := range lines {
		for key, value := range p.env {
			if strings.Index(key, "$") == -1 {
				key = "$" + key
			}
			line = strings.Replace(line, key, value, -1)
		}

		if p.cmd == "" {
			p.cmd = line
		} else {
			p.args = append(p.args, line)
		}
	}

	return p
}

// Burp tries to executed what has been yummed.
// returns an error if errors happen
func (p *Peanut) Burp() error {
	cmd := exec.Command(p.cmd, p.args...)
	defer func() {
		p.args = p.args[0:0]
		p.cmd = ""
	}()

	env := make([]string, 0)
	for key, value := range p.env {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	cmd.Env = env

	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(out.String())

		return err
	}

	return nil
}

// New creates Peanut to be consumed and yumed
// it accepts optional configuration
func New(options ...Option) (*Peanut, error) {
	peanut := &Peanut{}

	for _, option := range options {
		if err := option(peanut); err != nil {
			return nil, err
		}
	}

	return peanut, nil
}
