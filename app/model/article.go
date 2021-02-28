package model

import (
	"github.com/jinzhu/gorm"
	"github.com/maogou/ginapi/pkg/app"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleRow struct {
	ArticleId     uint32
	TagId         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (a Article) TableName() string {
	return "blog_article"
}

//针对swagger生成文档包含别的结构体无法显示的bug
type ArticleSwagger struct {
	List  []*Article
	pager *app.Page
}

//创建文章
func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

//更新文章
func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Updates(values).Where("id = ? AND is_del = 0", a.ID).Error; err != nil {
		return err
	}

	return nil
}

//获取文章
func (a Article) Get(db *gorm.DB) (Article, error) {

	var article Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

//删除文章
func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

//关联查询
func (a Article) ListByTagId(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article.title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}

	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+"AS ar").Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").Where("at.tag_id = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []*ArticleRow

	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleId, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagName); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").Where("at.tag_id =? AND ar.state = ? AND is_del = ?", tagID, a.State, 0).Count(&count).Error
	if err != nil {
		return 0, nil
	}

	return count, nil
}
