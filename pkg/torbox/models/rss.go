package models

// RSSFeed represents an RSS feed subscription
type RSSFeed struct {
	ID        int64  `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// RSSFeedItem represents an item from an RSS feed
type RSSFeedItem struct {
	ID          int64  `json:"id"`
	FeedID      int64  `json:"feed_id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	PubDate     string `json:"pub_date"`
	Downloaded  bool   `json:"downloaded"`
}

// AddRSSFeedRequest represents a request to add an RSS feed
type AddRSSFeedRequest struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// ModifyRSSFeedRequest represents a request to modify an RSS feed
type ModifyRSSFeedRequest struct {
	FeedID int64  `json:"feed_id"`
	Name   string `json:"name,omitempty"`
	Active *bool  `json:"active,omitempty"`
}

// ControlRSSFeedRequest represents a request to control an RSS feed
type ControlRSSFeedRequest struct {
	FeedID    int64  `json:"feed_id"`
	Operation string `json:"operation"` // "delete", "refresh", etc.
}

// RSSFeedsResponse represents the response from the RSS feeds list endpoint
type RSSFeedsResponse struct {
	BaseResponse
	Data []RSSFeed `json:"data"`
}

// RSSFeedItemsResponse represents the response from the RSS feed items endpoint
type RSSFeedItemsResponse struct {
	BaseResponse
	Data []RSSFeedItem `json:"data"`
}
