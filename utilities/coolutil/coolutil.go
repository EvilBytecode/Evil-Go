package utilities

import (
    "fmt"
    "os/exec"
    "runtime"
)

func ExecuteCommand() {
    os := runtime.GOOS

    if os == "windows" {
        cmd := exec.Command("cmd", "/C", "start", "cmd.exe")
        cmd.Start()
        cmd.Wait()
        fmt.Println("Command finished successfully.")
    } else if os == "linux" {
        fmt.Println("Executing command for Linux.")
    } else {
        fmt.Printf("Unsupported operating system: %s\n", os)
    }
}
