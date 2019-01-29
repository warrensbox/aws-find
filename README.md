[![Build Status](https://travis-ci.org/warrensbox/aws-find.svg?branch=master)](https://travis-ci.org/warrensbox/aws-find)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/aws-find)](https://goreportcard.com/report/github.com/warrensbox/aws-find)
[![CircleCI](https://circleci.com/gh/warrensbox/aws-find/tree/master.svg?style=shield&circle-token=55ddceec95ff67eb38269152282f8a7d761c79a5)](https://circleci.com/gh/warrensbox/aws-find)

# AWS Find

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/awsfind/smallerlogo.png" alt="drawing" width="120" height="130"/>

<!-- ![gopher](https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/awsfind/logo.png =100x20) -->

The `awsfind` command line tool lets you switch between different versions of [terraform](https://www.terraform.io/). 
If you do not have a particular version of terraform installed, `awsfind` will download the version you desire.
The installation is minimal and easy. 
Once installed, simply select the version you require from the dropdown and start using terraform. 

See installation guide here: [awsfind installation](https://warrensbox.github.io/aws-find/)

## Installation

`awsfind` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/). 


```ruby
brew install warrensbox/tap/awsfind
```

### Linux

Installation for other linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/aws-find/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/aws-find/releases) 

## How to use:
### Use dropdown menu to select version
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/awsfind/awsfind.gif" alt="drawing" style="width: 180px;"/>

1.  You can switch between different versions of terraform by typing the command `awsfind` on your terminal. 
2.  Select the version of terraform you require by using the up and down arrow.
3.  Hit **Enter** to select the desired version.

The most recently selected versions are presented at the top of the dropdown.

### Supply version on command line
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/awsfind/awsfind-v4.gif" alt="drawing" style="width: 170px;"/>

1. You can also supply the desired version as an argument on the command line.
2. For example, `awsfind 0.10.5` for version 0.10.5 of terraform.
3. Hit **Enter** to switch.

## Additional Info

See how to *upgrade*, *uninstall*, *troubleshoot* here:[More info](https://warrensbox.github.io/aws-find/additional)


## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/aws-find/issues)








