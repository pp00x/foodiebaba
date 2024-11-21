package logger

import (
    "os"

    "github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// Init initializes the global logger
func Init() {
    Log = logrus.New()

    // Set the output destination
    Log.Out = os.Stdout

    // Set the log level (you can change this as needed)
    Log.SetLevel(logrus.InfoLevel)

    // Set the formatter
    Log.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
        PrettyPrint:     false,
    })
}