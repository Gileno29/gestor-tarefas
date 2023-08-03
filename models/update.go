package models

import "github.com/Gileno29/gestor-tarefas/db"

func Update(id int64, todo Todo) (int64, error) {

	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err

	}
	defer conn.Close()

	res, err := conn.Exec(`ÙPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)

	if err != nil {
		return 0, err

	}

	return res.RowsAffected()

}