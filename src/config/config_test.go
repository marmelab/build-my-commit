package config

import "os"
import "testing"

func TestGetGitHubToken(t *testing.T) {
	expected := "foo"
	os.Setenv("GITHUB_API", expected)

	token := GetGitHubToken()

	if token != expected {
		t.Errorf("GetGitHubToken() == %q, expected %q", token, expected)
	}
}
