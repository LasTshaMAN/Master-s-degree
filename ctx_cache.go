package main

// ctxCache helps us to avoid making unnecessary requests for context
type ctxCache struct {
	storage map[string]Context
}

func newCtxCache() ctxCache {
	return ctxCache{
		storage: map[string]Context{},
	}
}

func (c *ctxCache) GetCtx(word string) Context {
	if ctx, ok := c.storage[word]; ok {
		return ctx
	}
	ctx := getContext(word)
	c.storage[word] = ctx
	return ctx
}
