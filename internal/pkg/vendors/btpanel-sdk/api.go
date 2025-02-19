package btpanelsdk

func (c *Client) ConfigSavePanelSSL(req *ConfigSavePanelSSLRequest) (*ConfigSavePanelSSLResponse, error) {
	resp := ConfigSavePanelSSLResponse{}
	err := c.sendRequestWithResult("/config?action=SavePanelSSL", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SiteSetSSL(req *SiteSetSSLRequest) (*SiteSetSSLResponse, error) {
	resp := SiteSetSSLResponse{}
	err := c.sendRequestWithResult("/site?action=SetSSL", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SystemServiceAdmin(req *SystemServiceAdminRequest) (*SystemServiceAdminResponse, error) {
	resp := SystemServiceAdminResponse{}
	err := c.sendRequestWithResult("/system?action=ServiceAdmin", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SSLCertSaveCert(req *SSLCertSaveCertRequest) (*SSLCertSaveCertResponse, error) {
	resp := SSLCertSaveCertResponse{}
	err := c.sendRequestWithResult("/ssl/cert/save_cert", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SSLSetBatchCertToSite(req *SSLSetBatchCertToSiteRequest) (*SSLSetBatchCertToSiteResponse, error) {
	resp := SSLSetBatchCertToSiteResponse{}
	err := c.sendRequestWithResult("/ssl?action=SetBatchCertToSite", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
