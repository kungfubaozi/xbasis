package gs_commons_constants

const (
	PermissionService     = "gs.svc.permission"
	MessageService        = "gs.svc.message"
	ConnectionService     = "gs.svc.connection"
	UserService           = "gs.svc.user"
	ApplicationService    = "gs.svc.application"
	SafetyService         = "gs.svc.safety"
	AuthenticationService = "gs.svc.authentication"
)

const (
	ExtPermissionVerification = "gs.ext.svc.permission.verification"
	ExtUserService            = "gs.ext.svc.user"
	ExtSafetyService          = "gs.ext.svc.safety"
	ExtApplicationService     = "gs.ext.svc.application"
	ExtAuthenticationService  = "gs.ext.svc.authentication"
)

const (
	MessageChannel          = "gs.channel.message"
	ConnectionFanoutChannel = "gs.channel.connection.fanout"
)

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
)

const (
	Enabled = 14 << 6
	Closed  = 14 << 7
)

const (
	UserGroupTypeOfRoot   = 7 << 5
	UserGroupTypeOfNormal = 7 << 6
)

const (
	BlacklistOfIP         = 5 << 4
	BlacklistOfUserDevice = 5 << 6
)

const (
	ZKWatchInitializeConfigPath        = "/_gosion-initialize-config"
	ZKWatchInitializeVersionListenPath = "/_gosion-initialize-version"
	GosionConfiguration                = "/_gosion-configuration"
)

const (
	DatTypeOfSendToWeixin  = 40801
	DatTypeOfSendToDefault = 40803
)

const (
	AccessToken  = 8 >> 10
	RefreshToken = 9
)
