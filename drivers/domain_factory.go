package drivers

import (
	userDB "Capstone/drivers/mysql/users"
	userDomain "Capstone/models/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
