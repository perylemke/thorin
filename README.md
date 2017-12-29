**Thorin**
===================
Thorin is a simple Go CLI to do a snapshots on [Digital Ocean](https://www.digitalocean.com).

Why? I'm study Go and I'm a Lazy SysAdmin.

### ToDo's

* Writing tests.
* Improve the code
* Implement Travis CI

### Libs

* bytes
* encoding/json
* fmt
* net/http
* os
* time

### How to use

* Clone the repo and access Thorin folder:
```bash
$ go get github.com/perylemke/thorin.git
```

* Open your ~/.bashrc or ~/.zshrc and input a Digital Ocean API Key and Droplet ID
```bash
$ vim ~/.bashrc # Or vim ~/.zshrc
$ export DO_API_KEY=YOUR_KEY_HERE
$ export DO_DROPLET_ID=YOUR_DROPLET_ID_HERE
```

* Execute thorin
```bash
thorin
```
* At the end your snapshot was created.

Feedbacks are welcome and if you have a improvement fork us and pull us! :)

### That's it folks!