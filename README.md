# 🔬 Go YAML research

Research YAML parsers written on Go.

## Motivation

In developing [✨Sparkle](https://sparkle.wiki/), I've encountered
a need to select an appropriate library for handling YAML data
because it is a format of front matter.

## Requirements

1. **Deterministic serialization/deserialization:** The library
must ensure deterministic outcomes in serialization and
deserialization processes. Given the same input, the library
should consistently produce the same output every time,
ensuring predictability and reliability in processing data.

2. **Support for complex data types:** The library must be able
to handle complex data types. This includes basic types and
more intricate structures such as nested objects, arrays, and
custom types. Such support is essential for dealing with
diverse data formats and systems typically encountered in YAML files.

3. **Concise and user-friendly API:** The library should offer
a straightforward and intuitive API. This will facilitate ease
of use and enable developers to perform tasks with minimal code,
enhancing productivity and reducing the learning curve.
A well-designed API makes it easier to implement YAML processing
functionalities without delving into unnecessary complexities.

4. **No runtime panics:** Stability is a crucial requirement,
and the library should be designed to avoid runtime panics.
It should handle errors gracefully and provide clear error messages,
allowing for robust error handling by the developer. Ensuring
the library operates smoothly under various scenarios and inputs
is vital to maintaining the overall reliability of the application.

## Resources

- https://github.com/go-yaml/yaml
  - github.com/goccy/go-yaml
- https://github.com/goccy/go-yaml
  - gopkg.in/yaml.v3

## The result

Interim results or a detailed follow-up.[^1]

<p align="right">made with ❤️ for everyone by <a href="https://www.octolab.org/">OctoLab</a></p>

[^1]: work in progress
