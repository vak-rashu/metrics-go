package metrics

func readProcDiskstat() {
	path := procPath("diskstats")
	reader, err := openPath(path)

}
