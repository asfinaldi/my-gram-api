package usecase

import (
	"context"
	"final-project/models"
)

type photoUseCase struct {
	photoRepository models.PhotoRepository
}

func NewPhotoUseCase(photoRepository models.PhotoRepository) *photoUseCase {
	return &photoUseCase{photoRepository}
}

func (photoUseCase *photoUseCase) Fetch(ctx context.Context, photos *[]models.Photo) (err error) {
	if err = photoUseCase.photoRepository.Fetch(ctx, photos); err != nil {
		return err
	}

	return
}

func (photoUseCase *photoUseCase) Store(ctx context.Context, photo *models.Photo) (err error) {
	if err = photoUseCase.photoRepository.Store(ctx, photo); err != nil {
		return err
	}

	return
}

func (photoUseCase *photoUseCase) GetByID(ctx context.Context, photo *models.Photo, id string) (err error) {
	if err = photoUseCase.photoRepository.GetByID(ctx, photo, id); err != nil {
		return err
	}

	return
}

func (photoUseCase *photoUseCase) Update(ctx context.Context, photo models.Photo, id string) (p models.Photo, err error) {
	if p, err = photoUseCase.photoRepository.Update(ctx, photo, id); err != nil {
		return p, err
	}

	return p, nil
}

func (photoUseCase *photoUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = photoUseCase.photoRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
