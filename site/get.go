package site

import "context"

type Site struct {
	// ID is a unique ID for the site.
	ID int `json:"id"`
	// URL is the site's URL.
	URL string `json:"url"`
}

//encore:api public method=GET path=/site/:siteID
func (s *Service) Get(ctx context.Context, siteID int) (*Site, error) {
	var site Site
	if err := s.db.Where("id = $1", siteID).First(&site).Error; err != nil {
		return nil, err
	}
	return &site, nil
}
