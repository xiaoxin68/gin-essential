package controller

import (
	"gin-essential/common"
	"gin-essential/dto"
	"gin-essential/model"
	"gin-essential/response"
	"gin-essential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

//注册
func Register(ctx *gin.Context) {
	DB := common.GetDB() //获取数据库引擎
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	//获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称为空给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//如果电话号码存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//对密码进行加密
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasPassword),
	}

	DB.Create(&newUser)

	//发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

//登录
func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		//c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		//log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "登陆成功")
}

//获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	//dto.ToUserDao:用来优化用户信息的返回，因为有些信息不应该显示
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDao(user.(model.User))}})
}
