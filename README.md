# TrackerJacker
"Tracker jackers are genetically engineered wasps, conceived and created in the Capitol. They are genetically coded to attack anyone or anything that disturbs their nest."

## Questions
For questions please open up an issue in Github that has the question along with any related lines of code noted! This will aid in tracking and discussing questions.

## Major Todos
- [ ] ensure that capitalization errors will not cause issues in validation. i.e: shift everything to lowercase and just for good measure trim whitespace.

## Checks
### Cross
1. Files
2. Hosts
3. Users (Basic)

#### Files
* Existence
* Hash

**Todos**
- [ ] Add check for text based questions from users, or decide if the app should host a mini server to provide http traffic to the user.

#### Hosts
* Host Existence
* Ip Existence

#### Users
* Existence

**Todos**
- [ ] Add meta checks for feature parody with Windows

### Windows
1. Users (Comprehensive)
2. Groups (Completed/Basic)
3. Software (Comprehensive)
4. Services (Comprehensive) 
5. Shares (Comprehensive)
6. Processes (In Process)
7. Policies (Planned)
8. Firewalls (Planned)

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