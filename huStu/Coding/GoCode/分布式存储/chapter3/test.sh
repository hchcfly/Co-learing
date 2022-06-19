#!/bin/bash

curl -XPUT localhost:9200/metadata -H 'content-Type:application/json'  -d'{"mappings":{"objects":
{"properties":{"name":{"type":"text"},"version":
{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"text"}}}}}'

curl -XPUT localhost:9200/metadata -H 'content-Type:application/json'  -d'{"mappings":{"objects":                                                                                                                   {"properties":{"name":{"type":"text"},"version":
{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"text"}}}}}'