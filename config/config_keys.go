package config

import (
    "time"
)

//config keys
const (
    //EnableConnectionPool ...
    EnableConnectionPool = "EnableConncetionPool"
    //MaxConnectionsPerHost ...
    MaxConnectionsPerHost = "MaxConnectionsPerHost"
    //MaxTotalConnections ...
    MaxTotalConnections = "MaxTotalConnections"
    //ConnectTimeout ...
    ConnectTimeout = "ConnectTimeout"
    //ReadTimeout ...
    ReadTimeout = "ReadTimeout"
    //MaxAutoRetries ...
    MaxAutoRetries = "MaxAutoRetries"
    //MaxAutoRetriesNextServer ...
    MaxAutoRetriesNextServer = "MaxAutoRetriesNextServer"
    //OKToRetryOnAllOperations ...
    OKToRetryOnAllOperations = "OKToRetryOnAllOperations"
    //Port ...
    Port = "Port"
    //ListOfServers ...
    ListOfServers = "ListOfServers"
    //ConnectionFailureThreshold ...
    ConnectionFailureThreshold = "ConnectionFailureThreshold"
    //CircuitTrippedTimeoutFactor ...
    CircuitTrippedTimeoutFactor = "CircuitTrippedTimeoutFactor"
    //CircuitTripMaxTimeout ...
    CircuitTripMaxTimeout = "CircuitTripMaxTimeout"
    //PingInterval ...
    PingInterval = "PingInterval"
)

//config default value
const (
    //DefaultEnableConnectionPool ...
    DefaultEnableConnectionPool bool = true
    //DefaultMaxConnectionsPerHost ...
    DefaultMaxConnectionsPerHost int = 50
    //DefaultMaxTotalConnections ...
    DefaultMaxTotalConnections int = 200
    //DefaultConnectTimeout ...
    DefaultConnectTimeout = 2000 * time.Millisecond //ms
    //DefaultReadTimeout ...
    DefaultReadTimeout = 5000 * time.Millisecond //ms
    //DefaultMaxAutoRetries
    DefaultMaxAutoRetries int = 0
    //DefaultMaxAutoRetriesNextServer ...
    DefaultMaxAutoRetriesNextServer int = 1
    //DefaultOKToRetryOnAllOperations ...
    DefaultOKToRetryOnAllOperations bool = false
    //DefaultPort ...
    DefaultPort int = 80
    //DefaultListOfServers ...
    DefaultListOfServers string = ""
    //DefaultConnectionFailureThreshold ...
    DefaultConnectionFailureThreshold = 3
    //DefaultCircuitTrippedTimeoutFactor ...
    DefaultCircuitTrippedTimeoutFactor = 10
    //DefaultCircuitTripMaxTimeout ...
    DefaultCircuitTripMaxTimeout = 30 * time.Second
    //DefaultPingInterval ...
    DefaultPingInterval = 5 * time.Second

)