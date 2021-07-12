package v1

// NOTE: The Enum validation on OutputSpec.Type must be updated if the list of
// known types changes.

// Output type constants, must match JSON tags of OutputTypeSpec fields.
const (
	OutputTypeElasticsearch  = "elasticsearch"
	OutputTypeFluentdForward = "fluentdForward"
	OutputTypeSyslog         = "syslog"
	OutputTypeKafka          = "kafka"
	OutputTypeLoki           = "loki"
)

// OutputTypeSpec is a union of optional additional configuration specific to an
// output type. The fields of this struct define the set of known output types.
type OutputTypeSpec struct {
	// +optional
	Syslog *Syslog `json:"syslog,omitempty"`
	// +optional
	FluentdForward *FluentdForward `json:"fluentdForward,omitempty"`
	// +optional
	Elasticsearch *Elasticsearch `json:"elasticsearch,omitempty"`
	// +optional
	Kafka *Kafka `json:"kafka,omitempty"`
	// +optional
	Loki *Loki `json:"loki,omitempty"`
}

// Syslog provides optional extra properties for output type `syslog`
type Syslog struct {
	// Severity to set on outgoing syslog records.
	//
	// Severity values are defined in https://tools.ietf.org/html/rfc5424#section-6.2.1
	// The value can be a decimal integer or one of these case-insensitive keywords:
	//
	//     Emergency Alert Critical Error Warning Notice Informational Debug
	//
	// +optional
	Severity string `json:"severity,omitempty"`

	// Facility to set on outgoing syslog records.
	//
	// Facility values are defined in https://tools.ietf.org/html/rfc5424#section-6.2.1.
	// The value can be a decimal integer. Facility keywords are not standardized,
	// this API recognizes at least the following case-insensitive keywords
	// (defined by https://en.wikipedia.org/wiki/Syslog#Facility_Levels):
	//
	//     kernel user mail daemon auth syslog lpr news
	//     uucp cron authpriv ftp ntp security console solaris-cron
	//     local0 local1 local2 local3 local4 local5 local6 local7
	//
	// +optional
	Facility string `json:"facility,omitempty"`

	// TrimPrefix is a prefix to trim from the tag.
	//
	// +optional
	TrimPrefix string `json:"trimPrefix,omitempty"`

	// Tag specifies a record field to use as tag.
	//
	// +optional
	Tag string `json:"tag,omitempty"`

	// PayloadKey specifies record field to use as payload.
	//
	// +optional
	PayloadKey string `json:"payloadKey,omitempty"`

	// AddLogSource adds log's source information to the log message
	// If the logs are collected from a process; namespace_name, pod_name, container_name is added to the log
	//
	// +optional
	AddLogSource bool `json:"addLogSource,omitempty"`

	// Rfc specifies the rfc to be used for sending syslog
	//
	// Rfc values can be one of:
	//  - RFC3164 (https://tools.ietf.org/html/rfc3164)
	//  - RFC5424 (https://tools.ietf.org/html/rfc5424)
	//
	// If unspecified, RFC5424 will be assumed.
	//
	// +kubebuilder:validation:Enum:=RFC3164;RFC5424
	// +kubebuilder:default:=RFC5424
	// +optional
	RFC string `json:"rfc,omitempty"`

	// AppName is APP-NAME part of the syslog-msg header
	//
	// AppName needs to be specified if using rfc5424
	//
	// +optional
	AppName string `json:"appName,omitempty"`

	// ProcID is PROCID part of the syslog-msg header
	//
	// ProcID needs to be specified if using rfc5424
	//
	// +optional
	ProcID string `json:"procID,omitempty"`

	// MsgID is MSGID part of the syslog-msg header
	//
	// MsgID needs to be specified if using rfc5424
	//
	// +optional
	MsgID string `json:"msgID,omitempty"`
}

// Kafka provides optional extra properties for `type: kafka`
type Kafka struct {
	// Topic specifies the target topic to send logs to.
	//
	// +optional
	Topic string `json:"topic,omitempty"`

	// Brokers specifies the list of brokers
	// to register in addition to the main output URL
	// on initial connect to enhance reliability.
	//
	// +optional
	Brokers []string `json:"brokers,omitempty"`
}

type FluentdForward struct{}

type Elasticsearch struct {
	// StructuredTypeKey specifies the metadata key to be used as name of elasticsearch index
	// It takes precedence over StructuredTypeName
	//
	// +optional
	StructuredTypeKey string `json:"structuredTypeKey,omitempty"`

	// StructuredTypeName specifies the name of elasticsearch schema
	//
	// +optional
	StructuredTypeName string `json:"structuredTypeName,omitempty"`
}

// Loki provides optional extra properties for `type: loki`
type Loki struct {
	// LabelKeys is a list of meta-data field keys to replace the default Loki labels.
	//
	// Loki label names must match the regular expression "[a-zA-Z_:][a-zA-Z0-9_:]*".
	// Illegal characters in meta-data keys are replaced with "_" to form the label name.
	// For example the meta-data key "kubernetes.labels.foo" becomes Loki label "kubernetes_labels_foo".
	//
	// If LabelKeys is not set, the default labels are:
	// [log_type, kubernetes_namespace_name, kubernetes_pod_name, kubernetes_host]
	//
	// Note: the label kubernetes_host is always included, even if not requested.
	// Loki requires log streams to be correctly ordered by timestamp.
	// Including kubernetes_host in the label set ensures each stream originates from a single host,
	// and timestamps cannot become disordered due to clock differences on different hosts.
	//
	// +optional
	LabelKeys []string `json:"labelKeys,omitempty"`
}
