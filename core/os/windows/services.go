 package windows

 import (
	 "encoding/json"
	 "github.com/bugsnag/bugsnag-go"
	 wapi "github.com/iamacarpet/go-win64api"
	 "os"
	 "strconv"
 )

func ServiceExist(svc string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	services, err := wapi.GetServices()
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

	for _, s := range services {
		if s.SCName == svc {
			retBool = true
			if out, err := json.Marshal(s); err == nil {
				return retBool, string(out)
			}
		}
	}

	out, _ := json.Marshal(services)
	return retBool, string(out)
}

func ServiceDisplayName(svc string, display string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	services, err := wapi.GetServices()
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

	for _, s := range services {
		if s.SCName == svc {
			if s.DisplayName == display {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(services)
	return retBool, string(out)
}

func ServiceStatusText(svc string, text string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	services, err := wapi.GetServices()
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

	for _, s := range services {
		if s.SCName == svc {
			if s.StatusText == text {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(services)
	return retBool, string(out)
}

func ServiceStatus(svc string, status string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	services, err := wapi.GetServices()
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

	val, err := strconv.ParseUint(status, 10, 32)
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

	for _, s := range services {
		if s.SCName == svc {
			if s.Status == uint32(val) {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(services)
	return retBool, string(out)
}

func ServiceAcceptStop(svc string, stop string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	services, err := wapi.GetServices()
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

	val, err := strconv.ParseBool(stop)
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

	for _, s := range services {
		if s.SCName == svc {
			if s.AcceptStop == val {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(services)
	return retBool, string(out)
}

func ServiceRunning(svc string, running string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	services, err := wapi.GetServices()
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

	val, err := strconv.ParseBool(running)
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

	for _, s := range services {
		if s.SCName == svc {
			if s.IsRunning == val {
				retBool = true
				if out, err := json.Marshal(s); err == nil {
					return retBool, string(out)
				}
			}
		}
	}

	out, _ := json.Marshal(services)
	return retBool, string(out)
}

func ServiceRunningPid(svc string, pid string) (retBool bool, retData string) {
	retBool = false
	retData = ""

	services, err := wapi.GetServices()
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

	val, err := strconv.ParseUint(pid, 10, 64)
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

 	for _, s := range services {
 		if s.SCName == svc {
 			if s.RunningPid == uint32(val) {
			 	retBool = true
			 	if out, err := json.Marshal(s); err == nil {
			 		return retBool, string(out)
			 	}
		 	}
	 	}
 	}

	out, _ := json.Marshal(services)
	return retBool, string(out)
}

 func ServiceType(svc string, typ string) (retBool bool, retData string) {
	 retBool = false
	 retData = ""

	 services, err := wapi.GetServices()
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

	 val, err := strconv.ParseUint(typ, 10, 64)
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

	 for _, s := range services {
		 if s.SCName == svc {
			 if s.ServiceType == uint32(val) {
				 retBool = true
				 if out, err := json.Marshal(s); err == nil {
					 return retBool, string(out)
				 }
			 }
		 }
	 }

	 out, _ := json.Marshal(services)
	 return retBool, string(out)
 }