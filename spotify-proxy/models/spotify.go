package models

// SearchResult contains spotify search results.
type SearchResult struct {
	Tracks []Track
}

// Track contains basic info about a track.
type Track struct {
	// The name of the track.
	Name string `json:"name" example:"Seventeen"`
	// The SpotifyID for the track.
	ID ID `json:"id" example:"3ctBlLC81pc2TMriIxnXxN"`
	// The Spotify URI for the track.
	URI URI `json:"uri" example:"spotify:track:3ctBlLC81pc2TMriIxnXxN"`
	// The album on which the track appears. The album object includes a link in href to full information about the album.
	Album Album `json:"album"`
	// A slice of the track album Artists
	Artists []Artist `json:"artists"`
	// A list of the countries in which the track can be played,
	// identified by their ISO 3166-1 alpha-2 codes. i.e. ["AD","AE","AR","AT"]
	AvailableMarkets []string `json:"availableMarkets"`
	// The length of the track, in milliseconds.
	Duration int `json:"duration" example:"271125"`
	// A link to the Web API endpoint providing full details for this track.
	// Popularity of the track.  The value will be between 0 and 100,
	// with 100 being the most popular.  The popularity is calculated from
	// both total plays and most recent plays.
	Popularity int `json:"popularity" example:"50"`
	// A link to the Web API enpoint providing full
	// details of the track.
	Endpoint string `json:"endpoint" example:"https://api.spotify.com/v1/tracks/3ctBlLC81pc2TMriIxnXxN"`
	// External URLs for this track. i.e. "spotify": "https://open.spotify.com/track/3ctBlLC81pc2TMriIxnXxN"
	ExternalURLs map[string]string `json:"externalUrls"`
	// A URL to a 30 second preview (MP3) of the track.
	PreviewURL string `json:"previewUrl" example:"https://p.scdn.co/mp3-preview/fcf7c5d0ba73fd4c3648f5ed7300ecb07022a3fc?cid=783fd35ed6094d45888a78e9404f42ed"`
}

// Artist contains basic info about an artist.
type Artist struct {
	// The name of the artist.
	Name string `json:"name" example:"Tomberlin"`
	// The SpotifyID for the artist.
	ID ID `json:"id" example:"0jzaoSt5gOC04OWBqN78VS"`
	// The Spotify URI for the artist.
	URI URI `json:"uri" example:"spotify:artist:0jzaoSt5gOC04OWBqN78VS"`
	// A link to the Web API enpoint providing full details of the artist.
	Endpoint string `json:"endpoint" example:"https://api.spotify.com/v1/artists/0jzaoSt5gOC04OWBqN78VS"`
	// External URLs for this artist. i.e. "spotify": "https://open.spotify.com/artist/0jzaoSt5gOC04OWBqN78VS"
	ExternalURLs map[string]string `json:"externalUrls"`
}

// Album contains basic info about an album.
type Album struct {
	// The name of the album.
	Name string `json:"name" example:"At Weddings"`
	// The SpotifyID for the album.
	ID ID `json:"id" example:"1v55LXxVAFvFV02xGMImwP"`
	// The SpotifyURI for the album.
	URI URI `json:"uri" example:"spotify:album:1v55LXxVAFvFV02xGMImwP"`
	// The type of the album: one of "album",
	// "single", or "compilation".
	AlbumType string `json:"albumType" example:"album"`
	// The cover art for the album in various sizes,
	// widest first.
	Images []Image `json:"images"`
	// A link to the Web API enpoint providing full
	// details of the album.
	Endpoint string `json:"endpoint" example:"https://api.spotify.com/v1/albums/1v55LXxVAFvFV02xGMImwP"`
	// Known external URLs for this album. i.e. "spotify": "https://open.spotify.com/album/1v55LXxVAFvFV02xGMImwP"
	ExternalURLs map[string]string `json:"externalUrls"`
}

// Image identifies an image associated with an item.
type Image struct {
	// The image height, in pixels.
	Height int `json:"height" example:"640"`
	// The image width, in pixels.
	Width int `json:"width" example:"640"`
	// The source URL of the image.
	URL string `json:"url" example:"https://i.scdn.co/image/941a28b9a99eab76991c4b9d48fc8887b63689a1"`
}

// URI identifies an artist, album, track, or category.  For example,
// spotify:track:6rqhFgbbKwnb9MLmUQDhG6
type URI string

// ID is a base-62 identifier for an artist, track, album, etc.
// It can be found at the end of a spotify.URI.
type ID string
