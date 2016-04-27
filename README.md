# gospinner
A simple command line spinner Go library.

### Installation
    go get github.com/fredyw/gospinner

### Example
	spinner := gospinner.NewSpinner()
	spinner.Start(func(start, stop chan bool) {
		fmt.Println("Processing...")
		start <- true
		time.Sleep(3 * time.Second)
		stop <- true
	})
	fmt.Println("Done")
