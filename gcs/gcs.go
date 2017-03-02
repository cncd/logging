package google

//
// import (
// 	"context"
// 	"fmt"
// 	"sync"
//
// 	"github.com/cncd/logging"
//
// 	"cloud.google.com/go/logging"
// )
//
// type subscriber struct {
// 	handler stream.Handler
// }
//
// type channel struct {
// 	sync.Mutex
//
// 	path string
// 	hist []*stream.Entry
// 	subs map[*subscriber]struct{}
// 	done chan struct{}
// }
//
// type mux struct {
// 	sync.Mutex
//
// 	client  *logging.Client
// 	logger  *logging.Logger
// 	project string
//
// 	list map[string]*channel
// }
//
// // New returns a new stream multiplexer.
// func New(client *logging.Client) (stream.Mux, error) {
// 	return &mux{
// 		client: client,
// 		logger: client.Logger("cncd-stream"),
// 	}, nil
// }
//
// func (m *mux) Open(c context.Context, path string) error {
//
// 	return nil
// }
//
// func (m *mux) Write(c context.Context, path string, entry *stream.Entry) error {
// 	p := fmt.Sprintf("projects/%s/logs/%s", m.project, path)
// 	e := logging.Entry{Payload: entry, LogName: p}
// 	m.logger.Log(e)
// 	return nil
// }
//
// func (m *mux) Tail(c context.Context, path string, handler stream.Handler) error {
// 	return nil
// }
//
// func (m *mux) Close(c context.Context, path string) error {
// 	return nil
// }
