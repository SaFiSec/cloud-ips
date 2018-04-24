package cmd

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
	"html/template"
	"time"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/pkg/errors"
	"github.com/brotherpowers/ipsubnet"

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

// RouteItem is the struct passed to the template.
type RouteItem struct {
	CIDR    string
	Mask    string
	Comment string
}

// TemplateData is the struct passed to the template.
type TemplateData struct {
	Timestamp  string
	RouteItems []RouteItem
}

type cmdRender struct {
	TemplatePath string
	OutputPath   string
	EndpointURL  string
}

func (cmd *cmdRender) run(c *kingpin.ParseContext) error {
	fmt.Println("Fetching list of CloudFront edge location CIDRs")

	var resp Response
	r, err := http.Get(cmd.EndpointURL)
	if err != nil {
		return errors.Wrap(err, "Failed to get ip ranges")
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		return errors.Wrap(err, "Failed to decode ip ranges")
	}

	fmt.Println(fmt.Sprintf("Loading template %s", cmd.TemplatePath))
	now := time.Now()
	data := TemplateData{
		Timestamp: now.Format(time.RFC822Z),
	}
	for _, r := range resp.Prefixes {
		if r.Service != "CLOUDFRONT" {
			continue
		}

		parts := strings.Split(r.Prefix, "/")
		mask, _ := strconv.Atoi(parts[1])
		sub := ipsubnet.SubnetCalculator(parts[0], mask)

		data.RouteItems = append(data.RouteItems, RouteItem{
			CIDR: r.Prefix,
			Mask: sub.GetSubnetMask(),
			Comment: fmt.Sprintf("%s - %s", r.Service, r.Region),
		})
	}

	fmt.Println(fmt.Sprintf("Rendering output %s", cmd.OutputPath))
	_, err = os.Stat(cmd.OutputPath)
	if os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("File doesn't exist - creating", cmd.OutputPath))
		var file, _ = os.Create(cmd.OutputPath)
		defer file.Close()
	}

	w, err := os.OpenFile(cmd.OutputPath, os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, "Could not open output file")
	}
	defer w.Close()

	tmpl := template.Must(template.ParseFiles(cmd.TemplatePath))
	tmpl.Execute(w, data)

	return nil
}

// Render declares the "version" sub command.
func Render(app *kingpin.Application) {
	cmd := new(cmdRender)
	r := app.Command("render", fmt.Sprintf("Takes a template and renders to target file", app.Name)).Action(cmd.run)
	r.Flag("template", "path to template file").Required().StringVar(&cmd.TemplatePath)
	r.Flag("output", "path to output file").Required().StringVar(&cmd.OutputPath)
	r.Flag("endpoint", "URL of cloudfront edge location JSON feed").Default("https://ip-ranges.amazonaws.com/ip-ranges.json").Hidden().StringVar(&cmd.EndpointURL)
}