package service

func (dbc *DbClient) IfHostOrIpVerified(clientIP, clientHost string) (bool, error) {
	return dbc.dao.IfHostVerified(clientIP, clientHost)
}
