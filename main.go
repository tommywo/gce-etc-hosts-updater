package main
import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
  "log"
  "github.com/lextoumbourou/goodhosts"
  "time"
)

func main() {
        // Use oauth2.NoContext if there isn't a good context to pass in.
        ctx := oauth2.NoContext



        client, err := google.DefaultClient(ctx, compute.ComputeScope)
        if err != nil {
          log.Fatal(err)
        }
        service, err := compute.New(client)
        if err != nil {
          log.Fatal(err)
        }

        projectId := "api-project-773019699287"
        zone:= "europe-west1-d"

        // Show the current images that are available.


        for {
            res, err := service.Instances.List(projectId, zone).Do()
            if err != nil {
                log.Fatal(err)
            }
            instances := res.Items
            for _,element := range instances {
                hosts, err := goodhosts.NewHosts()
                if err != nil {
                    log.Fatal(err)
                }
                if hosts.Has(element.NetworkInterfaces[0].NetworkIP, element.Name) {
                    continue
                } else {
                  // Note that nothing will be added to the hosts file until ``hosts.Flush`` is called.
                  hosts.Add(element.NetworkInterfaces[0].NetworkIP, element.Name)
                  if err := hosts.Flush(); err != nil {
                      panic(err)
                  }
                }
            }
            time.Sleep(3*time.Second)
        }


}
