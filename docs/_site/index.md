# AWS Instance Finder 

The `aws-find` command line tool lets you lookup AWS instances' IP address and other information by their tag name. 
Once installed, the `aws-find` lets you provide the tag name and value as an argument to the command [see example](#how-to-use). 

<hr>

## Installation

`aws-find` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/){:target="_blank"}. 


```ruby
brew install warrensbox/tap/awsfind
```

### Linux

Installation for Linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/aws-find/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from the source [here](https://github.com/warrensbox/aws-find/releases) 

<hr>

### How to use

<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/awsfindemo.gif" alt="drawing" style="width: 480px;" /> 

1. You can pass parameters to your `awsfind` command on your terminal
2. Pass -t or --tag for the tag name. And, -n and --name for the tag value
3. Hit **Enter** to see the list of instaces 

 <img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/awsfindemo1.gif" alt="drawing" style="width: 480px;" /> 

1. Optionally, you don't have to pass any parameters to see the list of all the instances in region
2. Pass -r or --region to specify another region 
3. Hit **Enter** to see the list of instaces 

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/aws-find/issues){:target="_blank"}

<hr>

See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)