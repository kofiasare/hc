package checkout

type store struct {
	Name        string
	Tagline     string
	PostAddress string
	Phone       string
	LogoURL     string
	WebsiteURL  string
}

func (s *store) info() *store {
	return s
}
