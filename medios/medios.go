package medios

import "strconv"

//* IMAGE STRUCT.
type Image struct {
	Title   string
	Format  string
	Channel string
}

// Show images.
func (i *Image) Show() string {
	return "\nIMAGE\n-Title: " + i.Title +
		"\n-Format: " + i.Format + "\n-Channel: " + i.Channel + "\n"
}

//* AUDIO STRUCT.
type Audio struct {
	Title string
	Format string
	Duration int64
}

// Show audio.
func (a *Audio) Show() string {
	return "\nAUDIO\n-Title: " + a.Title +
		"\n-Format: " + a.Format + "\n-Duration: " + strconv.FormatInt(a.Duration, 10) + "\n"
}

//* VIDEO STRUCT.
type Video struct {
	Title string
	Format string
	Frames int64
}

// Show video.
func (v *Video) Show() string {
	return "\nVIDEO\n-Title: " + v.Title +
		"\n-Format: " + v.Format + "\n-Frames: " + strconv.FormatInt(v.Frames, 10) + "\n"
}

//* INTERFACE for encapsulate the Show function.
type Multimedia interface {
	Show() string
}

//* CONTENT-WEB STRUCT for multiple registers.
type ContentWeb struct {
	Multimedia []Multimedia
}

// Show all register.
func (c *ContentWeb) Show() string {
	register := ""
	for _, m := range c.Multimedia {
		register += m.Show()
	}
	return register
}