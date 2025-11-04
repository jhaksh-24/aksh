package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"strconv"
)

func main() {
	for {
		currdir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Fprint(os.Stdout, currdir,"\n$ ")
		cmdLine, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		
		cmdLine = strings.TrimSpace(cmdLine)

		parts := strings.Fields(cmdLine)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		if cmd == "exit" {
			code := 0
			if len(parts) > 1{
				code, err = strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Invalid exit code:", parts[1])
				}
			}

			os.Exit(code)

		} else if cmd == "echo" {

			fmt.Println(Echo(parts))

		} else if cmd == "type" {

			Type(parts)

		} else if cmd == "cd" {

			ChangeDirectory(parts[1:]...)
			
		} else {

			Execute(cmd, parts[1:]...)

		}
	}
}

func InvalidCmd(cmd string) error {
	return fmt.Errorf("%s: command not found",cmd)
}

func Echo (parts []string) string {
	if len(parts) <= 1 {
		return ""
	}
	return strings.Join(parts[1:], " ")
}

func Type (parts []string) {
	if len(parts) <= 1 {
		return
	}
	
	builtins := []string{"exit", "echo", "type", "cd"}

	for _, val := range builtins {
		if val == parts[1] {
			fmt.Println(val, "is a shell builtin")
			return
		}
	}
	
	if file, exists := searchBinInPath(parts[1]); exists {
		fmt.Fprintf(os.Stdout, "%s is %s\n", parts[1], file)
		return
	}

	fmt.Printf("%s: not found\n", parts[1])
}

func searchBinInPath(bin string) (string, bool) {
	pathEnv := os.Getenv("PATH")

	dirs := strings.Split(pathEnv, ":")
	for _, dir := range dirs {
		fullPath := dir + "/" + bin

		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}
	return "", false
}

func Execute(cmd string, args ...string) {
	if file,exists := searchBinInPath(cmd); exists {		
		command := exec.Command(file, args...)

		output, err := command.CombinedOutput()
		fmt.Println(string(output))

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	} else {
		err := InvalidCmd(cmd)
		fmt.Println(err)
	}
}

func ChangeDirectory(path ...string) error {
	if len(path) > 1 {
		return fmt.Errorf("Too many args for cd command\n")
	}

	var target string
	var err error

	if len(path) == 0 || path[0] == "~" || path[0] == "" {
		if target, err = os.UserHomeDir(); err != nil {
			return fmt.Errorf("Error getting home directory: %v\n", err)
		} 
		err = os.Chdir(target)
		if err != nil {
			return fmt.Errorf("Error changing directory: %v\n", err)
		}
	} else {
		target = path[0]
		err = os.Chdir(target)
		if err != nil {
			return fmt.Errorf("Error changing directory: %v\n", err)
		}
	}
	return nil
}
