package main

import (
	"assignee/first-domain-implement/domain/entity"
	"assignee/first-domain-implement/domain/repository"
	"assignee/first-domain-implement/mysql_repository"
	"assignee/first-domain-implement/pkg/mysql_connection"
	"context"
	"fmt"
)

var (
	connectionDatabase     = mysql_connection.InitMysqlDB()
	articleRepositoryMysql = mysql_repository.NewArticleRepositoryMysqlInteractor(connectionDatabase)
)

type ArticleLogicFactoryHandler struct {
	articleRepository repository.ArticleRepository
}

func NewArticleLogicFactoryHandler(repoArticleImplementation repository.ArticleRepository) *ArticleLogicFactoryHandler {
	return &ArticleLogicFactoryHandler{articleRepository: repoArticleImplementation}
}

func main() {
	ctx := context.Background()
	CreateFirstArticle := entity.DTONewCreateArticle{
		TitleOriginal: "Apa itu Lorem Ipsum?",
		TextOriginal:  "Lorem Ipsum hanyalah teks tiruan dari industri percetakan dan penyusunan huruf. Lorem Ipsum telah menjadi teks dummy standar industri sejak tahun 1500-an, ketika seorang pencetak yang tidak dikenal mengambil sekumpulan tipe dan mengacaknya untuk membuat buku spesimen tipe. Ini telah bertahan tidak hanya lima abad, tetapi juga lompatan ke pengaturan huruf elektronik, pada dasarnya tetap tidak berubah. Itu dipopulerkan pada 1960-an dengan merilis lembar Letraset yang berisi bagian-bagian Lorem Ipsum, dan baru-baru ini dengan perangkat lunak penerbitan desktop seperti Aldus PageMaker termasuk versi Lorem Ipsum.",
		Banner:        "example.jpg",
		Author:        "Taupik Pirdian",
		Thumbs:        "thumbs-example.jpg",
		IsHighLight:   false,
		Translation: []entity.DTOTranslation{
			{
				CodeLanguage: "ENG",
				Title:        "What is Lorem Ipsum?",
				Text:         "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			},
			{
				CodeLanguage: "GER",
				Title:        "Was ist Lorem Ipsum?",
				Text:         "Lorem Ipsum ist einfach Blindtext der Druck- und Satzindustrie. Lorem Ipsum ist seit den 1500er Jahren der Standard-Dummy-Text der Branche, als ein unbekannter Drucker eine Reihe von Typen nahm und daraus ein Musterbuch für Typen erstellte. Sie hat nicht nur fünf Jahrhunderte, sondern auch den Sprung in den elektronischen Satz überstanden und ist im Wesentlichen unverändert geblieben. Es wurde in den 1960er Jahren mit der Veröffentlichung von Letraset-Blättern mit Passagen von Lorem Ipsum und in jüngerer Zeit mit Desktop-Publishing-Software wie Aldus PageMaker, einschließlich Versionen von Lorem Ipsum, populär.",
			},
		},
	}

	// create dan generate id code number dan year modeling
	FirstArticle, errCheckDomainArticle := entity.NewCreateArticle(CreateFirstArticle)
	if errCheckDomainArticle != nil {
		fmt.Println("GAGAL CREATE ARTICLE KARENA WRONG DOMAIN")
		panic(errCheckDomainArticle)

	}

	handlerRepo := NewArticleLogicFactoryHandler(articleRepositoryMysql)
	errStoreRepo := handlerRepo.articleRepository.Store(ctx, FirstArticle)
	if errStoreRepo != nil {
		fmt.Println("GAGAL CREATE ARTICLE ADA KESALAHAN DALAM PENYIMPANAN")
		panic(errStoreRepo)
	}
}
