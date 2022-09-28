module github.com/eyedeekay/about.i2p

replace cerca => ../../cblgh/cerca 

replace cerca/html => ./about/html

go 1.19

require (
	cerca v0.0.0-00010101000000-000000000000
	github.com/eyedeekay/onramp v0.0.0-20220829050101-64cb1842d0f0
)

require (
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/carlmjohnson/requests v0.22.1 // indirect
	github.com/cretz/bine v0.2.0 // indirect
	github.com/eyedeekay/i2pkeys v0.0.0-20220804220722-1048b5ce6ba7 // indirect
	github.com/eyedeekay/sam3 v0.33.3 // indirect
	github.com/gomarkdown/markdown v0.0.0-20211212230626-5af6ad2f47df // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	github.com/microcosm-cc/bluemonday v1.0.17 // indirect
	github.com/synacor/argon2id v0.0.0-20190318165710-18569dfc600b // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220805013720-a33c5aa5df48 // indirect
	golang.org/x/sys v0.0.0-20220804214406-8e32c043e418 // indirect
)
