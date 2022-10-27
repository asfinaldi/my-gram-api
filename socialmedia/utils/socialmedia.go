package utils

import "time"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SocialMedia struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         string     `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	User           *User      `json:"user"`
}

type SocialMedias struct {
	SocialMedias []SocialMedia `json:"social_medias"`
}

type FetchedSocialMedia struct {
	SocialMedias interface{} `json:"social_medias"`
}

type ResponseDataFetchedSocialMedia struct {
	Status string       `json:"status"`
	Data   SocialMedias `json:"data"`
}

type AddSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type AddedSocialMedia struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         string     `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type ResponseDataAddedSocialMedia struct {
	Status string           `json:"status"`
	Data   AddedSocialMedia `json:"data"`
}

type UpdateSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type UpdatedSocialMedia struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         string     `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type ResponseDataUpdatedSocialMedia struct {
	Status string             `json:"status"`
	Data   UpdatedSocialMedia `json:"data"`
}

type ResponseMessageDeletedSocialMedia struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseMessage struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
