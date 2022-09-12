package main

import (
	"bytes"
	"errors"
	"log"
	"strconv"

	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
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

type CommandStep struct {
	Command           string
	CommandExecutable string
	CommandArgs       []string
}

func readjson() {
	data, err := os.ReadFile("commands.json")
	if err != nil {
		fmt.Println("ERROR reading commands.json")
		os.Exit(1)
	}

	var commands []CommandStep
	err = json.Unmarshal(data, &commands)
	if err != nil {
		fmt.Println("json.Unmarshal error")
	}

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadByte()
	// fmt.Println(text)

	for _, cmdData := range commands {
		cmd := exec.Command("sh", "-c", cmdData.Command)
		fmt.Println(cmdData.Command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("ERROR: %s\n", out)
			continue
		}
		fmt.Printf("%s\n\n", out)
	}

}

func prompt() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func lookPath() {
	path, err := exec.LookPath("go")
	if err != nil {
		log.Fatalf("installing fortune is in your future %s", err)
	}
	fmt.Printf("fortune is available at %s\n", path)
}

func login() {
	cmd := "echo"
	login := exec.Command(cmd, "login")

	buffer := bytes.Buffer{}
	buffer.Write([]byte("username\npassword\n"))
	login.Stdin = &buffer

	login.Stdout = os.Stdout
	login.Stderr = os.Stderr

	err := login.Run()
	if err != nil {
		fmt.Printf("error occurred %s\n", err)
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Printf("successfully executed\n")
}

func main() {
	login()
}
