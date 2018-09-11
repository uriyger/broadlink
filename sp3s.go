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

func (bd *BaseDevice) CheckPower() (bool, error) {
	command := NewCheckPowerCommand()
	response, err := bd.SendCommand(command)

	//response = self.send_packet(0x6a, packet)
	//err = response[0x22] | (response[0x23] << 8)
	//if err == 0:
	//payload = self.decrypt(bytes(response[0x38:]))

	//if ord(payload[0x4]) == 1 or ord(payload[0x4]) == 3:
	//state = True
	//else:
	//state = False
	//return state

	return int(response[0x3]) == 1 || int(response[0x3]) == 3, err
}

func (bd *Sp3sDevice) GetEnergy() (float32, error) {
	command := NewCommandGetEnergy()
	payload, err := bd.SendCommand(command)
	if err != nil {
		return 0, err
	}

	energy := float32(payload[0x07])*float32(256) + float32(payload[0x06]) + float32(payload[0x05])/100.0

	return energy, nil
}

//packet = bytearray([8, 0, 254, 1, 5, 1, 0, 0, 0, 45])
//response = self.send_packet(0x6a, packet)
//err = response[0x22] | (response[0x23] << 8)
//if err == 0:
//payload = self.decrypt(bytes(response[0x38:]))
//if type(payload[0x07]) == int:
//energy = int(hex(payload[0x07] * 256 + payload[0x06])[2:]) + int(hex(payload[0x05])[2:])/100.0
//else:
//energy = int(hex(ord(payload[0x07]) * 256 + ord(payload[0x06]))[2:]) + int(hex(ord(payload[0x05]))[2:])/100.0
//return energy
