package loki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	loggingv1 "github.com/openshift/cluster-logging-operator/pkg/apis/logging/v1"
	"github.com/openshift/cluster-logging-operator/test/client"
	"github.com/openshift/cluster-logging-operator/test/helpers"
	"github.com/openshift/cluster-logging-operator/test/helpers/loki"
	"github.com/openshift/cluster-logging-operator/test/runtime"
)

func TestLogForwardingToLoki(t *testing.T) {
	c := client.NewTest()
	defer c.Close()
	rcv := loki.NewReceiver(c.NS.Name, "loki-receiver")
	e2e := helpers.NewE2ETestFramework()
	defer e2e.Cleanup()

	cl := helpers.NewClusterLogging(helpers.ComponentTypeCollector)
	clf := runtime.NewClusterLogForwarder()
	clf.Spec = loggingv1.ClusterLogForwarderSpec{
		Outputs: []loggingv1.OutputSpec{{
			Name: rcv.Name,
			Type: loggingv1.OutputTypeLoki,
			URL:  fmt.Sprintf(rcv.InternalURL("").String()),
		}},
		Pipelines: []loggingv1.PipelineSpec{
			{
				Name:       "test-app",
				InputRefs:  []string{loggingv1.InputNameApplication},
				OutputRefs: []string{rcv.Name},
			},
			{
				Name:       "test-audit",
				InputRefs:  []string{loggingv1.InputNameAudit},
				OutputRefs: []string{rcv.Name},
			},
			{
				Name:       "test-infra",
				InputRefs:  []string{loggingv1.InputNameInfrastructure},
				OutputRefs: []string{rcv.Name},
			},
		},
	}

	// Start independent components in parallel to speed up the test.
	var g errgroup.Group
	g.Go(func() error { return e2e.DeployLogGenerator() })
	g.Go(func() error { return rcv.Create(c.Client) })
	g.Go(func() error { return e2e.CreateClusterLogging(cl) })
	g.Go(func() error { return e2e.CreateClusterLogForwarder(clf) })
	require.NoError(t, g.Wait())
	require.NoError(t, e2e.WaitFor(helpers.ComponentTypeCollector))

	// Now the actual test.
	for _, logType := range []string{"application", "infrastructure", "audit"} {
		r, err := rcv.QueryUntil(fmt.Sprintf(`{log_type=%q}`, logType), 1)
		assert.NoError(t, err, "failed query for %v", logType)
		assert.Len(t, r, 1, "single log stream")
		assert.NotEmpty(t, r[0].Lines(), "no log lines for %v", logType)
		assert.Equal(t, r[0].Stream["log_type"], logType, "wrong type of logs")
	}
}
