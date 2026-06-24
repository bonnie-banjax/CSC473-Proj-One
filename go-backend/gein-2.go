package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Global template cache to prevent parsing files on every single user request
var tmpl *template.Template

// --- 1. CORE DATA MODELS ---

type PageState struct {
	Meta    MetaConfig
	Nav     NavConfig
	Hero    TextConfig
	About   AboutConfig
	Menu    MenuConfig
	Gallery GalleryConfig
	Contact ContactConfig
	Cart    CartConfig
	Footer  FooterConfig
}

type MetaConfig struct {
	PageTitle     string
	ThemeClass    string
	CurrencyLabel string
}

type NavConfig struct {
	LogoText     string
	HamBtnIcon   string
	HomeLabel    string
	MenuLabel    string
	AboutLabel   string
	ContactLabel string
}

type TextConfig struct {
	Title    string
	Subtitle string
}

type AboutConfig struct {
	Title     string
	AlertText string
	Content   string
}

type MenuItem struct {
	ID     string
	Name   string
	Price  string
	Desc   string
	ImgSrc string
}

type MenuConfig struct {
	Title         string
	ActionBtnText string
	Items         []MenuItem
}

type GalleryConfig struct {
	Title  string
	Images []string
}

type ContactConfig struct {
	Title         string
	NameLabel     string
	EmailLabel    string
	MessageLabel  string
	SubmitBtnText string
	MapUrl        string
}

type CartItem struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type CartConfig struct {
	Title         string
	TriggerLabel  string
	RemoveBtnText string
	EmptyText     string
	TotalLabel    string
	ClearBtnText  string
	TotalItems    int
	TotalPrice    float64
	Items         []CartItem
}

type FooterConfig struct {
	CorpName      string
	Tagline       string
	HoursTitle    string
	HoursLines    []string
	SocialsTitle  string
	SocialLinks   map[string]string
	CopyrightLine string
}

// --- 2. SAMPLE SERVICE / DATA STORE SIMULATION ---

// GetBistroState acts as your standard service query layer.
// In production, this would query a database, read environment variables, or parse JSON configurations.
func GetBistroState() PageState {
	state := PageState{}
	state.Meta.PageTitle = "The Monolith Bistro"
	state.Meta.ThemeClass = "theme-corporate"
	state.Meta.CurrencyLabel = "$"

	state.Nav.LogoText = "The Monolith Bistro"
	state.Nav.HamBtnIcon = "☰"
	state.Nav.HomeLabel = "Home"
	state.Nav.MenuLabel = "Menu"
	state.Nav.AboutLabel = "About"
	state.Nav.ContactLabel = "Contact"

	state.Hero.Title = "The Monolith Bistro"
	state.Hero.Subtitle = "Exquisite Dining Rendered in a Single Architecture"

	state.About.Title = "Our Story"
	state.About.Content = "Founded in 2026, The Monolith Bistro was born out of a desire to simplify gastronomy while maximizing depth..."

	state.Menu.Title = "Our Menu"
	state.Menu.ActionBtnText = "Add to Cart"
	state.Menu.Items = []MenuItem{
		{ID: "1", Name: "Truffle Burger", Price: "18.00", Desc: "Prime wagyu beef", ImgSrc: "https://images.unsplash.com/..."},
	}

	state.Cart.Title = "Your Order"
	state.Cart.TriggerLabel = "🛒 Cart"
	state.Cart.TotalLabel = "Total:"
	state.Cart.ClearBtnText = "Clear Cart"

	state.Footer.CorpName = "Monolith Bistro"
	state.Footer.CopyrightLine = "© 2026 The Monolith Bistro. All rights reserved."

	return state
}

// --- 3. PARAMETERIZED HANDLERS ---

// handlePage renders the initial static view using the full template structure.
func handlePage(w http.ResponseWriter, r *http.Request) {
	// Fetch fresh, contextual state mapping specifically to this request cycle
	currentData := GetBistroState()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute standard output using the base layout file
	err := tmpl.Execute(w, currentData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Template compilation failure: %s", err), http.StatusInternalServerError)
	}
}

// handleCartUpdate simulates a Datastar fragment update event.
func handleCartUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return}

	// 1. Process business logic (e.g., parsing a Datastar payload and altering an active cart context)
	currentData := GetBistroState()
	currentData.Cart.Items = []CartItem{{ID: "1", Name: "Truffle Burger", Price: 18.00, Quantity: 1}}
	currentData.Cart.TotalItems = 1
	currentData.Cart.TotalPrice = 18.00

	// 2. Set content type appropriately for Datastar SSE (Server-Sent Events) or raw fragment rendering
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 3. TARGET THE HOOKED BLOCK EXACTLY:
	// Instead of rebuilding the page template layout, you extract only the cart HTML fragment.
	err := tmpl.ExecuteTemplate(w, "cart-sidebar", currentData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// --- 4. RUNTIME BOOTSTRAP ---

func main() {
	// Initialize templates globally ONCE at program bootstrap.
	// template.Must panics early if your target html document contains broken syntax tags.
	var err error
	tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Printf("Fatal error compiling HTML source components: %s\n", err)
		os.Exit(1)
	}

	// Attach endpoints to the server multiplexer
	http.HandleFunc("/", handlePage)             // Normal initial layout load
	http.HandleFunc("/api/cart/add", handleCartUpdate) // Target patch block update endpoint

	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failure: %s\n", err)
	}
}