package dao

import (
	"fmt"
	"gin_demo/model"
)

func (d *Dao) FindAllUsers() (users []*model.User, err error) {
	fmt.Println("FindAllUsers dao", d.db)
	users = []*model.User{}
	rows, err := d.db.Query("SELECT first_name, last_name, points FROM users")
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	for rows.Next() {
		user := &model.User{}
		if err = rows.Scan(&user.FirstName, &user.LastName, &user.Points); err != nil {
			fmt.Println("Err", err.Error())
			return
		}
		fmt.Printf("User: %v\n", user)
		users = append(users, user)
	}
	return
}
