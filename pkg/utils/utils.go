package utils

import (
	"context"
)

type TrackingKey string

const TrackingIDKey TrackingKey = "log_tracking_id"

func GetTrackingIdFromContext(c context.Context) string {
	val := c.Value(TrackingIDKey)
	if val == nil {
		return ""
	}
	trackingID, ok := val.(string)
	if !ok {
		return ""
	}
	return trackingID
}
