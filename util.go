package pulsarzap

import "github.com/apache/pulsar-client-go/pulsar/log"

func pulsarFieldsToKVSlice(f log.Fields) []interface{} {
	ret := make([]interface{}, 0, len(f)*2)
	for k, v := range f {
		ret = append(ret, k, v)
	}
	return ret
}
