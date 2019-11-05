package processor

import (
	"bytes"
	"os/exec"
	"time"
	"syscall"
	"tcf/logger"
)

func RunProcessor(args string) (stdout, stderr string){
	command_template := "cd ../../examples/apps/generic_client/; python generic_client.py --uri 'http://localhost:1947' --workload_id 'heart-disease-eval' --in_data 'Data: "+args+"' -m 1"
	stdout,stderr,killed,_:=runCommandWithTimeout(60*10, "docker", "exec","tcf", "/bin/bash","-c",command_template)	//stdout,stderr,killed,_:=runCommandWithTimeout(30, "docker exec", "tcf bash -c",
	logger.Log.Printf("killed:%v,stdout:%v",killed,stdout)
	logger.Log.Printf("killed:%v,stderr:%v",killed,stderr)
	if stdout == "" {
		stdout = stderr
	}
	return
}


// qiye add, wait for your python script until timeout..
func runCommandWithTimeout(timeout int, command string, args ...string) (stdout, stderr string, isKilled bool, err error) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	cmd.Start()
	done := make(chan error)

	err = nil
	go func() {
		done <- cmd.Wait()
	}()
	after := time.After(time.Duration(timeout) * time.Second)
	select {
	case <-after:
		cmd.Process.Signal(syscall.SIGINT)
		time.Sleep(10 * time.Millisecond)
		cmd.Process.Kill()
		isKilled = true
	case err = <-done:
		isKilled = false

	}
	stdout = string(bytes.TrimSpace(stdoutBuf.Bytes())) // Remove \n
	stderr = string(bytes.TrimSpace(stderrBuf.Bytes())) // Remove \n
	return
}
