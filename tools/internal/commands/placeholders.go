package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func RunMockServer(args []string) {
	fmt.Println("Starting Mock Server...")

	cmd := exec.Command("go", append([]string{"run", "../mockserver"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Mock server error: %v\n", err)
		os.Exit(1)
	}
}
