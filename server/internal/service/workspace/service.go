package workspace

type Repository interface {
	Create()
	Get()
	AddUser()
}

type workspaceService struct {
	workspaceRepository Repository
}

func NewWorkspaceService(repo Repository) *workspaceService {
	return &workspaceService{
		workspaceRepository: repo,
	}
}
