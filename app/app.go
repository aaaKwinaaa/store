package app

type App struct {
	Controller SetupController
}

func New() *App {
	repository := &SetupRepository{}
	repository.Setup()

	service := &SetupService{Repo: repository}
	service.Setup()

	controller := &SetupController{Service: service}
	controller.Setup()

	app := App{
		Controller: *controller,
	}
	return &app
}
