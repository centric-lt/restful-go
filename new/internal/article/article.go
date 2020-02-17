package article

type Article struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}

type Articles []Article
