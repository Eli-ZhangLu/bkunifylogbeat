package main

import (
	"fmt"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/libgse/beat"
	"os"
	"path/filepath"
	"testing"

	"github.com/TencentBlueKing/bkunifylogbeat/beater"
	"github.com/elastic/beats/libbeat/cmd/instance"
	"github.com/elastic/beats/libbeat/publisher/processing"
	"github.com/stretchr/testify/assert"
)

func TestLogBeat(t *testing.T) {
	absPath, err := filepath.Abs("./tests/conf/")
	assert.NotNil(t, absPath)
	assert.NoError(t, err)
	configFile := absPath + "/winlog.conf"
	os.Args = []string{"cmd", "-c", configFile}

	//step 1: 初始化采集器
	settings := instance.Settings{
		Processing: processing.MakeDefaultSupport(false),
	}
	config, err := beat.InitWithPublishConfig(beatName, version, beat.PublishConfig{
		PublishMode: beat.GuaranteedSend,
		ACKEvents:   beater.AckEvents,
	}, settings)
	if err != nil {
		fmt.Printf("Init filed with error: %s\n", err.Error())
		os.Exit(1)
	}

	// step 2：加载配置
	bt, err := beater.New(config)
	if err != nil {
		fmt.Printf("New failed with error: %s\n\n", err.Error())
		os.Exit(1)
	}
	// step 3：主动开启采集器
	//go func() {
	//	_, err = os.Stat("/data/bkunifylogbeat/conf/task.conf")
	//	if err != nil {
	//		close(beat.Done)
	//	}
	//}()

	err = bt.Run()
	assert.NoError(t, err)
}
