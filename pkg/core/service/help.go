package common

type Helper struct {
	Name string
}

func NewHelper() *Helper {
	return &Helper{
		Name: "Some name",
	}
}