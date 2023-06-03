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

func (j *jsonManager) Output(ctx context.Context, out interface{}) error {
	if m, ok := out.(proto.Message); ok {
		return j.printProto(ctx, m)
	}

	return j.printInterface(ctx, out)
}

func (j *jsonManager) printProto(ctx context.Context, m proto.Message) error {
	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	if j.pretty {
		opts.Multiline = true
		opts.Indent = "  "
	}

	outBytes, err := opts.Marshal(m)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(os.Stdout, string(outBytes))
	if err != nil {
		return err
	}

	return nil
}

func (j *jsonManager) printInterface(ctx context.Context, data interface{}) error {
	if j.pretty {
		prettyJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(os.Stdout, string(prettyJSON))
		if err != nil {
			return err
		}
		return nil
	}

	plainJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(os.Stdout, string(plainJSON))
	if err != nil {
		return err
	}

	return nil
}
