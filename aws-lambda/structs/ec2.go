package structs

type EC2Instance struct {
	InstanceId       string
	ImageId          string
	PrivateDnsName   string
	KeyName          string
	InstanceType     string
	VpcId            string
	PrivateIpAddress string
	Tags             map[string]string
}
