package main

import (
	"fmt"
	"sort"
	"strings"

	"kubevirt.io/containerized-data-importer/pkg/monitoring"
)

// constant parts of the file
const (
	title      = "# Containerized Data Importer metrics\n"
	background = "This document aims to help users that are not familiar with metrics exposed by the Containerized Data Importer.\n" +
		"All metrics documented here are auto-generated by the utility tool `tools/metricsdocs` and reflects exactly what is being exposed.\n\n"

	KVSpecificMetrics = "## Containerized Data Importer Metrics List\n"

	opening = title +
		background +
		KVSpecificMetrics

	// footer
	footerHeading = "## Developing new metrics\n"
	footerContent = "After developing new metrics or changing old ones, please run `make generate-doc` to regenerate this document.\n\n" +
		"If you feel that the new metric doesn't follow these rules, please change `tools/metricsdocs` with your needs.\n"

	footer = footerHeading + footerContent
)

func main() {
	metricsList := recordRulesDescToMetricList(monitoring.GetRecordRulesDesc(""))
	for _, opts := range monitoring.MetricOptsList {
		metricsList = append(metricsList, opts)
	}

	sort.Slice(metricsList, func(i, j int) bool {
		return metricsList[i].Name < metricsList[j].Name
	})

	writeToFile(metricsList)
}

func writeToFile(metricsList metricList) {
	fmt.Print(opening)
	metricsList.writeOut()
	fmt.Print(footer)
}

func recordRulesDescToMetricList(mdl []monitoring.RecordRulesDesc) metricList {
	res := make([]monitoring.MetricOpts, len(mdl))
	for i, md := range mdl {
		res[i] = metricDescriptionToMetric(md)
	}

	return res
}

func metricDescriptionToMetric(rrd monitoring.RecordRulesDesc) monitoring.MetricOpts {
	return monitoring.MetricOpts{
		Name: rrd.Opts.Name,
		Help: rrd.Opts.Help,
	}
}

func writeOut(m monitoring.MetricOpts) {
	fmt.Println("###", m.Name)
	fmt.Println(m.Help)
}

type metricList []monitoring.MetricOpts

// Len implements sort.Interface.Len
func (m metricList) Len() int {
	return len(m)
}

// Less implements sort.Interface.Less
func (m metricList) Less(i, j int) bool {
	return m[i].Name < m[j].Help
}

// Swap implements sort.Interface.Swap
func (m metricList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *metricList) add(line string) {
	split := strings.Split(line, " ")
	name := split[2]
	split[3] = strings.Title(split[3])
	description := strings.Join(split[3:], " ")
	*m = append(*m, monitoring.MetricOpts{Name: name, Help: description})
}

func (m metricList) writeOut() {
	for _, met := range m {
		writeOut(met)
	}
}
