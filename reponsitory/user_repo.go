package reponsitory

import (
	"context"

	"github.com/HaAnhTu2/code01/entities"
)

type UserRepo interface {
	SaveUser(context context.Context, user entities.User) (entities.User, error)
}
