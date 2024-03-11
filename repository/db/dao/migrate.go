package dao

import "ChargPiles/repository/db/model"

func migrate() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Pile{}, &model.Station{},
			&model.StationPiles{}, &model.Order{})

	return
}
