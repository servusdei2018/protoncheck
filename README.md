# protoncheck 
[![](https://goreportcard.com/badge/github.com/servusdei2018/protoncheck)](https://goreportcard.com/report/github.com/servusdei2018/protoncheck) [![Go Reference](https://pkg.go.dev/badge/github.com/servusdei2018/protoncheck.svg)](https://pkg.go.dev/github.com/servusdei2018/protoncheck) [![](https://sonarcloud.io/api/project_badges/measure?project=servusdei2018_protoncheck&metric=alert_status)](https://sonarcloud.io/project/overview?id=servusdei2018_protoncheck)  [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go)

awesome waybar/polybar/yabar/i3blocks module to check the amount of unread emails in a ProtonMail inbox.

## Installation
You may compile protoncheck yourself or download a copy from the [Releases](https://github.com/servusdei2018/protoncheck/releases/latest).

## Usage
Run `protoncheck`, using the `--username` and `--password` flags to supply your ProtonMail username and password, respectively. It will output the amount of unread emails in your inbox.

> ⚠️  _Protoncheck doesn't work with accounts with 2FA enabled._

#### waybar Config
Below is a sample Waybar configuration to check your ProtonMail inbox every 30 seconds.

~/.config/waybar/config:
```json
"modules-right": {
  "custom/protoncheck"
}

"custom/protoncheck": {
		"exec": "~/bin/protoncheck --username USER@protonmail.com --password PASSWORD",
		"format": " {}",
		"interval": 30,
		"on-click": "xdg-open https://mail.protonmail.com/inbox"
}
```

#### polybar Config
Below is a sample Polybar configuration to check your ProtonMail inbox every 30 seconds.

~/.config/polybar/config:
```yaml
modules-right = protoncheck

[module/protoncheck]
type = custom/script
exec = ~/bin/protoncheck --username USER@protonmail.com --password PASSWORD | sed  's/^/ /'
interval = 30
click-left = xdg-open https://mail.protonmail.com/inbox
```

#### yabar Config
Below is a sample yabar configuration to check your ProtonMail inbox every 30 seconds.

~/.config/yabar/yabar.conf:
```
	title: {
		exec: ~/bin/protoncheck --username USER@protonmail.com --password PASSWORD | sed  's/^/ /';
		type: "periodic";
		interval: 30;
	}
```

#### i3blocks Config
Below is a sample i3blocks configuration to check your ProtonMail inbox every 30 seconds.

~/.config/i3blocks/config:
```
[protoncheck]
command=~/bin/protoncheck --username USER@protonmail.com --password PASSWORD | sed  's/^/ /'
interval=30
color=#A4C2F4
```

## Credits
 - Special thank you to [@crabvk](https://github.com/crabvk) for making [bar-protonmail](https://github.com/crabvk/bar-protonmail), without which I wouldn't have had the inspiration to make protoncheck.

## Sponsoring
Work on `protoncheck` is done by volunteers in their free time, and takes time and effort. To show your appreciation and to provide continued motivation for work on this and other open source software, please send a donation.

[![](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://www.paypal.com/donate?business=S9REMLHZB64NQ&no_recurring=0&item_name=Help+fund+my+college+tuition+and+open-source+work&currency_code=USD)
