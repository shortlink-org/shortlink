package link

func NewURL(link string) (Link, error) {
	newLink := Link{Url: link}
	return newLink, nil
}

func Init() (*LinkList, error) {
	return &LinkList{
		links: make(map[string]Link),
	}, nil
}

func (l *LinkList) Get(link Link) (*Link, error) {
	l.mu.Lock()
	response := l.links[link.Url]
	l.mu.Unlock()

	return &response, nil
}

func (l *LinkList) Add(link Link) error {
	l.mu.Lock()
	l.links[link.Url] = link
	l.mu.Unlock()

	return nil
}

func (l *LinkList) Update(link Link) (*Link, error) {
	return nil, nil
}

func (l *LinkList) Delete(link Link) error {
	l.mu.Lock()
	delete(l.links, link.Url)
	l.mu.Unlock()

	return nil
}
