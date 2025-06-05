package main

import (
	"cli-text-editor/editor"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

var ACTIONS = []string{
	"Create New File",
	"Edit Existing File",
	"Delete File",
	"Exit",
}

func main() {
	for {
		prompt := promptui.Select{
			Label: "Text Editor Menu",
			Items: ACTIONS,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Println("Prompt failed", err)
			os.Exit(1)
		}

		switch result {
		case ACTIONS[0]:
			editor.Create_File()
		case ACTIONS[3]:
			return
		}
	}
}
