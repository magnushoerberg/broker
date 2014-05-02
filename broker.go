package broker

type Broker struct {
    channels map[string]([]chan interface{})
}

func (b Broker) getChannels(routingKey string) ([]chan interface{}) {
    channels := b.channels[routingKey]
    if channels == nil {
        channels = make([]chan interface{}, 0)
    }
    return channels
}
func (b Broker) createChannel(routingKey string) (chan interface{}) {
    channels := b.getChannels(routingKey)
    channel := make(chan interface{})
    channels = append(channels, channel)
    b.channels[routingKey] = channels
    return channel
}
func (b Broker) Pub(routingKey string, msg interface{}) {
    channels := b.getChannels(routingKey)
    for _, c := range channels {
        c <- msg
    }
}
func (b Broker) Sub(routingKey string, callback func(interface{})) {
    channel := b.createChannel(routingKey)
    go func(channel chan interface{}) {
        for msg := range channel {
            callback(msg)
        }
    }(channel)
}
func NewBroker() (Broker) {
    return Broker{
        channels: make(map[string][]chan interface{}),
    }
}
