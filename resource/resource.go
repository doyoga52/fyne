// Package resource manages bundled assets in Fyne and Fyne applications
package resource

// Resource represents a single bundled resource.
// A resource has an identifying name and byte array content.
// The serialised path of a resource can be obtained which may result in a
// blocking filesystem write operation.
type Resource struct {
	Name    string
	Content []byte
}

// CachePath will return the cached location of a resource.
// If the resource has not previously been written to a cache this operation
// will block until the data is available at the returned location.
func (r *Resource) CachePath() string {
	path := filePath(r.Name)
	if !pathExists(path) {
		toFile(r)
	}

	return path
}

// NewResource returns a new resource object with the specified name and content.
// Creating a new resource in memory results in sharable binary data that may be
// serialised to the location returned by Path().
func NewResource(name string, content []byte) *Resource {
	return &Resource{
		Name:    name,
		Content: content,
	}
}
