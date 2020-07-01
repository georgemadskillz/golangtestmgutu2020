package ioctrl

import (
	"bufio"
	"flydb/cui"
	"flydb/datamdl"
	"os"
	"strings"
)

// FlyDbIO is a common type for I/O actions
type FlyDbIO struct {
	db        datamdl.FlyDb
	flyHandle FileHandler
	airHandle FileHandler
	prcHandle FileHandler
	CuiPtr    *cui.UIctrl
}

// FileHandler is
type FileHandler struct {
	path      string
	available bool
}

// Init initializes io fr FlyDB
func (io *FlyDbIO) Init(flyPath, airPath, prcPath string) {
	io.db.Init(100)

	io.flyHandle.path = "database/flights.fdb"
	io.airHandle.path = "database/airports.fdb"
	io.prcHandle.path = "database/prices.fdb"

	io.LoadFlyTable()
}

// GetRange is
func (io *FlyDbIO) GetRange(fromIndex, toIndex int) []datamdl.Flight {
	//r := toIndex - fromIndex
	// if range < 0 ????

	flights := make([]datamdl.Flight, 0)

	for i := fromIndex; i <= toIndex; i++ {
		flights = append(flights, io.db.GetFlight(i))
	}

	return flights
}

// LoadFlyTable is
func (io *FlyDbIO) LoadFlyTable() error {
	file, err := os.Open(io.flyHandle.path)
	defer file.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		io.CuiPtr.DebugPrintln("Got line: %v", line)
		flight := parseFdbRow(line)
		io.db.AppendFlight(flight)
	}

	return nil
}

func parseFdbRow(line string) datamdl.Flight {
	tokens := strings.Split(line, ",")

	flight := datamdl.Flight{}
	flight.TimeFrom = tokens[0]
	flight.FlightFrom = tokens[1]
	flight.FlightTo = tokens[2]
	flight.TimeTo = tokens[3]

	return flight
}
