package todo

// Todo strcture
type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TableName return
func (b *Todo) TableName() string {
	return "todo"
}
