package mysql_repository

import (
	"assignee/first-domain-implement/domain/entity"
	"assignee/first-domain-implement/domain/repository"
	"context"
	"database/sql"
	"time"
)

// interaksi dengan DB
type ArticleRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

// build structnya, yang mengacu ke connection dan kontrak interface di repository
func NewArticleRepositoryMysqlInteractor(connectionDatabse *sql.DB) repository.ArticleRepository {
	return &ArticleRepositoryMysqlInteractor{dbConn: connectionDatabse}
}

// implementasi dari interface kontrak dalam bentuk method receiver
func (repo *ArticleRepositoryMysqlInteractor) Store(ctx context.Context, dataArticle *entity.Article) error {
	var (
		errMysql            error
		errMysqlTranslation error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// query insert to table article
	insertQuery := "INSERT INTO article(code_article, title_original, text_original, date, " +
		"banner, author, thumbs, is_highlight) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	_, errMysql = repo.dbConn.Exec(insertQuery, dataArticle.GetCodeArtikel(), dataArticle.GetTitleArtikel(), dataArticle.GetTextArtikel(),
		dataArticle.GetTanggalArtikel(), dataArticle.GetBannerArtikel(), dataArticle.GetAuthorArtikel(), dataArticle.GetThumbsArtikel(), dataArticle.GetIsHighLightArtikel())

	if errMysql != nil {
		return errMysql
	}

	// query insert to table translation
	insertQueryTranslation := "INSERT INTO translation(code_article, code_language, title, text) VALUES(?, ?, ?, ?)"

	_, errMysqlTranslation = repo.dbConn.Exec(insertQueryTranslation, dataArticle.GetCodeArtikel(), dataArticle.GetTranslationCodeLanguage(), dataArticle.GetTranslationTitle(),
		dataArticle.GetTranslationText())

	if errMysqlTranslation != nil {
		return errMysqlTranslation
	}

	return nil
}
