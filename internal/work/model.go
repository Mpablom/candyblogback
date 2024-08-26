package work

type Work struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Gallery     []Gallery `json:"gallery" gorm:"foreignKey:WorkID"`
}

type Gallery struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	WorkID string `json:"work_id"`
	Image  uint   `json:"image"`
}
