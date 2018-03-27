package structs

import "time"

type EC2Instance struct {
	InstanceId            string
	ImageId               string
	PrivateDnsName        string
	KeyName               string
	InstanceType          string
	VpcId                 string
	PrivateIpAddress      string
	IamInstanceProfile    string
	Tags                  map[string]string
	LaunchTime            time.Time
	SpotInstanceRequestId string
	State                 string
}
