package consts

const (
	TplDependencyService           = "dependencyService(%+q)"
	TplDependencyTag               = "dependencyTag(%+q)"
	TplDependencyValue             = "dependencyValue(%s)"
	TplDependencyProvider          = "dependencyProvider(%s)"
	TplDependencyConcatenateChunks = "dependencyProvider(func () (string, error) { return concatenateChunks(%s) })"

	TplTokenGetParam = "func() (interface{}, error) { return getParam(%+q) }"
	TplTokenProvider = "func() (r interface{}, err error) { %s }"

	BuiltInGetEnv    = "getEnv"
	BuiltinGetEnvInt = "getEnvInt"
	BuiltInParamTodo = "paramTodo"
)
