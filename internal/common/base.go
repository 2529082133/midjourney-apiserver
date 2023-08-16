package common

import (
	"github.com/2529082133/midjourney-api/midjourney-go/midjourney"
	"github.com/2529082133/midjourney-apiserver/internal/config"
	"github.com/2529082133/midjourney-apiserver/pkg/store"
	"github.com/bwmarrin/discordgo"
)

type Base struct {
	*discordgo.Session
	Store    *store.Store
	MJClient *midjourney.Client
	Config   *config.Config
}
