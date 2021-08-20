package service

import (
	"context"
	"myapp/graph/model"
)

var users []*model.User = make([]*model.User, 0)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	users = append(users, (*model.User)(&input))
	return (*model.User)(&input), nil
}

func UserGetAll(ctx context.Context) ([]*model.User, error) {
	return users, nil
}
