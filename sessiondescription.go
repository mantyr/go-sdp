package sdp

type SessionDescription struct {
    Version           int
    Origin            Origin
    SessionName       string
    Info              string
    Uri               string
    Emails            []Email
    Phones            []Phone
    Connection        Connection
    Bandwidths        []Bandwidth
    Times             []TimeField
    Key               Key
    Attributes        map[string]string
    MediaDescriptions []MediaDescription
}

type Origin struct {
    Username       string
    SessionId      string
    SessionVersion string
    NetType        string
    AddrType       string
    UnicastAddr    string
}

type Email struct {
    Address string
    Name    string
}

type Phone struct {
    Address string
    Name    string
}

type Connection struct {
    NetType  string
    AddrType string
    Address  string
}

type Bandwidth struct {
    Type string
    Bandwidth string
}

type SessionTime struct {
    Start time.*Time
    Stop time.*Time
    Repeats []Repeat
    Zones []Zone
}

type Repeat struct {
    Interval time.*Duration
    Active time.*Duration
    Offsets []time.*Time
}

type Zone struct {
    Time Time
    TypedTime
}

type MediaDescription struct {
    MediaField
    Information
    []Connection
    []Bandwidth
    Key
    Attributes map[string]string
}

type MediaField struct {
    Type string
    Port int
    Proto string
}