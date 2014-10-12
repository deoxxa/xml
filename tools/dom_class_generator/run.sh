#!/bin/sh

# must have github.com/pointlander/peg installed

peg parser.peg && go build && ./dom_class_generator doml3.idl
