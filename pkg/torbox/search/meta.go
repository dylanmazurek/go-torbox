package search

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *SearchService) GetMeta(idType string, id string) (*models.Torrent, error) {
	params := &url.Values{}
	params.Add("metadata", "true")
	params.Add("check_cache", "true")
	params.Add("check_owned", "true")

	path := fmt.Sprintf("%s/%s:%s", constants.PATH_SEARCH_META, idType, id)
	req, err := s.newRequest(http.MethodGet, path, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetMetaResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
