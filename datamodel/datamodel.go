package datamodel

// Flight stores one row of flights table
type Flight struct {
	timeFrom   string
	flightFrom string
	flightTo   string
	timeTo     string
}

// Airport stores one row of airports table
type Airport struct {
	airID   string
	airCity string
	airName string
}

// Price stores one row from prices table
type Price struct {
	flightID string
	seat     string
	price    string
}

// FlyDb is a common type of database
type FlyDb struct {
	flyTbl  []Flight
	airTbl  []Airport
	prcTbl  []Price
	maxRows int
}

// Init initializes FlyDb
func (db *FlyDb) Init(maxRows int) {
	flyTbl = make([]Flight, maxRows)
	airTbl = make([]Airport, maxRows)
	prcTbl = make([]Price, maxRows)
}
