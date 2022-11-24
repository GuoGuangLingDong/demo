package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"

	"demo/internal/dao"
	"demo/internal/model"
	"demo/internal/model/do"
	"demo/internal/model/entity"
	"demo/internal/service"
)

type (
	SUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() *SUser {
	return &SUser{}
}

// Create creates user account.
func (s *SUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, generate one
	if in.Nickname == "" {
		in.Nickname = fmt.Sprintf("wesoul-%v", in.UId[:6])
	}
	// Username checks.
	available, err := s.UsernameLegalCheck(ctx, in.Username)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Uid:         in.UId,
			PhoneNumber: in.PhoneNumebr,
			Username:    in.Username,
			Password:    in.Password,
			Nickname:    in.Nickname,
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *SUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *SUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Username: in.Username,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`Passport or Password not correct`)
	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
	})
	return nil
}

// SignOut removes the session for current signed-in user.
func (s *SUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// UsernameLegalCheck checks and returns given nickname is available for signing up.
func (s *SUser) UsernameLegalCheck(ctx context.Context, username string) (bool, error) {
	res, err := dao.User.Ctx(ctx).One(do.User{
		Username: username,
	})
	if err != nil {
		return false, err
	}
	return len(res) == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *SUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx)
}
