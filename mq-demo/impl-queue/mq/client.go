package mq

// Client ...
type Client struct {
	bro *BrokerImpl
}

// NewClient ...
func NewClient() *Client {
	return &Client{
		bro: NewBroker(),
	}
}

// SetConditions ...
func (c *Client) SetConditions(capacity int) {
	c.bro.setConditions(capacity)
}

// Publish ...
func (c *Client) Publish(topic string, msg interface{}) error {
	return c.bro.publish(topic, msg)
}

// Subscribe ...
func (c *Client) Subscribe(topic string) (<-chan interface{}, error) {
	return c.bro.subscribe(topic)
}

// Unsubscribe ...
func (c *Client) Unsubscribe(topic string, sub <-chan interface{}) error {
	return c.bro.unsubscribe(topic, sub)
}

// Close ...
func (c *Client) Close() {
	c.bro.close()
}

// GetPayLoad ...
func (c *Client) GetPayLoad(sub <-chan interface{}) interface{} {
	for val := range sub {
		if val != nil {
			return val
		}
	}
	return nil
}
