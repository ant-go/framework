package icodec

type CodecInterface interface {
	Marshal(value any) (data []byte, err error)
	Unmarshal(data []byte, value any) (err error)
}
