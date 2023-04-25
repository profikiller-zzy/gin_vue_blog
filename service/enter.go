package service

import (
	"gin_vue_blog_AfterEnd/service/image_service"
	"gin_vue_blog_AfterEnd/service/user_service"
)

type Service struct {
	UserServiceApp  user_service.UserService
	ImageServiceApp image_service.ImageService
}

var ServiceApp = new(Service)
