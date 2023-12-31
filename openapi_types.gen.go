// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package main

import (
	"time"
)

// Defines values for VirtualMachineStatus.
const (
	Failed  VirtualMachineStatus = "failed"
	Working VirtualMachineStatus = "working"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// VirtualMachine defines model for VirtualMachine.
type VirtualMachine struct {
	CreatedAt time.Time            `json:"created_at"`
	Id        int                  `json:"id"`
	Name      *string              `json:"name,omitempty"`
	Status    VirtualMachineStatus `json:"status"`
}

// VirtualMachineStatus defines model for VirtualMachine.Status.
type VirtualMachineStatus string

// CreateVMJSONRequestBody defines body for CreateVM for application/json ContentType.
type CreateVMJSONRequestBody = VirtualMachine
