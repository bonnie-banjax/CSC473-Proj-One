package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Item struct {
	ID    int
	Name  string
	Price float64
}

var catalog = map[int]Item{
	1: {ID: 1, Name: "Truffle Burger", Price: 18.00},
	2: {ID: 2, Name: "Artisanal Pizza", Price: 16.50},
	3: {ID: 3, Name: "Harvest Salmon Bowl", Price: 21.00},
}

type CartSession struct {
	sync.Mutex
	Items map[int]int
}

var session = &CartSession{Items: make(map[int]int)}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing target file argument.")
		fmt.Println("Usage: ./server <path_to_html_file>")
		os.Exit(1)
	}
	targetFile := os.Args[1]

	if _, err := os.Stat(targetFile); os.IsNotExist(err) {
		fmt.Printf("Error: Target file path '%s' does not exist.\n", targetFile)
		os.Exit(1)
	}

	// 1. Root Route
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, targetFile)
	})

	// 2. Structured Static File Serving
	// This captures everything under /static/ and strips the prefix before checking local disk files
	fs := http.FileServer(http.Dir("."))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// 3. API Endpoints
	http.HandleFunc("POST /api/cart/add", handleAddToCart)
	http.HandleFunc("POST /api/cart/clear", handleClearCart)

	fmt.Printf("Server running on http://localhost:8080 serving layout: %s\n", targetFile)
	http.ListenAndServe(":8080", nil)
}

func setSSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

func handleAddToCart(w http.ResponseWriter, r *http.Request) {
	setSSEHeaders(w)

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || catalog[id].ID == 0 {
		return
	}

	session.Lock()
	session.Items[id]++

	totalItemsCount := 0
	totalPrice := 0.0
	var itemsHTML string

	for k, qty := range session.Items {
		item := catalog[k]
		totalItemsCount += qty
		totalPrice += item.Price * float64(qty)

		itemsHTML += fmt.Sprintf(`
			<div class="cart-item">
				<div class="cart-item-info">
					<h4>%s</h4>
					<span>$%0.2f x %d</span>
				</div>
			</div>`, item.Name, item.Price, qty)
	}
	session.Unlock()

	fmt.Fprintf(w, "event: dspl\n")
	fmt.Fprintf(w, "data: merge {\"cart_count\": %d, \"cart_total\": \"$%0.2f\", \"cart_open\": true}\n\n", totalItemsCount, totalPrice)

	fmt.Fprintf(w, "event: dshf\n")
	fmt.Fprintf(w, "data: selector #cart-items-container\n")
	fmt.Fprintf(w, "data: <div id=\"cart-items-container\" class=\"cart-items\">%s</div>\n\n", itemsHTML)
}

func handleClearCart(w http.ResponseWriter, r *http.Request) {
	setSSEHeaders(w)

	session.Lock()
	session.Items = make(map[int]int)
	session.Unlock()

	fmt.Fprintf(w, "event: dspl\n")
	fmt.Fprintf(w, "data: merge {\"cart_count\": 0, \"cart_total\": \"$0.00\"}\n\n")

	fmt.Fprintf(w, "event: dshf\n")
	fmt.Fprintf(w, "data: selector #cart-items-container\n")
	fmt.Fprintf(w, "data: <div id=\"cart-items-container\" class=\"cart-items\"><p style=\"text-align:center; color:#888; margin-top:20px;\">Your cart is empty.</p></div>\n\n")
}