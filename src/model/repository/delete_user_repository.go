package repository

import "github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"

func (ur *userRepository) DeleteUserRepository(id int) *rest_err.RestErr {
	result, err := ur.database.Exec(sqlDeleteUser, id)
	if err != nil {
		//	todo:logs
		return rest_err.NewInternalServerError(err.Error())
	}

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		//	todo:logs
		return rest_err.NewInternalServerError(err.Error())
	}
	println(rowsaffected)

	return nil
}
