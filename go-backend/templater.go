type PageState struct {
	Meta struct {
		PageTitle     string
		ThemeClass    string // "theme-cyberpunk" or "theme-corporate"
		CurrencyLabel string // "creds" or "USD"
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
		Title          string
		NameLabel      string
		EmailLabel     string
		MessageLabel   string
		SubmitBtnText  string
		MapUrl         string
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