Go bindings for bitcoin secp256k1
======
[![Documentation Status](https://readthedocs.org/projects/ansicolortags/badge/?version=latest)](http://ansicolortags.readthedocs.io/?badge=latest) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)


### Important: Be sure that you use golang < 1.9.4 due to the fact of this [Issue](https://github.com/golang/go/issues/23739)

### How to install
====
```
cd $GOPATH/src/github.com/0x4139/secp256k1
git submodule update
cd c-secp256k1
make distclean && ./autogen.sh && ./configure && make
cd ..
go clean && go install
```


### Docs
====
Due to the fact that it implements a part of the functions of the bitcoin secp256k1 you can look directly in to the 
[Header File](https://github.com/bitcoin-core/secp256k1/blob/3087bc4d75ec17287e71a36bda5df52a9ab8d854/include/secp256k1.h)

### License
=====

````
  DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2018 <s0x4139@gmail.com>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
````