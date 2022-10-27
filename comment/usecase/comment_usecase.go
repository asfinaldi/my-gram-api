package usecase

import (
	"context"
	"final-project/models"
)

type commentUseCase struct {
	commentRepository models.CommentRepository
}

func NewCommentUseCase(commentRepository models.CommentRepository) *commentUseCase {
	return &commentUseCase{commentRepository}
}

func (commentUseCase *commentUseCase) Fetch(ctx context.Context, comments *[]models.Comment, userID string) (err error) {
	if err = commentUseCase.commentRepository.Fetch(ctx, comments, userID); err != nil {
		return err
	}

	return
}

func (commentUseCase *commentUseCase) Store(ctx context.Context, comment *models.Comment) (err error) {
	if err = commentUseCase.commentRepository.Store(ctx, comment); err != nil {
		return err
	}

	return
}

func (commentUseCase *commentUseCase) GetByID(ctx context.Context, comment *models.Comment, id string) (err error) {
	if err = commentUseCase.commentRepository.GetByID(ctx, comment, id); err != nil {
		return err
	}

	return
}

func (commentUseCase *commentUseCase) Update(ctx context.Context, comment models.Comment, id string) (photo models.Photo, err error) {
	if photo, err = commentUseCase.commentRepository.Update(ctx, comment, id); err != nil {
		return photo, err
	}

	return photo, nil
}

func (commentUseCase *commentUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = commentUseCase.commentRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
