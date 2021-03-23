package user
import (
	// "context"
	"context"
	"log"
)

type Service interface {
	CreateUser(ctx context.Context, user *User) (User, error)
}

type userService struct {
	repo            UserRepository;
	// ctx 			context.Context

}

func (uss *userService) CreateUser(ctx context.Context, user *User) (User, error) {
	// res, err := bis.repository.GetBankURL(ctx, bankCode)
	// if err != nil {
	// 	return "", liberror.Builder().SetOperation("Service.GetBankURL").SetCause(err).Build()
	// }
	log.Printf("first Name  "+user.FirstName)
	log.Printf("last Name  "+user.LastName)
	// dc:=config.NewDBConfig()
	// ndh,_:=db.NewDBHandler(dc).GetDB()
	uss.repo.AddUser(ctx,user)
	return *user, nil

}
// func NewUserService(ctx context.Context) Service {
// 	return &userService{
// 		ctx : ctx,
// 	}
// }

func NewUserService(repo UserRepository) Service {
	return &userService{
		repo: repo,
			}	
}