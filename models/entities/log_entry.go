package entities

type LogEntry struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"not null"`
	Action    string `gorm:"not null"`
	Route     string `gorm:"not null"`
	Timestamp int64  `gorm:"not null"`
	// Add more fields as needed based on your logging requirements
}
