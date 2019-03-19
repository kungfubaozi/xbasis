package gs_commons_config

type GosionConfiguration struct {
	AccessTokenExpiredTime   int64  //def: 10min
	RefreshTokenExpiredTime  int64  //def: 7day
	EmailValidateExpiredTime int64  //def: 10min
	PhoneValidateExpiredTime int64  //def: 10min
	EmailValidateTemplate    string //no def
	PhoneValidateTemplate    string //no def
}
