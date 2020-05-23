package repository

import (
	"gin-essential/common"
	"gin-essential/model"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}

	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) DeleteById(id int) error {
	var category model.Category
	if err := c.DB.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}

func (c CategoryRepository) GetList(pageNum, pageSize int) ([]*model.Category, error) {
	var categories []*model.Category
	if err := c.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c CategoryRepository) Count() int {
	var count int
	c.DB.Model(model.Category{}).Count(&count)
	return count
}
