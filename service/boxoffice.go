package service

type BoxOfficeService interface {
	PullData() interface{} // todo 拉取票房数据
}

type boxOfficeService struct {
}

func NewBoxOfficeService() BoxOfficeService {
	return &boxOfficeService{}
}

func (srv *boxOfficeService) PullData() interface{} {

	return ""
}
