package storage

import (
	"Lists-app/internal/storage/api/notification"
	user22 "Lists-app/internal/storage/api/user"
)

type Storage interface {
	User() user22.User
	Notification() notification.Notification
}
