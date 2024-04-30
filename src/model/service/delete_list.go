package service

import "github.com/ThalesMonteir0/go-mvc-api-my-filmes/src/configuration/rest_err"

func (ls *listService) DeleteList(id int) *rest_err.RestErr {
	if _, err := ls.findListByID(id); err != nil {
		return err
	}

	if err := ls.repository.DeleteList(id); err != nil {
		return err
	}

	return nil
}
