package cdb

const (
	startingHash = 5381
	size = 4
)

type hashImpl struct {
	uint32
}

func (h *hashImpl) Sum32() uint32 {
	return h.uint32
}

func (h *hashImpl) Write(data []byte) (int, error) {
	var val uint32
	val = startingHash

	for _, c := range data {
		val = ((val << 5) + val) ^ uint32(c)
	}

	h.uint32 = val

	return len(data), nil
}

func (h *hashImpl) Reset() {}

func (h *hashImpl) Sum(b []byte) []byte {
	s := h.Sum32()
	return append(b, byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

func (h *hashImpl) BlockSize() int {
	return 1
}

func (h *hashImpl) Size() int {
	return size
}
