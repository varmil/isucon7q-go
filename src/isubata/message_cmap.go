package main

import (
	"strconv"
	"sync"

	cmap "github.com/orcaman/concurrent-map"
)

// MessageCmap contains cmap
type MessageCmap struct {
	// wg *sync.WaitGroup
	mx *sync.Mutex
	r  cmap.ConcurrentMap
}

// NewMessageCmap returns the instance
func NewMessageCmap() *MessageCmap {
	return &MessageCmap{
		// wg: &sync.WaitGroup{},
		mx: &sync.Mutex{},
		r:  cmap.New(),
	}
}

// Store the instance with sorted by messageID DESC
func (s *MessageCmap) Store(channelID int64, message *Message) {
	var slice []*Message

	s.mx.Lock()
	defer s.mx.Unlock()

	// s.wg.Wait()
	// s.wg.Add(1)
	// defer s.wg.Done()

	slice = *(s.LoadWithoutLock(channelID))

	// ID採番。MySQLの採番は無視。
	message.ID = int64(len(slice) + 1)

	// 無理やりsliceにしてprepend表現
	slice = append([]*Message{message}, slice...)
	// sort.Slice(slice, func(i, j int) bool { return slice[i].ID > slice[j].ID })
	s.r.Set(toString(channelID), &slice)

}

// LoadWithoutLock does not use Lock
func (s *MessageCmap) LoadWithoutLock(channelID int64) *[]*Message {
	t, ok := s.r.Get(toString(channelID))
	if !ok {
		return &[]*Message{}
	}
	return t.(*[]*Message)
}

// Load the instance, return empty slice if not exists
func (s *MessageCmap) Load(channelID int64) (*[]*Message, bool) {
	// s.mx.Lock()
	// defer s.mx.Unlock()
	// s.wg.Wait()

	t, ok := s.r.Get(toString(channelID))

	if !ok {
		return &[]*Message{}, false
	}
	return t.(*[]*Message), true
}

// Delete the instance
func (s *MessageCmap) Delete(channelID int64) {
	s.r.Remove(toString(channelID))
}

// Count all of the messages for channelID
func (s *MessageCmap) Count(channelID int64) int64 {
	t, _ := s.Load(channelID)
	return int64(len(*t))
}

func toString(n int64) string {
	return strconv.FormatInt(n, 10)
}
