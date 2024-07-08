package pkg

import (
	"io"
	"reflect"
	"sync"

	"github.com/wildberries-tech/sgroups/v2/pkg/api/sgroups"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// ClosableClient closable client
type ClosableClient[T any] struct {
	C     T
	close func() error
}

// MustInit panics when Init returns error
func (sc *ClosableClient[T]) MustInit(conn grpc.ClientConnInterface) {
	if err := sc.Init(conn); err != nil {
		panic(err)
	}
}

// Init initialize closable client
func (sc *ClosableClient[T]) Init(conn grpc.ClientConnInterface) error {
	if cc := (clientConstructor[T])(nil); cc.load() {
		sc.C = cc(conn)
		s, err := new(sync.Once), new(error)
		sc.close = func() error {
			s.Do(func() {
				if c, _ := conn.(io.Closer); c != nil {
					*err = c.Close()
				}
			})
			return *err
		}
		return nil
	}
	return errors.WithMessagef(ErrConstructorNotReg, "when init '%s'",
		reflect.TypeOf(sc).Elem().Name())
}

// Close impl Closable
func (sc ClosableClient[T]) Close() error {
	return sc.close()
}

var (
	//ErrConstructorNotReg constructor is not found
	ErrConstructorNotReg = errors.New("constructor is not registered")

	clientConstructors = make(map[reflect.Type]reflect.Value)
)

type clientConstructor[T any] func(grpc.ClientConnInterface) T

func (cc *clientConstructor[T]) load() bool {
	ty := reflect.TypeOf(cc).Elem()
	if x, ok := clientConstructors[ty]; ok {
		reflect.ValueOf(cc).Elem().Set(x)
	}
	return *cc != nil
}

func (cc clientConstructor[T]) reg(f func(grpc.ClientConnInterface) T) {
	ty := reflect.TypeOf(cc)
	clientConstructors[ty] = reflect.ValueOf(f)
}

func init() {
	(clientConstructor[sgroups.SecGroupServiceClient])(nil).
		reg(sgroups.NewSecGroupServiceClient)
}
