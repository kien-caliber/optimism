package plasma

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-redis/redis"
)

// ErrNotFound is returned when the server could not find the input.
var ErrNotFound = errors.New("not found")

// ErrCommitmentMismatch is returned when the server returns the wrong input for the given commitment.
var ErrCommitmentMismatch = errors.New("commitment mismatch")

// ErrInvalidInput is returned when the input is not valid for posting to the DA storage.
var ErrInvalidInput = errors.New("invalid input")

// DAClient is an HTTP client to communicate with a DA storage service.
// It creates commitments and retrieves input data + verifies if needed.
// Currently only supports Keccak256 commitments but may be extended eventually.
type DAClient struct {
	redisClient *redis.Client
}

func NewDAClient(redisUrl string, verify bool) *DAClient {
	return &DAClient{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     redisUrl,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// GetInput returns the input data for the given commitment bytes.
func (c *DAClient) GetInput(ctx context.Context, key []byte) ([]byte, error) {
	keyHex := fmt.Sprintf("0x%x", key)
	fmt.Println("GetInput", keyHex)
	val, err := c.redisClient.Get(keyHex).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return []byte(val), nil
}

// SetInput sets the input data and returns the keccak256 hash commitment.
func (c *DAClient) SetInput(ctx context.Context, img []byte) ([]byte, error) {
	if len(img) == 0 {
		return nil, ErrInvalidInput
	}
	key := crypto.Keccak256(img)

	// save to redis
	keyHex := fmt.Sprintf("0x%x", key)
	fmt.Println("SetInput", keyHex)
	err := c.redisClient.Set(keyHex, img, 0).Err()
	if err != nil {
		return nil, err
	}
	return key, nil
}
