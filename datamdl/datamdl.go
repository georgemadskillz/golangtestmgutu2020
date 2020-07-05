package datamdl

import "fmt"

const (
	// FdbFly is a special fdb key for Flight table
	FdbFly = iota
	// FdbAir is a special fdb key for Airport table
	FdbAir = iota
	// FdbPrc is a special fdb key for Price table
	FdbPrc = iota
)

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

type flyTbl struct {
	tbl      []Flight
	curIndex int
	maxIndex int
}

func (t *flyTbl) init(maxRows int) error {
	if maxRows < 1 {
		return fmt.Errorf("invalid maxRows=%v! flyTbl hasn't been initialized", maxRows)
	}

	t.maxIndex = maxRows
	t.tbl = make([]Flight, t.maxIndex, t.maxIndex)
	t.curIndex = -1
	return nil
}

func (t *flyTbl) appendRow(row Flight) error {
	inx := t.curIndex
	inx++
	if inx >= t.maxIndex {
		return fmt.Errorf("cannot append new row, flights table overflow at index=%v", inx)
	}
	t.curIndex = inx
	t.tbl[inx] = Flight(row)
	return nil
}

func (t *flyTbl) getRow(inx int) (Flight, error) {
	if inx >= t.maxIndex {
		return Flight{}, fmt.Errorf("cannot get row, at index=%v > max index=%v", inx, t.maxIndex)
	}

	return t.tbl[inx], nil
}

type airTbl struct {
	tbl      []Airport
	curIndex int
	maxIndex int
}

func (t *airTbl) init(maxRows int) error {
	if maxRows < 1 {
		return fmt.Errorf("invalid maxRows=%v! airTbl hasn't been initialized", maxRows)
	}

	t.maxIndex = maxRows
	t.tbl = make([]Airport, t.maxIndex, t.maxIndex)
	t.curIndex = -1
	return nil
}

func (t *airTbl) appendRow(row Airport) error {
	inx := t.curIndex
	inx++
	if inx >= t.maxIndex {
		return fmt.Errorf("cannot append new row, airports table overflow at index=%v", inx)
	}
	t.curIndex = inx
	t.tbl[inx] = Airport(row)
	return nil
}

func (t *airTbl) getRow(inx int) (Airport, error) {
	if inx >= t.maxIndex {
		return Airport{}, fmt.Errorf("cannot get row, at index=%v > max index=%v", inx, t.maxIndex)
	}

	return t.tbl[inx], nil
}

type prcTbl struct {
	tbl      []Price
	curIndex int
	maxIndex int
}

func (t *prcTbl) init(maxRows int) error {
	if maxRows < 1 {
		return fmt.Errorf("invalid maxRows=%v! prcTbl hasn't been initialized", maxRows)
	}

	t.maxIndex = maxRows
	t.tbl = make([]Price, t.maxIndex, t.maxIndex)
	t.curIndex = -1
	return nil
}

func (t *prcTbl) appendRow(row Price) error {
	inx := t.curIndex
	inx++
	if inx >= t.maxIndex {
		return fmt.Errorf("cannot append new row, prices table overflow at index=%v", inx)
	}
	t.curIndex = inx
	t.tbl[inx] = Price(row)
	return nil
}

func (t *prcTbl) getRow(inx int) (Price, error) {
	if inx >= t.maxIndex {
		return Price{}, fmt.Errorf("cannot get row, at index=%v > max index=%v", inx, t.maxIndex)
	}

	return t.tbl[inx], nil
}

// FlyDb is a common type of database
type FlyDb struct {
	flyTable flyTbl
	airTable airTbl
	prcTable prcTbl
	maxRows  int
}

// Init initializes FlyDb
func (db *FlyDb) Init(maxRows int) error {
	maxR := maxRows

	db.flyTable = flyTbl{}
	if err := db.flyTable.init(maxR); err != nil {
		return err
	}

	db.airTable = airTbl{}
	if err := db.airTable.init(maxR); err != nil {
		return err
	}

	db.prcTable = prcTbl{}
	if err := db.prcTable.init(maxR); err != nil {
		return err
	}

	db.maxRows = maxRows

	return nil
}

// AppendRow is
func (db *FlyDb) AppendRow(row interface{}) error {

	switch r := row.(type) {
	case Flight:
		return db.flyTable.appendRow(r)
	case Airport:
		return db.airTable.appendRow(r)
	case Price:
		return db.prcTable.appendRow(r)
	default:
		return fmt.Errorf("unknown data base typein row=%#v", r)
	}
}

// GetRow is
func (db *FlyDb) GetRow(key int, index int) (row interface{}, err error) {
	if index >= db.maxRows {
		return nil, fmt.Errorf("index")
	}

	switch key {
	case FdbFly:
		return db.flyTable.getRow(index)
	case FdbAir:
		return db.airTable.getRow(index)
	case FdbPrc:
		return db.prcTable.getRow(index)
	default:
		return nil, fmt.Errorf("unknown key=%#v", key)
	}
}
