package model

import (
	"fmt" 

	"github.com/go-pg/pg"
)

func ConnectToDb(user, password, database, addr  string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr: addr,
	})
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	if n == 1 {
		fmt.Println("Connected")
	}
	return db
}

func Getbyid(db *pg.DB, id int16) (data Post) {
	data = Post{Id: id}
	err := db.Select(&data)
	if err != nil {
		panic(err)
	}
	return data
}
func Getall(db *pg.DB) (data []Post) {
	_ = db.Model(&data).Select()
	return data
}
func Getpublished(db *pg.DB) (data []Post) {
	_ = db.Model(&data).Where("published LIKE ?", "publish%").Select()
	return data
}
func Edit(db *pg.DB, idpost int16) (data Post) {
	data = Post{Id: idpost}
	_ = db.Select(&data)
	return data
}
func Insertdata(db *pg.DB, title, alias, intro_text, full_text, image string, published string, category []string, created_at string) (data Post) {
	fmt.Println(category)
	data = Post{Title: title, Alias: alias, Intro_text: intro_text,
		Full_text: full_text, Image: image, Published: published, Categories: category,  Created_at: created_at}
	err := db.Insert(&data)
	if err != nil {
		panic(err)
	}
	return data
}
func Deletedata(db *pg.DB, id int16){
	data := Post{Id: id}
	err := db.Delete(&data)
	if err != nil {
		panic(err)
	}
}
func Update(db *pg.DB, title, alias, intro_text, full_text, image string, published string, category []string, created_at string, id int16){
	var post Post
	_, _ = db.Model(&post).
	Set("title = ?, alias = ?, intro_text = ?, full_text = ? , image = ?,  published = ?, categories = ?", title, alias, intro_text, full_text, image,  published, category).
	Where("id = ?", id).Update()
}
func Updatename(db *pg.DB, name string, id int16){
	var post Post
	_, _ = db.Model(&post).
	Set("title = ?", name).
	Where("id = ?", id).Update()
}

func Getbycategory(db *pg.DB, cate string) (data []Post) {
	_ = db.Model(&data).Where("categories ? '" +  cate + "'").Select()
	fmt.Println(data)
	return data
}