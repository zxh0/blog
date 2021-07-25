package bcapi

type ParallelClient struct {
	client SimpleClient

	startHeight uint32
	nGoroutines uint32

	hChans []chan uint32
	bChans []chan string
}

func NewParallelClient(startHeight, nGoroutines uint32) Client {
	c := &ParallelClient{
		client:      SimpleClient{debug: false},
		startHeight: startHeight,
		nGoroutines: nGoroutines,
	}
	c.start()
	return c
}

func (c *ParallelClient) start() {
	c.hChans = make([]chan uint32, c.nGoroutines)
	c.bChans = make([]chan string, c.nGoroutines)

	for i := uint32(0); i < c.nGoroutines; i++ {
		c.hChans[i] = make(chan uint32)
		c.bChans[i] = make(chan string)

		go func(hChan chan uint32, bChan chan string) {
			for {
				h := <-hChan
				b, err := c.client.GetRawBlockByHeight(h)
				if err != nil {
					panic(err)
				}

				bChan <- b
			}
		}(c.hChans[i], c.bChans[i])
	}

	for i := uint32(0); i < c.nGoroutines; i++ {
		c.hChans[i] <- c.startHeight + i
	}
}

func (c *ParallelClient) GetRawTx(hash string) (string, error) {
	return c.client.GetRawTx(hash)
}
func (c *ParallelClient) GetRawBlockByHash(hash string) (string, error) {
	return c.client.GetRawBlockByHash(hash)
}

func (c *ParallelClient) GetRawBlockByHeight(height uint32) (string, error) {
	idx := (height - c.startHeight) % c.nGoroutines
	b := <-c.bChans[idx]
	c.hChans[idx] <- height + c.nGoroutines
	return b, nil
}
