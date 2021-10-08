//// +build kafka



package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
)

type PartitionStrategy int

const (
	RANDOM      PartitionStrategy = iota
	ROUND_ROBIN PartitionStrategy = iota
	EXPRESSION  PartitionStrategy = iota
	DEFAULT     PartitionStrategy = iota
)

func getPartitionerConstructor(strategy PartitionStrategy) (sarama.PartitionerConstructor, error) {
	switch strategy {
	case DEFAULT:
		return sarama.NewManualPartitioner, nil
	case RANDOM:
		return sarama.NewRandomPartitioner, nil
	case ROUND_ROBIN:
		return sarama.NewRoundRobinPartitioner, nil
	default:
		return nil, errors.New("Unsupported/Unrecognized Partitioner Type")
	}
}
