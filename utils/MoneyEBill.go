package utils

import (
	"Managing-home-energy/constants"
	"Managing-home-energy/dto"
	"math"
	"strings"
	"time"
)

const day = 31
const (
	layout = "02-01-2006" // định dạng dd-mm-yyyy
)

var (
	UnitFamilyQuantityL1 = constants.UnitFamilyLevel1.Quantity / day
	UnitFamilyQuantityL2 = constants.UnitFamilyLevel2.Quantity / day
	UnitFamilyQuantityL3 = constants.UnitFamilyLevel3.Quantity / day
	UnitFamilyQuantityL4 = constants.UnitFamilyLevel4.Quantity / day
	UnitFamilyQuantityL5 = constants.UnitFamilyLevel5.Quantity / day
	UnitFamilyPriceL1    = constants.UnitFamilyLevel1.UnitPrice
	UnitFamilyPriceL2    = constants.UnitFamilyLevel2.UnitPrice
	UnitFamilyPriceL3    = constants.UnitFamilyLevel3.UnitPrice
	UnitFamilyPriceL4    = constants.UnitFamilyLevel4.UnitPrice
	UnitFamilyPriceL5    = constants.UnitFamilyLevel5.UnitPrice
	UnitFamilyPriceL6    = constants.UnitFamilyLevel6.UnitPrice
)

var (
	BusinessUnitPriceLevel1Low = constants.UnitBusinessLowLevel1.UnitPrice
	BusinessUnitPriceLevel2Low = constants.UnitBusinessLowLevel2.UnitPrice
	BusinessUnitPriceLevel3Low = constants.UnitBusinessLowLevel3.UnitPrice

	BusinessUnitPriceLevel1Medium = constants.UnitBusinessMediumLevel1.UnitPrice
	BusinessUnitPriceLevel2Medium = constants.UnitBusinessMediumLevel2.UnitPrice
	BusinessUnitPriceLevel3Medium = constants.UnitBusinessMediumLevel3.UnitPrice

	BusinessUnitPriceLevel1High = constants.UnitBusinessHighLevel1.UnitPrice
	BusinessUnitPriceLevel2High = constants.UnitBusinessHighLevel2.UnitPrice
	BusinessUnitPriceLevel3High = constants.UnitBusinessHighLevel3.UnitPrice
)

var (
	IndustrialUnitPriceLevel1Low = constants.UnitIndustrialLowLevel1.UnitPrice
	IndustrialUnitPriceLevel2Low = constants.UnitIndustrialLowLevel2.UnitPrice
	IndustrialUnitPriceLevel3Low = constants.UnitIndustrialLowLevel2.UnitPrice

	IndustrialUnitPriceLevel1Medium = constants.UnitIndustrialMediumLevel1.UnitPrice
	IndustrialUnitPriceLevel2Medium = constants.UnitIndustrialMediumLevel2.UnitPrice
	IndustrialUnitPriceLevel3Medium = constants.UnitIndustrialMediumLevel3.UnitPrice

	IndustrialUnitPriceLevel1High = constants.UnitIndustrialHighLevel1.UnitPrice
	IndustrialUnitPriceLevel2High = constants.UnitIndustrialHighLevel1.UnitPrice
	IndustrialUnitPriceLevel3High = constants.UnitIndustrialHighLevel1.UnitPrice

	IndustrialUnitPriceLevel1Max = constants.UnitIndustrialMaxLevel1.UnitPrice
	IndustrialUnitPriceLevel2Max = constants.UnitIndustrialMaxLevel1.UnitPrice
	IndustrialUnitPriceLevel3Max = constants.UnitIndustrialMaxLevel1.UnitPrice
)

var (
	AdministrativeUnitPriceLevel1Low  = constants.UnitAdministrativeLowLevel1.UnitPrice
	AdministrativeUnitPriceLevel1High = constants.UnitAdministrativeHighLevel1.UnitPrice
)

func MoneyFamily(TotalEUsed float64, days float64) float64 {
	var TotalMoneyBeforeTax = 0.0
	QuantityL1, QuantityL2, QuantityL3, QuantityL4, QuantityL5 := math.Round(UnitFamilyQuantityL1*days), math.Round(UnitFamilyQuantityL2*days), math.Round(UnitFamilyQuantityL3*days), math.Round(UnitFamilyQuantityL4*days), math.Round(UnitFamilyQuantityL5*days)

	if TotalEUsed <= QuantityL1 {
		TotalMoneyBeforeTax = UnitFamilyPriceL1 * TotalEUsed
	} else if TotalEUsed <= (QuantityL2 + QuantityL1) {
		TotalMoneyBeforeTax = QuantityL1*UnitFamilyPriceL1 + (TotalEUsed-QuantityL1)*UnitFamilyPriceL2
	} else if TotalEUsed <= (QuantityL3 + QuantityL2 + QuantityL1) {
		TotalMoneyBeforeTax = QuantityL1*UnitFamilyPriceL1 + QuantityL2*UnitFamilyPriceL2 + (TotalEUsed-QuantityL1-QuantityL2)*UnitFamilyPriceL3
	} else if TotalEUsed <= (QuantityL1 + QuantityL2 + QuantityL4 + QuantityL3) {
		TotalMoneyBeforeTax = QuantityL1*UnitFamilyPriceL1 + QuantityL2*UnitFamilyPriceL2 + QuantityL3*UnitFamilyPriceL3 + (TotalEUsed-QuantityL1-QuantityL2-QuantityL3)*UnitFamilyPriceL4
	} else if TotalEUsed <= (QuantityL1 + QuantityL2 + QuantityL3 + QuantityL4 + QuantityL5) {
		TotalMoneyBeforeTax = QuantityL1*UnitFamilyPriceL1 + QuantityL2*UnitFamilyPriceL2 + QuantityL3*UnitFamilyPriceL3 + QuantityL4*UnitFamilyPriceL4 + (TotalEUsed-QuantityL1-QuantityL2-QuantityL3-QuantityL4)*UnitFamilyPriceL5
	} else {
		TotalMoneyBeforeTax = QuantityL1*UnitFamilyPriceL1 + QuantityL2*UnitFamilyPriceL2 + QuantityL3*UnitFamilyPriceL3 + QuantityL4*UnitFamilyPriceL4 + QuantityL5*UnitFamilyPriceL5 + (TotalEUsed-QuantityL1-QuantityL2-QuantityL3-QuantityL4-QuantityL5)*UnitFamilyPriceL6
	}

	return TotalMoneyBeforeTax
}

func MoneyBusiness(TotalEUsed *dto.ListEUsedResponse, days float64) float64 {
	var TotalMoneyBeforeTax = 0.0

	if TotalEUsed.Total < constants.UnitBusinessLowLevel1.Quantity/day*days {
		TotalMoneyBeforeTax = TotalEUsed.Low*BusinessUnitPriceLevel1Low + TotalEUsed.Normal*BusinessUnitPriceLevel2Low + TotalEUsed.High*BusinessUnitPriceLevel3Low
	} else if TotalEUsed.Total > constants.UnitBusinessHighLevel1.Quantity/day*days {
		TotalMoneyBeforeTax = TotalEUsed.Low*BusinessUnitPriceLevel1High + TotalEUsed.Normal*BusinessUnitPriceLevel2High + TotalEUsed.High*BusinessUnitPriceLevel3High
	} else {
		TotalMoneyBeforeTax = TotalEUsed.Low*BusinessUnitPriceLevel1Medium + TotalEUsed.Normal*BusinessUnitPriceLevel2Medium + TotalEUsed.High*BusinessUnitPriceLevel3Medium
	}
	return TotalMoneyBeforeTax
}

func MoneyIndustrial(TotalEUsed *dto.ListEUsedResponse, days float64) float64 {
	var TotalMoneyBeforeTax = 0.0
	if TotalEUsed.Total < constants.UnitIndustrialLowLevel1.Quantity/day*days {
		TotalMoneyBeforeTax = TotalEUsed.Low*IndustrialUnitPriceLevel1Low + TotalEUsed.Normal*IndustrialUnitPriceLevel2Low + TotalEUsed.High*IndustrialUnitPriceLevel3Low
	} else if TotalEUsed.Total < constants.UnitIndustrialMediumLevel1.Quantity/day*days {
		TotalMoneyBeforeTax = TotalEUsed.Low*IndustrialUnitPriceLevel1Medium + TotalEUsed.Normal*IndustrialUnitPriceLevel2Medium + TotalEUsed.High*IndustrialUnitPriceLevel3Medium
	} else if TotalEUsed.Total < constants.UnitIndustrialHighLevel1.Quantity/day*days {
		TotalMoneyBeforeTax = TotalEUsed.Low*IndustrialUnitPriceLevel1High + TotalEUsed.Normal*IndustrialUnitPriceLevel2High + TotalEUsed.High*IndustrialUnitPriceLevel3High
	} else {
		TotalMoneyBeforeTax = TotalEUsed.Low*IndustrialUnitPriceLevel1Max + TotalEUsed.Normal*IndustrialUnitPriceLevel2Max + TotalEUsed.High*IndustrialUnitPriceLevel3Max
	}

	return TotalMoneyBeforeTax
}

func MoneyAdministrative(TotalEUsed float64, days float64) float64 {
	var TotalMoneyBeforeTax = 0.0
	if TotalEUsed < constants.UnitAdministrativeHighLevel1.Quantity/day*days {
		TotalMoneyBeforeTax = TotalEUsed * AdministrativeUnitPriceLevel1Low
	} else {
		TotalMoneyBeforeTax = TotalEUsed * AdministrativeUnitPriceLevel1High
	}
	return TotalMoneyBeforeTax
}

func StringSQL(hours []constants.TimeRange) (string, []interface{}) {
	var orConditions []string
	var args []interface{}
	for _, tr := range hours {
		orConditions = append(orConditions, "(start_hour >= ? AND end_hour <= ?)")
		args = append(args, tr.Start, tr.End)
	}

	orSQL := "(" + strings.Join(orConditions, " OR ") + ")"
	return orSQL, args
}

func CheckDate(StartDate string, EndDate string) (time.Time, time.Time, error) {
	StartD, errS := time.Parse(layout, StartDate)
	EndD, errD := time.Parse(layout, EndDate)
	if errS != nil || errD != nil {
		return time.Time{}, time.Time{}, dto.ErrDataFormatWrong
	}
	if StartD.After(EndD) {
		return time.Time{}, time.Time{}, dto.ErrStartAfterEnd
	}
	return StartD, EndD, nil
}
