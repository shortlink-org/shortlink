package logger

// Test input/output
// -> log.Info("Run HTTP-CHI API")
// -> {"level":"info","msg":"Run HTTP-CHI API","time":"2019-11-02T14:29:24+03:00"}

// Test setlevel
// setLevel error
// -> log.Info("Run HTTP-CHI API")
// -> ...
// setLevel info
// -> log.Info("Run HTTP-CHI API")
// -> {"level":"info","msg":"Run HTTP-CHI API","time":"2019-11-02T14:29:24+03:00"}

// Test fields
//var fields = logger.Fields{
//	"hello": "world",
//}
// -> log.Info("Run HTTP-CHI API", fields)
// -> {"level":"info","msg":"Run HTTP-CHI API","time":"2019-11-02T14:29:24+03:00", "hello": "world"}
