package pushover

// Reference https://pushover.net/api#sounds
const (
	// SoundPushover is the default sound
	SoundPushover = "pushover"

	SoundBike         = "bike"
	SoundBugle        = "bugle"
	SoundCashRegister = "cashregister"
	SoundClassical    = "classical"
	SoundCosmic       = "cosmic"
	SoundFalling      = "falling"
	SoundGamelan      = "gamelan"
	SoundIncoming     = "incoming"
	SoundIntermission = "intermission"
	SoundMagic        = "magic"
	SoundMechanical   = "mechanical"
	SoundPianoBar     = "pianobar"
	SoundSiren        = "siren"
	SoundSpaceAlarm   = "spacealarm"
	SoundTugBoat      = "tugboat"
	SoundAlien        = "alien"
	SoundClimb        = "climb"
	SoundPersistent   = "persistent"
	SoundPushoverEcho = "echo"
	SoundUpDown       = "updown"
	SoundVibrate      = "vibrate"
	SoundNone         = "none"
)

// https://pushover.net/api#priority
const (
	PriorityLowest    = Priority(-2)
	PriorityLow       = Priority(-1)
	PriorityNormal    = Priority(0)
	PriorityHigh      = Priority(1)
	PriorityEmergency = Priority(2)
)
