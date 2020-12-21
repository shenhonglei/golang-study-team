package models

import (
	"context"
	"gin-demo/initHypervisor"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"reflect"
)

type VmSummary struct {
	Id	string `json:"uuid,string"`
	Name string `json:"vmname,string"`
	OSType string `json:"guestfullname,string"`
}

func (mysummary VmSummary)QueryVM(vmname string) VmSummary { //通过vmname传值
	ctx,_ := context.WithCancel(context.Background())
	c := initHypervisor.Conn()
	m := view.NewManager(c)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		panic(err)
	}

	// Retrieve summary property for all machines
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		panic(err)
	}
	// Print summary per vm (see also: govc/vm/info.go)
	for _, vm := range vms {
		// 判断虚拟机名称是否相同，相同的话，vm 就是查找到主机
		if vm.Summary.Config.Name == vmname {
			mysummary.Name = reflect.ValueOf(vm.Summary.Config.Name).String()
			mysummary.Id = reflect.ValueOf(vm.Summary.Config.Uuid).String()
			mysummary.OSType = reflect.ValueOf(vm.Summary.Config.GuestFullName).String()
			break
		}
	}
	return mysummary
}