package store

import "time"

type (
	Status string
	Type   string
	Key    string
)

const (
	StatusWaitingToStart Status = "Waiting to start"
	StatusJobQueued      Status = "Job queued"
	StatusProcessing     Status = "Processing"
	StatusComplete       Status = "Complete"
	StatusUnknown        Status = "Unknown"

	TypeImagine          Type          = "Imagine"
	TypeUpscale          Type          = "Upscale"
	TypeUnknown          Type          = "Unknown"
	DrmMidjourneyCaptcha               = `drm:midjourney_server`
	Expired              time.Duration = 3 * time.Hour
)
