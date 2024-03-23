package model

type Workspace struct {
	Id   string
	Name string
}

type WorkspaceWithoutId struct {
	Name string
}
