package common

import (
	"github.com/2529082133/midjourney-api/midjourney-go/midjourney"
	"github.com/bwmarrin/discordgo"
	"github.com/hongliang5316/midjourney-apiserver/internal/config"
	"github.com/hongliang5316/midjourney-apiserver/pkg/store"
)

type Base struct {
	*discordgo.Session
	Store    *store.Store
	MJClient *midjourney.Client
	Config   *config.Config
}
