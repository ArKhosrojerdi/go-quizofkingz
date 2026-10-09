package mysql

import (
	"database/sql"
	"fmt"
	"quizofkingz/entity"

	_ "github.com/go-sql-driver/mysql"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number=?`, phoneNumber)

	_, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("mysql query error: %w", err)
	}

	return false, nil
}

func (d *MysqlDB) RegisterUser(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`INSERT INTO users(name, phone_number, password) VALUES(?, ?, ?)`, user.Name, user.PhoneNumber, user.Password)

	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return entity.User{}, fmt.Errorf("can't get inserted user ID: %w", err)
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
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("mysql query error: %w", err)
	}

	return user, true, nil
}

func (d *MysqlDB) GetUserByID(userID uint) (entity.User, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE id=?`, userID)

	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("record not found")
		}

		return entity.User{}, fmt.Errorf("can't scan query results: %w", err)
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	return user, err
}
