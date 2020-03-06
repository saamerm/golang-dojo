package client

import (
	"fmt" 
)

type VaultClient struct {
    password string
}

func NewVaultClient(vaultAddress string, tokenString string) (*VaultClient, error){
	p := VaultClient{password: vaultAddress}
	p.password = email
	return &p, error

}

func (p VaultClient) Print(){
	fmt.Println(p.make + " " + p.model + " ")
	fmt.Println(p.year)
	fmt.Println(p.colour)
}