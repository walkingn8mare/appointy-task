package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "strings"
    "io/ioutil"
    "strconv"
    "time"
)
type Article struct {
    Id string `json:"Id"`
    Title string `json:"Title"`
    SubTitle string `json:"SubTitle"`
    Content string `json:"content"`
    CreatedAt time.Time `json:"createdAt`
}

var Articles []Article

func articleHandler(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body  
    switch r.Method {
	case "GET":
        fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
    return
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
        var article Article 
        json.Unmarshal(reqBody, &article)
        // update our global Articles array to include
        // our new Article
        article.Id=strconv.Itoa(len(Articles) +1)
        article.CreatedAt=time.Now()
        Articles = append(Articles, article)
        json.NewEncoder(w).Encode(article)
    return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}  
    
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Endpoint Hit: homePage")
}
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    parts := strings.Split(r.URL.String(), "/")
    key:=parts[2]
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("article not found"))
		return
}
func searchArticle(w http.ResponseWriter, r *http.Request){
    query := r.URL.Query()
    q := query.Get("q")
    for _, article := range Articles {
        if article.Title == q {
            json.NewEncoder(w).Encode(article)
            return
        }else if article.SubTitle==q{
            json.NewEncoder(w).Encode(article)
            return
        }else if article.Content ==q{
            json.NewEncoder(w).Encode(article)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("article not found"))
		return
}

func handleRequests() {

    http.HandleFunc("/", homePage)
    http.HandleFunc("/articles", articleHandler)
    http.HandleFunc(`/articles/`,returnSingleArticle)
    http.HandleFunc("/articles/search",searchArticle)

    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
   
    Articles = []Article{
        Article{Id: "1", Title: "Hello", SubTitle: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", SubTitle: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}
