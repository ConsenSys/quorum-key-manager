package client

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type AwsKmsClient struct {
	client *kms.KMS
	cfg    *Config
}

func NewKmsClient(cfg *Config) (*AwsKmsClient, error) {
	newSession, err := session.NewSession(cfg.ToAWSConfig())
	if err != nil {
		return nil, err
	}

	return &AwsKmsClient{
		client: kms.New(newSession),
		cfg:    cfg,
	}, nil
}

func (c *AwsKmsClient) CreateKey(ctx context.Context, id, keyType string, tags []*kms.Tag) (*kms.CreateKeyOutput, error) {
	// Always create with same usage for key now (sign & verify)
	keyUsage := kms.KeyUsageTypeSignVerify

	out, err := c.client.CreateKey(&kms.CreateKeyInput{
		CustomerMasterKeySpec: &keyType,
		KeyUsage:              &keyUsage,
		Tags:                  tags,
	})
	if err != nil {
		return nil, err
	}

	_, err = c.client.CreateAlias(&kms.CreateAliasInput{
		AliasName:   &id,
		TargetKeyId: out.KeyMetadata.KeyId,
	})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return out, nil
}

func (c *AwsKmsClient) GetPublicKey(ctx context.Context, keyID string) (*kms.GetPublicKeyOutput, error) {
	out, err := c.client.GetPublicKey(&kms.GetPublicKeyInput{
		KeyId: &keyID,
	})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return out, nil
}

func (c *AwsKmsClient) DescribeKey(ctx context.Context, id string) (*kms.DescribeKeyOutput, error) {
	out, err := c.client.DescribeKey(&kms.DescribeKeyInput{KeyId: &id})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return out, nil
}

func (c *AwsKmsClient) ListKeys(ctx context.Context) ([]string, error) {
	var keys []string
	err := c.client.ListKeysPages(&kms.ListKeysInput{}, func(res *kms.ListKeysOutput, lastPage bool) bool {
		for _, k := range res.Keys {
			keys = append(keys, *k.KeyArn)
		}

		return !lastPage
	})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return keys, nil
}

func (c *AwsKmsClient) GetAlias(ctx context.Context, keyID string) (string, error) {
	out, err := c.client.ListAliases(&kms.ListAliasesInput{
		KeyId: aws.String(keyID),
		Limit: aws.Int64(1),
	})
	if err != nil {
		return "", parseKmsErrorResponse(err)
	}

	return *out.Aliases[0].AliasName, nil
}

func (c *AwsKmsClient) ListTags(ctx context.Context, id, marker string) (*kms.ListResourceTagsOutput, error) {
	input := &kms.ListResourceTagsInput{KeyId: &id}
	if len(marker) > 0 {
		input.Marker = &marker
	}

	tags, err := c.client.ListResourceTags(input)
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return tags, nil
}

func (c *AwsKmsClient) Sign(ctx context.Context, keyID string, msg []byte, signingAlgorithm string) (*kms.SignOutput, error) {
	// Message type is always digest
	msgType := kms.MessageTypeDigest
	out, err := c.client.Sign(&kms.SignInput{
		KeyId:            &keyID,
		Message:          msg,
		MessageType:      &msgType,
		SigningAlgorithm: &signingAlgorithm,
	})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return out, nil
}

func (c *AwsKmsClient) DeleteKey(ctx context.Context, keyID string) (*kms.ScheduleKeyDeletionOutput, error) {
	out, err := c.client.ScheduleKeyDeletion(&kms.ScheduleKeyDeletionInput{
		KeyId: &keyID,
	})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return out, nil
}

func (c *AwsKmsClient) RestoreKey(ctx context.Context, keyID string) (*kms.CancelKeyDeletionOutput, error) {
	out, err := c.client.CancelKeyDeletion(&kms.CancelKeyDeletionInput{
		KeyId: &keyID,
	})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return out, nil
}

func (c *AwsKmsClient) UpdateKey(ctx context.Context, keyID string, tags []*kms.Tag) (*kms.TagResourceOutput, error) {
	outTagResource, err := c.client.TagResource(&kms.TagResourceInput{
		KeyId: &keyID,
		Tags:  tags,
	})
	if err != nil {
		return nil, parseKmsErrorResponse(err)
	}

	return outTagResource, nil
}
