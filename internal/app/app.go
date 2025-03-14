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
	pgClient := dbclient.NewClientPG(cfg.DB)

	log.Println("initializing repositories")
	empRepo := postgres.NewEmployeesRepository(pgClient)
	grpRepo := postgres.NewGroupsRepository(pgClient)
	lsnRepo := postgres.NewLessonsRepository(pgClient)
	posRepo := postgres.NewPositionsRepository(pgClient)
	sbjRepo := postgres.NewSubjectsRepository(pgClient)
	specRepo := postgres.NewSpecialRepository(pgClient)
	studRepo := postgres.NewStudentsRepository(pgClient)
	marksRepo := postgres.NewMarksRepository(pgClient)
	lsnTpsRepo := postgres.NewLessonTypesRepository(pgClient)
	empSbjRepo := postgres.NewEmployeesSubjectsRepository(pgClient)

	log.Println("application initialized")

	return App{
		cfg: cfg,
		repository: &repository.Repository{
			Employees:         empRepo,
			Groups:            grpRepo,
			LessonTypes:       lsnTpsRepo,
			Lessons:           lsnRepo,
			Marks:             marksRepo,
			Positions:         posRepo,
			Students:          studRepo,
			Subjects:          sbjRepo,
			EmployeesSubjects: empSbjRepo,
			Special:           specRepo,
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
