package follows

import "context"

type followsUsecase struct {
	followsRepository Repository
}

func FollowsUsecase(repository Repository) Usecase {
	return &followsUsecase{
		followsRepository: repository,
	}
}

func (usecase *followsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.followsRepository.GetAll(ctx)
}

func (usecase *followsUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.followsRepository.GetByID(ctx, id)
}

func (usecase *followsUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
	return usecase.followsRepository.GetByName(ctx, name)
}

func (usecase *followsUsecase) Create(ctx context.Context, followsDomain *Domain) (Domain, error) {
	return usecase.followsRepository.Create(ctx, followsDomain)
}

func (usecase *followsUsecase) Update(ctx context.Context, followsDomain *Domain, id string) (Domain, error) {
	return usecase.followsRepository.Update(ctx, followsDomain, id)
}

func (usecase *followsUsecase) Delete(ctx context.Context, id string) error {
	return usecase.followsRepository.Delete(ctx, id)
}

func (usecase *followsUsecase) Restore(ctx context.Context, id string) (Domain, error) {
	return usecase.followsRepository.Restore(ctx, id)
}

func (usecase *followsUsecase) ForceDelete(ctx context.Context, id string) error {
	return usecase.followsRepository.ForceDelete(ctx, id)
}
