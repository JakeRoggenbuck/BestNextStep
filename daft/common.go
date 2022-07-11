package main

import (
	"database/sql"
	"fmt"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"log"
)

func createDefaultElements(db *sql.DB) {
	stepRepository := step.NewSQLiteRepository(db)

	if err := stepRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	stepOne := step.Step{
		Name:  "Step One",
		Desc:  "The first step.",
		Left:  -1,
		Right: 2,
		Owner: 1,
	}
	stepTwo := step.Step{
		Name:  "Step Two",
		Desc:  "The second step.",
		Left:  1,
		Right: -1,
		Owner: 1,
	}

	createdStepOne, err := stepRepository.Create(stepOne)
	if err != nil {
		fmt.Println(err)
	}

	createdStepTwo, err := stepRepository.Create(stepTwo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(createdStepOne)
	fmt.Println(createdStepTwo)
}
