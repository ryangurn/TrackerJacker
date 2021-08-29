package windows

import (
	"encoding/json"
	wapi "github.com/iamacarpet/go-win64api"
	"strconv"
	"time"
)

func SoftwareExist(software string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			retBool = true
			if out, err := json.Marshal(v); err == nil {
				return retBool, string(out)
			}
			return
		}
	}

	return
}

func SoftwareArch(software string, arch string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.Arch == arch {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwarePublisher(software string, publisher string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.Publisher == publisher {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareInstall(software string, date string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			val, err := time.Parse("2016-01-02 15:04", date)
			if err != nil {
				return
			}

			if v.InstallDate == val {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareEstimatedSize(software string, size string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			val, err := strconv.ParseUint(size, 10, 64)
			if err != nil {
				return
			}

			if v.EstimatedSize == uint64(val) {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareContact(software string, contact string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.Contact == contact {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareHelplink(software string, link string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.HelpLink == link {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareInstallSource(software string, source string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.InstallSource == source {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareInstallLocation(software string, location string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.InstallLocation == location {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareMajorVersion(software string, version string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			val, err := strconv.ParseUint(version, 10, 64)
			if err != nil {
				return
			}

			if v.VersionMajor == uint64(val) {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}

func SoftwareMinorVersion(software string, version string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			val, err := strconv.ParseUint(version, 10, 64)
			if err != nil {
				return
			}

			if v.VersionMajor == uint64(val) {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
				return
			}
		}
	}

	return
}