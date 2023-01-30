package common

import (
	"log"
	"strconv"

	oscgo "github.com/outscale/osc-sdk-go/v2"
)

func buildOscNetFilters(input map[string]string) oscgo.FiltersNet {
	var filters oscgo.FiltersNet
	for k, v := range input {
		filterValue := []string{v}
		switch name := k; name {
		case "ip-range":
			filters.SetIpRanges(filterValue)
		case "dhcp-options-set-id":
			filters.SetDhcpOptionsSetIds(filterValue)
		case "is-default":
			if isDefault, err := strconv.ParseBool(v); err == nil {
				filters.SetIsDefault(isDefault)
			}
		case "state":
			filters.SetStates(filterValue)
		case "tag-key":
			filters.SetTagKeys(filterValue)
		case "tag-value":
			filters.SetTagValues(filterValue)
		default:
			log.Printf("[Debug] Unknown Filter Name: %s.", name)
		}
	}
	return filters
}

func buildOscSubnetFilters(input map[string]string) oscgo.FiltersSubnet {
	var filters oscgo.FiltersSubnet
	for k, v := range input {
		filterValue := []string{v}
		switch name := k; name {
		case "available-ips-counts":
			if ipCount, err := strconv.Atoi(v); err == nil {
				filters.SetAvailableIpsCounts([]int32{int32(ipCount)})
			}
		case "ip-ranges":
			filters.SetIpRanges(filterValue)
		case "net-ids":
			filters.SetNetIds(filterValue)
		case "states":
			filters.SetStates(filterValue)
		case "subnet-ids":
			filters.SetSubnetIds(filterValue)
		case "sub-region-names":
			filters.SetSubregionNames(filterValue)
		default:
			log.Printf("[Debug] Unknown Filter Name: %s.", name)
		}
	}
	return filters
}

func buildOSCOMIFilters(input map[string]string) oscgo.FiltersImage {
	var filters oscgo.FiltersImage
	for k, v := range input {
		filterValue := []string{v}

		switch name := k; name {
		case "account-alias":
			filters.SetAccountAliases(filterValue)
		case "account-id":
			filters.SetAccountIds(filterValue)
		case "architecture":
			filters.SetArchitectures(filterValue)
		case "image-id":
			filters.SetImageIds(filterValue)
		case "image-name":
			filters.SetImageNames(filterValue)
		// case "image-type":
		// 	filters.ImageTypes = filterValue
		case "virtualization-type":
			filters.SetVirtualizationTypes(filterValue)
		case "root-device-type":
			filters.SetRootDeviceTypes(filterValue)
		// case "block-device-mapping-volume-type":
		// 	filters.BlockDeviceMappingVolumeType = filterValue
		//Some params are missing.
		default:
			log.Printf("[WARN] Unknown Filter Name: %s.", name)
		}
	}
	return filters
}
