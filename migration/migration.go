package migration

import (
	"myapp/config"
	"myapp/graph/model"
)

func MigrateTable() {
	db := config.GetDb()

	db.AutoMigrate(&model.User{})
}
