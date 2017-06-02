# Golang API Client for GRR Rapid Response

## What?

This package contains Go bindings for the GRR Rapid Response API. It
provides a way to talk to the GRR AdminUI server and is intended for
automation and integration tasks.

It currently targets the stable 3.1.0.2 release of GRR. This is the
version for which the GRR project has provided an Ubuntu/Xenial
package. There is also a
[package](https://packages.debian.org/sid/grr-server) of the 3.1.0.2
version in Debian/unstable which can be used on Debian 9.x ("stretch")
systems.

## Why?

During the
[March 8 2017 Meetup](https://www.youtube.com/watch?v=SIvf7-Lzp2M), it
was announced that someone inside Google had been working on a Golang
API client and that it would be published "soon". While a Python
client has been added to the
[api_client/python](https://github.com/google/grr/tree/master/api_client/python)
directory, no "official" Go client has been published as of late
May 2017. Until the GRR team at Google decides to publish their code,
I hope that this will be useful.

## Hacking / Work in progress

Data structures have been generated from the Protobuf definitions that
are shipped with GRR 3.1.0.2. Some functinos that are not implemented
in terms of the API have been reimplemented in `legacy.go`.

This package does not yet cover all functions and API endpoints. There
are templates for three basic patterns, see `generate.go`,
`generate/gen-functions.go` for details:

- simple GET: receive Protobuf Message
- simple POST: provide Protobuf Message, receive simple answer
- GET/POST: provide Protobuf Message, receive Protobuf Message

Remaining issues

- A GET -> binary stream function template (for the routes annotated
  with `@BinaryStream`)
- Some functions are not annotated with a result type at all
- How can `AnyType` be mapped to something useful, do we need the type
  information that is omitted by passing `strip_type_info=1`?
- Adding most of the functions / API routes that have been added since
  GRR 3.1.0.2 should be possible.

## License

Apache 2.0, see LICENSE file in the source distribution.

## Author

Hilko Bengen <bengen@hilluzination.de>
