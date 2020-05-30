package features

import "fmt"

type SBCPlatform interface {
	getConfig() string
}

type AWS struct {
	isCloud   bool
	diskSpace int
}

type KVM struct {
	isNto1Supported bool
}

type Openstack struct {
	isAutomatedUpgradeSupported bool
	template                    string
}

func (aws *AWS) getConfig() string {
	return "AWS config"
}

func (kvm *KVM) getConfig() string {
	return "KVM config"
}

func (openstack *Openstack) getConfig() string {
	return "Openstack config"
}

func loadConfig(platform SBCPlatform) {
	res := platform.getConfig()
	fmt.Printf("Loading %s\n", res)
}

func InterfaceDemo() {
	aws := AWS{isCloud: true, diskSpace: 40}
	kvm := KVM{isNto1Supported: false}
	openstack := Openstack{isAutomatedUpgradeSupported: true, template: "heatRgActive.yaml"}

	loadConfig(&aws)
	loadConfig(&kvm)
	loadConfig(&openstack)
}
