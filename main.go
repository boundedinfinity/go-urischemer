package urischemer

//go:generate enumer -path=./main.go

// https://en.wikipedia.org/wiki/List_of_URI_schemes

type UriScheme string

const (
	Aaa                        UriScheme = "aaa"                     // Diameter Protocol
	Aaas                       UriScheme = "aaas"                    // Secure Diameter Protocol
	About                      UriScheme = "about"                   // About (used in browsers)
	Acap                       UriScheme = "acap"                    // Application Configuration Access Protocol
	Acct                       UriScheme = "acct"                    // Identifying user account
	Acr                        UriScheme = "acr"                     //	Anonymous Customer Reference
	Adiumxtra                  UriScheme = "adiumxtra"               //	Direct installation of Adium Xtras (plugins)
	Aft                        UriScheme = "afp"                     //	Apple Filing Protocol
	Afs                        UriScheme = "afs"                     // Andrew File System
	Aim                        UriScheme = "aim"                     // AOL Instant Messenger
	Apt                        UriScheme = "apt"                     //  APT protocol
	Attachment                 UriScheme = "attachment"              // 	Attaching resources to MHTML pages
	Aw                         UriScheme = "aw"                      //	Link to an Active Worlds world
	Amss                       UriScheme = "amss"                    //	Identifier for an AMSS broadcast
	Barion                     UriScheme = "barion"                  // Barion e-money wallet
	Beshare                    UriScheme = "beshare"                 //	Open a search query on a BeShare server
	Bitcoin                    UriScheme = "bitcoin"                 // Bitcoin address
	Blob                       UriScheme = "blob"                    //	Binary data access in browsers
	Bolo                       UriScheme = "bolo"                    // Bolo game
	Callto                     UriScheme = "callto"                  // Skype call
	Cap                        UriScheme = "cap"                     //Calendar access protocol
	Chrome                     UriScheme = "chrome"                  //	Specifies user interfaces built using XUL in Mozilla-based browsers
	ChromeExtension            UriScheme = "chrome-extension"        //	Manage Chrome extensions
	ComEventbriteAttendee      UriScheme = "com-eventbrite-attendee" //	Provisional	IANA registration template
	Cid                        UriScheme = "cid"                     // Referencing individual parts of an SMTP/MIME message	Permanent
	Coaps                      UriScheme = "coaps"                   //	Constrained Application Protocol
	Content                    UriScheme = "content"                 //	Accessing an Android content provider
	Crid                       UriScheme = "crid"                    // TV-Anytime Content Reference Identifier
	Cvs                        UriScheme = "cvs"                     // Concurrent Versions System (CVS) Repository
	Dab                        UriScheme = "dab"                     //	Identifier for a DAB broadcast
	Data                       UriScheme = "data"                    //	Inclusion of small data items inline
	Dav                        UriScheme = "dav"                     //	HTTP Extensions for Distributed Authoring (WebDAV)
	Dict                       UriScheme = "dict"                    // Dictionary service protocol
	DlnaPlaySingle             UriScheme = "dlna-playsingle"
	DlnaPlayContainer          UriScheme = "dlna-playcontainer"
	Dns                        UriScheme = "dns"  // Domain Name System
	Dntp                       UriScheme = "dntp" // Direct Network Transfer Protocol
	Doi                        UriScheme = "doi"  // Digital object identifier, a digital identifier for any object of intellectual property.
	Drm                        UriScheme = "drm"  // Identifier for a DRM broadcastUriScheme = "
	Dtn                        UriScheme = "dtn"  //	DTNRG research and development
	Dvb                        UriScheme = "dvb"
	Ed2k                       UriScheme = "ed2k" // 	Resources available using the eDonkey2000 network
	Example                    UriScheme = "example"
	Facetime                   UriScheme = "facetime"     //	FaceTime is a video conferencing software
	Fax                        UriScheme = "fax"          // telefacsimile numbers
	Feed                       UriScheme = "feed"         // web feed subscription
	File                       UriScheme = "file"         //	Addressing files on local or network file systems
	Filesystem                 UriScheme = "filesystem"   //	Abandoned part of File API
	Finger                     UriScheme = "finger"       //	Querying user information using the Finger protocol
	Fish                       UriScheme = "fish"         //	Accessing another computer's files using the SSH protocol
	Fm                         UriScheme = "fm"           //	Identifier for an FM broadcast
	Ftp                        UriScheme = "ftp"          //FTP resources
	Gemini                     UriScheme = "gemini"       //Used with the Gemini protocol
	Geo                        UriScheme = "geo"          // 	A Uniform Resource Identifier for Geographic Locations
	Gg                         UriScheme = "gg"           // 	Starting chat with Gadu-Gadu user
	Git                        UriScheme = "git"          // link to a GIT repository
	GizmoProject               UriScheme = "gizmoproject" //	Gizmo5 calling link
	Go                         UriScheme = "go"           // Common Name Resolution Protocol
	Gopher                     UriScheme = "gopher"       // Gopher protocol
	Gtalk                      UriScheme = "gtalk"        // Start a chat with a Google Talk user
	H323                       UriScheme = "h323"         // Used with H.323 multimedia communications
	Http                       UriScheme = "http"         // HTTP resources
	Https                      UriScheme = "https"        // HTTP connections secured using SSL/TLS
	Iax                        UriScheme = "iax"          // Inter-Asterisk eXchange protocol version 2
	Icap                       UriScheme = "icap"         // Internet Content Adaptation Protocol
	Icon                       UriScheme = "icon"         // Provisional	IETF Draft
	Im                         UriScheme = "im"           // Instant messaging protocol
	Imap                       UriScheme = "imap"         // Accessing e-mail resources through IMAP
	Info                       UriScheme = "info"         // Information Assets with Identifiers in Public Namespaces
	IotDisco                   UriScheme = "iotdisco"     // Identify things on Internet of Things
	Ipn                        UriScheme = "ipn"
	Ipp                        UriScheme = "ipp"       // Internet Printing Protocol
	Ipps                       UriScheme = "ipps"      // Internet Printing Protocol over HTTPS
	Irc                        UriScheme = "irc"       // Internet Relay Chat server
	Irc6                       UriScheme = "irc6"      // IPv6 equivalent of irc
	Ircs                       UriScheme = "ircs"      // Secure equivalent of irc
	Iris                       UriScheme = "iris"      // Internet Registry Information Service
	IrisBeep                   UriScheme = "iris.beep" // Internet Registry Information Service
	IrisXpc                    UriScheme = "iris.xpc"  // Internet Registry Information Service
	IrisXpcs                   UriScheme = "iris.xpcs" // Internet Registry Information Service
	IrisLws                    UriScheme = "iris.lws"  // Internet Registry Information Service
	Itms                       UriScheme = "itms"      // iTunes Music Store
	Jabber                     UriScheme = "jabber"
	Jar                        UriScheme = "jar"     // Compressed archive
	Jms                        UriScheme = "jms"     // Java Message Service
	Keyparc                    UriScheme = "keyparc" // Keyparc encrypt/decrypt resource
	Lastfm                     UriScheme = "lastfm"
	Ldap                       UriScheme = "ldap"       // LDAP directory request
	Ldaps                      UriScheme = "ldaps"      // Secure equivalent of ldap
	Magnet                     UriScheme = "magnet"     // magnet links
	MailServer                 UriScheme = "mailserver" // Access to data available from mail servers
	Mailto                     UriScheme = "mailto"     // SMTP e-mail addresses and default content
	Maps                       UriScheme = "maps"       // map links
	Market                     UriScheme = "market"     // Opens Google Play
	Message                    UriScheme = "message"    // Direct link to specific email message
	Mid                        UriScheme = "mid"        // Referencing SMTP/MIME messages, or parts of messages
	Mms                        UriScheme = "mms"        // Windows streaming media
	Modem                      UriScheme = "modem"
	MsHelp                     UriScheme = "ms-help"                      // help page on Microsoft Windows Help and Support Center
	MsSettings                 UriScheme = "ms-settings"                  // application Settings in Windows
	MsSettingsAirplanemode     UriScheme = "ms-settings-airplanemode"     // application Settings in Windows
	MsSettingsBluetooth        UriScheme = "ms-settings-bluetooth"        // application Settings in Windows
	MsSettingsCamera           UriScheme = "ms-settings-camera"           // application Settings in Windows
	MsSettingsCellular         UriScheme = "ms-settings-cellular"         // application Settings in Windows
	MsSettingsCloudStorage     UriScheme = "ms-settings-cloudstorage"     // application Settings in Windows
	MsSettingsEmailAndAccounts UriScheme = "ms-settings-emailandaccounts" // application Settings in Windows
	MsSettingsLanguage         UriScheme = "ms-settings-language"         // application Settings in Windows
	MsSettingsLocation         UriScheme = "ms-settings-location"         // application Settings in Windows
	MsSettingsLock             UriScheme = "ms-settings-lock"             // application Settings in Windows
	MsSettingsNfcTransactions  UriScheme = "ms-settings-nfctransactions"  // application Settings in Windows
	MsSettingsNotifications    UriScheme = "ms-settings-notifications"    // application Settings in Windows
	MsSettingsPower            UriScheme = "ms-settings-power"            // application Settings in Windows
	MsSettingsPrivacy          UriScheme = "ms-settings-privacy"          // application Settings in Windows
	MsSettingsProximity        UriScheme = "ms-settings-proximity"        // application Settings in Windows
	MsSettingsScreenRotation   UriScheme = "ms-settings-screenrotation"   // application Settings in Windows
	MsSettingsWifi             UriScheme = "ms-settings-wifi"             // application Settings in Windows
	MsSettingsWorkplace        UriScheme = "ms-settings-workplace"        // application Settings in Windows
	MsnIm                      UriScheme = "msnim"                        // Windows Live Messenger
	Msrp                       UriScheme = "msrp"                         //
	Msrps                      UriScheme = "msrps"                        // Message Session Relay Protocol
	Mtqp                       UriScheme = "mtqp"                         //	Message Tracking Query Protocol
	Mumble                     UriScheme = "mumble"                       //
	Mupdate                    UriScheme = "mupdate"                      // Mailbox Update Protocol
	Mvn                        UriScheme = "mvn"                          // Access Apache Maven repository artifacts
	News                       UriScheme = "news"                         // (Usenet) newsgroups and postings
	Nfs                        UriScheme = "nfs"                          // Network File System resources
	Ni                         UriScheme = "ni"                           // Naming Things with Hashes
	Nih                        UriScheme = "nih"                          // Naming Things with Hashes
	Nntp                       UriScheme = "nntp"                         // Usenet NNTP
	Oid                        UriScheme = "oid"
	OpaqueLockToken            UriScheme = "opaquelocktoken"
	OpenPgp4Fpr                UriScheme = "openpgp4fpr" // OpenPGP version 4 public keys
	Pack                       UriScheme = "pack"
	Palm                       UriScheme = "palm"       // system services in HP webOS applications
	Paparazzi                  UriScheme = "paparazzi"  // the application "Paparazzi!" (Mac only)
	Payto                      UriScheme = "payto"      // Designate target for payments
	Pkcs11                     UriScheme = "pkcs11"     // PKCS #11
	Platform                   UriScheme = "platform"   // Eclipse platform resources
	Pop                        UriScheme = "pop"        // Accessing mailbox through POP3	Permanent
	Pres                       UriScheme = "pres"       // Used in Common Profile for Presence (CPP) to identify presence
	Prospero                   UriScheme = "prospero"   // Prospero Directory Service
	Proxy                      UriScheme = "proxy"      // Alter proxy settings in the FoxyProxy application
	Psyc                       UriScheme = "psyc"       // Used to identify or locate a person, group, place or a service and specify its ability to communicate
	Query                      UriScheme = "query"      // Opens a filesystem query
	Redis                      UriScheme = "redis"      // Redis database
	Rediss                     UriScheme = "rediss"     // Redis database secure
	Reload                     UriScheme = "reload"     // REsource LOcation And Discovery Protocol
	Res                        UriScheme = "res"        // Used by Internet Explorer to display error pages
	Resource                   UriScheme = "resource"   // Creating mapping for resource protocol aliases generted by the resource instruction. Used by Firefox
	Rmi                        UriScheme = "rmi"        // Look up a Java object in an RMI registry
	Rsync                      UriScheme = "rsync"      // rsync
	Rtmfp                      UriScheme = "rtmfp"      // Real Time Media Flow Protocol
	Rtmp                       UriScheme = "rtmp"       // Real Time Messaging Protocol
	Rtsp                       UriScheme = "rtsp"       // Real Time Streaming Protocol
	S3                         UriScheme = "s3"         // Amazon S3 bucket
	SecondLife                 UriScheme = "secondlife" // Open the Map floater in Second Life application to teleport the resident to the location.	Provisional	IANA registration template
	Service                    UriScheme = "service"
	Session                    UriScheme = "session" // Media Resource Control Protocol
	Sftp                       UriScheme = "sftp"    // SFTP file transfers (not be to confused with FTPS (FTP/SSL))
	Sgn                        UriScheme = "sgn"     // Social Graph Node Mapper
	Shc                        UriScheme = "shc"     // SMART Health Cards Framework
	Shttp                      UriScheme = "shttp"   // Secure HTTP
	Sieve                      UriScheme = "sieve"   // ManageSieve protocol
	Sip                        UriScheme = "sip"     // Session Initiation Protocol (SIP)
	Sips                       UriScheme = "sips"    // Secure equivalent of sip
	Skype                      UriScheme = "skype"   // Launching Skype call (see also callto:)
	Smb                        UriScheme = "smb"     // Accessing SMB/CIFS shares
	Sms                        UriScheme = "sms"     // SMS capable devices for composing and sending messages
	Snews                      UriScheme = "snews"   // NNTP over SSL/TLS
	Snmp                       UriScheme = "snmp"    // Simple Network Management Protocol
	SoapBeep                   UriScheme = "soap.beep"
	SoapBeeps                  UriScheme = "soap.beeps"
	Soldat                     UriScheme = "soldat"
	Spotify                    UriScheme = "spotify" // Load a track, album, artist, search, or playlist in Spotify
	Ssh                        UriScheme = "ssh"     // SSH connections
	Steam                      UriScheme = "steam"   // Interact with Steam: install apps, purchase games, run games, etc
	Stun                       UriScheme = "stun"    // Session Traversal Utilities for NAT (STUN)
	Stuns                      UriScheme = "stuns"   // Session Traversal Utilities for NAT (STUN)
	Svn                        UriScheme = "svn"     // Subversion (SVN) source control repository
	Tag                        UriScheme = "tag"     // tag URI scheme
	TeamSpeak                  UriScheme = "teamspeak"
	Tel                        UriScheme = "tel"         // telephone numbers
	Telnet                     UriScheme = "telnet"      // telnet
	Tftp                       UriScheme = "tftp"        // Trivial File Transfer Protocol
	Things                     UriScheme = "things"      // Interact with Things: create new to-dos or go to a specific list
	ThisMessage                UriScheme = "thismessage" // multipart/related relative reference resolution
	Tn3270                     UriScheme = "tn3270"      // Interactive 3270 emulation sessions
	Tip                        UriScheme = "tip"         // Transaction Internet Protocol
	Turn                       UriScheme = "turn"        // Traversal Using Relays around NAT (TURN)
	Turns                      UriScheme = "turns"       // Secure Traversal Using Relays around NAT (TURN)
	Tv                         UriScheme = "tv"          // TV Broadcasts
	Udp                        UriScheme = "udp"         // BitTorrent tracker protocol based on UDP
	Udp2                       UriScheme = "udp"         // MPEG Transport Stream over UDP
	Unreal                     UriScheme = "unreal"
	Urn                        UriScheme = "urn" // Uniform Resource Names
	Ut2004                     UriScheme = "ut2004"
	Vemmi                      UriScheme = "vemmi" // Versatile Multimedia Interface
	Ventrilo                   UriScheme = "ventrilo"
	VideoTex                   UriScheme = "videotex"
	ViewSource                 UriScheme = "view-source" // Shows a web page as code
	Vnc                        UriScheme = "vnc"         // Virtual Network Computing
	Wasi                       UriScheme = "wais"        // Wide area information server (WAIS)
	WebCal                     UriScheme = "webcal"      // Calendars in iCalendar format
	Ws                         UriScheme = "ws"          // WebSocket protocol
	Wss                        UriScheme = "wss"         // Secure WebSocket protocol
	Wtai                       UriScheme = "wtai"        // Wireless Telephony Application Interface
	Wyciwyg                    UriScheme = "wyciwyg"     // What You Cache Is What You Get
	Xcon                       UriScheme = "xcon"
	XconUserId                 UriScheme = "xcon-userid"
	Xfire                      UriScheme = "xfire"
	XmlRpcBeep                 UriScheme = "xmlrpc.beep"
	XmlRpcBeeps                UriScheme = "xmlrpc.beeps"
	Xmpp                       UriScheme = "xmpp"          // XMPP
	Xri                        UriScheme = "xri"           // eXtensible Resource Identifier (XRI)
	Ymsgr                      UriScheme = "ymsgr"         // a Yahoo! Contact
	Z3950                      UriScheme = "z39.50"        // Z39.50 information access
	Z3950r                     UriScheme = "z39.50r"       // Z39.50 retrieval
	Z3950ss                    UriScheme = "z39.50s"       // Z39.50 session
	Admin                      UriScheme = "admin"         // URL scheme in the GNOME desktop environment to access file(s)
	App                        UriScheme = "app"           // URL scheme can be used by packaged applications to obtain resources that are inside a container
	Javascript                 UriScheme = "javascript"    // Execute JavaScript code
	Jdbc                       UriScheme = "jdbc"          // Java Database Connectivity technology
	MsTeams                    UriScheme = "msteams"       // Microsoft to launch the Microsoft Teams desktop client
	MsAccess                   UriScheme = "ms-access"     // Used by Microsoft to launch Microsoft Office applications
	MsExcel                    UriScheme = "ms-excel"      // Used by Microsoft to launch Microsoft Office applications
	MsInfoPath                 UriScheme = "ms-infopath"   // Used by Microsoft to launch Microsoft Office applications
	MsPowerPoint               UriScheme = "ms-powerpoint" // Used by Microsoft to launch Microsoft Office applications
	MsProject                  UriScheme = "ms-project"    // Used by Microsoft to launch Microsoft Office applications
	MsPublisher                UriScheme = "ms-publisher"  // Used by Microsoft to launch Microsoft Office applications
	MsSpd                      UriScheme = "ms-spd"        // Used by Microsoft to launch Microsoft Office applications
	MsVisio                    UriScheme = "ms-visio"      // Used by Microsoft to launch Microsoft Office applications
	MsWord                     UriScheme = "ms-word"       // Used by Microsoft to launch Microsoft Office applications
	Odbc                       UriScheme = "odbc"          // Open Database Connectivity
	Rdar                       UriScheme = "rdar"          // URL scheme used by Apple's internal issue-tracking system
	TrueConf                   UriScheme = "trueconf"      // Used by TrueConf Server
	Slack                      UriScheme = "slack"         // Slack to launch the Slack client
	Stratum                    UriScheme = "stratum"       // the Stratum protocol
	ZoomMtg                    UriScheme = "zoommtg"       // Used by Zoom conferencing software to launch the Zoom client
	ZoomUs                     UriScheme = "zoomus"        // Used by Zoom conferencing software to launch the Zoom client
)
