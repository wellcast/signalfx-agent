package hostid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GoogleComputeID generates a unique id for the compute instance that the
// agent is running on.  It returns a blank string if we are not running on GCP
// or there was an error getting the metadata.
func GoogleComputeID() string {
	projectID := getMetadata("project/project-id")
	if projectID == "" {
		return ""
	}
	instanceID := getMetadata("instance/id")
	if instanceID == "" {
		return ""
	}

	return fmt.Sprintf("%s_%s", projectID, instanceID)
}

func getMetadata(path string) string {
	url := fmt.Sprintf("http://metadata.google.internal/computeMetadata/v1/%s", path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// This would only be due to a programming bug
		panic("Could not construct request for GCP ID")
	}

	req.Header.Add("Metadata-Flavor", "Google")

	c := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		return ""
	}

	if resp.Header.Get("Metadata-Flavor") != "Google" {
		return ""
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(data)
}
