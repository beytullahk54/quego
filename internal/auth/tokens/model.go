package tokens

type Token struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  string `json:"user_id"`  // Gönderilecek endpoint
	TokenID string `json:"token_id"` // Bu job hangi API token ile ilişkili
}
