// Initialize PageState matching the original cyber-themed variant
data := PageState{}

// --- Meta Section ---
data.Meta.PageTitle = "NEO-MONOLITH // NUTRITION DISPENSARY"
data.Meta.ThemeClass = "theme-coporate" // Retained the source file's typo "theme-coporate"
data.Meta.CurrencyLabel = "creds"

// --- Nav Section ---
data.Nav.LogoText = "NEO-MONOLITH // v1.0.4"
data.Nav.HamBtnIcon = "🚨"
data.Nav.HomeLabel = "Terminal"
data.Nav.MenuLabel = "Rations"
data.Nav.AboutLabel = "Log-File"
data.Nav.ContactLabel = "Geolocate"

// --- Hero Section ---
data.Hero.Title = "CYBERNETIC NUTRIENT MATRIX"
data.Hero.Subtitle = "// CONSUME. SUBMIT. RENDER STATE."

// --- About Section ---
data.About.Title = "SYSTEM_LOG.TXT"
data.About.AlertText = "[!] ARCHITECTURAL ANOMALY DETECTED"
data.About.Content = "This section features generic, heartfelt copy optimized for search algorithms. Every restaurant layout on the web utilizes this exact layout pattern. It took an LLM approximately 0.4 seconds to generate this entire business structure. 94% of all local business websites are structural duplicates operating under the illusion of unique identity. You are browsing inside a deterministic cage of standard boxes styled to look distinct. Enjoy your simulated dining paradigm."

// --- Menu Section ---
data.Menu.Title = "Sludge Allocations"
data.Menu.ActionBtnText = "Download Mass"
data.Menu.Items = []MenuItem{
	{
		ID:     "1",
		Name:   "Synth-Wagyu Cuboid",
		Price:  "18.00",
		Desc:   "Reconstituted beef proteins suspended in artificial lipid-emulsion on high-density carbohydrate bun. Enhanced with synthetic truffle flavor-code.",
		ImgSrc: "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?auto=format&fit=crop&w=500&q=80",
	},
	{
		ID:     "2",
		Name:   "Thermal Dough Disc",
		Price:  "16.50",
		Desc:   "Flatbread mechanism loaded with coagulated cow lipids and high-sodium plant extract fluids. Broiled at 450°C inside corporate kilns.",
		ImgSrc: "https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=500&q=80",
	},
	{
		ID:     "3",
		Name:   "Heavy-Metal Salmon Base",
		Price:  "21.00",
		Desc:   "Slices of aquaculture flesh loaded with Omega-3 and trace ocean microplastics, stacked perfectly upon bleached carbohydrate grains.",
		ImgSrc: "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?auto=format&fit=crop&w=500&q=80",
	},
}

// --- Gallery Section ---
data.Gallery.Title = "Optic Buffers"
data.Gallery.Images = []string{
	"https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1555396273-367ea4eb4db5?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1414235077428-338989a2e8c0?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1559339352-11d035aa65de?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1514933651103-005eec06c04b?auto=format&fit=crop&w=800&q=80",
	"https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=800&q=80",
}

// --- Contact Section ---
data.Contact.Title = "Telemetry Upload"
data.Contact.NameLabel = "Designation (Name)"
data.Contact.EmailLabel = "Comm Net-Node (Email)"
data.Contact.MessageLabel = "Manifest Manifestations (Message)"
data.Contact.SubmitBtnText = "Broadcast Signal"
data.Contact.MapUrl = "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3022.6175392743824!2d-73.9878531242371!3d40.74844447138905!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c259a9b3117469%3A0xd134e199a405a163!2sEmpire%20State%20Building!5e0!3m2!1sen!2sus!4v1710000000000!5m2!1sen!2sus"

// --- Cart Section ---
data.Cart.Title = "MANIFEST REQUISITION BUFFER"
data.Cart.TriggerLabel = "[SLOT_ALLOCATION]"
data.Cart.RemoveBtnText = "Purge"            // Built-in template fallback match
data.Cart.EmptyText = "No items allocated." // Built-in template fallback match
data.Cart.TotalLabel = "CREDIT DEBT:"
data.Cart.ClearBtnText = "Purge Buffer State"
data.Cart.TotalItems = 0
data.Cart.TotalPrice = 0.00
data.Cart.Items = []CartItem{}

// --- Footer Section ---
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