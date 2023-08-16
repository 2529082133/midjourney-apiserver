package service

import (
	"context"
	"fmt"
	"github.com/2529082133/midjourney-api/midjourney-go/midjourney"
	"github.com/2529082133/midjourney-apiserver/pkg/api"
	"github.com/2529082133/midjourney-apiserver/pkg/store"
	"github.com/google/uuid"
	"log"
	"time"
)

/*
flow:
1. create mesasge id: 3
2. update message id: 2
3. create message id: 4 -> contains attachments
4. delete message id: 3
*/
func (s *Service) Variation(ctx context.Context, in *api.VariationRequest) (*api.VariationResponse, error) {
	if in.RequestId == "" {
		in.RequestId = uuid.NewString()
	}

	metaData, err := s.Store.GetMetaData(ctx, in.TaskId)
	if err != nil {
		e := err.(store.Error)
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      e.Code,
			Msg:       e.Msg,
		}, nil
	}

	if metaData == nil {
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      api.Codes_CODES_INVALID_PARAMETER_ERROR,
			Msg:       fmt.Sprintf("id: %s, not found", in.TaskId),
		}, nil
	}

	if metaData.Type != store.TypeImagine {
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      api.Codes_CODES_INVALID_PARAMETER_ERROR,
			Msg:       fmt.Sprintf("id: %s, the type is not `Imagine`, %s", in.TaskId, metaData.Type),
		}, nil
	}

	if metaData.Status != store.StatusComplete {
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      api.Codes_CODES_INVALID_PARAMETER_ERROR,
			Msg:       fmt.Sprintf("id: %s, the status is not `Complete`, %s", in.TaskId, metaData.Status),
		}, nil
	}

	url, err := metaData.GetImageURL()
	if err != nil {
		e := err.(store.Error)
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      e.Code,
			Msg:       e.Msg,
		}, nil
	}

	key := store.GetKey(metaData.Prompt)

	log.Printf("Upscale, key: %s, len: %d", key, len(key))

	if !KeyChan.Init(key) {
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      api.Codes_CODES_INVALID_PARAMETER_ERROR,
			Msg:       fmt.Sprintf("The same prompt is being processed, please try again later."),
		}, nil
	}

	defer KeyChan.Del(key)

	if err := s.MJClient.Variation(ctx, &midjourney.VariationRequest{
		Index:       in.Index,
		GuildID:     s.Config.Midjourney.GuildID,
		ChannelID:   s.Config.Midjourney.ChannelID,
		MessageID:   metaData.CompleteMessageID,
		MessageHash: midjourney.GetMessageHash(url),
	}); err != nil {
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      api.Codes_CODES_SERVER_INTERNAL_ERROR,
			Msg:       fmt.Sprint(err),
		}, nil
	}

	select {
	case <-time.After(10 * time.Second):
		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      api.Codes_CODES_SERVER_INTERNAL_ERROR,
			Msg:       "timeout",
		}, nil
	case msgInfo := <-KeyChan.Get(key):
		if msgInfo.Error != nil {
			code := api.Codes_CODES_SERVER_INTERNAL_ERROR

			switch msgInfo.Error.Title {
			case "Invalid parameter":
				code = api.Codes_CODES_INVALID_PARAMETER_ERROR
			}

			return &api.VariationResponse{
				RequestId: in.RequestId,
				Code:      code,
				Msg:       msgInfo.Error.Description,
			}, nil
		}

		if err := s.Store.SaveWebhook(ctx, msgInfo.ID, in.Webhook, in.RequestId); err != nil {
			return &api.VariationResponse{
				RequestId: in.RequestId,
				Code:      api.Codes_CODES_SERVER_INTERNAL_ERROR,
				Msg:       fmt.Sprint(err),
			}, nil
		}

		return &api.VariationResponse{
			RequestId: in.RequestId,
			Code:      api.Codes_CODES_SUCCESS,
			Msg:       "success",
			Data: &api.VariationResponseData{
				TaskId:    msgInfo.ID,
				StartTime: msgInfo.StartTime,
			},
		}, nil
	}
}
