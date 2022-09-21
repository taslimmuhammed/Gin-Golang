package service

import (
	entity "taslimmuhammed/golang-Gin/Entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	Videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(Video entity.Video) entity.Video {
	service.Videos = append(service.Videos, Video)
	return Video
}

func (service *videoService) FindAll() []entity.Video {
	return service.Videos
}
