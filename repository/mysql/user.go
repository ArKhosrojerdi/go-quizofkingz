package mysql

import (
	"database/sql"
	"errors"
	"quizofkingz/entity"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/richerror"

	_ "github.com/go-sql-driver/mysql"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number=?`, phoneNumber)

	_, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}

		return false, richerror.New("mysql.IsPhoneNumberUnique").
			WithErr(err).
			WithMessage(errmsg.ErrorMsgMysqlQuery).
			WithKind(richerror.KindUnexpected)
	}

	return false, nil
}

func (d *MysqlDB) RegisterUser(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`INSERT INTO users(name, phone_number, password) VALUES(?, ?, ?)`, user.Name, user.PhoneNumber, user.Password)

	if err != nil {
		return entity.User{}, richerror.New("mysql.RegisterUser").
			WithErr(err).
			WithMessage(errmsg.ErrorMsgCantExecuteCommand).
			WithKind(richerror.KindUnexpected)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return entity.User{}, richerror.New("mysql.RegisterUser").
			WithErr(err).
			WithMessage(errmsg.ErrorMsgCantGetInsertedID).
			WithKind(richerror.KindUnexpected)
	}

	user = entity.User{
		ID:          uint(id),
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	}

	return user, nil
}

func (d *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number=?`, phoneNumber)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, richerror.New("mysql.GetUserByPhoneNumber").
			WithErr(err).
			WithMessage(errmsg.ErrorMsgMysqlQuery).
			WithKind(richerror.KindUnexpected)
	}

	return user, true, nil
}

func (d *MysqlDB) GetUserByID(userID uint) (entity.User, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE id=?`, userID)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New("mysql.GetUserByID").
				WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return entity.User{}, richerror.New("mysql.GetUserByID").
			WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	return user, err
}
