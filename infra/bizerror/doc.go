/*

bizerror 包装了业务上捕捉到的错误/异常。

必须的参数包括 Code 和 Msg，同时也可以提供 **错误详情** 及 **底层 error**。

+ Code 和 Msg 将通过应用层协议(HTTP/gRPC/...)返回给客户端。
    + HTTP Code 为 200, Code/Msg 在返回的 Response Body 的顶层 ErrCode/ErrMsg
    + gRPC Code 为 OK, Code/Msg 在 ErrorResponse 的 1/2 位置(见 error.proto)
+ 底层 error 有以下用途
    + 提供更多信息供开发排查问题，即用 github.com/pkg/errors.Wrap 包装
    + 通过追加 error 类型，提供信息来修改框架的行为，如 Silence 可以忽略 Prometheus Metric

Example:
    New(1000, "服务异常")
    Newf(1000, "服务异常", "更多信息")
    Newf(1000, "服务异常", "更多信息 UserID: %d", userID)
    Wrap(1000, "服务异常", sql.ErrNoRows)
    Wrapf(1000, "服务异常", sql.ErrNoRwos, "UserId: %d", userID)

    // Silence 在 error 上追加了 Silence 信息，通常在 Interceptor/Middleware 中使用
    err := Wrap(1000, "服务异常", Silence(sql.ErrNoRows))
    IsSilence(err) // true

在业务层面可以自行提供一层 Wrap，比如

Example:
    ErrorMap = map[int]error {
        1: errors.New("服务异常"),
        2: Silence(errors.New("服务异常"),
    }

    func New(code int)
    func Newf(code int, format string, args ...interface{})
    func Wrap(code int, err error)
    func Wrapf(code int, err error, format string, args ...interface{})

推荐阅读

+ https://dave.cheney.net/practical-go/presentations/qcon-china.html#_error_handling
+ https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

*/
package bizerror
