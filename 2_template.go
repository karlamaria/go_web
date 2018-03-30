package main

import (
  "html/template"
  "net/http"
  "os"
)

//Data Structures
type Job struct {
  Title   string
  Company string
}

type Portfolio struct {
  Name string
  Jobs []Job
}

func renderPortfolio(w http.ResponseWriter, r *http.Request) {
  data := Portfolio{
    Name: "Karla Garcia",
    Jobs: []Job{
      {Title: "Software Developer", Company: "RD"},
      {Title: "Systems Analyst", Company: "Digitro Tecnologia"},
      {Title: "SysAdmin", Company: "VIAS"},
    },
  }

  tmpl, _ := template.ParseFiles("templates/layout.html")
  // Accepts an io.Writer for writing out the template and an interface{} to pass data into the template
  tmpl.Execute(w, data)
}

func configPort() string {
  port := os.Getenv("PORT")
  if port != "" {
    return ":" + port
  }
  return ":9999"
}

func main() {
  http.HandleFunc("/", renderPortfolio)

  // fs := http.FileServer(http.Dir("templates/styles/"))
  // http.Handle("/styles/", http.StripPrefix("/styles/", fs))

  http.ListenAndServe(":9999", nil)
  // http.ListenAndServe(configPort(), nil)
}