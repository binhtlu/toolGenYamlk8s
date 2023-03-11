package helps

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
	apiv1 "k8s.io/api/core/v1"
)

type Endpoints struct {
	// end *apiv1.Endpoints
}

func (this *Endpoints) NewEndPoints(ip, name string) *apiv1.Endpoints {
	var tcp string
	tcp = "TCP"
	endpoint := &apiv1.Endpoints{}
	endpoint.TypeMeta.Kind = "Endpoints"
	endpoint.TypeMeta.APIVersion = "v1"
	endpoint.ObjectMeta.Namespace = "crypto-currency"
	endpoint.ObjectMeta.Name = name
	endpoint.Subsets = []apiv1.EndpointSubset{
		{
			Addresses: []apiv1.EndpointAddress{
				{
					IP: ip,
				},
			},
			Ports: []apiv1.EndpointPort{
				{

					Port:        18143,
					Protocol:    apiv1.ProtocolTCP,
					AppProtocol: &tcp,
					Name:        name,
				},
			},
		},
	}
	return endpoint

}

func (this *Endpoints) GenFile(filepath string) {
	var ip, name string
	var end *apiv1.Endpoints
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		a := findString(scanner.Text(), "#", len(scanner.Text()))

		if a == false {
			s := strings.Join(strings.Fields(scanner.Text()), " ")
			ip, name = HandleString(s)
			if ip != "" && name != "" {
				end = this.NewEndPoints(ip, name)

				bData, err := yaml.Marshal(end)
				if err != nil {
					os.Exit(1)
				}

				err1 := ioutil.WriteFile(name+".yaml", bData, 0777)
				if err1 != nil {
					fmt.Println(err)
				}
			}
		}
	}

}
