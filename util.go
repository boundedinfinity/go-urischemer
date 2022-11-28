package urischemer

import (
	"net/url"
	"path"

	"github.com/boundedinfinity/go-commoner/slicer"
)

func Detect(uri string) (UriScheme, error) {
	var scheme UriScheme
	parsed, err := url.Parse(uri)

	if err != nil {
		return scheme, ErrUriSchemeNotFoundv(uri)
	}

	return Parse(parsed.Scheme)
}

func DetectOneOf(uri string, schemes ...UriScheme) (UriScheme, error) {
	found, err := Detect(uri)

	if err != nil {
		return found, err
	}

	if !slicer.Exist(schemes, found) {
		return found, ErrUriSchemeNotFoundv(uri)
	}

	return found, err
}

func Join(uri string, elems ...string) (string, error) {
	var new string

	parsed, err := url.Parse(uri)

	if err != nil {
		return new, err
	}

	parsed.Path = path.Join(elems...)

	return parsed.String(), nil
}

func Break(uri string) (UriScheme, string, error) {
	var scheme UriScheme
	var path string

	parsed, err := url.Parse(uri)

	if err != nil {
		return scheme, path, err
	}

	scheme, err = Parse(parsed.Scheme)

	if err != nil {
		return scheme, path, err
	}

	path = parsed.Path

	return scheme, path, nil
}

func Combine(scheme UriScheme, elem ...string) string {
	parsed := &url.URL{}
	parsed.Scheme = scheme.String()
	parsed.Path = path.Join(elem...)
	return parsed.String()
}

func Clean(uri string) (string, error) {
	parsed, err := url.Parse(uri)

	if err != nil {
		return "", err
	}

	parsed.Path = path.Clean(parsed.Path)
	return parsed.String(), nil
}
