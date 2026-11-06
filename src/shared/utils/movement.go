package veUtils

type MovementType string

const (
	MovementSend            MovementType = "send"
	MovementReceive         MovementType = "receive"
	MovementExternalSend    MovementType = "external_send"
	MovementExternalReceive MovementType = "external_receive"
	MovementFee             MovementType = "fee"
)

func (m MovementType) Result(initialValue, movementValue Currency) Currency {
	switch m {
		case MovementReceive, MovementExternalReceive:
			return initialValue + movementValue;

		case MovementSend, MovementExternalSend:
			return initialValue - movementValue;
	}
	panic("Invalid movement type: " + m)
}
