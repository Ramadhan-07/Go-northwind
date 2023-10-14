package services

import "codeid.northwind/repositories"

type Servicemanager struct {
	CategoryService
	//ProductService
}

// constructor
func NewServiceManager(repoMgr *repositories.RepositoryManager) *Servicemanager{
	return &Servicemanager{
		CategoryService: *NewCategoryService(repoMgr),
		//ProductService: *NewProductService(&repoMgr.ProductCategory),
}
}