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
    fmt.Print("Enter Git command: ")
    git_cmd, _ := reader.ReadString('\n')
    git_cmd     = strings.TrimSpace(git_cmd)

    fmt.Print("Enter repo folder that you have already downloaded: ")
    dir, _ := reader.ReadString('\n')
    dir     = strings.TrimSpace(dir)

    cmd := exec.Command("git", git_cmd)
    cmd.Dir = "/home/user/repos/" + dir + "/"
    fmt.Println(cmd.Dir)

    res, err := cmd.Output()
    if err != nil {
        log.Fatalf("%v\n", err)
    }
    fmt.Println(string(res))
}
