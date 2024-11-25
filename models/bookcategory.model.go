package models

import (
	"archcalculator.github.io/db"
	"github.com/go-playground/validator"
	"net/http"
)

type BookCategory struct {
	Id              int    `json:"id"`
	Book_Id   string `json:"book_id"`
	Category_Id string `json:"category_id"`
}

func ReadBookByCategoryId(id string)(Response, error){
	var obj BookCategory
	var arrObj []BookCategory
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from book_category where category_id= "+id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Book_Id, &obj.Category_Id)

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

func ReadCategoryByBookId(id string)(Response, error){
	var obj BookCategory
	var arrObj []BookCategory
	var res Response
	con:= db.CreateCon()

	sqlStatement := "SELECT * from book_category where book_id= "+id
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Book_Id, &obj.Category_Id)

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

func CreateBookCategory(book_id string, category_id string) (Response, error) {
	var res Response

	v := validator.New()

	book_category := BookCategory{
		Book_Id :   book_id,
		Category_Id :   category_id,
	}

	err := v.Struct(book_category)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO book_category(book_id,category_id) VALUES (?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(book_id,category_id)

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

func DeleteBookCategory(id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "DELETE FROM book_category WHERE id=?"
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