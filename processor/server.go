package processor

import (
	"bytes"
	"os/exec"
	"time"
	"syscall"
	"tcf/logger"
)

func RunProcessor(args string) (stdout string){
	stdout,stderr,killed,err:=runCommandWithTimeout(30, "sudo docker exec tcf bash -c",
		"cd ../../examples/apps/generic_client/; python generic_client.py", "--uri", "http://localhost:1947",
		"heart-disease-eval","--in_data ","Data:",args)
	if err == nil{
		logger.Log.Printf("killed:%v,stdout:%v",killed,stdout)
	}else{
		logger.Log.Printf("killed:%v,stderr:%v",killed,stderr)
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
