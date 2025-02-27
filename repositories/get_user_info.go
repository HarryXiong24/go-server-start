package repositories

import (
	"fmt"
	"go-server-start/db"
	"go-server-start/models"
)

func GetUserInfo(name string) (*models.User, error) {
	var res models.User
	q := db.GetTableUser()

	if name != "" {
		q = q.Select("Id, name").Where("name = ? ", name)

		if err := q.Find(&res).Error; err != nil {
			fmt.Printf("Fail to query, err: %v\n", err)
			return nil, err
		}
	}

	return &res, nil
}
