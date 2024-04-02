package repo_impl

import (
	"context"
	"time"

	"github.com/HaAnhTu2/code01/abc"
	"github.com/HaAnhTu2/code01/db"
	"github.com/HaAnhTu2/code01/entities"
	"github.com/HaAnhTu2/code01/log"
	"github.com/HaAnhTu2/code01/reponsitory"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) reponsitory.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}
func (u *UserRepoImpl) SaveUser(context context.Context, user entities.User) (entities.User, error) {
	statement := `
	INSERT INTO users(id, name, email, password, created_at, updated_at)
	Values(:id, :name, :email, :password, :created_at, :updated_at)
	`
	user.Created_At = time.Now()
	user.Updated_At = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, abc.UserConfligct
			}
		}
		return user, abc.SignUpFail
	}
	return user, nil
}

func (u *UserRepoImpl) SaveProduct(context context.Context, product entities.Product) (entities.Product, error) {
	statement := `
	INSERT INTO products(id,name,price,quantity,created_at,update_at)
	Values (:id, :name, :price, :quantity, :created_at, :updated_at)
	`

	product.Created_At = time.Now()
	product.Updated_At = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, product)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return product, abc.UserConfligct
			}
		}
		return product, abc.SignUpFail
	}
	return product, nil
}
