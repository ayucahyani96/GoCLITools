/*
Created Date : Oct, 15th 2021
Created By : Ayu Cahyani Febryanti
Source code reference : https://medium.com/easyread/today-i-learned-belajar-membuat-aplikasi-interactive-shell-sederhana-di-golang-2ef013003393
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ mytools")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(commandStr string) error {
	oldLocation := "/var/log/nginx/error.log"
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)

	case "-h":
		fmt.Println("Here is the manual")

	case "/var/log/nginx/error.log -o /User/johnmayer/Desktop/nginxlog.txt":
		newLocation := "/User/johnmayer/Desktop/nginxlog.txt"
		err := os.Rename(oldLocation, newLocation)
		if err != nil {
			log.Fatal(err)
		}

	case "/var/log/nginx/error.log -o /User/johnmayer/Desktop/nginxlog.json":
		newLocation := "/User/johnmayer/Desktop/nginxlog.json"
		err := os.Rename(oldLocation, newLocation)
		if err != nil {
			log.Fatal(err)
		}

	case "/var/log/nginx/error.log -t json":
		fmt.Println("File converted to Json")

	case "/var/log/nginx/error.log -t text":
		fmt.Println("File converted to Plain Text")

	case "/var/log/nginx/error.log":
		fmt.Println("File converted to Plain Text as Default")

	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}


