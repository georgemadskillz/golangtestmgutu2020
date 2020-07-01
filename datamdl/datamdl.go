package datamdl

// Flight stores one row of flights table
type Flight struct {
	TimeFrom   string
	FlightFrom string
	FlightTo   string
	TimeTo     string
}

// Airport stores one row of airports table
type Airport struct {
	AirID   string
	AirCity string
	AirName string
}

// Price stores one row from prices table
type Price struct {
	FlightID string
	Seat     string
	Price    string
}

// FlyDb is a common type of database
type FlyDb struct {
	flyTbl   []Flight
	flyIndex int
	airTbl   []Airport
	airIndex int
	prcTbl   []Price
	prcIndex int

	maxRows int
}

// AppendFlight is
func (db *FlyDb) AppendFlight(f Flight) {
	db.flyTbl[db.flyIndex] = Flight(f)
	db.flyIndex++
	if db.flyIndex >= db.maxRows { // For now it is circular rewriting, maybe need to do error when DB overflow??
		db.flyIndex = 0
	}
}

// GetFlight is
func (db *FlyDb) GetFlight(index int) Flight {
	if index >= db.maxRows {
		return Flight{}
	}

	return db.flyTbl[index]
}

// Init initializes FlyDb
func (db *FlyDb) Init(maxRows int) {
	db.maxRows = maxRows
	db.flyTbl = make([]Flight, maxRows, maxRows)
	db.airTbl = make([]Airport, maxRows, maxRows)
	db.prcTbl = make([]Price, maxRows, maxRows)
}
