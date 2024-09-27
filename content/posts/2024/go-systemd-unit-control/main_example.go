package main

import (
	"fmt"
	"os"
	"time"

	"context"

	"github.com/coreos/go-systemd/v22/dbus"
)

const (
	exitCodeSuccess                  = 0
	exitCodeFailedToListUnits        = 11
	exitCodeUnitNotFound             = 12
	exitCodeUnitNotActive            = 13
	exitCodeUnitNotRunning           = 14
	exitCodeErrorWhileWaiting        = 15
	exitCodeTimeoutWaitingForRestart = 16
)

func main() { os.Exit(mainWithExitCode()) }

func mainWithExitCode() int {
	// What service are we looking at?
	targetSystemdUnit := "fakemill.service"

	ctx := context.Background()
	// Connect to systemd
	// Specifically this will look DBUS_SYSTEM_BUS_ADDRESS environment variable
	// For example: `unix:path=/run/dbus/system_bus_socket`
	systemdConnection, err := dbus.NewSystemConnectionContext(ctx)
	if err != nil {
		fmt.Printf("Failed to connect to systemd: %v\n", err)
		panic(err)
	}
	defer systemdConnection.Close()

	listOfUnits, err := systemdConnection.ListUnitsContext(ctx)
	if err != nil {
		fmt.Printf("Failed to list units: %v\n", err)
		return exitCodeFailedToListUnits
	}

	found := false
	targetUnit := dbus.UnitStatus{}
	for _, unit := range listOfUnits {
		if unit.Name == targetSystemdUnit {
			fmt.Printf("Found systemd unit %s\n", targetSystemdUnit)
			found = true
			targetUnit = unit
			break
		}
	}
	if !found {
		fmt.Printf("Expected systemd unit %s not found\n", targetSystemdUnit)
		return exitCodeUnitNotFound
	}

	// Validate it's current state
	desiredActiveState := "active"
	desiredSubState := "running"
	fmt.Printf("Unit %s is in state %s and substate %s\n", targetUnit.Name, targetUnit.ActiveState, targetUnit.SubState)
	if targetUnit.ActiveState != desiredActiveState {
		fmt.Printf("Expected systemd unit %s to be active, but it is %s\n", targetSystemdUnit, targetUnit.ActiveState)
		return exitCodeUnitNotActive
	}

	if targetUnit.SubState != desiredSubState {
		fmt.Printf("Expected systemd unit %s to be running, but it is %s\n", targetSystemdUnit, targetUnit.SubState)
		return exitCodeUnitNotRunning
	}

	// Restart the service..
	restartMode := "replace"
	// What is the `mode` parameter? TLDR: I think we want replace
	// see: https://www.freedesktop.org/wiki/Software/systemd/dbus/
	// StartUnit() enqeues a start job, and possibly depending jobs. Takes the unit to activate, plus a mode string.
	// The mode needs to be one of replace, fail, isolate, ignore-dependencies, ignore-requirements.
	// If "replace" the call will start the unit and its dependencies, possibly replacing already queued jobs that conflict with this.
	// If "fail" the call will start the unit and its dependencies, but will fail if this would change an already queued job.
	// If "isolate" the call will start the unit in question and terminate all units that aren't dependencies of it.
	// If "ignore-dependencies" it will start a unit but ignore all its dependencies. If "ignore-requirements" it will start a unit
	//  but only ignore the requirement dependencies. It is not recommended to make use of the latter two options. Returns the newly created job object.
	completedRestartCh := make(chan string)
	jobID, err := systemdConnection.RestartUnitContext(
		ctx,
		targetSystemdUnit,
		restartMode,
		completedRestartCh,
	)

	if err != nil {
		fmt.Printf("Failed to restart unit: %v\n", err)
		panic(err)
	}
	fmt.Printf("Restart job id: %d\n", jobID)

	// Wait for the restart to complete
	select {
	case <-completedRestartCh:
		fmt.Printf("Restart job completed for unit: %s\n", targetSystemdUnit)
	case <-time.After(30 * time.Second):
		fmt.Printf("Timed out waiting for restart job to complete for unit: %s\n", targetSystemdUnit)
	}

	// Wait for the service to reach a running state
	channelBuffer := 10

	// Configure which changes we care about
	isRelevantChangeFunc := func(before *dbus.UnitStatus, after *dbus.UnitStatus) bool {
		if before.ActiveState != after.ActiveState {
			fmt.Printf("Active state changed from %s to %s\n", before.ActiveState, after.ActiveState)
			return true
		}
		if before.SubState != after.SubState {
			fmt.Printf("Sub state changed from %s to %s\n", before.SubState, after.SubState)
			return true
		}
		return false
	}

	// Ignore any services that we don't care about by filtering them out
	filterUnits := func(unit string) bool {
		return unit != targetSystemdUnit
	}

	// Subscribe to the changes
	changeCh, errorCh := systemdConnection.SubscribeUnitsCustom(time.Millisecond*10, channelBuffer, isRelevantChangeFunc, filterUnits)

	// Wait for the service to be active and running or give up
	for {
		select {
		case changedUnits := <-changeCh:
			unitStatus := changedUnits[targetSystemdUnit]
			fmt.Printf("Unit %s has changed\n", targetSystemdUnit)
			fmt.Printf("UnitStatus dump: %+v \n", unitStatus)
			if unitStatus.ActiveState == desiredActiveState && unitStatus.SubState == desiredSubState {
				fmt.Printf("Unit %s is now active and running\n", targetSystemdUnit)
				return exitCodeSuccess
			}
		case <-errorCh:
			fmt.Printf("Error while waiting for unit %s to change\n", targetSystemdUnit)
			return exitCodeErrorWhileWaiting
		case <-time.After(30 * time.Second):
			fmt.Printf("Timed out waiting for restart job to complete for unit: %s\n", targetSystemdUnit)
			return exitCodeTimeoutWaitingForRestart
		}
	}
}