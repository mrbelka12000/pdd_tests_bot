package repo

type Repo struct {
	Case *CaseRepo
	User *UserRepo
}

func New() *Repo {
	return &Repo{
		Case: NewCaseRepo(),
		User: NewUserRepo(),
	}
}
