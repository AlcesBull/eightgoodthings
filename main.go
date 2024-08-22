package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Which eight good things are you going to share about today?")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	items := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		fmt.Printf("Enter item %d: ", i+1)
		item, _ := reader.ReadString('\n')
		items = append(items, strings.TrimSpace(item))
	}

	updateHTMLFiles(category, items)
	commitAndPushChanges()
	fmt.Println("Things saved and sent to web!")
}

func updateHTMLFiles(category string, items []string) {
	updateIndexHTML(category, items)
	updateArchiveHTML(category, items)
	panic("unimplemented")
}

func updateIndexHTML(category string, items []string) {
	content := fmt.Sprintf(`
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Eight Good Things</title>
        <link rel="stylesheet" href="style.css">
    </head>
    <body>
        <h1>Eight Good Things</h1>
        <h2>Category: %s</h2>
        <ul>`, category)

	for _, item := range items {
		content += fmt.Sprintf("<li>%s</li>\n", item)
	}

	content += `</ul>
        <a href="archive.html">View Archive</a>
    </body>
    </html>`

	ioutil.WriteFile("index.html", []byte(content), 0644)
}

func updateArchiveHTML(category string, items []string) {
	file, err := os.OpenFile("archive.html", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening archive.html:", err)
		return
	}
	defer file.Close()

	content := fmt.Sprintf("<li><strong>Category: %s</strong>\n<ul>", category)
	for _, item := range items {
		content += fmt.Sprintf("<li>%s</li>\n", item)
	}
	content += "</ul></li>\n"

	file.WriteString(content)
}

func commitAndPushChanges() {
	cmd := exec.Command("git", "add", ".")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error adding files:", err)
		return
	}

	cmd = exec.Command("git", "commit", "-m", "Update website with new list")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	cmd = exec.Command("git", "push")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error pushing changes:", err)
		return
	}

	fmt.Println("Changes pushed to GitHub Pages!")
}
