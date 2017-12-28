**Thorin**
===================
Thorin is a simple python script to do a snapshots on [Digital Ocean](https://www.digitalocean.com).

Why? I'm study Python and I'm a Lazy SysAdmin.

### ToDo's

* Writing tests.
* Improve the code
* Retire api_key and droplet_id in the code.

### Libs

* sys
* requests
* json
* datetime

### How to use

* Clone the repo and access Thorin folder:
```bash
$ git clone https://github.com/perylemke/thorin.git
$ cd thorin
```

* Open a thorin.py and input a Digital Ocean API Key and Droplet ID
```bash
$ vim thorin.py
>> api_key = 'YOUR KEY HERE!'
>> droplet_id = 'YOUR DROPLET ID HERE!'
```

* Execute thorin.py
```bash
./thorin.py
```

* At the end your snapshot was created.

Feedbacks are welcome and if you have a improvement fork us and pull us! :)

### That's it folks!