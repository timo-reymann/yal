package model

import "github.com/timo-reymann/yal/pkg/assets"

// Section represents a section consisting of title and items
type Section struct {
	// Entries contained by section
	Entries []Entry `json:"entries"`
	// Title gives a name to the section
	Title string `json:"title"`
}

// Entry represents a clickable link with an icon, title and short explanation for more context
type Entry struct {
	// Text is the displayed text right next to the icon
	Text string `json:"text"`
	// Link that is opened when the entry is selected
	Link string `json:"link"`
	// Icon is a local path to an image file or a remote publicly available web url
	Icon string `json:"icon"`
	// Description provides a short and concise description about the entry
	Description string `json:"description"`
}

func (e *Entry) InlineIcon() string {
	return assets.Must(assets.InlineIcon(e.Icon))
}
