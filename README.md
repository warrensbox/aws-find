[![Build Status](https://travis-ci.org/warrensbox/aws-find.svg?branch=master)](https://travis-ci.org/warrensbox/aws-find)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/aws-find)](https://goreportcard.com/report/github.com/warrensbox/aws-find)
[![CircleCI](https://circleci.com/gh/warrensbox/aws-find/tree/release.svg?style=shield&circle-token=518d496e953ed4d63075c0fd84b7bac7af68ac7f)](https://circleci.com/gh/warrensbox/aws-find/tree/release)


# AWS Find

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/smallerlogo.png" alt="drawing" width="95" height="125"/>

<!-- ![gopher](https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/logo.png =100x20) -->

The `aws-find` command line tool lets you lookup AWS instances'information by their tag name. 
Once installed, the `aws-find` lets you provide the tag name and value as an argument to the command [see example](#how-to-use). 

See installation guide here: [aws-find installation](https://warrensbox.github.io/aws-find/)

## Installation

`aws-find` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/). 


```ruby
brew install warrensbox/tap/aws-find
```

### Linux

Installation for other linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/aws-find/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/aws-find/releases) 

## How to use:
### Pass parameters optional
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/aws-findemo.gif" alt="drawing" style="width: 480px;" /> 

1. You can pass parameters to your `aws-find` command on your terminal. 
2. Pass -t or --tag for the tag name. And, -n and --name for the tag value. 
3. Hit **Enter** to see the list of instaces 

### Without parameters
 <img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/aws-findemo1.gif" alt="drawing" style="width: 480px;" /> 

1. Optionally, you don't have to pass any parameters to see the list of all the instances in region.</li>
2. Pass -r or --region to specify another region</li>
3. Hit **Enter** to see the list of instaces 

## Additional Info

See how to *upgrade*, *uninstall*, *troubleshoot* here:[More info](https://warrensbox.github.io/aws-find/additional)


## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/aws-find/issues)








