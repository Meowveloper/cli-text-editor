package editor

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func Create_File() {
	prompt := promptui.Prompt{
		Label: "Enter file name (with extension)",
	}

	file_name, err := prompt.Run()

	if err != nil {
		fmt.Println("Prompt failed: ", err)
		return
	}
	fmt.Println("Enter your content (type `:wq` on a new line to save and exit):")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == ":wq" {
			break
		}
		lines = append(lines, line)
	}
	err = os.WriteFile(file_name, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Failed to save file:", err)
		return
	}
	fmt.Println("File saved successfully.")
}
