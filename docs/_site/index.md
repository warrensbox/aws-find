# AWS Instance Finder 

The `aws-find` command line tool lets you lookup AWS instances' IP address other information by their tag name. 
Once installed, the `aws-find` requires the tag name and value as an argument to the command [see example](#how-to-use). 

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

<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/awsfind.gif" alt="drawing" style="width: 480px;"/>

1.  You can switch between different versions of terraform by typing the command `awsfind` on your terminal. 
2.  Select the version of terraform you require by using the up and down arrow.
3.  Hit **Enter** to select the desired version

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/aws-find/issues){:target="_blank"}

<hr>

See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)