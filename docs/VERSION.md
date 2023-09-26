# Version

```yaml
version: 0.1.0
```

To avoid issues related to using incompatible versions, define the `version` attribute in the schema.
Gontainer automatically checks whether the current build is compatible with the given configuration.
Whenever the version is not given, this check will be skipped.
Whenever the current build does not have defined the proper version
(e.g. Gontainer has been built manually without proper `ldflags`), this check will be skipped.
