package service

import (
	"github.com/labstack/echo"
	
	"golang_blog_demo_noORM/module"
	"net/http"
	"strconv"
)

type ResultPayload struct {
	State   bool        `json:"state"`
	Content interface{} `json:"contents"`
	Message string      `json:"message"`
}

func _queryAllBlogs(c echo.Context) error {
	//查询返回所有博客
	articles, err := module.GetAllBlogs()
	if err != nil {
		return errorHandler(400, err.Error())
	}
	r := &ResultPayload{
		State:   true,
		Content: articles,
		Message: "success",
	}
	return c.JSON(http.StatusOK, r)
}
func _postComment(c echo.Context) error {
	//对博客评论回复
	comment := c.FormValue("comment")
	id := c.Param("id")
	if comment == "" {
		return errorHandler(400, "comment can't be empty")
	}
	idnum, err := idParser(id)
	if err != nil {
		return errorHandler(400, id+"is not a number")
	}
	err = new(module.Comment).PostCommentById(idnum, comment, c.RealIP()) //插入评论
	if err != nil {
		return errorHandler(http.StatusNotFound, err.Error())
	}
	r := &ResultPayload{
		State:   true,
		Content: comment,
		Message: "success",
	}
	return c.JSON(http.StatusOK, r)
}

func _getComments(c echo.Context) error { //获取文章下所有评论
	articleId := c.Param("id")
	idnum, err := idParser(articleId)
	if err != nil {
		return errorHandler(400, articleId+" is not a number")
	}
	err, comments := module.Comment{}.GetCommentsById(idnum)
	if err != nil {
		return errorHandler(400, err.Error())
	}
	r := &ResultPayload{
		State:   true,
		Content: comments,
		Message: "success",
	}
	return c.JSON(http.StatusOK, r)
}
func createArticle(c echo.Context) error {
	requestBinder := new(module.PreRequestBinder)
	if err := c.Bind(requestBinder); err != nil {
		return err
	}
	if requestBinder.Title == "" {
		return errorHandler(400, "Title can't be empty")
	}
	if len([]rune(requestBinder.Title)) > 50 {
		return errorHandler(400, "Title can't be over 50 chars")
	}
	if requestBinder.Author == "" {
		return errorHandler(400, "Author can't be empty")
	}
	if len([]rune(requestBinder.Author)) > 20 {
		return errorHandler(400, "Author can't be over 20 chars")
	}
	if requestBinder.Content == "" {
		return errorHandler(400, "Content can't be empty")
	}
	article := new(module.Article)
	err, _ := article.NewBlog(requestBinder)
	if err != nil {
		return errorHandler(http.StatusInternalServerError, err.Error())
	}
	err = article.Insert()
	if err != nil {
		return errorHandler(http.StatusInternalServerError, err.Error())
	}
	r := &ResultPayload{
		State:   true,
		Content: nil,
		Message: "success",
	}
	return c.JSON(http.StatusOK, r)
}

func _queryBlog(c echo.Context) error {
	id := c.Param("id")
	idnum, err := idParser(id)
	if err != nil {
		return errorHandler(400, "articleId: "+id+"is not a number")
	}
	article := new(module.Article)
	err = article.GetBlogById(idnum)
	if err != nil {
		return errorHandler(http.StatusOK, "not article")
	}
	r := &ResultPayload{
		State:   true,
		Content: article,
		Message: "success",
	}
	return c.JSON(http.StatusOK, r)
}

func errorHandler(code int, message string) error {
	return echo.NewHTTPError(code, message)
}

func idParser(id string) (int64, error) {
	idnum, err := strconv.ParseInt(id, 10, 64)
	return idnum, err
}
