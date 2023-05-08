package main

import (
	"fmt"
	"gomock/event"
	"google.golang.org/grpc"
	"net"
	"sync"
)

// EventServer is a gRPC server for event subscription.
type EventServer struct {
	mu          sync.RWMutex
	eventSource map[string][]chan<- *event.Event
}

// Subscribe subscribes to events from the given event source.
func (s *EventServer) Subscribe(source *event.EventSource, stream event.EventService_SubscribeServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create a channel for the stream.
	ch := make(chan *event.Event)
	s.eventSource[source.Name] = append(s.eventSource[source.Name], ch)

	// Send events from the channel to the stream.
	for event := range ch {
		if err := stream.Send(event); err != nil {
			// Remove the channel from the subscriber list.
			for i, c := range s.eventSource[source.Name] {
				if c == ch {
					s.eventSource[source.Name] = append(s.eventSource[source.Name][:i], s.eventSource[source.Name][i+1:]...)
					break
				}
			}
			return err
		}
	}
	return nil
}

func (s *EventServer) Dispatch(source *event.EventSource, event *event.Event) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, ch := range s.eventSource[source.Name] {
		ch <- event
	}
}

func main() {
	server := grpc.NewServer()

	eventServer := &EventServer{
		eventSource: make(map[string][]chan<- *event.Event),
	}
	fmt.Println(eventServer)
	// event.RegisterEventServiceServer(server, eventServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {

	}
	if err := server.Serve(listener); err != nil {

	}
}
