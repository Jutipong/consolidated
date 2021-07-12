package repository

import "consolidated/feature/inwardfee/model"

type InwardFee struct{}

func (i *InwardFee) FindById(u *[]model.Request) (err error) {
	// rows, err := config.DB.Table("User").Find(&u).Rows()
	// if err = config.DB.Table("User").Find(&u).Error; err != nil {
	// 	return err
	// }
	// defer rows.Close()
	return nil
}
