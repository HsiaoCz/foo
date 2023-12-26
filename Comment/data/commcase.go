package data

type CommUseCase struct{}

func NewCommUseCase() CommentRepo {
	return &CommUseCase{}
}
