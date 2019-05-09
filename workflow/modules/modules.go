package modules

type Modules interface {
	History() IHistory

	Instance() IInstance

	Process() IProcesses

	Runtime() IRuntime

	Form() IForm

	User() IUser
}
