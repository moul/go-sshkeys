# go-sshkeys
:wrench: Golang SSH Keys manipulation library

[![GoDoc](https://godoc.org/github.com/moul/go-sshkeys?status.svg)](https://godoc.org/github.com/moul/go-sshkeys) [![GuardRails badge](https://badges.production.guardrails.io/moul/go-sshkeys.svg)](https://www.guardrails.io)

## Install

```console
$ go get github.com/moul/go-sshkeys
```

## Example

```go
import "github.com/moul/go-sshkeys"

sshkey, _ := sshkeys.NewSSHKey([]byte("ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEApvPvDbWDY50Lsx4WyUInw407379iERte63OTTNae6+JgAeYsn52Z43Oeks/2qC0gxweq+sRY9ccqhfReie+r+mvl756T4G8lxX1ND8m6lZ9kM30Rvk0piZn3scF45spmLNzCNXza/Hagxy53P82ej2vq2ewXtjVdvW20G3cMHVLkcdgKJN+2s+UkSYlASW6enUj3no+bukT+6M8lJtlT0/0mZtnBRJtqCCvF0cm9xU0uxILrhIfdYAJ1XqaoqIQLFSDLVo5lILMzDNwV+CfAotRMWIKvWomCszhVQYHCQo2Z+b2Gs0TL4DRb23fRMdeaRufnVhh5ZMlNkb2ajaL6sw== m@42.am"))

fmt.Println(sshkey.Hash())
// Output: 7d:c7:69:37:2b:5f:84:36:db:75:98:44:db:b8:1a:36
```

## License

MIT
