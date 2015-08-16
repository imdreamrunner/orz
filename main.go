package main

import (
    "log"
    "net/http"
    "database/sql"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "github.com/russross/blackfriday"
    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
    StartDb()
    StartServer()
}

func StartDb() {
    conn, err := sql.Open("sqlite3", "idea.db")
    if err != nil {
        panic(err)
    }

    _, err = conn.Exec(`
        create table if not exists idea (
            id integer primary key,
            name text,
            email text,
            link text,
            content text,
            create_time timestamp default current_timestamp
        )`)
    if err != nil {
        panic(err)
    }

    db = conn
}

func StartServer() {
    m := martini.Classic()

    m.Use(martini.Static("_static"))

    m.Use(render.Renderer(render.Options{
        Extensions: []string{".html"},
    }))


    m.Post("/post", PostIdeaHandler)
    m.Get("/ideas", GetIdeasHandler)

    m.Run()
}

func PostIdeaHandler(req *http.Request, render render.Render, log *log.Logger) {
    log.Println("Receive post data.")
    name := req.FormValue("name")
    email := req.FormValue("email")
    link := req.FormValue("link")
    content := req.FormValue("content")
    log.Println("Name: " + name)
    log.Println("Email: " + email)
    log.Println("Link: " + link)
    log.Println("Content: " + content)
    if len(name) == 0 {
        render.JSON(200, map[string]interface{}{"error": 1, "message": "Empty name."})
        return
    }
    if len(email) == 0 {
        render.JSON(200, map[string]interface{}{"error": 2, "message": "Empty email."})
        return
    }
    if len(content) == 0 {
        render.JSON(200, map[string]interface{}{"error": 2, "message": "Empty content."})
        return
    }
    AddIdea(name, email, link, content)
    render.JSON(200, map[string]interface{}{"error": 0})
}

func GetIdeasHandler(req *http.Request, render render.Render, log *log.Logger) {
    render.JSON(200, GetIdeas())
}

func AddIdea(name string, email string, link string, content string) {
    insert, _ := db.Prepare("insert into idea(name, email, link, content) values(?, ?, ?, ?)")
    _, _ = insert.Exec(name, email, link, content)
}

func GetIdeas() []map[string]interface{} {
    var (
        id int
        name string
        email string
        link string
        content string
        timestamp int
    )

    ideas := make([]map[string]interface{}, 0)

    rows, _ := db.Query(`select id, name, email, link, content, strftime("%s", create_time) from idea order by id desc`)
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&id, &name, &email, &link, &content, &timestamp)
        if err != nil {
            log.Fatal(err)
        }
        log.Println("Original content: " + content)
        html := string(blackfriday.MarkdownCommon([]byte(content)))
        log.Println("Converted: " + html)
        ideas = append(ideas, map[string]interface{}{
            "id": id,
            "name": name,
            "email": email,
            "link": link,
            "content": content,
            "html": html,
            "timestamp": timestamp,
        })
    }
    return ideas
}
