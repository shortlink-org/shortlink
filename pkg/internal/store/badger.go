package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/batazor/shortlink/pkg/internal/link"
	"github.com/dgraph-io/badger"
)

type BadgerLinkList struct {
	client *badger.DB
}

func (b *BadgerLinkList) Init() error {
	var err error
	b.client, err = badger.Open(badger.DefaultOptions("/tmp/links.badger"))
	if err != nil {
		return err
	}

	return nil
}

func (b *BadgerLinkList) Get(id string) (*link.Link, error) {
	var valCopy []byte

	err := b.client.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			return nil
		})
		if err != nil {
			return err
		}

		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	var response link.Link

	err = json.Unmarshal(valCopy, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: errors.New(fmt.Sprintf("Not found id: %s", id))}
	}

	return &response, nil
}

func (b *BadgerLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = b.client.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(data.Hash), payload)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *BadgerLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

func (b *BadgerLinkList) Delete(id string) error {
	err := b.client.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(id))
		return err
	})
	return err
}
