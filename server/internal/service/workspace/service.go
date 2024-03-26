package workspace

type Repository interface {
	Create(w string, userId int) (int, error)
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
