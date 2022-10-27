package utils

import (
	"time"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type FetchedPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *User      `json:"user"`
}

type ResponseDataFetchedPhoto struct {
	Status string         `json:"status"`
	Data   []FetchedPhoto `json:"data"`
}

type AddPhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type AddedPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type ResponseDataAddedPhoto struct {
	Status string     `json:"status"`
	Data   AddedPhoto `json:"data"`
}

type UpdatePhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type UpdatedPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ResponseDataUpdatedPhoto struct {
	Status string       `json:"status"`
	Data   UpdatedPhoto `json:"data"`
}

type ResponseMessageDeletedPhoto struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseMessage struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
