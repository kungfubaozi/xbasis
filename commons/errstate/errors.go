package errstate

import commons "konekko.me/xbasis/commons/dto"

var (
	Success                          = &commons.State{Ok: true, Code: 0, Message: "success"}
	SuccessTraceCheck                = &commons.State{Ok: true, Code: 1001, Message: "success"}
	FinishedInvited                  = &commons.State{Ok: true, Code: 1002, Message: "success"}
	ErrRequest                       = &commons.State{Ok: false, Code: 20001, Message: "err request"}
	ErrInvalidStateField             = &commons.State{Ok: false, Code: 20002, Message: "err return type, no found 'State' file"}
	ErrRoleAlreadyExists             = &commons.State{Ok: false, Code: 20003, Message: "err role already exists"}
	ErrGroupAlreadyExists            = &commons.State{Ok: false, Code: 20004, Message: "err group already exists"}
	ErrBlacklistAlreadyExists        = &commons.State{Ok: false, Code: 20005, Message: "err blacklist already exists"}
	ErrApiAlreadyExists              = &commons.State{Ok: false, Code: 20006, Message: "err api already exists"}
	ErrFunctionBindGroupId           = &commons.State{Ok: false, Code: 20007, Message: "err no found function groupId"}
	ErrFunctionAuthType              = &commons.State{Ok: false, Code: 20008, Message: "err function auth type"}
	ErrApplicationAlreadyExists      = &commons.State{Ok: false, Code: 20009, Message: "err application already exists"}
	ErrApplicationClosed             = &commons.State{Ok: false, Code: 20010, Message: "err application closed"}
	ErrSystem                        = &commons.State{Ok: false, Code: 20011, Message: "err system"}
	ErrDurationAccessExpired         = &commons.State{Ok: false, Code: 20012, Message: "err function access key"}
	ErrDurationAccess                = &commons.State{Ok: false, Code: 20013, Message: "err function access"}
	ErrAuthorization                 = &commons.State{Ok: false, Code: 20014, Message: "err authorization"}
	ErrDurationAccessTokenBusy       = &commons.State{Ok: false, Code: 20015, Message: "err get function access key busy"}
	ErrNotFoundDurationAccessToken   = &commons.State{Ok: false, Code: 20016, Message: "err not found duration access key"}
	ErrVerificationCode              = &commons.State{Ok: false, Code: 20017, Message: "err verification code"}
	ErrInvalidUsernameOrPassword     = &commons.State{Ok: false, Code: 20018, Message: "err invalid username or password"}
	ErrLoginFailed                   = &commons.State{Ok: false, Code: 20019, Message: "err login failed"}
	ErrRefreshToken                  = &commons.State{Ok: false, Code: 20020, Message: "err refresh token"}
	ErrAccessToken                   = &commons.State{Ok: false, Code: 20021, Message: "err access token"}
	ErrAccessTokenExpired            = &commons.State{Ok: false, Code: 20022, Message: "err access token expired"}
	ErrRefreshTokenExpired           = &commons.State{Ok: false, Code: 20023, Message: "err refresh token expired"}
	ErrOperate                       = &commons.State{Ok: false, Code: 20024, Message: "err operate"}
	ErrClientClosed                  = &commons.State{Ok: false, Code: 20025, Message: "err client closed"}
	ErrAccessTokenOrClient           = &commons.State{Ok: false, Code: 20026, Message: "err access token or client"}
	ErrUserPermission                = &commons.State{Ok: false, Code: 20027, Message: "err access function without permission"}
	ErrInvalidClientId               = &commons.State{Ok: false, Code: 20030, Message: "err invalid client"}
	ErrUserAppPermission             = &commons.State{Ok: false, Code: 20031, Message: "err access application without permission"}
	ErrRoutePlatform                 = &commons.State{Ok: false, Code: 20032, Message: "err route platform"}
	ErrRouteNotMainClient            = &commons.State{Ok: false, Code: 20033, Message: "err not main client"}
	ErrRouteSameApplication          = &commons.State{Ok: false, Code: 20034, Message: "err route same application"}
	ErrOperateBusy                   = &commons.State{Ok: false, Code: 20035, Message: "err operate busy"}
	ErrUserSyncToApp                 = &commons.State{Ok: false, Code: 20035, Message: "err operate busy"}
	ErrApplicationRedirectUrl        = &commons.State{Ok: false, Code: 20036, Message: "err app url"}
	ErrApplicationSyncUrl            = &commons.State{Ok: false, Code: 20037, Message: "err sync url"}
	ErrDurationAccessCredential      = &commons.State{Ok: false, Code: 20038, Message: "err duration access credential"}
	ErrDurationAccessTarget          = &commons.State{Ok: false, Code: 20039, Message: "err duration access target"}
	ErrFormatEmail                   = &commons.State{Ok: false, Code: 20040, Message: "err format email"}
	ErrFormatPhone                   = &commons.State{Ok: false, Code: 20041, Message: "err format phone"}
	ErrDurationAccessUnsentCode      = &commons.State{Ok: false, Code: 20042, Message: "err duration access unsent code"} //
	ErrDurationAccessCode            = &commons.State{Ok: false, Code: 20043, Message: "err duration access code"}        //
	ErrUnbindEmail                   = &commons.State{Ok: false, Code: 20044, Message: "err unbind email"}                //
	ErrUnbindPhone                   = &commons.State{Ok: false, Code: 20045, Message: "err unbind phone"}
	ErrAlreadyLocking                = &commons.State{Ok: false, Code: 20046, Message: "err user already locking"}
	ErrNotFound                      = &commons.State{Ok: false, Code: 20047, Message: "err not found"}
	ErrUserAlreadyBindRole           = &commons.State{Ok: false, Code: 20048, Message: "err user already bind role"}
	ErrFunctionAlreadyBindRole       = &commons.State{Ok: false, Code: 20049, Message: "err function already bind role"}
	ErrOAuthTypeNotFound             = &commons.State{Ok: false, Code: 20050, Message: "err oauth type not found"}
	ErrUserAlreadyRegister           = &commons.State{Ok: false, Code: 20051, Message: "err user already register"}
	ErrValidationCode                = &commons.State{Ok: false, Code: 20052, Message: "err validation code"}
	ErrPasswordLength                = &commons.State{Ok: false, Code: 20053, Message: "err password length"}
	ErrUserNotAuthorize              = &commons.State{Ok: false, Code: 20055, Message: "err user not authorize"}
	ErrHasInvited                    = &commons.State{Ok: false, Code: 20056, Message: "err has invited"}
	ErrUserNeedSetUsername           = &commons.State{Ok: false, Code: 20057, Message: "err need username"}
	ErrInvalidServiceNode            = &commons.State{Ok: false, Code: 20058, Message: "err No valid node was found"}          //未发现可用的服务节点
	ErrInvalidApplicationServiceName = &commons.State{Ok: false, Code: 20059, Message: "err invalid application service name"} //应用服务名称未设置
	ErrServiceRequestTimeout         = &commons.State{Ok: false, Code: 20059, Message: "err service request timeout"}          //请求超时
)
