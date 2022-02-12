package common

import (
	"time"

	v1 "k8s.io/api/core/v1"
)

// Common types and constants used by the importer and controller.
// TODO: maybe the vm cloner can use these common values

const (
	// CDILabelKey provides a constant for CDI PVC labels
	CDILabelKey = "app"
	// CDILabelValue provides a constant  for CDI PVC label values
	CDILabelValue = "containerized-data-importer"
	// CDILabelSelector provides a constant to use for the selector to identify CDI objects in list
	CDILabelSelector = CDILabelKey + "=" + CDILabelValue

	// CDIComponentLabel can be added to all CDI resources
	CDIComponentLabel = "cdi.kubevirt.io"
	// CDIControllerName is the CDI controller name
	CDIControllerName = "cdi-controller"

	// AppKubernetesPartOfLabel is the Kubernetes recommended part-of label
	AppKubernetesPartOfLabel = "app.kubernetes.io/part-of"
	// AppKubernetesVersionLabel is the Kubernetes recommended version label
	AppKubernetesVersionLabel = "app.kubernetes.io/version"
	// AppKubernetesManagedByLabel is the Kubernetes recommended managed-by label
	AppKubernetesManagedByLabel = "app.kubernetes.io/managed-by"
	// AppKubernetesComponentLabel is the Kubernetes recommended component label
	AppKubernetesComponentLabel = "app.kubernetes.io/component"

	// PrometheusLabelKey provides the label to indicate prometheus metrics are available in the pods.
	PrometheusLabelKey = "prometheus.cdi.kubevirt.io"
	// PrometheusLabelValue provides the label value which shouldn't be empty to avoid a prometheus WIP issue.
	PrometheusLabelValue = "true"
	// PrometheusServiceName is the name of the prometheus service created by the operator.
	PrometheusServiceName = "cdi-prometheus-metrics"
	// KubePersistentVolumeFillingUpSuppressLabelKey is the label name that helps suppress this alert for our PVCs
	KubePersistentVolumeFillingUpSuppressLabelKey = "alerts.k8s.io/KubePersistentVolumeFillingUp"
	// KubePersistentVolumeFillingUpSuppressLabelValue is the label value that helps suppress this alert for our PVCs
	KubePersistentVolumeFillingUpSuppressLabelValue = "disabled"

	// UploadTargetLabel has the UID of upload target PVC
	UploadTargetLabel = CDIComponentLabel + "/uploadTarget"

	// DataImportCronLabel has the name of the DataImportCron responsible for the labeled DataSource or DataVolume
	DataImportCronLabel = CDIComponentLabel + "/dataImportCron"
	// DataImportCronCleanupLabel tells whether to delete the resource when its DataImportCron is deleted
	DataImportCronCleanupLabel = DataImportCronLabel + ".cleanup"

	// ImporterVolumePath provides a constant for the directory where the PV is mounted.
	ImporterVolumePath = "/data"
	// DiskImageName provides a constant for our importer/datastream_ginkgo_test and to build ImporterWritePath
	DiskImageName = "disk.img"
	// ImporterWritePath provides a constant for the cmd/cdi-importer/importer.go executable
	ImporterWritePath = ImporterVolumePath + "/" + DiskImageName
	// WriteBlockPath provides a constant for the path where the PV is mounted.
	WriteBlockPath = "/dev/cdi-block-volume"
	// NbdkitLogPath provides a constant for the path in which the nbdkit log messages are stored.
	NbdkitLogPath = "/tmp/nbdkit.log"
	// PodTerminationMessageFile is the name of the file to write the termination message to.
	PodTerminationMessageFile = "/dev/termination-log"
	// ImporterPodName provides a constant to use as a prefix for Pods created by CDI (controller only)
	ImporterPodName = "importer"
	// ImporterDataDir provides a constant for the controller pkg to use as a hardcoded path to where content is transferred to/from (controller only)
	ImporterDataDir = "/data"
	// ScratchDataDir provides a constant for the controller pkg to use as a hardcoded path to where scratch space is located.
	ScratchDataDir = "/scratch"
	// ImporterS3Host provides an S3 string used by importer/dataStream.go only
	ImporterS3Host = "s3.amazonaws.com"
	// ImporterCertDir is where the configmap containing certs will be mounted
	ImporterCertDir = "/certs"
	// DefaultPullPolicy imports k8s "IfNotPresent" string for the import_controller_gingko_test and the cdi-controller executable
	DefaultPullPolicy = string(v1.PullIfNotPresent)
	// ImportProxyConfigMapName provides the name of the ConfigMap in the cdi namespace containing a CA certificate bundle
	ImportProxyConfigMapName = "trusted-ca-proxy-bundle-cm"
	// ImportProxyConfigMapKey provides the key name of the ConfigMap in the cdi namespace containing a CA certificate bundle
	ImportProxyConfigMapKey = "ca.pem"
	// ImporterProxyCertDir is where the configmap containing proxy certs will be mounted
	ImporterProxyCertDir = "/proxycerts/"

	// PullPolicy provides a constant to capture our env variable "PULL_POLICY" (only used by cmd/cdi-controller/controller.go)
	PullPolicy = "PULL_POLICY"
	// ImporterSource provides a constant to capture our env variable "IMPORTER_SOURCE"
	ImporterSource = "IMPORTER_SOURCE"
	// ImporterContentType provides a constant to capture our env variable "IMPORTER_CONTENTTYPE"
	ImporterContentType = "IMPORTER_CONTENTTYPE"
	// ImporterEndpoint provides a constant to capture our env variable "IMPORTER_ENDPOINT"
	ImporterEndpoint = "IMPORTER_ENDPOINT"
	// ImporterAccessKeyID provides a constant to capture our env variable "IMPORTER_ACCES_KEY_ID"
	ImporterAccessKeyID = "IMPORTER_ACCESS_KEY_ID"
	// ImporterSecretKey provides a constant to capture our env variable "IMPORTER_SECRET_KEY"
	ImporterSecretKey = "IMPORTER_SECRET_KEY"
	// ImporterImageSize provides a constant to capture our env variable "IMPORTER_IMAGE_SIZE"
	ImporterImageSize = "IMPORTER_IMAGE_SIZE"
	// ImporterCertDirVar provides a constant to capture our env variable "IMPORTER_CERT_DIR"
	ImporterCertDirVar = "IMPORTER_CERT_DIR"
	// InsecureTLSVar provides a constant to capture our env variable "INSECURE_TLS"
	InsecureTLSVar = "INSECURE_TLS"
	// ImporterDiskID provides a constant to capture our env variable "IMPORTER_DISK_ID"
	ImporterDiskID = "IMPORTER_DISK_ID"
	// ImporterUUID provides a constant to capture our env variable "IMPORTER_UUID"
	ImporterUUID = "IMPORTER_UUID"
	// ImporterReadyFile provides a constant to capture our env variable "IMPORTER_READY_FILE"
	ImporterReadyFile = "IMPORTER_READY_FILE"
	// ImporterDoneFile provides a constant to capture our env variable "IMPORTER_DONE_FILE"
	ImporterDoneFile = "IMPORTER_DONE_FILE"
	// ImporterBackingFile provides a constant to capture our env variable "IMPORTER_BACKING_FILE"
	ImporterBackingFile = "IMPORTER_BACKING_FILE"
	// ImporterThumbprint provides a constant to capture our env variable "IMPORTER_THUMBPRINT"
	ImporterThumbprint = "IMPORTER_THUMBPRINT"
	// ImporterCurrentCheckpoint provides a constant to capture our env variable "IMPORTER_CURRENT_CHECKPOINT"
	ImporterCurrentCheckpoint = "IMPORTER_CURRENT_CHECKPOINT"
	// ImporterPreviousCheckpoint provides a constant to capture our env variable "IMPORTER_PREVIOUS_CHECKPOINT"
	ImporterPreviousCheckpoint = "IMPORTER_PREVIOUS_CHECKPOINT"
	// ImporterFinalCheckpoint provides a constant to capture our env variable "IMPORTER_FINAL_CHECKPOINT"
	ImporterFinalCheckpoint = "IMPORTER_FINAL_CHECKPOINT"
	// Preallocation provides a constant to capture out env variable "PREALLOCATION"
	Preallocation = "PREALLOCATION"
	// ImportProxyHTTP provides a constant to capture our env variable "http_proxy"
	ImportProxyHTTP = "http_proxy"
	// ImportProxyHTTPS provides a constant to capture our env variable "https_proxy"
	ImportProxyHTTPS = "https_proxy"
	// ImportProxyNoProxy provides a constant to capture our env variable "no_proxy"
	ImportProxyNoProxy = "no_proxy"
	// ImporterProxyCertDirVar provides a constant to capture our env variable "IMPORTER_PROXY_CERT_DIR"
	ImporterProxyCertDirVar = "IMPORTER_PROXY_CERT_DIR"
	// InstallerPartOfLabel provides a constant to capture our env variable "INSTALLER_PART_OF_LABEL"
	InstallerPartOfLabel = "INSTALLER_PART_OF_LABEL"
	// InstallerVersionLabel provides a constant to capture our env variable "INSTALLER_VERSION_LABEL"
	InstallerVersionLabel = "INSTALLER_VERSION_LABEL"
	// ImporterExtraHeader provides a constant to include extra HTTP headers, as the prefix to a format string
	ImporterExtraHeader = "IMPORTER_EXTRA_HEADER_"
	// ImporterSecretExtraHeadersDir is where the secrets containing extra HTTP headers will be mounted
	ImporterSecretExtraHeadersDir = "/extraheaders"

	// CloningLabelValue provides a constant to use as a label value for pod affinity (controller pkg only)
	CloningLabelValue = "host-assisted-cloning"
	// CloningTopologyKey  (controller pkg only)
	CloningTopologyKey = "kubernetes.io/hostname"
	// ClonerSourcePodName (controller pkg only)
	ClonerSourcePodName = "cdi-clone-source"
	// ClonerMountPath (controller pkg only)
	ClonerMountPath = "/var/run/cdi/clone/source"
	// ClonerSourcePodNameSuffix (controller pkg only)
	ClonerSourcePodNameSuffix = "-source-pod"

	// KubeVirtAnnKey is part of a kubevirt.io key.
	KubeVirtAnnKey = "kubevirt.io/"
	// CDIAnnKey is part of a kubevirt.io key.
	CDIAnnKey = "cdi.kubevirt.io/"

	// SmartClonerCDILabel is the label applied to resources created by the smart-clone controller
	SmartClonerCDILabel = "cdi-smart-clone"

	// UploadPodName (controller pkg only)
	UploadPodName = "cdi-upload"
	// UploadServerCDILabel is the label applied to upload server resources
	UploadServerCDILabel = "cdi-upload-server"
	// UploadServerPodname is name of the upload server pod container
	UploadServerPodname = UploadServerCDILabel
	// UploadServerDataDir is the destination directoryfor uploads
	UploadServerDataDir = ImporterDataDir
	// UploadServerServiceLabel is the label selector for upload server services
	UploadServerServiceLabel = "service"
	// UploadImageSize provides a constant to capture our env variable "UPLOAD_IMAGE_SIZE"
	UploadImageSize = "UPLOAD_IMAGE_SIZE"

	// FilesystemOverheadVar provides a constant to capture our env variable "FILESYSTEM_OVERHEAD"
	FilesystemOverheadVar = "FILESYSTEM_OVERHEAD"
	// DefaultGlobalOverhead is the amount of space reserved on Filesystem volumes by default
	DefaultGlobalOverhead = "0.055"

	// ConfigName is the name of default CDI Config
	ConfigName = "config"

	// OwnerUID provides the UID of the owner entity (either PVC or DV)
	OwnerUID = "OWNER_UID"

	// KeyAccess provides a constant to the accessKeyId label using in controller pkg and transport_test.go
	KeyAccess = "accessKeyId"
	// KeySecret provides a constant to the secretKey label using in controller pkg and transport_test.go
	KeySecret = "secretKey"

	// DefaultResyncPeriod sets a 10 minute resync period, used in the controller pkg and the controller cmd executable
	DefaultResyncPeriod = 10 * time.Minute
	// InsecureRegistryConfigMap is the name of the ConfigMap for insecure registries
	InsecureRegistryConfigMap = "cdi-insecure-registries"

	// ScratchSpaceNeededExitCode is the exit code that indicates the importer pod requires scratch space to function properly.
	ScratchSpaceNeededExitCode = 42

	// ScratchNameSuffix (controller pkg only)
	ScratchNameSuffix = "scratch"

	// UploadTokenIssuer is the JWT issuer of upload tokens
	UploadTokenIssuer = "cdi-apiserver"

	// CloneTokenIssuer is the JWT issuer for clone tokens
	CloneTokenIssuer = "cdi-apiserver"

	// ExtendedCloneTokenIssuer is the JWT issuer for clone tokens
	ExtendedCloneTokenIssuer = "cdi-deployment"

	// QemuSubGid is the gid used as the qemu group in fsGroup
	QemuSubGid = int64(107)

	// ControllerServiceAccountName is the name of the CDI controller service account
	ControllerServiceAccountName = "cdi-sa"

	// VddkConfigMap is the name of the ConfigMap with a reference to the VDDK image
	VddkConfigMap = "v2v-vmware"
	// VddkConfigDataKey is the name of the ConfigMap key of the VDDK image reference
	VddkConfigDataKey = "vddk-init-image"
	// AwaitingVDDK is a Pending condition reason that indicates the PVC is waiting for a VDDK image
	AwaitingVDDK = "AwaitingVDDK"

	// UploadContentTypeHeader is the header upload clients may use to set the content type explicitly
	UploadContentTypeHeader = "x-cdi-content-type"

	// FilesystemCloneContentType is the content type when cloning a filesystem
	FilesystemCloneContentType = "filesystem-clone"

	// BlockdeviceClone is the content type when cloning a block device
	BlockdeviceClone = "blockdevice-clone"

	// UploadPathSync is the path to POST CDI uploads
	UploadPathSync = "/v1beta1/upload"

	// UploadPathAsync is the path to POST CDI uploads in async mode
	UploadPathAsync = "/v1beta1/upload-async"

	// UploadArchivePath is the path to POST CDI archive uploads
	UploadArchivePath = "/v1beta1/upload-archive"

	// UploadArchiveAlphaPath is the path to POST CDI alpha archive uploads
	UploadArchiveAlphaPath = "/v1alpha1/upload-archive"

	// UploadFormSync is the path to POST CDI uploads as form data
	UploadFormSync = "/v1beta1/upload-form"

	// UploadFormAsync is the path to POST CDI uploads as form data in async mode
	UploadFormAsync = "/v1beta1/upload-form-async"

	// PreallocationApplied is a string inserted into importer's/uploader's exit message
	PreallocationApplied = "Preallocation applied"

	// SecretHeader is the key in a secret containing a sensitive extra header for HTTP data sources
	SecretHeader = "secretHeader"

	// UnusualRestartCountThreshold is the number of pod restarts that we consider unusual and would like to alert about
	UnusualRestartCountThreshold = 3

	// CDIControllerLeaderElectionHelperName is the name of the configmap that is used as a helper for controller leader election
	CDIControllerLeaderElectionHelperName = "cdi-controller-leader-election-helper"
)

// ProxyPaths are all supported paths
var ProxyPaths = append(
	append(SyncUploadPaths, AsyncUploadPaths...),
	append(SyncUploadFormPaths, AsyncUploadFormPaths...)...,
)

// SyncUploadPaths are paths to POST CDI uploads
var SyncUploadPaths = []string{
	UploadPathSync,
	"/v1alpha1/upload",
}

// AsyncUploadPaths are paths to POST CDI uploads in async mode
var AsyncUploadPaths = []string{
	UploadPathAsync,
	"/v1alpha1/upload-async",
}

// ArchiveUploadPaths are paths to POST CDI uploads of archive
var ArchiveUploadPaths = []string{
	UploadArchivePath,
	UploadArchiveAlphaPath,
}

// SyncUploadFormPaths are paths to POST CDI uploads as form data
var SyncUploadFormPaths = []string{
	UploadFormSync,
	"/v1alpha1/upload-form",
}

// AsyncUploadFormPaths are paths to POST CDI uploads as form data in async mode
var AsyncUploadFormPaths = []string{
	UploadFormAsync,
	"/v1alpha1/upload-form-async",
}
