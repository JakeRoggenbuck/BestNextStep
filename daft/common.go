package main

import (
	"database/sql"
	"github.com/jakeroggenbuck/BestNextStep/daft/col"
	"github.com/jakeroggenbuck/BestNextStep/daft/step"
	"log"
)

func createDefaultElements(db *sql.DB) {
	stepRepository := step.NewSQLiteRepository(db)

	if err := stepRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	stepOne := step.Step{
		Name:       "Step One",
		Desc:       "The first step.",
		Collection: 1,
		Owner:      1,
	}

	stepTwo := step.Step{
		Name:       "Step Two",
		Desc:       "The second step.",
		Collection: 1,
		Owner:      1,
	}

	stepRepository.Create(stepOne)
	stepRepository.Create(stepTwo)

	colRepository := col.NewSQLiteRepository(db)

	if err := stepRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	colOne := col.Col{
		Name:  "First Task",
		Desc:  "The first task that I have to do.",
		Owner: 1,
	}

	colRepository.Create(colOne)
}
