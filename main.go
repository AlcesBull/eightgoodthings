package main

import (
	"eightgoodthings/models"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("logs/dev.log", "8gt")
	if err != nil {
		log.Fatalf("issue opening log file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(models.NewForm(
		"Which eight good things are you going to share about today?",
		8,
	), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// package main
//
// import (
//
//	"bufio"
//	"eightgoodthings/html"
//	"fmt"
//	"os"
//	"strings"
//
// )
//
//	func main() {
//		reader := bufio.NewReader(os.Stdin)
//
//		fmt.Println("Which eight good things are you going to share about today?")
//		category, _ := reader.ReadString('\n')
//		category = strings.TrimSpace(category)
//
//		items := make([]string, 0, 8)
//		for i := 0; i < 8; i++ {
//			fmt.Printf("Enter item %d: ", i+1)
//			item, _ := reader.ReadString('\n')
//			items = append(items, strings.TrimSpace(item))
//		}
//
//		err := html.UpdateFiles(category, items)
//		if err != nil {
//			fmt.Println("Error updating HTML files:", err)
//			return
//		}
//
//		html.CommitAndPushChanges(category)
//		fmt.Println("Things saved and sent to web!")
//	}
