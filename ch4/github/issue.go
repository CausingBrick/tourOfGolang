package github

// IssuesURL The github Issues API.
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult The results of issues form github API .
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue store the values for issues.
type Issue struct {
	NUmber     int
	HTMLURL    string `json:"html_url"`
	Title      string
	State      string
	User       *User
	CreateadAt string `json:"Created_at"`
	Body       string //markdown formats.
}

// User store the values for User Info.
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
