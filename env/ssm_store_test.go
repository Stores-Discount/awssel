package env

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func TestSSMStore(t *testing.T) {

	t.Run("Get a Client", func(t *testing.T) {
		actual := NewSSMStore()
		assert.NotNil(t, actual)
	})

	t.Run("Get a Client with options", func(t *testing.T) {
		actual := NewSSMStore(SSMStoreOptions{
			roleARN:  "arn:aws:iam::721728311103:role/service-role/homer-fn-role",
			profile:  "fake-profile",
			endpoint: aws.String("fake.ssm.us-east-2.amazonaws.com"),
			region:   "us-west-1",
		})
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.sess)
	})

}
