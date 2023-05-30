package comments

import "context"

type commentsUsecase struct {
	commentsRepository Repository
}

func CommentsUsecase(repository Repository) Usecase {
	return &commentsUsecase{
		commentsRepository: repository,
	}
}

func (usecase *commentsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.commentsRepository.GetAll(ctx)
}

func (usecase *commentsUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.commentsRepository.GetByID(ctx, id)
}

func (usecase *commentsUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
	return usecase.commentsRepository.GetByName(ctx, name)
}

func (usecase *commentsUsecase) Create(ctx context.Context, commentsDomain *Domain) (Domain, error) {
	return usecase.commentsRepository.Create(ctx, commentsDomain)
}

func (usecase *commentsUsecase) Update(ctx context.Context, commentsDomain *Domain, id string) (Domain, error) {
	return usecase.commentsRepository.Update(ctx, commentsDomain, id)
}

func (usecase *commentsUsecase) Delete(ctx context.Context, id string) error {
	return usecase.commentsRepository.Delete(ctx, id)
}

func (usecase *commentsUsecase) Restore(ctx context.Context, id string) (Domain, error) {
	return usecase.commentsRepository.Restore(ctx, id)
}

func (usecase *commentsUsecase) ForceDelete(ctx context.Context, id string) error {
	return usecase.commentsRepository.ForceDelete(ctx, id)
}
