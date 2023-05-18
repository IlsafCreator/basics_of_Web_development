package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
	MostRecent    []recentPostData
}

type postData struct {
	Title     string `db:"title"`
	Subtitle  string `db:"subtitle"`
	PostImage string `db:"image_url"`
	Content   string `db:"content"`
}

type featuredPostData struct {
	ID          string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostURL     string
}
type recentPostData struct {
	ID          string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	PostImage   string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostURL     string
}

// type orderListData struct {
// 	OrderID string `db:"order_id"`
// 	Title   string `db:"title"`
// }

// type orderData struct {
// 	Title   string `db:"title"`
// 	Content string `db:"content"`
// }

// Главная страница блога
func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featured_posts, err := getFeaturedPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		recently_posts, err := getMostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := indexPage{
			Title:         "Escape",
			FeaturedPosts: featured_posts,
			MostRecent:    recently_posts,
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"] // Получаем postID в виде строки из параметров урла

		postID, err := strconv.Atoi(postIDStr) // Конвертируем строку postID в число
		if err != nil {
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := getPostByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что пост не найден
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error1", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error2", 500)
			log.Println(err.Error())
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

// Страница просмотра постов
func getFeaturedPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			post_id,
			title,
			subtitle,
			image_url,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 1
	`

	var posts []featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func getMostRecentPosts(db *sqlx.DB) ([]recentPostData, error) {
	const query = `
		SELECT
			post_id,
			title,
			subtitle,
			image_url,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 0
	`

	var posts []recentPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func getPostByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			image_url,
			content
		FROM
			post
		WHERE
			post_id = ?
	`

	var postContent postData

	// Обязательно нужно передать в параметрах postID
	err := db.Get(&postContent, query, postID)
	if err != nil {
		return postData{}, err
	}

	return postContent, nil
}
