package framework

import (
	"strings"

	"github.com/google/uuid"
	"github.com/onsi/gomega"

	cdiv1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
	"kubevirt.io/containerized-data-importer/pkg/common"
	"kubevirt.io/containerized-data-importer/tests/utils"
)

// CreateVddkWarmImportDataVolume fetches snapshot information from vcsim and returns a multi-stage VDDK data volume
func (f *Framework) CreateVddkWarmImportDataVolume(dataVolumeName, size, url string) *cdiv1.DataVolume {
	// Find vcenter-simulator pod
	pod, err := utils.FindPodByPrefix(f.K8sClient, f.CdiInstallNs, "vcenter-deployment", "app=vcenter")
	gomega.Expect(err).ToNot(gomega.HaveOccurred())
	gomega.Expect(pod).ToNot(gomega.BeNil())

	// Get test VM UUID
	id, err := f.RunKubectlCommand("exec", "-n", pod.Namespace, pod.Name, "--", "cat", "/tmp/vmid")
	gomega.Expect(err).To(gomega.BeNil())
	vmid, err := uuid.Parse(strings.TrimSpace(id))
	gomega.Expect(err).To(gomega.BeNil())

	// Get snapshot 1 ID
	previousCheckpoint, err := f.RunKubectlCommand("exec", "-n", pod.Namespace, pod.Name, "--", "cat", "/tmp/vmsnapshot1")
	gomega.Expect(err).To(gomega.BeNil())
	previousCheckpoint = strings.TrimSpace(previousCheckpoint)
	gomega.Expect(err).To(gomega.BeNil())

	// Get snapshot 2 ID
	currentCheckpoint, err := f.RunKubectlCommand("exec", "-n", pod.Namespace, pod.Name, "--", "cat", "/tmp/vmsnapshot2")
	gomega.Expect(err).To(gomega.BeNil())
	currentCheckpoint = strings.TrimSpace(currentCheckpoint)
	gomega.Expect(err).To(gomega.BeNil())

	// Get disk name
	disk, err := f.RunKubectlCommand("exec", "-n", pod.Namespace, pod.Name, "--", "cat", "/tmp/vmdisk")
	gomega.Expect(err).To(gomega.BeNil())
	disk = strings.TrimSpace(disk)
	gomega.Expect(err).To(gomega.BeNil())

	// Create VDDK login secret
	stringData := map[string]string{
		common.KeyAccess: "user",
		common.KeySecret: "pass",
	}
	backingFile := disk
	secretRef := "vddksecret"
	thumbprint := "testprint"
	finalCheckpoint := true
	s, _ := utils.CreateSecretFromDefinition(f.K8sClient, utils.NewSecretDefinition(nil, stringData, nil, f.Namespace.Name, secretRef))

	return utils.NewDataVolumeWithVddkWarmImport(dataVolumeName, size, backingFile, s.Name, thumbprint, url, vmid.String(), currentCheckpoint, previousCheckpoint, finalCheckpoint)
}
