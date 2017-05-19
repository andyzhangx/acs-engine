package armhelpers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
)

// ResourceName returns the last segment (the resource name) for the specified resource identifier.
func ResourceName(ID string) (string, error) {
	parts := strings.Split(ID, "/")
	name := parts[len(parts)-1]
	if len(name) == 0 {
		return "", fmt.Errorf("resource name was missing from identifier")
	}

	return name, nil
}

// SplitBlobURI returns a decomposed blob URI parts: accountName, containerName, blobName.
func SplitBlobURI(URI string) (string, string, string, error) {
	uri, err := url.Parse(URI)
	if err != nil {
		return "", "", "", err
	}

	accountName := strings.Split(uri.Host, ".")[0]
	urlParts := strings.Split(uri.Path, "/")

	containerName := urlParts[1]
	blobPath := strings.Join(urlParts[2:], "/")

	return accountName, containerName, blobPath, nil
}

// VMNameParts returns parts of Linux VM name e.g: k8s-agentpool1-11290731-0
func VMNameParts(vmName string) (string, string, string, int, error) {
	vmNameParts := strings.Split(vmName, "-")

	if len(vmNameParts) != 4 {
		return "", "", "", -1, fmt.Errorf("resource name was missing from identifier")
	}

	vmNum, err := strconv.Atoi(vmNameParts[3])

	if err != nil {
		return "", "", "", -1, fmt.Errorf("Error parsing VM Name: %v", err)
	}

	return vmNameParts[0], vmNameParts[1], vmNameParts[2], vmNum, nil
}

// ByVMNameOffset implements sort.Interface for []VirtualMachine based on
// the Name field.
type ByVMNameOffset []compute.VirtualMachine

func (vm ByVMNameOffset) Len() int      { return len(vm) }
func (vm ByVMNameOffset) Swap(i, j int) { vm[i], vm[j] = vm[j], vm[i] }
func (vm ByVMNameOffset) Less(i, j int) bool {
	vm1NameParts := strings.Split(*vm[i].Name, "-")
	vm1Num := vm1NameParts[len(vm1NameParts)-1]

	vm2NameParts := strings.Split(*vm[j].Name, "-")
	vm2Num := vm2NameParts[len(vm2NameParts)-1]

	return vm1Num < vm2Num
}
