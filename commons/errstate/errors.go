package errstate

import "konekko.me/gosion/commons/dto"

var (
	Success                        = &gs_commons_dto.State{Ok: true, Code: 0, Message: "success"}
	SuccessTraceCheck              = &gs_commons_dto.State{Ok: true, Code: 1001, Message: "success"}
	ErrRequest                     = &gs_commons_dto.State{Ok: false, Code: 20001, Message: "err request"}
	ErrInvalidStateField           = &gs_commons_dto.State{Ok: false, Code: 20002, Message: "err return type, no found 'State' file"}
	ErrRoleAlreadyExists           = &gs_commons_dto.State{Ok: false, Code: 20003, Message: "err role already exists"}
	ErrGroupAlreadyExists          = &gs_commons_dto.State{Ok: false, Code: 20004, Message: "err group already exists"}
	ErrBlacklistAlreadyExists      = &gs_commons_dto.State{Ok: false, Code: 20005, Message: "err blacklist already exists"}
	ErrApiAlreadyExists            = &gs_commons_dto.State{Ok: false, Code: 20006, Message: "err api already exists"}
	ErrFunctionBindGroupId         = &gs_commons_dto.State{Ok: false, Code: 20007, Message: "err no found function groupId"}
	ErrFunctionAuthType            = &gs_commons_dto.State{Ok: false, Code: 20008, Message: "err function auth type"}
	ErrApplicationAlreadyExists    = &gs_commons_dto.State{Ok: false, Code: 20009, Message: "err application already exists"}
	ErrApplicationClosed           = &gs_commons_dto.State{Ok: false, Code: 20010, Message: "err application closed"}
	ErrSystem                      = &gs_commons_dto.State{Ok: false, Code: 20011, Message: "err system"}
	ErrDurationAccessExpired       = &gs_commons_dto.State{Ok: false, Code: 20012, Message: "err function access key"}
	ErrDurationAccess              = &gs_commons_dto.State{Ok: false, Code: 20013, Message: "err function access"}
	ErrAuthorization               = &gs_commons_dto.State{Ok: false, Code: 20014, Message: "err authorization"}
	ErrDurationAccessTokenBusy     = &gs_commons_dto.State{Ok: false, Code: 20015, Message: "err get function access key busy"}
	ErrNotFoundDurationAccessToken = &gs_commons_dto.State{Ok: false, Code: 20016, Message: "err not found duration access key"}
	ErrVerificationCode            = &gs_commons_dto.State{Ok: false, Code: 20017, Message: "err verification code"}
	ErrInvalidUsernameOrPassword   = &gs_commons_dto.State{Ok: false, Code: 20018, Message: "err invalid username or password"}
	ErrLoginFailed                 = &gs_commons_dto.State{Ok: false, Code: 20019, Message: "err login failed"}
	ErrRefreshToken                = &gs_commons_dto.State{Ok: false, Code: 20020, Message: "err refresh token"}
	ErrAccessToken                 = &gs_commons_dto.State{Ok: false, Code: 20021, Message: "err access token"}
	ErrAccessTokenExpired          = &gs_commons_dto.State{Ok: false, Code: 20022, Message: "err access token expired"}
	ErrRefreshTokenExpired         = &gs_commons_dto.State{Ok: false, Code: 20023, Message: "err refresh token expired"}
	ErrOperate                     = &gs_commons_dto.State{Ok: false, Code: 20024, Message: "err operate"}
	ErrClientClosed                = &gs_commons_dto.State{Ok: false, Code: 20025, Message: "err client closed"}
	ErrAccessTokenOrClient         = &gs_commons_dto.State{Ok: false, Code: 20026, Message: "err access token or client"}
	ErrUserPermission              = &gs_commons_dto.State{Ok: false, Code: 20027, Message: "err access function without permission"}
	ErrInvalidClientId             = &gs_commons_dto.State{Ok: false, Code: 20030, Message: "err invalid client"}
	ErrUserAppPermission           = &gs_commons_dto.State{Ok: false, Code: 20031, Message: "err access application without permission"}
	ErrRoutePlatform               = &gs_commons_dto.State{Ok: false, Code: 20032, Message: "err route platform"}
	ErrRouteNotMainClient          = &gs_commons_dto.State{Ok: false, Code: 20033, Message: "err not main client"}
	ErrRouteSameApplication        = &gs_commons_dto.State{Ok: false, Code: 20034, Message: "err route same application"}
	ErrOperateBusy                 = &gs_commons_dto.State{Ok: false, Code: 20035, Message: "err operate busy"}
	ErrUserSyncToApp               = &gs_commons_dto.State{Ok: false, Code: 20035, Message: "err operate busy"}
	ErrApplicationRedirectUrl      = &gs_commons_dto.State{Ok: false, Code: 20036, Message: "err app url"}
	ErrApplicationSyncUrl          = &gs_commons_dto.State{Ok: false, Code: 20037, Message: "err sync url"}
	ErrDurationAccessCredential    = &gs_commons_dto.State{Ok: false, Code: 20038, Message: "err duration access credential"}
	ErrDurationAccessTarget        = &gs_commons_dto.State{Ok: false, Code: 20039, Message: "err duration access target"}
	ErrFormatEmail                 = &gs_commons_dto.State{Ok: false, Code: 20040, Message: "err format email"}
	ErrFormatPhone                 = &gs_commons_dto.State{Ok: false, Code: 20041, Message: "err format phone"}
	ErrDurationAccessUnsentCode    = &gs_commons_dto.State{Ok: false, Code: 20042, Message: "err duration access unsent code"} //
	ErrDurationAccessCode          = &gs_commons_dto.State{Ok: false, Code: 20043, Message: "err duration access code"}        //
	ErrUnbindEmail                 = &gs_commons_dto.State{Ok: false, Code: 20044, Message: "err unbind email"}                //
	ErrUnbindPhone                 = &gs_commons_dto.State{Ok: false, Code: 20045, Message: "err unbind phone"}
	ErrAlreadyLocking              = &gs_commons_dto.State{Ok: false, Code: 20046, Message: "err user already locking"}
	ErrNotFound                    = &gs_commons_dto.State{Ok: false, Code: 20047, Message: "err not found"}
	ErrUserAlreadyBindRole         = &gs_commons_dto.State{Ok: false, Code: 20048, Message: "err user already bind role"}
	ErrFunctionAlreadyBindRole     = &gs_commons_dto.State{Ok: false, Code: 20049, Message: "err function already bind role"}
	ErrOAuthTypeNotFound           = &gs_commons_dto.State{Ok: false, Code: 20050, Message: "err oauth type not found"}
	ErrUserAlreadyRegister         = &gs_commons_dto.State{Ok: false, Code: 20051, Message: "err user already register"}
	ErrValidationCode              = &gs_commons_dto.State{Ok: false, Code: 20052, Message: "err validation code"}
	ErrPasswordLength              = &gs_commons_dto.State{Ok: false, Code: 20053, Message: "err password length"}
	ErrUserNotSync                 = &gs_commons_dto.State{Ok: false, Code: 20054, Message: "err user not sync"}
)
