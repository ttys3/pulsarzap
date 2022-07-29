package pulsarlogger

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/apache/pulsar-client-go/pulsar"
)

func TestPulsarLogger(t *testing.T) {
	lc := zap.NewDevelopmentConfig()
	lc.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	lc.DisableStacktrace = true
	lc.Encoding = "console"
	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := lc.Build(zap.WithCaller(true))
	log := New(logger.Sugar())

	log.Infof("this is a pulsarlogger with custom config")

	log.Debugf("begin pulsar.NewClient")

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://pulsar-broker.service.dc1.consul:6650",
		OperationTimeout:  3 * time.Second,
		ConnectionTimeout: 3 * time.Second,
		Logger:            log,
	})
	if err != nil {
		log.Errorf("Could not instantiate Pulsar client: %v", err)
		return
	}
	log.Info("Pulsar client instantiated")

	defer client.Close()

	log.Info("begin Subscribe ", "my-topic")
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:                       "my-topic",
		SubscriptionName:            "my-sub",
		Name:                        "consumer-name001",
		Type:                        pulsar.Shared,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
	})
	if err != nil {
		log.Error(err)
		return
	}
	defer consumer.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
		// time.Sleep(time.Millisecond * 100)
	}

	if err := consumer.Unsubscribe(); err != nil {
		log.Error(err)
	}
}

func TestPulsarLoggerNewCustom(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	plog := New(sugar)
	plog.Infof("hello %s", "world")

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://pulsar-broker.service.dc1.consul:6650",
		OperationTimeout:  6 * time.Second,
		ConnectionTimeout: 6 * time.Second,
		Logger:            plog,
	})
	if err != nil {
		plog.Errorf("Could not instantiate Pulsar client: %v", err)
		return
	}
	plog.Info("Pulsar client instantiated")

	defer client.Close()
}

func TestPulsarLoggerDefault(t *testing.T) {
	plog := NewDefault()
	plog.Infof("hello %s", "world")
	plog.Info("hello world", " this", " is", " a", " test")

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://pulsar-broker.service.dc1.consul:6650",
		OperationTimeout:  6 * time.Second,
		ConnectionTimeout: 6 * time.Second,
		Logger:            plog,
	})
	if err != nil {
		plog.Errorf("Could not instantiate Pulsar client: %v", err)
		return
	}
	plog.Info("Pulsar client instantiated")

	defer client.Close()
}
