package module

import (
	"go_blog/utils"
)

type PreRequestBinder struct {
	Title   string `json:"title" form:"title" query:"title"`
	Author  string `json:"author" form:"author" query:"author"`
	Content string `json:"content" form:"content" query:"content"`
}

type Article struct {
	Id        int64      `json:"id" form:"id" query:"id"`
	Title     string     `json:"title" form:"title" query:"title"`
	Author    string     `json:"author" form:"author" query:"author"`
	Timestamp string     `json:"timestamp" form:"timestamp" query:"timestamp"`
	Content   string     `json:"content" form:"content" query:"content"`
	Comments  []*Comment `json:"comments " form:"comments" query:"comments"`
}

type Comment struct {
	Content   string `json:"content" form:"content" query:"content"`
	Timestamp string `json:"timestamp" form:"timestamp" query:"timestamp"`
	FormIP    string `json:"fromip" form:"fromip" query:"fromip"`
}

func (c *Comment) PostCommentById(id int64, message string, ip string) error {
	c.Content = message
	c.Timestamp = utils.GetNowTime()
	c.FormIP = ip
	err := c.insert(id, c)
	return err
}
func (c Comment) GetCommentsById(id int64) (error, []*Comment) {
	return c.queryById(id)
}
func GetAllBlogs() ([]*Article, error) {
	return Article{}.queryALL()
}

func (a *Article) GetBlogById(id int64) error {
	var err error
	err = a.queryById(id)
	if err != nil {
		return err
	}
	err, c := new(Comment).queryById(id)
	a.Comments = c
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) NewBlog(re *PreRequestBinder) (error, *Article) {
	a.Init()
	a.Title = re.Title
	a.Content = re.Content
	a.Author = re.Author
	return nil, a
}

func (a *Article) Init() *Article {
	//a.Id = utils.GenRandoms().Int64()
	a.Timestamp = utils.GetNowTime()
	return a
}
