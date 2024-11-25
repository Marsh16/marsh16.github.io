package models

import (
	"archcalculator.github.io/db"
	"net/http"

	"github.com/go-playground/validator"
)

type Member struct {
	Id              int    `json:"id"`
	Name   string `json:"name"`
	Phone_Number   string `json:"phone_number"`
	Email   string `json:"email"`
	Birthday   string `json:"birthday"`
	No_Ktp   string `json:"no_ktp"`
}

func CreateMember(name string, phone_number string, email string, birthday string, no_ktp string) (Response, error) {
	var res Response

	v := validator.New()

	member := Member{
		Name :   name,
		Phone_Number: phone_number,
		Email: email,
		Birthday: birthday,
		No_Ktp: no_ktp,
	}

	err := v.Struct(member)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO member(name, phone_number, email, birthday, no_ktp) VALUES (?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(name, phone_number, email, birthday, no_ktp)

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

func DeleteMember(id string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "DELETE FROM member WHERE id=?"
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

func EditMember(id string, name string, phone_number string, email string, birthday string, no_ktp string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE member SET name=?,phone_number=?,email=?, birthday=?,no_ktp=? WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, phone_number, email, birthday, no_ktp, id)

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

func ReadAllMember()(Response, error){
	var obj Member
	var arrObj []Member
	var res Response

	con:= db.CreateCon()

	sqlStatement := "SELECT * FROM member"
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil{
		return res,err
	}
 
	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name,&obj.Phone_Number,&obj.Email, &obj.Birthday,&obj.No_Ktp)

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