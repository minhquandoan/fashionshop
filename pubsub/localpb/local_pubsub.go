package localpb

import (
	"context"
	"log"
	"sync"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/pubsub"
)

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel map[pubsub.Topic][]chan *pubsub.Message
	locker *sync.RWMutex
}

func NewLocalPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel: make(map[pubsub.Topic][]chan *pubsub.Message),
		locker: new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func(pb *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		pb.messageQueue <- data
		log.Println("New event published: ", data.String(), " with data ", data.Data())
	}()
	return nil
}

func(pb *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	// Lock to write in mapChannel, prevent program from deadlock
	pb.locker.Lock()

	if val, ok := pb.mapChannel[topic]; ok {
		val = append(val, c)
		pb.mapChannel[topic] = val
	} else {
		pb.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	pb.locker.Unlock()

	// Return the created channel, with unsubsribe function
	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := pb.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					pb.locker.Lock()
					pb.mapChannel[topic] = chans
					pb.locker.Unlock()
					break
				}
			}
		}
		
	}
}

func(pb *localPubSub) run() {
	log.Println("PubSub started ...")

	go func() {
		for {
			mess := <-pb.messageQueue
			log.Println("Message Dequeue: ", mess)

			if chans, ok := pb.mapChannel[mess.Channel()]; ok {
				for channel := range chans {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(chans[channel])
				}
			}
		}
	}()
}