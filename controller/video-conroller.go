package controller

import (
	"net/http"
	entity "taslimmuhammed/golang-Gin/Entity"
	service "taslimmuhammed/golang-Gin/Service"
	Validators1 "taslimmuhammed/golang-Gin/Validators1"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct{
	service service.VideoService
}

var validate *validator.Validate
func New(s1 service.VideoService) VideoController{
	validate = validator.New()
	validate.RegisterValidation("is-cool", Validators1.ValidateCoolTitle)
	return controller{
		service: s1,
	}
}

func (c1 controller) FindAll() []entity.Video{
	return c1.service.FindAll()
}
func (c1 controller) Save(ctx *gin.Context) error{
	var video  entity.Video
	err := ctx.ShouldBindJSON(&video) //checkin  for json type
	if err != nil{
		return err
	}
	err = validate.Struct(video)
	if err != nil{
		return err
	}
	c1.service.Save(video)
	return nil
}

func (c controller) ShowAll(ctx *gin.Context){
	videos := c.service.FindAll()
	data := gin.H{
		"title":"Video Page",
		"videos":videos,
	}
	ctx.HTML(http.StatusOK, "index.html",data)
}