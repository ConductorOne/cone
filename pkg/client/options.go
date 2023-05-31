package client

type clientConfig struct {
	Debug bool
}

type ClientOptionFunc func(option *clientConfig)

func WithDebug(debug bool) ClientOptionFunc {
	return func(option *clientConfig) {
		option.Debug = debug
	}
}

func applyOpts(opts []ClientOptionFunc) clientConfig {
	option := clientConfig{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&option)
	}
	return option
}
