package cloudwatchevents

type EventDetail struct {
	AutoScalingGroupName string
	StatusCode           string
	EndTime              string
	EC2InstanceId        string
	StartTime            string
	Cause                string
}
