package events

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-community/go-cfclient"
	sonde_events "github.com/cloudfoundry/sonde-go/events"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

//go:generate counterfeiter -o mocks/mock_registerer.go . prometheus.Registerer

var _ = Describe("AppWatcher", func() {
	var (
		appWatcher *AppWatcher
		// Apps        []cfclient.App
		registerer *prometheus.Registry
	)

	BeforeEach(func() {

		config := &cfclient.Config{
			ApiAddress:        "some/endpoint",
			SkipSslValidation: true,
			Username:          "barry",
			Password:          "password",
			ClientID:          "dummy_client_id",
			ClientSecret:      "dummy_client_secret",
		}

		apps := []cfclient.App{
			{Guid: "33333333-3333-3333-3333-333333333333", Instances: 1, Name: "foo", SpaceURL: "/v2/spaces/123"},
		}

		log.Printf("app: %v", apps)
		log.Printf("app: %v", config)
		registerer = prometheus.NewRegistry()
		appWatcher = NewAppWatcher(config, apps[0], registerer)

		log.Printf("app: %v", appWatcher)

	})
	AfterEach(func() {})

	Describe("AppName", func() {
		It("knows the name of its application", func() {
			Expect(appWatcher.AppName()).To(Equal("foo"))
		})
	})

	Describe("Run", func() {
		It("Registers metrics on startup", func() {
			go appWatcher.Run()
			defer {
				appWatcher.Close()
			}

			Expect(registerer.MustRegister()).ToHaveBeenCalledWith(appWatcher.metricsForInstance[0].cpu)
		})
	})

	Describe("Run", func() {
		It("Unregisters metrics on close", func() {
			//
		})
	})

	Describe("Run", func() {
		It("Registers more metrics when new instances are created", func() {
			//
		})
	})

	Describe("Run", func() {
		It("Unregisters some metrics when old instances are deleted", func() {
			//
		})
	})

	Describe("processContainerMetrics", func() {
		It("sets a CPU metric on an instance", func() {
			cpuPercentage := 10.0
			var instanceIndex int32 = 0
			containerMetric := sonde_events.ContainerMetric{
				CpuPercentage: &cpuPercentage,
				InstanceIndex: &instanceIndex,
			}
			appWatcher.processContainerMetric(&containerMetric)
			cpuGauge := appWatcher.metricsForInstance[instanceIndex].cpu

			Expect(testutil.ToFloat64(cpuGauge)).To(Equal(cpuPercentage))
		})
	})
})
