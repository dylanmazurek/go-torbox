package general

// func (s *GeneralService) CreateWebDownload(r models.CreateWebDownloadRequest) (*models.WebDownload, error) {
// 	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_WEBDL_CREATE, request.WithJSONBody(r))
// 	if err != nil {
// 		return nil, err
// 	}

// 	var resp models.CreateWebDownloadResponse
// 	err = s.internalClient.Do(req, &resp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !resp.Success {
// 		return nil, fmt.Errorf("failed to create web download: %s", resp.Detail)
// 	}

// 	return resp.Data, nil
// }

// func (s *GeneralService) ControlWebDownload(webId int64, operation constants.ControlWebDownloadOperation) error {
// 	r := models.ControlWebDownloadRequest{
// 		WebID:     webId,
// 		Operation: operation,
// 	}

// 	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_WEBDL_CONTROL, request.WithJSONBody(r))
// 	if err != nil {
// 		return err
// 	}

// 	var resp models.BaseResponse
// 	err = s.internalClient.Do(req, &resp)
// 	if err != nil {
// 		return err
// 	}

// 	if !resp.Success {
// 		return fmt.Errorf("failed to control web download: %s", resp.Detail)
// 	}

// 	return nil
// }
