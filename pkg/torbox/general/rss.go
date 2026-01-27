package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetRSSFeeds() ([]models.RSSFeed, error) {
	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_RSS_GET_FEEDS)
	if err != nil {
		return nil, err
	}

	var resp models.RSSFeedsResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get RSS feeds: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) GetRSSFeedItems(feedId *int64) ([]models.RSSFeedItem, error) {
	params := make(url.Values)
	if feedId != nil {
		params.Add("feed_id", fmt.Sprintf("%d", *feedId))
	}

	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_RSS_GET_ITEMS, request.WithURLParams(&params))
	if err != nil {
		return nil, err
	}

	var resp models.RSSFeedItemsResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get RSS feed items: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) AddRSSFeed(url string, name string) error {
	r := models.AddRSSFeedRequest{
		URL:  url,
		Name: name,
	}

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_RSS_ADD_FEED, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to add RSS feed: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) ModifyRSSFeed(feedId int64, name *string, active *bool) error {
	r := models.ModifyRSSFeedRequest{
		FeedID: feedId,
		Name:   *name,
		Active: active,
	}

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_RSS_MODIFY_FEED, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to modify RSS feed: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) ControlRSSFeed(feedId int64, operation string) error {
	r := models.ControlRSSFeedRequest{
		FeedID:    feedId,
		Operation: operation,
	}

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_RSS_CONTROL, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to control RSS feed: %s", resp.Detail)
	}

	return nil
}
