package html

import (
	"fmt"
	"os"
	"strings"
)

func UpdateHTMLFiles(category string, items []string) error {
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

	return os.WriteFile("index.html", []byte(content), 0644)
}

func updateArchiveHTML(category string, items []string) error {
	// Read the entire file
	content, err := os.ReadFile("archive.html")
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
	return os.WriteFile("archive.html", []byte(newContent), 0644)
}
