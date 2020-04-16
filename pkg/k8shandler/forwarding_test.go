package k8shandler

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
	logging "github.com/openshift/cluster-logging-operator/pkg/apis/logging/v1"
	"github.com/openshift/cluster-logging-operator/pkg/apis/logging/v1/outputs"
	. "github.com/openshift/cluster-logging-operator/test"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/yaml"
)

const (
	aNamespace      = "aNamespace"
	otherTargetName = "someothername"
)

// Match condition by type, status and reason if reason != "".
func MatchCondition(ctype logging.ConditionType, status bool, reason logging.ConditionReason, messageRegex string) types.GomegaMatcher {
	fields := Fields{"Type": Equal(ctype), "Status": Equal(logging.ConditionStatusOf(status))}
	if reason != "" {
		fields["Reason"] = Equal(reason)
	}
	if messageRegex != "" {
		fields["Message"] = MatchRegexp(messageRegex)
	}
	return MatchFields(IgnoreExtras, fields)
}

func HaveCondition(t logging.ConditionType, status bool, r logging.ConditionReason, messageRegex string) types.GomegaMatcher {
	return ContainElement(MatchCondition(t, status, r, messageRegex))
}

var _ = Describe("Normalizing forwarder", func() {

	var (
		cluster     *logging.ClusterLogging
		output      logging.OutputSpec
		otherOutput logging.OutputSpec
		request     *ClusterLoggingRequest
	)
	BeforeEach(func() {
		output = logging.OutputSpec{
			Name: "myOutput",
			Type: "elasticsearch",
			URL:  "anOutPut",
		}
		otherOutput = logging.OutputSpec{
			Name: otherTargetName,
			Type: "elasticsearch",
			URL:  "someotherendpoint",
		}
		request = &ClusterLoggingRequest{
			client: fake.NewFakeClient(),
			cluster: &logging.ClusterLogging{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: aNamespace,
				},
			},
			ForwarderRequest: logging.ClusterLogForwarder{},
		}
		cluster = request.cluster
	})

	Context("while validating ", func() {

		BeforeEach(func() {
			request.ForwarderSpec = logging.ClusterLogForwarderSpec{
				Outputs: []logging.OutputSpec{
					output,
					otherOutput,
				},
				Pipelines: []logging.PipelineSpec{
					{
						Name:       "aPipeline",
						OutputRefs: []string{output.Name, otherOutput.Name},
						InputRefs:  []string{logging.InputNameApplication},
					},
				},
			}
		})
		Context("pipelines", func() {

			It("should only include inputs if there is at least one valid pipeline", func() {
				request.ForwarderSpec.Pipelines = []logging.PipelineSpec{
					{
						Name:       "aPipeline",
						OutputRefs: []string{"someotherendpoint"},
						InputRefs:  []string{logging.InputNameApplication},
					},
				}
				spec, status := request.normalizeForwarder()
				Expect(spec.Pipelines).To(BeEmpty(), "Exp. all pipelines to be dropped")
				Expect(status.Inputs).To(BeEmpty())
			})

			It("should drop pipelines that do not have unique names", func() {
				request.ForwarderSpec.Pipelines = append(request.ForwarderSpec.Pipelines,
					logging.PipelineSpec{
						Name:       "aPipeline",
						OutputRefs: []string{output.Name, otherOutput.Name},
						InputRefs:  []string{logging.InputNameApplication},
					})
				spec, status := request.normalizeForwarder()
				Expect(spec.Pipelines).To(HaveLen(1))
				Expect(status.Pipelines).To(HaveKey("pipeline[1]"))
				Expect(status.Pipelines["pipeline[1]"]).To(HaveCondition(logging.ConditionReady, false, "Invalid", "duplicate"))
				Expect(status.Pipelines).To(HaveLen(2))
			})

			It("should allow pipelines with empty/missing names", func() {
				request.ForwarderSpec.Pipelines = append(request.ForwarderSpec.Pipelines,
					logging.PipelineSpec{
						OutputRefs: []string{output.Name},
						InputRefs:  []string{logging.InputNameInfrastructure},
					})
				spec, _ := request.normalizeForwarder()
				Expect(spec.Pipelines).To(HaveLen(2), "Exp. all pipelines to be ok")
				Expect(spec.Pipelines[0].Name).To(Equal("aPipeline"))
				Expect(spec.Pipelines[1].Name).To(Equal("pipeline[1]"))
			})

			It("should drop pipelines that have unrecognized inputRefs", func() {
				request.ForwarderSpec.Pipelines = []logging.PipelineSpec{
					{
						Name:       "someDefinedPipeline",
						OutputRefs: []string{output.Name, otherOutput.Name},
						InputRefs:  []string{"foo"},
					}}
				spec, status := request.normalizeForwarder()
				conds := status.Pipelines["someDefinedPipeline"]
				Expect(spec.Pipelines).To(BeEmpty(), "Exp. all pipelines to be dropped")
				Expect(conds).To(HaveCondition(logging.ConditionReady, false, logging.ReasonInvalid, `inputs:.*\[foo]`))
			})

			It("should drop pipelines that have no outputRefs", func() {
				request.ForwarderSpec.Pipelines = append(request.ForwarderSpec.Pipelines,
					logging.PipelineSpec{
						Name:       "someDefinedPipeline",
						OutputRefs: []string{},
						InputRefs:  []string{logging.InputNameApplication},
					})
				spec, status := request.normalizeForwarder()
				conds := status.Pipelines["someDefinedPipeline"]
				Expect(conds).To(HaveCondition(logging.ConditionReady, false, logging.ReasonInvalid, "no valid outputs"))
				Expect(spec.Pipelines).NotTo(ContainElement(
					MatchFields(IgnoreExtras, Fields{"Name": Equal("someDefinedPipeline")})))
				Expect(spec.Pipelines).To(HaveLen(1))
			})

			It("should degrade pipelines where there are fewer outputs then defined outputRefs", func() {
				request.ForwarderSpec.Pipelines = append(request.ForwarderSpec.Pipelines,
					logging.PipelineSpec{
						Name:       "someDefinedPipeline",
						OutputRefs: []string{output.Name, otherOutput.Name, "aMissingOutput"},
						InputRefs:  []string{logging.InputNameApplication},
					})
				spec, status := request.normalizeForwarder()
				Expect(spec.Pipelines).To(HaveLen(2), "Exp. all defined pipelines")
				Expect(status.Pipelines).To(HaveLen(2), "Exp. all defined pipelines")
				Expect(status.Pipelines).To(HaveKey("someDefinedPipeline"))
				conds := status.Pipelines["someDefinedPipeline"]
				Expect(conds).To(HaveCondition(logging.ConditionDegraded, true, "Invalid", "aMissingOutput"))
				Expect(conds).To(HaveCondition(logging.ConditionReady, true, "", ""))
			})

		})

		Context("outputs", func() {

			It("should drop outputs that do not have unique names", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name: "myOutput",
					Type: "elasticsearch",
					URL:  "anOutPut",
				})
				//sanity check
				Expect(request.ForwarderSpec.Outputs).To(HaveLen(3))
				spec, status := request.normalizeForwarder()
				Expect(spec.Outputs).To(HaveLen(2), "Exp. non-unique outputs to be dropped")
				Expect(status.Outputs["myOutput"]).To(HaveCondition(logging.ConditionReady, true, "", ""))
				Expect(status.Outputs["output[2]"]).To(HaveCondition(logging.ConditionReady, false, logging.ReasonInvalid, "duplicate"))
			})

			It("should drop outputs that have empty names", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name: "",
					Type: "elasticsearch",
					URL:  "anOutPut",
				})
				spec, status := request.normalizeForwarder()
				Expect(spec.Outputs).To(HaveLen(2), "Exp. outputs with an empty name to be dropped")
				Expect(status.Outputs["output[2]"]).To(HaveCondition("Ready", false, "Invalid", "must have a name"))
			})

			It("should drop outputs that conflict with the internally reserved name", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name: logging.OutputNameDefault, // FIXME(alanconway) case insensitivity
					Type: "Elasticsearch",
					URL:  "anOutPut",
				})
				spec, status := request.normalizeForwarder()
				Expect(spec.Outputs).To(HaveLen(2), "Exp. outputs with an internal name conflict to be dropped")
				Expect(status.Outputs).To(HaveKey("output[2]"))
				Expect(status.Outputs["output[2]"]).To(HaveCondition("Ready", false, "Invalid", "reserved"))
			})

			It("should drop outputs that have empty types", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name: "aName",
					URL:  "anOutPut",
				})
				spec, status := request.normalizeForwarder()
				Expect(len(spec.Outputs)).To(Equal(2), "Exp. outputs with an empty type to be dropped")
				Expect(status.Outputs["aName"]).To(HaveCondition("Ready", false, "Invalid", "unknown.*\"\""))
			})
			It("should drop outputs that have unrecognized types", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name: "aName",
					Type: "foo",
					URL:  "anOutPut",
				})
				spec, status := request.normalizeForwarder()
				Expect(len(spec.Outputs)).To(Equal(2), "Exp. outputs with an unrecognized type to be dropped")
				Expect(status.Outputs["aName"]).To(HaveCondition("Ready", false, "Invalid", "unknown.*\"foo\""))
			})

			It("should drop outputs that have empty URL", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name: "aName",
					Type: "fluentForward",
					URL:  "",
				})
				spec, status := request.normalizeForwarder()
				Expect(len(spec.Outputs)).To(Equal(2), "Exp. outputs with an empty endpoint to be dropped")
				Expect(status.Outputs["aName"]).To(HaveCondition("Ready", false, "Invalid", "missing URL"))
			})

			It("should drop outputs that have an invalid URL", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name: "aName",
					Type: "fluentForward",
					URL:  "",
				})
				spec, status := request.normalizeForwarder()
				Expect(len(spec.Outputs)).To(Equal(2), "Exp. outputs with an empty endpoint to be dropped")
				Expect(status.Outputs["aName"]).To(HaveCondition("Ready", false, "Invalid", "missing URL"))
			})

			It("should drop outputs that have secrets with no names", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name:   "aName",
					Type:   "elasticsearch",
					URL:    "an output",
					Secret: &logging.OutputSecretSpec{},
				})
				spec, status := request.normalizeForwarder()
				Expect(len(spec.Outputs)).To(Equal(2), "Exp. outputs with empty secrets to be dropped")
				Expect(status.Outputs["aName"]).To(HaveCondition("Ready", false, "Invalid", "secretRef must have a name"))
			})

			It("should drop outputs that have secrets which don't exist", func() {
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name:   "aName",
					Type:   "elasticsearch",
					URL:    "an output",
					Secret: &logging.OutputSecretSpec{Name: "mysecret"},
				})
				spec, status := request.normalizeForwarder()
				Expect(len(spec.Outputs)).To(Equal(2), "Exp. outputs with non-existent secrets to be dropped")
				Expect(status.Outputs["aName"]).To(HaveCondition("Ready", false, "MissingResource", "secret.*not found"))
			})

			It("should accept well formed outputs", func() {
				request.client = fake.NewFakeClient(&corev1.Secret{
					TypeMeta: metav1.TypeMeta{
						Kind:       "Secret",
						APIVersion: corev1.SchemeGroupVersion.String(),
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "mysecret",
						Namespace: aNamespace,
					},
				})
				request.ForwarderSpec.Outputs = append(request.ForwarderSpec.Outputs, logging.OutputSpec{
					Name:   "aName",
					Type:   "elasticsearch",
					URL:    "an output",
					Secret: &logging.OutputSecretSpec{Name: "mysecret"},
				})
				spec, status := request.normalizeForwarder()
				Expect(status.Outputs["aName"]).To(HaveCondition("Ready", true, "", ""))
				Expect(spec.Outputs).To(HaveLen(3))
			})
		})
	})

	Context("with empty forwarder spec", func() {
		BeforeEach(func() { request.ForwarderSpec = logging.ClusterLogForwarderSpec{} })

		It("returns bad status on default output with no default logstore", func() {
			spec, status := request.normalizeForwarder()
			Expect(spec.Pipelines).To(BeEmpty())
			Expect(status.Conditions).To(HaveCondition("Ready", false, "", ""))
			Expect(status.Outputs).To(HaveKey(logging.OutputNameDefault))
			Expect(status.Outputs[logging.OutputNameDefault]).To(HaveCondition("Ready", false, "MissingResource", "log store"))
			Expect(spec).To(Equal(&logging.ClusterLogForwarderSpec{}))
		})

		It("forwards all logs to the default logstore", func() {
			cluster.Spec.LogStore = &logging.LogStoreSpec{
				Type: logging.LogStoreTypeElasticsearch,
			}
			spec, status := request.normalizeForwarder()

			Expect(spec.Outputs).To(HaveLen(1))
			Expect(spec.Outputs[0].Name).To(Equal(logging.OutputNameDefault))

			Expect(spec.Pipelines).To(HaveLen(1))
			pipeline := spec.Pipelines[0]
			Expect(pipeline.InputRefs).To(ConsistOf("application", "infrastructure", "audit"))
			Expect(pipeline.OutputRefs).To(ConsistOf("default"))

			Expect(status.Conditions).To(HaveCondition("Ready", true, "", ""))
			Expect(status.Pipelines["pipeline[0]"]).To(HaveCondition("Ready", true, "", ""))
			Expect(status.Outputs[logging.OutputNameDefault]).To(HaveCondition("Ready", true, "", ""))
			Expect(status.Inputs[logging.InputNameApplication]).To(HaveCondition("Ready", true, "", ""))
			Expect(status.Inputs[logging.InputNameInfrastructure]).To(HaveCondition("Ready", true, "", ""))
		})
	})

	It("parses spec with syslog", func() {
		request.ForwarderSpec = logging.ClusterLogForwarderSpec{
			Inputs: []logging.InputSpec{
				{
					Name: "in",
					Type: "application",
				},
			},
			Outputs: []logging.OutputSpec{
				{
					Name: "out",
					Type: "syslog",
					URL:  "tls:blahblah",
					OutputTypeSpec: logging.OutputTypeSpec{
						Syslog: &outputs.Syslog{},
					},
				},
			},
			Pipelines: []logging.PipelineSpec{
				{
					Name:       "test",
					InputRefs:  []string{"in"},
					OutputRefs: []string{"out"},
				},
			},
		}
		// FIXME(alanconway)
		// forwarder.Spec.Outputs[0].Syslog.RFC = "RFC5424"
		spec, status := request.normalizeForwarder()
		Expect(status.Conditions).To(HaveCondition("Ready", true, "", ""), "unexpected "+YAMLString(status))
		_ = spec // FIXME(alanconway)
	})
})

var _ = DescribeTable("Normalizing valid forwarder YAML specs",

	func(yamlSpec string, goSpec logging.ClusterLogForwarderSpec) {
		request := ClusterLoggingRequest{
			client: fake.NewFakeClient(),
			cluster: &logging.ClusterLogging{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: aNamespace,
				},
			},
		}
		Expect(yaml.Unmarshal([]byte(yamlSpec), &request.ForwarderSpec)).To(Succeed())
		spec, status := request.normalizeForwarder()
		Expect(status.Conditions).To(HaveCondition("Ready", true, "", ""), spec)
		Expect(goSpec).To(EqualDiff(*spec))
	},
	Entry("application to syslog", `
outputs:
- name: out
  type: syslog
  syslog:
    severity: Alert
  url: syslog-receiver.openshift-logging.svc:24224
pipelines:
  - inputRefs: [ application ]
    outputRefs: [ out ]
`,
		logging.ClusterLogForwarderSpec{
			Outputs: []logging.OutputSpec{
				{
					Name: "out",
					Type: "syslog",
					URL:  "syslog-receiver.openshift-logging.svc:24224",
					OutputTypeSpec: logging.OutputTypeSpec{
						Syslog: &outputs.Syslog{Severity: "Alert"},
					},
				},
			},
			Pipelines: []logging.PipelineSpec{
				{
					Name:       "pipeline[0]",
					InputRefs:  []string{"application"},
					OutputRefs: []string{"out"},
				},
			},
		},
	))
