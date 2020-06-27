package ioctrl

// IOcontroller handling DB files write/read
func IOcontroller() {
	var io FlyDbIO

	io.Init("database/flights.fdb", "database/airports.fdb", "database/prices.fdb")

	for {

	}
}

// FlyDbIO is a common type for I/O actions
type FlyDbIO struct {
	flyFileName string
	airFileName string
	prcFilename string
}

// Init initializes io fr FlyDB
func (io *FlyDbIO) Init(flyFile, airFile, prcFile string) {

	io.flyFileName = flyFile
	io.airFileName = airFile
	io.prcFilename = prcFile

}

// func (io *FlyDbIO) ReadDb () {
// 	// // Check flights tale file
// 	// if _, err := os.Stat(flyFile); err == nil {
// 	// 	file, err := os.Open(flyFile)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	defer file.Close()

// 	// } else if os.IsNotExist(err) {
// 	// 	// path/to/whatever does *not* exist

// 	// }
// }

// func (io *FlyDbIO) WriteDb () {

// }

// // CheckFileFormat checks file format and returns number of rows
// func (io *FlyDbIO) CheckFileFormat() rowsAmnt int {

// }
