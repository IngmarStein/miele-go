package miele

const (
	DEVICE_TYPE_WASHING_MACHINE                  = 1
	DEVICE_TYPE_TUMBLE_DRYER                     = 2
	DEVICE_TYPE_DISHWASHER                       = 7
	DEVICE_TYPE_OVEN                             = 12
	DEVICE_TYPE_OVEN_MICROWAVE                   = 13
	DEVICE_TYPE_HOB_HIGHLIGHT                    = 14
	DEVICE_TYPE_STEAM_OVEN                       = 15
	DEVICE_TYPE_MICROWAVE                        = 16
	DEVICE_TYPE_COFFEE_SYSTEM                    = 17
	DEVICE_TYPE_HOOD                             = 18
	DEVICE_TYPE_FRIDGE                           = 19
	DEVICE_TYPE_FREEZER                          = 20
	DEVICE_TYPE_FRIDGE_FREEZER_COMBINATION       = 21
	DEVICE_TYPE_VACUUM_CLEANER                   = 23
	DEVICE_TYPE_WASHER_DRYER                     = 24
	DEVICE_TYPE_DISH_WARMER                      = 25
	DEVICE_TYPE_HOB_INDUCTION                    = 27
	DEVICE_TYPE_STEAM_OVEN_COMBINATION           = 31
	DEVICE_TYPE_WINE_CABINET                     = 32
	DEVICE_TYPE_WINE_CONDITIONING_UNIT           = 33
	DEVICE_TYPE_WINE_STORAGE_CONDITIONING_UNIT   = 34
	DEVICE_TYPE_STEAM_OVEN_MICROWAVE_COMBINATION = 45
	DEVICE_TYPE_VACUUM_DRAWER                    = 48
	DEVICE_TYPE_DIALOGOVEN                       = 67
	DEVICE_TYPE_WINE_CABINET_FREEZER_COMBINATION = 68
)

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

const (
	VENTILATION_STEP1 = 1
	VENTILATION_STEP2 = 2
	VENTILATION_STEP3 = 3
	VENTILATION_STEP4 = 4
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

type ShortDevice struct {
	FabNumber  string `json:"fabNumber"`
	State      string `json:"state"`
	Type       string `json:"type"`
	DeviceName string `json:"deviceName"`
	Details    string `json:"details"`
}

type DeviceAction struct {
	ProcessAction     []int   `json:"processAction"`
	Light             []int   `json:"light"`
	StartTime         [][]int `json:"startTime"`
	VentilationStep   []int   `json:"ventilationStep"`
	ProgramId         []int   `json:"programId"`
	TargetTemperature []struct {
		Zone int `json:"zone"`
		Min  int `json:"min"`
		Max  int `json:"max"`
	} `json:"targetTemperature"`
	DeviceName bool `json:"deviceName"`
	PowerOff   bool `json:"powerOff"`
	PowerOn    bool `json:"powerOn"`
}
