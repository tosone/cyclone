package common

const (
	// AdminTenant is name of the system admin tenant, it's a default tenant created when Cyclone
	// start, and resources shared among all tenants would be placed in this tenant, such as stage
	// templates.
	AdminTenant = "admin"

	// TenantNamespacePrefix is the prefix of namespace which related to a specific tenant
	TenantNamespacePrefix = "cyclone-"

	// TenantPVCPrefix is the prefix of pvc which related to a specific tenant
	TenantPVCPrefix = "cyclone-pvc-"

	// DefaultPVCSize is the default size of pvc
	DefaultPVCSize = "5Gi"

	// AnnotationTenant is the annotation key used for namespace to relate tenant information
	AnnotationTenant = "cyclone.io/tenant-info"

	// LabelProjectName is the label key used to indicate the project which the resources belongs to
	LabelProjectName = "cyclone.io/project-name"

	// LabelWorkflowName is the label key used to indicate the workflow which the resources belongs to
	LabelWorkflowName = "cyclone.io/workflow-name"

	// LabelIntegrationType is the label key used to indicate type of integration
	LabelIntegrationType = "cyclone.io/integration-type"

	// LabelClusterOn is the label key used to indicate the cluster is a worker for the tenant
	LabelClusterOn = "cyclone.io/cluster-worker"

	// LabelTrueValue is the label value used to represent true
	LabelTrueValue = "true"

	// LabelFalseValue is the label value used to represent false
	LabelFalseValue = "false"

	// LabelOwner is the label key used to indicate namespaces created by cyclone
	LabelOwner = "cyclone.io/owner"

	// OwnerCyclone is the label value used to indicate namespaces created by cyclone
	OwnerCyclone = "cyclone"

	// LabelBuiltin is the label key used to represent cyclone built in resources
	LabelBuiltin = "cyclone.io/builtin"

	// LabelScene is the label key used to indicate cyclone scenario
	LabelScene = "cyclone.io/scene"

	// SceneCICD is the label value used to indicate cyclone CI/CD scenario
	SceneCICD = "cicd"

	// SceneAI is the label value used to indicate cyclone AI scenario
	SceneAI = "ai"

	// LabelStageTemplate is the label key used to represent a stage is a stage template
	LabelStageTemplate = "cyclone.io/stage-template"

	// LabelAcceleration is the label key used to indicate a workflowrun turned on acceleration
	LabelAcceleration = "workflowrun.cyclone.io/acceleration"

	// AnnotationAlias is the annotation key used to indicate the alias of resources
	AnnotationAlias = "cyclone.io/alias"

	// AnnotationDescription is the annotation key used to describe resources
	AnnotationDescription = "cyclone.io/description"

	// QuotaCPULimit represents default value of 'limits.cpu'
	QuotaCPULimit = "2"
	// QuotaCPURequest represents default value of 'requests.cpu'
	QuotaCPURequest = "1"
	// QuotaMemoryLimit represents default value of 'limits.memory'
	QuotaMemoryLimit = "4Gi"
	// QuotaMemoryRequest represents default value of 'requests.memory'
	QuotaMemoryRequest = "2Gi"

	// SecretKeyIntegration is the key of the secret dada to indicate its value is about integration information.
	SecretKeyIntegration = "integration"
	// SecretKeyRepos is the key of the secret dada to records webhooks created for workflowtriggers.
	// Only for SCM integration secrets.
	SecretKeyRepos = "repos"

	// ControlClusterName is the name of control cluster
	ControlClusterName = "control-cluster"

	// PodLabelSelector is selector used to select pod created by Cyclone stages
	// TODO(robin) Copy from pkg/workflow/common/const.go, need to create package pkg/common
	// to merge pkg/workflow/common and pkg/server/common.
	PodLabelSelector = "cyclone.io/workflow==true"

	// CachePrefixPath is the prefix path of acceleration caches
	CachePrefixPath = "caches"
)
