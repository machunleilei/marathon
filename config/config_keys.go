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
	//ReadWriteTimeout ...
	ReadWriteTimeout = "ReadWriteTimeout"
	//RequestTimeout ...
	RequestTimeout = "RequestTimeout"
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
	//FailureCountSlidingWindowInterval ...
	FailureCountSlidingWindowInterval = "FailureCountSlidingWindowInterval"
	//PingInterval ...
	PingInterval = "PingInterval"
	//PingStrategy ...
	PingStrategy = "PingStrategy"
	//LoadBalancerRule ...
	LoadBalancerRule = "LoadBalancerRule"
	//LoadBalancerKey
	LoadBalancerKey = "LoadBalancerKey"
	//ListOfServersPollingInterval ...
	ListOfServersPollingInterval = "ListOfServersPollingInterval"
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
	DefaultConnectTimeout = 200 * time.Millisecond
	//DefaultReadTimeout ...
	DefaultReadWriteTimeout = 500 * time.Millisecond
	//DefaultRequestTimeout ...
	DefaultRequestTimeout = 500 * time.Millisecond
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
	DefaultConnectionFailureThreshold = 5
	//DefaultCircuitTrippedTimeoutFactor ...
	DefaultCircuitTrippedTimeoutFactor = 10
	//DefaultCircuitTripMaxTimeout ...
	DefaultCircuitTripMaxTimeout = 30 * time.Second
	//DefaultFailureCountSlidingWindowInterval ...
	DefaultFailureCountSlidingWindowInterval = 10 * time.Second
	//DefaultPingInterval ...
	DefaultPingInterval = 5 * time.Second
	//DefaultPingStrategy ...
	DefaultPingStrategy = "SerialPingStrategy"
	//DefaultLoadBalancerRule ...
	DefaultLoadBalancerRule = "RandomRule"
	//DefaultLoadBalancerKey ...
	DefaultLoadBalancerKey = "marathon"
	//DefaultListOfServersPollingInterval ...
	DefaultListOfServersPollingInterval = 30 * time.Second
)

//PingStrategy ...
const (
	//SerialPingStrategy ...
	SerialPingStrategy = "SerialPingStrategy"
	//ParallelPingStrategy ...
	ParallelPingStrategy = "ParallelPingStrategy"
)

//LoadBalancer Rule
const (
	//RandomLBRule ...
	RandomLBRule = "RandomRule"
	//RoundRobinLBRule ...
	RoundRobinLBRule = "RoundRobinRule"
)
