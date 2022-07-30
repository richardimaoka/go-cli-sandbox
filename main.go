package main

import (
    "fmt"
    "bufio"
    "os"
    "os/exec"
)

func main() {
    exec.Command("")
    fmt.Println("gloと入力して下さい。")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        if scanner.Text() == "glo" {
	    out, err := exec.Command("glo").Output()
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
