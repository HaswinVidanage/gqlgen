package models

import "github.com/HaswinVidanage/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
