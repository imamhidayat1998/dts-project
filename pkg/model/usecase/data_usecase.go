package usecase

import (
	"dts-project/pkg/model"
	"dts-project/pkg/model/request"
)

type ServiceCenterUsecase struct {
	serviceCenterRepo model.DataRepository
}

func NewServiceCenter(ServiceDataRepository model.DataRepository) ServiceCenterUsecase {
	return ServiceCenterUsecase{
		serviceCenterRepo: ServiceDataRepository,
	}
}

func (r ServiceCenterUsecase)CreateUsers(params request.User) (string,error) {
	Result , err := r.serviceCenterRepo.CreateUser(params)
	if err != nil{
		return "",err
	}
	return Result,err

}
func (r ServiceCenterUsecase)UpdateUsers(params request.User)error  {
	err := r.serviceCenterRepo.UpdateUser(params)
	if err != nil {
		return err
	}
	return nil


}
