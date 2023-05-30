package savedthreads

import "context"

type savedthreadsUsecase struct {
	savedthreadsRepository Repository
}

func NewSavedsavedthreadsUsecase(repository Repository) Usecase {
	return &savedthreadsUsecase{
		savedthreadsRepository: repository,
	}
}

func (usecase *savedthreadsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.savedthreadsRepository.GetAll(ctx)
}

func (usecase *savedthreadsUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.savedthreadsRepository.GetByID(ctx, id)
}

func (usecase *savedthreadsUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
	return usecase.savedthreadsRepository.GetByName(ctx, name)
}

func (usecase *savedthreadsUsecase) Create(ctx context.Context, savedthreadsDomain *Domain) (Domain, error) {
	return usecase.savedthreadsRepository.Create(ctx, savedthreadsDomain)
}

func (usecase *savedthreadsUsecase) Update(ctx context.Context, savedthreadsDomain *Domain, id string) (Domain, error) {
	return usecase.savedthreadsRepository.Update(ctx, savedthreadsDomain, id)
}

func (usecase *savedthreadsUsecase) Delete(ctx context.Context, id string) error {
	return usecase.savedthreadsRepository.Delete(ctx, id)
}

func (usecase *savedthreadsUsecase) Restore(ctx context.Context, id string) (Domain, error) {
	return usecase.savedthreadsRepository.Restore(ctx, id)
}

func (usecase *savedthreadsUsecase) ForceDelete(ctx context.Context, id string) error {
	return usecase.savedthreadsRepository.ForceDelete(ctx, id)
}
