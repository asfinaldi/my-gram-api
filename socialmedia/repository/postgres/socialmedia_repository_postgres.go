package repository

import (
	"context"
	"fmt"
	"final-project/models"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (socialMediaRepository *socialMediaRepository) Fetch(ctx context.Context, socialMedias *[]models.SocialMedia, userID string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = socialMediaRepository.db.WithContext(ctx).Where("user_id = ?", userID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username", "ProfileImageUrl")
	}).Find(&socialMedias).Error; err != nil {
		return err
	}

	return
}

func (socialMediaRepository *socialMediaRepository) Store(ctx context.Context, socialMedia *models.SocialMedia) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	ID, _ := nanoid.New(16)

	socialMedia.ID = fmt.Sprintf("socialmedia-%s", ID)

	if err = socialMediaRepository.db.WithContext(ctx).Create(&socialMedia).Error; err != nil {
		return err
	}

	return
}

func (socialMediaRepository *socialMediaRepository) GetByID(ctx context.Context, socialMedia *models.SocialMedia, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = socialMediaRepository.db.WithContext(ctx).First(&socialMedia, &id).Error; err != nil {
		return err
	}

	return
}

func (socialMediaRepository *socialMediaRepository) Update(ctx context.Context, socialMedia models.SocialMedia, id string) (socmed models.SocialMedia, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	socmed = models.SocialMedia{}

	if err = socialMediaRepository.db.WithContext(ctx).First(&socmed, &id).Error; err != nil {
		return socmed, err
	}

	if err = socialMediaRepository.db.WithContext(ctx).Model(&socmed).Updates(socialMedia).Error; err != nil {
		return socmed, err
	}

	return socmed, nil
}

func (socialMediaRepository *socialMediaRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = socialMediaRepository.db.WithContext(ctx).First(&models.SocialMedia{}, &id).Error; err != nil {
		return err
	}

	if err = socialMediaRepository.db.WithContext(ctx).Delete(&models.SocialMedia{}, &id).Error; err != nil {
		return err
	}

	return
}
