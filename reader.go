package mutago

// Title() returns the "TIT2" tag from the loaded tags.
func (m *Metadata) Title() (string, error) {
	val, err := m.tagger.Get("TIT2")
	if err != nil {
		return "", err
	}

	return val, nil
}

// Album() returns the "TALB" tag from the loaded tags.
func (m *Metadata) Album() (string, error) {
	val, err := m.tagger.Get("TALB")
	if err != nil {
		return "", err
	}

	return val, nil
}

// Artist() returns the "TPE1" tag from the loaded tags.
func (m *Metadata) Artist() (string, error) {
	val, err := m.tagger.Get("TPE1")
	if err != nil {
		return "", err
	}

	return val, nil
}

// // Albumart() returns the "APIC" tag from the loaded tags.
// func (m *Metadata) Albumart() (*v2.Albumart, error) {
// 	val, err := m.tagger.Get("APIC")
// 	if err != nil {
// 		return "", err
// 	}

// 	return val, nil
// }

// Get() returns an arbitrary tag from the loaded tags.
func (m *Metadata) Get(tag string) (string, error) {
	val, err := m.tagger.Get(tag)
	if err != nil {
		return "", err
	}

	return val, nil
}
