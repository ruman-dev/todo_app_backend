package todos

type Todos struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}
