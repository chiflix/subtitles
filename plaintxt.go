package subtitles

// AsPlainTXT export subtitle as plain txt
func (s *Subtitle) AsPlainTXT() (ret string) {
	for _, cap := range s.Captions {
		for _, line := range cap.Text {
			ret += line + "\n"
		}
	}
	return
}
