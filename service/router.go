package service

import (
	"github.com/labstack/echo"
	"net/http"
)

//CURD api映射
const (
	index         = "/"            //get 首页
	enterNewBlog  = "/postarticle" //post 创建新博客文章
	queryAllBlogs = "/articles"    //get 请求所有文章列表
	queryBlog     = "/article/:id" //get 请求单独一页文章
	getComments   = "/comment/:id" //get 取得文章评论
	postComment   = "/comment/:id" //post 对当前id文章评论
)

func Router(e *echo.Echo) {
	e.GET(index, func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome!")
	})
	e.GET(queryAllBlogs, _queryAllBlogs)
	e.POST(postComment, _postComment)
	e.GET(queryBlog, _queryBlog)
	e.GET(getComments, _getComments)
	e.POST(enterNewBlog, createArticle)
}
