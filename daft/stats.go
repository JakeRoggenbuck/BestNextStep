package main

import (
	"fmt"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"github.com/jakeroggenbuck/BestNextStep/daft/user"
)

func userCount(repo *user.SQLiteRepository) int {
	all, err := repo.All()
	if err != nil {
		fmt.Println(err)
	}
	return len(all)
}

func stepCount(repo *step.SQLiteRepository) int {
	all, err := repo.All()
	if err != nil {
		fmt.Println(err)
	}
	return len(all)
}
