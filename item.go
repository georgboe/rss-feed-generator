package podcast

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/georgboe/rss-feed-generator/html2text"
)

// Item represents a single entry in a podcast.
//
// Article minimal requirements are:
// - Title
// - Description
// - Link
//
// Audio minimal requirements are:
// - Title
// - Description
// - Enclosure (HREF, Type and Length all required)
//
// Recommendations:
// - Setting the minimal fields sets most of other fields, including iTunes.
// - Use the Published time.Time setting instead of PubDate.
// - Always set an Enclosure.Length, to be nice to your downloaders.
// - Use Enclosure.Type instead of setting TypeFormatted for valid extensions.
type Item struct {
	XMLName            xml.Name `xml:"item"`
	GUID               *GUID
	Title              string `xml:"title"`
	Link               string `xml:"link"`
	Description        *Description
	EncodedDescription *EncodedContent
	AuthorFormatted    string `xml:"author,omitempty"`
	Category           string `xml:"category,omitempty"`
	Comments           string `xml:"comments,omitempty"`
	Source             string `xml:"source,omitempty"`
	PubDate            string `xml:"pubDate,omitempty"`
	Enclosure          *Enclosure

	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	IAuthor            string `xml:"itunes:author,omitempty"`
	ITitle             string `xml:"itunes:title,omitempty"`
	SeasonNumber       string `xml:"itunes:season,omitempty"`
	EpisodeNumber      string `xml:"itunes:episode,omitempty"`
	EpisodeType        string `xml:"itunes:episodeType,omitempty"`
	ISubtitle          string `xml:"itunes:subtitle,omitempty"`
	ISummary           *ISummary
	IImage             *IImage
	IBlock             string `xml:"itunes:block,omitempty"`
	IDuration          string `xml:"itunes:duration,omitempty"`
	IExplicit          string `xml:"itunes:explicit,omitempty"`
	IIsClosedCaptioned string `xml:"itunes:isClosedCaptioned,omitempty"`
	IOrder             string `xml:"itunes:order,omitempty"`
}

func (i *Item) AddGUID(guid string) {
	if len(guid) <= 0 {
		return
	}

	i.GUID = &GUID{
		IsPermaLink: false,
		Value:       guid,
	}
}

func (i *Item) AddTitle(title string) {
	if len(title) <= 0 {
		return
	}

	i.Title = title
}

func (i *Item) AddLink(link string) {
	if len(link) <= 0 {
		return
	}

	i.Link = link
}

func (i *Item) AddDescription(description Description) {
	if len(description.Text) <= 0 {
		return
	}

	i.Description = &description
	i.EncodedDescription = &EncodedContent{
		Text: description.Text,
	}
}

// AddEnclosure adds the downloadable asset to the podcast Item.
func (i *Item) AddEnclosure(
	url string, enclosureType EnclosureType, enclosureTypeString string, lengthInBytes int64) {
	i.Enclosure = &Enclosure{
		URL:           url,
		Type:          enclosureType,
		TypeFormatted: enclosureTypeString,
		Length:        lengthInBytes,
	}
}

func (i *Item) AddEpisodeNumber(episodeNumber int64) {
	if episodeNumber <= 0 {
		return
	}

	i.EpisodeNumber = strconv.FormatInt(episodeNumber, 10)
}

func (i *Item) AddEpisodeType(episodeType string) {
	if len(episodeType) == 0 {
		return
	} else if episodeType == EpisodeTypeFull || episodeType == EpisodeTypeTrailer || episodeType == EpisodeTypeBonus {
		i.EpisodeType = episodeType
	}

	return
}

// AddImage adds the image as an iTunes-only IImage.  RSS 2.0 does not have
// the specification of Images at the Item level.
//
// Podcast feeds contain artwork that is a minimum size of
// 1400 x 1400 pixels and a maximum size of 3000 x 3000 pixels,
// 72 dpi, in JPEG or PNG format with appropriate file
// extensions (.jpg, .png), and in the RGB colorspace. To optimize
// images for mobile devices, Apple recommends compressing your
// image files.
func (i *Item) AddImage(url string) {
	if len(url) > 0 {
		i.IImage = &IImage{HREF: url}
	}
}

func (i *Item) AddItunesBlock(block string) {
	if block == "hide" {
		i.IBlock = "Yes"
	} else {
		i.IBlock = "No"
	}
}

func (i *Item) AddItunesTitle(title string) {
	if len(title) == 0 {
		return
	}

	i.ITitle = title
}

func (i *Item) AddParentalAdvisory(parentalAdvisory string) {
	if parentalAdvisory == ParentalAdvisoryExplicit {
		i.IExplicit = "yes"
	} else if parentalAdvisory == ParentalAdvisoryClean {
		i.IExplicit = "no"
	}

	return
}

func (i *Item) AddPubDate(datetime string) {

	if len(datetime) == 0 {
		return
	}

	i.PubDate = datetime
}

func (i *Item) AddSeasonNumber(seasonNumber int64) {
	if seasonNumber <= 0 {
		return
	}

	i.SeasonNumber = strconv.FormatInt(seasonNumber, 10)
}

// AddSummary adds the iTunes summary.
//
// Limit: 4000 characters
//
// Note that this field is a CDATA encoded field which allows for rich text
// such as html links: `<a href="http://www.apple.com">Apple</a>`.
func (i *Item) AddSummary(summary string) {
	count := utf8.RuneCountInString(summary)
	if count > 4000 {
		s := []rune(summary)
		summary = string(s[0:4000])
	}
	i.ISummary = &ISummary{
		Text: html2text.HTML2Text(summary),
	}
}

// AddDuration adds the duration to the iTunes duration field.
func (i *Item) AddDuration(durationInSeconds int64) {
	if durationInSeconds <= 0 {
		return
	}
	i.IDuration = fmt.Sprint(durationInSeconds)
}

var parseDuration = func(duration int64) string {
	h := duration / 3600
	duration = duration % 3600

	m := duration / 60
	duration = duration % 60

	s := duration

	// HH:MM:SS
	if h > 9 {
		return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	}

	// H:MM:SS
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}

	// MM:SS
	if m > 9 {
		return fmt.Sprintf("%02d:%02d", m, s)
	}

	// M:SS
	return fmt.Sprintf("%d:%02d", m, s)
}
