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
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/awsfindemo.gif" alt="drawing" style="width: 480px;" /> 

1. You can pass parameters to your `awsfind` command on your terminal. 
2. Pass -t or --tag for the tag name. And, -n and --name for the tag value. 
3. Hit **Enter** to see the list of instaces 

 <img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/awsfindemo1.gif" alt="drawing" style="width: 480px;" /> 

1. Optionally, you don't have to pass any parameters to see the list of all the instances in region.</li>
2. Pass -r or --region to specify another region</li>
3. Hit **Enter** to see the list of instaces 

## Additional Info

See how to *upgrade*, *uninstall*, *troubleshoot* here:[More info](https://warrensbox.github.io/aws-find/additional)


## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/aws-find/issues)








