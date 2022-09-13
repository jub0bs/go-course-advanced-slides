package magicnumbers

import (
	"fmt"
	"net/http"
	"net/url"
)

// START OMIT
func IsAvailableOnGitHub(username string) (bool, error) {
	const tmpl = "https://github.com/%s"
	endpoint := fmt.Sprintf(tmpl, url.PathEscape(username))
	req, err := http.NewRequest(http.MethodGet, endpoint, nil) // HL
	if err != nil {
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusNotFound, nil // HL
}

// END OMIT
