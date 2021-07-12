package repository

import "consolidated/feature/outwardfee/model"

type OutwardFee struct{}

func (i *OutwardFee) FindById(u *[]model.Request) (err error) {
	// rows, err := config.DB.Table("User").Find(&u).Rows()
	// if err = config.DB.Table("User").Find(&u).Error; err != nil {
	// 	return err
	// }
	// defer rows.Close()
	return nil
}
