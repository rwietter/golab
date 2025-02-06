package base

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Snippet struct {
	Id          int
	Title       string
	Content     string
	Language    string
	Description string
}

type SnippetRepository interface {
	GetAll() []Snippet
	GetById(id int) Snippet
	Create(snippet Snippet) Snippet
	Update(snippet Snippet) Snippet
	Delete(id int)
}

func (s Snippet) GetAll(db *sql.DB) []Snippet {
	snippets := []Snippet{}
	result, err := db.Query("SELECT * FROM snippets")
	if err != nil {
		panic(err)
	}
	defer result.Close()

	for result.Next() {
		snippet := Snippet{}
		err := result.Scan(&snippet.Id, &snippet.Title, &snippet.Content, &snippet.Language, &snippet.Description)
		if err != nil {
			panic(err)
		}
		snippets = append(snippets, snippet)
	}
	return snippets
}

func (s Snippet) Create(db *sql.DB) Snippet {
	_, err := db.Exec("INSERT INTO snippets (title, content, language, description) VALUES (?, ?, ?, ?)", s.Title, s.Content, s.Language, s.Description)
	if err != nil {
		panic(err)
	}
	return s
}

func createSnippet(s Snippet, db *sql.DB) {
	s.Create(db)
}

func GetAllSnippets(db *sql.DB) []Snippet {
	snippet := Snippet{}
	return snippet.GetAll(db)
}

func CreateSnippetInterface() {
	db, err := sql.Open("sqlite3", "snippets.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS snippets (id INTEGER PRIMARY KEY, title TEXT, content TEXT, language TEXT, description TEXT)")
	if err != nil {
		panic(err)
	}

	snippet := Snippet{
		Id:          1,
		Title:       "Hello, World",
		Content:     "fmt.Println('Hello, World')",
		Language:    "go",
		Description: "Prints Hello, World in the console",
	}

	createSnippet(snippet, db)

	selectedSnippet := GetAllSnippets(db)
	fmt.Println("selectedSnippet: ", selectedSnippet[1])
}
