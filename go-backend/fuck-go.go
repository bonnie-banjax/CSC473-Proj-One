package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Item represents our invariant menu entity data
type Item struct {
	ID    int
	Name  string
	Price float64
}

// Global database of items matching the template layout targets
var catalog = map[int]Item{
	1: {ID: 1, Name: "Truffle Burger", Price: 18.00},
	2: {ID: 2, Name: "Artisanal Pizza", Price: 16.50},
	3: {ID: 3, Name: "Harvest Salmon Bowl", Price: 21.00},
}

// InMemory Session State Tracker
type CartSession struct {
	sync.Mutex
	Items map[int]int // Maps Item ID -> Quantity
}

var session = &CartSession{Items: make(map[int]int)}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/api/cart/add", handleAddToCart)
	http.HandleFunc("/api/cart/clear", handleClearCart)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Helper to configure required Server-Sent Event (SSE) stream headers
func setSSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

func handleAddToCart(w http.ResponseWriter, r *http.Request) {
	setSSEHeaders(w)

	// Extract parameters
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || catalog[id].ID == 0 {
		return
	}

	// Mutate session state
	session.Lock()
	session.Items[id]++

	// Compute new calculations out of the mutated state
	totalItemsCount := 0
	totalPrice := 0.0
	var itemsHTML string

	for k, qty := range session.Items {
		item := catalog[k]
		totalItemsCount += qty
		totalPrice += item.Price * float64(qty)

		// Render dynamic hypermedia lists on the fly
		itemsHTML += fmt.Sprintf(`
			<div class="cart-item">
				<div class="cart-item-info">
					<h4>%s</h4>
					<span>$%0.2f x %d</span>
				</div>
			</div>`, item.Name, item.Price, qty)
	}
	session.Unlock()

	// 1. Send State Patch (dspl) event to instantly update Datastar client store variables
	fmt.Fprintf(w, "event: dspl\n")
	fmt.Fprintf(w, "data: merge {\"cart_count\": %d, \"cart_total\": \"$%0.2f\", \"cart_open\": true}\n\n", totalItemsCount, totalPrice)

	// 2. Send Hypermedia Fragment (dshf) event to swap out the items container list
	fmt.Fprintf(w, "event: dshf\n")
	fmt.Fprintf(w, "data: selector #cart-items-container\n")
	fmt.Fprintf(w, "data: <div id=\"cart-items-container\" class=\"cart-items\">%s</div>\n\n", itemsHTML)
}

func handleClearCart(w http.ResponseWriter, r *http.Request) {
	setSSEHeaders(w)

	// Reset state memory cleanly
	session.Lock()
	session.Items = make(map[int]int)
	session.Unlock()

	// 1. Reset client store parameters back to default state
	fmt.Fprintf(w, "event: dspl\n")
	fmt.Fprintf(w, "data: merge {\"cart_count\": 0, \"cart_total\": \"$0.00\"}\n\n")

	// 2. Return fallback fragment structure
	fmt.Fprintf(w, "event: dshf\n")
	fmt.Fprintf(w, "data: selector #cart-items-container\n")
	fmt.Fprintf(w, "data: <div id=\"cart-items-container\" class=\"cart-items\"><p style=\"text-align:center; color:#888; margin-top:20px;\">Your cart is empty.</p></div>\n\n")
}