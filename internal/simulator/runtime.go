package simulator

import (
    "fmt"
    "os/exec"
)

// RuntimeManager abstracts the underlying simulation runtime.
type RuntimeManager struct{}

func NewRuntimeManager() *RuntimeManager { return &RuntimeManager{} }

func (r *RuntimeManager) Ensure(kind, name, desired string) error {
    if _, err := exec.LookPath("docker"); err != nil {
        fmt.Printf("docker not available; simulated reconcile for %s/%s: %s\n", kind, name, desired)
        return nil
    }
    return exec.Command("docker", "info").Run()
}
