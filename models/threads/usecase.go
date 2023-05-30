package threads

import "context"

type threadsUsecase struct {
	threadsRepository Repository
}

func NewThreadsUsecase(repository Repository) Usecase {
	return &threadsUsecase{
		threadsRepository: repository,
	}
}

func (usecase *threadsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.threadsRepository.GetAll(ctx)
}

func (usecase *threadsUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.threadsRepository.GetByID(ctx, id)
}

func (usecase *threadsUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
	return usecase.threadsRepository.GetByName(ctx, name)
}

func (usecase *threadsUsecase) Create(ctx context.Context, threadsDomain *Domain) (Domain, error) {
	return usecase.threadsRepository.Create(ctx, threadsDomain)
}

func (usecase *threadsUsecase) Update(ctx context.Context, threadsDomain *Domain, id string) (Domain, error) {
	return usecase.threadsRepository.Update(ctx, threadsDomain, id)
}

func (usecase *threadsUsecase) Delete(ctx context.Context, id string) error {
	return usecase.threadsRepository.Delete(ctx, id)
}

func (usecase *threadsUsecase) Restore(ctx context.Context, id string) (Domain, error) {
	return usecase.threadsRepository.Restore(ctx, id)
}

func (usecase *threadsUsecase) ForceDelete(ctx context.Context, id string) error {
	return usecase.threadsRepository.ForceDelete(ctx, id)
}
