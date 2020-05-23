package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

//使用uuid作为主键
type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category   *Category //关联查询  //注意：这里的外键是符合命名规范的，所以不用指定，默认是CategoryId。否则需要自己指定外键
	Title      string    `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string    `json:"head_img"`
	Content    string    `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time      `json:"updated_at" gorm:"type:timestamp"`
}

func (p *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
