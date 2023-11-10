package service

func (c *Client) IfHostOrIpVerified(clientIP, clientHost string) (bool, error) {
	return c.db.IfHostOrIpVerified(clientIP, clientHost)
}
