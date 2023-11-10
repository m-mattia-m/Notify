package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"message-proxy/internal/model"
)

func (dc *DaoClient) IfHostVerified(clientIP, clientHost string) (bool, error) {
	ctx := context.Background()
	filter := bson.M{
		"host": bson.M{
			"$in": []string{clientIP, clientHost},
		},
	}
	hostsResponse, err := dc.engine.Database(dc.dbName).Collection("host").Find(ctx, filter)
	if err != nil {
		return false, err
	}

	defer hostsResponse.Close(ctx)

	var hosts []model.Host
	for hostsResponse.Next(ctx) {
		var host model.Host
		err = hostsResponse.Decode(&hosts)
		if err != nil {
			return false, err
		}
		hosts = append(hosts, host)
	}

	if len(hosts) > 0 {
		return true, nil
	}
	return false, nil
}
