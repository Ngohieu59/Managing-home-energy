package mysql

import "github.com/samber/do"

func Inject(di *do.Injector) {
	do.Provide(di, newUserRepo)
	do.Provide(di, newEbillRepo)
}
