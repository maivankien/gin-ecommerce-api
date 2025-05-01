package routers

import (
	"github.com/maivankien/go-ecommerce-api/internal/routers/manage"
	"github.com/maivankien/go-ecommerce-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
