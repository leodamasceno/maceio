package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func main() {

    cmd := exec.Command("/bin/sh", "-c", "ls")
    cmd_output, _ := cmd.CombinedOutput()

    output := string(cmd_output)

    format_msg := "```"+output+"```"
    msg := `{"text": "`+format_msg+`"}`


    fmt.Println(strings.Replace(msg,"\n","\\n",-1))

}
