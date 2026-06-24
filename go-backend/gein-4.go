package main

import (
	"fmt"
	"html/template"
	"os"
)

// --- 1. CORE STRUCTURE DEFINITIONS ---

type PageState struct {
	Meta struct {
		PageTitle     string
		ThemeClass    string
		CurrencyLabel string
	}
	Nav struct {
		LogoText     string
		HamBtnIcon   string
		HomeLabel    string
		MenuLabel    string
		AboutLabel   string
		ContactLabel string
	}
	Hero struct {
		Title    string
		Subtitle string
	}
	About struct {
		Title     string
		AlertText string
		Content   string
	}
	Menu struct {
		Title         string
		ActionBtnText string
		Items         []MenuItem
	}
	Gallery struct {
		Title  string
		Images []string
	}
	Contact struct {
		Title         string
		NameLabel     string
		EmailLabel    string
		MessageLabel  string
		SubmitBtnText string
		MapUrl        string
	}
	Cart struct {
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
	Footer struct {
		CorpName      string
		Tagline       string
		HoursTitle    string
		HoursLines    []string
		SocialsTitle  string
		SocialLinks   map[string]string
		CopyrightLine string
	}
}

type MenuItem struct {
	ID     string
	Name   string
	Price  string
	Desc   string
	ImgSrc string
}

type CartItem struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

// --- 2. EXTRACTED DATA INPUT BLOCKS ---

func loadBistroInput() PageState {
	data := PageState{}
	data.Meta.PageTitle = "The Monolith Bistro"
	data.Meta.ThemeClass = "theme-corporate"
	data.Meta.CurrencyLabel = "$"

	data.Nav.LogoText = "The Monolith Bistro"
	data.Nav.HamBtnIcon = "☰"
	data.Nav.HomeLabel = "Home"
	data.Nav.MenuLabel = "Menu"
	data.Nav.AboutLabel = "About"
	data.Nav.ContactLabel = "Contact"

	data.Hero.Title = "The Monolith Bistro"
	data.Hero.Subtitle = "Exquisite Dining Rendered in a Single Architecture"

	data.About.Title = "Our Story"
	data.About.AlertText = ""
	data.About.Content = "Founded in 2026, The Monolith Bistro was born out of a desire to simplify gastronomy while maximizing depth. We believe that great meals—much like great applications—don't require dozens of disconnected configurations. By sourcing every ingredient directly from local micro-farms and treating our kitchen as a single ecosystem, we create a unified culinary masterpiece on every single plate."

	data.Menu.Title = "Our Menu"
	data.Menu.ActionBtnText = "Add to Cart"
	data.Menu.Items = []MenuItem{
		{ID: "1", Name: "Truffle Burger", Price: "18.00", Desc: "Prime wagyu beef patty topped with white truffle aioli, aged gruyère, and wild arugula on a brioche bun.", ImgSrc: "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?auto=format&fit=crop&w=500&q=80"},
		{ID: "2", Name: "Artisanal Pizza", Price: "16.50", Desc: "Hand-tossed sourdough crust with house-made san marzano tomato sauce, fresh buffalo mozzarella, and basil.", ImgSrc: "https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=500&q=80"},
		{ID: "3", Name: "Harvest Salmon Bowl", Price: "21.00", Desc: "Pan-seared Atlantic salmon over wild rice, roasted sweet potatoes, avocado, and a citrus-ginger drizzle.", ImgSrc: "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?auto=format&fit=crop&w=500&q=80"},
	}

	data.Gallery.Title = "Gallery"
	data.Gallery.Images = []string{
		"https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1555396273-367ea4eb4db5?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1414235077428-338989a2e8c0?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1559339352-11d035aa65de?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1514933651103-005eec06c04b?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=800&q=80",
	}

	data.Contact.Title = "Contact Us"
	data.Contact.NameLabel = "Name"
	data.Contact.EmailLabel = "Email"
	data.Contact.MessageLabel = "Message"
	data.Contact.SubmitBtnText = "Send Message"
	data.Contact.MapUrl = "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3022.6175392743824!2d-73.9878531242371!3d40.74844447138905!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c259a9b3117469%3A0xd134e199a405a163!2sEmpire%20State%20Building!5e0!3m2!1sen!2sus!4v1710000000000!5m2!1sen!2sus"

	data.Cart.Title = "Your Order"
	data.Cart.TriggerLabel = "🛒 Cart"
	data.Cart.RemoveBtnText = "Remove"
	data.Cart.EmptyText = "Your cart is empty."
	data.Cart.TotalLabel = "Total:"
	data.Cart.ClearBtnText = "Clear Cart"
	data.Cart.TotalItems = 0
	data.Cart.TotalPrice = 0.00
	data.Cart.Items = []CartItem{}

	data.Footer.CorpName = "Monolith Bistro"
	data.Footer.Tagline = "Crafting stateful culinary experiences packed inside pristine architecture."
	data.Footer.HoursTitle = "Hours of Operation"
	data.Footer.HoursLines = []string{
		"Monday - Thursday: 4:00 PM - 10:00 PM",
		"Friday - Sunday: 12:00 PM - 11:00 PM",
	}
	data.Footer.SocialsTitle = "Follow Us"
	data.Footer.SocialLinks = map[string]string{
		"Facebook":    "#",
		"Instagram":   "#",
		"X / Twitter": "#",
	}
	data.Footer.CopyrightLine = "© 2026 The Monolith Bistro. All rights reserved."
	return data
}

func loadMonolithInput() PageState {
	data := PageState{}
	data.Meta.PageTitle = "NEO-MONOLITH // NUTRITION DISPENSARY"
	data.Meta.ThemeClass = "theme-coporate"
	data.Meta.CurrencyLabel = "creds"

	data.Nav.LogoText = "NEO-MONOLITH // v1.0.4"
	data.Nav.HamBtnIcon = "🚨"
	data.Nav.HomeLabel = "Terminal"
	data.Nav.MenuLabel = "Rations"
	data.Nav.AboutLabel = "Log-File"
	data.Nav.ContactLabel = "Geolocate"

	data.Hero.Title = "CYBERNETIC NUTRIENT MATRIX"
	data.Hero.Subtitle = "// CONSUME. SUBMIT. RENDER STATE."

	data.About.Title = "SYSTEM_LOG.TXT"
	data.About.AlertText = "[!] ARCHITECTURAL ANOMALY DETECTED"
	data.About.Content = "This section features generic, heartfelt copy optimized for search algorithms. Every restaurant layout on the web utilizes this exact layout pattern. It took an LLM approximately 0.4 seconds to generate this entire business structure. 94% of all local business websites are structural duplicates operating under the illusion of unique identity. You are browsing inside a deterministic cage of standard boxes styled to look distinct. Enjoy your simulated dining paradigm."

	data.Menu.Title = "Sludge Allocations"
	data.Menu.ActionBtnText = "Download Mass"
	data.Menu.Items = []MenuItem{
		{ID: "1", Name: "Synth-Wagyu Cuboid", Price: "18.00", Desc: "Reconstituted beef proteins suspended in artificial lipid-emulsion on high-density carbohydrate bun. Enhanced with synthetic truffle flavor-code.", ImgSrc: "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?auto=format&fit=crop&w=500&q=80"},
		{ID: "2", Name: "Thermal Dough Disc", Price: "16.50", Desc: "Flatbread mechanism loaded with coagulated cow lipids and high-sodium plant extract fluids. Broiled at 450°C inside corporate kilns.", ImgSrc: "https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=500&q=80"},
		{ID: "3", Name: "Heavy-Metal Salmon Base", Price: "21.00", Desc: "Slices of aquaculture flesh loaded with Omega-3 and trace ocean microplastics, stacked perfectly upon bleached carbohydrate grains.", ImgSrc: "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?auto=format&fit=crop&w=500&q=80"},
	}

	data.Gallery.Title = "Optic Buffers"
	data.Gallery.Images = []string{
		"https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1555396273-367ea4eb4db5?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1414235077428-338989a2e8c0?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1559339352-11d035aa65de?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1514933651103-005eec06c04b?auto=format&fit=crop&w=800&q=80",
		"https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=800&q=80",
	}

	data.Contact.Title = "Telemetry Upload"
	data.Contact.NameLabel = "Designation (Name)"
	data.Contact.EmailLabel = "Comm Net-Node (Email)"
	data.Contact.MessageLabel = "Manifest Manifestations (Message)"
	data.Contact.SubmitBtnText = "Broadcast Signal"
	data.Contact.MapUrl = "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3022.6175392743824!2d-73.9878531242371!3d40.74844447138905!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c259a9b3117469%3A0xd134e199a405a163!2sEmpire%20State%20Building!5e0!3m2!1sen!2sus!4v1710000000000!5m2!1sen!2sus"

	data.Cart.Title = "MANIFEST REQUISITION BUFFER"
	data.Cart.TriggerLabel = "[SLOT_ALLOCATION]"
	data.Cart.RemoveBtnText = "Purge"
	data.Cart.EmptyText = "No items allocated."
	data.Cart.TotalLabel = "CREDIT DEBT:"
	data.Cart.ClearBtnText = "Purge Buffer State"
	data.Cart.TotalItems = 0
	data.Cart.TotalPrice = 0.00
	data.Cart.Items = []CartItem{}

	data.Footer.CorpName = "MONOLITH CONGLOMERATE"
	data.Footer.Tagline = "We abstract your hunger via modular array loops."
	data.Footer.HoursTitle = "Uptime Parameters"
	data.Footer.HoursLines = []string{
		"Cycles 1-4: 1600 - 2200 Hours",
		"Cycles 5-7: 1200 - 2300 Hours",
	}
	data.Footer.SocialsTitle = "Signal Channels"
	data.Footer.SocialLinks = map[string]string{
		"DeepWeb/FB":   "#",
		"NeuralNet/IG": "#",
		"X-Sector":     "#",
	}
	data.Footer.CopyrightLine = "© 2026 NEO-MONOLITH AUTOMATIONS. Built without React. Built without shame."
	return data
}

// --- 3. EXECUTION PIPELINE ---

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Missing parameters.")
		fmt.Println("Usage: go run render.go [bistro|monolith] [output_filename.html]")
		os.Exit(1)
	}

	// 1. Parameter Assignment Switch
	var inputData PageState
	mode := os.Args[1]

	switch mode {
	case "bistro":
		inputData = loadBistroInput()
	case "monolith":
		inputData = loadMonolithInput()
	default:
		fmt.Printf("Unknown variant input mapping: '%s'. Choose 'bistro' or 'monolith'.\n", mode)
		os.Exit(1)
	}

	outputFile := os.Args[2]

	// 2. Load the system template component file
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Printf("Syntax compilation mistake in template structure: %s\n", err)
		os.Exit(1)
	}

	// 3. Create file descriptor pointer for output target destination
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Failed to generate file workspace path targets: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// 4. Resolve full template tree directly to disk footprint
	err = tmpl.Execute(file, inputData)
	if err != nil {
		fmt.Printf("Resolution error while mapping structural text patterns: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Success. Matrix compiled via configuration option '%s' out to '%s'.\n", mode, outputFile)
}