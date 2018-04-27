package cmd

import (
	"fmt"
	"net/http"
	"encoding/json"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/pkg/errors"
	"strings"
)

// Response is the structure of the main json response.
type Response struct {
	SyncToken  string   `json:"syncToken"`
	CreateDate string   `json:"createDate"`
	Prefixes   []Prefix `json:"prefixes"`
}

// Prefix is the substructure of the json response.
type Prefix struct {
	Prefix  string `json:"ip_prefix"`
	Region  string `json:"region"`
	Service string `json:"service"`
}

// CloudIpRange is the structure of the json output.
type CloudIpRange struct {
	CIDR     string `json:"cidr"`
	Provider string `json:"provider"`
	Region   string `json:"region"`
	Service  string `json:"service"`
}

type cmdGet struct {
	Providers []string
	Regions   []string
	Services  []string
}

const AwsEndpoint = "https://ip-ranges.amazonaws.com/ip-ranges.json"

func (cmd *cmdGet) run(c *kingpin.ParseContext) error {
	filteredIpRanges := []CloudIpRange{}

	for _, provider := range cmd.Providers {
		switch provider {
		case "aws":
			break;
		default:
			return errors.New("Provider not currently supported")
		}
	}

	var resp Response
	r, err := http.Get(AwsEndpoint)
	if err != nil {
		return errors.Wrap(err, "Failed to get ip ranges")
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return errors.Wrap(err, "Failed to decode ip ranges")
	}

	// Loop through AWS responses.
	for _, r := range resp.Prefixes {
		var serviceIncluded bool = false
		var regionIncluded bool = false

		if len(cmd.Services) >= 1 {
			for _, service := range cmd.Services {
				if strings.ToLower(service) == strings.ToLower(r.Service) {
					serviceIncluded = true
				}
			}
		} else {
			serviceIncluded = true
		}

		if serviceIncluded == false {
			continue
		}

		if len(cmd.Regions) >= 1 {
			for _, region := range cmd.Regions {
				if strings.ToLower(region) == strings.ToLower(r.Region) {
					regionIncluded = true
				}
			}
		} else {
			regionIncluded = true
		}

		if regionIncluded == false {
			continue
		}

		filteredIpRanges = append(filteredIpRanges, CloudIpRange{
			CIDR: r.Prefix,
			Region: r.Region,
			Service: r.Service,
			Provider: "aws",
		})
	}

	json, _ := json.MarshalIndent(filteredIpRanges, "", "  ")

	fmt.Println(string(json))

	return nil
}

// Get declares the "get" sub command.
func Get(app *kingpin.Application) {
	cmd := new(cmdGet)
	g := app.Command("get", fmt.Sprintf("Fetches IP ranges for given set of constraints", app.Name)).Action(cmd.run)

	g.Flag("provider", "name of cloud provider(s)").Default("aws").StringsVar(&cmd.Providers)
	g.Flag("region", "region names of cloud providers (returns all results if omitted)").StringsVar(&cmd.Regions)
	g.Flag("service", "name of service (returns all results if omitted)").StringsVar(&cmd.Services)
}
