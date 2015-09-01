package sshkeys

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type SSHKey struct {
	Key []byte

	Pubkey  ssh.PublicKey
	Comment string
	Options []string
	Rest    []byte
	Type    string
}

func NewSSHKey(key []byte) (*SSHKey, error) {
	sshkey := SSHKey{
		Key: key,
	}
	err := sshkey.parse()
	if err != nil {
		return nil, err
	}
	return &sshkey, nil
}

func (s *SSHKey) parse() error {
	var err error
	s.Pubkey, s.Comment, s.Options, s.Rest, err = ssh.ParseAuthorizedKey(s.Key)
	if err != nil {
		return err
	}
	s.Type = s.Pubkey.Type()

	// FIXME: compute bits
	// FIXME: check min/max key length
	return nil
}

func (s *SSHKey) Hash() string {
	output := ""

	hash := fmt.Sprintf("%x", md5.Sum(s.Pubkey.Marshal()))
	hashLength := len(hash)
	for i, r := range []rune(hash) {
		output += string(r)
		if (i+1)%2 == 0 && i < hashLength-1 {
			output += ":"
		}
	}
	return output
}
