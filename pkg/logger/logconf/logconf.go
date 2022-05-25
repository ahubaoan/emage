package logconf

type OutFileRotate struct {
	Enable     bool   `json:"enable" yaml:"enable"`
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"max_size" yaml:"max_size"`       // maximum size in megabytes of the log file
	MaxBackups int    `json:"max_backups" yaml:"max_backups"` // maximum number of old log files to retain
	MaxAge     int    `json:"max_age" yaml:"max_age"`         // maximum number of days to retain old log files
	Compress   bool   `json:"compress" yaml:"compress"`       // the rotated log files should be compressed using gzip
}

type OutputElasticSearch struct {
	Enable bool
}

type EncodeTimeType uint

const (
	EncodeTimeTypeEpoch EncodeTimeType = iota
	EncodeTimeTypeISO8601
)

const (
	FormatTypeJson    = "json"
	FormatTypeConsole = "console"
)

type LogConf struct {
	UseLog        bool           `json:"use_log" yaml:"use_log"`
	TimeType      EncodeTimeType `json:"time_type" yaml:"time_type"` // default
	FormatType    string         `json:"format_type" yaml:"format_type"`
	DebugLevel    uint8          `json:"debug_level" yaml:"debug_level"`
	OutStdout     bool           `json:"out_stdout" yaml:"out_stdout"`
	OutFileRotate OutFileRotate  `json:"out_file_rotate" yaml:"out_file_rotate"`
}
