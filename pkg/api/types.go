package api

type addRequest struct {
	Url      string
	Describe string
}

type getRequest struct {
	Url      string
	Describe string
}

type deleteRequest struct {
	Url string
}
