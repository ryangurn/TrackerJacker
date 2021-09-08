package windows

import (
	"encoding/json"
	"github.com/bugsnag/bugsnag-go"
	wapi "github.com/iamacarpet/go-win64api"
	"os"
	"strconv"
	"time"
)

func SoftwareExist(software string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			retBool = true
			if out, err := json.Marshal(v); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareArch(software string, arch string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.Arch == arch {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwarePublisher(software string, publisher string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.Publisher == publisher {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareInstall(software string, date string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := time.Parse("2006-01-02 15:04", date)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.InstallDate == val {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareEstimatedSize(software string, size string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := strconv.ParseUint(size, 10, 64)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.EstimatedSize == uint64(val) {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareContact(software string, contact string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.Contact == contact {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareHelplink(software string, link string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.HelpLink == link {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareInstallSource(software string, source string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.InstallSource == source {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareInstallLocation(software string, location string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.InstallLocation == location {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareMajorVersion(software string, version string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := strconv.ParseUint(version, 10, 64)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.VersionMajor == uint64(val) {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}

func SoftwareMinorVersion(software string, version string) (retBool bool, retData string){
	retBool = false

	softwares, err := wapi.InstalledSoftwareList()
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	val, err := strconv.ParseUint(version, 10, 64)
	if err != nil {
		bugsnag.Notify(err, bugsnag.HandledState{
			SeverityReason:   bugsnag.SeverityReasonHandledError,
			OriginalSeverity: bugsnag.SeverityWarning,
			Unhandled:      false,
		}, bugsnag.MetaData{
			"ENV": {
				"AUTH_TOKEN": os.Getenv("AUTH_TOKEN"),
				"BUGSNAG_KEY": os.Getenv("BUGSNAG_KEY"),
				"IMAGE": os.Getenv("IMAGE"),
				"SCORING_METHOD": os.Getenv("SCORING_METHOD"),
				"SERVER": os.Getenv("SERVER"),
			},
		})
		return
	}

	for _, v := range softwares {
		if v.DisplayName == software {
			if v.VersionMajor == uint64(val) {
				retBool = true
				if out, err := json.Marshal(v); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(softwares)
	return retBool, string(out)
}