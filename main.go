package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Product struct {
	Name        string
	Price       string
	Description string
}

func product(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Product service is running on - " + servant))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	birdJson := `[{"name":"iPhone 14","price":"LKR. 145,000.00", "description":"Brandnew condition"},{"name":"Samsumg S7","price":"LKR. 145,000.00", "description":"Brandnew condition"},{"name":"Nokia Lumia 9","price":"LKR. 145,000.00", "description":"Brandnew condition"}]`
	var birds []Product
	json.Unmarshal([]byte(birdJson), &birds)
	var buf []byte
	buf, err := json.Marshal(birds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(buf))
}

func main() {
	http.HandleFunc("/product", product)
	http.HandleFunc("/product/get", getProducts)
	log.Fatal(http.ListenAndServe(":8070", nil))
}
