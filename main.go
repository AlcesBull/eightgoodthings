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

	err := updateHTMLFiles(category, items)
	if err != nil {
		fmt.Println("Error updating HTML files:", err)
		return
	}

	commitAndPushChanges(category)
	fmt.Println("Things saved and sent to web!")
}

func updateHTMLFiles(category string, items []string) error {
	if err := updateIndexHTML(category, items); err != nil {
		return err
	}
	if err := updateArchiveHTML(category, items); err != nil {
		return err
	}
	return nil
}

func updateIndexHTML(category string, items []string) error {
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

	return ioutil.WriteFile("index.html", []byte(content), 0644)
}

func updateArchiveHTML(category string, items []string) error {
	// Read the entire file
	content, err := ioutil.ReadFile("archive.html")
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error reading archive.html: %w", err)
	}

	// If the file doesn't exist, create a basic HTML structure
	if os.IsNotExist(err) {
		content = []byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Eight Good Things Archive</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <h1>Archive of Lists</h1>
    <ul>
    </ul>
</body>
</html>`)
	}

	// Create the new entry
	newEntry := fmt.Sprintf("<li><strong>Category: %s</strong>\n<ul>", category)
	for _, item := range items {
		newEntry += fmt.Sprintf("<li>%s</li>\n", item)
	}
	newEntry += "</ul></li>\n"

	// Find the position to insert the new entry (after the opening <ul> tag)
	insertPos := strings.Index(string(content), "<ul>") + 4

	// Combine the parts
	newContent := string(content[:insertPos]) + newEntry + string(content[insertPos:])

	// Write the entire content back to the file
	return ioutil.WriteFile("archive.html", []byte(newContent), 0644)
}

func commitAndPushChanges(category string) {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error adding files:", err)
		return
	}

	commitMessage := fmt.Sprintf("Update website with new list: %s", category)
	cmd = exec.Command("git", "commit", "-m", commitMessage)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	cmd = exec.Command("git", "push")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error pushing changes:", err)
		return
	}

	fmt.Println("Changes pushed to GitHub Pages!")
}
