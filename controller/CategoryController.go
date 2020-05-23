package controller

import (
	"gin-essential/model"
	"gin-essential/repository"
	"gin-essential/response"
	"gin-essential/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

//CategoryController实现了ICategoryController接口，必须实现它里面的所有方法
type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "添加成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取path中的参数,将字符串强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	//更新分类
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数，字符串强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "展示成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数,将字符串强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}
	response.Success(ctx, nil, "删除成功")
}

func (c CategoryController) PageList(ctx *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	//分页查询
	categories, err := c.Repository.GetList(pageNum, pageSize)
	if err != nil {
		response.Fail(ctx, nil, "查询失败")
	}

	//查询记录总条数
	count := c.Repository.Count()

	response.Success(ctx, gin.H{"categories": categories, "count": count}, "查询成功")
}
