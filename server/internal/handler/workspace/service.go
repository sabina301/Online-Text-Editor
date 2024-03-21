package workspace

type Service interface {
	Create()
	Get()
	AddUser()
	Delete()
}

type Implementation struct {
	workspaceService Service
}

func NewImplementation(workspaceService Service) *Implementation {
	return &Implementation{workspaceService: workspaceService}
}
