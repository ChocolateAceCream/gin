package dao

import (
	"database/sql"
	"fmt"
	"gin_demo/model"
)

func (d *Dao) FindAllUsers() (users []*model.User, err error) {
	fmt.Println("FindAllUsers dao", d.db)
	users = []*model.User{}
	var rows *sql.Rows
	if rows, err = d.db.Query("SELECT first_name, last_name, points FROM users"); err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := &model.User{}
		if err = rows.Scan(&user.FirstName, &user.LastName, &user.Points); err != nil {
			fmt.Println("Err", err.Error())
			return
		}
		users = append(users, user)
	}
	return
}

func (d *Dao) AddUser(user *model.User) (id int64, err error) {
	r, err := d.db.Exec("insert into users(first_name, last_name, points)values(?, ?, ?)", user.FirstName, user.LastName, user.Points)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ, new id:", id)
	return

}

func (d *Dao) EditUser(user *model.User) (count int64, err error) {
	r, err := d.db.Exec("update users set points=? where id=?", user.Points, user.ID)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	count, err = r.RowsAffected()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("update succ:", count)
	return

}

func (d *Dao) DeleteUser(user *model.User) (count int64, err error) {
	r, err := d.db.Exec("delete from users where id=?", user.ID)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	count, err = r.RowsAffected()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("update succ:", count)
	return

}
