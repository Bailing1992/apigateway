package http

type HeaderAdapter struct {
	keys    map[string]string  // normalizedKey -> key
	headers *map[string]string // key -> header value  = &requestData.Headers

	extraKeys    map[string]string  // normalizedKey -> key
	extraHeaders *map[string]string // key -> header value  = &requestData.ExtraHeaders
}

func (h *HeaderAdapter) NewHeaderAdapter(headers *map[string]string, extraHeaders *map[string]string) {
	h.headers = headers
	h.extraHeaders = extraHeaders
	h.keys = make(map[string]string, 20)
	h.extraKeys = make(map[string]string, 20)
}

func (h *HeaderAdapter) ExtraHeaders() map[string]string {
	return *h.extraHeaders
}
