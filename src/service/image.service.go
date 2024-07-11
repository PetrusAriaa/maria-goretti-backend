package service

type ImageListResponse struct {
	Data []string `json:"data"`
}

func (s *Service) GetImageList() *ImageListResponse {
	obj := s.repository.GetImageList()

	var res ImageListResponse
	res.Data = append(res.Data, obj...)
	return &res
}
