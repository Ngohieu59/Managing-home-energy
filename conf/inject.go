package conf

import "github.com/samber/do"

func Inject(di *do.Injector) {
	do.Provide(di, NewConfig) // Registering a dependency and how to create it
}
