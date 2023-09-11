package service

import (
	"context"
	"gostore/entity"
	userError "gostore/helper/domain/errorModel"
	"gostore/middleware"
	"gostore/repo"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Repo repo.UserRepository
}

type UserService interface {
	GetUser(ctx context.Context, id int, username string) (entity.UserResponse, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	ChangePasswordUser(ctx context.Context, userChange entity.UserChangePassword) error
	DeleteUser(ctx context.Context) error
}

func NewUserService(r repo.UserRepository) UserService {
	return &userService{
		Repo: r,
	}
}

func (u *userService) GetUser(ctx context.Context, id int, username string) (entity.UserResponse, error) {
	return u.Repo.GetUser(ctx, id, username)
}

func (u *userService) CreateUser(ctx context.Context, user *entity.User) error {
	var (
		detailError = make(map[string]any)
	)

	if user.Name == "" {
		detailError["name"] = "this field is missing input"
	}

	if user.Username == "" {
		detailError["username"] = "this field is missing input"
	} else if len(user.Username) > 50 {
		detailError["username"] = "username lenght max 50"
	} else if isCorrectUsername(user.Username) {
		detailError["username"] = "username only alphabets and numberic allowed"
	}

	if user.Password == "" {
		detailError["password"] = "this field is missing input"
	} else if !isCorrectPassword(user.Password) {
		detailError["password"] = "need at least one number, lowercase, uppercase, and special character"
	} else if len(user.Password) < 8 || len(user.Password) > 30 {
		detailError["password"] = "minimum length is 8 and maximum length is 30"
	}

	if user.Email == "" {
		detailError["email"] = "this field is missing input"
	}

	if user.Phonenumber == nil {
		detailError["phonenumber"] = "this field is missing input"
	} else if *user.Phonenumber < 0 {
		detailError["phonenumber"] = "this field must not negative"
	}

	if len(detailError) > 0 {
		return userError.ErrUserInput.AttachDetail(detailError)
	}

	// generate password entry to be encrypt
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPwd)
	user.CreateAt = time.Now()

	err := u.Repo.CreateUser(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			return userError.ErrExistUser
		}
		return err
	}
	return nil
}

func (u *userService) UpdateUser(ctx context.Context, user *entity.User) error {
	var (
		detailError = make(map[string]any)
		id          = ctx.Value(middleware.GOSTORE_USERID).(int)
	)

	if user.Name == "" {
		detailError["name"] = "this field is missing input"
	}

	if user.Username == "" {
		detailError["username"] = "this field is missing input"
	}

	if user.Email == "" {
		detailError["email"] = "this field is missing input"
	}

	if user.Phonenumber == nil {
		detailError["phonenumber"] = "this field is missing input"
	} else if *user.Phonenumber < 0 {
		detailError["phonenumber"] = "this field must not negative"
	}

	if len(detailError) > 0 {
		return userError.ErrUserInput.AttachDetail(detailError)
	}

	checkUser, err := u.Repo.GetUser(ctx, id, "")
	if err != nil {
		return err
	} else if checkUser.Id != id {
		return userError.ErrInvalidUser
	}

	return u.Repo.UpdateUser(ctx, id, user)
}

func (u *userService) ChangePasswordUser(ctx context.Context, userChange entity.UserChangePassword) error {
	var (
		detailError = make(map[string]any)
		userId      = ctx.Value(middleware.GOSTORE_USERID).(int)
	)

	if userChange.OldPassword == "" {
		detailError["old_password"] = "this field is missing input"
	}

	if userChange.NewPassword == "" {
		detailError["new_password"] = "this field is missing input"
	} else if !isCorrectPassword(userChange.NewPassword) {
		detailError["password"] = "need at least one number, lowercase, uppercase, and special character"
	} else if len(userChange.NewPassword) < 8 || len(userChange.NewPassword) > 30 {
		detailError["password"] = "minimum length is 8 and maximum length is 30"
	}

	if len(detailError) > 0 {
		return userError.ErrUserInput.AttachDetail(detailError)
	}

	checkUser, err := u.Repo.GetUser(ctx, userId, "")
	if err != nil {
		return err
	}

	// validation old password
	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(userChange.OldPassword)); err != nil {
		return userError.ErrWrongPassword
	}

	// generate hash of new password
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(userChange.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userChange.NewPassword = string(hashPwd)
	return u.Repo.ChangePasswordUser(ctx, userId, userChange.NewPassword)
}

func (u *userService) DeleteUser(ctx context.Context) error {
	id := ctx.Value(middleware.GOSTORE_USERID).(int)
	checkUser, err := u.Repo.GetUser(ctx, id, "")
	if err != nil {
		return err
	} else if checkUser.Id != id {
		return userError.ErrInvalidUser
	}

	return u.Repo.DeleteUser(ctx, id)
}

// Username validation, username only allowed alpabhet and numberik
func isCorrectUsername(s string) bool {
	for _, r := range s {
		// jika username bukan alphabet atau numerik, maka tidak valid
		if !unicode.IsLetter(r) || !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}

// Password validation, password must have minimum 1 number, lowercase, uppercase, and special character
func isCorrectPassword(s string) bool {
	numeric := regexp.MustCompile(`[0-9]`).MatchString(s)
	lowcase := regexp.MustCompile(`[a-z]`).MatchString(s)
	upcase := regexp.MustCompile(`[A-Z]`).MatchString(s)
	specialChar := func(s string) bool {
		for _, c := range s {
			if unicode.IsPunct(c) || unicode.IsMark(c) {
				return true
			}
		}
		return false
	}

	// validation
	if !numeric {
		return false
	}

	if !lowcase {
		return false
	}

	if !upcase {
		return false
	}

	if !specialChar(s) {
		return false
	}

	return true
}
