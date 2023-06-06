package nginx

type InstanceSpec struct {
	Revision      int
	State         string
	ConfigVersion string
}

type InstanceStatus struct {
	State         string
	ConfigVersion string
}

type Instance struct {
	ID     string
	Spec   *InstanceSpec
	Status *InstanceStatus
}
