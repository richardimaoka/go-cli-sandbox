package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func readFromStdIo() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "glo" {
			out, err := exec.Command("sh -c glo").Output()
			if err != nil {
				fmt.Println("ERROR occurred")
				fmt.Println(err)
			}
			fmt.Println("SUCCESS")
			fmt.Println(out)
			break
		}
		fmt.Println("git status と入力されるまでループします。")
	}
	fmt.Println("終了します。")
}

func main() {
	type CommandData struct {
		Command           string
		CommandExecutable string
		CommandArgs       []string
	}

	data, err := os.ReadFile("commands.json")
	if err != nil {
		fmt.Println("ERROR reading commands.json")
	}

	var commandDataArray []CommandData
	err = json.Unmarshal(data, &commandDataArray)
	if err != nil {
		fmt.Println("json.Unmarshal error")
	}

	for _, cmdData := range commandDataArray {
		cmd := exec.Command("sh", "-c", cmdData.Command)
		fmt.Println(cmdData.Command)
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}
		fmt.Printf("%s\n", out)
	}

}
