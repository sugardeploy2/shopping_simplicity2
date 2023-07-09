package repositories

import (
	"fmt"
	"shop_khordad/models/entities"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoggingRepository struct {
	DB *gorm.DB
}

func NewLoggingRepository(db *gorm.DB) *LoggingRepository {
	return &LoggingRepository{
		DB: db,
	}
}

func (lr *LoggingRepository) SaveLogEntry(logEntry *entities.LogEntry) error {
	result := lr.DB.Create(logEntry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lr *LoggingRepository) GetLogsByUserID(userID int) ([]entities.LogEntry, error) {
	var logs []entities.LogEntry
	result := lr.DB.Where("user_id = ?", userID).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func (lr *LoggingRepository) GetLogsByAction(action string) ([]entities.LogEntry, error) {
	var logs []entities.LogEntry
	result := lr.DB.Where("action = ?", action).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func (lr *LoggingRepository) GetLogsByRoute(route string) ([]entities.LogEntry, error) {
	var logs []entities.LogEntry
	result := lr.DB.Where("route = ?", route).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func (lr *LoggingRepository) GetLogsByTimestampRange(start, end int64) ([]entities.LogEntry, error) {
	var logs []entities.LogEntry
	result := lr.DB.Where("timestamp >= ? AND timestamp <= ?", start, end).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func (lr *LoggingRepository) DeleteLog(logEntry *entities.LogEntry) error {
	result := lr.DB.Delete(logEntry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (lr *LoggingRepository) GetUserIDFromSession(c *gin.Context) (uint, error) {
	// Retrieve the user ID from the session
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		return 0, fmt.Errorf("User ID not found in session")
	}

	// Convert the user ID to an unsigned integer
	userIDUint, err := strconv.ParseUint(fmt.Sprintf("%v", userID), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid user ID")
	}

	return uint(userIDUint), nil
}
