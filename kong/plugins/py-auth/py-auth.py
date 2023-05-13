#!/usr/bin/env python3
import kong_pdk.pdk.kong as kong

Schema = (
)

version = '0.1.0'
priority = 0

class Plugin(object):
    def __init__(self, config):
        self.config = config

    def access(self, kong: kong.kong):
        pass