package config

import "os"
import "testing"

func TestGetGitHubToken(t *testing.T){
  os.Setenv("GITHUB_API", "foo")

  token := GetGitHubToken()

  if token != "foo" {
    t.Errorf("GetGitHubToken() == %q, expected %q", token, "foo")
  }
}
