package users

import (
	"main/datasource/mysql/users_db"
	"main/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryGetUserByEmail = "SELECT id, first_name, last_name, email, password FROM users WHERE email=?;"
	queryGetUserByID    = "SELECT id, first_name, last_name, email FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerErrror("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerErrror("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerErrror("database error")
	}
	user.ID = userID
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.NewInternalServerErrror("invalid email")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerErrror("database error")
	}
	return nil
}

func (user *User) GetByID() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewInternalServerErrror("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerErrror("database error")
	}
	return nil
}
