package helper

import (
	"ms-go-gin/config"

	"github.com/go-resty/resty/v2"
)

func GetRestyClient(cfg *config.Config) *resty.Client {
	var client = resty.New().
		SetBaseURL(cfg.SupabaseUrl).
		SetHeader("apikey", cfg.SupabaseApikey).SetDebug(true)

	return client
}
