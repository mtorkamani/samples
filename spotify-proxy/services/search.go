package services

import (
	"samples/spotify-proxy/models"

	"github.com/zmb3/spotify"
)

type SearchService interface {
	Search(ctx *Context, term string) (models.SearchResult, error)
}

type searchService struct {
	logger Logger
}

func NewSearchService(logger Logger) SearchService {
	return &searchService{
		logger: logger,
	}
}

func (svc *searchService) Search(ctx *Context, term string) (models.SearchResult, error) {
	svc.logger.Log("Start searching for term %s", term)
	searchRes, err := ctx.SpotifyContext.client.Search(term,
		spotify.SearchTypeAlbum|
			spotify.SearchTypeArtist|
			spotify.SearchTypeTrack)
	if err != nil {
		return models.SearchResult{}, err
	}
	tracks := []models.Track{}
	for _, trk := range searchRes.Tracks.Tracks {
		tracks = append(tracks, newTrack(trk))
	}

	return models.SearchResult{
		Tracks: tracks,
	}, nil
}

func newTrack(trk spotify.FullTrack) models.Track {
	return models.Track{
		ID:               models.ID(trk.ID),
		Name:             trk.Name,
		URI:              models.URI(trk.URI),
		Album:            newAlbum(trk.Album),
		Artists:          newArtists(trk.Artists),
		AvailableMarkets: trk.AvailableMarkets,
		Duration:         trk.Duration,
		Popularity:       trk.Popularity,
		Endpoint:         trk.Endpoint,
		ExternalURLs:     trk.ExternalURLs,
		PreviewURL:       trk.PreviewURL,
	}
}

func newArtists(arts []spotify.SimpleArtist) []models.Artist {
	results := []models.Artist{}
	for _, art := range arts {
		artist := models.Artist{
			ID:           models.ID(art.ID),
			Name:         art.Name,
			URI:          models.URI(art.URI),
			Endpoint:     art.Endpoint,
			ExternalURLs: art.ExternalURLs,
		}
		results = append(results, artist)
	}
	return results
}

func newAlbum(albm spotify.SimpleAlbum) models.Album {
	return models.Album{
		ID:           models.ID(albm.ID),
		Name:         albm.Name,
		URI:          models.URI(albm.URI),
		AlbumType:    albm.AlbumType,
		Images:       newImages(albm.Images),
		Endpoint:     albm.Endpoint,
		ExternalURLs: albm.ExternalURLs,
	}
}

func newImages(imgs []spotify.Image) []models.Image {
	result := []models.Image{}
	for _, img := range imgs {
		result = append(result, models.Image(img))
	}
	return result
}
