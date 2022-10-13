package cmd

import "os/exec"

func main() {
	configSetup()

	cmd := exec.Command("cd", "..")
	err := cmd.Run()
	catchError(err)
	cmd = exec.Command("go", "install")
	err = cmd.Run()
	catchError(err)
}
