package user

import (
	"context"
	"log"
	// "errors"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
	// "payment-svc/pkg/liberror"
	// "payment-svc/pkg/user/model"
	"time"
)

type UserRepository interface {
	AddUser(ctx context.Context,user *User) error
}

type gormUserRepository struct {
	db *gorm.DB
}

func (ur *gormUserRepository) AddUser(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	log.Print("in repository")
	db := ur.db.WithContext(ctx).Create(&user)
	if db.Error != nil {
		return db.Error
	}

	// if db.RowsAffected == 0 {
	// 	return liberror.Builder().SetOperation("CreateBankInfo.db.Create").SetKind(liberror.ResourceConflictError).SetCause(errors.New("duplicate record")).Build()
	// }

	return nil
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}
