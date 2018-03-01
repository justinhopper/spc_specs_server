package spc_app

import (
    "fmt"
    "net/http"
    "encoding/json"
)

const base_path = "/specs"

type provider struct {
    Name     string   `json:"Name"`
    Label    string   `json:"Label"`
    Machines []string `json:"Machines"`
}
var provider_specs = map[string]provider{
                            "aws": provider{
                                Name: "aws", 
                                Label: "AWS",
                                Machines: []string{
                                        "t2.medium",
                                        "t2.large",
                                        "t2.xlarge",
                                        "t2.2xlarge",
                                        "c4.large",
                                        "c4.xlarge",
                                        "c4.2xlarge",
                                        "c4.4xlarge",
                                        "c4.8xlarge",
                                        "c5.large",
                                        "c5.xlarge",
                                        "c5.2xlarge",
                                        "c5.4xlarge",
                                        "c5.9xlarge",
                                        "r4.large",
                                        "r4.xlarge",
                                        "r4.2xlarge",
                                        "r4.4xlarge",
                                        "r4.8xlarge",
                                        "r4.16xlarge",
                                        "m3.large",
                                        "m3.xlarge",
                                        "m3.2xlarge",
                                        "m4.large",
                                        "m4.xlarge",
                                        "m4.2xlarge",
                                        "m4.4xlarge",
                                        "m4.10xlarge",
                                        "m4.16xlarge",
                                        "m5.large",
                                        "m5.xlarge",
                                        "m5.2xlarge",
                                        "m5.4xlarge",
                                        "m5.12xlarge",
                                        "m5.24xlarge" },
                                },
                            "azure": provider{
                                Name: "azure",
                                Label: "Azure",
                                Machines: []string{
                                        "standard_a2",
                                        "standard_a3",
                                        "standard_a4",
                                        "standard_a5",
                                        "standard_a6",
                                        "standard_a7",
                                        "standard_a8",
                                        "standard_a9",
                                        "standard_a10",
                                        "standard_a11",
                                        "basic_a2",
                                        "basic_a3",
                                        "basic_a4",
                                        "standard_a2_v2",
                                        "standard_a4_v2",
                                        "standard_a8_v2",
                                        "standard_a2m_v2",
                                        "standard_a4m_v2",
                                        "standard_a8m_v2",
                                        "standard_d1",
                                        "standard_d2",
                                        "standard_d3",
                                        "standard_d4",
                                        "standard_d11",
                                        "standard_d12",
                                        "standard_d13",
                                        "standard_d14",
                                        "standard_ds1",
                                        "standard_ds2",
                                        "standard_ds3",
                                        "standard_ds4",
                                        "standard_d1_v2",
                                        "standard_d2_v2",
                                        "standard_d3_v2",
                                        "standard_d4_v2",
                                        "standard_d5_v2",
                                        "standard_d11_v2",
                                        "standard_d12_v2",
                                        "standard_d13_v2",
                                        "standard_d14_v2",
                                        "standard_d15_v2",
                                        "standard_ds1_v2",
                                        "standard_ds2_v2",
                                        "standard_ds3_v2",
                                        "standard_ds4_v2",
                                        "standard_ds5_v2",
                                        "standard_ds11_v2",
                                        "standard_ds12_v2",
                                        "standard_ds13_v2",
                                        "standard_ds14_v2",
                                        "standard_ds15_v2",
                                        "standard_f1",
                                        "standard_f2",
                                        "standard_f4",
                                        "standard_f8",
                                        "standard_f16",
                                        "standard_f1s",
                                        "standard_f2s",
                                        "standard_f4s",
                                        "standard_f8s",
                                        "standard_f16s",
                                        "standard_g1",
                                        "standard_g2",
                                        "standard_g3",
                                        "standard_g4",
                                        "standard_g5",
                                        "standard_gs1",
                                        "standard_gs2",
                                        "standard_gs3",
                                        "standard_gs4",
                                        "standard_gs5" },
                                },
                            "do": provider{
                                Name: "do",
                                Label: "Digital Ocean",
                                Machines: []string{
                                        "2gb",
                                        "4gb",
                                        "8gb",
                                        "16gb",
                                        "32gb",
                                        "48gb",
                                        "64gb",
                                        "c-2",
                                        "c-4",
                                        "c-8",
                                        "c-16",
                                        "c-32",
                                        "m-16gb",
                                        "m-32gb",
                                        "m-64gb",
                                        "m-128gb",
                                        "m-224gb" },
                                },
                            "gce": provider{
                                Name: "gce",
                                Label: "GCE",
                                Machines: []string{
                                        "n1-standard-1",
                                        "n1-standard-2",
                                        "n1-standard-4",
                                        "n1-standard-8",
                                        "n1-standard-16",
                                        "n1-standard-32",
                                        "n1-standard-64",
                                        "n1-highcpu-2",
                                        "n1-highcpu-4",
                                        "n1-highcpu-8",
                                        "n1-highcpu-16",
                                        "n1-highcpu-32",
                                        "n1-highcpu-64",
                                        "n1-highmem-2",
                                        "n1-highmem-4",
                                        "n1-highmem-8",
                                        "n1-highmem-16",
                                        "n1-highmem-32",
                                        "n1-highmem-64" },
                                },
                            "gke": provider{
                                Name: "gke",
                                Label: "GKE",
                                Machines: []string{ "(no machine types)" },
                                },
                            "packet": provider{
                                Name: "packet",
                                Label: "Packet",
                                Machines: []string{ "(no machine types)" },
                                },
                            "oneandone": provider{
                                Name: "oneandone",
                                Label: "1&1",
                                Machines: []string{
                                        "m",
                                        "mx",
                                        "l",
                                        "xl",
                                        "xxl",
                                        "3xl",
                                        "4xl",
                                        "m_beta",
                                        "l_beta",
                                        "xl_beta",
                                        "xxl_beta",
                                        "3xl_beta",
                                        "4xl_beta" },
                                },
                            "profitbricks": provider{
                                Name: "profitbricks",
                                Label: "ProfitBricks",
                                Machines: []string{ "(no machine types)" },
                                },
                     }

func init() {
    for key := range provider_specs {
        // Make copy of key here, otherwise Go will re-use function instead of creating one for each key
        key2 := key
        http.HandleFunc(base_path + "/" + key2 + "/", 
            func(w http.ResponseWriter, req *http.Request) {
                fmt.Printf("Saw request (%v): %v\n", key2, req.URL.Path) 
	        w.Header().Set("Content-Type", "application/json")
                d, err := json.Marshal(provider_specs[key2]);
                if err != nil {
                    fmt.Printf("Error marshalling: %v\n", err)
                }
                w.Write(d)
            })
    }
}
