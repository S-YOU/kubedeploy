package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"k8s.io/kubernetes/pkg/api"
)

func printPodsTable(pods []api.Pod) {
	color.Red("=== Pod ===")
	data := [][]string{}
	for _, pod := range pods {
		data = append(data, []string{
			pod.Name,
			pod.Spec.Containers[0].Image,
			pod.Namespace,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Name",
		"Image",
		"Namespace",
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	fmt.Println()
}

func printServicesTable(services []api.Service) {
	color.Blue("=== Service ===")
	data := [][]string{}
	for _, service := range services {
		data = append(data, []string{
			service.Name,
			service.Namespace,
			service.Spec.Selector["color"],
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Name",
		"Namespace",
		"Color",
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	fmt.Println()

}

func printReplace(old, new string) {
	fmt.Println("Replace: " + old + " => " + new)
}