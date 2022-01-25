package ismgExporter

type Port struct {
	Value int
}

type PortError struct {
	port    int
	message string
}

func newPortError(port int) *PortError {
	return &PortError{port, "Not valid port value"}
}

func (err *PortError) Error() string {
	return err.message
}

func (toTest *Port) isValid() error {
	if toTest.Value < 1000 || toTest.Value > 65535 {
		return newPortError(toTest.Value)
	}
	return nil
}

func (Port) Create(port int) (*Port, error) {
	p := &Port{port}
	err := p.isValid()

	if err != nil {
		return nil, err
	}

	return p, nil
}
