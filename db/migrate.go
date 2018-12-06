package db

import "github.com/y-ogura/gorm-transaction/account"

// Migrate migration db
func Migrate() {
	DB.AutoMigrate(
		&account.Account{},
	)
}

// DropTable drop table
func DropTable() {
	DB.DropTableIfExists(
		&account.Account{},
	)
}
