---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dns_auth_nsgs Data Source - terraform-provider-bloxone"
subcategory: "DNS"
description: |-
  Retrieves information about existing Authoritative DNS Server Groups.
---

# bloxone_dns_auth_nsgs (Data Source)

Retrieves information about existing Authoritative DNS Server Groups.

## Example Usage

```terraform
# Get Auth NSGs filtered by an attribute
data "bloxone_dns_auth_nsgs" "example_by_attribute" {
  filters = {
    "name" = "example_auth_nsg"
  }
}

# Get Auth NSGs filtered by tag
data "bloxone_dns_auth_nsgs" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

# Get all Auth NSGs
data "bloxone_dns_auth_nsgs" "example_all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `tag_filters` (Map of String) Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Required:

- `name` (String) Name of the object.

Optional:

- `comment` (String) Optional. Comment for the object.
- `external_primaries` (Attributes List) Optional. DNS primaries external to BloxOne DDI. Order is not significant. (see [below for nested schema](#nestedatt--results--external_primaries))
- `external_secondaries` (Attributes List) DNS secondaries external to BloxOne DDI. Order is not significant. (see [below for nested schema](#nestedatt--results--external_secondaries))
- `internal_secondaries` (Attributes List) Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant. (see [below for nested schema](#nestedatt--results--internal_secondaries))
- `nsgs` (List of String) The resource identifier.
- `tags` (Map of String) Tagging specifics.

Read-Only:

- `id` (String) The resource identifier.

<a id="nestedatt--results--external_primaries"></a>
### Nested Schema for `results.external_primaries`

Required:

- `type` (String) Allowed values: * _nsg_, * _primary_.

Optional:

- `address` (String) Optional. Required only if _type_ is _server_. IP Address of nameserver.
- `fqdn` (String) Optional. Required only if _type_ is _server_. FQDN of nameserver.
- `nsg` (String) The resource identifier.
- `tsig_enabled` (Boolean) Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--results--external_primaries--tsig_key))

Read-Only:

- `protocol_fqdn` (String) FQDN of nameserver in punycode.

<a id="nestedatt--results--external_primaries--tsig_key"></a>
### Nested Schema for `results.external_primaries.tsig_key`

Required:

- `key` (String) The resource identifier.

Read-Only:

- `algorithm` (String) TSIG key algorithm.

  Possible values:
  * _hmac_sha256_
  * _hmac_sha1_
  * _hmac_sha224_
  * _hmac_sha384_
  * _hmac_sha512_
- `comment` (String) Comment for TSIG key.
- `name` (String) TSIG key name, FQDN.
- `protocol_name` (String) TSIG key name in punycode.
- `secret` (String, Sensitive) TSIG key secret, base64 string.



<a id="nestedatt--results--external_secondaries"></a>
### Nested Schema for `results.external_secondaries`

Required:

- `address` (String) IP Address of nameserver.
- `fqdn` (String) FQDN of nameserver.

Optional:

- `stealth` (Boolean) If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_
- `tsig_enabled` (Boolean) If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.  Default: _false_
- `tsig_key` (Attributes) (see [below for nested schema](#nestedatt--results--external_secondaries--tsig_key))

Read-Only:

- `protocol_fqdn` (String) FQDN of nameserver in punycode.

<a id="nestedatt--results--external_secondaries--tsig_key"></a>
### Nested Schema for `results.external_secondaries.tsig_key`

Required:

- `key` (String) The resource identifier.

Read-Only:

- `algorithm` (String) TSIG key algorithm.

  Possible values:
  * _hmac_sha256_
  * _hmac_sha1_
  * _hmac_sha224_
  * _hmac_sha384_
  * _hmac_sha512_
- `comment` (String) Comment for TSIG key.
- `name` (String) TSIG key name, FQDN.
- `protocol_name` (String) TSIG key name in punycode.
- `secret` (String, Sensitive) TSIG key secret, base64 string.



<a id="nestedatt--results--internal_secondaries"></a>
### Nested Schema for `results.internal_secondaries`

Required:

- `host` (String) The resource identifier.
