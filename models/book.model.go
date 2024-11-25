package models

import (
	"archcalculator.github.io/db"
	"github.com/go-playground/validator"
	"net/http"
)

type Book struct {
	Id              int    `json:"id"`
	Title   string `json:"title"`
	Synopsis string `json:"synopsis"`
	Cover_Image  string `json:"cover_image"`
	Author  string `json:"author"`
	Publish_Date             string `json:"publish_date"`
	Member_Id           string `json:"member_id"`
}

func CreateBook(title string, synopsis string, cover_image string, author string, publish_date string, member_id string) (Response, error) {
	var res Response

	v := validator.New()

	book := Book{
		Title:   title,
		Synopsis: synopsis,
		Cover_Image:  cover_image,
		Author:  author,
		Publish_Date: publish_date,
		Member_Id: member_id,
	}

	err := v.Struct(book)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO book(title, synopsis, cover_image, author, publish_date, member_id) VALUES (?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(title, synopsis, cover_image, author, publish_date, member_id)

	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedID,
	}
	return res, nil

}

func DeleteBook(id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "DELETE FROM book WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil

}

func EditBook(id string, title string, synopsis string, cover_image string, author string, publish_date string, member_id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE book SET title=?, synopsis=?, cover_image=?, author=?, publish_date=?, member_id=? WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, synopsis, cover_image, author, publish_date, member_id, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil

}

func ReadAllBook()(Response, error){
	var obj Book
	var arrObj []Book
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM book"
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Title, &obj.Synopsis, &obj.Cover_Image, &obj.Author, &obj.Publish_Date, &obj.Member_Id)

		if err != nil{
			return res,err
		}
		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message="Success"
	res.Data = arrObj

	return res, nil
}

func ReadBookByBookId(id string)(Response, error){
	var obj Book
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from book where id= "+id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Title, &obj.Synopsis, &obj.Cover_Image, &obj.Author, &obj.Publish_Date, &obj.Member_Id)

		if err != nil{
			return res,err
		}
	}
	res.Status = http.StatusOK
	res.Message="Success"
	res.Data = obj

	return res, nil
}

func ReadBookByMemberId(id string)(Response, error){
	var obj Book
	var arrObj []Book
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from book where member_id= "+id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Title, &obj.Synopsis, &obj.Cover_Image, &obj.Author, &obj.Publish_Date, &obj.Member_Id)

		if err != nil{
			return res,err
		}
		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message="Success"
	res.Data = arrObj

	return res, nil
}

