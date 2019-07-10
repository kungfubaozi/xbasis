package xbasisconstants

const (
	PermissionService     = "xbasis.external.permission"
	GatewayService        = "xbasis.external.gateway"
	AnalysisService       = "xbasis.external.analytical"
	MessageService        = "xbasis.external.message"
	ConnectionService     = "xbasis.external.connection"
	UserService           = "xbasis.external.user"
	WorkflowService       = "xbasis.external.workflow"
	ApplicationService    = "xbasis.external.apps"
	SafetyService         = "xbasis.external.safety"
	AuthenticationService = "xbasis.external.authentication"
)

const (
	InternalPermission            = "xbasis.internal.permission" //permission
	InternalUserService           = "xbasis.internal.user"
	InternalSafetyService         = "xbasis.internal.safety"
	InternalApplicationService    = "xbasis.internal.application"
	InternalAuthenticationService = "xbasis.internal.authentication"
)

const (
	MessageChannel          = "xbasis.channel.message"
	ConnectionFanoutChannel = "xbasis.channel.connection.fanout"
)

const (
	StateOk int64 = 3 << 3
)

const (
	InviteStateOfRegister  int64 = 2 << 2
	InviteStateOfAuthorize int64 = 2 << 3
)

func GetStateString(state int64) string {
	value := ""
	switch state {
	case InviteStateOfRegister:
		value = "(等待注册)"
		break
	case InviteStateOfAuthorize:
		value = "(等待授权)"
		break
	}
	return value
}

const (
	//send email or phone validate code, next request user usage _stat
	AuthTypeOfValcode = 11 << 5

	//face
	AuthTypeOfFace = 11 << 8

	//token, check user status, blacklist, token and so on
	AuthTypeOfToken = 11 << 9

	//mobile confirm
	AuthTypeOfMobileConfirm = 11 << 15

	////weixin mini program --- Gosion
	AuthTypeOfMiniProgramCodeConfirm = 11 << 17
	//
	////weixin mini program
	AuthTypeOfMiniProgramUserConfirm = 11 << 19
)

const (
	OpenModeOfSelfOrganization = 12 << 9

	OpenModeOfAllOrganization = 12 << 11

	OpenModeOfCompletely = 12 << 13
)

const (
	//android
	PlatformOfAndroid = 13 << 11

	//ios
	PlatformOfIOS = 13 << 12

	//windows
	PlatformOfWindows = 13 << 13

	//macos
	PlatfromOfMacOS = 13 << 14

	//web
	PlatformOfWeb = 13 << 15

	PlatformOfLinux = 13 << 16

	PlatformOfFuchsia = 13 << 17
)

func GetPlatformName(v int64) string {
	switch v {
	case PlatformOfIOS:
		return "iOS"
	case PlatformOfWeb:
		return "Web"
	case PlatformOfAndroid:
		return "Android"
	case PlatformOfLinux:
		return "Linux"
	case PlatformOfFuchsia:
		return "Fuchsia"
	case PlatformOfWindows:
		return "Windows"
	case PlatfromOfMacOS:
		return "MacOS"
	default:
		return ""
	}
}

const (
	AppTypeRoute int64 = 9 << 3

	AppTypeUser int64 = 9 << 5

	AppTypeManage int64 = 9 << 7

	AppTypeSafe int64 = 9 << 9
)

const (
	Enabled int64 = 14 << 6
	Closed  int64 = 14 << 7
)

const (
	BlacklistOfIP       = 5 << 4
	BlacklistOfDevice   = 5 << 6
	BlacklistOfRegister = 5 << 7
	BlacklistOfArea     = 5 << 8
	BlacklistOfUser     = 5 << 9
)

const (
	ZKWatchInitializeConfigPath        = "/xbasis-initialize-configuration"
	ZKWatchInitializeVersionListenPath = "/xbasis-initialize-version"
	XBasisConfiguration                = "/xbasis-configuration"
	ZKAutonomyRegister                 = "/xbasis-autonomy-register"
)

const (
	DatTypeOfSendToWeixin  = 40801
	DatTypeOfSendToDefault = 40803
)

const (
	AccessToken  int64 = 8 >> 10
	RefreshToken int64 = 9 >> 9
)

const (
	ContractTypeOfPhone = 5 << 2
	ContractTypeOfEmail = 5 << 3
)

const (
	AppUserGroup          = "$app"
	AppMainStructureGroup = "$org"
)

var IndexMapping = `{
	"mappings":{
		"_doc":{
			"properties":{
				"join_field":{
					"type":"join",
					"relations":{
						"relation":"child"
					}
				}
			}
		}
	}
}`
