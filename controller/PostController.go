package controller

import (
	"gin-essential/model"
	"gin-essential/repository"
	"gin-essential/response"
	"gin-essential/vo"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
}

type PostController struct {
	Repository repository.PostRepository
}

func NewPostController() IPostController {
	repository := repository.NewPostRepository()
	repository.DB.AutoMigrate(model.Post{})
	return PostController{Repository: repository}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	//数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "数据验证错误 ")
		return
	}

	//获取登录用户 user
	user, _ := ctx.Get("user")

	//创建post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	//插入数据
	if _, err := p.Repository.Create(post); err != nil {
		panic(err)
		return
	}

	response.Success(ctx, nil, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestPost vo.CreatePostRequest

	//数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "数据验证错误 ")
		return
	}

	//获取path中的id
	postID := ctx.Params.ByName("id")

	//根据id查询post
	post := p.Repository.SelectById(postID)
	if post == nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	//获取登录用户 user
	user, _ := ctx.Get("user")

	//判断当前用户是否为文章的作者
	if post.UserId != user.(model.User).ID {
		response.Fail(ctx, nil, "文章不属于您，没有权限操作")
		return
	}

	//更新文章内容
	post, err := p.Repository.Update(*post, requestPost)
	if err != nil {
		response.Fail(ctx, nil, "更新文章失败")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "修改成功")
}

func (p PostController) Show(ctx *gin.Context) {
	//获取path中的id
	postId := ctx.Params.ByName("id")

	post := p.Repository.SelectById(postId)
	if post == nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "查询成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	//获取path中的id
	postId := ctx.Params.ByName("id")

	//判断文章是否存在
	post := p.Repository.SelectById(postId)
	if post == nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	//获取登录用户 user
	user, _ := ctx.Get("user")

	//判断当前用户是否为文章的作者
	if post.UserId != user.(model.User).ID {
		response.Fail(ctx, nil, "文章不属于您，没有权限操作")
		return
	}

	//删除文章
	err := p.Repository.Delete(postId)
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
	}

	response.Success(ctx, nil, "删除成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	//分页查询
	posts, err := p.Repository.GetList(pageNum, pageSize)
	if err != nil {
		response.Fail(ctx, nil, "查询失败")
	}

	//查询记录总条数
	count := p.Repository.Count()

	response.Success(ctx, gin.H{"posts": posts, "count": count}, "查询成功")
}
