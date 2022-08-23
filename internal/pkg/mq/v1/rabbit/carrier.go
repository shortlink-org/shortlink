package rabbit

type amqpHeadersCarrier map[string]interface{}

func (c amqpHeadersCarrier) Get(key string) string {
	val, ok := c[key]
	if !ok {
		return ""
	}

	return val.(string)
}

func (c amqpHeadersCarrier) Keys() []string {
	var keys []string

	for k := range c {
		keys = append(keys, k)
	}

	return keys
}

// ForeachKey conforms to the TextMapReader interface.
func (c amqpHeadersCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, val := range c {
		v, ok := val.(string)
		if !ok {
			continue
		}
		if err := handler(k, v); err != nil {
			return err
		}
	}

	return nil
}

// Set implements the TextMapWriter interface.
func (c amqpHeadersCarrier) Set(key, val string) {
	c[key] = val
}
