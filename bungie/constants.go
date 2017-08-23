package bungie

// Constant API endpoints
const (
	//GetCurrentAccountEndpoint = "http://localhost:8000/account.json"
	//ItemsEndpointFormat       = "http://localhost:8000/%d/%s/items.json"
	GetCurrentAccountEndpoint         = "https://www.bungie.net/Platform/User/GetCurrentBungieAccount/"
	ItemsEndpointFormat               = "https://www.bungie.net/d1/Platform/Destiny/%d/Account/%s/Items"
	MembershipIDFromDisplayNameFormat = "https://www.bungie.net/d1/Platform/Destiny/SearchDestinyPlayer/%d/%s/"
	TransferItemEndpointURL           = "https://www.bungie.net/d1/Platform/Destiny/TransferItem/"
	TrialsCurrentEndpoint             = "https://api.destinytrialsreport.com/currentMap"
)

const (
	UnknownTier  = 0
	CurrencyTier = 1
	BasicTier    = 2
	CommonTier   = 3
	RareTier     = 4
	SuperiorTier = 5
	ExoticTier   = 6
)

// Hash values for different class types 'classHash' JSON key
const (
	WARLOCK = 2271682572
	TITAN   = 3655393761
	HUNTER  = 671679327
)

var classHashToName = map[uint]string{
	WARLOCK: "warlock",
	TITAN:   "titan",
	HUNTER:  "hunter",
}

var classNameToHash = map[string]uint{
	"warlock": WARLOCK,
	"titan":   TITAN,
	"hunter":  HUNTER,
}

// Class Enum value passed in some of the Destiny API responses
const (
	TitanEnum        = 0
	HunterEnum       = 1
	WarlockEnum      = 2
	UnknownClassEnum = 3
)

// Hash values for Race types 'raceHash' JSON key
const (
	AWOKEN = 2803282938
	HUMAN  = 3887404748
	EXO    = 898834093
)

// Hash values for Gender 'genderHash' JSON key
const (
	MALE   = 3111576190
	FEMALE = 2204441813
)

// Gender Enum values used in some of the Bungie API responses
const (
	MaleEnum          = 0
	FemaleEnum        = 1
	UnknownGenderEnum = 2
)

// Membership type constant values
const (
	XBOX = 1
	PSN  = 2
)

// Alexa doesn't understand some of the dsetiny items or splits them into separate words
// This will allow us to translate to the correct name before doing the lookup.
var commonAlexaItemTranslations = map[string]string{
	"spin metal":     "spinmetal",
	"spin mental":    "spinmetal",
	"passage coins":  "passage coin",
	"strange coins":  "strange coin",
	"exotic shards":  "exotic shard",
	"worm spore":     "wormspore",
	"3 of coins":     "three of coins",
	"worms for":      "wormspore",
	"worm for":       "wormspore",
	"motes":          "mote of light",
	"motes of light": "mote of light",
	"spin middle":    "spinmetal",
}

var commonAlexaClassNameTrnaslations = map[string]string{
	"fault": "vault",
	"tatum": "titan",
}
