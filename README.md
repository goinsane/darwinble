# darwinble

**darwinble** implements Go BLE (Bluetooth Low Energy) bindings for Darwin (MacOS and IOS)
with [CoreBluetooth](https://developer.apple.com/documentation/corebluetooth?language=objc).

## Documentation

For documentation, see the [CoreBluetooth docs](https://developer.apple.com/documentation/corebluetooth?language=objc).

Examples are in the `examples` directory.

## Scope

**darwinble** aims to implement all functionality that is supported in macOS 10.13.

## Naming

Function and type names in **darwinble** are intended to match the corresponding CoreBluetooth functionality as closely
as possible. There are a few (consistent) deviations:

- All **darwinble** identifiers start with a capital letter to make them public.
- Named arguments in CoreBluetooth functions are eliminated.
- Properties are implemented as a pair of functions (`PropertyName` and `SetPropertyName`).

## Issues

There are definitely memory leaks. ARC is not compatible with cgo, so objective C memory has to be managed manually.
Contributors didn't see a set of consistent guidelines for object ownership in the CoreBluetooth documentation, so *
*darwinble** errs on the side of leaking. Hopefully this is only an issue for very long running processes!  Any fixes
here are much appreciated.

## Source Code Chain

- The repository is forked from [github.com/tinygo-org/cbgo](https://github.com/tinygo-org/cbgo)
- [github.com/tinygo-org/cbgo](https://github.com/tinygo-org/cbgo) is forked
  from [github.com/JuulLabs-OSS/cbgo](https://github.com/JuulLabs-OSS/cbgo)
