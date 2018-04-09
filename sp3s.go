package broadlink

type Sp3sDevice struct {
	*BaseDevice
}

func newSp3s(dev *BaseDevice) *Sp3sDevice {
	return &Sp3sDevice{
		BaseDevice: dev,
	}
}

func (rm *Sp3sDevice) Check() {

}

func (rm *Sp3sDevice) Send(data []byte) {

}

func (rm *Sp3sDevice) EnterLearning() {

}
