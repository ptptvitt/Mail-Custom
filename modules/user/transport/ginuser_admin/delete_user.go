package ginuser_admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task1/component"
	"task1/modules/user/biz/adminc_role"
	usermodel "task1/modules/user/model"
	storageuser "task1/modules/user/storage"
)

func DeleteUserByAdmin(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		user_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := storageuser.NewSQLStore(appCtx.GetMainDbConnection())
		biz := adminc_role.NewDeleteUserBiz(store)
		if err := biz.DeleteUser(c.Request.Context(), user_id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"mail": data.Email})
	}
}