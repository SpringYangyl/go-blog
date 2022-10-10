package dao

type Article struct {
	Id          int64
	UserId      int64  `json:"user_id" gorm:"column:user_id"'`
	ArticleName string `json:"article_name" gorm:"column:article_name"`
	Description string `json:"description" gorm:"column:description"`
	Content     string `json:"content" gorm:"column:content"`
	LikeCount   int64  `json:"like_count" gorm:"column:like_count"`
	WatchCount  int64  `json:"watch_count" gorm:"watch_count"`
}

func GetArticleContent(id int64) (Article, error) {
	article := Article{
		Id: id,
	}
	if err := Db.First(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}
func GetAllArticle() ([]Article, error) {
	var article []Article
	if err := Db.Find(&article).Error; err != nil {
		return article, err
	}
	return article, nil
}
func CreateArticle(article Article) bool {
	if err := Db.Create(&article).Error; err != nil {
		return false
	}
	return true
}
func (article Article) TableName() string {
	return "article"
}
