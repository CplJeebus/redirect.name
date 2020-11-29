# Why? 

So I had used [redirect.name](http://redirect.name/) previously for some projects, but was concerned about having too much details in  TXT records. The big driver was Zoom embedding a `pwd` in your personal Zoom room. 

Also I just wanted to play with `go` and google `run`. 

## What this does. 
Checks for a TXT recored to be redirected, once it gets the record substitutes out tokens with the values stored in a `secrets.yaml` file. 

With this in mind you will need to self host this instance or redirect.name to manage the secrets.yaml. 

# This is a Fork of the excellent project below. 

## [redirect.name](http://redirect.name/) [![Build status][ci-image]][ci-url]

Please refer to [redirect.name](http://redirect.name/) for documentation.

[ci-image]: https://img.shields.io/circleci/project/holic/redirect.name/master.svg?style=flat
[ci-url]: https://circleci.com/gh/holic/redirect.name
