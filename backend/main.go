package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Link struct {
	Link  string `json:"link"`
	Error string `json:"error"`
}

type Server struct {
	db *sqlx.DB
}

func main() {
	enverr := godotenv.Load(".env")
	if enverr != nil {
		log.Fatalf("Shit went down")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
			dbName,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	s := Server{db: db}

	s.handleRequests()
}

func (s *Server) homePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://add.beanl.ink", http.StatusSeeOther)
}

func (s *Server) addNewLink(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var link Link
	json.Unmarshal(reqBody, &link)
	reg := regexp.MustCompile(`^.*beanconqueror.com.*$`)
	match := reg.FindStringSubmatch(link.Link)
	if match == nil {
		json.NewEncoder(w).Encode(Link{Error: "Mismatched url"})
		return
	}

	id_link := String(10)

	tx := s.db.MustBegin()
	tx.MustExec(
		`
		INSERT INTO links (id_link, return_link)
		VALUES (?, ?)`,
		id_link,
		link.Link,
	)
	tx.Commit()
	returnLink := fmt.Sprintf("https://backend.beanl.ink/%s", id_link)

	json.NewEncoder(w).Encode(Link{Link: returnLink})
}

func (s *Server) returnLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	link := vars["id"]

	var returnLink string

	log.Printf("Link in request is: %v\n", link)

	s.db.Get(&returnLink, "SELECT return_link from links WHERE id_link = ?", link)

	http.Redirect(w, r, returnLink, http.StatusSeeOther)
}

func (s *Server) handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", s.homePage)
	myRouter.HandleFunc("/add", s.addNewLink)
	myRouter.HandleFunc("/{id}", s.returnLink)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "x-api-token"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":10033", handlers.CORS(headersOk, methodsOk)(myRouter)))
}
