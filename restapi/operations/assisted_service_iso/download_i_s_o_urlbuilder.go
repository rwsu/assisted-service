// Code generated by go-swagger; DO NOT EDIT.

package assisted_service_iso

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// DownloadISOURL generates an URL for the download i s o operation
type DownloadISOURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DownloadISOURL) WithBasePath(bp string) *DownloadISOURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DownloadISOURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DownloadISOURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/assisted-service-iso/data"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/assisted-install/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DownloadISOURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DownloadISOURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DownloadISOURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DownloadISOURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DownloadISOURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DownloadISOURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
