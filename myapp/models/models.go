package models

import (
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME   = "beeblog.db"
	_PG_DRIVER = "sqlite3"
)

//分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

//文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Labels          string
	Category        string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	Replytime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

//评论
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_PG_DRIVER, orm.DRPostgres)
	orm.RegisterDataBase("default", _PG_DRIVER, _DB_NAME, 10)

}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return nil
	}

	_, err = o.Insert(cate)

	if err != nil {
		return err
	}
	return nil
}

//获取所有category
func GetAllCategory() (cates []*Category, err error) {
	o := orm.NewOrm()
	cates = make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err = qs.All(&cates)
	return
}
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	var oldCate string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}

	if len(oldCate) > 0 {
		cate := &Category{}
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	return err
}

func GetAllTopic(cate, label string, isDecs bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")
	var err error
	if isDecs {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := &Topic{}
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)

	topic.Labels = strings.Replace(strings.Replace(
		topic.Labels, "#", " ", -1), "$", "", -1)

	return topic, err
}

func AddTopic(title, content, label, category, attachment string) (err error) {
	//
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Category:   category,
		Content:    content,
		Labels:     label,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		Replytime:  time.Now(),
	}

	_, err = o.Insert(topic)
	if err != nil {
		return err
	}

	//更新分类统计
	cate := &Category{}
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		// 不存在错误
		cate.TopicCount++
		_, err = o.Update(cate)
	}

	return err
}

func ModifyTopic(tid, title, content, label, category, attachment string) (err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return
	}
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	var oldCate, oldAttach string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		topic.Labels = label
		topic.Attachment = attachment
		topic.Category = category
		_, err = o.Update(topic)
		if err != nil {
			return
		}
	}

	//更新旧的分类统计
	if len(oldCate) > 0 {
		cate := &Category{}
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	//删除旧的附件
	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}
	cate := &Category{}
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return
}

func AddReply(tid, nickname, content string) (err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Replytime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return

}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	var tidNum int64
	var lenReplies int64
	o := orm.NewOrm()
	reply := &Comment{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}

	replies := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tidNum}
	lenReplies = int64(len(replies))
	if o.Read(topic) == nil && lenReplies > 0 {
		topic.Replytime = replies[0].Created
		topic.ReplyCount = lenReplies
		_, err = o.Update(topic)
	}
	return err
}
