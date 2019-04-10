package userhandlers

import (
	"github.com/jinzhu/gorm"
)

func InitializeTables(db *gorm.DB) {
	if !db.HasTable("user_index_2_1") {

		ct := func(t int) {
			for i := 0; i <= 5; i++ {
				db.CreateTable(&userIndex{Type: t, Code: i})
			}
		}

		ct(emailIndexType)
		ct(phoneIndexType)
		ct(accountIndexType)
	}

}
