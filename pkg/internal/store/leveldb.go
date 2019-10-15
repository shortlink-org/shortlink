package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/batazor/shortlink/pkg/internal/link"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBLinkList struct {
	client *leveldb.DB
}

func (l *LevelDBLinkList) Init() error {
	var err error
	l.client, err = leveldb.OpenFile("/tmp/links.db", nil)
	if err != nil {
		return err
	}

	return nil
}

func (l *LevelDBLinkList) Get(id string) (*link.Link, error) {
	value, err := l.client.Get([]byte(id), nil)
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	var response link.Link

	err = json.Unmarshal(value, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	return &response, nil
}

func (l *LevelDBLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = l.client.Put([]byte(data.Hash), payload, nil)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (l *LevelDBLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

func (l *LevelDBLinkList) Delete(id string) error {
	err := l.client.Delete([]byte(id), nil)
	return err
}
