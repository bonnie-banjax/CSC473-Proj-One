package main

import (
	"html/template"
	"os"
)

// Paste the models you provided
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

func main() {
	// 1. Create dummy debugging data fulfilling your PageState struct
	data := PageState{}

	data.Meta.PageTitle = "Morphic Mono-Culture"
	data.Meta.ThemeClass = "theme-cyberpunk"
	data.Meta.CurrencyLabel = "creds"

	data.Nav.LogoText = "⚡ NEO-TOKYO ⚡"
	data.Nav.HamBtnIcon = "☰"
	data.Nav.HomeLabel = "Grid"
	data.Nav.MenuLabel = "Rations"
	data.Nav.AboutLabel = "Manifesto"
	data.Nav.ContactLabel = "Comms"

	data.Hero.Title = "Welcome to the Neon Jungle"
	data.Hero.Subtitle = "High tech, low life food options."

	data.About.Title = "Our Manifesto"
	data.About.AlertText = "SYSTEM ANOMALY DETECTED"
	data.About.Content = "We serve synthetic nutrients optimized for street samurai."

	data.Menu.Title = "Daily Rations"
	data.Menu.ActionBtnText = "Load to Deck"
	data.Menu.Items = []MenuItem{
		{ID: "item_1", Name: "Synth-Noodles", Price: "45", Desc: "Rehydrated carb-strings with spicy broth.", ImgSrc: "noodles.jpg"},
		{ID: "item_2", Name: "Caffeine-Bolt", Price: "12", Desc: "Pure alertness in a reusable can.", ImgSrc: "coffee.jpg"},
	}

	data.Gallery.Title = "Visual Fragments"
	data.Gallery.Images = []string{"pic1.jpg", "pic2.jpg"}

	data.Contact.Title = "Secure Comms"
	data.Contact.NameLabel = "Handle"
	data.Contact.EmailLabel = "Matrix Node"
	data.Contact.MessageLabel = "Encrypted Data"
	data.Contact.SubmitBtnText = "Transmit"
	data.Contact.MapUrl = "https://maps.example.com"

	data.Cart.Title = "Your Deck Cache"
	data.Cart.TriggerLabel = "Payload"
	data.Cart.RemoveBtnText = "Purge"
	data.Cart.EmptyText = "Cache empty. Go scavenge."
	data.Cart.TotalLabel = "Total Debt"
	data.Cart.ClearBtnText = "Wipe Cache"
	data.Cart.TotalItems = 1
	data.Cart.TotalPrice = 45.0
	data.Cart.Items = []CartItem{
		{ID: "item_1", Name: "Synth-Noodles", Price: 45.0, Quantity: 1},
	}

	data.Footer.CorpName = "Megacorp Inc."
	data.Footer.Tagline = "Obey. Consume. Compile."
	data.Footer.HoursTitle = "Uptime"
	data.Footer.HoursLines = []string{"Cycle 01-05: 24hrs", "Cycle 06-07: Offline"}
	data.Footer.SocialsTitle = "Frequencies"
	data.Footer.SocialLinks = map[string]string{
		"Matrix": "https://matrix.example.com",
		"Sub-net": "https://subnet.example.com",
	}
	data.Footer.CopyrightLine = "© 2026 Megacorp. All rights reserved."

	// 2. Parse your template file
	// Note: We use template.Must to panic if your template has syntax errors.
	tmpl := template.Must(template.ParseFiles("index.html"))

	// 3. Execute the template and stream output to standard console output (os.Stdout)
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}