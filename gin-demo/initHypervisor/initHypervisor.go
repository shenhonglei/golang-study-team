package initHypervisor

import (
	"context"
	"gin-demo/settings"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"log"
	"net/url"
)

func Conn() *vim25.Client {
	u, err := soap.ParseURL(settings.AppConfig.Hypervisor.Url)
	ctx,_ := context.WithCancel(context.Background())
	u.User = url.UserPassword(settings.AppConfig.Hypervisor.Username, settings.AppConfig.Hypervisor.Password)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	client, err := govmomi.NewClient(ctx, u, settings.AppConfig.Hypervisor.Insecure)
	if err != nil {
		panic(err)
	}
	log.Println(client)
	return client.Client

}
