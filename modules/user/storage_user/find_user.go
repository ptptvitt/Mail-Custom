package storageuser

import (
	"context"
	"gorm.io/gorm"
	"task1/common"
	usermodel "task1/modules/user/model_user"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error) {
	db := s.db

	var data usermodel.User
	// find user,
	if err := db.Table(data.TableName()).First(&data, conditions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}