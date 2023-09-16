package surge

import "github.com/zapsaang/conf_factory/repo"

type Repo struct {
	repo.Repo
	BaseDir   string
	ResultDir string
}

func New() *Repo {
	return &Repo{
		Repo:      repo.New(repo.RoleIntegrator),
		BaseDir:   "/base/surge",
		ResultDir: "/result/surge",
	}
}

func GetSurgeBaseList() []string {
	return []string{}
}

func GetSurgeBaseContent(baseName string) []byte {
	return []byte{}
}
