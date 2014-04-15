package argparser

import "github.com/sublee/argparser"

func main() {

	parser := argparser.NewParser()
	parser.SetSummary("Test...")

	device, arg := parser.AddString("DEVICE", nil, "blah blah")
	arg.Require(true)
	arg.SetChoices([]string{"queue", "forwarder", "streamer"})

	long, _ := parser.AddFlag("--long", "-l", "hello, world")
	noTraffic, _ := parser.AddAntiFlag("--no-traffic", nil, nil)

	subparser := parser.AddSubparser("device")
	subparser.AddBool("--hello", nil, nil)

	parser.Parse(nil)

}
