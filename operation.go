package s5cmd

import "fmt"

type Operation int

const (
	OP_ABORT Operation = iota
	OP_DOWNLOAD
	OP_UPLOAD
	OP_COPY
	OP_MOVE
	OP_DELETE
	OP_LOCAL_COPY
	OP_LOCAL_MOVE
	OP_LOCAL_DELETE
	OP_SHELL_EXEC
)

type ParamType int

const (
	PARAM_UNCHECKED ParamType = iota
	PARAM_UNCHECKED_ONE_OR_MORE
	PARAM_S3OBJ
	PARAM_S3OBJORDIR
	PARAM_FILEOBJ
	PARAM_FILEORDIR
)

type commandMap struct {
	keyword   string
	operation Operation
	params    []ParamType
}

var commands = []commandMap{
	{"exit", OP_ABORT, []ParamType{}},
	{"exit", OP_ABORT, []ParamType{PARAM_UNCHECKED}},
	{"get", OP_DOWNLOAD, []ParamType{PARAM_S3OBJ}},
	{"get", OP_DOWNLOAD, []ParamType{PARAM_S3OBJ, PARAM_FILEORDIR}},
	{"put", OP_UPLOAD, []ParamType{PARAM_FILEOBJ, PARAM_S3OBJORDIR}},
	{"cp", OP_COPY, []ParamType{PARAM_S3OBJ, PARAM_S3OBJORDIR}},
	{"mv", OP_MOVE, []ParamType{PARAM_S3OBJ, PARAM_S3OBJORDIR}},
	{"del", OP_DELETE, []ParamType{PARAM_S3OBJ}},
	{"local-cp", OP_LOCAL_COPY, []ParamType{PARAM_FILEOBJ, PARAM_FILEORDIR}},
	{"local-mv", OP_LOCAL_MOVE, []ParamType{PARAM_FILEOBJ, PARAM_FILEORDIR}},
	{"local-rm", OP_LOCAL_DELETE, []ParamType{PARAM_FILEOBJ}},
	{"exec", OP_SHELL_EXEC, []ParamType{PARAM_UNCHECKED_ONE_OR_MORE}},
}

func (o Operation) String() string {
	switch o {
	case OP_ABORT:
		return "abort"
	case OP_DOWNLOAD:
		return "download"
	case OP_UPLOAD:
		return "upload"
	case OP_COPY:
		return "copy"
	case OP_MOVE:
		return "move"
	case OP_DELETE:
		return "delete"
	case OP_LOCAL_COPY:
		return "local-copy"
	case OP_LOCAL_MOVE:
		return "local-move"
	case OP_LOCAL_DELETE:
		return "local-delete"
	case OP_SHELL_EXEC:
		return "shell-exec"
	}

	return fmt.Sprintf("Unknown:%d", o)
}