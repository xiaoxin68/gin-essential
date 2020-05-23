package repository

import (
	"gin-essential/common"
	"gin-essential/model"
	"gin-essential/vo"
	"github.com/jinzhu/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository() PostRepository {
	return PostRepository{DB: common.GetDB()}
}

func (p PostRepository) Create(post model.Post) (*model.Post, error) {
	if err := p.DB.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (p PostRepository) Exist(id string) bool {
	var post model.Post
	if p.DB.Where("id=?", id).First(&post).RecordNotFound() {
		return false
	}
	return true
}

func (p PostRepository) SelectById(id string) *model.Post {
	var post model.Post
	if p.DB.Preload("Category").Where("id=?", id).First(&post).RecordNotFound() {
		return nil
	}
	return &post
}

func (p PostRepository) Update(post model.Post, requestPost vo.CreatePostRequest) (*model.Post, error) {
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (p PostRepository) Delete(id string) error {
	var post model.Post
	if err := p.DB.Where("id=?", id).Delete(&post).Error; err != nil {
		return err
	}
	return nil
}

func (p PostRepository) GetList(pageNum, pageSize int) ([]*model.Post, error) {
	var posts []*model.Post
	if err := p.DB.Preload("Category").Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p PostRepository) Count() int {
	var count int
	p.DB.Model(model.Post{}).Count(&count)
	return count
}
