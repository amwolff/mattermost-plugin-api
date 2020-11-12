package cluster

import (
	"time"

	"github.com/mattermost/mattermost-server/v5/plugin"
)

func ExampleScheduleOnce() {
	// Use p.API from your plugin instead.
	pluginAPI := plugin.API(nil)

	callback := func(key string) {
		if key == "the key i'm watching for" {
			// Work to do only once per cluster
		}
	}

	// Get the scheduler, which you can pass throughout the plugin...
	scheduler := GetJobOnceScheduler(pluginAPI)

	// And add callbacks to, near the parts of the code that will handle those jobs...
	_, _ = scheduler.AddCallback(callback)

	// After the plugin has added its callbacks, start the scheduler, which schedules all waiting jobs.
	_ = scheduler.Start()

	// main thread

	// add a job
	_, _ = scheduler.ScheduleOnce("the key i'm watching for", time.Now().Add(2*time.Hour))

	// Maybe you want to check the scheduled jobs, or close them:
	jobs, _ := scheduler.ListScheduledJobs()
	defer func() {
		for _, j := range jobs {
			scheduler.Close(j.Key)
		}
	}()
}
