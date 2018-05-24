package response

type TEXTResponse struct {
	*Response
}

func (r *TEXTResponse) Xpath() interface{} {
	return nil
}
