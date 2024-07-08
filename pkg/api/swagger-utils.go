package pkg

import (
	"embed"
	"encoding/json"
	"io"
	"reflect"

	"github.com/wildberries-tech/sgroups/v2/pkg/api/sgroups"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
)

var (
	//go:embed */*.swagger.json
	swaggerStore embed.FS
)

var (
	//ErrSwaggerNotExist when document is not found
	ErrSwaggerNotExist = errors.New("swagger doc is no exist")
)

// SwaggerUtil ...
type SwaggerUtil[T any] struct{}

func (u SwaggerUtil[T]) reg(p string) {
	swaggerPaths[reflect.TypeOf((*T)(nil)).Elem()] = p
}

// GetSpec ...
func (u SwaggerUtil[T]) GetSpec() (*spec.Swagger, error) {
	res := new(spec.Swagger)
	ty := reflect.TypeOf((*T)(nil)).Elem()
	p := swaggerPaths[ty]
	err := whenFindSwagger(p, func(r io.Reader) error {
		return json.NewDecoder(r).Decode(res)
	})
	return res, err
}

// GetRaw ...
func (u SwaggerUtil[T]) GetRaw() (json.RawMessage, error) {
	var ret json.RawMessage
	ty := reflect.TypeOf((*T)(nil)).Elem()
	p := swaggerPaths[ty]
	err := whenFindSwagger(p, func(reader io.Reader) error {
		data, e := io.ReadAll(reader)
		ret = data
		return e
	})
	return ret, err
}

var (
	swaggerPaths = make(map[reflect.Type]string)
)

func whenFindSwagger(p string, f func(reader io.Reader) error) error {
	data, e := swaggerStore.Open(p)
	if e != nil {
		return ErrSwaggerNotExist
	}
	return f(data)
}

func init() {
	const (
		apiSGroups = "sgroups/api.swagger.json"
	)

	SwaggerUtil[sgroups.SecGroupServiceServer]{}.
		reg(apiSGroups)

	SwaggerUtil[sgroups.SecGroupServiceClient]{}.
		reg(apiSGroups)
}
