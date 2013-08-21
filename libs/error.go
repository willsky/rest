package libs

import (
    "fmt"
    "strconv"
)

func (this *Error) Error() string {
    return fmt.Sprintf("Errorcode: %d Error: %s", this.code, this.err)
}

func RuntimeError(args ... string) *Error{
    err := new(Error)
    var (
        msg string = "unknown"
        code int = 0
    )

    if len(args) > 0 {
        msg = args[0]
    }

    if len(args) > 1 {
        if errcode, err := strconv.Atoi(args[1]); nil == err {
            code = errcode
        }
    }

    err.err = msg
    err.code = code
    return err
}

