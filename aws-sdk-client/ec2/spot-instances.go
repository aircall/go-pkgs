package ec2

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type SpotInstance struct {
	InstanceId            string
	SpotInstanceRequestId string
	State                 string
	Tags                  EC2Tags
	LaunchTime            time.Time
	ValidUntil            time.Time
}

func (cli Sdk) CancelSpotInstanceRequest(requestId string) error {
	input := &ec2.CancelSpotInstanceRequestsInput{
		SpotInstanceRequestIds: []*string{aws.String(requestId)},
	}
	_, err := cli.ec2.CancelSpotInstanceRequests(input)
	return err
}

func (cli Sdk) GetSpotInstanceRequests() ([]SpotInstance, error) {
	res := []SpotInstance{}

	typeFilter := ec2.Filter{}
	typeFilter.SetName("type")
	typeFilter.SetValues([]*string{aws.String("one-time")})

	activeFilter := ec2.Filter{}
	activeFilter.SetName("state")
	activeFilter.SetValues([]*string{aws.String("active")})

	input := &ec2.DescribeSpotInstanceRequestsInput{
		Filters: []*ec2.Filter{&typeFilter, &activeFilter},
	}

	result, err := cli.ec2.DescribeSpotInstanceRequests(input)
	if err != nil {
		return res, err
	}

	for _, s := range result.SpotInstanceRequests {
		r := SpotInstance{
			InstanceId:            aws.StringValue(s.InstanceId),
			SpotInstanceRequestId: aws.StringValue(s.SpotInstanceRequestId),
			State:      aws.StringValue(s.State),
			ValidUntil: aws.TimeValue(s.ValidUntil),
			LaunchTime: aws.TimeValue(s.CreateTime),
		}
		r.Tags = func(tx []*ec2.Tag) EC2Tags {
			res := make(EC2Tags)
			for _, t := range tx {
				res[aws.StringValue(t.Key)] = aws.StringValue(t.Value)
			}
			return res
		}(s.Tags)
		res = append(res, r)
	}

	return res, err
}
