# pulsarzap
zap logger wrapper for pulsar Logger

## Usage

```go
import "github.com/ttys3/pulsarzap"
```

### simple usage

```go
client, err := pulsar.NewClient(pulsar.ClientOptions{
    URL:               "pulsar://pulsar-broker.service.dc1.consul:6650",
    Logger:            pulsarzap.NewDefault(),
})
```

### use custom zap logger config

console encoding logger for development

```go
lc := zap.NewDevelopmentConfig()
lc.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
lc.DisableStacktrace = true
lc.Encoding = "console"
lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
logger, _ := lc.Build(zap.WithCaller(true))
log := pulsarzap.New(logger.Sugar())

client, err := pulsar.NewClient(pulsar.ClientOptions{
    URL:               "pulsar://pulsar-broker.service.dc1.consul:6650",
    Logger:            log,
})
```

json encoding logger for production

```go
lc := zap.NewProductionConfig()
lc.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
logger, _ := lc.Build(zap.WithCaller(true))
log := pulsarzap.New(logger.Sugar())

client, err := pulsar.NewClient(pulsar.ClientOptions{
    URL:               "pulsar://pulsar-broker.service.dc1.consul:6650",
    Logger:            log,
})
```
