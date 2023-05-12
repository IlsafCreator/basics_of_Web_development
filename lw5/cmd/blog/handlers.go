package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
	MostRecent    []recentPostData
}
type postPage struct {
	Title    string
	Subtitle string
}

type featuredPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}
type recentPostData struct {
	Title       string
	Subtitle    string
	PostImage   string
	Author      string
	AuthorImg   string
	PublishDate string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Escape",
		FeaturedPosts: featuredPosts(),
		MostRecent:    mostRecent(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/the-road-ahead.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := postPage{
		Title:    "The Road Ahead",
		Subtitle: "The road ahead might be paved - it might not be.",
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "featured-posts__post_the-road-ahead",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.png",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you've never been before.",
			ImgModifier: "featured-posts__post_from-top-down",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.png",
			PublishDate: "September 25, 2015",
		},
	}
}

func mostRecent() []recentPostData {
	return []recentPostData{
		{
			PostImage:   "static/img/balloons.png",
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage:   "static/img/golden_gate.png",
			Title:       "Sunny Side Upl",
			Subtitle:    "No place is ever as bad as they tell you it's going to be.",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage:   "static/img/lake.png",
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage:   "static/img/sea_waves.png",
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			Author:      "William Wong",
			AuthorImg:   "static/img/william_wong.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage:   "static/img/fog.png",
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage:   "static/img/waterfall.png",
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
	}
}
