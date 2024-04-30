package repository

import "github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"

func (lr *listRepository) DeleteList(id int) *rest_err.RestErr {
	_, err := lr.DB.Exec(deleteListByID, id)
	if err != nil {
		return rest_err.NewInternalServerError("unable to delete the list")
	}

	return nil
}
