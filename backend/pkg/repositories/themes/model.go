package theme_repo

type Model struct {
	BackgroundColor string `json:"backgroundColor"`
	TextColor       string `json:"textColor"`
	PrimaryColor    string `json:"primaryColor"`
	SecondaryColor  string `json:"secondaryColor"`
	ErrorColor      string `json:"errorColor"`
	SuccessColor    string `json:"successColor"`
}
