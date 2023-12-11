//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

//salah
//func InitializedHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHello)
//	return nil
//}

var helloSet = wire.NewSet(
	NewSayHello,
	wire.Bind(new(SayHelloInterface), new(*SayHello)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

var fooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializedFooBar() *FooBar {
	wire.Build(fooBarSet, wire.Struct(new(FooBar), "*"))
	return nil
}

var fooValuw = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValuw), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
