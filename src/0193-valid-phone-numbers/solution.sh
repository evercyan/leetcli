#!/bin/sh

# [(xxx) ]|[xxx-]xxx-xxxx

awk '/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/' file.txt
