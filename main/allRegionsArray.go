package main

import "github.com/aws/aws-sdk-go/aws/endpoints"

var allRegions = [18]string{
	endpoints.ApEast1RegionID,      // Asia Pacific (Hong Kong).
	endpoints.ApNortheast1RegionID, // Asia Pacific (Tokyo).
	endpoints.ApNortheast2RegionID, // Asia Pacific (Seoul).
	endpoints.ApSouth1RegionID,     // Asia Pacific (Mumbai).
	endpoints.ApSoutheast1RegionID, // Asia Pacific (Singapore).
	endpoints.ApSoutheast2RegionID, // Asia Pacific (Sydney).
	endpoints.CaCentral1RegionID,   // Canada (Central).
	endpoints.EuCentral1RegionID,   // EU (Frankfurt).
	endpoints.EuNorth1RegionID,     // EU (Stockholm).
	endpoints.EuWest1RegionID,      // EU (Ireland).
	endpoints.EuWest2RegionID,      // EU (London).
	endpoints.EuWest3RegionID,      // EU (Paris).
	endpoints.MeSouth1RegionID,     // Middle East (Bahrain).
	endpoints.SaEast1RegionID,      // South America (Sao Paulo).
	endpoints.UsEast1RegionID,      // US East (N. Virginia).
	endpoints.UsEast2RegionID,      // US East (Ohio).
	endpoints.UsWest1RegionID,      // US West (N. California).
	endpoints.UsWest2RegionID}      // US West (Oregon).
