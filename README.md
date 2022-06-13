<!-- README template from https://github.com/dthung1602/terraform-provider-instant-proxy -->


[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/dthung1602/terraform-provider-instant-proxy">
    <img src="./logo.png" width="600">
  </a>

<h3 align="center">InstantProxy Terraform Provider</h3>

<p align="center">
   An <b>unofficial</b> Terraform provider for Instant Proxy API
</p>


## Installation

Hashicorp provider repo: https://registry.terraform.io/providers/dthung1602/instant-proxy/latest

Installation: see usage below

## Usage

```terraform

terraform {
  required_providers {
    instantproxy = {
      version = "0.1.2"
      source  = "dthung1602/instant-proxy"
    }
  }
}

provider "instantproxy" {
  /**
  For testing you can use the fake instant proxy server available at 
  https://github.com/dthung1602/instant-proxy-api-client
  Not providing endpoint or setting it to empty string will cause 
  the provider make real requests to InstantProxy server
  */
  endpoint = "http://localhost:3000"
  username = "username"
  password = "password"
}

/** Get list of all available proxies for this account*/
data "instantproxy_proxies" "proxies" {}

/** Sample output:
proxies = {
  "id" = "1655031199"
  "proxies" = tolist([
    {
      "address" = "67.123.80.92:8800"
      "ip" = "67.123.80.92"
      "port" = 8800
    },
    {
      "address" = "145.37.250.71:8800"
      "ip" = "145.37.250.71"
      "port" = 8800
    }
  ])
}
*/
output "proxies" {
  value = data.instantproxy_proxies.proxies
}

/** Get list of all authorized ips */
data "instantproxy_authorized_ips" "data_auth_ips" {}

/** Sample output:
data_auth_ips = {
  "id" = "1655031201"
  "value" = tolist([
    "1.1.1.1",
    "127.0.0.1",
  ])
}
*/
output "data_auth_ips" {
  value = data.instantproxy_authorized_ips.data_auth_ips
}

/** Set value for authorized ips */
resource "instantproxy_authorized_ips" "res_auth_ips" {
  value = ["127.0.0.1", "8.8.8.8"]
}

/** Sample output:
data_auth_ips = {
  "id" = "1655031201"
  "value" = tolist([
    "127.0.0.1",
    "8.8.8.8",
  ])
}
*/
output "res_auth_ips" {
  value = resource.instantproxy_authorized_ips.res_auth_ips
}

```

## Development

1. Clone repo: `git clone https://github.com/dthung1602/terraform-provider-instant-proxy.git`
2. Go to `Makefile` and change `OS_ARCH` to match your operating system and architecture
3. Run `go mod tidy` to install golang dependencies
4. Make sure you have terraform >= 0.14 installed on your system
5. When any changes are make to source code:
   - Run `make install` to build & make your provider available in your system
   - Clear terraform lock file `.terraform.lock.hcl`


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact

Duong Thanh Hung - [dthung1602@gmail.com](mailto:dthung1602@gmail.com)

Project Link: [https://github.com/dthung1602/terraform-provider-instant-proxy](https://github.com/dthung1602/terraform-provider-instant-proxy)


<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Best README template](https://github.com/othneildrew/Best-README-Template)
* [Img Shields](https://shields.io)



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/dthung1602/terraform-provider-instant-proxy.svg?style=flat-square
[contributors-url]: https://github.com/dthung1602/terraform-provider-instant-proxy/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/dthung1602/terraform-provider-instant-proxy.svg?style=flat-square
[forks-url]: https://github.com/dthung1602/terraform-provider-instant-proxy/network/members
[stars-shield]: https://img.shields.io/github/stars/dthung1602/terraform-provider-instant-proxy.svg?style=flat-square
[stars-url]: https://github.com/dthung1602/terraform-provider-instant-proxy/stargazers
[issues-shield]: https://img.shields.io/github/issues/dthung1602/terraform-provider-instant-proxy.svg?style=flat-square
[issues-url]: https://github.com/dthung1602/terraform-provider-instant-proxy/issues
[license-shield]: https://img.shields.io/github/license/dthung1602/terraform-provider-instant-proxy.svg?style=flat-square
[license-url]: https://github.com/dthung1602/terraform-provider-instant-proxy/blob/master/LICENSE
