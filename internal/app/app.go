package app

import (
	"log"
	"university-db-admin/internal/config"
	"university-db-admin/internal/repository"
	"university-db-admin/internal/repository/postgres"
	"university-db-admin/internal/ui"
	"university-db-admin/pkg/dbclient"
)

type App struct {
	cfg        *config.Config
	repository *repository.Repository
}

func NewApp() App {
	log.Println("initializing application")

	log.Println("initializing config")
	cfg := config.LoadConfig()

	log.Println("initializing database")
	pg := dbclient.NewClientPG(cfg.DB)

	log.Println("initializing repositories")
	employees := postgres.NewEmployeesRepository(pg)
	groups := postgres.NewGroupsRepository(pg)
	lessons := postgres.NewLessonsRepository(pg)
	positions := postgres.NewPositionsRepository(pg)
	subjects := postgres.NewSubjectsRepository(pg)
	lessonTypes := postgres.NewLessonTypesRepository(pg)
	marks := postgres.NewMarksRepository(pg)
	students := postgres.NewStudentsRepository(pg)
	employeesSubjects := postgres.NewEmployeesSubjectsRepository(pg)

	log.Println("application initialized")

	return App{
		cfg: cfg,
		repository: &repository.Repository{
			Employees:         employees,
			Groups:            groups,
			Lessons:           lessons,
			Positions:         positions,
			Subjects:          subjects,
			LessonTypes:       lessonTypes,
			Marks:             marks,
			Students:          students,
			EmployeesSubjects: employeesSubjects,
		},
	}
}

func (a *App) startUI() {
	ui.Run(a.repository)
}

func Run() {
	app := NewApp()

	log.Println("application started")
	app.startUI()
}
