package broadlink

type CommandCode uint16

type Command interface {
	Code() CommandCode
	Bytes() ([]byte, error)
}

type baseCommand struct {
	code CommandCode
}

func (bc *baseCommand) Code() CommandCode {
	return bc.code
}

type commandAuth struct {
	baseCommand
}

func NewAuthCommand() Command {
	return &commandAuth{
		baseCommand{0x65},
	}
}

func (ac *commandAuth) Bytes() ([]byte, error) {
	ba := [0x50]byte{0}

	ba[0x04] = 0x31
	ba[0x05] = 0x31
	ba[0x06] = 0x31
	ba[0x07] = 0x31
	ba[0x08] = 0x31
	ba[0x09] = 0x31
	ba[0x0a] = 0x31
	ba[0x0b] = 0x31
	ba[0x0c] = 0x31
	ba[0x0d] = 0x31
	ba[0x0e] = 0x31
	ba[0x0f] = 0x31
	ba[0x10] = 0x31
	ba[0x11] = 0x31
	ba[0x12] = 0x31
	ba[0x1e] = 0x01
	ba[0x2d] = 0x01
	ba[0x30] = 'T'
	ba[0x31] = 'e'
	ba[0x32] = 's'
	ba[0x33] = 't'
	ba[0x34] = ' '
	ba[0x35] = ' '
	ba[0x36] = '1'

	return ba[:], nil
}

// Enter Learning

type commandEnterLearning struct {
	baseCommand
}

func NewEnterLearningCommand() Command {
	return &commandEnterLearning{
		baseCommand{0x006a},
	}
}

func (ac *commandEnterLearning) Bytes() ([]byte, error) {
	ba := [0x10]byte{0}

	ba[0x00] = 0x03

	return ba[:], nil
}

// Check Data

type commandCheckData struct {
	baseCommand
}

func NewCheckDataCommand() Command {
	return &commandCheckData{
		baseCommand{0x006a},
	}
}

func (ac *commandCheckData) Bytes() ([]byte, error) {
	ba := [0x10]byte{0}

	ba[0x00] = 0x04

	return ba[:], nil
}

// Send Data

type commandSendData struct {
	baseCommand
	data []byte
}

func NewSendDataCommand(data []byte) Command {
	return &commandSendData{
		baseCommand{0x006a},
		data,
	}
}

func (ac *commandSendData) Bytes() ([]byte, error) {
	ba := [0x04]byte{0x02, 0x00, 0x00, 0x00}

	bs := append(ba[:], ac.data...)

	return bs, nil
}

// Check Power

type commandCheckPower struct {
	baseCommand
}

func NewCheckPowerCommand() Command {
	return &commandCheckPower{
		baseCommand{0x006a},
	}
}

func (ac *commandCheckPower) Bytes() ([]byte, error) {
	ba := [0x10]byte{0}

	ba[0x00] = 0x01

	return ba[:], nil
}

// Get Energy

type commandGetEnergy struct {
	baseCommand
}

func NewCommandGetEnergy() Command {
	return &commandCheckPower{
		baseCommand{0x6a},
	}
}

func (ac *commandGetEnergy) Bytes() ([]byte, error) {
	ba := [0x10]byte{8, 0, 254, 1, 5, 1, 0, 0, 0, 45}

	return ba[:], nil
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

//Enter Learning
//def enter_learning(self):
//packet = bytearray(16)
//packet[0] = 3
//self.send_packet(0x6a, packet)

// Check Power
//def check_power(self):
//"""Returns the power state of the smart plug."""
//packet = bytearray(16)
//packet[0] = 1
