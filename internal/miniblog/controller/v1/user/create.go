package user

import (
	"github.com/asaskevich/govalidator"

	"github.com/cjhw/miniblog/internal/pkg/core"
	"github.com/cjhw/miniblog/internal/pkg/errno"
	"github.com/cjhw/miniblog/internal/pkg/log"
	v1 "github.com/cjhw/miniblog/pkg/api/miniblog/v1"
	"github.com/gin-gonic/gin"
)

// Create 创建一个新的用户.
func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
