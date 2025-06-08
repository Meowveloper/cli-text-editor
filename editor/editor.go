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

func Edit_File() {
	prompt := promptui.Prompt{
		Label: "Enter the file name to edit (with extension)",
	}
	file_name, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed: ", err)
		return
	}

	data, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println("Reading file failed: ", err)
	}
	fmt.Println("Current content (press Enter to keep, or enter new content):")
	fmt.Println("----------------------------------")
	fmt.Println(string(data))
	fmt.Println("----------------------------------")
	fmt.Println("Enter your new content (or press Enter to keep existing, type `:wq` to save):")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == ":wq" {
			break
		}
		lines = append(lines, line)
	}

	if len(lines) > 0 {
		err := os.WriteFile(file_name, []byte(strings.Join(lines, "\n")), 0644)
		if err != nil {
			fmt.Println("Error saving file :", err)
			return
		}
		fmt.Println("File updated successfully")
	} else {
		fmt.Println("No changes made")
	}
}

func Delete_File() {
	yes_no := []string{"NO", "YES"}

	prompt := promptui.Prompt{
		Label: "Enter filename to delete",
	}
	file_name, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed")
		return
	}

	confirm_prompt := promptui.Select{
		Label: "Are you sure you want to delete?",
		Items: yes_no,
	}
	_, result, err := confirm_prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed")
		return
	}

	switch result {
	case yes_no[0]:
		fmt.Println("Canceled")
		return
	case yes_no[1]:
		err := os.Remove(file_name)
		if err != nil {
			fmt.Println("Failed to delete: ", err)
			return
		}
		fmt.Println("File deleted")
	}
}
