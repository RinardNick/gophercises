package link

// Link represents link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse pulls links from reference
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
