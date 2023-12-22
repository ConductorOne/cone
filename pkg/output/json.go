package output

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type jsonManager struct {
	pretty bool
}

func (j *jsonManager) Output(ctx context.Context, out interface{}, opts ...outputOption) error {
	outBytes, err := MakeJSON(ctx, out, j.pretty)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(os.Stdout, string(outBytes))
	if err != nil {
		return err
	}

	return nil
}

func MakeJSON(ctx context.Context, out interface{}, pretty bool) ([]byte, error) {
	if m, ok := out.(proto.Message); ok {
		return MakeJSONFromProto(ctx, m, pretty)
	}

	return MakeJSONFromInterface(ctx, out, pretty)
}

func MakeJSONFromProto(ctx context.Context, m proto.Message, pretty bool) ([]byte, error) {
	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	if pretty {
		opts.Multiline = true
		opts.Indent = "  "
	}

	outBytes, err := opts.Marshal(m)
	if err != nil {
		return nil, err
	}

	return outBytes, nil
}

func MakeJSONFromInterface(ctx context.Context, data interface{}, pretty bool) ([]byte, error) {
	if pretty {
		prettyJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return nil, err
		}
		return prettyJSON, nil
	}

	plainJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return plainJSON, nil
}
