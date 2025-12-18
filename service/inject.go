package service

import "github.com/samber/do"

func Inject(di *do.Injector) {
	do.Provide(di, newAuthService)
	do.Provide(di, newUserService)
	do.Provide(di, newEBillService)
}

func InjectStaff(di *do.Injector) {
	do.Provide(di, newStaffService)
}
