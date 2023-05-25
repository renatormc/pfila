package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// URL of the PowerShell script on GitHub
	scriptURL := "https://raw.githubusercontent.com/renatormc/pfila/main/update.ps1"

	// Destination file path to save the script
	scriptPath := filepath.Join(os.TempDir(), "script.ps1")

	// Download the script file
	err := downloadFile(scriptURL, scriptPath)
	if err != nil {
		fmt.Println("Error downloading the script:", err)
		return
	}

	// Make the script file executable (optional)
	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		fmt.Println("Error making the script executable:", err)
		return
	}

	// Execute the PowerShell script
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing the script:", err)
		return
	}

	fmt.Println("Script executed successfully!")
}

func downloadFile(url string, destPath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
