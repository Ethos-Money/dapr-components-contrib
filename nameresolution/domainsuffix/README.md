# Domain Suffix Name Resolution

The Domain Suffix name resolver provides a simple way to resolve service names by appending a configurable domain suffix to the app name. This is useful in scenarios where you want to map service names to predictable DNS names.

## Configuration Format

To use the Domain Suffix name resolver, create a configuration in your Dapr environment:

```yaml
apiVersion: dapr.io/v1alpha1
kind: Configuration
metadata:
  name: appconfig
spec:
  nameResolution:
    component: "domainsuffix"
    configuration:
      domainSuffix: ".example.dev"  # Replace with your desired domain suffix
```

## Configuration Fields

| Field        | Required | Details                                                                                   | Example         |
|--------------|----------|-------------------------------------------------------------------------------------------|-----------------|
| domainSuffix | Y        | The domain suffix to append to service names. Should include a leading dot | ".example.dev" |

## Example

When configured with `domainSuffix: ".example.dev"`, the resolver will transform service names as follows:

- Service ID "myapp" → "myapp.example.dev"
- Service ID "another" → "another.example.dev"

## Notes

- Empty service IDs are not allowed and will result in an error
- The domain suffix must be provided in the configuration
- Leading dots in the domain suffix are expected