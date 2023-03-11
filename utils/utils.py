#!/usr/bin/python
# -*- coding: utf-8 -*-

import os


def read_file(filename):
    with open(filename, "r", encoding="utf-8") as fh:
        return fh.read()


def write_file(filename, s):
    with open(filename, "w", encoding="utf-8") as fh:
        fh.write(s)


def popen(cmd):
    r = os.popen(cmd)
    text = r.read()
    r.close()
    return text


def in_running(pattern):
    cmd_ps = "ps -ef|grep '{}'".format(pattern) + "|grep -v grep|awk '{print $2}'"
    if popen(cmd_ps).strip():
        return True
    return False


if __name__ == '__main__':
    pass
