package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)


func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /books")
	keys, ok := r.URL.Query()["genre"]
	query_string := "SELECT * FROM books"

    if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'genre' is missing")
	}
	fmt.Println(keys)
	if(ok){
		query_string = "SELECT * FROM books WHERE book_genre = '"+keys[0]+"'";
	}
	//localhost:8080/books?genre=test

	rows, err := db.Query(query_string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer rows.Close()
	var books []Book
	for rows.Next() {
		var book Book

		err := rows.Scan(&book.ISBN, &book.Book_title, &book.Page_num, &book.Book_price, &book.Inventory_count, &book.Restock_threshold, &book.Book_genre, &book.Publisher_sale_percentage, &book.Publisher_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		books = append(books, book)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(books)
}


func getBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("GET /books/"+ id)
	rows, err := db.Query(`SELECT * FROM books WHERE isbn = $1`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	var book Book
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ISBN, &book.Book_title, &book.Page_num, &book.Book_price, &book.Inventory_count, &book.Restock_threshold, &book.Book_genre, &book.Publisher_sale_percentage, &book.Publisher_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		book_publisher, err := db.Query(`SELECT * FROM publishers WHERE publisher_id = $1`, book.Publisher_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer book_publisher.Close()

		for book_publisher.Next() {
			var publisher Publisher
			err := book_publisher.Scan(&publisher.Publisher_id, &publisher.Publisher_name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			book.Publisher = publisher
		}

	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if (len(book.ISBN) <= 0){
		json.NewEncoder(w).Encode("Error: Invalid ISBN: "+id)
		fmt.Println("Error: Invalid ISBN: "+id)
	}else{
		json.NewEncoder(w).Encode(book)
	}

}

func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /books")
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	_, err := db.Exec(`INSERT INTO books (isbn, book_title, page_num, book_price, inventory_count, restock_threshold, book_genre, publisher_sale_percentage, publisher_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, book.ISBN, book.Book_title, book.Page_num, book.Book_price, book.Inventory_count, book.Restock_threshold, book.Book_genre, book.Publisher_sale_percentage, book.Publisher_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Book Succesfully Created")
}