package config

import "os"

const ENV_GITHUB_API = "GITHUB_API"

func GetGitHubToken() string {
  return os.Getenv(ENV_GITHUB_API)
}
