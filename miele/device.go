package miele

const (
	DEVICE_STATUS_OFF                         = 1
	DEVICE_STATUS_ON                          = 2
	DEVICE_STATUS_PROGRAMMED                  = 3
	DEVICE_STATUS_PROGRAMMED_WAITING_TO_START = 4
	DEVICE_STATUS_RUNNING                     = 5
	DEVICE_STATUS_PAUSE                       = 6
	DEVICE_STATUS_END_PROGRAMMED              = 7
	DEVICE_STATUS_FAILURE                     = 8
	DEVICE_STATUS_PROGRAMME_INTERRUPTED       = 9
	DEVICE_STATUS_IDLE                        = 10
	DEVICE_STATUS_RINSE_HOLD                  = 11
	DEVICE_STATUS_SERVICE                     = 12
	DEVICE_STATUS_SUPERFREEZING               = 13
	DEVICE_STATUS_SUPERCOOLING                = 14
	DEVICE_STATUS_SUPERHEATING                = 15
	DEVICE_STATUS_SUPERCOOLING_SUPERFREEZING  = 146
	DEVICE_STATUS_NOT_CONNECTED               = 255
)

const (
	ACTION_START               = 1
	ACTION_STOP                = 2
	ACTION_PAUSE               = 3
	ACTION_START_SUPERFREEZING = 4
	STOP_SUPERFREEZING         = 5
	START_SUPERCOOLING         = 6
	STOP_SUPERCOOLING          = 7
)

const (
	LIGHT_ENABLE  = 1
	LIGHT_DISABLE = 2
)

type LocalizedValue struct {
	ValueRaw       int    `json:"value_raw"`
	ValueLocalized string `json:"value_localized"`
	KeyLocalized   string `json:"key_localized"`
}

type DeviceIdentLabel struct {
	FabNumber string `json:"fabNumber"`
	FabIndex  string `json:"fabIndex"`
	TechType  string `json:"techType"`
	MatNumber string `json:"matNumber"`
}

type Ident struct {
	Typ              LocalizedValue   `json:"type"`
	DeviceName       string           `json:"deviceName"`
	DeviceIdentLabel DeviceIdentLabel `json:"deviceIdentLabel"`
	XkmIdentLabel    XkmIdentLabel    `json:"xkmIdentLabel"`
}

type XkmIdentLabel struct {
	TechType       string `json:"techType"`
	ReleaseVersion string `json:"releaseVersion"`
}

type State struct {
	Status        LocalizedValue `json:"status"`
	ProgramType   LocalizedValue `json:"programType"`
	ProgramPhase  LocalizedValue `json:"programPhase"`
	SignalInfo    bool           `json:"signalInfo"`
	SignalFailure bool           `json:"signalFailure"`
	SignalDoor    bool           `json:"signalDoor"`
	RemoteEnable  struct {
		FullRemoteControl bool `json:"fullRemoteControl"`
		SmartGrid         bool `json:"smartGrid"`
	} `json:"remoteEnable"`
}

type Device struct {
	Ident Ident `json:"ident"`
	State State `json:"state"`
}
