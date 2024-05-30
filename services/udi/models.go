// Code is generated by ucloud-model, DO NOT EDIT IT.

package udi

/*
FunctionTemplate -
*/
type FunctionTemplate struct {

	//
	FunctionName string

	//
	Id string

	//
	Name string
}

/*
ParamCustom -
*/
type ParamCustom struct {

	//
	Max int

	//
	Min int

	//
	ParamName string

	//
	ParamType string
}

/*
ParamOption -
*/
type ParamOption struct {

	//
	DisplayName string

	//
	OptionalValues []string

	//
	ParamName string

	//
	Required bool

	//
	WhenValueCustom []ParamCustom
}

/*
Function -
*/
type Function struct {

	//
	DisplayName string

	//
	FunctionName string

	//
	InputType string

	//
	OutputType string

	//
	SupportParams []ParamOption
}

/*
MediaTask -
*/
type MediaTask struct {

	//
	Function string

	//
	Id string

	//
	Status string
}