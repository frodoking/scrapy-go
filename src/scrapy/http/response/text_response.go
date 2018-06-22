package response

type TEXTResponse struct {
	*BaseResponse
}

func (r *TEXTResponse) Xpath() interface{} {
	return nil
}
