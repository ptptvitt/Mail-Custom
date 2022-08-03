package storagerole

import (
	"context"
	"task1/common"
	"task1/modules/role/model_role"
)

func (s *sqlStore) CreateRole(ctx context.Context, data *model_role.Role) error {
	db := s.db

	// create new role
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
