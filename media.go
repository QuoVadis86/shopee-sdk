package shopee

import (
	"context"
)

type MediaService struct {
	client *Client
}

func NewMediaService(client *Client) *MediaService {
	return &MediaService{client: client}
}

func (s *MediaService) CancelVideoUpload(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathMediaCancelVideoUpload, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *MediaService) CompleteVideoUpload(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathMediaCompleteVideoUpload, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *MediaService) GetVideoUploadResult(videouploadid string) (*BaseResponse, error) {
    q := map[string]string{
		"video_upload_id": videouploadid,
	}
    result := &BaseResponse{}
    if err := s.client.DoGet(context.Background(), PathMediaGetVideoUploadResult, q, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *MediaService) InitVideoUpload(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathMediaInitVideoUpload, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *MediaService) UploadImage(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathMediaUploadImage, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *MediaService) UploadVideoPart(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathMediaUploadVideoPart, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}
