package exchange

type TestGenerator struct {
	exchange string
	stopCh   chan struct{}
}

func NewTestGenerator(exchange string) *TestGenerator {
	return &TestGenerator{
		exchange: exchange,
		stopCh:   make(chan struct{}),
	}
}


