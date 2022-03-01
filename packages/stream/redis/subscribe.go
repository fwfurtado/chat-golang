package redis

func (s redisStream) Subscribe(id, channel string) error {
	s.mutex.RLock()

	subs, ok := s.subscriptions[id]

	if ok {
		err := subs.close()

		if err != nil {
			return err
		}
	}

	subscriber := newSubscription(id, channel, s.client)

	s.mutex.RUnlock()

	s.mutex.Lock()
	s.subscriptions[id] = subscriber
	s.mutex.Unlock()

	return nil

}
