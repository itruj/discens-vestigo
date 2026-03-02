package store

type Card struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type CardStore interface {
	CreateCard(title string, done bool) error
	GetCard(title string) (*Card, error)
}
