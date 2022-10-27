package usecase

import (
	"context"
	"final-project/models"
)

type userUseCase struct {
	userRepository models.UserRepository
}

func NewUserUseCase(userRepository models.UserRepository) *userUseCase {
	return &userUseCase{userRepository}
}

func (userUseCase *userUseCase) Register(ctx context.Context, user *models.User) (err error) {
	if err = userUseCase.userRepository.Register(ctx, user); err != nil {
		return err
	}

	return
}

func (userUseCase *userUseCase) Login(ctx context.Context, user *models.User) (err error) {
	if err = userUseCase.userRepository.Login(ctx, user); err != nil {
		return err
	}

	return
}

func (userUseCase *userUseCase) Update(ctx context.Context, user models.User) (u models.User, err error) {
	if u, err = userUseCase.userRepository.Update(ctx, user); err != nil {
		return u, err
	}

	return u, nil
}

func (userUseCase *userUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = userUseCase.userRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
