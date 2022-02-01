package database

import (
	"bloggin/entity/admin"
)

var adminUser admin.Admin

func GetAdminUser() admin.Admin {
	return adminUser
}
