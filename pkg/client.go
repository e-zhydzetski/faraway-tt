package farawaytt

import "github.com/e-zhydzetski/faraway-tt/internal/app"

var NewClient = app.NewClient                     // nolint:gochecknoglobals // API
var DefaultClientConfig = app.DefaultClientConfig // nolint:gochecknoglobals // API

type ClientConfig = app.ClientConfig
type Client = app.Client
