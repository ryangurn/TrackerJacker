# TrackerJacker
"Tracker jackers are genetically engineered wasps, conceived and created in the Capitol. They are genetically coded to attack anyone or anything that disturbs their nest."

## Questions
For questions please open up an issue in Github that has the question along with any related lines of code noted! This will aid in tracking and discussing questions.

## Major Todos
- [ ] ensure that all tests are using proper truthy values, this will need to use the provided getters that are utilized throughout the application already.
- [ ] develop a web-server that can be run from any port on the client
- [ ] ensure that capitalization errors will not cause issues in validation. i.e: shift everything to lowercase and just for good measure trim whitespace.
- [ ] add in description and category into the input.json file to provide feedback to user via a web interface
- [ ] utilize goroutines and channels to ensure proper parallel processing is taking place.
- [ ] implement optimization techniques when pulling data from powershell, table-driven design might help.

## Checks
### Cross
1. Files
2. Hosts
3. Users (Basic)

#### Files
* Existence
* Hash

**Todos**
- [ ] Add ability for text based questions from users, or decide if the app should host a mini server to provide http traffic to the user. (this will be included in the graphical user interface presented via the web-server.)

#### Hosts
* Host Existence
* Ip Existence

#### Users
* Existence

**Todos**
- [ ] Add meta checks for feature parody with Windows

### Windows
1. Users (Comprehensive)
2. Groups (Basic)
3. Software (Comprehensive)
4. Services (Comprehensive) 
5. Shares (Comprehensive)
6. Processes (Comprehensive)
7. Policies (Comprehensive)
8. Firewalls (Comprehensive)
9. Updates (Comprehensive)
10. Bitlocker (Basic)
11. Profiles (Planned)

#### Users
* Existence
* BadPasswordCount
* FullName
* IsAdmin
* IsEnabled
* IsLocked
* LastLogon
* NoChangePassword
* NumberOfLogons
* PasswordAge
* PasswordNeverExpires

**Todos**
- [ ] add sessions into the checks for users on windows [documentation](https://github.com/iamacarpet/go-win64api/blob/master/sessions.go).

#### Groups
* Existence
* Comment

**Todos**
- [ ] Add check for membership

#### Software
* Existence
* DisplayVersion
* Arch
* Publisher
* InstallDate
* EstimatedSize
* Contact
* HelpLink
* InstallSource
* InstallLocation
* VersionMajor
* VersionMinor

**Todos**
- [ ] add list of windows features

#### Services
* Existence
* DisplayName
* StatusText
* Status
* AcceptStop
* IsRunning
* RunningPid
* ServiceType

#### Shares
* Existence
* Status
* Caption
* Description
* Path
* AllowMaximum
* Type

#### Processes
* Existence
* Pid
* Ppid
* Username
* Executable

#### Firewalls
* Profile
* Enabled
* DefaultInboundAction
* DefaultOutboundAction
* AllowInboundRules
* AllowLocalFirewallRules
* AllowLocalIPsecRules
* AllowUserApps
* AllowUserPorts
* AllowUnicastResponseToMulticast
* NotifyOnListen
* EnableStealthModeForIPsec
* LogMaxSizeKilobytes
* LogAllowed
* LogBlocked
* LogIgnored
* Caption
* Description
* ElementName
* InstanceID
* DisabledInterfaceAliases
* LogFileName
* Name
* PSComputerName

#### Policies
* String Policy Values
* String array Policy Values
* Integer64 Policy Values
* Binary Policy Values
* MUI Policy Values

#### Updates
* Updated
* Update History Exists

#### Bitlocker
* Existence
* DeviceID
* PersistentVolumeID
* ConversionStatus
* ProtectionStatus

#### Profiles
* User Profile Directory
* Global Profile Directory