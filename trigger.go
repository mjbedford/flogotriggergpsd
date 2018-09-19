package flogotriggergpsd

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/stratoberry/go-gpsd"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("flogotriggergpsd")

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct{
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata:md}
}

// New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config:config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
}

// Initialize implements trigger.Init.Initialize
func (t *MyTrigger) Initialize(ctx trigger.InitContext) error {
	return nil
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	log.Info("Trigger.Start")
	filterName :=t.config.GetSetting("filter")
	//filterName := context.GetInput("filter").(string)
	log.Info("Filter : " + filterName)
	if gps, err := gpsd.Dial(gpsd.DefaultAddress); err != nil {
		log.Error(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}
	log.Info("After Dial")
	/* if err != nil{
		log.Info("Connected to gpsd")
		if filterName == "TPV"{
			tpvFilter := func(r interface{}) {
			report := r.(*gpsd.TPVReport)
			log.Info("Location updated", report.Lat, report.Lon)
			}

			gps.AddFilter("TPV", tpvFilter)
		} else {
			log.Info("Filter not TPV")
		}

	} else {
	log.Error(err)
	log.Info(err) */
	}
// start the trigger
	return nil
}



// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}
