package database

import (
	"bloggin/internal/entity/admin"
)

var adminUser admin.Admin

func GetAdminUser() admin.Admin {
	return adminUser
}
