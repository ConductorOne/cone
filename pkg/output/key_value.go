package output

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type keyValueManager struct {
}

func (k *keyValueManager) Output(ctx context.Context, out interface{}) error {
	if m, ok := out.(proto.Message); ok {
		return k.printProto(ctx, m)
	}

	return k.printInterface(ctx, out)
}

func (k *keyValueManager) printProto(ctx context.Context, m proto.Message) error {

	return nil
}

func (k *keyValueManager) printInterface(ctx context.Context, data interface{}) error {

	return nil
}
