package client

import (
	"fmt"

	"github.com/supernova0730/rickandmortysync/pkg/rest"
)

type Client struct {
	rest *rest.Client
	url  string
}

func New(rest *rest.Client) *Client {
	return &Client{
		rest: rest,
		url:  "https://rickandmortyapi.com/api/character/",
	}
}

func (p *Client) GetInfo() (info Info, err error) {
	var result Response
	err = p.rest.Get(p.url, &result)
	if err != nil {
		err = fmt.Errorf("failed to get info: %v", err)
		return
	}

	info = result.Info
	return
}

func (p *Client) GetCharactersByPage(page int) (characters []Character, err error) {
	var response Response
	url := fmt.Sprintf(p.url+"?page=%d", page)
	err = p.rest.Get(url, &response)
	if err != nil {
		err = fmt.Errorf("failed to get characters: %v", err)
		return
	}

	characters = response.Data
	return
}
