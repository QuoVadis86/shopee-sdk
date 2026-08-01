package shopee

import (
	"context"
)

type PushService struct {
	client *Client
}

func NewPushService(client *Client) *PushService {
	return &PushService{client: client}
}

func (s *PushService) ConfirmConsumedLostPushMessage(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathPushConfirmConsumedLostPushMessage, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}
