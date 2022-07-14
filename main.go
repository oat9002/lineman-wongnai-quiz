package main

import "encoding/base64"

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	for i := len(sd) - 1; i >= 0; i-- {
		whatIsIt += string(sd[i])
	}

	println(whatIsIt) // Join:us:at:LINE:MAN:Wongnai
}
