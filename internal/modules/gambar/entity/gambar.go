package entity

type Gambar struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Reference    string `gorm:"type:varchar(50);not null" json:"reference"`
	ForeignID    int    `gorm:"not null" json:"foreign_id"`
	Path         string `gorm:"type:text;not null" json:"path"`
	OriginalName string `gorm:"type:varchar(255);not null" json:"original_name"`
	MimeType     string `gorm:"type:varchar(100);not null" json:"mime_type"`
}

func (Gambar) TableName() string {
	return "gambars"
}

type GetGambarRequest struct {
	Reference string `query:"reference"`
	ForeignID int    `query:"foreign_id"`
}

type DeleteGambarRequest struct {
	DeleteImageIds []int `json:"delete_image_ids"`
}
