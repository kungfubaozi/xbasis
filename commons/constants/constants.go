package gs_commons_constants

const (
	PermissionService     = "gs.svc.permission"
	MessageService        = "gs.svc.message"
	ConnectionService     = "gs.svc.connection"
	UserService           = "gs.svc.service"
	ApplicationService    = "gs.svc.application"
	SafetyService         = "gs.svc.safety"
	AuthenticationService = "gs.svc.authentication"
	AuthWrapperService    = "inside.gs.svc.auth.wrapper"
	GoMicroApi            = "go.micro.api"
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

	//uncheck
	AuthTypeOfNone = 11 << 11

	//password
	AuthTypeOfPassword = 11 << 12
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
