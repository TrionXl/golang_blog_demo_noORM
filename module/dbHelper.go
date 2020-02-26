package module

import (
	"errors"
)

type articleDB interface {
	Insert() error
	queryALL() ([]*Article, error)
	queryById(id int64)
}

type commentDB interface {
	insert(id int64, comment *Comment) error
	queryById(id int64) (error, []*Comment)
}

func (c *Comment) insert(id int64, comment *Comment) error { //按文章id插入评论
	if !commentIdTest(id) {
		return errors.New("article not exists")
	}
	tx, err := tsStart()
	defer tx.Commit()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO comment(id, fromIP,timestamp,content) VALUES (?,?,?,?)", id, c.FormIP, c.Timestamp, c.Content)
	if err != nil {
		return err
	}
	return nil
}

func (c *Comment) queryById(id int64) (error, []*Comment) { //按文章id获取评论
	tx, _ := tsStart()
	defer tx.Commit()
	rows, err := tx.Query("SELECT fromIP,timestamp,content FROM comment WHERE id = ?", id)
	if err != nil {
		return err, nil
	}
	defer rows.Close()
	commentList := make([]*Comment, 1)
	for rows.Next() {
		_c := new(Comment)
		if err := rows.Scan(&_c.FormIP, &_c.Timestamp, &_c.Content); err != nil {
			//log.Fatal(err)
			return err,nil
		}
		commentList = append(commentList, _c)
	}
	return err, commentList
}

func (a *Article) Insert() error { //插入文章
	if !titleTest(a.Title) {
		return errors.New("article title duplicated")
	}
	tx, err := tsStart()
	defer tx.Commit()
	_, err = tx.Exec("INSERT INTO article(title,author,timestamp,content) VALUES (?,?,?,?)", a.Title, a.Author, a.Timestamp, a.Content)
	if err != nil {
		print(err)
		return err
	}
	return nil
}

func (a Article) queryALL() ([]*Article, error) { //获取所有文章
	tx, _ := tsStart()
	defer tx.Commit()
	rows, err := tx.Query("SELECT id,title,author,timestamp,content FROM article")
	defer rows.Close()
	articleList := make([]*Article, 1)
	for rows.Next() {
		_a := new(Article)
		if err := rows.Scan(&_a.Id, &_a.Title, &_a.Author, &_a.Timestamp, &_a.Content); err != nil {
			//log.Fatal(err)
			return nil,err
		}
		articleList = append(articleList, _a)
	}
	return articleList, err
}
func (a *Article) queryById(id int64) error { //按id获取文章
	tx, _ := tsStart()
	defer tx.Commit()
	rows := tx.QueryRow("SELECT id,title,author,timestamp,content FROM article where id=?", id)
	if err := rows.Scan(&a.Id, &a.Title, &a.Author, &a.Timestamp, &a.Content); err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func commentIdTest(id int64) bool { // 测试目标文章是否存在
	tx, _ := tsStart()
	defer tx.Commit()
	row, err := tx.Query("SELECT id FROM article where id=?", id)
	if err != nil {
		return false
	}
	defer row.Close()
	return row.Next()
}
func titleTest(title string) bool { //检测新文章标题是否重复
	tx, _ := tsStart()
	defer tx.Commit()
	row, err := tx.Query("SELECT title FROM article where title=?", title)
	if err != nil {
		return false
	}
	defer row.Close()
	return !row.Next()
}
