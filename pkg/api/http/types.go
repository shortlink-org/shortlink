package http

type addRequest struct {
	Url      string
	Describe string
}

type getRequest struct {
	Hash     string
	Describe string
}

type deleteRequest struct {
	Hash string
}
