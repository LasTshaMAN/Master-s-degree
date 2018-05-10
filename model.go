package main

type meaning struct {
	representations []string
	googleRank      string
	translation     string
	rank            int
}

func newMeaning(representations []string, rank string, translation string) meaning {
	return meaning{
		representations: representations,
		googleRank:      rank,
		translation:     translation,
	}
}

func (m *meaning) Representations() []string {
	return m.representations
}

func (m *meaning) IncreaseRank(value int) {
	m.rank += value
}

func (m *meaning) IsNotRare() bool {
	return m.googleRank != "8"
}

func (m *meaning) Translation() string {
	return m.translation
}

func (m *meaning) Rank() int {
	return m.rank
}

type Context struct {
	words []string
}

func newContext(words []string) Context {
	return Context{
		words: words,
	}
}

func (c *Context) Words() []string {
	return c.words
}
