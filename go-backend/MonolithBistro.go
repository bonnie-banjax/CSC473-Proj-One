// Create and populate the initial PageState instance
data := PageState{}

// --- Meta Section ---
data.Meta.PageTitle = "The Monolith Bistro"
data.Meta.ThemeClass = "theme-corporate" // matches <body class="theme-corporate">
data.Meta.CurrencyLabel = "$"

// --- Nav Section ---
data.Nav.LogoText = "The Monolith Bistro"
data.Nav.HamBtnIcon = "☰"
data.Nav.HomeLabel = "Home"
data.Nav.MenuLabel = "Menu"
data.Nav.AboutLabel = "About"
data.Nav.ContactLabel = "Contact"

// --- Hero Section ---
data.Hero.Title = "The Monolith Bistro"
data.Hero.Subtitle = "Exquisite Dining Rendered in a Single Architecture"

// --- About Section ---
data.About.Title = "Our Story"
data.About.AlertText = "" // No alert active in the provided source file
data.About.Content = "Founded in 2026, The Monolith Bistro was born out of a desire to simplify gastronomy while maximizing depth. We believe that great meals—much like great applications—don't require dozens of disconnected configurations. By sourcing every ingredient directly from local micro-farms and treating our kitchen as a single ecosystem, we create a unified culinary masterpiece on every single plate."

// --- Menu Section ---
data.Menu.Title = "Our Menu"
data.Menu.ActionBtnText = "Add to Cart"
data.Menu.Items = []MenuItem{
	{
		ID:     "1",
		Name:   "Truffle Burger",
		Price:  "18.00",
		Desc:   "Prime wagyu beef patty topped with white truffle aioli, aged gruyère, and wild arugula on a brioche bun.",
		ImgSrc: "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?auto=format&fit=crop&w=500&q=80",
	},
	{
		ID:     "2",
		Name:   "Artisanal Pizza",
		Price:  "16.50",
		Desc:   "Hand-tossed sourdough crust with house-made san marzano tomato sauce, fresh buffalo mozzarella, and basil.",
		ImgSrc: "https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=500&q=80",
	},
	{
		ID:     "3",
		Name:   "Harvest Salmon Bowl",
		Price:  "21.00",
		Desc:   "Pan-seared Atlantic salmon over wild rice, roasted sweet potatoes, avocado, and a citrus-ginger drizzle.",
		ImgSrc: "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?auto=format&fit=crop&w=500&q=80",
	},
}

// --- Gallery Section ---
data.Gallery.Title = "Gallery"
data.Gallery.Images = []string{
	"https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1555396273-367ea4eb4db5?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1414235077428-338989a2e8c0?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1559339352-11d035aa65de?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1514933651103-005eec06c04b?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=800&q=80",
}

// --- Contact Section ---
data.Contact.Title = "Contact Us"
data.Contact.NameLabel = "Name"
data.Contact.EmailLabel = "Email"
data.Contact.MessageLabel = "Message"
data.Contact.SubmitBtnText = "Send Message"
data.Contact.MapUrl = "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3022.6175392743824!2d-73.9878531242371!3d40.74844447138905!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c259a9b3117469%3A0xd134e199a405a163!2sEmpire%20State%20Building!5e0!3m2!1sen!2sus!4v1710000000000!5m2!1sen!2sus"

// --- Cart Section ---
data.Cart.Title = "Your Order"
data.Cart.TriggerLabel = "🛒 Cart"
data.Cart.RemoveBtnText = "Remove" // implied fallback configuration option
data.Cart.EmptyText = "Your cart is empty." // implied fallback configuration option
data.Cart.TotalLabel = "Total:"
data.Cart.ClearBtnText = "Clear Cart"
data.Cart.TotalItems = 0
data.Cart.TotalPrice = 0.00
data.Cart.Items = []CartItem{} // Starts fully empty based on initial page state

// --- Footer Section ---
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