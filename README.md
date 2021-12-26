# protoncheck  
lightweight, fast waybar/polybar/yabar/i3blocks module to check the amount of unread emails in a ProtonMail inbox.

## Installation
You may compile protoncheck yourself or download a copy from the [Releases](https://github.com/servusdei2018/protoncheck/releases/latest).

## Usage
Run `protoncheck`, using the `--username` and `--password` flags to supply your ProtonMail username and password, respectively. It will output the amount of unread emails in your inbox.

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

[![](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://paypal.me/njtbracy)
